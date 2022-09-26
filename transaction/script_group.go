package transaction

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/nervosnetwork/ckb-sdk-go/types"
)

type TransactionWithScriptGroups struct {
	TxView       *types.Transaction `json:"tx_view"`
	ScriptGroups []*ScriptGroup     `json:"script_groups"`
}

type ScriptGroup struct {
	Script        *types.Script    `json:"script"`
	GroupType     types.ScriptType `json:"group_type"`
	InputIndices  []uint32         `json:"input_indices"`
	OutputIndices []uint32         `json:"output_indices"`
}

func (r *ScriptGroup) UnmarshalJSON(input []byte) error {
	var jsonObj struct {
		Script        *types.Script    `json:"script"`
		GroupType     types.ScriptType `json:"group_type"`
		InputIndices  []hexutil.Uint   `json:"input_indices"`
		OutputIndices []hexutil.Uint   `json:"output_indices"`
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
