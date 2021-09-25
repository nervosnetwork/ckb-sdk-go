package resp

type BlockInfoResponse struct {
	BlockNumber     uint64             `json:"block_number"`
	BlockHash       string             `json:"block_hash"`
	ParentBlockHash string             `json:"parent_block_hash"`
	Timestamp       uint64             `json:"timestamp"`
	Transactions    []*TransactionInfo `json:"transactions"`
}
