package model

import "github.com/nervosnetwork/ckb-sdk-go/mercury/model/req"

type DaoClaimPayload struct {
	From    []*req.Item `json:"from"`
	To      string      `json:"to,omitempty"`
	FeeRate uint64      `json:"fee_rate,omitempty"`
}
