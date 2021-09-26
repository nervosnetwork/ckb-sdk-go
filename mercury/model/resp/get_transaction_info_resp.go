package resp

import (
	"encoding/json"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/common"
	"github.com/nervosnetwork/ckb-sdk-go/types"
	"strings"
)

type GetTransactionInfoResponse struct {
	Transaction  *TransactionInfo        `json:"transaction"`
	Status       types.TransactionStatus `json:"status"`
	RejectReason uint8                   `json:"reject_reason"`
}

type TransactionInfo struct {
	TxHash  string      `json:"tx_hash"`
	Records []Record    `json:"records"`
	Fee     int64       `json:"fee"`
	Burn    []*BurnInfo `json:"burn"`
}

type BurnInfo struct {
	UdtHash string      `json:"udt_hash"`
	Amount  *model.U128 `json:"amount"`
}

type Record struct {
	Id                string                    `json:"id"`
	AddressOrLockHash *common.AddressOrLockHash `json:"address_or_lock_hash"`
	Amount            *model.U128               `json:"amount"`
	Occupied          *model.U128               `json:"occupied"`
	AssetInfo         *common.AssetInfo         `json:"asset_info"`
	Status            RecordStatus              `json:"status"`
	Extra             ExtraFilter               `json:"extra"`
	BlockNumber       uint64                    `json:"block_number"`
	EpochNumber       []byte                    `json:"epoch_number"`
}

type RecordStatus struct {
	Status      AssetStatus
	BlockNumber uint64
}

func (r *RecordStatus) UnmarshalJSON(bytes []byte) error {
	recordData := make(map[string]interface{})
	json.Unmarshal(bytes, &recordData)

	if _, ok := recordData["Claimable"]; ok {
		blockNumber := recordData["Claimable"].(float64)
		r.BlockNumber = uint64(blockNumber)
		r.Status = Claimable
	} else {
		blockNumber := recordData["Fixed"].(float64)
		r.BlockNumber = uint64(blockNumber)
		r.Status = Fixed

	}

	return nil
}

type ExtraFilter struct {
	DaoInfo   *DaoInfo
	ExtraType ExtraType
}

func (e *ExtraFilter) UnmarshalJSON(bytes []byte) error {
	if strings.Contains(string(bytes), "null") {
		return nil
	}

	if strings.Contains(string(bytes), "CellBase") {
		e.ExtraType = CellBase
	} else {
		ExtraFilterData := make(map[string]interface{})
		json.Unmarshal(bytes, &ExtraFilterData)

		DaoData := ExtraFilterData["Dao"].(map[string]interface{})
		stateData := DaoData["state"].(map[string]interface{})
		var depositBlockNumber uint64
		var withdrawBlockNumber uint64
		var state DaoState

		if _, ok := stateData["Deposit"]; ok {
			depositNumber := stateData["Deposit"].(float64)
			depositBlockNumber = uint64(depositNumber)
			state = Deposit

		} else {
			withdraw := stateData["Withdraw"].([]interface{})
			depositNumber := withdraw[0].(float64)
			withdrawNumber := withdraw[1].(float64)
			depositBlockNumber = uint64(depositNumber)
			withdrawBlockNumber = uint64(withdrawNumber)
			state = Withdraw
		}

		reward := DaoData["reward"].(float64)

		e.DaoInfo = &DaoInfo{
			DepositBlockNumber:  depositBlockNumber,
			WithdrawBlockNumber: withdrawBlockNumber,
			DaoState:            state,
			Reward:              uint64(reward),
		}

		e.ExtraType = Dao

	}

	return nil
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

type ExtraType string

const (
	Dao      ExtraType = "Dao"
	CellBase ExtraType = "CellBase"
)

type AssetStatus string

const (
	Claimable AssetStatus = "Claimable"
	Fixed     AssetStatus = "Fixed"
)
