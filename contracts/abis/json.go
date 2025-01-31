package abis

const ValidatorSetJSONABI = `[
    {
        "inputs":
        [],
        "stateMutability": "nonpayable",
        "type": "constructor"
    },
    {
        "anonymous": false,
        "inputs":
        [
            {
                "indexed": true,
                "internalType": "address",
                "name": "sender",
                "type": "address"
            },
            {
                "indexed": true,
                "internalType": "address",
                "name": "oldToken",
                "type": "address"
            },
            {
                "indexed": true,
                "internalType": "address",
                "name": "newToken",
                "type": "address"
            }
        ],
        "name": "GovernorTokenSet",
        "type": "event"
    },
    {
        "anonymous": false,
        "inputs":
        [
            {
                "indexed": true,
                "internalType": "address",
                "name": "sender",
                "type": "address"
            },
            {
                "indexed": true,
                "internalType": "uint256",
                "name": "oldMinimum",
                "type": "uint256"
            },
            {
                "indexed": true,
                "internalType": "uint256",
                "name": "newMinimum",
                "type": "uint256"
            }
        ],
        "name": "MinimumSet",
        "type": "event"
    },
    {
        "anonymous": false,
        "inputs":
        [
            {
                "indexed": true,
                "internalType": "address",
                "name": "previousOwner",
                "type": "address"
            },
            {
                "indexed": true,
                "internalType": "address",
                "name": "newOwner",
                "type": "address"
            }
        ],
        "name": "OwnershipTransferred",
        "type": "event"
    },
    {
        "anonymous": false,
        "inputs":
        [
            {
                "indexed": true,
                "internalType": "address",
                "name": "sender",
                "type": "address"
            },
            {
                "indexed": true,
                "internalType": "uint256",
                "name": "amount",
                "type": "uint256"
            }
        ],
        "name": "Staked",
        "type": "event"
    },
    {
        "anonymous": false,
        "inputs":
        [
            {
                "indexed": true,
                "internalType": "address",
                "name": "sender",
                "type": "address"
            },
            {
                "indexed": true,
                "internalType": "uint256",
                "name": "oldThreshold",
                "type": "uint256"
            },
            {
                "indexed": true,
                "internalType": "uint256",
                "name": "newThreshold",
                "type": "uint256"
            }
        ],
        "name": "ThresholdSet",
        "type": "event"
    },
    {
        "anonymous": false,
        "inputs":
        [
            {
                "indexed": true,
                "internalType": "address",
                "name": "sender",
                "type": "address"
            },
            {
                "indexed": true,
                "internalType": "uint256",
                "name": "amount",
                "type": "uint256"
            }
        ],
        "name": "Unstaked",
        "type": "event"
    },
    {
        "anonymous": false,
        "inputs":
        [
            {
                "indexed": true,
                "internalType": "address",
                "name": "sender",
                "type": "address"
            },
            {
                "indexed": true,
                "internalType": "address",
                "name": "account",
                "type": "address"
            }
        ],
        "name": "ValidatorAdded",
        "type": "event"
    },
    {
        "anonymous": false,
        "inputs":
        [
            {
                "indexed": true,
                "internalType": "address",
                "name": "sender",
                "type": "address"
            },
            {
                "indexed": true,
                "internalType": "address",
                "name": "account",
                "type": "address"
            }
        ],
        "name": "ValidatorDeleted",
        "type": "event"
    },
    {
        "inputs":
        [
            {
                "internalType": "uint256",
                "name": "",
                "type": "uint256"
            }
        ],
        "name": "_validators",
        "outputs":
        [
            {
                "internalType": "address",
                "name": "",
                "type": "address"
            }
        ],
        "stateMutability": "view",
        "type": "function"
    },
    {
        "inputs":
        [
            {
                "internalType": "address",
                "name": "account",
                "type": "address"
            }
        ],
        "name": "accountStake",
        "outputs":
        [
            {
                "internalType": "uint256",
                "name": "",
                "type": "uint256"
            }
        ],
        "stateMutability": "view",
        "type": "function"
    },
    {
        "inputs":
        [
            {
                "internalType": "address",
                "name": "account",
                "type": "address"
            }
        ],
        "name": "addValidator",
        "outputs":
        [],
        "stateMutability": "nonpayable",
        "type": "function"
    },
    {
        "inputs":
        [
            {
                "internalType": "address",
                "name": "account",
                "type": "address"
            }
        ],
        "name": "deleteValidator",
        "outputs":
        [],
        "stateMutability": "nonpayable",
        "type": "function"
    },
    {
        "inputs":
        [],
        "name": "governor",
        "outputs":
        [
            {
                "internalType": "address",
                "name": "",
                "type": "address"
            }
        ],
        "stateMutability": "view",
        "type": "function"
    },
    {
        "inputs":
        [
            {
                "internalType": "address",
                "name": "account",
                "type": "address"
            }
        ],
        "name": "isValidator",
        "outputs":
        [
            {
                "internalType": "bool",
                "name": "",
                "type": "bool"
            }
        ],
        "stateMutability": "view",
        "type": "function"
    },
    {
        "inputs":
        [],
        "name": "minimum",
        "outputs":
        [
            {
                "internalType": "uint256",
                "name": "",
                "type": "uint256"
            }
        ],
        "stateMutability": "view",
        "type": "function"
    },
    {
        "inputs":
        [],
        "name": "owner",
        "outputs":
        [
            {
                "internalType": "address",
                "name": "",
                "type": "address"
            }
        ],
        "stateMutability": "view",
        "type": "function"
    },
    {
        "inputs":
        [],
        "name": "renounceOwnership",
        "outputs":
        [],
        "stateMutability": "nonpayable",
        "type": "function"
    },
    {
        "inputs":
        [
            {
                "internalType": "address",
                "name": "token",
                "type": "address"
            }
        ],
        "name": "setGovernorToken",
        "outputs":
        [],
        "stateMutability": "nonpayable",
        "type": "function"
    },
    {
        "inputs":
        [
            {
                "internalType": "uint256",
                "name": "number",
                "type": "uint256"
            }
        ],
        "name": "setMinimum",
        "outputs":
        [],
        "stateMutability": "nonpayable",
        "type": "function"
    },
    {
        "inputs":
        [
            {
                "internalType": "uint256",
                "name": "amount",
                "type": "uint256"
            }
        ],
        "name": "setThreshold",
        "outputs":
        [],
        "stateMutability": "nonpayable",
        "type": "function"
    },
    {
        "inputs":
        [
            {
                "internalType": "uint256",
                "name": "amount",
                "type": "uint256"
            }
        ],
        "name": "stake",
        "outputs":
        [],
        "stateMutability": "nonpayable",
        "type": "function"
    },
    {
        "inputs":
        [],
        "name": "stakedAmount",
        "outputs":
        [
            {
                "internalType": "uint256",
                "name": "",
                "type": "uint256"
            }
        ],
        "stateMutability": "view",
        "type": "function"
    },
    {
        "inputs":
        [],
        "name": "threshold",
        "outputs":
        [
            {
                "internalType": "uint256",
                "name": "",
                "type": "uint256"
            }
        ],
        "stateMutability": "view",
        "type": "function"
    },
    {
        "inputs":
        [
            {
                "internalType": "address",
                "name": "newOwner",
                "type": "address"
            }
        ],
        "name": "transferOwnership",
        "outputs":
        [],
        "stateMutability": "nonpayable",
        "type": "function"
    },
    {
        "inputs":
        [],
        "name": "unstake",
        "outputs":
        [],
        "stateMutability": "nonpayable",
        "type": "function"
    },
    {
        "inputs":
        [],
        "name": "validators",
        "outputs":
        [
            {
                "internalType": "address[]",
                "name": "",
                "type": "address[]"
            }
        ],
        "stateMutability": "view",
        "type": "function"
    }
]`

const BridgeJSONABI = `[
    {
        "inputs":
        [],
        "stateMutability": "nonpayable",
        "type": "constructor"
    },
    {
        "anonymous": false,
        "inputs":
        [
            {
                "indexed": true,
                "internalType": "address",
                "name": "receiver",
                "type": "address"
            },
            {
                "indexed": true,
                "internalType": "uint256",
                "name": "amount",
                "type": "uint256"
            },
            {
                "indexed": false,
                "internalType": "string",
                "name": "txid",
                "type": "string"
            },
            {
                "indexed": false,
                "internalType": "string",
                "name": "sender",
                "type": "string"
            }
        ],
        "name": "Deposited",
        "type": "event"
    },
    {
        "anonymous": false,
        "inputs":
        [
            {
                "indexed": true,
                "internalType": "address",
                "name": "sender",
                "type": "address"
            },
            {
                "indexed": true,
                "internalType": "uint256",
                "name": "oldMinimumThreshold",
                "type": "uint256"
            },
            {
                "indexed": true,
                "internalType": "uint256",
                "name": "newMinimumThreshold",
                "type": "uint256"
            }
        ],
        "name": "MinimumThresholdSet",
        "type": "event"
    },
    {
        "anonymous": false,
        "inputs":
        [
            {
                "indexed": true,
                "internalType": "address",
                "name": "previousOwner",
                "type": "address"
            },
            {
                "indexed": true,
                "internalType": "address",
                "name": "newOwner",
                "type": "address"
            }
        ],
        "name": "OwnershipTransferred",
        "type": "event"
    },
    {
        "anonymous": false,
        "inputs":
        [
            {
                "indexed": true,
                "internalType": "address",
                "name": "sender",
                "type": "address"
            },
            {
                "indexed": true,
                "internalType": "uint256",
                "name": "oldRate",
                "type": "uint256"
            },
            {
                "indexed": true,
                "internalType": "uint256",
                "name": "newRate",
                "type": "uint256"
            }
        ],
        "name": "RateSet",
        "type": "event"
    },
    {
        "anonymous": false,
        "inputs":
        [
            {
                "indexed": true,
                "internalType": "address",
                "name": "sender",
                "type": "address"
            },
            {
                "indexed": true,
                "internalType": "address",
                "name": "signer",
                "type": "address"
            }
        ],
        "name": "SignerDeleted",
        "type": "event"
    },
    {
        "anonymous": false,
        "inputs":
        [
            {
                "indexed": true,
                "internalType": "address",
                "name": "sender",
                "type": "address"
            },
            {
                "indexed": true,
                "internalType": "address",
                "name": "signer",
                "type": "address"
            }
        ],
        "name": "ValidatorAdded",
        "type": "event"
    },
    {
        "anonymous": false,
        "inputs":
        [
            {
                "indexed": true,
                "internalType": "address",
                "name": "sender",
                "type": "address"
            },
            {
                "indexed": true,
                "internalType": "uint256",
                "name": "amount",
                "type": "uint256"
            },
            {
                "indexed": true,
                "internalType": "uint256",
                "name": "fee",
                "type": "uint256"
            },
            {
                "indexed": false,
                "internalType": "string",
                "name": "receiver",
                "type": "string"
            }
        ],
        "name": "Withdrawn",
        "type": "event"
    },
    {
        "inputs":
        [
            {
                "internalType": "address",
                "name": "account",
                "type": "address"
            }
        ],
        "name": "addSigner",
        "outputs":
        [],
        "stateMutability": "nonpayable",
        "type": "function"
    },
    {
        "inputs":
        [
            {
                "internalType": "address",
                "name": "account",
                "type": "address"
            }
        ],
        "name": "deleteSigner",
        "outputs":
        [],
        "stateMutability": "nonpayable",
        "type": "function"
    },
    {
        "inputs":
        [
            {
                "internalType": "address",
                "name": "receiver",
                "type": "address"
            },
            {
                "internalType": "uint256",
                "name": "amount",
                "type": "uint256"
            },
            {
                "internalType": "string",
                "name": "txid",
                "type": "string"
            },
            {
                "internalType": "string",
                "name": "sender",
                "type": "string"
            }
        ],
        "name": "deposit",
        "outputs":
        [],
        "stateMutability": "nonpayable",
        "type": "function"
    },
    {
        "inputs":
        [],
        "name": "getMinimumThreshold",
        "outputs":
        [
            {
                "internalType": "uint256",
                "name": "",
                "type": "uint256"
            }
        ],
        "stateMutability": "view",
        "type": "function"
    },
    {
        "inputs":
        [],
        "name": "getSigners",
        "outputs":
        [
            {
                "internalType": "address[]",
                "name": "",
                "type": "address[]"
            }
        ],
        "stateMutability": "view",
        "type": "function"
    },
    {
        "inputs":
        [
            {
                "internalType": "address",
                "name": "account",
                "type": "address"
            }
        ],
        "name": "isSigner",
        "outputs":
        [
            {
                "internalType": "bool",
                "name": "",
                "type": "bool"
            }
        ],
        "stateMutability": "view",
        "type": "function"
    },
    {
        "inputs":
        [],
        "name": "owner",
        "outputs":
        [
            {
                "internalType": "address",
                "name": "",
                "type": "address"
            }
        ],
        "stateMutability": "view",
        "type": "function"
    },
    {
        "inputs":
        [],
        "name": "rate",
        "outputs":
        [
            {
                "internalType": "uint256",
                "name": "",
                "type": "uint256"
            }
        ],
        "stateMutability": "view",
        "type": "function"
    },
    {
        "inputs":
        [],
        "name": "renounceOwnership",
        "outputs":
        [],
        "stateMutability": "nonpayable",
        "type": "function"
    },
    {
        "inputs":
        [
            {
                "internalType": "uint256",
                "name": "newMinimumThreshold",
                "type": "uint256"
            }
        ],
        "name": "setMinimumThreshold",
        "outputs":
        [],
        "stateMutability": "nonpayable",
        "type": "function"
    },
    {
        "inputs":
        [
            {
                "internalType": "uint256",
                "name": "newRate",
                "type": "uint256"
            }
        ],
        "name": "setRate",
        "outputs":
        [],
        "stateMutability": "nonpayable",
        "type": "function"
    },
    {
        "inputs":
        [
            {
                "internalType": "address",
                "name": "receiver",
                "type": "address"
            },
            {
                "internalType": "uint256",
                "name": "amount",
                "type": "uint256"
            },
            {
                "internalType": "string",
                "name": "txid",
                "type": "string"
            },
            {
                "internalType": "string",
                "name": "sender",
                "type": "string"
            }
        ],
        "name": "signatures",
        "outputs":
        [
            {
                "internalType": "address[]",
                "name": "",
                "type": "address[]"
            }
        ],
        "stateMutability": "view",
        "type": "function"
    },
    {
        "inputs":
        [],
        "name": "totalSupply",
        "outputs":
        [
            {
                "internalType": "uint256",
                "name": "",
                "type": "uint256"
            }
        ],
        "stateMutability": "view",
        "type": "function"
    },
    {
        "inputs":
        [
            {
                "internalType": "address",
                "name": "newOwner",
                "type": "address"
            }
        ],
        "name": "transferOwnership",
        "outputs":
        [],
        "stateMutability": "nonpayable",
        "type": "function"
    },
    {
        "inputs":
        [
            {
                "internalType": "string",
                "name": "receiver",
                "type": "string"
            }
        ],
        "name": "withdraw",
        "outputs":
        [],
        "stateMutability": "payable",
        "type": "function"
    }
]`

const VaultJSONABI = `[
    {
        "inputs":
        [],
        "stateMutability": "nonpayable",
        "type": "constructor"
    },
    {
        "anonymous": false,
        "inputs":
        [
            {
                "indexed": true,
                "internalType": "address",
                "name": "previousOwner",
                "type": "address"
            },
            {
                "indexed": true,
                "internalType": "address",
                "name": "newOwner",
                "type": "address"
            }
        ],
        "name": "OwnershipTransferred",
        "type": "event"
    },
    {
        "anonymous": false,
        "inputs":
        [
            {
                "indexed": true,
                "internalType": "address",
                "name": "to",
                "type": "address"
            },
            {
                "indexed": true,
                "internalType": "uint256",
                "name": "amount",
                "type": "uint256"
            }
        ],
        "name": "RewardTo",
        "type": "event"
    },
    {
        "anonymous": false,
        "inputs":
        [
            {
                "indexed": true,
                "internalType": "address",
                "name": "from",
                "type": "address"
            },
            {
                "indexed": true,
                "internalType": "uint256",
                "name": "amount",
                "type": "uint256"
            }
        ],
        "name": "receiveReward",
        "type": "event"
    },
    {
        "inputs":
        [],
        "name": "balance",
        "outputs":
        [
            {
                "internalType": "uint256",
                "name": "",
                "type": "uint256"
            }
        ],
        "stateMutability": "view",
        "type": "function"
    },
    {
        "inputs":
        [
            {
                "internalType": "address payable",
                "name": "to",
                "type": "address"
            },
            {
                "internalType": "uint256",
                "name": "amount",
                "type": "uint256"
            }
        ],
        "name": "claimRewards",
        "outputs":
        [],
        "stateMutability": "nonpayable",
        "type": "function"
    },
    {
        "inputs":
        [],
        "name": "owner",
        "outputs":
        [
            {
                "internalType": "address",
                "name": "",
                "type": "address"
            }
        ],
        "stateMutability": "view",
        "type": "function"
    },
    {
        "inputs":
        [],
        "name": "renounceOwnership",
        "outputs":
        [],
        "stateMutability": "nonpayable",
        "type": "function"
    },
    {
        "inputs":
        [
            {
                "internalType": "address",
                "name": "newOwner",
                "type": "address"
            }
        ],
        "name": "transferOwnership",
        "outputs":
        [],
        "stateMutability": "nonpayable",
        "type": "function"
    },
    {
        "stateMutability": "payable",
        "type": "receive"
    }
]`

const StressTestJSONABI = `[
    {
      "inputs": [],
      "stateMutability": "nonpayable",
      "type": "constructor"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": false,
          "internalType": "uint256",
          "name": "number",
          "type": "uint256"
        }
      ],
      "name": "txnDone",
      "type": "event"
    },
    {
      "inputs": [],
      "name": "getCount",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "string",
          "name": "sName",
          "type": "string"
        }
      ],
      "name": "setName",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    }
  ]`
