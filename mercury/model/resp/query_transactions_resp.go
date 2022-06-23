package resp

type PaginationResponseTransactionView struct {
	Response   []*TransactionViewWrapper `json:"response"`
	Count      uint64                    `json:"count,omitempty"`
	NextCursor uint64                    `json:"next_cursor,omitempty"`
}

type PaginationResponseTransactionInfo struct {
	Response   []*TransactionInfoWrapper `json:"response"`
	Count      uint64                    `json:"count,omitempty"`
	NextCursor uint64                    `json:"next_cursor,omitempty"`
}
