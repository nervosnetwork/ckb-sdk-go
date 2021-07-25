package test

import (
	"encoding/json"
	"fmt"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/example/constant"
	"testing"
)

func TestRegisterAddresses(t *testing.T) {
	// acp address
	addresses := []string{"ckt1qyp07nuu3fpu9rksy677uvchlmyv9ce5saes824qjq"}
	scriptHashes, err := constant.GetMercuryApiInstance().RegisterAddresses(addresses)
	if err != nil {
		t.Error(err)
	}

	json, err := json.Marshal(scriptHashes)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(string(json))
}
