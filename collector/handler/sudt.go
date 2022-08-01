package handler

import (
	"github.com/nervosnetwork/ckb-sdk-go/collector"
	"github.com/nervosnetwork/ckb-sdk-go/transaction"
	"github.com/nervosnetwork/ckb-sdk-go/types"
	"reflect"
)

type SudtScriptHandler struct {
	cellDep *types.CellDep
	network types.Network
}

func NewSudtScriptHandler(network types.Network) *SudtScriptHandler {
	var txHash types.Hash
	if network == types.NetworkMain {
		txHash = types.HexToHash("0xc7813f6a415144643970c2e88e0bb6ca6a8edc5dd7c1022746f628284a9936d5")
	} else if network == types.NetworkTest {
		txHash = types.HexToHash("0xe12877ebd2c3c364dc46c5c992bcfaf4fee33fa13eebdf82c591fc9825aab769")
	} else {
		return nil
	}
	return &SudtScriptHandler{
		cellDep: &types.CellDep{
			OutPoint: &types.OutPoint{
				TxHash: txHash,
				Index:  0,
			},
			DepType: types.DepTypeCode,
		},
		network: network,
	}
}

func (r *SudtScriptHandler) isMatched(script *types.Script) bool {
	if script == nil {
		return false
	}
	codeHash := types.GetCodeHash(types.BuiltinScriptSudt, r.network)
	return reflect.DeepEqual(script.CodeHash, codeHash)
}

func (r *SudtScriptHandler) BuildTransaction(builder collector.TransactionBuilder, group *transaction.ScriptGroup, context interface{}) (bool, error) {
	if group == nil || !r.isMatched(group.Script) {
		return false, nil
	}
	builder.AddCellDep(r.cellDep)
	return true, nil
}
