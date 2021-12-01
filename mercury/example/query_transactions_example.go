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
	fmt.Println(len(transactions.Response))
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
	fmt.Println(len(transactions.Response))
}

func TestQueryTransactionsWithAll(t *testing.T) {
	item, _ := req.NewIdentityItemByCkb(constant.QUERY_TRANSACTION_KEY_PUBKEY)
	payload := model.NewQueryTransactionsPayloadBuilder().
		SetItem(item).
		Build()
	printJson(payload)

	transactions, err := constant.GetMercuryApiInstance().QueryTransactionsWithTransactionView(payload)
	if err != nil {
		t.Error(err)
	}
	printJson(transactions)
	fmt.Println(len(transactions.Response))
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
	fmt.Println(len(transactions.Response))
}

func TestQueryTransactionsWithCellbase(t *testing.T) {
	item, _ := req.NewAddressItem("ckt1qzda0cr08m85hc8jlnfp3zer7xulejywt49kt2rr0vthywaa50xwsqw6vjzy9kahx3lyvlgap8dp8ewd8g80pcgcexzrj")
	extra := common.ExtraFilterCellBase
	payload := model.NewQueryTransactionsPayloadBuilder().
		SetItem(item).
		AddAssetInfo(common.NewCkbAsset()).
		SetExtra((*common.ExtraFilterType)(&extra)).
		Build()

	printJson(payload)

	transactions, err := constant.GetMercuryApiInstance().QueryTransactionsWithTransactionView(payload)
	if err != nil {
		t.Error(err)
	}
	printJson(transactions)
}

func TestQueryTransactionsWithDao(t *testing.T) {
	item, _ := req.NewAddressItem(constant.TEST_ADDRESS3)
	extra := common.ExtraFilterDao
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

func TestQueryTransactionsWithFromBlockAndToBlock(t *testing.T) {
	item, _ := req.NewIdentityItemByAddress(constant.QUERY_TRANSACTION_ADDRESS)
	payload := model.NewQueryTransactionsPayloadBuilder().
		SetItem(item).
		AddBlockRange(&model.BlockRange{
			From: 2778110,
			To:   2778201,
		}).
		Build()

	printJson(payload)

	transactions, err := constant.GetMercuryApiInstance().QueryTransactionsWithTransactionView(payload)
	if err != nil {
		t.Error(err)
	}
	printJson(transactions)
	fmt.Println(len(transactions.Response))
}

func TestQueryTransactionsWithLimit(t *testing.T) {
	item, _ := req.NewIdentityItemByAddress(constant.QUERY_TRANSACTION_ADDRESS)
	payload := model.NewQueryTransactionsPayloadBuilder().
		SetItem(item).
		SetLimit(2).
		Build()

	printJson(payload)

	transactions, err := constant.GetMercuryApiInstance().QueryTransactionsWithTransactionView(payload)
	if err != nil {
		t.Error(err)
	}
	printJson(transactions)
	fmt.Println(len(transactions.Response))
}

func TestQueryTransactionsWithOrder(t *testing.T) {
	item, _ := req.NewIdentityItemByAddress(constant.QUERY_TRANSACTION_ADDRESS)
	payload := model.NewQueryTransactionsPayloadBuilder().
		SetItem(item).
		SetOrder(model.ASC).
		Build()

	printJson(payload)

	transactions, err := constant.GetMercuryApiInstance().QueryTransactionsWithTransactionView(payload)
	if err != nil {
		t.Error(err)
	}
	printJson(transactions)
	fmt.Println(len(transactions.Response))
}

func TestQueryTransactionsWithPage1(t *testing.T) {
	item, _ := req.NewIdentityItemByCkb(constant.QUERY_TRANSACTION_KEY_PUBKEY)
	builder := model.NewQueryTransactionsPayloadBuilder().
		SetItem(item).
		SetLimit(1)
	payload := builder.Build()

	//printJson(payload)

	transactions, err := constant.GetMercuryApiInstance().QueryTransactionsWithTransactionView(payload)
	if err != nil {
		t.Error(err)
	}

	printJson(transactions)

	for {
		if transactions.NextCursor == nil {
			break
		}
		payload.Pagination.Cursor = transactions.NextCursor
		transactions, err = constant.GetMercuryApiInstance().QueryTransactionsWithTransactionView(payload)
		if err != nil {
			t.Error(err)
		}

		printJson(transactions)
	}

}

func TestQueryTransactionsWithPage2(t *testing.T) {
	item, _ := req.NewIdentityItemByAddress(constant.QUERY_TRANSACTION_ADDRESS)
	payload := model.NewQueryTransactionsPayloadBuilder().
		SetItem(item).
		SetLimit(1).
		SetPageNumber(1).
		Build()

	printJson(payload)

	transactions, err := constant.GetMercuryApiInstance().QueryTransactionsWithTransactionView(payload)
	if err != nil {
		t.Error(err)
	}
	printJson(transactions)
	fmt.Println(len(transactions.Response))
}

func printJson(i interface{}) {
	marshal, err := json.Marshal(i)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(marshal))
}
