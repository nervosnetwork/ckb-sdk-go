package model

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/nervosnetwork/ckb-sdk-go/indexer"
	"github.com/nervosnetwork/ckb-sdk-go/types"
)

type SearchKeyBuilder struct {
	Script     *types.Script
	ScriptType indexer.ScriptType
	ArgsLen    uint
	Filter     *indexer.CellsFilter
}

func (builder *SearchKeyBuilder) AddScript(script *types.Script)  {
	builder.Script = script
}

func (builder *SearchKeyBuilder) AddScriptType(scriptType indexer.ScriptType) {
	builder.ScriptType = scriptType
}

func (builder *SearchKeyBuilder) AddArgsLen(argsLen uint) {
	builder.ArgsLen = argsLen
}

func (builder *SearchKeyBuilder) AddFilterScript(script *types.Script) {
	builder.initFilter()
	builder.Filter.Script = script
}

func (builder *SearchKeyBuilder) AddFilterOutputDataLenRange(inclusive, exclusive uint64)  {
	builder.initFilter()
	builder.Filter.OutputDataLenRange = &[2]uint64{inclusive, exclusive}
}

func (builder *SearchKeyBuilder) AddFilterOutputCapacityRange(inclusive, exclusive uint64)  {
	builder.initFilter()
	builder.Filter.OutputCapacityRange = &[2]uint64{inclusive, exclusive}
}

func (builder *SearchKeyBuilder) AddFilterBlockRange(inclusive, exclusive uint64)  {
	builder.initFilter()
	builder.Filter.BlockRange = &[2]uint64{inclusive, exclusive}
}

func (builder *SearchKeyBuilder) initFilter()  {
	if builder.Filter == nil {
		builder.Filter = &indexer.CellsFilter{}
	}
}

func (builder *SearchKeyBuilder) Build() *indexer.SearchKey {
	return &indexer.SearchKey{
		builder.Script,
		builder.ScriptType,
		builder.ArgsLen,
		builder.Filter,
	}
}

func BuildScript(hash string, hashType types.ScriptHashType, args string) *types.Script {
	return &types.Script{
		types.HexToHash(hash),
		types.HashTypeType,
		common.FromHex(args),
	}
}

