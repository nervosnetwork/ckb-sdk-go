package model

import (
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/common"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/req"
	"math/big"
)

type TransferPayload struct {
	AssetInfo              *common.AssetInfo      `json:"asset_info,omitempty"`
	From                   []*req.Item            `json:"from"`
	To                     []*ToInfo                `json:"to"`
	OutputCapacityProvider OutputCapacityProvider `json:"output_capacity_provider,omitempty"`
	PayFee                 PayFee                 `json:"pay_fee,omitempty"`
	FeeRate                uint64                 `json:"fee_rate,omitempty"`
	Since                  *SinceConfig           `json:"since,omitempty"`
}

type PayFee string
type OutputCapacityProvider string

const (
	PayFeeFrom                 PayFee                 = "From"
	PayFeeTo                   PayFee                 = "To"
	OutputCapacityProviderFrom OutputCapacityProvider = "From"
	OutputCapacityProviderTo   OutputCapacityProvider = "To"
)

type ToInfo struct {
	Address string   `json:"address"`
	Amount  *big.Int `json:"amount"`
}

type SinceConfig struct {
	Flag  SinceFlag `json:"flag"`
	Type  SinceType `json:"type_"`
	Value uint64    `json:"value"`
}

type SinceFlag string

const (
	Relative SinceFlag = "Relative"
	Absolute SinceFlag = "Absolute"
)

type SinceType string

const (
	BlockNumber SinceType = "BlockNumber"
	EpochNumber SinceType = "EpochNumber"
	Timestamp   SinceType = "Timestamp"
)
