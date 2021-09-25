package resp

import (
	. "github.com/nervosnetwork/ckb-sdk-go/mercury/model/common"
)

type GetBlockInfoResponse struct {
	BlockNumber     BlockNumber         `json:"block_number"`
	BlockHash       H256                `json:"block_hash"`
	ParentBlockHash H256                `json:"parent_block_hash"`
	Timestamp       uint64              `json:"timestamp"`
	Transactions    []*TransactionInfo  `json:"transactions"`
}
