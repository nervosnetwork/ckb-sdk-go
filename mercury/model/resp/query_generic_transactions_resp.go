package resp

type QueryGenericTransactionsResponse struct {
	Txs        []*GenericTransactionResp `json:"txs"`
	TotalCount uint64                    `json:"total_count"`
	NextOffset uint64                    `json:"next_offset"`
}
