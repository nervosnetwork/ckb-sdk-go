package collector_test

import (
	"reflect"

	"github.com/nervosnetwork/ckb-sdk-go/v2/collector"
	"github.com/nervosnetwork/ckb-sdk-go/v2/collector/builder"
	"github.com/nervosnetwork/ckb-sdk-go/v2/transaction"
	"github.com/nervosnetwork/ckb-sdk-go/v2/types"
)

// SimpleLockScriptHandler is an example script handler to add specified cell dep and prefill the witness.
type SimpleLockScriptHandler struct {
	CellDep            *types.CellDep
	WitnessPlaceholder []byte
	CodeHash           types.Hash
}

func (r *SimpleLockScriptHandler) isMatched(script *types.Script) bool {
	if script == nil {
		return false
	}
	return reflect.DeepEqual(script.CodeHash, r.CodeHash)
}

func (r *SimpleLockScriptHandler) BuildTransaction(builder collector.TransactionBuilder, group *transaction.ScriptGroup, context interface{}) (bool, error) {
	// Only run on matched groups
	if group == nil || !r.isMatched(group.Script) {
		return false, nil
	}
	index := group.InputIndices[0]
	// set the witness placeholder
	if err := builder.SetWitness(uint(index), types.WitnessTypeLock, r.WitnessPlaceholder); err != nil {
		return false, err
	}
	// CkbTransactionBuilder.AddCellDep will remove duplications automatically.
	builder.AddCellDep(r.CellDep)
	return true, nil
}

func ExampleScriptHandler() {
	txHash := "0x1234"
	typeScriptHash := "0xabcd"

	s := builder.SimpleTransactionBuilder{}
	s.Register(&SimpleLockScriptHandler{
		CellDep: &types.CellDep{
			OutPoint: &types.OutPoint{
				TxHash: types.HexToHash(txHash),
				Index:  0,
			},
			DepType: types.DepTypeCode,
		},
		CodeHash:           types.HexToHash(typeScriptHash),
		WitnessPlaceholder: make([]byte, 8),
	})
}
