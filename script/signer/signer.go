package signer

import (
	"fmt"
	"github.com/nervosnetwork/ckb-sdk-go/crypto/blake2b"
	"github.com/nervosnetwork/ckb-sdk-go/script"
	"github.com/nervosnetwork/ckb-sdk-go/types"
)

type ScriptSigner interface {
	SignTransaction(transaction *types.Transaction, group *ScriptGroup, ctx *Context) (bool, error)
}

type TransactionSigner struct {
	signers map[types.Hash]ScriptSigner
}

func NewTransactionSigner() *TransactionSigner {
	return &TransactionSigner{signers: make(map[types.Hash]ScriptSigner)}
}

var testInstance *TransactionSigner
var mainInstance *TransactionSigner

func GetTransactionSignerInstance(network types.Network) *TransactionSigner {
	var instance *TransactionSigner
	var isInitialized = true
	if network == types.NetworkTest {
		if testInstance == nil {
			testInstance = NewTransactionSigner()
			isInitialized = false
		}
		instance = testInstance
	} else if network == types.NetworkMain {
		if mainInstance == nil {
			mainInstance = NewTransactionSigner()
			isInitialized = false
		}
		instance = mainInstance
	} else {
		return nil
	}

	if !isInitialized {
		instance.RegisterLockSigner(
			script.GetCodeHash(network, script.SystemScriptSecp256k1Blake160SighashAll), &Secp256k1Blake160SighashAllSigner{})
		instance.RegisterLockSigner(
			script.GetCodeHash(network, script.SystemScriptSecp256k1Blake160MultisigAll), &Secp256k1Blake160MultisigAllSigner{})
		instance.RegisterLockSigner(
			script.GetCodeHash(network, script.SystemScriptAnyoneCanPay), &AnyCanPaySigner{})
		instance.RegisterLockSigner(
			script.GetCodeHash(network, script.SystemScriptPwLock), &PWLockSigner{})
	}
	return instance
}

func (r *TransactionSigner) RegisterSigner(codeHash types.Hash, scriptType types.ScriptType, signer ScriptSigner) {
	hash := hash(codeHash, scriptType)
	r.signers[hash] = signer
}

func (r *TransactionSigner) RegisterTypeSigner(codeHash types.Hash, signer ScriptSigner) {
	r.RegisterSigner(codeHash, types.ScriptTypeType, signer)
}

func (r *TransactionSigner) RegisterLockSigner(codeHash types.Hash, signer ScriptSigner) {
	r.RegisterSigner(codeHash, types.ScriptTypeLock, signer)
}

func hash(codeHash types.Hash, scriptType types.ScriptType) types.Hash {
	data := codeHash.Bytes()
	data = append(data, []byte(scriptType)...)
	return types.BytesToHash(blake2b.Blake256(data))
}

func (r *TransactionSigner) SignTransactionByPrivateKeys(tx *TransactionWithScriptGroups, privKeys ...string) ([]int, error) {
	ctxs := NewContexts()
	if err := ctxs.AddByPrivateKeys(privKeys...); err != nil {
		return nil, err
	}
	return r.SignTransaction(tx, ctxs)
}

func (r *TransactionSigner) SignTransaction(tx *TransactionWithScriptGroups, contexts Contexts) ([]int, error) {
	var err error
	signedIndex := make([]int, 0)
	for i, group := range tx.ScriptGroups {
		if err := checkScriptGroup(group); err != nil {
			return signedIndex, err
		}
		key := hash(group.Script.CodeHash, group.GroupType)
		signer := r.signers[key]
		if signer != nil {
			for _, ctx := range contexts {
				var signed bool
				if signed, err = signer.SignTransaction(tx.TxView, group, ctx); err != nil {
					return signedIndex, err
				}
				if signed {
					signedIndex = append(signedIndex, i)
					break
				}
			}
		}
	}
	return signedIndex, nil
}

func checkScriptGroup(group *ScriptGroup) error {
	if group == nil {
		return fmt.Errorf("nil ScriptGroup")
	}
	switch group.GroupType {
	case types.ScriptTypeType:
		if len(group.OutputIndices)+len(group.InputIndices) < 0 {
			return fmt.Errorf("groupType is Type but OutputIndices and InputIndices are empty")
		}
	case types.ScriptTypeLock:
		if len(group.InputIndices) == 0 {
			return fmt.Errorf("groupType is Lock but InputIndices is empty")
		}
	default:
		return fmt.Errorf("unknown group type %s", group.GroupType)
	}
	return nil
}
