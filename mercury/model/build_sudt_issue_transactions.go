package model

import (
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/mode"
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

type buildSudtIssueTransactionPayloadBuilder struct {
	Owner                  string
	From                   []*req.Item
	To                     []*ToInfo
	OutputCapacityProvider OutputCapacityProvider
	FeeRate                uint64
	Since                  *SinceConfig
}

func (builder *buildSudtIssueTransactionPayloadBuilder) AddOwner(owner string) {
	builder.Owner = owner
}

// TODO fix
func (builder *buildSudtIssueTransactionPayloadBuilder) AddTo(mode mode.Mode, toInfos ...*ToInfo) {
	//builder.To = &To{
	//	ToInfos: toInfos,
	//	Mode:    mode,
	//}
}

func (builder *buildSudtIssueTransactionPayloadBuilder) AddFeeRate(feeRate uint64) {
	builder.FeeRate = feeRate
}
func (builder *buildSudtIssueTransactionPayloadBuilder) AddSince(since *SinceConfig) {
	builder.Since = since
}
func (builder *buildSudtIssueTransactionPayloadBuilder) Build() *BuildSudtIssueTransactionPayload {
	return &BuildSudtIssueTransactionPayload{
		Owner:                  builder.Owner,
		From:                   builder.From,
		To:                     builder.To,
		OutputCapacityProvider: builder.OutputCapacityProvider,
		FeeRate:                builder.FeeRate,
		Since:                  builder.Since,
	}
}

func NewBuildSudtIssueTransactionPayloadBuilder() *buildSudtIssueTransactionPayloadBuilder {
	return &buildSudtIssueTransactionPayloadBuilder{
		FeeRate: 1000,
	}
}
