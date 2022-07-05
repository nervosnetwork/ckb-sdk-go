package transaction

import (
	"errors"
	"github.com/nervosnetwork/ckb-sdk-go/crypto/blake2b"
	"github.com/nervosnetwork/ckb-sdk-go/types"
)

type ScriptSigner interface {
	SignTransaction(transaction *types.Transaction, group *ScriptGroup, ctx *Context) (bool, error)
}

type TransactionSigner struct {
	signers map[types.Hash]ScriptSigner
}

func (r *TransactionSigner) RegisterSigner(codeHash types.Hash, scriptType ScriptType, signer ScriptSigner) error {
	hash, err := hash(codeHash, scriptType)
	if err != nil {
		return err
	}
	r.signers[hash] = signer
	return nil
}

func (r *TransactionSigner) RegisterTypeSigner(codeHash types.Hash, signer ScriptSigner) error {
	return r.RegisterSigner(codeHash, ScriptTypeType, signer)
}

func (r *TransactionSigner) RegisterLockSigner(codeHash types.Hash, signer ScriptSigner) error {
	return r.RegisterSigner(codeHash, ScriptTypeLock, signer)
}

func hash(codeHash types.Hash, scriptType ScriptType) (types.Hash, error) {
	data := codeHash.Bytes()
	data = append(data, []byte(scriptType)...)
	hash, err := blake2b.Blake256(data)
	if err != nil {
		return types.Hash{}, err
	}
	return types.BytesToHash(hash), nil
}

func (r *TransactionSigner) signTransaction(transaction *TransactionWithScriptGroups, contexts Contexts) ([]int, error) {
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

func checkScriptGroup(group *ScriptGroup) error {
	if group == nil {
		return errors.New("nil ScriptGroup")
	}
	switch group.GroupType {
	case ScriptTypeType:
		if len(group.OutputIndices) == 0 {
			return errors.New("groupType is Lock but OutputIndices is empty")
		}
	case ScriptTypeLock:
		if len(group.InputIndices) == 0 {
			return errors.New("groupType is Type but InputIndices is empty")
		}
	default:
		return errors.New("unknown group type " + string(group.GroupType))
	}
	return nil
}
