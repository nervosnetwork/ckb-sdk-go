package resp

import (
	"math/big"
)

type QueryGenericTransactionsResponse struct {
	Txs        []*TransactionInfo `json:"txs"`
	TotalCount uint64             `json:"total_count"`
	NextOffset uint64             `json:"next_offset"`
}

type PaginationResponseTransactionView struct {
	Response   []TransactionViewWrapper `json:"response"`
	Count      big.Int                  `json:"count"`
	NextCursor []int                    `json:"next_cursor"`
}
type PaginationResponseTransactionInfo struct {
	Response   []TransactionInfoWrapper `json:"response"`
	Count      big.Int                  `json:"count"`
	NextCursor []int                    `json:"next_cursor"`
}
