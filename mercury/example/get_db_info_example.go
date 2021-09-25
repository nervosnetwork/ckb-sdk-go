package test

import (
	"encoding/json"
	"fmt"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/example/constant"
	"testing"
)

func TestGetDbInfo(t *testing.T) {
	dbInfo, err := constant.GetMercuryApiInstance().GetDbInfo()
	if err != nil {
		t.Error(err)
	}

	json, err := json.Marshal(dbInfo)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(json))
}
