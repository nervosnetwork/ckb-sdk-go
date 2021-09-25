package resp

import (
	"github.com/ethereum/go-ethereum/common/hexutil"
	. "github.com/nervosnetwork/ckb-sdk-go/mercury/model/common"
	"github.com/nervosnetwork/ckb-sdk-go/transaction"
	"github.com/nervosnetwork/ckb-sdk-go/types"
)

type TransferCompletionResponse struct {
	TxView    transactionResp   `json:"tx_view"`
	SigsEntry []*SignatureEntry `json:"sigs_entry"`
}

type SignatureEntry struct {
	Type     string `json:"type"`
	Index    int    `json:"index"`
	PubKey   string `json:"pub_key"`
	GroupLen int    `json:"group_len"`
}

type ScriptGroup struct {
	Group       []int
	WitnessArgs *types.WitnessArgs
	PubKey      string
}

func (self *TransferCompletionResponse) GetTransaction() *types.Transaction {
	return toTransaction(self.TxView)
}

func (self *TransferCompletionResponse) GetScriptGroup() []*ScriptGroup {
	groupScripts := make([]*ScriptGroup, len(self.SigsEntry))

	self.TxView.Witnesses = make([]hexutil.Bytes, len(self.TxView.Inputs))
	for i, _ := range self.TxView.Witnesses {
		self.TxView.Witnesses[i] = []byte{}
	}

	for index, entry := range self.SigsEntry {
		group := make([]int, entry.GroupLen)
		groupIndex := 0
		for i := entry.Index; i < entry.Index+entry.GroupLen; i++ {
			group[groupIndex] = i
			groupIndex += 1
		}

		groupScripts[index] = &ScriptGroup{
			Group:       group,
			WitnessArgs: transaction.EmptyWitnessArg,
			PubKey:      entry.PubKey,
		}

		self.TxView.Witnesses[entry.Index] = transaction.EmptyWitnessArgPlaceholder
	}
	return groupScripts
}

func toTransaction(tx transactionResp) *types.Transaction {
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

func toCellDeps(deps []CellDep) []*types.CellDep {
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

func toInputs(inputs []CellInput) []*types.CellInput {
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

func toOutputs(outputs []CellOutput) []*types.CellOutput {
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
