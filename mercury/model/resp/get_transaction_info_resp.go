package resp

import (
	"encoding/json"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/common"
	"github.com/nervosnetwork/ckb-sdk-go/types"
	"github.com/pkg/errors"
	"reflect"
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
	Type  ExtraFilterType `json:"type"`
	Value *DaoInfo        `json:"value"`
}

type DaoInfo struct {
	DepositBlockNumber  uint64   `json:"deposit_block_number,omitempty"`
	WithdrawBlockNumber uint64   `json:"withdraw_block_number,omitempty"`
	DaoState            DaoState `json:"state"`
	Reward              uint64   `json:"reward"`
}

type DaoState struct {
	Type  DaoStateType `json:"type"`
	Value []uint64     `json:"value"`
}

type DaoStateType = string

const (
	DaoStateDeposit  DaoStateType = "Deposit"
	DaoStateWithdraw DaoStateType = "Withdraw"
)

func (e *DaoState) UnmarshalJSON(bytes []byte) error {
	var data map[string]interface{}

	if err := json.Unmarshal(bytes, &data); err != nil {
		return err
	}

	e.Type = data["type"].(string)

	var v = data["value"]
	switch reflect.ValueOf(v).Kind() {
	case reflect.Float64:
		e.Value = make([]uint64, 1)
		e.Value[0] = uint64(v.(float64))
	case reflect.Slice:
		vv := v.([]interface{})
		e.Value = make([]uint64, len(vv))
		for i := range vv {
			e.Value[i] = uint64(vv[i].(float64))
		}
	default:
		return errors.New("invalid type while unmarshal DaoState")
	}

	return nil
}

type ExtraFilterType string

const (
	ExtraFilterDao      ExtraFilterType = "Dao"
	ExtraFilterCellBase                 = "CellBase"
	ExtraFilterFreeze                   = "Freeze"
)

type AssetStatus string

const (
	Claimable AssetStatus = "Claimable"
	Fixed     AssetStatus = "Fixed"
)
