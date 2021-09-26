package test

import (
	"encoding/json"
	"fmt"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/example/constant"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/common"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/req"
	"testing"
)

func TestQueryTransactionsWithCkb(t *testing.T) {

	payload := &model.QueryTransactionsPayload{
		Item:              req.NewAddressItem(constant.QUERY_TRANSACTION_KEY_PUBKEY),
		AssetInfos:        []*common.AssetInfo{common.NewCkbAsset()},
		Extra:             nil,
		BlockRange:        nil,
		PaginationRequest: model.PaginationRequest{},
	}

	//builder := model.NewQueryTransactionsPayloadBuilder()
	//builder.AddKeyAddress(&model.KeyAddress{constant.QUERY_TRANSACTION_ADDRESS})
	//
	marshal, _ := json.Marshal(payload)
	fmt.Println(string(marshal))

	transactions, err := constant.GetMercuryApiInstance().QueryTransactionsWithTransactionView(payload)
	if err != nil {
		t.Error(err)
	}

	json, err := json.Marshal(transactions)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(transactions.Response)
	fmt.Println(string(json))

}

//
//func TestQueryGenericTransactionsWithUdt(t *testing.T) {
//	builder := model.NewQueryTransactionsPayloadBuilder()
//	builder.AddKeyAddress(&model.KeyAddress{constant.QUERY_TRANSACTION_ADDRESS})
//	builder.AddUdtHash(constant.UDT_HASH)
//
//	marshal, _ := json.Marshal(builder.Build())
//	fmt.Println(string(marshal))
//
//	transactions, err := constant.GetMercuryApiInstance().QueryGenericTransactions(builder.Build())
//	if err != nil {
//		t.Error(err)
//	}
//
//	json, err := json.Marshal(transactions)
//	if err != nil {
//		t.Error(err)
//	}
//
//	fmt.Println(len(transactions.Txs))
//	fmt.Println(string(json))
//
//}
//
//func TestQueryGenericTransactionsWithAll(t *testing.T) {
//	builder := model.NewQueryTransactionsPayloadBuilder()
//	builder.AddKeyAddress(&model.KeyAddress{constant.QUERY_TRANSACTION_ADDRESS})
//	builder.AllTransactionType()
//
//	marshal, _ := json.Marshal(builder.Build())
//	fmt.Println(string(marshal))
//
//	transactions, err := constant.GetMercuryApiInstance().QueryGenericTransactions(builder.Build())
//	if err != nil {
//		t.Error(err)
//	}
//
//	json, err := json.Marshal(transactions)
//	if err != nil {
//		t.Error(err)
//	}
//
//	fmt.Println(len(transactions.Txs))
//	fmt.Println(string(json))
//
//}
//
//func TestQueryGenericTransactionsWithChequeAddress(t *testing.T) {
//	chequeAddress, err := address.GenerateChequeAddress(constant.TEST_ADDRESS0, constant.QUERY_TRANSACTION_ADDRESS)
//	assert.Nil(t, err)
//
//	builder := model.NewQueryTransactionsPayloadBuilder()
//	builder.AddNormalAddress(&model.NormalAddress{chequeAddress})
//	builder.AllTransactionType()
//
//	marshal, _ := json.Marshal(builder.Build())
//	fmt.Println(string(marshal))
//
//	transactions, err := constant.GetMercuryApiInstance().QueryGenericTransactions(builder.Build())
//	if err != nil {
//		t.Error(err)
//	}
//
//	json, err := json.Marshal(transactions)
//	if err != nil {
//		t.Error(err)
//	}
//
//	fmt.Println(len(transactions.Txs))
//	fmt.Println(string(json))
//
//}
//
//func TestQueryGenericTransactionsWithAcpAddress(t *testing.T) {
//	acpAddress, err := address.GenerateAcpAddress(constant.QUERY_TRANSACTION_ADDRESS)
//	assert.Nil(t, err)
//
//	builder := model.NewQueryTransactionsPayloadBuilder()
//	builder.AddNormalAddress(&model.NormalAddress{acpAddress})
//	builder.AllTransactionType()
//
//	marshal, _ := json.Marshal(builder.Build())
//	fmt.Println(string(marshal))
//
//	transactions, err := constant.GetMercuryApiInstance().QueryGenericTransactions(builder.Build())
//	if err != nil {
//		t.Error(err)
//	}
//
//	json, err := json.Marshal(transactions)
//	if err != nil {
//		t.Error(err)
//	}
//
//	fmt.Println(len(transactions.Txs))
//	fmt.Println(string(json))
//
//}
//
//func TestQueryGenericTransactionsWithFromBlock(t *testing.T) {
//	acpAddress, err := address.GenerateAcpAddress(constant.QUERY_TRANSACTION_ADDRESS)
//	assert.Nil(t, err)
//
//	builder := model.NewQueryTransactionsPayloadBuilder()
//	builder.AddKeyAddress(&model.KeyAddress{acpAddress})
//	builder.AllTransactionType()
//	builder.AddFromBlock(2224987)
//
//	marshal, _ := json.Marshal(builder.Build())
//	fmt.Println(string(marshal))
//
//	transactions, err := constant.GetMercuryApiInstance().QueryGenericTransactions(builder.Build())
//	if err != nil {
//		t.Error(err)
//	}
//
//	json, err := json.Marshal(transactions)
//	if err != nil {
//		t.Error(err)
//	}
//
//	fmt.Println(len(transactions.Txs))
//	fmt.Println(string(json))
//
//}
//
//func TestQueryGenericTransactionsWithToBlock(t *testing.T) {
//	acpAddress, err := address.GenerateAcpAddress(constant.QUERY_TRANSACTION_ADDRESS)
//	assert.Nil(t, err)
//
//	builder := model.NewQueryTransactionsPayloadBuilder()
//	builder.AddKeyAddress(&model.KeyAddress{acpAddress})
//	builder.AllTransactionType()
//	builder.AddToBlock(2224987)
//
//	marshal, _ := json.Marshal(builder.Build())
//	fmt.Println(string(marshal))
//
//	transactions, err := constant.GetMercuryApiInstance().QueryGenericTransactions(builder.Build())
//	if err != nil {
//		t.Error(err)
//	}
//
//	json, err := json.Marshal(transactions)
//	if err != nil {
//		t.Error(err)
//	}
//
//	fmt.Println(len(transactions.Txs))
//	fmt.Println(string(json))
//
//}
//
//func TestQueryGenericTransactionsWithFromBlockAndToBlock(t *testing.T) {
//	acpAddress, err := address.GenerateAcpAddress(constant.QUERY_TRANSACTION_ADDRESS)
//	assert.Nil(t, err)
//
//	builder := model.NewQueryTransactionsPayloadBuilder()
//	builder.AddKeyAddress(&model.KeyAddress{acpAddress})
//	builder.AllTransactionType()
//	builder.AddFromBlock(2224993)
//	builder.AddToBlock(2225023)
//
//	marshal, _ := json.Marshal(builder.Build())
//	fmt.Println(string(marshal))
//
//	transactions, err := constant.GetMercuryApiInstance().QueryGenericTransactions(builder.Build())
//	if err != nil {
//		t.Error(err)
//	}
//
//	json, err := json.Marshal(transactions)
//	if err != nil {
//		t.Error(err)
//	}
//
//	fmt.Println(len(transactions.Txs))
//	fmt.Println(string(json))
//
//}
//
//func TestQueryGenericTransactionsWithLimit(t *testing.T) {
//	acpAddress, err := address.GenerateAcpAddress(constant.QUERY_TRANSACTION_ADDRESS)
//	assert.Nil(t, err)
//
//	builder := model.NewQueryTransactionsPayloadBuilder()
//	builder.AddKeyAddress(&model.KeyAddress{acpAddress})
//	builder.AllTransactionType()
//	// default limit 50
//	builder.AddLimit(2)
//
//	marshal, _ := json.Marshal(builder.Build())
//	fmt.Println(string(marshal))
//
//	transactions, err := constant.GetMercuryApiInstance().QueryGenericTransactions(builder.Build())
//	if err != nil {
//		t.Error(err)
//	}
//
//	json, err := json.Marshal(transactions)
//	if err != nil {
//		t.Error(err)
//	}
//
//	fmt.Println(len(transactions.Txs))
//	fmt.Println(string(json))
//
//}
//
//func TestQueryGenericTransactionsWithOrder(t *testing.T) {
//	builder := model.NewQueryTransactionsPayloadBuilder()
//	builder.AddKeyAddress(&model.KeyAddress{constant.QUERY_TRANSACTION_ADDRESS})
//	builder.AllTransactionType()
//	// default order desc
//	builder.AddOrder(model.Asc)
//
//	marshal, _ := json.Marshal(builder.Build())
//	fmt.Println(string(marshal))
//
//	transactions, err := constant.GetMercuryApiInstance().QueryGenericTransactions(builder.Build())
//	if err != nil {
//		t.Error(err)
//	}
//
//	json, err := json.Marshal(transactions)
//	if err != nil {
//		t.Error(err)
//	}
//
//	fmt.Println(len(transactions.Txs))
//	fmt.Println(string(json))
//
//}
//
//func TestQueryGenericTransactionsWithOffset(t *testing.T) {
//	builder := model.NewQueryTransactionsPayloadBuilder()
//	builder.AddKeyAddress(&model.KeyAddress{constant.QUERY_TRANSACTION_ADDRESS})
//	builder.AllTransactionType()
//	builder.AddLimit(1)
//	// Offset start from 0
//	builder.AddOffset(1)
//
//	marshal, _ := json.Marshal(builder.Build())
//	fmt.Println(string(marshal))
//
//	transactions, err := constant.GetMercuryApiInstance().QueryGenericTransactions(builder.Build())
//	if err != nil {
//		t.Error(err)
//	}
//
//	json, err := json.Marshal(transactions)
//	if err != nil {
//		t.Error(err)
//	}
//
//	fmt.Println(len(transactions.Txs))
//	fmt.Println(string(json))
//
//}
