package test

import (
	"context"
	"fmt"
	"github.com/nervosnetwork/ckb-sdk-go/crypto/secp256k1"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/example/constant"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model"
	"github.com/nervosnetwork/ckb-sdk-go/transaction"
	"testing"
)

func TestAssetAccountCreationTransaction(t *testing.T) {
	mercuryApi := constant.GetMercuryApiInstance()
	ckbNode := constant.GetCkbNodeInstance()

	builder := model.NewCreateAssetAccountPayloadBuilder()
	builder.AddKeyAddress(constant.TEST_ADDRESS3)
	builder.AddUdtHash("0xf21e7350fa9518ed3cbb008e0e8c941d7e01a12181931d5608aa366ee22228bd")

	creationTransaction, err := mercuryApi.BuildAssetAccountCreationTransaction(builder.Build())
	if err != nil {
		t.Error(err)
	}

	creationTransaction.GetScriptGroup()
	tx := creationTransaction.GetTransaction()
	scriptGroups := creationTransaction.GetScriptGroup()
	for _, group := range scriptGroups {
		key, _ := secp256k1.HexToKey("2d4cf0546a1dc93092ad56f2e18fbe6e41ee477d9dec0575cf43b69740ce9f74")
		err = transaction.SingleSignTransaction(tx, group.Group, group.WitnessArgs, key)
	}

	hash, err := ckbNode.SendTransaction(context.Background(), tx)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(hash)
}
