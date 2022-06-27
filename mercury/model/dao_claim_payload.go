package model

import "github.com/nervosnetwork/ckb-sdk-go/mercury/model/req"

type DaoClaimPayload struct {
	From    []*req.Item `json:"from"`
	To      string      `json:"to,omitempty"`
	FeeRate uint64      `json:"fee_rate,omitempty"`
}

type daoClaimPayloadBuilder struct {
	From    *req.Item
	To      string
	FeeRate uint64
}

func (builder *daoClaimPayloadBuilder) AddItem(item *req.Item) {
	builder.From = item
}

func (builder *daoClaimPayloadBuilder) AddTo(to string) {
	builder.To = to
}

func (builder *daoClaimPayloadBuilder) AddFeeRate(feeRate uint64) {
	builder.FeeRate = feeRate
}

func (builder *daoClaimPayloadBuilder) Build() *DaoClaimPayload {
	return &DaoClaimPayload{
		From:    nil,
		To:      builder.To,
		FeeRate: builder.FeeRate,
	}
}

func NewDaoClaimPayloadBuilder() *daoClaimPayloadBuilder {
	// default fee rate
	return &daoClaimPayloadBuilder{
		FeeRate: 1000,
	}
}
