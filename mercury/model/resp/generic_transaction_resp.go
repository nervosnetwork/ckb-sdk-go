package resp

type GetGenericTransactionResponse struct {
	Transaction     GenericTransactionResp `json:"transaction"`
	Status          string                 `json:"status"`
	BlockHash       string                 `json:"block_hash"`
	BlockNumber     uint64                 `json:"block_number"`
	ConfirmedNumber uint64                 `json:"confirmed_number"`
}

type GenericTransactionResp struct {
	TxHash     string           `json:"tx_hash"`
	Operations []*OperationResp `json:"operations"`
}

type OperationResp struct {
	Id            uint       `json:"id"`
	KeyAddress    string     `json:"key_address"`
	NormalAddress string     `json:"normal_address"`
	Amount        AmountResp `json:"amount"`
}

type AmountResp struct {
	Value   string `json:"value"`
	UdtHash string `json:"udt_hash"`
	Status  string `json:"status"`
}
