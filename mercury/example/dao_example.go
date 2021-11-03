package test

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/example/constant"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/example/utils"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/req"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/source"
	"github.com/nervosnetwork/ckb-sdk-go/utils/amount"
	"testing"
)

func TestDaoDepositWithAddress(t *testing.T) {
	builder := model.NewDaoDepositPayloadBuilder()
	item, _ := req.NewAddressItem(constant.TEST_ADDRESS3)
	builder.AddFrom(source.Free, item)
	builder.AddAmount(amount.CkbToShannon(300).Uint64())

	transaction, err := constant.GetMercuryApiInstance().BuildDaoDepositTransaction(builder.Build())

	marshal, _ := json.Marshal(builder.Build())
	fmt.Println(string(marshal))

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

func TestDaoDepositWithIdentity(t *testing.T) {
	builder := model.NewDaoDepositPayloadBuilder()
	item, _ := req.NewIdentityItemByCkb(constant.TEST_PUBKEY3)
	builder.AddFrom(source.Free, item)
	builder.AddAmount(amount.CkbToShannon(300).Uint64())

	transaction, err := constant.GetMercuryApiInstance().BuildDaoDepositTransaction(builder.Build())

	marshal, _ := json.Marshal(builder.Build())
	fmt.Println(string(marshal))

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

func TestDaoWithdraw(t *testing.T) {
	builder := model.NewDaoWithdrawPayloadBuilder()
	item, _ := req.NewIdentityItemByCkb(constant.TEST_PUBKEY3)
	builder.AddItem(item)
	builder.AddPayFee(constant.TEST_ADDRESS1)

	transaction, err := constant.GetMercuryApiInstance().BuildDaoWithdrawTransaction(builder.Build())

	marshal, _ := json.Marshal(builder.Build())
	fmt.Println(string(marshal))

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

func TestDaoClaim(t *testing.T) {
	builder := model.NewDaoClaimPayloadBuilder()
	item, _ := req.NewIdentityItemByCkb(constant.TEST_PUBKEY3)
	builder.AddItem(item)

	transaction, err := constant.GetMercuryApiInstance().BuildDaoClaimTransaction(builder.Build())

	marshal, _ := json.Marshal(builder.Build())
	fmt.Println(string(marshal))

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
