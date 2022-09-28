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
	toUint32Array := func(in []hexutil.Uint) []uint32 {
		out := make([]uint32, len(in))
		for i, data := range in {
			out[i] = uint32(data)
		}
		return out
	}
	*r = ScriptGroup{
		Script:        jsonObj.Script,
		GroupType:     jsonObj.GroupType,
		InputIndices:  toUint32Array(jsonObj.InputIndices),
		OutputIndices: toUint32Array(jsonObj.OutputIndices),
	}
	return nil
}
