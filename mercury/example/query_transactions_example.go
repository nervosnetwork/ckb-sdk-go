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

func NewIdentityQueryTransactionsPayload(pubkey string) *model.QueryTransactionsPayload {
	item, _ := req.NewIdentityItemByCkb(pubkey)
	limit := uint64(3)
	tx := &model.QueryTransactionsPayload{
		Item:       item,
		AssetInfos: []*common.AssetInfo{},
		Extra:      nil,
		BlockRange: nil,
		Pagination: model.PaginationRequest{
			Order:       model.ASC,
			Limit:       &limit,
			ReturnCount: false,
		},
	}
	return tx
}

func TestQueryTransactionsWithCkb(t *testing.T) {
	payload := NewIdentityQueryTransactionsPayload(constant.QUERY_TRANSACTION_KEY_PUBKEY)
	payload.AddAssetInfo(common.NewCkbAsset())
	printJson(payload)

	transactions, err := constant.GetMercuryApiInstance().QueryTransactionsWithTransactionView(payload)
	if err != nil {
		t.Error(err)
	}
	printJson(transactions)
}

func TestQueryTransactionsWithUdt(t *testing.T) {
	payload := NewIdentityQueryTransactionsPayload(constant.QUERY_TRANSACTION_KEY_PUBKEY)
	payload.AddAssetInfo(common.NewUdtAsset(constant.UDT_HASH))
	printJson(payload)

	transactions, err := constant.GetMercuryApiInstance().QueryTransactionsWithTransactionView(payload)
	if err != nil {
		t.Error(err)
	}
	printJson(transactions)
}

func TestQueryTransactionsInfo(t *testing.T) {
	payload := NewIdentityQueryTransactionsPayload(constant.QUERY_TRANSACTION_KEY_PUBKEY)
	payload.AddAssetInfo(common.NewCkbAsset())
	printJson(payload)

	transactions, err := constant.GetMercuryApiInstance().QueryTransactionsWithTransactionInfo(payload)
	if err != nil {
		t.Error(err)
	}
	printJson(transactions)
}

func TestQueryTransactionsWithCellbase(t *testing.T) {
	payload := NewIdentityQueryTransactionsPayload(constant.QUERY_TRANSACTION_KEY_PUBKEY)
	payload.Item, _ = req.NewAddressItem("ckt1qyqd5eyygtdmwdr7ge736zw6z0ju6wsw7rssu8fcve")
	payload.AddAssetInfo(common.NewCkbAsset())
	extra := common.CellBase
	payload.Extra = &extra

	printJson(payload)

	transactions, err := constant.GetMercuryApiInstance().QueryTransactionsWithTransactionView(payload)
	if err != nil {
		t.Error(err)
	}
	printJson(transactions)
}

func TestQueryTransactionsWithPage(t *testing.T) {
	payload := NewIdentityQueryTransactionsPayload(constant.QUERY_TRANSACTION_KEY_PUBKEY)
	payload.AddAssetInfo(common.NewCkbAsset())
	limit := uint64(1)
	payload.Pagination.Limit = &limit

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
