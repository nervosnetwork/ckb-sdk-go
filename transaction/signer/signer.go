package signer

import (
	"errors"
	"github.com/nervosnetwork/ckb-sdk-go/crypto/blake2b"
	"github.com/nervosnetwork/ckb-sdk-go/transaction"
	"github.com/nervosnetwork/ckb-sdk-go/types"
)

type ScriptSigner interface {
	SignTransaction(transaction *types.Transaction, group *transaction.ScriptGroup, ctx *transaction.Context) (bool, error)
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
			types.GetCodeHash(types.BuiltinScriptSecp256k1Blake160SighashAll, network), &Secp256k1Blake160SighashAllSigner{})
		instance.RegisterLockSigner(
			types.GetCodeHash(types.BuiltinScriptSecp256k1Blake160MultisigAll, network), &Secp256k1Blake160MultisigAllSigner{})
		instance.RegisterLockSigner(
			types.GetCodeHash(types.BuiltinScriptAnyoneCanPay, network), &AnyCanPaySigner{})
		instance.RegisterLockSigner(
			types.GetCodeHash(types.BuiltinScriptPWLock, network), &PWLockSigner{})
	}
	return instance
}

func (r *TransactionSigner) RegisterSigner(codeHash types.Hash, scriptType transaction.ScriptType, signer ScriptSigner) error {
	hash, err := hash(codeHash, scriptType)
	if err != nil {
		return err
	}
	r.signers[hash] = signer
	return nil
}

func (r *TransactionSigner) RegisterTypeSigner(codeHash types.Hash, signer ScriptSigner) error {
	return r.RegisterSigner(codeHash, transaction.ScriptTypeType, signer)
}

func (r *TransactionSigner) RegisterLockSigner(codeHash types.Hash, signer ScriptSigner) error {
	return r.RegisterSigner(codeHash, transaction.ScriptTypeLock, signer)
}

func hash(codeHash types.Hash, scriptType transaction.ScriptType) (types.Hash, error) {
	data := codeHash.Bytes()
	data = append(data, []byte(scriptType)...)
	hash, err := blake2b.Blake256(data)
	if err != nil {
		return types.Hash{}, err
	}
	return types.BytesToHash(hash), nil
}

func (r *TransactionSigner) signTransaction(transaction *transaction.TransactionWithScriptGroups, contexts transaction.Contexts) ([]int, error) {
	signedIndex := make([]int, 0)
	for i, group := range transaction.ScriptGroups {
		if err := checkScriptGroup(group); err != nil {
			return signedIndex, err
		}
		key, err := hash(group.Script.CodeHash, group.GroupType)
		if err != nil {
			return signedIndex, err
		}
		signer := r.signers[key]
		if signer != nil {
			for _, ctx := range contexts {
				var signed bool
				if signed, err = signer.SignTransaction(transaction.TxView, group, ctx); err != nil {
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

func checkScriptGroup(group *transaction.ScriptGroup) error {
	if group == nil {
		return errors.New("nil ScriptGroup")
	}
	switch group.GroupType {
	case transaction.ScriptTypeType:
		if len(group.OutputIndices) == 0 {
			return errors.New("groupType is Lock but OutputIndices is empty")
		}
	case transaction.ScriptTypeLock:
		if len(group.InputIndices) == 0 {
			return errors.New("groupType is Type but InputIndices is empty")
		}
	default:
		return errors.New("unknown group type " + string(group.GroupType))
	}
	return nil
}
