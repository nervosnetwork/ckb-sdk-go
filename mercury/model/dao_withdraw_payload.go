package model

import "github.com/nervosnetwork/ckb-sdk-go/mercury/model/req"

type DaoWithdrawPayload struct {
	From    []*req.Item `json:"from"`
	FeeRate uint64      `json:"fee_rate,omitempty"`
}
