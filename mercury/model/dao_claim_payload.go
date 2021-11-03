package model

type DaoClaimPayload struct {
	From    interface{} `json:"from"`
	To      string      `json:"to,omitempty"`
	FeeRate uint64      `json:"fee_rate"`
}

type daoClaimPayloadBuilder struct {
	From    interface{}
	To      string
	FeeRate uint64
}

func (builder *daoClaimPayloadBuilder) AddItem(item interface{}) {
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
		From:    builder.From,
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
