# CKB SDK Golang

[![License](https://img.shields.io/badge/license-MIT-green)](https://github.com/nervosnetwork/ckb-sdk-go/blob/master/LICENSE)
[![Go version](https://img.shields.io/badge/go-1.11.5-blue.svg)](https://github.com/moovweb/gvm)
[![Telegram Group](https://cdn.rawgit.com/Patrolavia/telegram-badge/8fe3382b/chat.svg)](https://t.me/nervos_ckb_dev)

Golang SDK for Nervos [CKB](https://github.com/nervosnetwork/ckb).

The ckb-sdk-go is still under development and **NOT** production ready. You should get familiar with CKB transaction
structure and RPC before using it.

## WARNING

Module Indexer has been removed from [ckb_v0.40.0](https://github.com/nervosnetwork/ckb/releases/tag/v0.40.0): Please
use [ckb-indexer](https://github.com/nervosnetwork/ckb-indexer) as an alternate solution.

The following RPCs hash been removed from [ckb_v0.40.0](https://github.com/nervosnetwork/ckb/releases/tag/v0.40.0):

* `get_live_cells_by_lock_hash`
* `get_transactions_by_lock_hash`
* `index_lock_hash`
* `deindex_lock_hash`
* `get_lock_hash_index_states`
* `get_capacity_by_lock_hash`

Since [ckb_v0.36.0](https://github.com/nervosnetwork/ckb/releases/tag/v0.36.0) SDK
use [ckb-indexer](https://github.com/nervosnetwork/ckb-indexer) to collect cells, please
see [Collect cells](#5-collect-cells) for examples.

## Installation

### Minimum requirements

| Components | Version | Description |
|----------|-------------|-------------|
| [Golang](https://golang.org) | &ge; 1.11.5 | Go programming language |

### Install

```bash
go get -v github.com/nervosnetwork/ckb-sdk-go
```

## Quick start

### Sign and send transaction

```go
package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"log"

	"github.com/nervosnetwork/ckb-sdk-go/crypto/secp256k1"
	"github.com/nervosnetwork/ckb-sdk-go/rpc"
	"github.com/nervosnetwork/ckb-sdk-go/transaction"
	"github.com/nervosnetwork/ckb-sdk-go/types"
	"github.com/nervosnetwork/ckb-sdk-go/utils"
)

func main() {
	client, err := rpc.Dial("http://127.0.0.1:8114")
	if err != nil {
		log.Fatalf("create rpc client error: %v", err)
	}

	key, err := secp256k1.HexToKey(PRIVATE_KEY)
	if err != nil {
		log.Fatalf("import private key error: %v", err)
	}

	systemScripts, err := utils.NewSystemScripts(client)
	if err != nil {
		log.Fatalf("load system script error: %v", err)
	}

	toAddress, _ := hex.DecodeString("bf3e92da4911fa5f620e7b1fd27c2d0ddd0de744")
	changeScript, _ := key.Script(systemScripts)

	tx := transaction.NewSecp256k1SingleSigTx(systemScripts)
	tx.Outputs = append(tx.Outputs, &types.CellOutput{
		Capacity: 200000000000,
		Lock: &types.Script{
			CodeHash: types.HexToHash("0x9bd7e06f3ecf4be0f2fcd2188b23f1b9fcc88e5d4b65a8637b17723bbda3cce8"),
			HashType: types.HashTypeType,
			Args:     toAddress,
		},
	})
	tx.Outputs = append(tx.Outputs, &types.CellOutput{
		Capacity: 199999998000,
		Lock: &types.Script{
			CodeHash: types.HexToHash("0x9bd7e06f3ecf4be0f2fcd2188b23f1b9fcc88e5d4b65a8637b17723bbda3cce8"),
			HashType: types.HashTypeType,
			Args:     changeScript.Args,
		},
	})
	tx.OutputsData = [][]byte{{}, {}}

	group, witnessArgs, err := transaction.AddInputsForTransaction(tx, []*types.CellInput{
		{
			Since: 0,
			PreviousOutput: &types.OutPoint{
				TxHash: types.HexToHash("0x8e6d818c6e07e6cbd9fca51294030494ee23dc388d7f5276ba50b938d02cc015"),
				Index:  1,
			},
		},
	})

	if err != nil {
		log.Fatalf("add inputs to transaction error: %v", err)
	}

	err = transaction.SingleSignTransaction(tx, group, witnessArgs, key)
	if err != nil {
		log.Fatalf("sign transaction error: %v", err)
	}

	hash, err := client.SendTransaction(context.Background(), tx)
	if err != nil {
		log.Fatalf("send transaction error: %v", err)
	}

	fmt.Println(hash.String())
}
```

### Create a new address

In CKB world, a lock script can be represented as an address. `secp256k1_blake160` is the most common used address and
here we show how to generate it.

```go
addressGenerateResult, error := GenerateAddress(Testnet)
```

For more details please about CKB address refer
to [CKB rfc 0021](https://github.com/nervosnetwork/rfcs/blob/master/rfcs/0021-ckb-address-format/0021-ckb-address-format.md)
.

### Convert public key to address

Convert elliptic curve public key to an address (`secp256k1_blake160`)

```go
address, err := ConvertPublicToAddress(Mainnet, "0xb39bbc0b3673c7d36450bc14cfcdad2d559c6c64")
```

### Convert short/bech32 address to bech32m address

Short address and bech32 address are deprecated. The standard address format is bech32m-encoded long address, which can
be got from the short address or bech32 address as the following snippet code.

```go
bech32mFullAddress, err := ConvertDeprecatedAddressToBech32mFullAddress("ckt1qyqxgp7za7dajm5wzjkye52asc8fxvvqy9eqlhp82g")
```

### Parse and validate address

```go
parsedAddress, err := Parse("ckt1qg8mxsu48mncexvxkzgaa7mz2g25uza4zpz062relhjmyuc52ps3zn47dugwyk5e6mgxvlf5ukx7k3uyq9wlkkmegke")
```

### Mercury

[Mercury](https://github.com/nervosnetwork/mercury) is a development service in CKB ecosystem, providing many
useful [RPC APIs](https://github.com/nervosnetwork/mercury/blob/main/core/rpc/README.md) for development like querying
transaction or getting udt asset information. You need to deploy your own mercury and sync data with the network before
using it.

ckb-sdk-go also integrate with Mercury. For usage guide, please check the [example folder](./mercury/example).

## License

The SDK is available as open source under the terms of the [MIT License](https://opensource.org/licenses/MIT).