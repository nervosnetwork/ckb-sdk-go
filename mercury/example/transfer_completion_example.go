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
	"github.com/nervosnetwork/ckb-sdk-go/utils/amount"
	"testing"
)

func TestSingleFromSingleTo(t *testing.T) {
	mercuryApi := constant.GetMercuryApiInstance()

	builder := model.NewTransferBuilder()
	builder.AddAssetInfo(common.NewCkbAsset())
	item, _ := req.NewAddressItem(constant.TEST_ADDRESS2)
	builder.AddFrom(source.Free, item)
	builder.AddTo(mode.HoldByFrom, model.NewToInfo(constant.TEST_ADDRESS2, amount.CkbToShannon(100)))

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

func TestSingleFromMultiTo(t *testing.T) {
	mercuryApi := constant.GetMercuryApiInstance()

	builder := model.NewTransferBuilder()
	builder.AddAssetInfo(common.NewCkbAsset())
	item, _ := req.NewAddressItem(constant.TEST_ADDRESS1)
	builder.AddFrom(source.Free, item)
	builder.AddTo(mode.HoldByFrom, model.NewToInfo(constant.TEST_ADDRESS3, amount.CkbToShannon(100)),
		model.NewToInfo(constant.TEST_ADDRESS2, amount.CkbToShannon(100)))

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

func TestMultiFromSingleTo(t *testing.T) {
	mercuryApi := constant.GetMercuryApiInstance()

	builder := model.NewTransferBuilder()
	builder.AddAssetInfo(common.NewCkbAsset())
	item1, _ := req.NewAddressItem(constant.TEST_ADDRESS1)
	item2, _ := req.NewAddressItem(constant.TEST_ADDRESS2)
	builder.AddFrom(source.Free, item1, item2)
	builder.AddTo(mode.HoldByFrom, model.NewToInfo(constant.TEST_ADDRESS3, amount.CkbToShannon(100)))

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

func TestMultiFromMultiTo(t *testing.T) {
	mercuryApi := constant.GetMercuryApiInstance()

	builder := model.NewTransferBuilder()
	builder.AddAssetInfo(common.NewCkbAsset())

	item1, _ := req.NewAddressItem(constant.TEST_ADDRESS1)
	item2, _ := req.NewAddressItem(constant.TEST_ADDRESS2)
	builder.AddFrom(source.Free, item1, item2)

	builder.AddTo(mode.HoldByFrom, model.NewToInfo(constant.TEST_ADDRESS3, amount.CkbToShannon(100)),
		model.NewToInfo(constant.TEST_ADDRESS4, amount.CkbToShannon(100)))

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
