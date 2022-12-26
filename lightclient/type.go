package lightclient

import (
	"github.com/nervosnetwork/ckb-sdk-go/v2/indexer"
	"github.com/nervosnetwork/ckb-sdk-go/v2/types"
)

type ScriptDetail struct {
	Script      *types.Script    `json:"script"`
	ScriptType  types.ScriptType `json:"script_type"`
	BlockNumber uint64           `json:"block_number"`
}

type TransactionWithHeader struct {
	Transaction *types.Transaction `json:"transaction"`
	Header      *types.Header      `json:"header"`
}

type FetchStatus string

const (
	FetchStatusFetched  FetchStatus = "fetched"
	FetchStatusFetching FetchStatus = "fetching"
	FetchStatusAdded    FetchStatus = "added"
	FetchStatusNotFound FetchStatus = "not_found"
)

type FetchedHeader struct {
	Status    FetchStatus   `json:"status"`
	Data      *types.Header `json:"data"`
	FirstSent uint64        `json:"first_sent"`
	TimeStamp uint64        `json:"time_stamp"`
}

type FetchedTransaction struct {
	Status    FetchStatus            `json:"status"`
	Data      *TransactionWithHeader `json:"data"`
	FirstSent uint64                 `json:"first_sent"`
	TimeStamp uint64                 `json:"time_stamp"`
}

type TxWithCell struct {
	BlockNumber uint64             `json:"block_number"`
	IoIndex     uint               `json:"io_index"`
	IoType      indexer.IoType     `json:"io_type"`
	Transaction *types.Transaction `json:"transaction"`
	TxIndex     uint               `json:"tx_index"`
}

type TxWithCells struct {
	Transaction *types.Transaction `json:"transaction"`
	BlockNumber uint64             `json:"block_number"`
	TxIndex     uint               `json:"tx_index"`
	Cells       []*indexer.Cell    `json:"Cells"`
}

type TxsWithCell struct {
	LastCursor string        `json:"last_cursor"`
	Objects    []*TxWithCell `json:"objects"`
}

type TxsWithCells struct {
	LastCursor string         `json:"last_cursor"`
	Objects    []*TxWithCells `json:"objects"`
}
