package indexer

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/nervosnetwork/ckb-sdk-go/v2/types"
)

type jsonCellsFilter struct {
	Script              *types.Script      `json:"script"`
	ScriptLenRange      *[2]hexutil.Uint64 `json:"script_len_range,omitempty"`
	OutputDataLenRange  *[2]hexutil.Uint64 `json:"output_data_len_range,omitempty"`
	OutputCapacityRange *[2]hexutil.Uint64 `json:"output_capacity_range,omitempty"`
	BlockRange          *[2]hexutil.Uint64 `json:"block_range,omitempty"`
}

func (r Filter) MarshalJSON() ([]byte, error) {
	toUint64Array := func(a *[2]uint64) *[2]hexutil.Uint64 {
		if a == nil {
			return nil
		} else {
			result := [2]hexutil.Uint64{}
			result[0] = hexutil.Uint64(a[0])
			result[1] = hexutil.Uint64(a[1])
			return &result
		}
	}
	jsonObj := &jsonCellsFilter{
		Script:              r.Script,
		ScriptLenRange:      toUint64Array(r.ScriptLenRange),
		OutputDataLenRange:  toUint64Array(r.OutputDataLenRange),
		OutputCapacityRange: toUint64Array(r.OutputCapacityRange),
		BlockRange:          toUint64Array(r.BlockRange),
	}
	return json.Marshal(jsonObj)
}

type liveCellAlias LiveCell
type jsonLiveCell struct {
	liveCellAlias
	BlockNumber hexutil.Uint64 `json:"block_number"`
	OutputData  *hexutil.Bytes `json:"output_data"`
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
		TxIndex:     uint(jsonObj.TxIndex),
	}
	if jsonObj.OutputData != nil {
		r.OutputData = *jsonObj.OutputData
	}
	return nil
}

type jsonTxWithCell struct {
	BlockNumber hexutil.Uint64 `json:"block_number"`
	IoIndex     hexutil.Uint   `json:"io_index"`
	IoType      IoType         `json:"io_type"`
	TxHash      types.Hash     `json:"tx_hash"`
	TxIndex     hexutil.Uint   `json:"tx_index"`
}

func (r *TxWithCell) UnmarshalJSON(input []byte) error {
	var jsonObj jsonTxWithCell
	if err := json.Unmarshal(input, &jsonObj); err != nil {
		return err
	}
	*r = TxWithCell{
		BlockNumber: uint64(jsonObj.BlockNumber),
		IoIndex:     uint(jsonObj.IoIndex),
		IoType:      jsonObj.IoType,
		TxHash:      jsonObj.TxHash,
		TxIndex:     uint(jsonObj.TxIndex),
	}
	return nil
}

type jsonTxWithCells struct {
	TxHash      types.Hash     `json:"tx_hash"`
	BlockNumber hexutil.Uint64 `json:"block_number"`
	TxIndex     hexutil.Uint   `json:"tx_index"`
	Cells       []*Cell        `json:"Cells"`
}

func (r *TxWithCells) UnmarshalJSON(input []byte) error {
	var jsonObj jsonTxWithCells
	if err := json.Unmarshal(input, &jsonObj); err != nil {
		return err
	}
	*r = TxWithCells{
		BlockNumber: uint64(jsonObj.BlockNumber),
		TxHash:      jsonObj.TxHash,
		TxIndex:     uint(jsonObj.TxIndex),
		Cells:       jsonObj.Cells,
	}
	return nil
}

func (r *Cell) UnmarshalJSON(input []byte) error {
	var jsonObj struct {
		IoType  IoType       `json:"io_type"`
		IoIndex hexutil.Uint `json:"io_index"`
	}
	tmp := []interface{}{&jsonObj.IoType, &jsonObj.IoIndex}
	if err := json.Unmarshal(input, &tmp); err != nil {
		return err
	}
	*r = Cell{
		IoType:  jsonObj.IoType,
		IoIndex: uint(jsonObj.IoIndex),
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
