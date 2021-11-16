package test

import (
	"context"
	"fmt"
	"github.com/nervosnetwork/ckb-sdk-go/address"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/example/constant"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/example/utils"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/common"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/mode"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/req"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/source"
	"github.com/nervosnetwork/ckb-sdk-go/types"
	"github.com/nervosnetwork/ckb-sdk-go/utils/amount"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCkbInsufficientBalanceToPayTheFee1(t *testing.T) {
	from, _ := address.GenerateShortAddress(address.Testnet)
	to, _ := address.GenerateShortAddress(address.Testnet)

	builder := model.NewSimpleTransferPayloadBuilder()
	builder.AddAssetInfo(common.NewCkbAsset())
	builder.AddFrom(from.Address)
	builder.AddToInfo(to.Address, amount.CkbToShannon(100))

	_, err := constant.GetMercuryApiInstance().BuildSimpleTransferTransaction(builder.Build())

	assert.EqualError(
		t,
		err,
		"Mercury Rpc Error code -11009, error Asset type CKB hash 0000000000000000000000000000000000000000000000000000000000000000 token is not enough",
	)
}

func TestCkbInsufficientBalanceToPayTheFee2(t *testing.T) {
	from, _ := address.GenerateShortAddress(address.Testnet)

	builder := model.NewSimpleTransferPayloadBuilder()
	builder.AddAssetInfo(common.NewCkbAsset())
	builder.AddFrom(from.Address)
	builder.AddToInfo(constant.TEST_ADDRESS4, amount.CkbToShannon(100))

	_, err := constant.GetMercuryApiInstance().BuildSimpleTransferTransaction(builder.Build())

	assert.EqualError(
		t,
		err,
		"Mercury Rpc Error code -11009, error Asset type CKB hash 0000000000000000000000000000000000000000000000000000000000000000 token is not enough",
	)
}

func TestSourceByClaimable(t *testing.T) {
	initChequeCell()
	builder := model.NewSimpleTransferPayloadBuilder()
	builder.AddAssetInfo(common.NewUdtAsset(constant.UDT_HASH))
	builder.AddFrom(constant.TEST_ADDRESS2)
	builder.AddToInfo(constant.TEST_ADDRESS4, model.NewU128WithU64(20))

	tx, err := constant.GetMercuryApiInstance().BuildSimpleTransferTransaction(builder.Build())
	if err != nil {
		t.Error(err)
	}

	transaction := utils.Sign(tx)
	hash, err := constant.GetMercuryApiInstance().SendTransaction(context.Background(), transaction)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(hash)
}

func TestSendChequeCell(t *testing.T) {

	to, _ := address.GenerateShortAddress(address.Testnet)

	builder := model.NewSimpleTransferPayloadBuilder()
	builder.AddAssetInfo(common.NewUdtAsset(constant.UDT_HASH))
	builder.AddFrom(constant.TEST_ADDRESS2)
	builder.AddToInfo(to.Address, model.NewU128WithU64(20))

	tx, err := constant.GetMercuryApiInstance().BuildSimpleTransferTransaction(builder.Build())
	if err != nil {
		t.Error(err)
	}

	transaction := utils.Sign(tx)
	hash, err := constant.GetMercuryApiInstance().SendTransaction(context.Background(), transaction)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(hash)
}

func TestSourceByFree(t *testing.T) {
	builder := model.NewSimpleTransferPayloadBuilder()
	builder.AddAssetInfo(common.NewUdtAsset(constant.UDT_HASH))
	builder.AddFrom(constant.TEST_ADDRESS4)
	builder.AddToInfo(constant.TEST_ADDRESS1, model.NewU128WithU64(20))

	tx, err := constant.GetMercuryApiInstance().BuildSimpleTransferTransaction(builder.Build())
	if err != nil {
		t.Error(err)
	}

	transaction := utils.Sign(tx)
	hash, err := constant.GetMercuryApiInstance().SendTransaction(context.Background(), transaction)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(hash)
}

func TestModeByHoldyTo(t *testing.T) {

	builder := model.NewSimpleTransferPayloadBuilder()
	builder.AddAssetInfo(common.NewUdtAsset(constant.UDT_HASH))
	builder.AddFrom(constant.TEST_ADDRESS1)
	acpAddress, _ := address.GenerateAcpAddress(constant.TEST_ADDRESS4)
	builder.AddToInfo(acpAddress, model.NewU128WithU64(20))

	tx, err := constant.GetMercuryApiInstance().BuildSimpleTransferTransaction(builder.Build())
	if err != nil {
		t.Error(err)
	}

	transaction := utils.Sign(tx)
	hash, err := constant.GetMercuryApiInstance().SendTransaction(context.Background(), transaction)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(hash)
}

func initChequeCell() {
	mercuryApi := constant.GetMercuryApiInstance()

	builder := model.NewTransferBuilder()
	builder.AddAssetInfo(common.NewUdtAsset(constant.UDT_HASH))

	item, _ := req.NewIdentityItemByAddress(constant.TEST_ADDRESS1)
	builder.AddFrom(source.Free, item)
	builder.AddTo(mode.HoldByFrom, model.NewToInfo(constant.TEST_ADDRESS2, model.NewU128WithU64(100)))
	transferPayload := builder.Build()
	transferCompletion, err := mercuryApi.BuildTransferTransaction(transferPayload)
	if err != nil {
		fmt.Println(err)
	}

	tx := utils.Sign(transferCompletion)

	hash, err := mercuryApi.SendTransaction(context.Background(), tx)
	if err != nil {
		fmt.Println(err)
	}

	var txStatus types.TransactionStatus = "pending"
	for {
		transaction, _ := mercuryApi.GetTransaction(context.Background(), *hash)
		if transaction.TxStatus.Status != txStatus {
			break
		}
		fmt.Println("Awaiting transaction results")
		time.Sleep(1 * 1e9)

	}

	time.Sleep(60 * 1e9)
	fmt.Printf("send hash of cheque cell transactions: %s\n", hash.String())
}
