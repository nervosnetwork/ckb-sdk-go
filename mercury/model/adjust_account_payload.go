package model

import (
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/common"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/req"
)

type BuildAdjustAccountPayload struct {
	Item          *req.Item         `json:"item"`
	From          []*req.Item       `json:"from"`
	AssetInfo     *common.AssetInfo `json:"asset_info"`
	AccountNumber uint32            `json:"account_number,omitempty"`
	ExtraCKB      uint64            `json:"extra_ckb,omitempty"`
	FeeRate       uint64            `json:"fee_rate,omitempty"`
}

func NewBuildAdjustAccountPayload() *BuildAdjustAccountPayload{
	return &BuildAdjustAccountPayload{
		From:          make([]*req.Item, 0),
		AssetInfo:     common.NewCkbAsset(),
	}
}

func (r *BuildAdjustAccountPayload) SetItem(item *req.Item) {
	r.Item = item
}

func (r *BuildAdjustAccountPayload) AddFrom(item *req.Item) {
	r.From = append(r.From, item)
}

func (r *BuildAdjustAccountPayload) SetAssetInfo(assetInfo *common.AssetInfo) {
	r.AssetInfo = assetInfo
}

func (r *BuildAdjustAccountPayload) SetAccountNumber(accountNumber uint32) {
	r.AccountNumber = accountNumber
}

func (r *BuildAdjustAccountPayload) SetExtraCKB(extraCKB uint64) {
	r.ExtraCKB = extraCKB
}

func (r *BuildAdjustAccountPayload) SetFeeRate(feeRate uint64) {
	r.FeeRate = feeRate
}


type buildAdjustAccountPayloadBuilder struct {
	Item          *req.Item
	From          []*req.Item
	AssetInfo     *common.AssetInfo
	AccountNumber uint32
	ExtraCKB      uint64
	FeeRate       uint64
}

func (builder *buildAdjustAccountPayloadBuilder) AddItem(item *req.Item) {
	builder.Item = item
}
func (builder *buildAdjustAccountPayloadBuilder) AddFrom(items ...*req.Item) {
	builder.From = items
}

func (builder *buildAdjustAccountPayloadBuilder) AddAssetInfo(assetInfo *common.AssetInfo) {
	builder.AssetInfo = assetInfo
}

func (builder *buildAdjustAccountPayloadBuilder) AddAccountNumber(accountNumber uint32) {
	builder.AccountNumber = accountNumber
}

func (builder *buildAdjustAccountPayloadBuilder) AddExtraCKB(extraCKB uint64) {
	builder.ExtraCKB = extraCKB
}

func (builder *buildAdjustAccountPayloadBuilder) AddFeeRate(feeRate uint64) {
	builder.FeeRate = feeRate
}

func (builder *buildAdjustAccountPayloadBuilder) Build() *BuildAdjustAccountPayload {
	return &BuildAdjustAccountPayload{
		Item:          builder.Item,
		From:          builder.From,
		AssetInfo:     builder.AssetInfo,
		AccountNumber: builder.AccountNumber,
		ExtraCKB:      builder.ExtraCKB,
		FeeRate:       builder.FeeRate,
	}
}

func NewBuildAdjustAccountPayloadBuilder() *buildAdjustAccountPayloadBuilder {
	return &buildAdjustAccountPayloadBuilder{
		FeeRate: 1000,
		From:    make([]*req.Item, 0),
	}
}
