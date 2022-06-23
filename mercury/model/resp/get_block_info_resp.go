package resp

import "github.com/nervosnetwork/ckb-sdk-go/types"

type BlockInfo struct {
	BlockNumber  uint64             `json:"block_number"`
	BlockHash    types.Hash         `json:"block_hash"`
	ParentHash   types.Hash         `json:"parent_hash"`
	Timestamp    uint64             `json:"timestamp"`
	Transactions []*TransactionInfo `json:"transactions"`
}
