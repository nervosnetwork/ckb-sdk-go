package model

import (
	"github.com/nervosnetwork/ckb-sdk-go/types"
	"math/big"
)

type ExtraFilterType string
type Order string
type PayFee string
type OutputCapacityProvider string
type StructureType string
type SinceType string
type SinceFlag string

const (
	ExtraFilterTypeDao      ExtraFilterType = "Dao"
	ExtraFilterTypeCellbase ExtraFilterType = "Cellbase"
	ExtraFilterTypeFrozen   ExtraFilterType = "Frozen"

	OrderAsc                   Order                  = "Asc"
	OrderDesc                  Order                  = "Desc"
	PayFeeFrom                 PayFee                 = "From"
	PayFeeTo                   PayFee                 = "To"
	OutputCapacityProviderFrom OutputCapacityProvider = "From"
	OutputCapacityProviderTo   OutputCapacityProvider = "To"

	StructureTypeNative      StructureType = "Native"
	StructureTypeDoubleEntry StructureType = "DoubleEntry"

	SinceTypeBlockNumber SinceType = "BlockNumber"
	SinceTypeEpochNumber SinceType = "EpochNumber"
	SinceTypeTimestamp   SinceType = "Timestamp"
	SinceFlagRelative    SinceFlag = "Relative"
	SinceFlagAbsolute    SinceFlag = "Absolute"
)

type GetAccountInfoPayload struct {
	Item      *Item      `json:"item"`
	AssetInfo *AssetInfo `json:"asset_info"`
}

type GetBalancePayload struct {
	AssetInfos     []*AssetInfo     `json:"asset_infos"`
	TipBlockNumber uint64           `json:"tip_block_number,omitempty"`
	Item           *Item            `json:"item"`
	Extra          *ExtraFilterType `json:"extra,omitempty"`
}

type GetBlockInfoPayload struct {
	BlockNumber uint64     `json:"block_number,omitempty"`
	BlockHash   types.Hash `json:"block_hash,omitempty"`
}

type BuildAdjustAccountPayload struct {
	Item          *Item      `json:"item"`
	From          []*Item    `json:"from"`
	AssetInfo     *AssetInfo `json:"asset_info"`
	AccountNumber uint32     `json:"account_number,omitempty"`
	ExtraCKB      uint64     `json:"extra_ckb,omitempty"`
	FeeRate       uint64     `json:"fee_rate,omitempty"`
}

type BuildSudtIssueTransactionPayload struct {
	Owner                  string                 `json:"owner"`
	From                   []*Item                `json:"from"`
	To                     []*ToInfo              `json:"to"`
	OutputCapacityProvider OutputCapacityProvider `json:"output_capacity_provider,omitempty"`
	FeeRate                uint64                 `json:"fee_rate,omitempty"`
	Since                  *SinceConfig           `json:"since,omitempty"`
}

type SimpleTransferPayload struct {
	AssetInfo *AssetInfo   `json:"asset_info"`
	From      []string     `json:"from"`
	To        []*ToInfo    `json:"to"`
	FeeRate   uint64       `json:"fee_rate,omitempty"`
	Since     *SinceConfig `json:"since,omitempty"`
}

type TransferPayload struct {
	AssetInfo              *AssetInfo             `json:"asset_info,omitempty"`
	From                   []*Item                `json:"from"`
	To                     []*ToInfo              `json:"to"`
	OutputCapacityProvider OutputCapacityProvider `json:"output_capacity_provider,omitempty"`
	PayFee                 PayFee                 `json:"pay_fee,omitempty"`
	FeeRate                uint64                 `json:"fee_rate,omitempty"`
	Since                  *SinceConfig           `json:"since,omitempty"`
}

type ToInfo struct {
	Address string   `json:"address"`
	Amount  *big.Int `json:"amount"`
}

type SinceConfig struct {
	Flag  SinceFlag `json:"flag"`
	Type  SinceType `json:"type_"`
	Value uint64    `json:"value"`
}

type DaoDepositPayload struct {
	From    []*Item `json:"from"`
	To      string  `json:"to,omitempty"`
	Amount  uint64  `json:"amount"`
	FeeRate uint64  `json:"fee_rate,omitempty"`
}

type DaoWithdrawPayload struct {
	From    []*Item `json:"from"`
	FeeRate uint64  `json:"fee_rate,omitempty"`
}

type DaoClaimPayload struct {
	From    []*Item `json:"from"`
	To      string  `json:"to,omitempty"`
	FeeRate uint64  `json:"fee_rate,omitempty"`
}

type GetSpentTransactionPayload struct {
	OutPoint      types.OutPoint `json:"outpoint"`
	StructureType StructureType  `json:"structure_type"`
}

type QueryTransactionsPayload struct {
	Item          *Item              `json:"item"`
	AssetInfos    []*AssetInfo       `json:"asset_infos"`
	Extra         *ExtraFilterType   `json:"extra,omitempty"`
	BlockRange    *BlockRange        `json:"block_range,omitempty"`
	Pagination    *PaginationRequest `json:"pagination"`
	StructureType StructureType      `json:"structure_type"`
}

type BlockRange struct {
	From uint64 `json:"from"`
	To   uint64 `json:"to"`
}

type PaginationRequest struct {
	Cursor      uint64 `json:"cursor,omitempty"`
	Order       Order  `json:"order"`
	Limit       uint64 `json:"limit,omitempty"`
	ReturnCount bool   `json:"return_count"`
}
