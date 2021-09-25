package utils

import (
	"github.com/nervosnetwork/ckb-sdk-go/crypto/secp256k1"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/example/constant"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/types/resp"
	"github.com/nervosnetwork/ckb-sdk-go/transaction"
	"github.com/nervosnetwork/ckb-sdk-go/types"
)

func Sign(transferCompletion *resp.TransferCompletionResponse) *types.Transaction {
	transferCompletion.GetScriptGroup()
	tx := transferCompletion.GetTransaction()
	scriptGroups := transferCompletion.GetScriptGroup()
	for _, group := range scriptGroups {
		key, _ := secp256k1.HexToKey(constant.GetKey(group.PubKey))
		if err := transaction.SingleSignTransaction(tx, group.Group, group.WitnessArgs, key); err != nil {
			panic(err)
		}
	}
	return tx
}

func SignByKey(transferCompletion *resp.TransferCompletionResponse, privateKey string) *types.Transaction {
	transferCompletion.GetScriptGroup()
	tx := transferCompletion.GetTransaction()
	scriptGroups := transferCompletion.GetScriptGroup()
	for _, group := range scriptGroups {
		key, _ := secp256k1.HexToKey(privateKey)
		if err := transaction.SingleSignTransaction(tx, group.Group, group.WitnessArgs, key); err != nil {
			panic(err)
		}
	}
	return tx
}
