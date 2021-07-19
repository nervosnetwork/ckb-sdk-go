package model

type CreateWalletPayload struct {
	Ident   string        `json:"ident"`
	Info    []*WalletInfo `json:"info"`
	FeeRate uint          `json:"fee_rate"`
}
type WalletInfo struct {
	UdtHash string `json:"udt_hash"`
}

type createWalletPayloadBuilder struct {
	Ident   string        `json:"ident"`
	Info    []*WalletInfo `json:"info"`
	FeeRate uint          `json:"fee_rate"`
}

func NewCreateWalletPayloadBuilder() *createWalletPayloadBuilder {
	// default fee rate
	return &createWalletPayloadBuilder{
		FeeRate: 1000,
	}
}

func (walletPayload *createWalletPayloadBuilder) AddIdent(ident string) {
	walletPayload.Ident = ident
}

func (walletPayload *createWalletPayloadBuilder) AddFeeRate(feeRate uint) {
	walletPayload.FeeRate = feeRate
}

func (walletPayload *createWalletPayloadBuilder) AddInfo(udtHash string) {
	walletPayload.Info = append(walletPayload.Info, &WalletInfo{
		UdtHash: udtHash,
	})
}

func (walletPayload *createWalletPayloadBuilder) Build() *CreateWalletPayload {
	return &CreateWalletPayload{
		Ident:   walletPayload.Ident,
		Info:    walletPayload.Info,
		FeeRate: walletPayload.FeeRate,
	}
}
