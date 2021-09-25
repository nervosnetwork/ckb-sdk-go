package req

import (
	. "github.com/nervosnetwork/ckb-sdk-go/mercury/model/types"
	"math/big"
)

type SmartTransferPayload struct {
	AssetInfo *AssetInfo
	From      []string
	To        []*SmartTo
	Change    string
	FeeRate   uint
}

type SmartTo struct {
	Address string
	Amount  *big.Int
}

type smartTransferPayloadBuilder struct {
	AssetInfo *AssetInfo
	From      []string
	To        []*SmartTo
	Change    string
	FeeRate   uint
}

func (builder *smartTransferPayloadBuilder) AddAssetInfo(info *AssetInfo) {
	builder.AssetInfo = info
}

func (builder *smartTransferPayloadBuilder) AddFrom(from string) {
	builder.From = append(builder.From, from)
}

func (builder *smartTransferPayloadBuilder) AddSmartTo(address string, amount *big.Int) {
	builder.To = append(builder.To, &SmartTo{address, amount})
}

func (builder *smartTransferPayloadBuilder) AddChange(change string) {
	builder.Change = change
}

func (builder *smartTransferPayloadBuilder) AddFeeRate(feeRate uint) {
	builder.FeeRate = feeRate
}

func (builder *smartTransferPayloadBuilder) Build() *SmartTransferPayload {
	return &SmartTransferPayload{
		builder.AssetInfo,
		builder.From,
		builder.To,
		builder.Change,
		builder.FeeRate,
	}
}

func NewSmartTransferPayloadBuilder() *smartTransferPayloadBuilder {
	// default fee rate
	return &smartTransferPayloadBuilder{
		From:    make([]string, 0),
		To:      make([]*SmartTo, 0),
		FeeRate: 1000,
	}
}
