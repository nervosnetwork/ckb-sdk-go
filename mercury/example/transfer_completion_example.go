package test

import (
	"context"
	"fmt"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/example/constant"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/action"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/source"
	"testing"
)

func TestSingleFromSingleTo(t *testing.T) {
	mercuryApi := constant.GetMercuryApiInstance()
	ckbNode := constant.GetCkbNodeInstance()

	builder := model.NewTransferBuilder()
	builder.AddFrom([]string{constant.TEST_ADDRESS1}, source.Unconstrained)
	builder.AddItem(constant.TEST_ADDRESS2, action.Pay_by_from, 100)

	transferCompletion, err := mercuryApi.BuildTransferTransaction(builder.Build())
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

func TestSingleFromMultiTo(t *testing.T) {
	mercuryApi := constant.GetMercuryApiInstance()
	ckbNode := constant.GetCkbNodeInstance()

	builder := model.NewTransferBuilder()
	builder.AddFrom([]string{constant.TEST_ADDRESS1}, source.Unconstrained)
	builder.AddItem(constant.TEST_ADDRESS2, action.Pay_by_from, 100)
	builder.AddItem(constant.TEST_ADDRESS3, action.Pay_by_from, 100)

	transferCompletion, err := mercuryApi.BuildTransferTransaction(builder.Build())
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

func TestMultiFromSingleTo(t *testing.T) {
	mercuryApi := constant.GetMercuryApiInstance()
	ckbNode := constant.GetCkbNodeInstance()

	builder := model.NewTransferBuilder()
	builder.AddFrom([]string{constant.TEST_ADDRESS1, constant.TEST_ADDRESS2}, source.Unconstrained)
	builder.AddItem(constant.TEST_ADDRESS3, action.Pay_by_from, 100)

	transferCompletion, err := mercuryApi.BuildTransferTransaction(builder.Build())
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

func TestMultiFromMultiTo(t *testing.T) {
	mercuryApi := constant.GetMercuryApiInstance()
	ckbNode := constant.GetCkbNodeInstance()

	builder := model.NewTransferBuilder()
	builder.AddFrom([]string{constant.TEST_ADDRESS1, constant.TEST_ADDRESS2}, source.Unconstrained)
	builder.AddItem(constant.TEST_ADDRESS3, action.Pay_by_from, 100)
	builder.AddItem(constant.TEST_ADDRESS4, action.Pay_by_from, 100)

	transferCompletion, err := mercuryApi.BuildTransferTransaction(builder.Build())
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
