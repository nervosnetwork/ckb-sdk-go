package test

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/example/constant"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/action"
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
	udtHash                   = "0xf21e7350fa9518ed3cbb008e0e8c941d7e01a12181931d5608aa366ee22228bd"
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
	ckbNode := constant.GetCkbNodeInstance()

	builder := model.NewTransferBuilder()
	builder.AddUdtHash(udtHash)
	builder.AddFromKeyAddresses([]string{senderAddress}, source.Unconstrained)
	builder.AddToKeyAddressItem(chequeCellReceiverAddress, action.Lend_by_from, 100)
	transferPayload := builder.Build()
	transferCompletion, err := mercuryApi.BuildTransferTransaction(transferPayload)
	if err != nil {
		fmt.Println(err)
	}

	tx := constant.Sign(transferCompletion)

	hash, err := ckbNode.SendTransaction(context.Background(), tx)
	if err != nil {
		fmt.Println(err)
	}

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
	fmt.Printf("send hash of cheque cell transactions: %s\n", hash.String())
}

func claimChequeCell() {
	mercuryApi := constant.GetMercuryApiInstance()
	ckbNode := constant.GetCkbNodeInstance()

	builder := model.NewTransferBuilder()
	builder.AddUdtHash(udtHash)
	builder.AddFromKeyAddresses([]string{chequeCellReceiverAddress}, source.Fleeting)
	builder.AddToKeyAddressItem(receiverAddress, action.Pay_by_from, 100)
	transferPayload := builder.Build()
	transferCompletion, err := mercuryApi.BuildTransferTransaction(transferPayload)
	if err != nil {
		fmt.Println(err)
	}

	tx := constant.Sign(transferCompletion)

	hash, err := ckbNode.SendTransaction(context.Background(), tx)
	if err != nil {
		fmt.Println(err)
	}

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
	fmt.Printf("claim hash of cheque cell transactions: %s\n", hash.String())
}

func printBalance() {
	ckbBalanceA := getCkbBalance(senderAddress)
	udtBalanceA := getUdtBalance(senderAddress, udtHash)

	fmt.Printf("sender ckb balance: %s\n", getJsonStr(ckbBalanceA))
	fmt.Printf("sender udt balance: %s\n", getJsonStr(udtBalanceA))

	ckbBalanceB := getCkbBalance(chequeCellReceiverAddress)
	udtBalanceB := getUdtBalance(chequeCellReceiverAddress, udtHash)

	fmt.Printf("sender ckb balance: %s\n", getJsonStr(ckbBalanceB))
	fmt.Printf("sender udt balance: %s\n", getJsonStr(udtBalanceB))
}

func getCkbBalance(addr string) *resp.GetBalanceResponse {
	builder := model.NewGetBalancePayloadBuilder()
	builder.AddAddress(addr)
	payload, err := builder.Build()
	if err != nil {
		panic(err)
	}
	balance, _ := constant.GetMercuryApiInstance().GetBalance(payload)

	return balance
}

func getUdtBalance(addr, udtHash string) *resp.GetBalanceResponse {
	builder := model.NewGetBalancePayloadBuilder()

	builder.AddAddress(addr)
	builder.AddUdtHash(udtHash)
	payload, err := builder.Build()
	if err != nil {
		panic(err)
	}
	balance, _ := constant.GetMercuryApiInstance().GetBalance(payload)

	return balance
}

func getJsonStr(balance *resp.GetBalanceResponse) string {
	jsonStr, _ := json.Marshal(balance)
	return string(jsonStr)
}
