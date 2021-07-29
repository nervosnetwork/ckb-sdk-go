package test

import (
	"encoding/json"
	"fmt"
	"github.com/nervosnetwork/ckb-sdk-go/address"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/example/constant"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQueryGenericTransactionsWithCkb(t *testing.T) {
	builder := model.NewQueryGenericTransactionsPayloadBuilder()
	builder.AddAddress(constant.QUERY_TRANSACTION_ADDRESS)

	payload, err := builder.Build()
	if err != nil {
		t.Error(err)
	}

	marshal, _ := json.Marshal(payload)
	fmt.Println(string(marshal))

	transactions, err := constant.GetMercuryApiInstance().QueryGenericTransactions(payload)
	if err != nil {
		t.Error(err)
	}

	json, err := json.Marshal(transactions)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(len(transactions.Txs))
	fmt.Println(string(json))

}

func TestQueryGenericTransactionsWithUdt(t *testing.T) {
	builder := model.NewQueryGenericTransactionsPayloadBuilder()
	builder.AddAddress(constant.QUERY_TRANSACTION_ADDRESS)
	builder.AddUdtHash(constant.UdtHash)

	payload, err := builder.Build()
	if err != nil {
		t.Error(err)
	}

	marshal, _ := json.Marshal(payload)
	fmt.Println(string(marshal))

	transactions, err := constant.GetMercuryApiInstance().QueryGenericTransactions(payload)
	if err != nil {
		t.Error(err)
	}

	json, err := json.Marshal(transactions)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(len(transactions.Txs))
	fmt.Println(string(json))

}

func TestQueryGenericTransactionsWithAll(t *testing.T) {
	builder := model.NewQueryGenericTransactionsPayloadBuilder()
	builder.AddAddress(constant.QUERY_TRANSACTION_ADDRESS)
	builder.AllTransactionType()

	payload, err := builder.Build()
	if err != nil {
		t.Error(err)
	}

	marshal, _ := json.Marshal(payload)
	fmt.Println(string(marshal))

	transactions, err := constant.GetMercuryApiInstance().QueryGenericTransactions(payload)
	if err != nil {
		t.Error(err)
	}

	json, err := json.Marshal(transactions)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(len(transactions.Txs))
	fmt.Println(string(json))

}

func TestQueryGenericTransactionsWithChequeAddress(t *testing.T) {
	chequeAddress, err := address.GenerateChequeAddress(constant.TEST_ADDRESS0, constant.QUERY_TRANSACTION_ADDRESS)
	assert.Nil(t, err)

	builder := model.NewQueryGenericTransactionsPayloadBuilder()
	builder.AddAddress(chequeAddress)
	builder.AllTransactionType()

	payload, err := builder.Build()
	if err != nil {
		t.Error(err)
	}

	marshal, _ := json.Marshal(payload)
	fmt.Println(string(marshal))

	transactions, err := constant.GetMercuryApiInstance().QueryGenericTransactions(payload)
	if err != nil {
		t.Error(err)
	}

	json, err := json.Marshal(transactions)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(len(transactions.Txs))
	fmt.Println(string(json))

}

func TestQueryGenericTransactionsWithAcpAddress(t *testing.T) {
	acpAddress, err := address.GenerateAcpAddress(constant.QUERY_TRANSACTION_ADDRESS)
	assert.Nil(t, err)

	builder := model.NewQueryGenericTransactionsPayloadBuilder()
	builder.AddAddress(acpAddress)
	builder.AllTransactionType()

	payload, err := builder.Build()
	if err != nil {
		t.Error(err)
	}

	marshal, _ := json.Marshal(payload)
	fmt.Println(string(marshal))

	transactions, err := constant.GetMercuryApiInstance().QueryGenericTransactions(payload)
	if err != nil {
		t.Error(err)
	}

	json, err := json.Marshal(transactions)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(len(transactions.Txs))
	fmt.Println(string(json))

}

func TestQueryGenericTransactionsWithFromBlock(t *testing.T) {
	acpAddress, err := address.GenerateAcpAddress(constant.QUERY_TRANSACTION_ADDRESS)
	assert.Nil(t, err)

	builder := model.NewQueryGenericTransactionsPayloadBuilder()
	builder.AddAddress(acpAddress)
	builder.AllTransactionType()
	builder.AddFromBlock(2224987)

	payload, err := builder.Build()
	if err != nil {
		t.Error(err)
	}

	marshal, _ := json.Marshal(payload)
	fmt.Println(string(marshal))

	transactions, err := constant.GetMercuryApiInstance().QueryGenericTransactions(payload)
	if err != nil {
		t.Error(err)
	}

	json, err := json.Marshal(transactions)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(len(transactions.Txs))
	fmt.Println(string(json))

}

func TestQueryGenericTransactionsWithToBlock(t *testing.T) {
	acpAddress, err := address.GenerateAcpAddress(constant.QUERY_TRANSACTION_ADDRESS)
	assert.Nil(t, err)

	builder := model.NewQueryGenericTransactionsPayloadBuilder()
	builder.AddAddress(acpAddress)
	builder.AllTransactionType()
	builder.AddToBlock(2224987)

	payload, err := builder.Build()
	if err != nil {
		t.Error(err)
	}

	marshal, _ := json.Marshal(payload)
	fmt.Println(string(marshal))

	transactions, err := constant.GetMercuryApiInstance().QueryGenericTransactions(payload)
	if err != nil {
		t.Error(err)
	}

	json, err := json.Marshal(transactions)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(len(transactions.Txs))
	fmt.Println(string(json))

}

func TestQueryGenericTransactionsWithFromBlockAndToBlock(t *testing.T) {
	acpAddress, err := address.GenerateAcpAddress(constant.QUERY_TRANSACTION_ADDRESS)
	assert.Nil(t, err)

	builder := model.NewQueryGenericTransactionsPayloadBuilder()
	builder.AddAddress(acpAddress)
	builder.AllTransactionType()
	builder.AddFromBlock(2224993)
	builder.AddToBlock(2225023)

	payload, err := builder.Build()
	if err != nil {
		t.Error(err)
	}

	marshal, _ := json.Marshal(payload)
	fmt.Println(string(marshal))

	transactions, err := constant.GetMercuryApiInstance().QueryGenericTransactions(payload)
	if err != nil {
		t.Error(err)
	}

	json, err := json.Marshal(transactions)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(len(transactions.Txs))
	fmt.Println(string(json))

}

func TestQueryGenericTransactionsWithLimit(t *testing.T) {
	acpAddress, err := address.GenerateAcpAddress(constant.QUERY_TRANSACTION_ADDRESS)
	assert.Nil(t, err)

	builder := model.NewQueryGenericTransactionsPayloadBuilder()
	builder.AddAddress(acpAddress)
	builder.AllTransactionType()
	// default limit 50
	builder.AddLimit(2)

	payload, err := builder.Build()
	if err != nil {
		t.Error(err)
	}

	marshal, _ := json.Marshal(payload)
	fmt.Println(string(marshal))

	transactions, err := constant.GetMercuryApiInstance().QueryGenericTransactions(payload)
	if err != nil {
		t.Error(err)
	}

	json, err := json.Marshal(transactions)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(len(transactions.Txs))
	fmt.Println(string(json))

}

func TestQueryGenericTransactionsWithOrder(t *testing.T) {
	builder := model.NewQueryGenericTransactionsPayloadBuilder()
	builder.AddAddress(constant.QUERY_TRANSACTION_ADDRESS)
	builder.AllTransactionType()
	// default order desc
	builder.AddOrder("asc")

	payload, err := builder.Build()
	if err != nil {
		t.Error(err)
	}

	marshal, _ := json.Marshal(payload)
	fmt.Println(string(marshal))

	transactions, err := constant.GetMercuryApiInstance().QueryGenericTransactions(payload)
	if err != nil {
		t.Error(err)
	}

	json, err := json.Marshal(transactions)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(len(transactions.Txs))
	fmt.Println(string(json))

}

func TestQueryGenericTransactionsWithOffset(t *testing.T) {
	builder := model.NewQueryGenericTransactionsPayloadBuilder()
	builder.AddAddress(constant.QUERY_TRANSACTION_ADDRESS)
	builder.AllTransactionType()
	builder.AddLimit(1)
	// Offset start from 0
	builder.AddOffset(1)

	payload, err := builder.Build()
	if err != nil {
		t.Error(err)
	}

	marshal, _ := json.Marshal(payload)
	fmt.Println(string(marshal))

	transactions, err := constant.GetMercuryApiInstance().QueryGenericTransactions(payload)
	if err != nil {
		t.Error(err)
	}

	json, err := json.Marshal(transactions)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(len(transactions.Txs))
	fmt.Println(string(json))

}
