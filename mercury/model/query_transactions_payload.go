package model

import (
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/common"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/req"
)

type QueryTransactionsPayload struct {
	Item          *req.Item           `json:"item"`
	AssetInfos    []*common.AssetInfo `json:"asset_infos"`
	Extra         *ExtraFilterType    `json:"extra,omitempty"`
	BlockRange    *BlockRange         `json:"block_range,omitempty"`
	Pagination    *PaginationRequest  `json:"pagination"`
	StructureType StructureType       `json:"structure_type"`
}

type ExtraFilterType string

const (
	ExtraFilterDao      ExtraFilterType = "Dao"
	ExtraFilterCellBase ExtraFilterType = "CellBase"
)

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
