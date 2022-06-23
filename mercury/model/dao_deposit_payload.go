package model

import (
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/req"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/source"
)

type DaoDepositPayload struct {
	From    *req.Item `json:"from"`
	To      string    `json:"to,omitempty"`
	Amount  uint64    `json:"amount"`
	FeeRate uint64    `json:"fee_rate,omitempty"`
}

type daoDepositPayloadBuilder struct {
	From    *req.Item
	To      string
	Amount  uint64
	FeeRate uint64
}

// TODO: fix
func (builder *daoDepositPayloadBuilder) AddFrom(source source.Source, items ...interface{}) {
	//builder.From = &From{
	//	Items:  items,
	//	Source: source,
	//}
}

func (builder *daoDepositPayloadBuilder) AddTo(to string) {
	builder.To = to
}

func (builder *daoDepositPayloadBuilder) AddAmount(amount uint64) {
	builder.Amount = amount
}

func (builder *daoDepositPayloadBuilder) AddFeeRate(feeRate uint64) {
	builder.FeeRate = feeRate
}

func (builder *daoDepositPayloadBuilder) Build() *DaoDepositPayload {
	return &DaoDepositPayload{
		From:    builder.From,
		To:      builder.To,
		Amount:  builder.Amount,
		FeeRate: builder.FeeRate,
	}
}

func NewDaoDepositPayloadBuilder() *daoDepositPayloadBuilder {
	// default fee rate
	return &daoDepositPayloadBuilder{
		FeeRate: 1000,
	}
}
