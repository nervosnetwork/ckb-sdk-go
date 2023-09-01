package collector

import (
	"github.com/nervosnetwork/ckb-sdk-go/v2/transaction"
	"github.com/nervosnetwork/ckb-sdk-go/v2/types"
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

// The interface ScriptHandler is for scripts to register their building logic.
//
// The function BuildTransaction is the callback called by [TransactionBuilder]
// for each script group and each context passed in
// TransactionBuilder.Build. The context provides extra data for the script.
//
// Be calfully on when to run the logic for the script. TransactionBuilder will
// not check whether the script group matches the script.
//
// The callback often does two things:
//   - Fill witness placeholder to make fee calculation correct.
//   - Add cell deps for the script.
//
// Returns bool indicating whether the transaction has been modified.
type ScriptHandler interface {
	BuildTransaction(builder TransactionBuilder, group *transaction.ScriptGroup, context interface{}) (bool, error)
}
