package test

import (
	"encoding/json"
	"fmt"
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
	builder := model.NewGetBalancePayloadBuilder()
	item := req.NewOutpointItem(
		types.HexToHash("0x52b1cf0ad857d53e1a3552944c1acf268f6a6aea8e8fc85fe8febcb8127d56f0"),
		0)
	builder.AddItem(item)
	builder.AddAssetInfo(common.NewCkbAsset())

	balance, _ := constant.GetMercuryApiInstance().GetBalance(builder.Build())

	marshal, _ := json.Marshal(balance)
	fmt.Println(string(marshal))
	fmt.Println(len(balance.Balances))
}

func TestGetBalanceByRecordByScriptChequeCellReceiver(t *testing.T) {
	builder := model.NewGetBalancePayloadBuilder()
	item := req.NewOutpointItem(
		types.HexToHash("0x52b1cf0ad857d53e1a3552944c1acf268f6a6aea8e8fc85fe8febcb8127d56f0"),
		0)
	builder.AddItem(item)
	builder.AddAssetInfo(common.NewCkbAsset())

	balance, _ := constant.GetMercuryApiInstance().GetBalance(builder.Build())

	marshal, _ := json.Marshal(balance)
	fmt.Println(string(marshal))
	fmt.Println(len(balance.Balances))

}

func TestGetBalanceByRecordByAddress(t *testing.T) {
	builder := model.NewGetBalancePayloadBuilder()
	item := req.NewOutpointItem(
		types.HexToHash("0xfc43d8bdfff3051f3c908cd137e0766eecba4e88ae5786760c3e0e0f1d76c004"),
		2)
	fmt.Println(item)
	builder.AddItem(item)
	builder.AddAssetInfo(common.NewCkbAsset())

	fmt.Println(item)

	balance, _ := constant.GetMercuryApiInstance().GetBalance(builder.Build())

	marshal, _ := json.Marshal(balance)
	fmt.Println(string(marshal))
	fmt.Println(len(balance.Balances))

}
