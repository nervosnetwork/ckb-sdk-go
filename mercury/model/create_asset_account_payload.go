package model

type CreateAssetAccountPayload struct {
	KeyAddress string   `json:"key_address"`
	UdtHashes  []string `json:"udt_hashes"`
	FeeRate    uint     `json:"fee_rate"`
}

type createAssetAccountPayloadBuilder struct {
	KeyAddress string   `json:"key_address"`
	UdtHashes  []string `json:"udt_hashes"`
	FeeRate    uint     `json:"fee_rate"`
}

func (walletPayload *createAssetAccountPayloadBuilder) AddKeyAddress(keyAddress string) {
	walletPayload.KeyAddress = keyAddress
}

func (walletPayload *createAssetAccountPayloadBuilder) AddFeeRate(feeRate uint) {
	walletPayload.FeeRate = feeRate
}

func (walletPayload *createAssetAccountPayloadBuilder) AddUdtHash(udtHash string) {
	walletPayload.UdtHashes = append(walletPayload.UdtHashes, udtHash)
}

func (walletPayload *createAssetAccountPayloadBuilder) Build() *CreateAssetAccountPayload {
	return &CreateAssetAccountPayload{
		KeyAddress: walletPayload.KeyAddress,
		UdtHashes:  walletPayload.UdtHashes,
		FeeRate:    walletPayload.FeeRate,
	}
}

func NewCreateAssetAccountPayloadBuilder() *createAssetAccountPayloadBuilder {
	// default fee rate
	return &createAssetAccountPayloadBuilder{
		FeeRate: 1000,
	}
}
