package model

import (
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/common"
)

type SimpleTransferPayload struct {
	AssetInfo *common.AssetInfo `json:"asset_info"`
	From      []string          `json:"from"`
	To        []*ToInfo         `json:"to"`
	FeeRate   uint64            `json:"fee_rate"`
	Since     *SinceConfig      `json:"since,omitempty"`
}

type simpleTransferPayloadBuilder struct {
	AssetInfo *common.AssetInfo
	From      []string
	To        []*ToInfo
	PayFee    string
	Change    string
	FeeRate   uint64
	Since     *SinceConfig
}

func (builder *simpleTransferPayloadBuilder) AddAssetInfo(info *common.AssetInfo) {
	builder.AssetInfo = info
}

func (builder *simpleTransferPayloadBuilder) AddFrom(from string) {
	builder.From = append(builder.From, from)
}

func (builder *simpleTransferPayloadBuilder) AddToInfo(address string, amount *U128) {
	builder.To = append(builder.To, &ToInfo{address, &amount.Int})
}

func (builder *simpleTransferPayloadBuilder) AddPayFee(address string) {
	builder.PayFee = address
}

func (builder *simpleTransferPayloadBuilder) AddChange(change string) {
	builder.Change = change
}

func (builder *simpleTransferPayloadBuilder) AddFeeRate(feeRate uint64) {
	builder.FeeRate = feeRate
}

func (builder *simpleTransferPayloadBuilder) AddSince(since *SinceConfig) {
	builder.Since = since
}

func (builder *simpleTransferPayloadBuilder) Build() *SimpleTransferPayload {
	return &SimpleTransferPayload{
		AssetInfo: builder.AssetInfo,
		From:      builder.From,
		To:        builder.To,
		FeeRate:   builder.FeeRate,
		Since:     builder.Since,
	}
}

func NewSimpleTransferPayloadBuilder() *simpleTransferPayloadBuilder {
	// default fee rate
	return &simpleTransferPayloadBuilder{
		From:    make([]string, 0),
		To:      make([]*ToInfo, 0),
		FeeRate: 1000,
	}
}
