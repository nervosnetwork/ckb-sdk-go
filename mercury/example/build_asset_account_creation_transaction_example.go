package test

import (
	"context"
	"fmt"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/example/constant"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/example/utils"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model"
	"testing"
)

func TestAssetAccountCreationTransaction(t *testing.T) {
	mercuryApi := constant.GetMercuryApiInstance()

	builder := model.NewCreateAssetAccountPayloadBuilder()
	builder.AddKeyAddress(constant.TEST_ADDRESS3)
	builder.AddUdtHash(constant.UDT_HASH)

	creationTransaction, err := mercuryApi.BuildAssetAccountCreationTransaction(builder.Build())
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
