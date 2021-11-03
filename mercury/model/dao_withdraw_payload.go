package model

type DaoWithdrawPayload struct {
	From    interface{} `json:"from"`
	PayFee  string      `json:"pay_fee"`
	FeeRate uint64      `json:"fee_rate"`
}

type daoWithdrawPayloadBuilder struct {
	From    interface{}
	PayFee  string
	FeeRate uint64
}

func (builder *daoWithdrawPayloadBuilder) AddItem(item interface{}) {
	builder.From = item
}

func (builder *daoWithdrawPayloadBuilder) AddPayFee(address string) {
	builder.PayFee = address
}

func (builder *daoWithdrawPayloadBuilder) AddFeeRate(address uint64) {
	builder.FeeRate = address
}

func (builder *daoWithdrawPayloadBuilder) Build() *DaoWithdrawPayload {
	return &DaoWithdrawPayload{
		From:    builder.From,
		PayFee:  builder.PayFee,
		FeeRate: builder.FeeRate,
	}
}

func NewDaoWithdrawPayloadBuilder() *daoWithdrawPayloadBuilder {
	// default fee rate
	return &daoWithdrawPayloadBuilder{
		FeeRate: 1000,
	}
}
