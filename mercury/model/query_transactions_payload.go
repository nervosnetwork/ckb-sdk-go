package model

import (
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/common"
)

type QueryTransactionsPayload struct {
	Item          interface{}         `json:"item"`
	AssetInfos    []*common.AssetInfo `json:"asset_infos"`
	Extra         *common.ExtraType   `json:"extra"`
	BlockRange    *BlockRange         `json:"block_range"`
	Pagination    PaginationRequest   `json:"pagination"`
	StructureType StructureType       `json:"structure_type"`
}

func (v *QueryTransactionsPayload) AddAssetInfo(assetInfo *common.AssetInfo) {
	v.AssetInfos = append(v.AssetInfos, assetInfo)
}

type BlockRange struct {
	From *uint64 `json:"from"`
	To   *uint64 `json:"to"`
}

type PaginationRequest struct {
	Cursor      []int   `json:"cursor"`
	Order       Order   `json:"order"`
	Limit       *uint64 `json:"limit"`
	Skip        *uint64 `json:"skip"`
	ReturnCount bool    `json:"return_count"`
}

type Order string

const (
	ASC  Order = "asc"
	DESC Order = "desc"
)
