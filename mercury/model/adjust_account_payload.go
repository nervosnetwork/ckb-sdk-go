package model

import (
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/common"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/req"
)

type BuildAdjustAccountPayload struct {
	Item          *req.Item         `json:"item"`
	From          []*req.Item       `json:"from"`
	AssetInfo     *common.AssetInfo `json:"asset_info"`
	AccountNumber uint32            `json:"account_number,omitempty"`
	ExtraCKB      uint64            `json:"extra_ckb,omitempty"`
	FeeRate       uint64            `json:"fee_rate,omitempty"`
}
