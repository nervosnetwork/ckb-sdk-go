package test

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/nervosnetwork/ckb-sdk-go/address"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/example/constant"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/example/utils"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/common"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/req"
	"testing"
)

func TestCreateAsset(t *testing.T) {
	address, _ := address.GenerateShortAddress(address.Testnet)

	item, _ := req.NewIdentityItemByAddress(address.Address)
	builder := model.NewBuildAdjustAccountPayloadBuilder()
	builder.AddItem(item)
	from, _ := req.NewIdentityItemByCkb(constant.TEST_PUBKEY3)
	builder.AddFrom(from)
	builder.AddAssetInfo(common.NewUdtAsset(constant.UDT_HASH))
	builder.AddAccountNumber(1)

	transferCompletion, err := constant.GetMercuryApiInstance().BuildAdjustAccountTransaction(builder.Build())
	if err != nil {
		t.Error(err)
	}

	tx := utils.Sign(transferCompletion)

	hash, err := constant.GetMercuryApiInstance().SendTransaction(context.Background(), tx)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(hash)
}

func TestAdjustAssetAccountWithUdt(t *testing.T) {
	item, _ := req.NewIdentityItemByCkb(constant.TEST_PUBKEY4)
	builder := model.NewBuildAdjustAccountPayloadBuilder()
	builder.AddItem(item)
	builder.AddAssetInfo(common.NewUdtAsset(constant.UDT_HASH))
	builder.AddAccountNumber(1)

	marshal, _ := json.Marshal(builder.Build())
	fmt.Println(string(marshal))

	transferCompletion, err := constant.GetMercuryApiInstance().BuildAdjustAccountTransaction(builder.Build())
	if err != nil {
		t.Error(err)
	}

	if transferCompletion.TxView == nil {
		return
	}

	tx := utils.Sign(transferCompletion)

	hash, err := constant.GetMercuryApiInstance().SendTransaction(context.Background(), tx)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(hash)
}

func TestAdjustAssetExtraCkb(t *testing.T) {
	item, _ := req.NewIdentityItemByCkb(constant.TEST_PUBKEY4)
	builder := model.NewBuildAdjustAccountPayloadBuilder()
	builder.AddItem(item)
	builder.AddAssetInfo(common.NewUdtAsset(constant.UDT_HASH))
	builder.AddExtraCKB(100)
	builder.AddAccountNumber(3)

	marshal, _ := json.Marshal(builder.Build())
	fmt.Println(string(marshal))

	transferCompletion, err := constant.GetMercuryApiInstance().BuildAdjustAccountTransaction(builder.Build())
	if err != nil {
		t.Error(err)
	}

	if transferCompletion.TxView == nil {
		return
	}

	tx := utils.Sign(transferCompletion)

	hash, err := constant.GetMercuryApiInstance().SendTransaction(context.Background(), tx)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(hash)
}
