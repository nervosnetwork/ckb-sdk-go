package test

import (
	"context"
	"fmt"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/example/constant"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/example/utils"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/common"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/mode"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/req"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/source"
	"testing"
)

func TestDefaultFeeRate(t *testing.T) {
	mercuryApi := constant.GetMercuryApiInstance()

	builder := model.NewTransferBuilder()
	builder.AddAssetInfo(common.NewUdtAsset(constant.UDT_HASH))
	item, _ := req.NewAddressItem(constant.TEST_ADDRESS1)
	builder.AddFrom(source.Free, item)
	builder.AddTo(mode.HoldByTo, model.NewToInfo(constant.TEST_ADDRESS2, model.NewU128WithU64(100)))
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
	builder.AddAssetInfo(common.NewUdtAsset(constant.UDT_HASH))
	item, _ := req.NewAddressItem(constant.TEST_ADDRESS1)
	builder.AddFrom(source.Free, item)
	builder.AddTo(mode.HoldByTo, model.NewToInfo(constant.TEST_ADDRESS2, model.NewU128WithU64(100)))
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
