package indexer

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/nervosnetwork/ckb-sdk-go/indexer"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/example/constant"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model"
	"github.com/nervosnetwork/ckb-sdk-go/types"
	"testing"
)

func TestCells(t *testing.T) {
	key := model.SearchKeyBuilder{}
	key.AddScript(
		model.BuildScript("0x9bd7e06f3ecf4be0f2fcd2188b23f1b9fcc88e5d4b65a8637b17723bbda3cce8", types.HashTypeType, "0x0c24d18f16e3c43272695e5db006a22cb9ddde51"))

	key.AddScriptType(indexer.ScriptTypeLock)

	mercuryApi := constant.GetMercuryApiInstance()
	cells, err := mercuryApi.GetCells(context.Background(), key.Build(), indexer.SearchOrderAsc, indexer.SearchLimit, "")
	if err != nil {
		fmt.Println(err)
	}

	marshal, _ := json.Marshal(cells)
	fmt.Println(string(marshal))
}
