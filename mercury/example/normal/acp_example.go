package normal

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/example/constant"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/action"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/source"
	"testing"
)

const (
	acpAddress = "ckt1qyp07nuu3fpu9rksy677uvchlmyv9ce5saes824qjq"
	key        = "6aa38b72d55efc781c0c2bedcbd8adba2c946d90c1075189749d5049301ca84a"
)

func TestFromAcp(t *testing.T) {
	builder := model.NewTransferBuilder()
	builder.AddUdtHash(constant.UdtHash)
	builder.AddFromNormalAddresses([]string{acpAddress})
	builder.AddToKeyAddressItem(constant.TEST_ADDRESS2, action.Lend_by_from, 100)

	marshal, _ := json.Marshal(builder.Build())
	fmt.Println(string(marshal))

	transaction, err := constant.GetMercuryApiInstance().BuildTransferTransaction(builder.Build())
	if err != nil {
		t.Error(err)
	}

	tx := constant.SignByKey(transaction, key)

	hash, err := constant.GetCkbNodeInstance().SendTransaction(context.Background(), tx)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(hash)
}

func TestToAcp(t *testing.T) {
	builder := model.NewTransferBuilder()
	builder.AddUdtHash(constant.UdtHash)
	builder.AddFromKeyAddresses([]string{constant.TEST_ADDRESS1}, source.Unconstrained)
	builder.AddToNormalAddressItem(acpAddress, 100)

	marshal, _ := json.Marshal(builder.Build())
	fmt.Println(string(marshal))

	transaction, err := constant.GetMercuryApiInstance().BuildTransferTransaction(builder.Build())
	if err != nil {
		t.Error(err)
	}

	tx := constant.Sign(transaction)

	hash, err := constant.GetCkbNodeInstance().SendTransaction(context.Background(), tx)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(hash)
}
