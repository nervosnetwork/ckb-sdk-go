package test

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/nervosnetwork/ckb-sdk-go/address"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/example/constant"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/example/utils"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/action"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/source"
	"github.com/nervosnetwork/ckb-sdk-go/rpc"
	"github.com/nervosnetwork/ckb-sdk-go/types"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
	"time"
)

func TestFromKeyAddressAndToKeyAddressWithCkb(t *testing.T) {

	sendCKbTx()
	printCexCkbBalance()

	builder := model.NewCollectAssetPayloadBuilder()
	builder.AddFromKeyAddresses([]string{constant.CEX_ADDRESS}, source.Unconstrained)
	builder.AddToKeyAddressItem(constant.TEST_ADDRESS2, action.Pay_by_from)
	builder.AddFeePaidBy(constant.TEST_ADDRESS4)

	json, _ := json.Marshal(builder.Build())
	fmt.Println(string(json))

	transferCompletion, err := constant.GetMercuryApiInstance().BuildAssetCollectionTransaction(builder.Build())
	if err != nil {
		t.Error(err)
	}

	tx := utils.Sign(transferCompletion)

	hash, err := constant.GetMercuryApiInstance().SendTransaction(context.Background(), tx)
	if err != nil {
		t.Error(err)
	}

	waitTx(constant.GetMercuryApiInstance(), hash)
	printCexCkbBalance()
	fmt.Println(hash)

}

func TestFromNormalAddressesWithCkb(t *testing.T) {

	sendCKbTx()
	printCexCkbBalance()

	builder := model.NewCollectAssetPayloadBuilder()
	builder.AddFromNormalAddresses([]string{constant.CEX_ADDRESS})
	builder.AddToKeyAddressItem(constant.TEST_ADDRESS2, action.Pay_by_from)
	builder.AddFeePaidBy(constant.TEST_ADDRESS4)

	json, _ := json.Marshal(builder.Build())
	fmt.Println(string(json))

	transferCompletion, err := constant.GetMercuryApiInstance().BuildAssetCollectionTransaction(builder.Build())
	if err != nil {
		t.Error(err)
	}

	tx := utils.Sign(transferCompletion)

	hash, err := constant.GetMercuryApiInstance().SendTransaction(context.Background(), tx)
	if err != nil {
		t.Error(err)
	}

	waitTx(constant.GetMercuryApiInstance(), hash)
	printCexCkbBalance()
	fmt.Println(hash)
}

func TestToNormalAddressWithCkb(t *testing.T) {

	sendCKbTx()
	printCexCkbBalance()

	builder := model.NewCollectAssetPayloadBuilder()
	builder.AddFromKeyAddresses([]string{constant.CEX_ADDRESS}, source.Unconstrained)
	builder.AddToNormalAddressItem(constant.TEST_ADDRESS2)
	builder.AddFeePaidBy(constant.TEST_ADDRESS4)

	json, _ := json.Marshal(builder.Build())
	fmt.Println(string(json))

	transferCompletion, err := constant.GetMercuryApiInstance().BuildAssetCollectionTransaction(builder.Build())
	if err != nil {
		t.Error(err)
	}

	tx := utils.Sign(transferCompletion)

	hash, err := constant.GetMercuryApiInstance().SendTransaction(context.Background(), tx)
	if err != nil {
		t.Error(err)
	}

	waitTx(constant.GetMercuryApiInstance(), hash)
	printCexCkbBalance()
	fmt.Println(hash)
}

func TestFromNormalAddressesAndToNormalAddressWithCkb(t *testing.T) {

	sendCKbTx()
	printCexCkbBalance()

	builder := model.NewCollectAssetPayloadBuilder()
	builder.AddFromNormalAddresses([]string{constant.CEX_ADDRESS})
	builder.AddToNormalAddressItem(constant.TEST_ADDRESS2)
	builder.AddFeePaidBy(constant.TEST_ADDRESS4)

	json, _ := json.Marshal(builder.Build())
	fmt.Println(string(json))

	transferCompletion, err := constant.GetMercuryApiInstance().BuildAssetCollectionTransaction(builder.Build())
	if err != nil {
		t.Error(err)
	}

	tx := utils.Sign(transferCompletion)

	hash, err := constant.GetMercuryApiInstance().SendTransaction(context.Background(), tx)
	if err != nil {
		t.Error(err)
	}

	waitTx(constant.GetMercuryApiInstance(), hash)
	printCexCkbBalance()
	fmt.Println(hash)
}

func TestFromKeyAddressAndToKeyAddressWithUdt(t *testing.T) {

	printSenderUdtBalance()
	sendLendByFrom()
	printSenderUdtBalance()
	printCexUdtBalance()

	builder := model.NewCollectAssetPayloadBuilder()
	builder.AddUdtHash(constant.UDT_HASH)
	builder.AddFromKeyAddresses([]string{constant.CEX_ADDRESS}, source.Fleeting)
	builder.AddToKeyAddressItem(constant.TEST_ADDRESS3, action.Pay_by_to)
	builder.AddFeePaidBy(constant.TEST_ADDRESS4)

	requestJson, _ := json.Marshal(builder.Build())
	fmt.Println(string(requestJson))

	transferCompletion, err := constant.GetMercuryApiInstance().BuildAssetCollectionTransaction(builder.Build())
	if err != nil {
		t.Error(err)
	}

	responseJson, _ := json.Marshal(transferCompletion)
	fmt.Println(string(responseJson))

	tx := utils.Sign(transferCompletion)

	hash, err := constant.GetMercuryApiInstance().SendTransaction(context.Background(), tx)
	if err != nil {
		t.Error(err)
	}

	waitTx(constant.GetMercuryApiInstance(), hash)
	printCexUdtBalance()
	fmt.Println(hash)
}

func TestFromNormalAddressesWithUdt(t *testing.T) {

	printSenderUdtBalance()
	sendLendByFrom()
	printSenderUdtBalance()
	printCexUdtBalance()

	chequeAddress, err := address.GenerateChequeAddress(constant.TEST_ADDRESS0, constant.CEX_ADDRESS)
	assert.Nil(t, err)

	builder := model.NewCollectAssetPayloadBuilder()
	builder.AddUdtHash(constant.UDT_HASH)
	builder.AddFromNormalAddresses([]string{chequeAddress})
	builder.AddToKeyAddressItem(constant.TEST_ADDRESS3, action.Pay_by_to)
	builder.AddFeePaidBy(constant.TEST_ADDRESS4)

	requestJson, _ := json.Marshal(builder.Build())
	fmt.Println(string(requestJson))

	transferCompletion, err := constant.GetMercuryApiInstance().BuildAssetCollectionTransaction(builder.Build())
	if err != nil {
		t.Error(err)
	}

	responseJson, _ := json.Marshal(transferCompletion)
	fmt.Println(string(responseJson))

	tx := utils.Sign(transferCompletion)

	hash, err := constant.GetMercuryApiInstance().SendTransaction(context.Background(), tx)
	if err != nil {
		t.Error(err)
	}

	waitTx(constant.GetMercuryApiInstance(), hash)
	printCexUdtBalance()
	fmt.Println(hash)
}

func TestToNormalAddressWithUdt(t *testing.T) {

	printSenderUdtBalance()
	sendLendByFrom()
	printSenderUdtBalance()
	printCexUdtBalance()
	printreceiverUdtBalance()

	acpAddress, err := address.GenerateAcpAddress(constant.TEST_ADDRESS4)
	assert.Nil(t, err)

	builder := model.NewCollectAssetPayloadBuilder()
	builder.AddUdtHash(constant.UDT_HASH)
	builder.AddFromKeyAddresses([]string{constant.CEX_ADDRESS}, source.Fleeting)
	builder.AddToNormalAddressItem(acpAddress)
	builder.AddFeePaidBy(constant.TEST_ADDRESS4)

	requestJson, _ := json.Marshal(builder.Build())
	fmt.Println(string(requestJson))

	transferCompletion, err := constant.GetMercuryApiInstance().BuildAssetCollectionTransaction(builder.Build())
	if err != nil {
		t.Error(err)
	}

	responseJson, _ := json.Marshal(transferCompletion)
	fmt.Println(string(responseJson))

	tx := utils.Sign(transferCompletion)

	hash, err := constant.GetMercuryApiInstance().SendTransaction(context.Background(), tx)
	if err != nil {
		t.Error(err)
	}

	waitTx(constant.GetMercuryApiInstance(), hash)
	printCexUdtBalance()
	printreceiverUdtBalance()
	fmt.Println(hash)
}

func TestFromNormalAddressesAndToNormalAddressWithUdt(t *testing.T) {

	printSenderUdtBalance()
	sendLendByFrom()
	printSenderUdtBalance()
	printCexUdtBalance()

	acpAddress, err := address.GenerateAcpAddress(constant.TEST_ADDRESS4)
	assert.Nil(t, err)

	chequeAddress, err := address.GenerateChequeAddress(constant.TEST_ADDRESS0, constant.CEX_ADDRESS)
	assert.Nil(t, err)

	builder := model.NewCollectAssetPayloadBuilder()
	builder.AddUdtHash(constant.UDT_HASH)
	builder.AddFromNormalAddresses([]string{chequeAddress})
	builder.AddToNormalAddressItem(acpAddress)
	builder.AddFeePaidBy(constant.TEST_ADDRESS4)

	requestJson, _ := json.Marshal(builder.Build())
	fmt.Println(string(requestJson))

	transferCompletion, err := constant.GetMercuryApiInstance().BuildAssetCollectionTransaction(builder.Build())
	if err != nil {
		t.Error(err)
	}

	responseJson, _ := json.Marshal(transferCompletion)
	fmt.Println(string(responseJson))

	tx := utils.Sign(transferCompletion)

	hash, err := constant.GetMercuryApiInstance().SendTransaction(context.Background(), tx)
	if err != nil {
		t.Error(err)
	}

	waitTx(constant.GetMercuryApiInstance(), hash)
	printCexUdtBalance()
	fmt.Println(hash)
}

func sendLendByFrom() {
	mercuryApi := constant.GetMercuryApiInstance()

	builder := model.NewTransferBuilder()
	builder.AddUdtHash(constant.UDT_HASH)
	builder.AddFromKeyAddresses([]string{constant.TEST_ADDRESS0}, source.Unconstrained)
	builder.AddToKeyAddressItem(constant.CEX_ADDRESS, action.Lend_by_from, big.NewInt(100))
	transferPayload := builder.Build()
	transferCompletion, err := mercuryApi.BuildTransferTransaction(transferPayload)
	if err != nil {
		fmt.Println(err)
	}

	responseJson, _ := json.Marshal(transferCompletion)
	fmt.Println(string(responseJson))

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

func sendCKbTx() {
	mercuryApi := constant.GetMercuryApiInstance()

	transferPayload := getTransferPayload(constant.TEST_ADDRESS0, constant.CEX_ADDRESS, "", action.Pay_by_from)
	transferCompletion, err := mercuryApi.BuildTransferTransaction(transferPayload)
	if err != nil {
		fmt.Println(err)
	}

	tx := utils.Sign(transferCompletion)

	hash, err := mercuryApi.SendTransaction(context.Background(), tx)
	if err != nil {
		fmt.Println(err)
	}

	waitTx(mercuryApi, hash)
	fmt.Printf("send hash of transactions: %s\n", hash.String())
}

func waitTx(ckbNode rpc.Client, hash *types.Hash) {
	var txStatus types.TransactionStatus = "pending"
	for {
		transaction, _ := ckbNode.GetTransaction(context.Background(), *hash)
		if transaction.TxStatus.Status != txStatus {
			break
		}
		fmt.Println("Awaiting transaction results")
		time.Sleep(1 * 1e9)

	}

	time.Sleep(60 * 1e9)
}

func printCexCkbBalance() {
	ckbBalance := getCkbBalance(constant.CEX_ADDRESS)
	fmt.Printf("cex ckb balance: %s\n", getJsonStr(ckbBalance))
}

func printCexUdtBalance() {
	ckbBalance := getUdtBalance(constant.CEX_ADDRESS, constant.UDT_HASH)
	fmt.Printf("cex udt balance: %s\n", getJsonStr(ckbBalance))
}

func printreceiverUdtBalance() {
	ckbBalance := getUdtBalance(constant.TEST_ADDRESS3, constant.UDT_HASH)
	fmt.Printf("receiver udt balance: %s\n", getJsonStr(ckbBalance))
}

func printSenderUdtBalance() {
	ckbBalance := getUdtBalance(constant.TEST_ADDRESS0, constant.UDT_HASH)
	fmt.Printf("sender udt balance: %s\n", getJsonStr(ckbBalance))
}
