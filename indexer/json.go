package indexer

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/nervosnetwork/ckb-sdk-go/types"
)

type jsonCellsFilter struct {
	Script              *types.Script     `json:"script"`
	OutputDataLenRange  [2]hexutil.Uint64 `json:"output_data_len_range"`
	OutputCapacityRange [2]hexutil.Uint64 `json:"output_capacity_range"`
	BlockRange          [2]hexutil.Uint64 `json:"block_range"`
}

func (r CellsFilter) MarshalJSON() ([]byte, error) {
	toUint64Array := func(a *[2]uint64) [2]hexutil.Uint64 {
		result := [2]hexutil.Uint64{}
		result[0] = hexutil.Uint64(a[0])
		result[0] = hexutil.Uint64(a[1])
		return result
	}
	jsonObj := &jsonCellsFilter{
		Script:              r.Script,
		OutputDataLenRange:  toUint64Array(r.OutputDataLenRange),
		OutputCapacityRange: toUint64Array(r.OutputCapacityRange),
		BlockRange:          toUint64Array(r.BlockRange),
	}
	return json.Marshal(jsonObj)
}

type searchKeyAlias SearchKey
type jsonSearchKey struct {
	searchKeyAlias
	ArgsLen hexutil.Uint `json:"args_len,omitempty"`
}

func (r SearchKey) MarshalJSON() ([]byte, error) {
	var jsonObj = &jsonSearchKey{
		searchKeyAlias: searchKeyAlias(r),
		ArgsLen:        hexutil.Uint(r.ArgsLen),
	}
	return json.Marshal(jsonObj)
}
