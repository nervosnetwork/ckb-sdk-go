package normal

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/example/constant"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/example/utils"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/action"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/source"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/types/req"
	"math/big"
	"testing"
)

const (
	acpAddress = "ckt1qyp07nuu3fpu9rksy677uvchlmyv9ce5saes824qjq"
	key        = "6aa38b72d55efc781c0c2bedcbd8adba2c946d90c1075189749d5049301ca84a"
)

func TestFromAcp(t *testing.T) {
	builder := req.NewTransferBuilder()
	builder.AddUdtHash(constant.UDT_HASH)
	builder.AddFromNormalAddresses([]string{acpAddress})
	builder.AddToKeyAddressItem(constant.TEST_ADDRESS2, action.Lend_by_from, big.NewInt(100))

	marshal, _ := json.Marshal(builder.Build())
	fmt.Println(string(marshal))

	transaction, err := constant.GetMercuryApiInstance().BuildTransferTransaction(builder.Build())
	if err != nil {
		t.Error(err)
	}

	tx := utils.SignByKey(transaction, key)

	hash, err := constant.GetMercuryApiInstance().SendTransaction(context.Background(), tx)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(hash)
}

func TestToAcp(t *testing.T) {
	builder := req.NewTransferBuilder()
	builder.AddUdtHash(constant.UDT_HASH)
	builder.AddFromKeyAddresses([]string{constant.TEST_ADDRESS1}, source.Unconstrained)
	builder.AddToNormalAddressItem(acpAddress, big.NewInt(100))

	marshal, _ := json.Marshal(builder.Build())
	fmt.Println(string(marshal))

	transaction, err := constant.GetMercuryApiInstance().BuildTransferTransaction(builder.Build())
	if err != nil {
		t.Error(err)
	}

	tx := utils.Sign(transaction)

	hash, err := constant.GetMercuryApiInstance().SendTransaction(context.Background(), tx)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(hash)
}
