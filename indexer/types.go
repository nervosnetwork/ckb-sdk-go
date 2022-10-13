package indexer

import (
	"github.com/nervosnetwork/ckb-sdk-go/types"
)

type SearchOrder string
type IoType string

const (
	SearchOrderAsc  SearchOrder = "asc"
	SearchOrderDesc SearchOrder = "desc"

	IOTypeIn  IoType = "input"
	IOTypeOut IoType = "output"
)

type SearchKey struct {
	Script     *types.Script    `json:"script"`
	ScriptType types.ScriptType `json:"script_type"`
	Filter     *Filter          `json:"filter,omitempty"`
	WithData   bool             `json:"with_data"`
}

type Filter struct {
	Script              *types.Script `json:"script"`
	ScriptLenRange      *[2]uint64    `json:"script_len_range,omitempty"`
	OutputDataLenRange  *[2]uint64    `json:"output_data_len_range,omitempty"`
	OutputCapacityRange *[2]uint64    `json:"output_capacity_range,omitempty"`
	BlockRange          *[2]uint64    `json:"block_range,omitempty"`
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

type TransactionWithCells struct {
	TxHash      types.Hash `json:"tx_hash"`
	BlockNumber uint64     `json:"block_number"`
	TxIndex     uint       `json:"tx_index"`
	Cells       []Cell     `json:"Cells"`
}

type Cell struct {
	IoType  IoType `json:"io_type"`
	IoIndex uint   `json:"io_index"`
}

type Transactions struct {
	LastCursor string         `json:"last_cursor"`
	Objects    []*Transaction `json:"objects"`
}

type TransactionsWithCells struct {
	LastCursor string                  `json:"last_cursor"`
	Objects    []*TransactionWithCells `json:"objects"`
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
