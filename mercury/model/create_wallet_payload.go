package model

type CreateWalletPayload struct {
	Ident string        `json:"ident"`
	Info  []*WalletInfo `json:"info"`
	Fee   uint          `json:"fee"`
}
type WalletInfo struct {
	UdtHash string `json:"udt_hash"`
}

type CreateWalletPayloadBuilder struct {
	Ident string        `json:"ident"`
	Info  []*WalletInfo `json:"info"`
	Fee   uint          `json:"fee"`
}

func (walletPayload *CreateWalletPayload) AddIdent(ident string) {
	walletPayload.Ident = ident
}

func (walletPayload *CreateWalletPayload) AddFee(fee uint) {
	walletPayload.Fee = fee
}

func (walletPayload *CreateWalletPayload) AddInfo(udtHash string) {
	walletPayload.Info = append(walletPayload.Info, &WalletInfo{
		UdtHash: udtHash,
	})
}

func (walletPayload *CreateWalletPayload) Build() *CreateWalletPayload {
	return &CreateWalletPayload{
		Ident: walletPayload.Ident,
		Info:  walletPayload.Info,
		Fee:   walletPayload.Fee,
	}
}
