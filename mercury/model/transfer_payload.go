package model

import (
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/common"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/mode"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/source"
)

type TransferPayload struct {
	AssetInfo *common.AssetInfo `json:"asset_info,omitempty"`
	From      *From             `json:"from"`
	To        *To               `json:"to"`
	PayFee    string            `json:"pay_fee,omitempty"`
	Change    string            `json:"change,omitempty"`
	FeeRate   uint              `json:"fee_rate"`
	Since     *SinceConfig      `json:"since,omitempty"`
}

type From struct {
	Items  []interface{} `json:"items"`
	Source source.Source `json:"source"`
}

type ToInfo struct {
	Address string `json:"address"`
	Amount  *U128  `json:"amount"`
}

type To struct {
	ToInfos []*ToInfo `json:"to_infos"`
	Mode    mode.Mode `json:"mode"`
}

func NewToInfo(address string, amount *U128) *ToInfo {
	return &ToInfo{
		Address: address,
		Amount:  amount,
	}
}

type SinceConfig struct {
	Flag  SinceFlag `json:"flag"`
	Type  SinceType `json:"type_"`
	Value uint64    `json:"value"`
}

type SinceFlag string

const (
	Relative SinceFlag = "Relative"
	Absolute SinceFlag = "Absolute"
)

type SinceType string

const (
	BlockNumber SinceType = "BlockNumber"
	EpochNumber SinceType = "EpochNumber"
	Timestamp   SinceType = "Timestamp"
)

type transferBuilder struct {
	AssetInfo *common.AssetInfo
	From      *From
	To        *To
	PayFee    string
	Change    string
	FeeRate   uint
	Since     *SinceConfig
}

func (builder *transferBuilder) AddAssetInfo(assetInfo *common.AssetInfo) {
	builder.AssetInfo = assetInfo
}

func (builder *transferBuilder) AddFrom(source source.Source, items ...interface{}) {
	builder.From = &From{
		Items:  items,
		Source: source,
	}
}

func (builder *transferBuilder) AddTo(mode mode.Mode, toInfos ...*ToInfo) {
	builder.To = &To{
		ToInfos: toInfos,
		Mode:    mode,
	}
}

func (builder *transferBuilder) AddPayFee(address string) {
	builder.PayFee = address
}

func (builder *transferBuilder) AddChange(address string) {
	builder.Change = address
}

func (builder *transferBuilder) AddFeeRate(feeRate uint) {
	builder.FeeRate = feeRate
}

func (builder *transferBuilder) AddSince(since *SinceConfig) {
	builder.Since = since
}

func (builder *transferBuilder) Build() *TransferPayload {
	return &TransferPayload{
		AssetInfo: builder.AssetInfo,
		From:      builder.From,
		To:        builder.To,
		PayFee:    builder.PayFee,
		Change:    builder.Change,
		FeeRate:   builder.FeeRate,
		Since:     builder.Since,
	}
}

func NewTransferBuilder() *transferBuilder {
	// default fee rate
	return &transferBuilder{
		FeeRate: 1000,
	}
}
