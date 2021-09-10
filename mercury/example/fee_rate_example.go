package test

import (
	"context"
	"fmt"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/example/constant"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/example/utils"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/action"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/source"
	"math/big"
	"testing"
)

func TestDefaultFeeRate(t *testing.T) {
	mercuryApi := constant.GetMercuryApiInstance()

	builder := model.NewTransferBuilder()
	builder.AddUdtHash(constant.UDT_HASH)
	builder.AddFromKeyAddresses([]string{constant.TEST_ADDRESS1}, source.Unconstrained)
	builder.AddToKeyAddressItem(constant.TEST_ADDRESS2, action.Pay_by_to, big.NewInt(100))
	// default 1000 shannons/KB
	//builder.AddFeeRate(1000)

	transferCompletion, err := mercuryApi.BuildTransferTransaction(builder.Build())
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

func TestCustomizedFeeRate(t *testing.T) {
	mercuryApi := constant.GetMercuryApiInstance()

	builder := model.NewTransferBuilder()
	builder.AddUdtHash(constant.UDT_HASH)
	builder.AddFromKeyAddresses([]string{constant.TEST_ADDRESS1}, source.Unconstrained)
	builder.AddToKeyAddressItem(constant.TEST_ADDRESS2, action.Pay_by_to, big.NewInt(100))
	builder.AddFeeRate(10000)

	transferCompletion, err := mercuryApi.BuildTransferTransaction(builder.Build())
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
