package model

import (
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/common"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/req"
)

type GetBalancePayload struct {
	AssetInfos     []*common.AssetInfo `json:"asset_infos"`
	TipBlockNumber uint64              `json:"tip_block_number,omitempty"`
	Item           *req.Item           `json:"item"`
}
