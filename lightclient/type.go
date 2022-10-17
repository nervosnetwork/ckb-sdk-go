package lightclient

import "github.com/nervosnetwork/ckb-sdk-go/types"

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
