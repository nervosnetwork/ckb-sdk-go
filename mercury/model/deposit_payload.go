package model

import (
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/source"
)

type DepositPayload struct {
	From    *From  `json:"from"`
	To      string `json:"to,omitempty"`
	Amount  uint64 `json:"amount"`
	FeeRate uint64 `json:"fee_rate"`
}

type depositPayloadBuilder struct {
	From    *From
	To      string
	Amount  uint64
	FeeRate uint64
}

func (builder *depositPayloadBuilder) AddFrom(source source.Source, items ...interface{}) {
	builder.From = &From{
		Items:  items,
		Source: source,
	}
}

func (builder *depositPayloadBuilder) AddTo(to string) {
	builder.To = to
}

func (builder *depositPayloadBuilder) AddAmount(amount uint64) {
	builder.Amount = amount
}

func (builder *depositPayloadBuilder) AddFeeRate(feeRate uint64) {
	builder.FeeRate = feeRate
}

func (builder *depositPayloadBuilder) Build() *DepositPayload {
	return &DepositPayload{
		From:    builder.From,
		To:      builder.To,
		Amount:  builder.Amount,
		FeeRate: builder.FeeRate,
	}
}

func NewDepositPayloadBuilder() *depositPayloadBuilder {
	// default fee rate
	return &depositPayloadBuilder{
		FeeRate: 1000,
	}
}
