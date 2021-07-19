package test

import (
	"encoding/json"
	"fmt"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/example/constant"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model"
	"testing"
)

func TestGetBalance(t *testing.T) {
	builder := model.GetGetBalancePayloadBuilder()
	builder.AddAddress(constant.TEST_ADDRESS0)
	payload, err := builder.Build()
	if err != nil {
		panic(err)
	}
	balance, _ := constant.GetMercuryApiInstance().GetBalance(payload)

	marshal, _ := json.Marshal(balance)
	fmt.Println(string(marshal))
}

func TestGetSudtBalance(t *testing.T) {
	builder := model.GetGetBalancePayloadBuilder()
	builder.AddAddress(constant.TEST_ADDRESS0)
	builder.AddUdtHash("0xf21e7350fa9518ed3cbb008e0e8c941d7e01a12181931d5608aa366ee22228bd")

	payload, err := builder.Build()
	if err != nil {
		panic(err)
	}
	balance, _ := constant.GetMercuryApiInstance().GetBalance(payload)

	marshal, _ := json.Marshal(balance)
	fmt.Println(string(marshal))
}

func TestAllBalance(t *testing.T) {
	builder := model.GetGetBalancePayloadBuilder()
	builder.AddAddress(constant.TEST_ADDRESS0)
	builder.AllBalance()

	payload, err := builder.Build()
	if err != nil {
		panic(err)
	}
	balance, _ := constant.GetMercuryApiInstance().GetBalance(payload)

	marshal, _ := json.Marshal(balance)
	fmt.Println(string(marshal))

}
