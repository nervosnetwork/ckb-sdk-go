# CKB SDK Golang

[![License](https://img.shields.io/badge/license-MIT-green)](https://github.com/nervosnetwork/ckb-sdk-go/blob/master/LICENSE)
[![Go version](https://img.shields.io/badge/go-1.11.5-blue.svg)](https://github.com/moovweb/gvm)
[![Telegram Group](https://cdn.rawgit.com/Patrolavia/telegram-badge/8fe3382b/chat.svg)](https://t.me/nervos_ckb_dev)

Golang SDK for Nervos [CKB](https://github.com/nervosnetwork/ckb).

The ckb-sdk-go is still under development and considered to be a work in progress. You should get familiar with CKB transaction
structure and RPC before using it.

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

### Setup

ckb-sdk-go provides a convenient client to help you easily interact with [CKB](https://github.com/nervosnetwork/ckb), [CKB-indexer](https://github.com/nervosnetwork/ckb-indexer) or [Mercury](https://github.com/nervosnetwork/mercury) node.

```go
ckbClient, err := rpc.Dial("http://127.0.0.1:8114")
indexerClient, err := indexer.Dial("http://127.0.0.1:8114")
mercuryClient , err := mercury.Dial("http://127.0.0.1:8116")
```

You can call JSON-RPC APIs via these clients.

```go
block, err := ckbClient.GetBlock(context.Background(), types.HexToHash("0x77fdd22f6ae8a717de9ae2b128834e9b2a1424378b5fc95606ba017aab5fed75"))
```

For more details about JSON-RPC APIs, please check:

- [CKB RPC doc](https://github.com/nervosnetwork/ckb/blob/develop/rpc/README.md)
- [CKB-indexer RPC doc](https://github.com/nervosnetwork/ckb-indexer/blob/master/README.md)
- [Mercury RPC doc](https://github.com/nervosnetwork/mercury/blob/main/core/rpc/README.md).

### Build transaction with Mercury

[Mercury](https://github.com/nervosnetwork/mercury) is an application for better interaction with CKB chain, providing many useful [JSON-RPC APIs](https://github.com/nervosnetwork/mercury/blob/main/core/rpc/README.md) for development like querying transactions or getting UDT asset information. You need to deploy your own mercury server and sync data with the latest network before using it.

Mercury is another way to build transaction. With the help of Mercury, you can build a transaction by simply calling a JSON-RPC API. Here we show how to build a CKB transfer transaction with mercury.

```go
sender := "ckt1qzda0cr08m85hc8jlnfp3zer7xulejywt49kt2rr0vthywaa50xwsq0yvcdtsu5wcr2jldtl72fhkruf0w5vymsp6rk9r"
receiver := "ckt1qzda0cr08m85hc8jlnfp3zer7xulejywt49kt2rr0vthywaa50xwsqvglkprurm00l7hrs3rfqmmzyy3ll7djdsujdm6z"
ckbAmount := amount.CkbToShannon(100)  // Convert CKB to Shannon (1 CKB = 10^8 Shannon)
req := &model.SimpleTransferPayload{
	AssetInfo: model.NewCkbAsset(),
	From:       []string{sender},
	To:          []*model.ToInfo{{receiver, ckbAmount}},
	FeeRate:   1000,
}
// Get an unsigned raw transaction with the help of Mercury
resp, err := mercuryClient.BuildSimpleTransferTransaction(req)
```

For more use cases of Mercury, please refer to [Mercury test cases](./mercury/client_test.go) and [Mercury JSON-RPC documentation](https://github.com/nervosnetwork/mercury/blob/dev-0.4/core/rpc/README.md).

### Sign and send transaction

To send the built transaction to CKB network, you need to

1. sign transaction with your private key.
2. send signed transaction to CKB node, and wait it to be confirmed.

```go
var txWithScriptGroup *transaction.TransactionWithScriptGroups
// 0. Set your private key
privKey := "0xccb083b37aa346c5ce2e1f99a687a153baa04052f26db6ab3c26d6a4cc15c5f1"
// 1. Sign transaction with your private key
txSigner := signer.GetTransactionSignerInstance(types.NetworkTest)
txSigner.SignTransactionByPrivateKeys(txWithScriptGroup, privKey)
// 2. Send transaction to CKB node
txHash, err := ckbClient.SendTransaction(context.Background(), txWithScriptGroup.TxView)
```

Please note that before signing and sending transaction, you need to prepare a raw transaction represented by an instance of struct `TransactionWithScriptGroups`. You can get it [by Mercury](#Build-transaction-with-Mercury) or by ckb-indexer.

### Generate a new address
In CKB world, a lock script can be represented as an address. `secp256k1_blake160_signhash_all` is the most common used address and here we show how to generate it.

```go
// Generate a new address randomly
key, err := secp256k1.RandomNew()
if err != nil {
	// handle error
}
script := address.GenerateScriptSecp256K1Blake160SignhashAll(key)
addr := &address.Address{Script: script, Network: types.NetworkTest}
encodedAddr, err := addr.Encode()
```

For more details please about CKB address refer to [CKB rfc 0021](https://github.com/nervosnetwork/rfcs/blob/master/rfcs/0021-ckb-address-format/0021-ckb-address-format.md).

### Convert public key to address

Convert elliptic curve public key to an address (`secp256k1_blake160_signash_all`)

```go
// You should provide an elliptic curve public key of compressed format, with 33 bytes.
script, err := address.GenerateScriptSecp256K1Blake160SignhashAllByPublicKey("0x03a0a7a7597b019828a1dda6ed52ab25181073ec3a9825d28b9abbb932fe1ec83d")
if err != nil {
	// handle error
}
addr := &address.Address{Script: script, Network: types.NetworkTest}
```

### Parse address

Short address and full bech32 address are deprecated. The standard address encoded way is bech32m. You can still parse address
from an encoded string address and then get its network, script and encoded string of other format.

```go
addr, err := address.Decode("ckt1qyqxgp7za7dajm5wzjkye52asc8fxvvqy9eqlhp82g")
if err != nil {
	// handle error
}
script := addr.Script
network := addr.Network
```

## License

The SDK is available as open source under the terms of the [MIT License](https://opensource.org/licenses/MIT).