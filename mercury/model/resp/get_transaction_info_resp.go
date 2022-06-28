package resp

import (
	"encoding/json"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/common"
	"github.com/nervosnetwork/ckb-sdk-go/types"
	"github.com/pkg/errors"
	"math/big"
	"reflect"
)

type GetTransactionInfoResponse struct {
	Transaction *TransactionInfo        `json:"transaction,omitempty"`
	Status      types.TransactionStatus `json:"status"`
}

type TransactionInfo struct {
	TxHash    types.Hash  `json:"tx_hash"`
	Records   []*Record   `json:"records"`
	Fee       uint64      `json:"fee"`
	Burn      []*BurnInfo `json:"burn"`
	Timestamp uint64      `json:"timestamp"`
}

type BurnInfo struct {
	UdtHash types.Hash `json:"udt_hash"`
	Amount  *big.Int   `json:"amount"`
}

type Record struct {
	OutPoint    *types.OutPoint   `json:"out_point"`
	Ownership   string            `json:"ownership"`
	IoType      IoType            `json:"io_type"`
	Amount      *big.Int          `json:"amount"`
	Occupied    *big.Int          `json:"occupied"`
	AssetInfo   *common.AssetInfo `json:"asset_info"`
	Extra       *ExtraFilter      `json:"extra,omitempty"`
	BlockNumber uint64            `json:"block_number"`
	EpochNumber uint64            `json:"epoch_number"`
}

type RecordStatus struct {
	Type  RecordStatusType `json:"type"`
	Value uint64           `json:"value"`
}

type RecordStatusType string
type IoType string

const (
	RecordStatusFixed     RecordStatusType = "Fixed"
	RecordStatusClaimable RecordStatusType = "Claimable"
	IoTypeInput           IoType           = "Input"
	IoTypeOutput          IoType           = "Output"
)

type ExtraFilter struct {
	Type  ExtraFilterType `json:"type"`
	Value *DaoInfo        `json:"value,omitempty"`
}

type DaoInfo struct {
	State  DaoState `json:"state"`
	Reward uint64   `json:"reward"`
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
	ExtraFilterCellBase ExtraFilterType = "Cellbase"
	ExtraFilterFreeze   ExtraFilterType = "Frozen"
)
