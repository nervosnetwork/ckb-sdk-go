package indexer

import (
	"github.com/nervosnetwork/ckb-sdk-go/types"
)

type ScriptType string
type SearchOrder string
type IoType string

const (
	ScriptTypeLock ScriptType = "lock"
	ScriptTypeType ScriptType = "type"

	SearchOrderAsc  SearchOrder = "asc"
	SearchOrderDesc SearchOrder = "desc"

	IOTypeIn  IoType = "input"
	IOTypeOut IoType = "output"
)

type SearchKey struct {
	Script     *types.Script `json:"script"`
	ScriptType ScriptType    `json:"script_type"`
	ArgsLen    uint          `json:"args_len,omitempty"`
	Filter     *CellsFilter  `json:"filter,omitempty"`
}

type CellsFilter struct {
	Script              *types.Script `json:"script"`
	OutputDataLenRange  *[2]uint64    `json:"output_data_len_range"`
	OutputCapacityRange *[2]uint64    `json:"output_capacity_range"`
	BlockRange          *[2]uint64    `json:"block_range"`
}

type LiveCell struct {
	BlockNumber uint64            `json:"block_number"`
	OutPoint    *types.OutPoint   `json:"out_point"`
	Output      *types.CellOutput `json:"output"`
	OutputData  []byte            `json:"output_data"`
	TxIndex     uint              `json:"tx_index"`
}

type LiveCells struct {
	LastCursor string      `json:"last_cursor"`
	Objects    []*LiveCell `json:"objects"`
}

type Transaction struct {
	BlockNumber uint64     `json:"block_number"`
	IoIndex     uint       `json:"io_index"`
	IoType      IoType     `json:"io_type"`
	TxHash      types.Hash `json:"tx_hash"`
	TxIndex     uint       `json:"tx_index"`
}

type Transactions struct {
	LastCursor string         `json:"last_cursor"`
	Objects    []*Transaction `json:"objects"`
}

type TipHeader struct {
	BlockHash   types.Hash `json:"block_hash"`
	BlockNumber uint64     `json:"block_number"`
}

type Capacity struct {
	Capacity    uint64     `json:"capacity"`
	BlockHash   types.Hash `json:"block_hash"`
	BlockNumber uint64     `json:"block_number"`
}