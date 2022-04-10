package model

import (
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/common"
)

type GetAccountInfoPayload struct {
	Item      interface{}       `json:"item"`
	AssetInfo *common.AssetInfo `json:"asset_info"`
}

type GetAccountInfoPayloadBuilder struct {
	Item      interface{}       `json:"item"`
	AssetInfo *common.AssetInfo `json:"asset_info"`
}

func (b *GetAccountInfoPayloadBuilder) SetItem(item interface{}) *GetAccountInfoPayloadBuilder {
	b.Item = item
	return b
}
func (b *GetAccountInfoPayloadBuilder) AddAssetInfo(assetInfo *common.AssetInfo) *GetAccountInfoPayloadBuilder {
	b.AssetInfo = assetInfo
	return b
}

func NewGetAccountInfoPayloadBuilder() *GetAccountInfoPayloadBuilder {
	return &GetAccountInfoPayloadBuilder{
		Item:      nil,
		AssetInfo: nil,
	}
}

func (b GetAccountInfoPayloadBuilder) Build() *GetAccountInfoPayload {
	payload := &GetAccountInfoPayload{
		Item:      b.Item,
		AssetInfo: b.AssetInfo,
	}
	return payload
}
