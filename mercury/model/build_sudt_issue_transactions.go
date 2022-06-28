package model

import (
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/req"
)

type BuildSudtIssueTransactionPayload struct {
	Owner                  string                 `json:"owner"`
	From                   []*req.Item            `json:"from"`
	To                     []*ToInfo              `json:"to"`
	OutputCapacityProvider OutputCapacityProvider `json:"output_capacity_provider,omitempty"`
	FeeRate                uint64                 `json:"fee_rate,omitempty"`
	Since                  *SinceConfig           `json:"since,omitempty"`
}
