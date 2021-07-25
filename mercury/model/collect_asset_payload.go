package model

type CollectAssetPayload struct {
	UdtHash     string      `json:"udt_hash,omitempty"`
	FromAddress interface{} `json:"from_address,omitempty"`
	To          ToAddress   `json:"to,omitempty"`
	FeePaidBy   string      `json:"fee_paid_by,omitempty"`
	FeeRate     uint        `json:"fee_rate,omitempty"`
}

type collectAssetPayloadBuilder struct {
	UdtHash     string
	FromAddress interface{}
	To          ToAddress
	FeePaidBy   string
	FeeRate     uint
}

func (builder *collectAssetPayloadBuilder) AddUdtHash(udtHash string) {
	builder.UdtHash = udtHash
}

func (builder *collectAssetPayloadBuilder) AddFromKeyAddresses(keyAddr []string, source string) {
	builder.FromAddress = &FromKeyAddresses{
		KeyAddresses: &keyAddresses{
			KeyAddresses: keyAddr,
			Source:       source,
		},
	}
}

func (builder *collectAssetPayloadBuilder) AddFromNormalAddresses(normalAddress []string) {
	builder.FromAddress = &FromNormalAddresses{
		NormalAddresses: normalAddress,
	}
}

func (builder *collectAssetPayloadBuilder) AddToKeyAddressItem(addr, action string) {
	builder.To = &ToKeyAddress{
		KeyAddress: &keyAddress{
			KeyAddress: addr,
			Action:     action,
		},
	}
}

func (builder *collectAssetPayloadBuilder) AddToNormalAddressItem(addr string) {
	builder.To = &ToNormalAddress{
		NormalAddress: addr,
	}
}

func (builder *collectAssetPayloadBuilder) AddFeePaidBy(feePaidBy string) {
	builder.FeePaidBy = feePaidBy
}

func (builder *collectAssetPayloadBuilder) AddFeeRate(feeRate uint) {
	builder.FeeRate = feeRate
}

func (builder *collectAssetPayloadBuilder) Build() *CollectAssetPayload {
	return &CollectAssetPayload{
		UdtHash:     builder.UdtHash,
		FromAddress: builder.FromAddress,
		To:          builder.To,
		FeePaidBy:   builder.FeePaidBy,
		FeeRate:     builder.FeeRate,
	}
}

func NewCollectAssetPayloadBuilder() *collectAssetPayloadBuilder {
	// default fee rate
	return &collectAssetPayloadBuilder{
		FeeRate: 1000,
	}
}
