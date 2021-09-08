package test

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/example/constant"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/example/utils"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/action"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/source"
	"github.com/nervosnetwork/ckb-sdk-go/utils/amount"
	"math/big"
	"testing"
)

func TestTransferCompletionCkbWithPayByFrom(t *testing.T) {
	mercuryApi := constant.GetMercuryApiInstance()

	transferPayload := getTransferPayload(constant.TEST_ADDRESS0, constant.TEST_ADDRESS4, "", action.Pay_by_from)
	transferCompletion, err := mercuryApi.BuildTransferTransaction(transferPayload)
	if err != nil {
		t.Error(err)
	}

	tx := utils.Sign(transferCompletion)

	hash, err := mercuryApi.SendTransaction(context.Background(), tx)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(hash)
}

func TestTransferCompletionSudtWithPayByFrom(t *testing.T) {
	mercuryApi := constant.GetMercuryApiInstance()

	transferPayload := getTransferPayload(constant.TEST_ADDRESS0, constant.TEST_ADDRESS4, constant.UDT_HASH, action.Pay_by_from)
	transferCompletion, err := mercuryApi.BuildTransferTransaction(transferPayload)
	if err != nil {
		t.Error(err)
	}

	tx := utils.Sign(transferCompletion)

	hash, err := mercuryApi.SendTransaction(context.Background(), tx)
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

	transferPayload := getTransferPayload(constant.TEST_ADDRESS1, constant.TEST_ADDRESS2, constant.UDT_HASH, action.Lend_by_from)
	transferCompletion, err := mercuryApi.BuildTransferTransaction(transferPayload)
	if err != nil {
		t.Error(err)
	}

	tx := utils.Sign(transferCompletion)

	hash, err := mercuryApi.SendTransaction(context.Background(), tx)
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

	transferPayload := getTransferPayload(constant.TEST_ADDRESS1, constant.QUERY_TRANSACTION_ADDRESS, constant.UDT_HASH, action.Pay_by_to)
	transferCompletion, err := mercuryApi.BuildTransferTransaction(transferPayload)
	if err != nil {
		t.Error(err)
	}

	tx := utils.Sign(transferCompletion)

	hash, err := mercuryApi.SendTransaction(context.Background(), tx)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(hash)
}

func getTransferPayload(from, to, udtHash, action string) *model.TransferPayload {
	builder := model.NewTransferBuilder()
	builder.AddUdtHash(udtHash)
	builder.AddFromKeyAddresses([]string{from}, source.Unconstrained)
	if udtHash != "" {
		builder.AddToKeyAddressItem(to, action, big.NewInt(100))
	} else {
		builder.AddToKeyAddressItem(to, action, amount.CkbToShannon(100))
	}

	marshal, _ := json.Marshal(builder.Build())
	fmt.Println(string(marshal))

	return builder.Build()
}
