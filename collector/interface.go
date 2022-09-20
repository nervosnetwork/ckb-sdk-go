package collector

import (
	"github.com/nervosnetwork/ckb-sdk-go/transaction"
	"github.com/nervosnetwork/ckb-sdk-go/types"
)

type TransactionBuilder interface {
	SetVersion(version uint32)
	AddHeaderDep(headerDep types.Hash) int
	AddCellDep(cellDep *types.CellDep) int
	AddInput(input *types.CellInput) int
	SetSince(index uint, since uint64) error
	AddOutput(output *types.CellOutput, data []byte) int
	SetOutputData(index uint, data []byte) error
	SetWitness(index uint, witnessType types.WitnessType, data []byte) error
	AddScriptGroup(group *transaction.ScriptGroup) int
	Build(contexts ...interface{}) (*transaction.TransactionWithScriptGroups, error)
}

type ScriptHandler interface {
	BuildTransaction(builder TransactionBuilder, group *transaction.ScriptGroup, context interface{}) (bool, error)
}
