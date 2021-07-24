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

func TestDefaultFeeRate(t *testing.T) {
	mercuryApi := constant.GetMercuryApiInstance()
	ckbNode := constant.GetCkbNodeInstance()

	builder := model.NewTransferBuilder()
	builder.AddUdtHash("0xf21e7350fa9518ed3cbb008e0e8c941d7e01a12181931d5608aa366ee22228bd")
	builder.AddFromKeyAddresses([]string{constant.TEST_ADDRESS1}, source.Unconstrained)
	builder.AddKeyAddressItem(constant.TEST_ADDRESS2, action.Pay_by_to, 100)
	// default 1000 shannons/KB
	//builder.AddFeeRate(1000)

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

func TestCustomizedFeeRate(t *testing.T) {
	mercuryApi := constant.GetMercuryApiInstance()
	ckbNode := constant.GetCkbNodeInstance()

	builder := model.NewTransferBuilder()
	builder.AddUdtHash("0xf21e7350fa9518ed3cbb008e0e8c941d7e01a12181931d5608aa366ee22228bd")
	builder.AddFromKeyAddresses([]string{constant.TEST_ADDRESS1}, source.Unconstrained)
	builder.AddKeyAddressItem(constant.TEST_ADDRESS2, action.Pay_by_to, 100)
	builder.AddFeeRate(10000)

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
