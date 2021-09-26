package model

import (
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/common"
)

type GetBalancePayload struct {
	AssetInfos     []*common.AssetInfo `json:"asset_infos"`
	TipBlockNumber uint64              `json:"tip_block_number,omitempty"`
	Item           interface{}         `json:"item"`
}

type getBalancePayloadBuilder struct {
	item           interface{}
	assetInfos     []*common.AssetInfo
	TipBlockNumber uint64
}

func (builder *getBalancePayloadBuilder) AddAssetInfo(info *common.AssetInfo) {
	builder.assetInfos = append(builder.assetInfos, info)
}

func (builder *getBalancePayloadBuilder) AddItem(item interface{}) {
	builder.item = item
}

func (builder *getBalancePayloadBuilder) AddTipBlockNumber(tipBlockNumber uint64) {
	builder.TipBlockNumber = tipBlockNumber
}

func (builder *getBalancePayloadBuilder) Build() *GetBalancePayload {

	payload := &GetBalancePayload{
		Item:           builder.item,
		TipBlockNumber: builder.TipBlockNumber,
		AssetInfos:     builder.assetInfos,
	}

	return payload
}

func NewGetBalancePayloadBuilder() *getBalancePayloadBuilder {
	return &getBalancePayloadBuilder{
		assetInfos: make([]*common.AssetInfo, 0),
	}
}
