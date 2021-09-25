package test

import (
	"encoding/json"
	"fmt"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/types/req"
	"testing"

	"github.com/nervosnetwork/ckb-sdk-go/mercury/example/constant"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/types"
)

func TestGetCkbBalance(t *testing.T) {
	builder := req.NewGetBalancePayloadBuilder()
	builder.SetItemAsAddress(constant.TEST_ADDRESS4)
	builder.AddAssetInfo(types.NewCkbAsset())

	balance, _ := constant.GetMercuryApiInstance().GetBalance(builder.Build())

	marshal, _ := json.Marshal(balance)
	fmt.Println(string(marshal))
}

func TestGetSudtBalance(t *testing.T) {
	builder := req.NewGetBalancePayloadBuilder()
	builder.SetItemAsAddress(constant.TEST_ADDRESS4)
	builder.AddAssetInfo(types.NewUdtAsset(constant.UDT_HASH))
	balance, _ := constant.GetMercuryApiInstance().GetBalance(builder.Build())

	marshal, _ := json.Marshal(balance)
	fmt.Println(string(marshal))
}

func TestAllBalance(t *testing.T) {
	builder := req.NewGetBalancePayloadBuilder()
	builder.SetItemAsAddress(constant.TEST_ADDRESS4)

	balance, _ := constant.GetMercuryApiInstance().GetBalance(builder.Build())

	marshal, _ := json.Marshal(balance)
	fmt.Println(string(marshal))
	fmt.Println(len(balance.Balances))

}
