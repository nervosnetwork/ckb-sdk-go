package test

import (
	"context"
	"fmt"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/example/constant"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/action"
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
	builder.AddFrom([]string{senderAddress}, source.Unconstrained)
	builder.AddItem(chequeCellReceiverAddress, action.Lend_by_from, 100)
	transferPayload := builder.Build()
	transferCompletion, err := mercuryApi.BuildTransferTransaction(transferPayload)
	if err != nil {
		fmt.Println(err)
	}

	tx := sign(transferCompletion)

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
	builder.AddFrom([]string{chequeCellReceiverAddress}, source.Fleeting)
	builder.AddItem(receiverAddress, action.Pay_by_from, 100)
	transferPayload := builder.Build()
	transferCompletion, err := mercuryApi.BuildTransferTransaction(transferPayload)
	if err != nil {
		fmt.Println(err)
	}

	tx := sign(transferCompletion)

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
	mercuryApi := constant.GetMercuryApiInstance()
	ckbBalanceA, _ := mercuryApi.GetBalance(nil, senderAddress)
	udtBalanceA, _ := mercuryApi.GetBalance(udtHash, senderAddress)

	fmt.Printf("sender ckb balance: %+v\n", *ckbBalanceA)
	fmt.Printf("sender udt balance: %+v\n", *udtBalanceA)

	ckbBalanceB, _ := mercuryApi.GetBalance(nil, chequeCellReceiverAddress)
	udtBalanceB, _ := mercuryApi.GetBalance(udtHash, chequeCellReceiverAddress)

	fmt.Printf("sender ckb balance: %+v\n", *ckbBalanceB)
	fmt.Printf("sender udt balance: %+v\n", *udtBalanceB)
}
