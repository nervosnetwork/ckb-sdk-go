package test

import (
	"context"
	"fmt"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/example/constant"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/example/utils"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/common"
	"testing"
)

func TestBuildAdjustAccountTransaction(t *testing.T) {
	mercuryApi := constant.GetMercuryApiInstance()

	builder := model.NewAdjustAccountPayloadBuilder()
	builder.AddKeyAddress(constant.TEST_ADDRESS3)
	builder.AddAssetInfo(common.NewUdtAsset(constant.UDT_HASH))

	creationTransaction, err := mercuryApi.BuildAdjustAccountTransaction(builder.Build())
	if err != nil {
		t.Error(err)
	}

	tx := utils.Sign(creationTransaction)

	hash, err := mercuryApi.SendTransaction(context.Background(), tx)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(hash)
}
