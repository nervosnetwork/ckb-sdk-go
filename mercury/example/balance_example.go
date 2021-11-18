package test

import (
	"encoding/json"
	"fmt"
	"github.com/nervosnetwork/ckb-sdk-go/address"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/example/constant"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/common"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/req"
	"github.com/nervosnetwork/ckb-sdk-go/types"
	"testing"
)

func TestGetCkbBalance(t *testing.T) {
	builder := model.NewGetBalancePayloadBuilder()
	item, _ := req.NewIdentityItemByCkb(constant.TEST_PUBKEY4)
	builder.AddItem(item)
	builder.AddAssetInfo(common.NewCkbAsset())

	balance, err := constant.GetMercuryApiInstance().GetBalance(builder.Build())
	if err != nil {
		t.Error(err)
	}

	marshal, _ := json.Marshal(balance)
	fmt.Println(string(marshal))
}

func TestGetSudtBalance(t *testing.T) {
	builder := model.NewGetBalancePayloadBuilder()
	item, _ := req.NewIdentityItemByCkb(constant.TEST_PUBKEY4)
	builder.AddItem(item)
	builder.AddAssetInfo(common.NewUdtAsset(constant.UDT_HASH))
	balance, _ := constant.GetMercuryApiInstance().GetBalance(builder.Build())

	marshal, _ := json.Marshal(balance)
	fmt.Println(string(marshal))
}

func TestAllBalance(t *testing.T) {
	builder := model.NewGetBalancePayloadBuilder()
	item, _ := req.NewIdentityItemByCkb(constant.TEST_PUBKEY4)
	builder.AddItem(item)

	balance, _ := constant.GetMercuryApiInstance().GetBalance(builder.Build())

	marshal, _ := json.Marshal(balance)
	fmt.Println(string(marshal))
	fmt.Println(len(balance.Balances))

}

func TestGetBalanceByAddress(t *testing.T) {
	builder := model.NewGetBalancePayloadBuilder()
	item, _ := req.NewAddressItem(constant.TEST_ADDRESS4)
	builder.AddItem(item)
	builder.AddAssetInfo(common.NewCkbAsset())

	balance, _ := constant.GetMercuryApiInstance().GetBalance(builder.Build())

	marshal, _ := json.Marshal(balance)
	fmt.Println(string(marshal))
	fmt.Println(len(balance.Balances))

}

func TestGetBalanceByIdentity(t *testing.T) {
	builder := model.NewGetBalancePayloadBuilder()
	item, _ := req.NewIdentityItemByCkb(constant.TEST_PUBKEY4)
	builder.AddItem(item)
	builder.AddAssetInfo(common.NewCkbAsset())

	balance, _ := constant.GetMercuryApiInstance().GetBalance(builder.Build())

	marshal, _ := json.Marshal(balance)
	fmt.Println(string(marshal))
	fmt.Println(len(balance.Balances))

}

func TestGetBalanceByRecordByScriptByChequeCellSender(t *testing.T) {

	parse, _ := address.Parse(constant.TEST_ADDRESS1)
	script := parse.Script

	outPoint := &types.OutPoint{
		types.HexToHash("0xecfea4bdf6bf8290d8f8186ed9f4da9b0f8fbba217600b47632f5a72ff677d4d"),
		0,
	}

	builder := model.NewGetBalancePayloadBuilder()
	item, _ := req.NewRecordItemByScript(outPoint, script)
	builder.AddItem(item)
	builder.AddAssetInfo(common.NewCkbAsset())

	balance, _ := constant.GetMercuryApiInstance().GetBalance(builder.Build())

	marshal, _ := json.Marshal(balance)
	fmt.Println(string(marshal))
	fmt.Println(len(balance.Balances))

}

func TestGetBalanceByRecordByScriptChequeCellReceiver(t *testing.T) {

	parse, _ := address.Parse(constant.TEST_ADDRESS2)
	script := parse.Script

	outPoint := &types.OutPoint{
		types.HexToHash("0xecfea4bdf6bf8290d8f8186ed9f4da9b0f8fbba217600b47632f5a72ff677d4d"),
		0,
	}

	builder := model.NewGetBalancePayloadBuilder()
	item, _ := req.NewRecordItemByScript(outPoint, script)
	builder.AddItem(item)
	builder.AddAssetInfo(common.NewCkbAsset())

	balance, _ := constant.GetMercuryApiInstance().GetBalance(builder.Build())

	marshal, _ := json.Marshal(balance)
	fmt.Println(string(marshal))
	fmt.Println(len(balance.Balances))

}

func TestGetBalanceByRecordByAddress(t *testing.T) {

	outPoint := &types.OutPoint{
		types.HexToHash("0xfc43d8bdfff3051f3c908cd137e0766eecba4e88ae5786760c3e0e0f1d76c004"),
		2,
	}

	builder := model.NewGetBalancePayloadBuilder()
	item, _ := req.NewRecordItemByAddress(outPoint, constant.TEST_ADDRESS4)
	fmt.Println(item)
	builder.AddItem(item)
	builder.AddAssetInfo(common.NewCkbAsset())

	fmt.Println(item)

	balance, _ := constant.GetMercuryApiInstance().GetBalance(builder.Build())

	marshal, _ := json.Marshal(balance)
	fmt.Println(string(marshal))
	fmt.Println(len(balance.Balances))

}
