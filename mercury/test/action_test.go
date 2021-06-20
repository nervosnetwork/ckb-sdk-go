package test

import (
	"context"
	"fmt"
	"github.com/nervosnetwork/ckb-sdk-go/crypto/secp256k1"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/action"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/resp"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/source"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/test/constant"
	"github.com/nervosnetwork/ckb-sdk-go/transaction"
	"github.com/nervosnetwork/ckb-sdk-go/types"
	"testing"
)

func TestTransferCompletionCkbWithPayByFrom(t *testing.T) {
	mercuryApi := constant.GetMercuryApiInstance()
	ckbNode := constant.GetCkbNodeInstance()

	transferPayload := getTransferPayload(constant.TEST_ADDRESS0, constant.TEST_ADDRESS4, "", action.Pay_by_from)
	transferCompletion, err := mercuryApi.BuildTransferTransaction(transferPayload)
	if err != nil {
		t.Error(err)
	}

	tx := sign(transferCompletion)

	hash, err := ckbNode.SendTransaction(context.Background(), tx)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(hash)
}

func TestTransferCompletionSudtWithPayByFrom(t *testing.T) {
	mercuryApi := constant.GetMercuryApiInstance()
	ckbNode := constant.GetCkbNodeInstance()

	transferPayload := getTransferPayload(constant.TEST_ADDRESS0, constant.TEST_ADDRESS4, "0xf21e7350fa9518ed3cbb008e0e8c941d7e01a12181931d5608aa366ee22228bd", action.Pay_by_from)
	transferCompletion, err := mercuryApi.BuildTransferTransaction(transferPayload)
	if err != nil {
		t.Error(err)
	}

	tx := sign(transferCompletion)

	hash, err := ckbNode.SendTransaction(context.Background(), tx)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(hash)
}

func TestTransferCompletionCkbWithLendByFrom(t *testing.T) {
	mercuryApi := constant.GetMercuryApiInstance()

	transferPayload := getTransferPayload(constant.TEST_ADDRESS1, constant.TEST_ADDRESS2, "", action.Lend_by_from)
	_, err := mercuryApi.BuildTransferTransaction(transferPayload)
	if err != nil && err.Error() != "The transaction does not support ckb" {
		panic(err)
	}

}

func TestTransferCompletionSudtWithLendByFrom(t *testing.T) {
	mercuryApi := constant.GetMercuryApiInstance()
	ckbNode := constant.GetCkbNodeInstance()

	transferPayload := getTransferPayload(constant.TEST_ADDRESS1, constant.TEST_ADDRESS2, "0xf21e7350fa9518ed3cbb008e0e8c941d7e01a12181931d5608aa366ee22228bd", action.Lend_by_from)
	transferCompletion, err := mercuryApi.BuildTransferTransaction(transferPayload)
	if err != nil {
		t.Error(err)
	}

	tx := sign(transferCompletion)

	hash, err := ckbNode.SendTransaction(context.Background(), tx)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(hash)
}

func TestTransferCompletionCkbWithPayByTo(t *testing.T) {
	mercuryApi := constant.GetMercuryApiInstance()

	transferPayload := getTransferPayload(constant.TEST_ADDRESS1, constant.TEST_ADDRESS2, "", action.Pay_by_to)
	_, err := mercuryApi.BuildTransferTransaction(transferPayload)
	if err != nil && err.Error() != "The transaction does not support ckb" {
		t.Error(err)
	}
}

func TestTransferCompletionSudtWithPayByTo(t *testing.T) {
	mercuryApi := constant.GetMercuryApiInstance()
	ckbNode := constant.GetCkbNodeInstance()

	transferPayload := getTransferPayload(constant.TEST_ADDRESS1, constant.TEST_ADDRESS2, "0xf21e7350fa9518ed3cbb008e0e8c941d7e01a12181931d5608aa366ee22228bd", action.Pay_by_to)
	transferCompletion, err := mercuryApi.BuildTransferTransaction(transferPayload)
	if err != nil {
		t.Error(err)
	}

	tx := sign(transferCompletion)

	hash, err := ckbNode.SendTransaction(context.Background(), tx)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(hash)
}

func getTransferPayload(from, to, udtHash, action string) *model.TransferPayload {
	builder := new(model.TransferBuilder)
	builder.AddUdtHash(udtHash)
	builder.AddFrom([]string{from}, source.Unconstrained)
	builder.AddItem(to, action, 100)
	builder.AddFee(10000000)

	return builder.Build()
}

func sign(transferCompletion *resp.TransferCompletionResponse) *types.Transaction {
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
