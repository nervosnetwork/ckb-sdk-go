package indexer

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/nervosnetwork/ckb-sdk-go/indexer"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/example/constant"
	types2 "github.com/nervosnetwork/ckb-sdk-go/mercury/model/types"
	"github.com/nervosnetwork/ckb-sdk-go/types"
	"testing"
)

func TestCapacity(t *testing.T) {

	key := types2.SearchKeyBuilder{}
	key.AddScript(
		types2.BuildScript("0x9bd7e06f3ecf4be0f2fcd2188b23f1b9fcc88e5d4b65a8637b17723bbda3cce8", types.HashTypeType, "0x0c24d18f16e3c43272695e5db006a22cb9ddde51"))

	key.AddScriptType(indexer.ScriptTypeLock)

	mercuryApi := constant.GetMercuryApiInstance()
	cells, err := mercuryApi.GetCellsCapacity(context.Background(), key.Build())
	if err != nil {
		fmt.Println(err)
	}

	marshal, _ := json.Marshal(cells)
	fmt.Println(string(marshal))

}