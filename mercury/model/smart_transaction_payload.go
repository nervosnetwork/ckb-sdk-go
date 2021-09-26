package model

import (
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/common"
)

type SmartTransferPayload struct {
	AssetInfo *common.AssetInfo `json:"asset_info"`
	From      []string          `json:"from"`
	To        []*ToInfo         `json:"to"`
	PayFee    string            `json:"pay_fee,omitempty"`
	Change    string            `json:"change,omitempty"`
	FeeRate   uint64            `json:"fee_rate"`
	Since     *SinceConfig      `json:"since,omitempty"`
}

type smartTransferPayloadBuilder struct {
	AssetInfo *common.AssetInfo
	From      []string
	To        []*ToInfo
	PayFee    string
	Change    string
	FeeRate   uint64
	Since     *SinceConfig
}

func (builder *smartTransferPayloadBuilder) AddAssetInfo(info *common.AssetInfo) {
	builder.AssetInfo = info
}

func (builder *smartTransferPayloadBuilder) AddFrom(from string) {
	builder.From = append(builder.From, from)
}

func (builder *smartTransferPayloadBuilder) AddToInfo(address string, amount *U128) {
	builder.To = append(builder.To, &ToInfo{address, amount})
}

func (builder *smartTransferPayloadBuilder) AddPayFee(address string) {
	builder.PayFee = address
}

func (builder *smartTransferPayloadBuilder) AddChange(change string) {
	builder.Change = change
}

func (builder *smartTransferPayloadBuilder) AddFeeRate(feeRate uint64) {
	builder.FeeRate = feeRate
}

func (builder *smartTransferPayloadBuilder) AddSince(since *SinceConfig) {
	builder.Since = since
}

func (builder *smartTransferPayloadBuilder) Build() *SmartTransferPayload {
	return &SmartTransferPayload{
		AssetInfo: builder.AssetInfo,
		From:      builder.From,
		To:        builder.To,
		PayFee:    builder.PayFee,
		Change:    builder.Change,
		FeeRate:   builder.FeeRate,
		Since:     builder.Since,
	}
}

func NewSmartTransferPayloadBuilder() *smartTransferPayloadBuilder {
	// default fee rate
	return &smartTransferPayloadBuilder{
		From:    make([]string, 0),
		To:      make([]*ToInfo, 0),
		FeeRate: 1000,
	}
}
