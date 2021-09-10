package mercury

type rpcBalanceResp struct {
	KeyAddress    string `json:"key_address"`
	UdtHash       string `json:"udt_hash"`
	Unconstrained string `json:"unconstrained"`
	Fleeting      string `json:"fleeting"`
	Locked        string `json:"locked"`
}

type RpcGetBalanceResponse struct {
	Balances []*rpcBalanceResp `json:"balances"`
}
