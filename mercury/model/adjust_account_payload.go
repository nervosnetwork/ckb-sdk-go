package model

import (
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/common"
)

type BuildAdjustAccountPayload struct {
	Item          interface{}       `json:"item"`
	From          []interface{}     `json:"from"`
	AssetInfo     *common.AssetInfo `json:"asset_info"`
	AccountNumber uint              `json:"account_number"`
	ExtraCKB      *uint             `json:"extra_ckb"`
	FeeRate       uint              `json:"fee_rate"`
}

func NewBuildAdjustAccountPayload() *BuildAdjustAccountPayload {
	return &BuildAdjustAccountPayload{
		FeeRate: 1000,
		From:    []interface{}{},
	}
}

func (a *BuildAdjustAccountPayload) AddItemToFrom(item interface{}) {
	a.From = append(a.From, item)
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

//func (builder *adjustAccountPayloadBuilder) Build() *BuildAdjustAccountPayload {
//	payload := &BuildAdjustAccountPayload{
//		KeyAddress: builder.KeyAddress,
//		FeeRate:    builder.FeeRate,
//	}
//	for _, info := range builder.assetInfos {
//		payload.UdtHashes = append(payload.UdtHashes, info.UdtHash)
//	}
//
//	return payload
//}

func NewAdjustAccountPayloadBuilder() *adjustAccountPayloadBuilder {
	// default fee rate
	return &adjustAccountPayloadBuilder{
		FeeRate: 1000,
	}
}
