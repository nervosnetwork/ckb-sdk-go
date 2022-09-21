package handler

import (
	"github.com/nervosnetwork/ckb-sdk-go/collector"
	"github.com/nervosnetwork/ckb-sdk-go/script"
	"github.com/nervosnetwork/ckb-sdk-go/script/signer"
	"github.com/nervosnetwork/ckb-sdk-go/types"
	"reflect"
)

type Secp256k1Blake160SighashAllScriptHandler struct {
	CellDep  *types.CellDep
	CodeHash types.Hash
}

func NewSecp256k1Blake160SighashAllScriptHandler(network types.Network) *Secp256k1Blake160SighashAllScriptHandler {
	var txHash types.Hash
	if network == types.NetworkMain {
		txHash = types.HexToHash("0x71a7ba8fc96349fea0ed3a5c47992e3b4084b031a42264a018e0072e8172e46c")
	} else if network == types.NetworkTest {
		txHash = types.HexToHash("0xf8de3bb47d055cdf460d93a2a6e1b05f7432f9777c8c474abf4eec1d4aee5d37")
	} else {
		return nil
	}

	return &Secp256k1Blake160SighashAllScriptHandler{
		CellDep: &types.CellDep{
			OutPoint: &types.OutPoint{
				TxHash: txHash,
				Index:  0,
			},
			DepType: types.DepTypeDepGroup,
		},
		CodeHash: script.GetCodeHash(network, script.SystemScriptSecp256k1Blake160SighashAll),
	}
}

func (r *Secp256k1Blake160SighashAllScriptHandler) isMatched(script *types.Script) bool {
	if script == nil {
		return false
	}
	return reflect.DeepEqual(script.CodeHash, r.CodeHash)
}

func (r *Secp256k1Blake160SighashAllScriptHandler) BuildTransaction(builder collector.TransactionBuilder, group *signer.ScriptGroup, context interface{}) (bool, error) {
	if group == nil || !r.isMatched(group.Script) {
		return false, nil
	}
	index := group.InputIndices[0]
	lock := [65]byte{}
	if err := builder.SetWitness(uint(index), types.WitnessTypeLock, lock[:]); err != nil {
		return false, err
	}
	builder.AddCellDep(r.CellDep)
	return true, nil
}

type Secp256k1Blake160MultisigAllScriptHandler struct {
	cellDep *types.CellDep
	network types.Network
}

func NewSecp256k1Blake160MultisigAllScriptHandler(network types.Network) *Secp256k1Blake160MultisigAllScriptHandler {
	var txHash types.Hash
	if network == types.NetworkMain {
		txHash = types.HexToHash("0x71a7ba8fc96349fea0ed3a5c47992e3b4084b031a42264a018e0072e8172e46c")
	} else if network == types.NetworkTest {
		txHash = types.HexToHash("0xf8de3bb47d055cdf460d93a2a6e1b05f7432f9777c8c474abf4eec1d4aee5d37")
	} else {
		return nil
	}

	return &Secp256k1Blake160MultisigAllScriptHandler{
		cellDep: &types.CellDep{
			OutPoint: &types.OutPoint{
				TxHash: txHash,
				Index:  1,
			},
			DepType: types.DepTypeDepGroup,
		},
		network: network,
	}
}

func (r *Secp256k1Blake160MultisigAllScriptHandler) isMatched(s *types.Script) bool {
	if s == nil {
		return false
	}
	codeHash := script.GetCodeHash(r.network, script.SystemScriptSecp256k1Blake160MultisigAll)
	return reflect.DeepEqual(s.CodeHash, codeHash)
}

func (r *Secp256k1Blake160MultisigAllScriptHandler) BuildTransaction(builder collector.TransactionBuilder, group *signer.ScriptGroup, context interface{}) (bool, error) {
	if group == nil || !r.isMatched(group.Script) {
		return false, nil
	}
	var lock []byte
	switch context.(type) {
	case script.MultisigConfig, *script.MultisigConfig:
		var (
			config *script.MultisigConfig
			ok     bool
		)
		if config, ok = context.(*script.MultisigConfig); !ok {
			v, _ := context.(script.MultisigConfig)
			config = &v
		}
		lock = config.WitnessPlaceholderInLock()
	default:
		return false, nil
	}
	index := group.InputIndices[0]
	if err := builder.SetWitness(uint(index), types.WitnessTypeLock, lock[:]); err != nil {
		return false, err
	}
	builder.AddCellDep(r.cellDep)
	return true, nil
}
