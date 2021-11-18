package model

import (
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/common"
)

type QueryTransactionsPayload struct {
	Item          interface{}             `json:"item"`
	AssetInfos    []*common.AssetInfo     `json:"asset_infos"`
	Extra         *common.ExtraFilterType `json:"extra"`
	BlockRange    *BlockRange             `json:"block_range"`
	Pagination    PaginationRequest       `json:"pagination"`
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
	Cursor      []int  `json:"cursor"`
	Order       Order  `json:"order"`
	Limit       uint64 `json:"limit"`
	Skip        uint64 `json:"skip"`
	ReturnCount bool   `json:"return_count"`
}

type Order string

const (
	ASC  Order = "asc"
	DESC Order = "desc"
)

type QueryTransactionsPayloadBuilder struct {
	Item          interface{}
	AssetInfos    []*common.AssetInfo
	Extra         *common.ExtraFilterType
	BlockRange    *BlockRange
	Pagination    PaginationRequest
	StructureType StructureType
}

func (b *QueryTransactionsPayloadBuilder) SetItem(item interface{}) *QueryTransactionsPayloadBuilder {
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

func (b *QueryTransactionsPayloadBuilder) SetCursor(cursor []int) *QueryTransactionsPayloadBuilder {
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

func (b *QueryTransactionsPayloadBuilder) SetPageNumber(skip uint64) *QueryTransactionsPayloadBuilder {
	b.Pagination.Skip = skip
	return b
}

func NewQueryTransactionsPayloadBuilder() *QueryTransactionsPayloadBuilder {
	return &QueryTransactionsPayloadBuilder{
		Item:       nil,
		AssetInfos: []*common.AssetInfo{},
		Extra:      nil,
		BlockRange: nil,
		Pagination: PaginationRequest{
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
	if b.Pagination.Skip != 0 {
		x, _ := math.SafeMul(payload.Pagination.Limit, payload.Pagination.Skip)
		y, _ := math.SafeSub(x, payload.Pagination.Limit)
		payload.Pagination.Skip = y
	}
	if payload.Pagination.Cursor == nil && payload.Pagination.Order == DESC {
		payload.Pagination.Cursor = []int{127, 255, 255, 255, 255, 255, 255, 254}
	}
	return payload
}
