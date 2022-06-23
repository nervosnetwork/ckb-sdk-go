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

type liveCellAlias LiveCell
type jsonLiveCell struct {
	liveCellAlias
	BlockNumber hexutil.Uint64 `json:"block_number"`
	OutputData  hexutil.Bytes  `json:"output_data"`
	TxIndex     hexutil.Uint   `json:"tx_index"`
}

func (r *LiveCell) UnmarshalJSON(input []byte) error {
	var jsonObj jsonLiveCell
	if err := json.Unmarshal(input, &jsonObj); err != nil {
		return err
	}
	*r = LiveCell{
		BlockNumber: uint64(jsonObj.BlockNumber),
		OutPoint:    jsonObj.OutPoint,
		Output:      jsonObj.Output,
		OutputData:  jsonObj.OutputData,
		TxIndex:     uint(jsonObj.TxIndex),
	}
	return nil
}

type jsonTransaction struct {
	BlockNumber hexutil.Uint64 `json:"block_number"`
	IoIndex     hexutil.Uint   `json:"io_index"`
	IoType      IoType         `json:"io_type"`
	TxHash      types.Hash     `json:"tx_hash"`
	TxIndex     hexutil.Uint   `json:"tx_index"`
}

func (r *Transaction) UnmarshalJSON(input []byte) error {
	var jsonObj jsonTransaction
	if err := json.Unmarshal(input, &jsonObj); err != nil {
		return err
	}
	*r = Transaction{
		BlockNumber: uint64(jsonObj.BlockNumber),
		IoIndex:     uint(jsonObj.IoIndex),
		IoType:      jsonObj.IoType,
		TxHash:      jsonObj.TxHash,
		TxIndex:     uint(jsonObj.TxIndex),
	}
	return nil
}

func (r *Capacity) UnmarshalJSON(input []byte) error {
	var jsonObj struct {
		Capacity    hexutil.Uint64 `json:"capacity"`
		BlockHash   types.Hash     `json:"block_hash"`
		BlockNumber hexutil.Uint64 `json:"block_number"`
	}
	if err := json.Unmarshal(input, &jsonObj); err != nil {
		return err
	}
	*r = Capacity{
		Capacity:    uint64(jsonObj.Capacity),
		BlockHash:   jsonObj.BlockHash,
		BlockNumber: uint64(jsonObj.BlockNumber),
	}
	return nil
}

func (r *TipHeader) UnmarshalJSON(input []byte) error {
	var jsonObj struct {
		BlockHash   types.Hash     `json:"block_hash"`
		BlockNumber hexutil.Uint64 `json:"block_number"`
	}
	if err := json.Unmarshal(input, &jsonObj); err != nil {
		return err
	}
	*r = TipHeader{
		BlockHash:   jsonObj.BlockHash,
		BlockNumber: uint64(jsonObj.BlockNumber),
	}
	return nil
}
