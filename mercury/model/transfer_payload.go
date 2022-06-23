package model

import (
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/common"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/mode"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/req"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/source"
	"math/big"
)

type TransferPayload struct {
	AssetInfo              *common.AssetInfo      `json:"asset_info,omitempty"`
	From                   []*req.Item            `json:"from"`
	To                     *ToInfo                `json:"to"`
	OutputCapacityProvider OutputCapacityProvider `json:"output_capacity_provider,omitempty"`
	PayFee                 PayFee                 `json:"pay_fee,omitempty"`
	FeeRate                uint64                 `json:"fee_rate,omitempty"`
	Since                  *SinceConfig           `json:"since,omitempty"`
}

type PayFee string
type OutputCapacityProvider string

const (
	PayFeeFrom                 PayFee = "From"
	PayFeeTo                   PayFee = "To"
	OutputCapacityProviderFrom        = "From"
	OutputCapacityProviderTo          = "To"
)

type From struct {
	Items  []interface{} `json:"items"`
	Source source.Source `json:"source"`
}

type ToInfo struct {
	Address string   `json:"address"`
	Amount  *big.Int `json:"amount"`
}

type To struct {
	ToInfos []*ToInfo `json:"to_infos"`
	Mode    mode.Mode `json:"mode"`
}

// TODO: change method signature?
func NewToInfo(address string, amount *U128) *ToInfo {
	return &ToInfo{
		Address: address,
		Amount:  &amount.Int,
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
	AssetInfo              *common.AssetInfo
	From                   []*req.Item
	To                     *ToInfo
	OutputCapacityProvider OutputCapacityProvider
	PayFee                 PayFee
	FeeRate                uint64
	Since                  *SinceConfig
}

func (builder *transferBuilder) AddAssetInfo(assetInfo *common.AssetInfo) {
	builder.AssetInfo = assetInfo
}

// TODO: fix
func (builder *transferBuilder) AddFrom(source source.Source, items ...interface{}) {
	//builder.From = &From{
	//	Items:  items,
	//	Source: source,
	//}
}

// TODO: fix
func (builder *transferBuilder) AddTo(mode mode.Mode, toInfos ...*ToInfo) {
	//builder.To = &To{
	//	ToInfos: toInfos,
	//	Mode:    mode,
	//}
}

func (builder *transferBuilder) AddPayFee(payFee PayFee) {
	builder.PayFee = payFee
}

func (builder *transferBuilder) AddFeeRate(feeRate uint) {
	builder.FeeRate = uint64(feeRate)
}

func (builder *transferBuilder) AddSince(since *SinceConfig) {
	builder.Since = since
}

func (builder *transferBuilder) Build() *TransferPayload {
	return &TransferPayload{
		AssetInfo:              builder.AssetInfo,
		From:                   builder.From,
		To:                     builder.To,
		OutputCapacityProvider: builder.OutputCapacityProvider,
		PayFee:                 builder.PayFee,
		FeeRate:                builder.FeeRate,
		Since:                  builder.Since,
	}
}

func NewTransferBuilder() *transferBuilder {
	// default fee rate
	return &transferBuilder{
		FeeRate: 1000,
	}
}
