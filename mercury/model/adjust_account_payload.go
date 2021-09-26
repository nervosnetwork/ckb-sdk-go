package model

import (
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/common"
)

type BuildAdjustAccountPayload struct {
	Item          interface{}       `json:"item"`
	From          []interface{}     `json:"from"`
	AssetInfo     *common.AssetInfo `json:"asset_info"`
	AccountNumber uint              `json:"account_number"`
	ExtraCKB      *uint             `json:"extra_ckb"`
	FeeRate       uint              `json:"fee_rate"`
}

func NewBuildAdjustAccountPayload() *BuildAdjustAccountPayload {
	return &BuildAdjustAccountPayload{
		FeeRate: 1000,
		From:    []interface{}{},
	}
}

func (a *BuildAdjustAccountPayload) AddItemToFrom(item interface{}) {
	a.From = append(a.From, item)
}
