package test

//
//import (
//	"context"
//	"fmt"
//	"github.com/nervosnetwork/ckb-sdk-go/address"
//	"github.com/nervosnetwork/ckb-sdk-go/mercury/example/constant"
//	"github.com/nervosnetwork/ckb-sdk-go/mercury/example/utils"
//	"github.com/nervosnetwork/ckb-sdk-go/mercury/model"
//	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/mode"
//	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/common"
//	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/source"
//	"github.com/nervosnetwork/ckb-sdk-go/types"
//	"github.com/nervosnetwork/ckb-sdk-go/utils/amount"
//	"github.com/stretchr/testify/assert"
//	"math/big"
//	"testing"
//	"time"
//)
//
//func TestAccountNumber(t *testing.T) {
//	number, err := constant.GetMercuryApiInstance().GetAccountNumber(constant.TEST_ADDRESS4)
//	if err != nil {
//		t.Error(err)
//	}
//	fmt.Println(number)
//}
//
//func TestCkbInsufficientBalanceToPayTheFee1(t *testing.T) {
//	from, _ := address.GenerateShortAddress(address.Testnet)
//	to, _ := address.GenerateShortAddress(address.Testnet)
//
//	builder := model.NewSmartTransferPayloadBuilder()
//	builder.AddAssetInfo(common.NewCkbAsset())
//	builder.AddFrom(from.Address)
//	builder.AddSmartTo(to.Address, big.NewInt(100))
//
//	_, err := constant.GetMercuryApiInstance().BuildSmartTransferTransaction(builder.Build())
//
//	assert.EqualError(
//		t,
//		err,
//		"CKB Insufficient balance to pay the fee",
//	)
//}
//
//func TestCkbInsufficientBalanceToPayTheFee2(t *testing.T) {
//	from, _ := address.GenerateShortAddress(address.Testnet)
//
//	builder := model.NewSmartTransferPayloadBuilder()
//	builder.AddAssetInfo(common.NewCkbAsset())
//	builder.AddFrom(from.Address)
//	builder.AddSmartTo(constant.TEST_ADDRESS4, amount.CkbToShannon(100))
//
//	tx, err := constant.GetMercuryApiInstance().BuildSmartTransferTransaction(builder.Build())
//	if err != nil {
//		t.Error(err)
//	}
//
//	transaction := utils.Sign(tx)
//	hash, err := constant.GetMercuryApiInstance().SendTransaction(context.Background(), transaction)
//	if err != nil {
//		t.Error(err)
//	}
//
//	fmt.Println(hash)
//}
//
//func TestSourceByFleeting(t *testing.T) {
//	initChequeCell()
//
//	builder := model.NewSmartTransferPayloadBuilder()
//	builder.AddAssetInfo(common.NewUdtAsset(constant.UDT_HASH))
//	builder.AddFrom(constant.TEST_ADDRESS2)
//	builder.AddSmartTo(constant.TEST_ADDRESS4, big.NewInt(20))
//
//	tx, err := constant.GetMercuryApiInstance().BuildSmartTransferTransaction(builder.Build())
//	if err != nil {
//		t.Error(err)
//	}
//
//	transaction := utils.Sign(tx)
//	hash, err := constant.GetMercuryApiInstance().SendTransaction(context.Background(), transaction)
//	if err != nil {
//		t.Error(err)
//	}
//
//	fmt.Println(hash)
//}
//
//func TestSourceByUnconstrained(t *testing.T) {
//
//	builder := model.NewSmartTransferPayloadBuilder()
//	builder.AddAssetInfo(common.NewUdtAsset(constant.UDT_HASH))
//	builder.AddFrom(constant.TEST_ADDRESS4)
//	builder.AddSmartTo(constant.TEST_ADDRESS1, big.NewInt(20))
//
//	tx, err := constant.GetMercuryApiInstance().BuildSmartTransferTransaction(builder.Build())
//	if err != nil {
//		t.Error(err)
//	}
//
//	transaction := utils.Sign(tx)
//	hash, err := constant.GetMercuryApiInstance().SendTransaction(context.Background(), transaction)
//	if err != nil {
//		t.Error(err)
//	}
//
//	fmt.Println(hash)
//}
//
//func TestActionByPayByTo(t *testing.T) {
//
//	builder := model.NewSmartTransferPayloadBuilder()
//	builder.AddAssetInfo(common.NewUdtAsset(constant.UDT_HASH))
//	builder.AddFrom(constant.TEST_ADDRESS4)
//	builder.AddSmartTo(constant.TEST_ADDRESS1, big.NewInt(20))
//
//	tx, err := constant.GetMercuryApiInstance().BuildSmartTransferTransaction(builder.Build())
//	if err != nil {
//		t.Error(err)
//	}
//
//	transaction := utils.Sign(tx)
//	hash, err := constant.GetMercuryApiInstance().SendTransaction(context.Background(), transaction)
//	if err != nil {
//		t.Error(err)
//	}
//
//	fmt.Println(hash)
//}
//
//func TestActionByPayByFrom(t *testing.T) {
//	to, _ := address.GenerateShortAddress(address.Testnet)
//
//	builder := model.NewSmartTransferPayloadBuilder()
//	builder.AddAssetInfo(common.NewUdtAsset(constant.UDT_HASH))
//	builder.AddFrom(constant.TEST_ADDRESS4)
//	builder.AddSmartTo(to.Address, big.NewInt(20))
//
//	tx, err := constant.GetMercuryApiInstance().BuildSmartTransferTransaction(builder.Build())
//	if err != nil {
//		t.Error(err)
//	}
//
//	transaction := utils.Sign(tx)
//	hash, err := constant.GetMercuryApiInstance().SendTransaction(context.Background(), transaction)
//	if err != nil {
//		t.Error(err)
//	}
//
//	fmt.Println(hash)
//}
//
//func initChequeCell() {
//	mercuryApi := constant.GetMercuryApiInstance()
//
//	builder := model.NewTransferBuilder()
//	builder.AddAssetInfo(constant.UDT_HASH)
//	builder.AddFrom([]string{constant.TEST_ADDRESS1}, source.Unconstrained)
//	builder.AddTo(constant.TEST_ADDRESS2, mode.Lend_by_from, big.NewInt(100))
//	transferPayload := builder.Build()
//	transferCompletion, err := mercuryApi.BuildTransferTransaction(transferPayload)
//	if err != nil {
//		fmt.Println(err)
//	}
//
//	tx := utils.Sign(transferCompletion)
//
//	hash, err := mercuryApi.SendTransaction(context.Background(), tx)
//	if err != nil {
//		fmt.Println(err)
//	}
//
//	var txStatus types.TransactionStatus = "pending"
//	for {
//		transaction, _ := mercuryApi.GetTransaction(context.Background(), *hash)
//		if transaction.TxStatus.Status != txStatus {
//			break
//		}
//		fmt.Println("Awaiting transaction results")
//		time.Sleep(1 * 1e9)
//
//	}
//
//	time.Sleep(60 * 1e9)
//	fmt.Printf("send hash of cheque cell transactions: %s\n", hash.String())
//}
