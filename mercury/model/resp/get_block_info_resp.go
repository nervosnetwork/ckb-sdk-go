package resp

type BlockInfo struct {
	BlockNumber  uint64             `json:"block_number"`
	BlockHash    string             `json:"block_hash"`
	ParentHash   string             `json:"parent_hash"`
	Timestamp    uint64             `json:"timestamp"`
	Transactions []*TransactionInfo `json:"transactions"`
}
