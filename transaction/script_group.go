package transaction

import (
	"encoding/json"
	"errors"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/nervosnetwork/ckb-sdk-go/types"
	"strings"
)

type TransactionWithScriptGroups struct {
	TxView       *types.Transaction `json:"tx_view"`
	ScriptGroups []*ScriptGroup     `json:"script_groups"`
}

type ScriptGroup struct {
	Script        types.Script `json:"script"`
	GroupType     ScriptType   `json:"group_type"`
	InputIndices  []uint32     `json:"input_indices"`
	OutputIndices []uint32     `json:"output_indices"`
}

type ScriptType string

const (
	ScriptTypeLock ScriptType = "lock"
	ScriptTypeType ScriptType = "type"
)

func (r *ScriptType) UnmarshalJSON(input []byte) error {
	var jsonObj string
	if err := json.Unmarshal(input, &jsonObj); err != nil {
		return err
	}
	switch strings.ToLower(jsonObj) {
	case "":
		*r = ""
	case strings.ToLower(string(ScriptTypeLock)):
		*r = ScriptTypeLock
	case strings.ToLower(string(ScriptTypeType)):
		*r = ScriptTypeType
	default:
		return errors.New("can't unmarshal json from unknown script type " + string(input))
	}
	return nil
}

func (r *ScriptGroup) UnmarshalJSON(input []byte) error {
	var jsonObj struct {
		Script        types.Script   `json:"script"`
		GroupType     ScriptType     `json:"group_type"`
		InputIndices  []hexutil.Uint `json:"input_indices"`
		OutputIndices []hexutil.Uint `json:"output_indices"`
	}
	if err := json.Unmarshal(input, &jsonObj); err != nil {
		return err
	}
	toUint32Array := func(a []hexutil.Uint) []uint32 {
		result := make([]uint32, len(a))
		for i, data := range a {
			result[i] = uint32(data)
		}
		return result
	}
	*r = ScriptGroup{
		Script:        jsonObj.Script,
		GroupType:     jsonObj.GroupType,
		InputIndices:  toUint32Array(jsonObj.InputIndices),
		OutputIndices: toUint32Array(jsonObj.OutputIndices),
	}
	return nil
}
