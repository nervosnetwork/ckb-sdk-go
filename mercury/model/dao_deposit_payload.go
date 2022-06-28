package model

import (
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/req"
)

type DaoDepositPayload struct {
	From    []*req.Item `json:"from"`
	To      string      `json:"to,omitempty"`
	Amount  uint64      `json:"amount"`
	FeeRate uint64      `json:"fee_rate,omitempty"`
}
