package resp

import (
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/common"
	"github.com/nervosnetwork/ckb-sdk-go/types"
)

type TransferCompletionResponse struct {
	TxView           *transactionResp   `json:"tx_view"`
	SignatureActions []*SignatureAction `json:"signature_actions"`
}

type ScriptGroup struct {
	Action          *SignatureAction
	Transaction     *transactionResp
	OriginalWitness string
}

func (self *TransferCompletionResponse) GetTransaction() *types.Transaction {
	return toTransaction(self.TxView)
}

func (self *TransferCompletionResponse) GetScriptGroup() []*ScriptGroup {
	scriptGroups := make([]*ScriptGroup, len(self.SignatureActions))
	for i, v := range self.SignatureActions {
		scriptGroups[i] = NewScriptGroup(v, self.TxView)
	}
	return scriptGroups
}

func toTransaction(tx *transactionResp) *types.Transaction {
	return &types.Transaction{
		Version:     uint(tx.Version),
		Hash:        tx.Hash,
		CellDeps:    toCellDeps(tx.CellDeps),
		HeaderDeps:  tx.HeaderDeps,
		Inputs:      toInputs(tx.Inputs),
		Outputs:     toOutputs(tx.Outputs),
		OutputsData: toBytesArray(tx.OutputsData),
		Witnesses:   toBytesArray(tx.Witnesses),
	}
}

func toCellDeps(deps []common.CellDep) []*types.CellDep {
	result := make([]*types.CellDep, len(deps))
	for i := 0; i < len(deps); i++ {
		dep := deps[i]
		result[i] = &types.CellDep{
			OutPoint: &types.OutPoint{
				TxHash: dep.OutPoint.TxHash,
				Index:  uint(dep.OutPoint.Index),
			},
			DepType: dep.DepType,
		}
	}
	return result
}

func toInputs(inputs []common.CellInput) []*types.CellInput {
	result := make([]*types.CellInput, len(inputs))
	for i := 0; i < len(inputs); i++ {
		input := inputs[i]
		result[i] = &types.CellInput{
			Since: uint64(input.Since),
			PreviousOutput: &types.OutPoint{
				TxHash: input.PreviousOutput.TxHash,
				Index:  uint(input.PreviousOutput.Index),
			},
		}
	}
	return result
}

func toOutputs(outputs []common.CellOutput) []*types.CellOutput {
	result := make([]*types.CellOutput, len(outputs))
	for i := 0; i < len(outputs); i++ {
		output := outputs[i]
		result[i] = &types.CellOutput{
			Capacity: uint64(output.Capacity),
			Lock: &types.Script{
				CodeHash: output.Lock.CodeHash,
				HashType: output.Lock.HashType,
				Args:     output.Lock.Args,
			},
		}
		if output.Type != nil {
			result[i].Type = &types.Script{
				CodeHash: output.Type.CodeHash,
				HashType: output.Type.HashType,
				Args:     output.Type.Args,
			}
		}
	}
	return result
}

func toBytesArray(bytes []hexutil.Bytes) [][]byte {
	result := make([][]byte, len(bytes))
	for i, data := range bytes {
		result[i] = data
	}
	return result
}

func NewScriptGroup(action *SignatureAction, transaction *transactionResp) *ScriptGroup {
	return &ScriptGroup{
		Action:          action,
		Transaction:     transaction,
		OriginalWitness: transaction.Witnesses[action.SignatureLocation.Index].String(),
	}
}

func (g *ScriptGroup) GetOffSet() int {
	return g.Action.SignatureLocation.Offset
}

func (g *ScriptGroup) GetWitness() string {
	return g.OriginalWitness
}

func (g *ScriptGroup) GetWitnessIndex() int {
	return g.Action.SignatureLocation.Index
}

func (g *ScriptGroup) GetAddress() string {
	return g.Action.SignatureInfo.Address
}

func (g *ScriptGroup) GetGroupWitnesses() []string {
	var groupWitnesses []string
	groupWitnesses = append(groupWitnesses, g.OriginalWitness)
	for _, v := range g.Action.OtherIndexesInGroup {
		groupWitnesses = append(groupWitnesses, g.Transaction.Witnesses[v].String())
	}
	return groupWitnesses
}
