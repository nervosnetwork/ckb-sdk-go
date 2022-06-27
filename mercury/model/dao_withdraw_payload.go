package model

import "github.com/nervosnetwork/ckb-sdk-go/mercury/model/req"

type DaoWithdrawPayload struct {
	From    []*req.Item `json:"from"`
	FeeRate uint64      `json:"fee_rate,omitempty"`
}

type daoWithdrawPayloadBuilder struct {
	From    *req.Item
	FeeRate uint64
}

func (builder *daoWithdrawPayloadBuilder) AddItem(item *req.Item) {
	builder.From = item
}

// TODO: remove
func (builder *daoWithdrawPayloadBuilder) AddPayFee(address string) {
	//builder.PayFee = address
}

func (builder *daoWithdrawPayloadBuilder) AddFeeRate(address uint64) {
	builder.FeeRate = address
}

func (builder *daoWithdrawPayloadBuilder) Build() *DaoWithdrawPayload {
	return &DaoWithdrawPayload{
		From:    nil,
		FeeRate: builder.FeeRate,
	}
}

func NewDaoWithdrawPayloadBuilder() *daoWithdrawPayloadBuilder {
	// default fee rate
	return &daoWithdrawPayloadBuilder{
		FeeRate: 1000,
	}
}
