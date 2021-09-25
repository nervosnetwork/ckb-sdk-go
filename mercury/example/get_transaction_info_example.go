package test

import (
	"encoding/json"
	"fmt"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/example/constant"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/resp"
	"testing"
)

func TestGetTransaction(t *testing.T) {
	transaction, err := constant.GetMercuryApiInstance().GetTransactionInfo("0x83b849ef0c2fc02eab20faad0357026d0f94b98444a4fe947a11bcbafa01b4e8")
	if err != nil {
		t.Error(err)
	}

	json, err := json.Marshal(transaction)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(json))
}

func TestGetCellBaseTransactions(t *testing.T) {
	transaction, err := constant.GetMercuryApiInstance().GetTransactionInfo("0x9d3aff02b84c0d624d9eed265997d83c067cde554e9c4128806517fccc523e2f")
	if err != nil {
		t.Error(err)
	}

	json, err := json.Marshal(transaction)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(json))
}

func TestGetDaoTransactions(t *testing.T) {
	transaction, err := constant.GetMercuryApiInstance().GetTransactionInfo("0x94f30978b3d23b6d94dc559e8dc3e1b7a185712be26bba42f31543a08470035e")
	if err != nil {
		t.Error(err)
	}

	json, err := json.Marshal(transaction)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(json))
}

func TestDaoInfo1(t *testing.T) {
	jsonStr := "{\"Dao\":{\"state\":{\"Deposit\":100},\"reward\":100}}"
	d := &resp.ExtraFilter{}
	json.Unmarshal([]byte(jsonStr), &d)

	fmt.Println(d)

}

func TestDaoInfo2(t *testing.T) {
	jsonStr := "{\"Dao\":{\"state\":{\"Withdraw\":[100,1000]},\"reward\":1000}}"
	d := &resp.ExtraFilter{}
	json.Unmarshal([]byte(jsonStr), &d)

	fmt.Println(d)

}
