package model

import "github.com/nervosnetwork/ckb-sdk-go/mercury/model/common"

type AdjustAccountPayload struct {
	KeyAddress string   `json:"key_address"`
	UdtHashes  []string `json:"udt_hashes"`
	FeeRate    uint     `json:"fee_rate"`
}

type adjustAccountPayloadBuilder struct {
	KeyAddress string
	assetInfos []*common.AssetInfo
	FeeRate    uint
}

func (builder *adjustAccountPayloadBuilder) AddKeyAddress(keyAddress string) {
	builder.KeyAddress = keyAddress
}

func (builder *adjustAccountPayloadBuilder) AddFeeRate(feeRate uint) {
	builder.FeeRate = feeRate
}

func (builder *adjustAccountPayloadBuilder) AddAssetInfo(info *common.AssetInfo) {
	builder.assetInfos = append(builder.assetInfos, info)
}

func (builder *adjustAccountPayloadBuilder) Build() *AdjustAccountPayload {
	payload := &AdjustAccountPayload{
		KeyAddress: builder.KeyAddress,
		FeeRate:    builder.FeeRate,
	}
	for _, info := range builder.assetInfos {
		payload.UdtHashes = append(payload.UdtHashes, info.UdtHash)
	}

	return payload
}

func NewAdjustAccountPayloadBuilder() *adjustAccountPayloadBuilder {
	// default fee rate
	return &adjustAccountPayloadBuilder{
		FeeRate: 1000,
	}
}
