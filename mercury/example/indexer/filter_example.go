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

func TestFilterScript(t *testing.T) {
	key := types2.SearchKeyBuilder{}
	key.AddScript(
		types2.BuildScript("0x9bd7e06f3ecf4be0f2fcd2188b23f1b9fcc88e5d4b65a8637b17723bbda3cce8", types.HashTypeType, "0x0c24d18f16e3c43272695e5db006a22cb9ddde51"))

	key.AddScriptType(indexer.ScriptTypeLock)

	key.AddFilterScript(types2.BuildScript("0xc5e5dcf215925f7ef4dfaf5f4b4f105bc321c02776d6e7d52a1db3fcd9d011a4", types.HashTypeType, "0x7c7f0ee1d582c385342367792946cff3767fe02f26fd7f07dba23ae3c65b28bc"))

	mercuryApi := constant.GetMercuryApiInstance()
	cells, err := mercuryApi.GetCells(context.Background(), key.Build(), indexer.SearchOrderAsc, indexer.SearchLimit, "")
	if err != nil {
		fmt.Println(err)
	}

	marshal, _ := json.Marshal(cells)
	fmt.Println(string(marshal))
}

func TestFilterOutputCapacityRange(t *testing.T) {
	key := types2.SearchKeyBuilder{}
	key.AddScript(
		types2.BuildScript("0x9bd7e06f3ecf4be0f2fcd2188b23f1b9fcc88e5d4b65a8637b17723bbda3cce8", types.HashTypeType, "0x0c24d18f16e3c43272695e5db006a22cb9ddde51"))

	key.AddScriptType(indexer.ScriptTypeLock)

	key.AddFilterOutputCapacityRange(0, 1000000000000000000)

	mercuryApi := constant.GetMercuryApiInstance()
	cells, err := mercuryApi.GetCells(context.Background(), key.Build(), indexer.SearchOrderAsc, indexer.SearchLimit, "")
	if err != nil {
		fmt.Println(err)
	}

	marshal, _ := json.Marshal(cells)
	fmt.Println(string(marshal))
}

func TestFilterOutputDataLenRange(t *testing.T) {
	key := types2.SearchKeyBuilder{}
	key.AddScript(
		types2.BuildScript("0x9bd7e06f3ecf4be0f2fcd2188b23f1b9fcc88e5d4b65a8637b17723bbda3cce8", types.HashTypeType, "0x0c24d18f16e3c43272695e5db006a22cb9ddde51"))

	key.AddScriptType(indexer.ScriptTypeLock)

	key.AddFilterOutputDataLenRange(0, 32)

	mercuryApi := constant.GetMercuryApiInstance()
	cells, err := mercuryApi.GetCells(context.Background(), key.Build(), indexer.SearchOrderAsc, indexer.SearchLimit, "")
	if err != nil {
		fmt.Println(err)
	}

	marshal, _ := json.Marshal(cells)
	fmt.Println(string(marshal))
}

func TestFilterBlockRange(t *testing.T) {
	key := types2.SearchKeyBuilder{}
	key.AddScript(
		types2.BuildScript("0x9bd7e06f3ecf4be0f2fcd2188b23f1b9fcc88e5d4b65a8637b17723bbda3cce8", types.HashTypeType, "0x0c24d18f16e3c43272695e5db006a22cb9ddde51"))

	key.AddScriptType(indexer.ScriptTypeLock)

	key.AddFilterBlockRange(2003365, 2103365)

	mercuryApi := constant.GetMercuryApiInstance()
	cells, err := mercuryApi.GetCells(context.Background(), key.Build(), indexer.SearchOrderAsc, indexer.SearchLimit, "")
	if err != nil {
		fmt.Println(err)
	}

	marshal, _ := json.Marshal(cells)
	fmt.Println(string(marshal))
}