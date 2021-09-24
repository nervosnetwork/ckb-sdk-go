package resp

import "github.com/nervosnetwork/ckb-sdk-go/mercury/model/common"

type QueryGenericTransactionsResponse struct {
	Txs        []*common.TransactionInfo  `json:"txs"`
	TotalCount uint64                     `json:"total_count"`
	NextOffset uint64                     `json:"next_offset"`
}
