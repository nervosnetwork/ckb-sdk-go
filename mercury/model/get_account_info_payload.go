package model

import (
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/common"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/req"
)

type GetAccountInfoPayload struct {
	Item      *req.Item         `json:"item"`
	AssetInfo *common.AssetInfo `json:"asset_info"`
}
