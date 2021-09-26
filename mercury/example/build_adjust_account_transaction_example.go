package test

import (
	"encoding/json"
	"fmt"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/example/constant"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/common"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/req"
	"testing"
)

func TestBuildAdjustAccountTransaction(t *testing.T) {
	mercuryApi := constant.GetMercuryApiInstance()

	payload := model.NewBuildAdjustAccountPayload()
	item, _ := req.NewIdentityItemByCkb("0x512e97d31dcc9f012c550e880cbcc10daafb7aed")
	payload.Item = item
	from, _ := req.NewIdentityItemByCkb(constant.TEST_PUBKEY3)
	payload.AddItemToFrom(from)
	payload.AssetInfo = common.NewUdtAsset(constant.UDT_HASH)
	payload.AccountNumber = 1

	creationTransaction, err := mercuryApi.BuildAdjustAccountTransaction(payload)
	if err != nil {
		t.Error(err)
	}

	json, err := json.Marshal(creationTransaction)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(json))

	//tx := utils.Sign(creationTransaction)
	//
	//hash, err := mercuryApi.SendTransaction(context.Background(), tx)
	//if err != nil {
	//	t.Error(err)
	//}
	//
	//fmt.Println(hash)
}
