package txpool

import (
	"container/heap"
	"sync"
	"sync/atomic"

	"github.com/dogechain-lab/jury/types"
)

// txPricedList is a price-sorted heap to allow operating on transactions pool
// contents in a price-incrementing way. It's built opon the all transactions
// in txpool. All transactions will be considered for tracking, sorting, eviction, etc.
//
// Two heaps are used for sorting: the urgent heap (based on effective tip in the next
// block) and the floating heap (based on gasFeeCap). Always the bigger heap is chosen for
// eviction. Transactions evicted from the urgent heap are first demoted into the floating heap.
// In some cases (during a congestion, when blocks are full) the urgent heap can provide
// better candidates for inclusion while in other cases (at the top of the baseFee peak)
// the floating heap is better. When baseFee is decreasing they behave similarly.
type txPricedList struct {
	// Number of stale price points to (re-heap trigger).
	// This field is accessed atomically, and must be the first field
	// to ensure it has correct alignment for atomic.AddInt64.
	// See https://golang.org/pkg/sync/atomic/#pkg-note-BUG.
	stales int64

	all              *lookupMap   // Pointer to the map of all transactions
	urgent, floating minPriceHeap // Heaps of prices of all the stored transactions
	reheapMu         sync.Mutex   // Mutex asserts that only one routine is reheaping the list
}

const (
	// urgentRatio : floatingRatio is the capacity ratio of the two queues
	urgentRatio   = 4
	floatingRatio = 1
)

// newTxPricedList creates a new price-sorted transaction heap.
func newTxPricedList(all *lookupMap) *txPricedList {
	return &txPricedList{
		all: all,
	}
}

// Put inserts a new transaction into the heap.
func (l *txPricedList) Put(tx *types.Transaction, local bool) {
	if local {
		return
	}
	// Insert every new transaction to the urgent heap first; Discard will balance the heaps
	heap.Push(&l.urgent, tx)
}

// Removed notifies the prices transaction list that an old transaction dropped
// from the pool. The list will just keep a counter of stale objects and update
// the heap if a large enough ratio of transactions go stale.
func (l *txPricedList) Removed(count int) {
	// Bump the stale counter, but exit if still too low (< 25%)
	stales := atomic.AddInt64(&l.stales, int64(count))
	if int(stales) <= (len(l.urgent.list)+len(l.floating.list))/4 {
		return
	}
	// Seems we've reached a critical number of stale transactions, reheap
	l.Reheap()
}

// Underpriced checks whether a transaction is cheaper than (or as cheap as) the
// lowest priced transaction currently being tracked.
func (l *txPricedList) Underpriced(tx *types.Transaction) bool {
	// Note: with two queues, being underpriced is defined as being worse than the worst item
	// in all non-empty queues if there is any. If both queues are empty then nothing is underpriced.
	return (l.underpricedFor(&l.urgent, tx) || len(l.urgent.list) == 0) &&
		(l.underpricedFor(&l.floating, tx) || len(l.floating.list) == 0) &&
		(len(l.urgent.list) != 0 || len(l.floating.list) != 0)
}

// underpricedFor checks whether a transaction is cheaper than (or as cheap as) the
// lowest priced transaction in the given heap.
func (l *txPricedList) underpricedFor(h *minPriceHeap, tx *types.Transaction) bool {
	// Discard stale price points if found at the heap start
	for h.Len() > 0 {
		head := h.list[0]
		if _, exists := l.all.get(head.Hash); !exists {
			atomic.AddInt64(&l.stales, -1)
			heap.Pop(h)

			continue
		}

		break
	}
	// Check if the transaction is underpriced or not
	if len(h.list) == 0 {
		return false // There is no transaction at all.
	}
	// If the remote transaction is even cheaper than the
	// cheapest one tracked locally, reject it.
	return h.cmp(h.list[0], tx) >= 0
}

// Discard finds a number of most underpriced transactions, removes them from the
// priced list and returns them for further removal from the entire pool.
//
// Note local transaction won't be considered for eviction.
func (l *txPricedList) Discard(slots int, force bool) (types.Transactions, bool) {
	drop := make(types.Transactions, 0, slots) // Remote underpriced transactions to drop

	for slots > 0 {
		if len(l.urgent.list)*floatingRatio > len(l.floating.list)*urgentRatio || floatingRatio == 0 {
			// Discard stale transactions if found during cleanup
			// nolint: forcetypeassert
			tx := heap.Pop(&l.urgent).(*types.Transaction)
			if _, exists := l.all.get(tx.Hash); !exists { // Removed or migrated
				atomic.AddInt64(&l.stales, -1)

				continue
			}
			// Non stale transaction found, move to floating heap
			heap.Push(&l.floating, tx)
		} else {
			if len(l.floating.list) == 0 {
				// Stop if both heaps are empty
				break
			}
			// Discard stale transactions if found during cleanup
			// nolint: forcetypeassert
			tx := heap.Pop(&l.floating).(*types.Transaction)
			if _, exists := l.all.get(tx.Hash); !exists { // Removed or migrated
				atomic.AddInt64(&l.stales, -1)

				continue
			}
			// Non stale transaction found, discard it
			drop = append(drop, tx)
			slots -= int(slotsRequired(tx))
		}
	}
	// If we still can't make enough room for the new transaction
	if slots > 0 && !force {
		for _, tx := range drop {
			heap.Push(&l.urgent, tx)
		}

		return nil, false
	}

	return drop, true
}

// Reheap forcibly rebuilds the heap based on the current remote transaction set.
func (l *txPricedList) Reheap() {
	l.reheapMu.Lock()
	defer l.reheapMu.Unlock()

	atomic.StoreInt64(&l.stales, 0)
	l.urgent.list = make(types.Transactions, 0, l.all.Len())
	l.all.Range(func(hash types.Hash, tx *types.Transaction) bool {
		l.urgent.list = append(l.urgent.list, tx)

		return true
	})
	heap.Init(&l.urgent)

	// balance out the two heaps by moving the worse half of transactions into the
	// floating heap
	// Note: Discard would also do this before the first eviction but Reheap can do
	// is more efficiently. Also, Underpriced would work suboptimally the first time
	// if the floating queue was empty.
	floatingCount := len(l.urgent.list) * floatingRatio / (urgentRatio + floatingRatio)
	l.floating.list = make([]*types.Transaction, floatingCount)

	for i := 0; i < floatingCount; i++ {
		// nolint: forcetypeassert
		l.floating.list[i] = heap.Pop(&l.urgent).(*types.Transaction)
	}

	heap.Init(&l.floating)
}
