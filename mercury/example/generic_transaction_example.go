package test

import (
	"encoding/json"
	"fmt"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/example/constant"
	"testing"
)

func TestGetGenericTransaction(t *testing.T) {
	transaction, err := constant.GetMercuryApiInstance().GetGenericTransaction("0x83b849ef0c2fc02eab20faad0357026d0f94b98444a4fe947a11bcbafa01b4e8")
	if err != nil {
		t.Error(err)
	}

	json, err := json.Marshal(transaction)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(json))
}
