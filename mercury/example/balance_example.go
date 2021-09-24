package test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/nervosnetwork/ckb-sdk-go/mercury/example/constant"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/common"
)

func TestGetCkbBalance(t *testing.T) {
	builder := model.NewGetBalancePayloadBuilder()
	builder.SetItemAsAddress(constant.TEST_ADDRESS4)
	builder.AddAssetInfo(common.NewCkbAsset())

	balance, _ := constant.GetMercuryApiInstance().GetBalance(builder.Build())

	marshal, _ := json.Marshal(balance)
	fmt.Println(string(marshal))
}

func TestGetSudtBalance(t *testing.T) {
	builder := model.NewGetBalancePayloadBuilder()
	builder.SetItemAsAddress(constant.TEST_ADDRESS4)
	builder.AddAssetInfo(common.NewUdtAsset(constant.UDT_HASH))
	balance, _ := constant.GetMercuryApiInstance().GetBalance(builder.Build())

	marshal, _ := json.Marshal(balance)
	fmt.Println(string(marshal))
}

func TestAllBalance(t *testing.T) {
	builder := model.NewGetBalancePayloadBuilder()
	builder.SetItemAsAddress(constant.TEST_ADDRESS4)

	balance, _ := constant.GetMercuryApiInstance().GetBalance(builder.Build())

	marshal, _ := json.Marshal(balance)
	fmt.Println(string(marshal))
	fmt.Println(len(balance.Balances))

}
