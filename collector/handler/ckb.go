package handler

import (
	"github.com/nervosnetwork/ckb-sdk-go/v2/collector"
	"github.com/nervosnetwork/ckb-sdk-go/v2/systemscript"
	"github.com/nervosnetwork/ckb-sdk-go/v2/transaction"
	"github.com/nervosnetwork/ckb-sdk-go/v2/types"
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
	} else if network == types.NetworkPreview {
		txHash = types.HexToHash("0x0fab65924f2784f17ad7f86d6aef4b04ca1ca237102a68961594acebc5c77816")
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
		CodeHash: systemscript.GetCodeHash(network, systemscript.Secp256k1Blake160SighashAll),
	}
}

func (r *Secp256k1Blake160SighashAllScriptHandler) isMatched(script *types.Script) bool {
	if script == nil {
		return false
	}
	return reflect.DeepEqual(script.CodeHash, r.CodeHash)
}

func (r *Secp256k1Blake160SighashAllScriptHandler) BuildTransaction(builder collector.TransactionBuilder, group *transaction.ScriptGroup, context interface{}) (bool, error) {
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
	multisigVersion systemscript.MultisigVersion
	cellDep         *types.CellDep
	network         types.Network
}

func NewSecp256k1Blake160MultisigAllScriptHandler(network types.Network, multisigVersion systemscript.MultisigVersion) *Secp256k1Blake160MultisigAllScriptHandler {
	var txHash types.Hash
	var index uint32
	if multisigVersion == systemscript.MultisigLegacy {
		if network == types.NetworkMain {
			txHash = types.HexToHash("0x71a7ba8fc96349fea0ed3a5c47992e3b4084b031a42264a018e0072e8172e46c")
		} else if network == types.NetworkTest {
			txHash = types.HexToHash("0xf8de3bb47d055cdf460d93a2a6e1b05f7432f9777c8c474abf4eec1d4aee5d37")
			index = 1
		} else if network == types.NetworkPreview {
			txHash = types.HexToHash("0x0fab65924f2784f17ad7f86d6aef4b04ca1ca237102a68961594acebc5c77816")
			index = 1
		} else {
			return nil
		}
	} else if multisigVersion == systemscript.MultisigV2 {
		if network == types.NetworkMain {
			txHash = types.HexToHash("0x6888aa39ab30c570c2c30d9d5684d3769bf77265a7973211a3c087fe8efbf738")
			index = 0
		} else if network == types.NetworkTest {
			txHash = types.HexToHash("0x2eefdeb21f3a3edf697c28a52601b4419806ed60bb427420455cc29a090b26d5")
			index = 0
		} else {
			return nil
		}
	} else {
		return nil
	}

	return &Secp256k1Blake160MultisigAllScriptHandler{
		multisigVersion: multisigVersion,
		cellDep: &types.CellDep{
			OutPoint: &types.OutPoint{
				TxHash: txHash,
				Index:  index,
			},
			DepType: types.DepTypeDepGroup,
		},
		network: network,
	}
}

func (r *Secp256k1Blake160MultisigAllScriptHandler) isMatched(script *types.Script) bool {
	if script == nil {
		return false
	}

	var codeHash types.Hash
	var scriptHashType types.ScriptHashType
	if r.multisigVersion == systemscript.MultisigLegacy {
		codeHash = systemscript.GetCodeHash(r.network, systemscript.Secp256k1Blake160MultisigAllLegacy)
		scriptHashType = types.HashTypeType
	} else if r.multisigVersion == systemscript.MultisigV2 {
		codeHash = systemscript.GetCodeHash(r.network, systemscript.Secp256k1Blake160MultisigAllV2)
		scriptHashType = types.HashTypeData1
	} else {
		return false
	}
	return reflect.DeepEqual(script.CodeHash, codeHash) && reflect.DeepEqual(script.HashType, scriptHashType)
}

func (r *Secp256k1Blake160MultisigAllScriptHandler) BuildTransaction(builder collector.TransactionBuilder, group *transaction.ScriptGroup, context interface{}) (bool, error) {
	if group == nil || !r.isMatched(group.Script) {
		return false, nil
	}
	var lock []byte
	switch context.(type) {
	case systemscript.MultisigConfig, *systemscript.MultisigConfig:
		var (
			config *systemscript.MultisigConfig
			ok     bool
		)
		if config, ok = context.(*systemscript.MultisigConfig); !ok {
			v, _ := context.(systemscript.MultisigConfig)
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
