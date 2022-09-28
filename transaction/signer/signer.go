package signer

import (
	"fmt"
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

var testInstance = NewTransactionSigner()
var mainInstance = NewTransactionSigner()

func GetTransactionSignerInstance(network types.Network) *TransactionSigner {
	if network == types.NetworkTest {
		return testInstance
	} else if network == types.NetworkMain {
		return mainInstance
	} else {
		return nil
	}
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

func (r *TransactionSigner) SignTransactionByPrivateKeys(tx *transaction.TransactionWithScriptGroups, privKeys ...string) ([]int, error) {
	var ctxs []*transaction.Context
	for _, key := range privKeys {
		ctx, err := transaction.NewContext(key)
		if err != nil {
			return nil, err
		}
		ctxs = append(ctxs, ctx)
	}
	return r.SignTransaction(tx, ctxs...)
}

func (r *TransactionSigner) SignTransaction(tx *transaction.TransactionWithScriptGroups, contexts ...*transaction.Context) ([]int, error) {
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
