package test

import (
	"encoding/json"
	"fmt"
	"github.com/nervosnetwork/ckb-sdk-go/address"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/example/constant"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetBalance(t *testing.T) {
	builder := model.NewGetBalancePayloadBuilder()
	builder.AddKeyAddress(&model.KeyAddress{constant.TEST_ADDRESS4})

	balance, _ := constant.GetMercuryApiInstance().GetBalance(builder.Build())

	marshal, _ := json.Marshal(balance)
	fmt.Println(string(marshal))
}

func TestGetSudtBalance(t *testing.T) {
	builder := model.NewGetBalancePayloadBuilder()
	builder.AddKeyAddress(&model.KeyAddress{constant.TEST_ADDRESS4})
	builder.AddUdtHash(constant.UDT_HASH)

	balance, _ := constant.GetMercuryApiInstance().GetBalance(builder.Build())

	marshal, _ := json.Marshal(balance)
	fmt.Println(string(marshal))
}

func TestAllBalance(t *testing.T) {
	builder := model.NewGetBalancePayloadBuilder()
	builder.AddKeyAddress(&model.KeyAddress{constant.TEST_ADDRESS4})
	builder.AllBalance()

	balance, _ := constant.GetMercuryApiInstance().GetBalance(builder.Build())

	marshal, _ := json.Marshal(balance)
	fmt.Println(string(marshal))
	fmt.Println(len(balance.Balances))

}

func TestNormalAddressWithAcpAddress(t *testing.T) {
	acpAddress, err := address.GenerateAcpAddress(constant.TEST_ADDRESS4)
	assert.Nil(t, err)

	builder := model.NewGetBalancePayloadBuilder()
	builder.AddNormalAddress(&model.NormalAddress{acpAddress})
	builder.AddUdtHash(constant.UDT_HASH)

	balance, _ := constant.GetMercuryApiInstance().GetBalance(builder.Build())

	marshal, _ := json.Marshal(balance)
	fmt.Println(string(marshal))
}

func TestNormalAddressWithSecp256k1Address(t *testing.T) {
	builder := model.NewGetBalancePayloadBuilder()
	builder.AddNormalAddress(&model.NormalAddress{constant.TEST_ADDRESS4})
	builder.AddUdtHash(constant.UDT_HASH)

	balance, _ := constant.GetMercuryApiInstance().GetBalance(builder.Build())

	marshal, _ := json.Marshal(balance)
	fmt.Println(string(marshal))
}
