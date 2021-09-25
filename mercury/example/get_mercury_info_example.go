package test

import (
	"encoding/json"
	"fmt"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/example/constant"
	"testing"
)

func TestMercuryInfo(t *testing.T) {
	mercuryInfo, err := constant.GetMercuryApiInstance().GetMercuryInfo()
	if err != nil {
		t.Error(err)
	}

	json, err := json.Marshal(mercuryInfo)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(json))
}
