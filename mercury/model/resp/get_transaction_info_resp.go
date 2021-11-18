package resp

import (
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/common"
	"github.com/nervosnetwork/ckb-sdk-go/types"
)

type GetTransactionInfoResponse struct {
	Transaction  *TransactionInfo        `json:"transaction"`
	Status       types.TransactionStatus `json:"status"`
	RejectReason uint8                   `json:"reject_reason"`
}

type TransactionInfo struct {
	TxHash    string      `json:"tx_hash"`
	Records   []Record    `json:"records"`
	Fee       int64       `json:"fee"`
	Burn      []*BurnInfo `json:"burn"`
	Timestamp int64       `json:"timestamp"`
}

type BurnInfo struct {
	UdtHash string      `json:"udt_hash"`
	Amount  *model.U128 `json:"amount"`
}

type Record struct {
	Id          string            `json:"id"`
	Ownership   *common.Ownership `json:"ownership"`
	Amount      *model.U128       `json:"amount"`
	Occupied    *model.U128       `json:"occupied"`
	AssetInfo   *common.AssetInfo `json:"asset_info"`
	Status      RecordStatus      `json:"status"`
	Extra       ExtraFilter       `json:"extra"`
	BlockNumber uint64            `json:"block_number"`
	EpochNumber uint64            `json:"epoch_number"`
}

type RecordStatus struct {
	Type  RecordStatusType `json:"type"`
	Value uint64           `json:"value"`
}

type RecordStatusType string

const (
	RecordStatusFixed     RecordStatusType = "Fixed"
	RecordStatusClaimable                  = "Claimable"
)

type ExtraFilter struct {
	Type  common.ExtraFilterType `json:"type"`
	Value *DaoInfo               `json:"value"`
}

type DaoInfo struct {
	DepositBlockNumber  uint64
	WithdrawBlockNumber uint64
	DaoState            DaoState
	Reward              uint64
}

type DaoState = string

const (
	Deposit  DaoState = "Deposit"
	Withdraw DaoState = "Withdraw"
)

type AssetStatus string

const (
	Claimable AssetStatus = "Claimable"
	Fixed     AssetStatus = "Fixed"
)
