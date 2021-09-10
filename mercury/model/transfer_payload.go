package model

import (
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/action"
	"math/big"
)

type TransferPayload struct {
	UdtHash string          `json:"udt_hash,omitempty"`
	From    interface{}     `json:"from"`
	Items   []*TransferItem `json:"items"`
	Change  string          `json:"change,omitempty"`
	FeeRate uint            `json:"fee_rate"`
}

type FromKeyAddresses struct {
	KeyAddresses *keyAddresses `json:"key_addresses"`
}

type FromNormalAddresses struct {
	NormalAddresses []string `json:"normal_addresses"`
}

type keyAddresses struct {
	KeyAddresses []string `json:"key_addresses"`
	Source       string   `json:"source"`
}

type ToAddress interface {
	IsPayBayFrom() bool
}

type ToKeyAddress struct {
	KeyAddress *keyAddress `json:"key_address"`
}

func (address *ToKeyAddress) IsPayBayFrom() bool {
	if address.KeyAddress.Action == action.Pay_by_from {
		return true
	}

	return false
}

type ToNormalAddress struct {
	NormalAddress string `json:"normal_address"`
}

func (address *ToNormalAddress) IsPayBayFrom() bool {
	return false
}

type keyAddress struct {
	KeyAddress string `json:"key_address"`
	Action     string `json:"action"`
}

type TransferItem struct {
	To     ToAddress `json:"to"`
	Amount *big.Int  `json:"amount"`
}

type transferBuilder struct {
	UdtHash string          `json:"udt_hash,omitempty"`
	From    interface{}     `json:"from"`
	Items   []*TransferItem `json:"items"`
	Change  string          `json:"change,omitempty"`
	FeeRate uint            `json:"fee_rate"`
}

func (builder *transferBuilder) AddUdtHash(udtHash string) {
	builder.UdtHash = udtHash
}

func (builder *transferBuilder) AddFromKeyAddresses(keyAddr []string, source string) {
	builder.From = &FromKeyAddresses{
		KeyAddresses: &keyAddresses{
			KeyAddresses: keyAddr,
			Source:       source,
		},
	}
}

func (builder *transferBuilder) AddFromNormalAddresses(normalAddress []string) {
	builder.From = &FromNormalAddresses{
		NormalAddresses: normalAddress,
	}
}

func (builder *transferBuilder) AddToKeyAddressItem(addr, action string, amount *big.Int) {
	builder.Items = append(builder.Items, &TransferItem{Amount: amount, To: &ToKeyAddress{
		KeyAddress: &keyAddress{
			KeyAddress: addr,
			Action:     action,
		},
	}})
}

func (builder *transferBuilder) AddToNormalAddressItem(addr string, amount *big.Int) {
	builder.Items = append(builder.Items, &TransferItem{Amount: amount, To: &ToNormalAddress{
		NormalAddress: addr,
	}})
}

func (builder *transferBuilder) AddChange(change string) {
	builder.Change = change
}

func (builder *transferBuilder) AddFeeRate(feeRate uint) {
	builder.FeeRate = feeRate
}

func (builder *transferBuilder) Build() *TransferPayload {
	return &TransferPayload{
		builder.UdtHash,
		builder.From,
		builder.Items,
		builder.Change,
		builder.FeeRate,
	}
}

func NewTransferBuilder() *transferBuilder {
	// default fee rate
	return &transferBuilder{
		FeeRate: 1000,
	}
}
