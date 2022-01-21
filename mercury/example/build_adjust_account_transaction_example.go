package test

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/nervosnetwork/ckb-sdk-go/address"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/example/constant"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/example/utils"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/common"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/req"
	"testing"
)

func TestCreateAsset(t *testing.T) {
	address, _ := address.GenerateShortAddress(address.Testnet)

	item, _ := req.NewIdentityItemByAddress(address.Address)
	builder := model.NewBuildAdjustAccountPayloadBuilder()
	builder.AddItem(item)
	from, _ := req.NewIdentityItemByCkb(constant.TEST_PUBKEY3)
	builder.AddFrom(from)
	builder.AddAssetInfo(common.NewUdtAsset(constant.UDT_HASH))
	builder.AddAccountNumber(1)

	transferCompletion, err := constant.GetMercuryApiInstance().BuildAdjustAccountTransaction(builder.Build())
	if err != nil {
		t.Error(err)
	}

	tx := utils.Sign(transferCompletion)

	hash, err := constant.GetMercuryApiInstance().SendTransaction(context.Background(), tx)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(hash)
}

func TestAdjustAssetAccountWithUdt(t *testing.T) {
	item, _ := req.NewIdentityItemByCkb(constant.TEST_PUBKEY4)
	builder := model.NewBuildAdjustAccountPayloadBuilder()
	builder.AddItem(item)
	builder.AddAssetInfo(common.NewUdtAsset(constant.UDT_HASH))
	builder.AddAccountNumber(1)

	marshal, _ := json.Marshal(builder.Build())
	fmt.Println(string(marshal))

	transferCompletion, err := constant.GetMercuryApiInstance().BuildAdjustAccountTransaction(builder.Build())
	if err != nil {
		t.Error(err)
	}

	if transferCompletion.TxView == nil {
		return
	}

	tx := utils.Sign(transferCompletion)

	hash, err := constant.GetMercuryApiInstance().SendTransaction(context.Background(), tx)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(hash)
}

func TestAdjustAssetExtraCkb(t *testing.T) {
	item, _ := req.NewIdentityItemByCkb(constant.TEST_PUBKEY4)
	builder := model.NewBuildAdjustAccountPayloadBuilder()
	builder.AddItem(item)
	builder.AddAssetInfo(common.NewUdtAsset(constant.UDT_HASH))
	builder.AddExtraCKB(100)
	builder.AddAccountNumber(3)

	marshal, _ := json.Marshal(builder.Build())
	fmt.Println(string(marshal))

	transferCompletion, err := constant.GetMercuryApiInstance().BuildAdjustAccountTransaction(builder.Build())
	if err != nil {
		t.Error(err)
	}

	if transferCompletion.TxView == nil {
		return
	}

	tx := utils.Sign(transferCompletion)

	hash, err := constant.GetMercuryApiInstance().SendTransaction(context.Background(), tx)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(hash)
}

func TestPWLock(t *testing.T) {
	//item, _ := req.NewIdentityItemByCkb(constant.TEST_PUBKEY4)
	builder := model.NewBuildAdjustAccountPayloadBuilder()
	item, _ := req.NewAddressItem("ckt1qpvvtay34wndv9nckl8hah6fzzcltcqwcrx79apwp2a5lkd07fdxxqdd40lmnsnukjh3qr88hjnfqvc4yg8g0gskp8ffv")
	builder.AddItem(item)
	builder.AddAssetInfo(common.NewUdtAsset("0xf21e7350fa9518ed3cbb008e0e8c941d7e01a12181931d5608aa366ee22228bd"))
	item, _ = req.NewAddressItem("ckt1qzda0cr08m85hc8jlnfp3zer7xulejywt49kt2rr0vthywaa50xwsqthh5pum5pzqpssk47zk67hnd6lm28rnqs4cnj0w")
	builder.AddFrom(item)
	//builder.AddExtraCKB(100)
	builder.AddAccountNumber(3)

	marshal, _ := json.Marshal(builder.Build())
	fmt.Println(string(marshal))

	transferCompletion, err := constant.GetMercuryApiInstance().BuildAdjustAccountTransaction(builder.Build())
	if err != nil {
		t.Error(err)
	}

	if transferCompletion.TxView == nil {
		return
	}

	tx := utils.Sign(transferCompletion)
	fmt.Println(tx)

	hash, err := constant.GetMercuryApiInstance().SendTransaction(context.Background(), tx)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(hash)
}
