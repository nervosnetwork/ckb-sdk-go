package signer

import (
	"errors"
	"github.com/nervosnetwork/ckb-sdk-go/crypto/blake2b"
	"github.com/nervosnetwork/ckb-sdk-go/transaction"
	"github.com/nervosnetwork/ckb-sdk-go/types"
	"github.com/nervosnetwork/ckb-sdk-go/utils"
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
			utils.GetCodeHash(network, types.BuiltinScriptSecp256k1Blake160SighashAll), &Secp256k1Blake160SighashAllSigner{})
		instance.RegisterLockSigner(
			utils.GetCodeHash(network, types.BuiltinScriptSecp256k1Blake160MultisigAll), &Secp256k1Blake160MultisigAllSigner{})
		instance.RegisterLockSigner(
			utils.GetCodeHash(network, types.BuiltinScriptAnyoneCanPay), &AnyCanPaySigner{})
		instance.RegisterLockSigner(
			utils.GetCodeHash(network, types.BuiltinScriptPwLock), &PWLockSigner{})
	}
	return instance
}

func (r *TransactionSigner) RegisterSigner(codeHash types.Hash, scriptType transaction.ScriptType, signer ScriptSigner) {
	hash := hash(codeHash, scriptType)
	r.signers[hash] = signer
}

func (r *TransactionSigner) RegisterTypeSigner(codeHash types.Hash, signer ScriptSigner) {
	r.RegisterSigner(codeHash, transaction.ScriptTypeType, signer)
}

func (r *TransactionSigner) RegisterLockSigner(codeHash types.Hash, signer ScriptSigner) {
	r.RegisterSigner(codeHash, transaction.ScriptTypeLock, signer)
}

func hash(codeHash types.Hash, scriptType transaction.ScriptType) types.Hash {
	data := codeHash.Bytes()
	data = append(data, []byte(scriptType)...)
	return types.BytesToHash(blake2b.Blake256(data))
}

func (r *TransactionSigner) SignTransactionByPrivateKeys(tx *transaction.TransactionWithScriptGroups, privKeys ...string) ([]int, error) {
	ctxs := transaction.NewContexts()
	if err := ctxs.AddByPrivateKeys(privKeys...); err != nil {
		return nil, err
	}
	return r.SignTransaction(tx, ctxs)
}

func (r *TransactionSigner) SignTransaction(tx *transaction.TransactionWithScriptGroups, contexts transaction.Contexts) ([]int, error) {
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

func checkScriptGroup(group *transaction.ScriptGroup) error {
	if group == nil {
		return errors.New("nil ScriptGroup")
	}
	switch group.GroupType {
	case transaction.ScriptTypeType:
		if len(group.OutputIndices)+len(group.InputIndices) < 0 {
			return errors.New("groupType is Type but OutputIndices and InputIndices are empty")
		}
	case transaction.ScriptTypeLock:
		if len(group.InputIndices) == 0 {
			return errors.New("groupType is Lock but InputIndices is empty")
		}
	default:
		return errors.New("unknown group type " + string(group.GroupType))
	}
	return nil
}
