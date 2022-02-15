package resp

type AccountType string

const (
	Acp    AccountType = "Acp"
	PwLock AccountType = "PwLock"
)
type AccountInfo struct {
	AccountNumber     uint64             `json:"account_number"`
	AccountAddress    string             `json:"account_address"`
	AccountType       AccountType        `json:"account_type"`
}
