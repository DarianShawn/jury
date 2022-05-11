package txpool

import (
	"sync"

	"github.com/dogechain-lab/jury/types"
)

// Lookup map used to find transactions present in the pool
type lookupMap struct {
	sync.RWMutex
	all map[types.Hash]*types.Transaction
}

// add inserts the given transaction into the map. [thread-safe]
func (m *lookupMap) add(txs ...*types.Transaction) {
	m.Lock()
	defer m.Unlock()

	for _, tx := range txs {
		m.all[tx.Hash] = tx
	}
}

// remove removes the given transactions from the map. [thread-safe]
func (m *lookupMap) remove(txs ...*types.Transaction) {
	m.Lock()
	defer m.Unlock()

	for _, tx := range txs {
		delete(m.all, tx.Hash)
	}
}

// get returns the transaction associated with the given hash. [thread-safe]
func (m *lookupMap) get(hash types.Hash) (*types.Transaction, bool) {
	m.RLock()
	defer m.RUnlock()

	tx, ok := m.all[hash]
	if !ok {
		return nil, false
	}

	return tx, true
}

// Len returns the transaction length. [thread-safe]
func (m *lookupMap) Len() int {
	m.RLock()
	defer m.RUnlock()

	return len(m.all)
}

// Range calls f on each key and value present in the map. The callback passed
// should return the indicator whether the iteration needs to be continued.
// Callers need to specify which set (or both) to be iterated.
func (m *lookupMap) Range(f func(hash types.Hash, tx *types.Transaction) bool) {
	m.RLock()
	defer m.RUnlock()

	for key, value := range m.all {
		if !f(key, value) {
			return
		}
	}
}
