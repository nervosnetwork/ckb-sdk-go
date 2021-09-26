package test

import (
	"context"
	"encoding/json"
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

func TestTransferCompletionCkbWithHoldByFrom(t *testing.T) {
	mercuryApi := constant.GetMercuryApiInstance()

	builder := model.NewTransferBuilder()
	item, _ := req.NewIdentityItemByAddress(constant.TEST_ADDRESS0)
	builder.AddFrom(source.Free, item)
	builder.AddTo(mode.HoldByFrom, model.NewToInfo(constant.TEST_ADDRESS4, amount.CkbToShannon(100)))
	builder.AddAssetInfo(common.NewCkbAsset())

	transferCompletion, err := mercuryApi.BuildTransferTransaction(builder.Build())
	if err != nil {
		t.Error(err)
	}

	marshal, _ := json.Marshal(transferCompletion)
	fmt.Println(string(marshal))

	tx := utils.Sign(transferCompletion)

	hash, err := mercuryApi.SendTransaction(context.Background(), tx)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(hash)
}

func TestTransferCompletionSudtWithHoldByFrom(t *testing.T) {
	mercuryApi := constant.GetMercuryApiInstance()

	builder := model.NewTransferBuilder()
	item, _ := req.NewIdentityItemByAddress(constant.TEST_ADDRESS0)
	builder.AddFrom(source.Free, item)
	builder.AddTo(mode.HoldByFrom, model.NewToInfo(constant.TEST_ADDRESS4, model.NewU128WithU64(100)))
	builder.AddAssetInfo(common.NewUdtAsset(constant.UDT_HASH))

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

func TestTransferCompletionCkbWithClaimable(t *testing.T) {
	mercuryApi := constant.GetMercuryApiInstance()

	builder := model.NewTransferBuilder()
	item, _ := req.NewIdentityItemByAddress(constant.TEST_ADDRESS1)
	builder.AddFrom(source.Claimable, item)
	builder.AddTo(mode.HoldByFrom, model.NewToInfo(constant.TEST_ADDRESS2, amount.CkbToShannon(100)))
	builder.AddAssetInfo(common.NewCkbAsset())

	_, err := mercuryApi.BuildTransferTransaction(builder.Build())
	if err != nil && err.Error() != "The transaction does not support ckb" {
		panic(err)
	}

}

func TestTransferCompletionCkbWithHoldByTo(t *testing.T) {
	mercuryApi := constant.GetMercuryApiInstance()

	builder := model.NewTransferBuilder()
	item, _ := req.NewIdentityItemByAddress(constant.TEST_ADDRESS1)
	builder.AddFrom(source.Claimable, item)
	builder.AddTo(mode.HoldByTo, model.NewToInfo(constant.TEST_ADDRESS2, amount.CkbToShannon(100)))
	builder.AddAssetInfo(common.NewCkbAsset())

	_, err := mercuryApi.BuildTransferTransaction(builder.Build())
	if err != nil && err.Error() != "The transaction does not support ckb" {
		t.Error(err)
	}
}

func TestTransferCompletionSudtWithPayByTo(t *testing.T) {
	mercuryApi := constant.GetMercuryApiInstance()

	builder := model.NewTransferBuilder()
	item, _ := req.NewIdentityItemByAddress(constant.TEST_ADDRESS1)
	builder.AddFrom(source.Free, item)
	builder.AddTo(mode.HoldByTo, model.NewToInfo(constant.TEST_ADDRESS2, model.NewU128WithU64(100)))
	builder.AddAssetInfo(common.NewUdtAsset(constant.UDT_HASH))

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
