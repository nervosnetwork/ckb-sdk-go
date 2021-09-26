package test

import (
	"encoding/json"
	"fmt"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/example/constant"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/common"
	"github.com/nervosnetwork/ckb-sdk-go/types"
	"testing"
)

func TestGetSpentTransactionView(t *testing.T) {
	payload := &model.GetSpentTransactionPayload{
		OutPoint: common.OutPoint{
			types.HexToHash("0xb2e952a30656b68044e1d5eed69f1967347248967785449260e3942443cbeece"),
			01,
		},
	}

	transactionView, err := constant.GetMercuryApiInstance().GetSpentTransactionWithTransactionView(payload)
	if err != nil {
		t.Error(err)
	}

	json, err := json.Marshal(transactionView)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(json))
}

func TestGetSpentTransactionInfo(t *testing.T) {
	payload := &model.GetSpentTransactionPayload{
		OutPoint: common.OutPoint{
			types.HexToHash("0xb2e952a30656b68044e1d5eed69f1967347248967785449260e3942443cbeece"),
			01,
		},
	}

	transactionInfo, err := constant.GetMercuryApiInstance().GetSpentTransactionWithTransactionInfo(payload)
	if err != nil {
		t.Error(err)
	}

	json, err := json.Marshal(transactionInfo)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(json))
}
