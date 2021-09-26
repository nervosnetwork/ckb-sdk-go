package model

type WithdrawPayload struct {
	From    interface{} `json:"from"`
	PayFee  string      `json:"pay_fee"`
	FeeRate uint64      `json:"fee_rate"`
}

type withdrawPayloadBuilder struct {
	From    interface{}
	PayFee  string
	FeeRate uint64
}

func (builder *withdrawPayloadBuilder) AddItem(item interface{}) {
	builder.From = item
}

func (builder *withdrawPayloadBuilder) AddPayFee(address string) {
	builder.PayFee = address
}

func (builder *withdrawPayloadBuilder) AddFeeRate(address uint64) {
	builder.FeeRate = address
}

func (builder *withdrawPayloadBuilder) Build() *WithdrawPayload {
	return &WithdrawPayload{
		From:    builder.From,
		PayFee:  builder.PayFee,
		FeeRate: builder.FeeRate,
	}
}

func NewWithdrawPayloadBuilder() *withdrawPayloadBuilder {
	// default fee rate
	return &withdrawPayloadBuilder{
		FeeRate: 1000,
	}
}
