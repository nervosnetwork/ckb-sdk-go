package model

import "github.com/nervosnetwork/ckb-sdk-go/mercury/model/mode"

type BuildSudtIssueTransactionPayload struct {
	Owner   string       `json:"owner"`
	To      *To          `json:"to"`
	PayFee  interface{}  `json:"pay_fee,omitempty"`
	Change  string       `json:"change,omitempty"`
	FeeRate uint64       `json:"fee_rate,omitempty"`
	Since   *SinceConfig `json:"since,omitempty"`
}

type buildSudtIssueTransactionPayloadBuilder struct {
	Owner   string
	To      *To
	PayFee  interface{}
	Change  string
	FeeRate uint64
	Since   *SinceConfig
}

func (builder *buildSudtIssueTransactionPayloadBuilder) AddOwner(owner string) {
	builder.Owner = owner
}
func (builder *buildSudtIssueTransactionPayloadBuilder) AddTo(mode mode.Mode, toInfos ...*ToInfo) {
	builder.To = &To{
		ToInfos: toInfos,
		Mode:    mode,
	}
}
func (builder *buildSudtIssueTransactionPayloadBuilder) AddPayFee(payFee interface{}) {
	builder.PayFee = payFee
}
func (builder *buildSudtIssueTransactionPayloadBuilder) AddChange(Change string) {
	builder.Change = Change
}
func (builder *buildSudtIssueTransactionPayloadBuilder) AddFeeRate(feeRate uint64) {
	builder.FeeRate = feeRate
}
func (builder *buildSudtIssueTransactionPayloadBuilder) AddSince(since *SinceConfig) {
	builder.Since = since
}
func (builder *buildSudtIssueTransactionPayloadBuilder) Build() *BuildSudtIssueTransactionPayload {
	return &BuildSudtIssueTransactionPayload{
		Owner:   builder.Owner,
		To:      builder.To,
		PayFee:  builder.PayFee,
		Change:  builder.Change,
		FeeRate: builder.FeeRate,
		Since:   builder.Since,
	}
}

func NewBuildSudtIssueTransactionPayloadBuilder() *buildSudtIssueTransactionPayloadBuilder {
	return &buildSudtIssueTransactionPayloadBuilder{
		FeeRate: 1000,
	}
}
