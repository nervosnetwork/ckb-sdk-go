package model

import (
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/common"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/req"
)

type QueryTransactionsPayload struct {
	Item          *req.Item               `json:"item"`
	AssetInfos    []*common.AssetInfo     `json:"asset_infos"`
	Extra         *common.ExtraFilterType `json:"extra,omitempty"`
	BlockRange    *BlockRange             `json:"block_range,omitempty"`
	Pagination    *PaginationRequest      `json:"pagination"`
	StructureType StructureType           `json:"structure_type"`
}

func (v *QueryTransactionsPayload) AddAssetInfo(assetInfo *common.AssetInfo) {
	v.AssetInfos = append(v.AssetInfos, assetInfo)
}

type BlockRange struct {
	From uint64 `json:"from"`
	To   uint64 `json:"to"`
}

type PaginationRequest struct {
	Cursor      uint64 `json:"cursor,omitempty"`
	Order       Order  `json:"order"`
	Limit       uint64 `json:"limit,omitempty"`
	ReturnCount bool   `json:"return_count"`
}

type Order string

const (
	ASC  Order = "Asc"
	DESC Order = "Desc"
)

type QueryTransactionsPayloadBuilder struct {
	Item          *req.Item
	AssetInfos    []*common.AssetInfo
	Extra         *common.ExtraFilterType
	BlockRange    *BlockRange
	Pagination    *PaginationRequest
	StructureType StructureType
}

func (b *QueryTransactionsPayloadBuilder) SetItem(item *req.Item) *QueryTransactionsPayloadBuilder {
	b.Item = item
	return b
}
func (b *QueryTransactionsPayloadBuilder) AddAssetInfo(assetInfo *common.AssetInfo) *QueryTransactionsPayloadBuilder {
	b.AssetInfos = append(b.AssetInfos, assetInfo)
	return b
}
func (b *QueryTransactionsPayloadBuilder) SetExtra(extra *common.ExtraFilterType) *QueryTransactionsPayloadBuilder {
	b.Extra = extra
	return b
}

func (b *QueryTransactionsPayloadBuilder) AddBlockRange(blockRangerange *BlockRange) *QueryTransactionsPayloadBuilder {
	b.BlockRange = blockRangerange
	return b
}

func (b *QueryTransactionsPayloadBuilder) SetCursor(cursor uint64) *QueryTransactionsPayloadBuilder {
	b.Pagination.Cursor = cursor
	return b
}

func (b *QueryTransactionsPayloadBuilder) SetOrder(order Order) *QueryTransactionsPayloadBuilder {
	b.Pagination.Order = order
	return b
}

func (b *QueryTransactionsPayloadBuilder) SetLimit(limit uint64) *QueryTransactionsPayloadBuilder {
	b.Pagination.Limit = limit
	return b
}

func NewQueryTransactionsPayloadBuilder() *QueryTransactionsPayloadBuilder {
	return &QueryTransactionsPayloadBuilder{
		Item:       nil,
		AssetInfos: []*common.AssetInfo{},
		Extra:      nil,
		BlockRange: nil,
		Pagination: &PaginationRequest{
			Order:       DESC,
			Limit:       50,
			ReturnCount: false,
		},
	}
}

func (b QueryTransactionsPayloadBuilder) Build() *QueryTransactionsPayload {
	payload := &QueryTransactionsPayload{
		Item:          b.Item,
		AssetInfos:    b.AssetInfos,
		Extra:         b.Extra,
		BlockRange:    b.BlockRange,
		Pagination:    b.Pagination,
		StructureType: b.StructureType,
	}
	if payload.Pagination.Cursor == 0 && payload.Pagination.Order == DESC {
		payload.Pagination.Cursor = 0x7ffffffffffffffe
	}
	return payload
}
