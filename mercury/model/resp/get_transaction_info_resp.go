package resp

import (
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/common"
	"github.com/nervosnetwork/ckb-sdk-go/types"
)

type TransactionInfoWithStatusResponse struct {
	Transaction     *TransactionInfoResponse
	Status          types.TransactionStatus
	BlockHash       string
	BlockNumber     uint64
	ConfirmedNumber uint64
}

type TransactionInfoResponse struct {
	TxHash     string
	Operations []*RecordResponse
}

type RecordResponse struct {
	Id          uint
	Address     string
	Amount      string
	AssetInfo   *common.AssetInfo
	Status      AssetStatus
	BlockNumber uint
}

type AssetStatus string

const (
	CLAIMABLE AssetStatus = "claimable"
	FIXED     AssetStatus = "fixed"
)
