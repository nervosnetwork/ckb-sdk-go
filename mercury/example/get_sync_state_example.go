package test

import (
	"encoding/json"
	"fmt"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/example/constant"
	"testing"
)

func TestSyncState(t *testing.T) {
	mercurySyncState, err := constant.GetMercuryApiInstance().GetSyncState()
	if err != nil {
		t.Error(err)
	}

	json, err := json.Marshal(mercurySyncState)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(json))
}
