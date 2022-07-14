We roll out a brand new ckb-sdk-java after refactor work. The new sdk brings plenty of BREAKING CHANGES incompatible with the deprecated sdk `v1.0.*` and earlier releases.

The breaking changes include

- Type or name change of quite a few fields in RPC type representation.
- Unified address representation and operation.
- Transaction signing mechanism by `ScriptGroup`, `ScriptSigner`, and `TransactionSigner`.
- Clean some utils classes and unused classes.

Other underlying breaking changes that are possibly transparent to you

- More robust test.

In the following part, we list the aspects you need to care about most if you are ready to migrate from deprecated sdk to the new one.

## Address

The new sdk remove struct `ParsedAddress` and `AddressGenerateResult`, and use a unified struct `Address` to represent ckb address. Below code is the example of address operations.

Decode address

```go
encoded := "ckt1qzda..."
// Before
parsedAddr, err := address.Parse(encoded)
script := parsedAddr.Script
network := parsedAddr.Mode

// Now:
addr, err := address.Decode(encoded)
script := addr.Script
network := addr.Network
```

Encode address from script and network

```go
// Before:
encoded, err := address.ConvertScriptToAddress(address.Testnet, script)

// Now:
addr := &address.Address{Script: script, Network: types.NetworkTest}
encoded := addr.Encode()
```

## Packages migration

Some important packages moves

- `github.com/nervosnetwork/ckb-sdk-go/mercury/model/*` -> `github.com/nervosnetwork/ckb-sdk-go/mercury/model`

```go
// Before:
import (
    "github.com/nervosnetwork/ckb-sdk-go/mercury/model/common"
    "github.com/nervosnetwork/ckb-sdk-go/mercury/model/resp"
)

// Now:
import (
    "github.com/nervosnetwork/ckb-sdk-go/mercury/model"
)
```

## Sign transaction

The new go sdk introduces `ScriptGroup` for signing transactions. You can get the `ScriptGroup` with a raw transaction from mercury, or construct it by yourself. After that `TransactionSigner` can sign in the correct place as long as you provide the right secret information (e.g. private key).

```go
// Before:
addressWithKeys := make(map[string]string)
// omit the code to put private key to `addressWithKeys`
txWithGroup, err := mercuryClient.BuildSimpleTransferTransaction(req)
scriptGroups := txWithGroup.GetScriptGroup()
tx := txWithGroup.GetTransaction()
for _, group := range scriptGroups {
    key, _ := secp256k1.HexToKey(addressWithKeys[group.GetAddress()])
    resp.SignTransaction(tx, group, key)
}
txHash, err := ckbClient.SendTransaction(context.Background(), txWithScriptGroup.TxView)

// Now:
txWithScriptGroups, err := mercuryClient.BuildSimpleTransferTransaction(req)
txSigner := signer.GetTransactionSignerInstance(types.NetworkTest)
txSigner.SignTransaction(txWithScriptGroup, "0x6c9ed03816e31...")
txHash, err := ckbClient.SendTransaction(context.Background(), txWithScriptGroup.TxView)
```