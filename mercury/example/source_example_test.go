package test

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/example/constant"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/example/utils"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/common"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/mode"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/req"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/resp"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/source"
	"github.com/nervosnetwork/ckb-sdk-go/types"
	"testing"
	"time"
)

const (
	senderAddress             = constant.TEST_ADDRESS1
	chequeCellReceiverAddress = constant.TEST_ADDRESS2
	receiverAddress           = constant.TEST_ADDRESS3
)

func TestFleeting(t *testing.T) {
	printBalance()
	issuingChequeCell()
	printBalance()
	claimChequeCell()
	printBalance()
}

func issuingChequeCell() {
	mercuryApi := constant.GetMercuryApiInstance()

	builder := model.NewTransferBuilder()
	builder.AddAssetInfo(common.NewUdtAsset(constant.UDT_HASH))
	item, _ := req.NewIdentityItemByAddress(senderAddress)
	builder.AddFrom(source.Free, item)
	builder.AddTo(mode.HoldByFrom, model.NewToInfo(chequeCellReceiverAddress, model.NewU128WithU64(100)))
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

func claimChequeCell() {
	mercuryApi := constant.GetMercuryApiInstance()

	builder := model.NewTransferBuilder()
	builder.AddAssetInfo(common.NewUdtAsset(constant.UDT_HASH))
	item, _ := req.NewIdentityItemByAddress(chequeCellReceiverAddress)
	builder.AddFrom(source.Free, item)
	builder.AddTo(mode.HoldByFrom, model.NewToInfo(receiverAddress, model.NewU128WithU64(100)))
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
	fmt.Printf("claim hash of cheque cell transactions: %s\n", hash.String())
}

func printBalance() {
	ckbBalanceA := getCkbBalance(senderAddress)
	udtBalanceA := getUdtBalance(senderAddress, constant.UDT_HASH)

	fmt.Printf("sender ckb balance: %s\n", getJsonStr(ckbBalanceA))
	fmt.Printf("sender udt balance: %s\n", getJsonStr(udtBalanceA))

	ckbBalanceB := getCkbBalance(chequeCellReceiverAddress)
	udtBalanceB := getUdtBalance(chequeCellReceiverAddress, constant.UDT_HASH)

	fmt.Printf("sender ckb balance: %s\n", getJsonStr(ckbBalanceB))
	fmt.Printf("sender udt balance: %s\n", getJsonStr(udtBalanceB))
}

func getCkbBalance(addr string) *resp.GetBalanceResponse {
	builder := model.NewGetBalancePayloadBuilder()
	item, _ := req.NewIdentityItemByAddress(addr)
	builder.AddItem(item)
	builder.AddAssetInfo(common.NewCkbAsset())

	balance, _ := constant.GetMercuryApiInstance().GetBalance(builder.Build())

	return balance
}

func getUdtBalance(addr, udtHash string) *resp.GetBalanceResponse {
	builder := model.NewGetBalancePayloadBuilder()
	item, _ := req.NewIdentityItemByAddress(addr)
	builder.AddItem(item)
	builder.AddAssetInfo(common.NewUdtAsset(constant.UDT_HASH))

	balance, _ := constant.GetMercuryApiInstance().GetBalance(builder.Build())

	return balance
}

func getJsonStr(balance *resp.GetBalanceResponse) string {
	jsonStr, _ := json.Marshal(balance)
	return string(jsonStr)
}
