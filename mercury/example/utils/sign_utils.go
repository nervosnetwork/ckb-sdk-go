package utils

import (
	"github.com/nervosnetwork/ckb-sdk-go/crypto/secp256k1"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/example/constant"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/resp"
	"github.com/nervosnetwork/ckb-sdk-go/transaction"
	"github.com/nervosnetwork/ckb-sdk-go/types"
)

func Sign(transferCompletion *resp.TransferCompletionResponse) *types.Transaction {
	scriptGroups := transferCompletion.GetScriptGroup()
	tx := transferCompletion.GetTransaction()
	for _, group := range scriptGroups {
		key, _ := secp256k1.HexToKey(constant.GetKey(group.GetAddress()))
		transaction.SignTransaction(tx, group, key)
	}
	return tx
}
