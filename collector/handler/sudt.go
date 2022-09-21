package handler

import (
	"github.com/nervosnetwork/ckb-sdk-go/collector"
	"github.com/nervosnetwork/ckb-sdk-go/transaction/signer"
	"github.com/nervosnetwork/ckb-sdk-go/types"
	"github.com/nervosnetwork/ckb-sdk-go/utils"
	"reflect"
)

type SudtScriptHandler struct {
	CellDep  *types.CellDep
	CodeHash types.Hash
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
		CellDep: &types.CellDep{
			OutPoint: &types.OutPoint{
				TxHash: txHash,
				Index:  0,
			},
			DepType: types.DepTypeCode,
		},
		CodeHash: utils.GetCodeHash(network, types.BuiltinScriptSudt),
	}
}

func (r *SudtScriptHandler) isMatched(script *types.Script) bool {
	if script == nil {
		return false
	}
	return reflect.DeepEqual(script.CodeHash, r.CodeHash)
}

func (r *SudtScriptHandler) BuildTransaction(builder collector.TransactionBuilder, group *signer.ScriptGroup, context interface{}) (bool, error) {
	if group == nil || !r.isMatched(group.Script) {
		return false, nil
	}
	builder.AddCellDep(r.CellDep)
	return true, nil
}
