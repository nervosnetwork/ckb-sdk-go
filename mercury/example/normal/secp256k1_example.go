package normal

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/example/constant"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/example/utils"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/action"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/source"
	"math/big"
	"testing"
)

func TestFromSecp256k1(t *testing.T) {
	builder := model.NewTransferBuilder()
	builder.AddUdtHash(constant.UDT_HASH)
	builder.AddFromNormalAddresses([]string{constant.TEST_ADDRESS1})
	builder.AddToKeyAddressItem(constant.TEST_ADDRESS2, action.Lend_by_from, big.NewInt(100))

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

func TestToSecp256k1(t *testing.T) {
	builder := model.NewTransferBuilder()
	builder.AddUdtHash(constant.UDT_HASH)
	builder.AddFromKeyAddresses([]string{constant.TEST_ADDRESS1}, source.Unconstrained)
	builder.AddToNormalAddressItem(constant.TEST_ADDRESS2, big.NewInt(100))

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
