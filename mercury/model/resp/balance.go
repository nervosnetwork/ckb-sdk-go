package resp

type Balance struct {
	KeyAddress    string `json:"key_address"`
	UdtHash       string `json:"udt_hash"`
	Unconstrained string `json:"unconstrained"`
	Fleeting      string `json:"fleeting"`
	Locked        string `json:"locked"`
}

type GetBalanceResponse struct {
	BlockNum string     `json:"block_num,omitempty"`
	Balances []*Balance `json:"balances"`
}
