package model

import (
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/common"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/req"
)

type GetAccountInfoPayload struct {
	Item      *req.Item         `json:"item"`
	AssetInfo *common.AssetInfo `json:"asset_info"`
}

type GetAccountInfoPayloadBuilder struct {
	Item      *req.Item         `json:"item"`
	AssetInfo *common.AssetInfo `json:"asset_info"`
}

func (b *GetAccountInfoPayloadBuilder) SetItem(item *req.Item) *GetAccountInfoPayloadBuilder {
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
