package model

import (
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/common"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/req"
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
	ExtraFilterDao             ExtraFilterType        = "Dao"
	ExtraFilterCellBase        ExtraFilterType        = "CellBase"
	ASC                        Order                  = "Asc"
	DESC                       Order                  = "Desc"
	PayFeeFrom                 PayFee                 = "From"
	PayFeeTo                   PayFee                 = "To"
	OutputCapacityProviderFrom OutputCapacityProvider = "From"
	OutputCapacityProviderTo   OutputCapacityProvider = "To"
	Native                     StructureType          = "Native"
	DoubleEntry                StructureType          = "DoubleEntry"

	BlockNumber SinceType = "BlockNumber"
	EpochNumber SinceType = "EpochNumber"
	Timestamp   SinceType = "Timestamp"
	Relative    SinceFlag = "Relative"
	Absolute    SinceFlag = "Absolute"
)

type GetAccountInfoPayload struct {
	Item      *req.Item         `json:"item"`
	AssetInfo *common.AssetInfo `json:"asset_info"`
}

type GetBalancePayload struct {
	AssetInfos     []*common.AssetInfo `json:"asset_infos"`
	TipBlockNumber uint64              `json:"tip_block_number,omitempty"`
	Item           *req.Item           `json:"item"`
}

type GetBlockInfoPayload struct {
	BlockNumber uint64     `json:"block_number,omitempty"`
	BlockHash   types.Hash `json:"block_hash,omitempty"`
}

type BuildAdjustAccountPayload struct {
	Item          *req.Item         `json:"item"`
	From          []*req.Item       `json:"from"`
	AssetInfo     *common.AssetInfo `json:"asset_info"`
	AccountNumber uint32            `json:"account_number,omitempty"`
	ExtraCKB      uint64            `json:"extra_ckb,omitempty"`
	FeeRate       uint64            `json:"fee_rate,omitempty"`
}

type BuildSudtIssueTransactionPayload struct {
	Owner                  string                 `json:"owner"`
	From                   []*req.Item            `json:"from"`
	To                     []*ToInfo              `json:"to"`
	OutputCapacityProvider OutputCapacityProvider `json:"output_capacity_provider,omitempty"`
	FeeRate                uint64                 `json:"fee_rate,omitempty"`
	Since                  *SinceConfig           `json:"since,omitempty"`
}

type SimpleTransferPayload struct {
	AssetInfo *common.AssetInfo `json:"asset_info"`
	From      []string          `json:"from"`
	To        []*ToInfo         `json:"to"`
	FeeRate   uint64            `json:"fee_rate,omitempty"`
	Since     *SinceConfig      `json:"since,omitempty"`
}

type TransferPayload struct {
	AssetInfo              *common.AssetInfo      `json:"asset_info,omitempty"`
	From                   []*req.Item            `json:"from"`
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
	From    []*req.Item `json:"from"`
	To      string      `json:"to,omitempty"`
	Amount  uint64      `json:"amount"`
	FeeRate uint64      `json:"fee_rate,omitempty"`
}

type DaoWithdrawPayload struct {
	From    []*req.Item `json:"from"`
	FeeRate uint64      `json:"fee_rate,omitempty"`
}

type DaoClaimPayload struct {
	From    []*req.Item `json:"from"`
	To      string      `json:"to,omitempty"`
	FeeRate uint64      `json:"fee_rate,omitempty"`
}

type GetSpentTransactionPayload struct {
	OutPoint      types.OutPoint `json:"outpoint"`
	StructureType StructureType  `json:"structure_type"`
}

type QueryTransactionsPayload struct {
	Item          *req.Item           `json:"item"`
	AssetInfos    []*common.AssetInfo `json:"asset_infos"`
	Extra         *ExtraFilterType    `json:"extra,omitempty"`
	BlockRange    *BlockRange         `json:"block_range,omitempty"`
	Pagination    *PaginationRequest  `json:"pagination"`
	StructureType StructureType       `json:"structure_type"`
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
