package resp

type AccountInfo struct {
	AccountNumber     uint64             `json:"account_number"`
	AccountAddress    string             `json:"account_address"`
}
