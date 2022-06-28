package model

import (
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/common"
)

type SimpleTransferPayload struct {
	AssetInfo *common.AssetInfo `json:"asset_info"`
	From      []string          `json:"from"`
	To        []*ToInfo         `json:"to"`
	FeeRate   uint64            `json:"fee_rate,omitempty"`
	Since     *SinceConfig      `json:"since,omitempty"`
}
