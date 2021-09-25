package test

import (
	"encoding/json"
	"fmt"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/types/req"
	"testing"

	"github.com/nervosnetwork/ckb-sdk-go/address"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/example/constant"
	"github.com/stretchr/testify/assert"
)

func TestQueryGenericTransactionsWithCkb(t *testing.T) {
	builder := req.NewQueryGenericTransactionsPayloadBuilder()
	builder.AddKeyAddress(&req.Identity{constant.QUERY_TRANSACTION_ADDRESS})

	marshal, _ := json.Marshal(builder.Build())
	fmt.Println(string(marshal))

	transactions, err := constant.GetMercuryApiInstance().QueryGenericTransactions(builder.Build())
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
	builder := req.NewQueryGenericTransactionsPayloadBuilder()
	builder.AddKeyAddress(&req.Identity{constant.QUERY_TRANSACTION_ADDRESS})
	builder.AddUdtHash(constant.UDT_HASH)

	marshal, _ := json.Marshal(builder.Build())
	fmt.Println(string(marshal))

	transactions, err := constant.GetMercuryApiInstance().QueryGenericTransactions(builder.Build())
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
	builder := req.NewQueryGenericTransactionsPayloadBuilder()
	builder.AddKeyAddress(&req.Identity{constant.QUERY_TRANSACTION_ADDRESS})
	builder.AllTransactionType()

	marshal, _ := json.Marshal(builder.Build())
	fmt.Println(string(marshal))

	transactions, err := constant.GetMercuryApiInstance().QueryGenericTransactions(builder.Build())
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

	builder := req.NewQueryGenericTransactionsPayloadBuilder()
	builder.AddNormalAddress(&req.Address{chequeAddress})
	builder.AllTransactionType()

	marshal, _ := json.Marshal(builder.Build())
	fmt.Println(string(marshal))

	transactions, err := constant.GetMercuryApiInstance().QueryGenericTransactions(builder.Build())
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

	builder := req.NewQueryGenericTransactionsPayloadBuilder()
	builder.AddNormalAddress(&req.Address{acpAddress})
	builder.AllTransactionType()

	marshal, _ := json.Marshal(builder.Build())
	fmt.Println(string(marshal))

	transactions, err := constant.GetMercuryApiInstance().QueryGenericTransactions(builder.Build())
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

	builder := req.NewQueryGenericTransactionsPayloadBuilder()
	builder.AddKeyAddress(&req.Identity{acpAddress})
	builder.AllTransactionType()
	builder.AddFromBlock(2224987)

	marshal, _ := json.Marshal(builder.Build())
	fmt.Println(string(marshal))

	transactions, err := constant.GetMercuryApiInstance().QueryGenericTransactions(builder.Build())
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

	builder := req.NewQueryGenericTransactionsPayloadBuilder()
	builder.AddKeyAddress(&req.Identity{acpAddress})
	builder.AllTransactionType()
	builder.AddToBlock(2224987)

	marshal, _ := json.Marshal(builder.Build())
	fmt.Println(string(marshal))

	transactions, err := constant.GetMercuryApiInstance().QueryGenericTransactions(builder.Build())
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

	builder := req.NewQueryGenericTransactionsPayloadBuilder()
	builder.AddKeyAddress(&req.Identity{acpAddress})
	builder.AllTransactionType()
	builder.AddFromBlock(2224993)
	builder.AddToBlock(2225023)

	marshal, _ := json.Marshal(builder.Build())
	fmt.Println(string(marshal))

	transactions, err := constant.GetMercuryApiInstance().QueryGenericTransactions(builder.Build())
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

	builder := req.NewQueryGenericTransactionsPayloadBuilder()
	builder.AddKeyAddress(&req.Identity{acpAddress})
	builder.AllTransactionType()
	// default limit 50
	builder.AddLimit(2)

	marshal, _ := json.Marshal(builder.Build())
	fmt.Println(string(marshal))

	transactions, err := constant.GetMercuryApiInstance().QueryGenericTransactions(builder.Build())
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
	builder := req.NewQueryGenericTransactionsPayloadBuilder()
	builder.AddKeyAddress(&req.Identity{constant.QUERY_TRANSACTION_ADDRESS})
	builder.AllTransactionType()
	// default order desc
	builder.AddOrder("asc")

	marshal, _ := json.Marshal(builder.Build())
	fmt.Println(string(marshal))

	transactions, err := constant.GetMercuryApiInstance().QueryGenericTransactions(builder.Build())
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
	builder := req.NewQueryGenericTransactionsPayloadBuilder()
	builder.AddKeyAddress(&req.Identity{constant.QUERY_TRANSACTION_ADDRESS})
	builder.AllTransactionType()
	builder.AddLimit(1)
	// Offset start from 0
	builder.AddOffset(1)

	marshal, _ := json.Marshal(builder.Build())
	fmt.Println(string(marshal))

	transactions, err := constant.GetMercuryApiInstance().QueryGenericTransactions(builder.Build())
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
