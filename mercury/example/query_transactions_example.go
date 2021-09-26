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
	item, _ := req.NewIdentityItemByCkb(constant.QUERY_TRANSACTION_KEY_PUBKEY)
	payload := model.NewQueryTransactionsPayloadBuilder().
		SetItem(item).
		AddAssetInfo(common.NewCkbAsset()).
		Build()
	printJson(payload)

	transactions, err := constant.GetMercuryApiInstance().QueryTransactionsWithTransactionView(payload)
	if err != nil {
		t.Error(err)
	}
	printJson(transactions)
}

func TestQueryTransactionsWithUdt(t *testing.T) {
	item, _ := req.NewIdentityItemByCkb(constant.QUERY_TRANSACTION_KEY_PUBKEY)
	payload := model.NewQueryTransactionsPayloadBuilder().
		SetItem(item).
		AddAssetInfo(common.NewUdtAsset(constant.UDT_HASH)).
		Build()
	printJson(payload)

	transactions, err := constant.GetMercuryApiInstance().QueryTransactionsWithTransactionView(payload)
	if err != nil {
		t.Error(err)
	}
	printJson(transactions)
}

func TestQueryTransactionsInfo(t *testing.T) {
	item, _ := req.NewIdentityItemByCkb(constant.QUERY_TRANSACTION_KEY_PUBKEY)
	payload := model.NewQueryTransactionsPayloadBuilder().
		SetItem(item).
		AddAssetInfo(common.NewCkbAsset()).
		Build()

	printJson(payload)

	transactions, err := constant.GetMercuryApiInstance().QueryTransactionsWithTransactionInfo(payload)
	if err != nil {
		t.Error(err)
	}
	printJson(transactions)
}

func TestQueryTransactionsWithCellbase(t *testing.T) {
	item, _ := req.NewAddressItem("ckt1qyqd5eyygtdmwdr7ge736zw6z0ju6wsw7rssu8fcve")
	extra := common.CellBase
	payload := model.NewQueryTransactionsPayloadBuilder().
		SetItem(item).
		AddAssetInfo(common.NewCkbAsset()).
		SetExtra(&extra).
		Build()

	printJson(payload)

	transactions, err := constant.GetMercuryApiInstance().QueryTransactionsWithTransactionView(payload)
	if err != nil {
		t.Error(err)
	}
	printJson(transactions)
}

func TestQueryTransactionsWithPage(t *testing.T) {
	item, _ := req.NewIdentityItemByCkb(constant.QUERY_TRANSACTION_KEY_PUBKEY)
	builder := model.NewQueryTransactionsPayloadBuilder().
		SetItem(item).
		AddAssetInfo(common.NewUdtAsset(constant.UDT_HASH)).
		SetLimit(1)
	payload := builder.Build()

	printJson(payload)

	transactions, err := constant.GetMercuryApiInstance().QueryTransactionsWithTransactionView(payload)
	if err != nil {
		t.Error(err)
	}

	printJson(transactions)

	payload.Pagination.Cursor = transactions.NextCursor
	transactions, err = constant.GetMercuryApiInstance().QueryTransactionsWithTransactionView(payload)
	if err != nil {
		t.Error(err)
	}

	printJson(transactions)

}

func printJson(i interface{}) {
	marshal, err := json.Marshal(i)
	if err != nil {
		panic(err)
	}

	fmt.Println(i)
	fmt.Println(string(marshal))
}
