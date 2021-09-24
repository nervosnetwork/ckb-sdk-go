package resp

import "github.com/nervosnetwork/ckb-sdk-go/mercury/model/common"

type BlockInfoResponse struct {
	BlockNumber     uint64                     `json:"block_number"`
	BlockHash       string                     `json:"block_hash"`
	ParentBlockHash string                     `json:"parent_block_hash"`
	Timestamp       uint64                     `json:"timestamp"`
	Transactions    []*common.TransactionInfo  `json:"transactions"`
}
