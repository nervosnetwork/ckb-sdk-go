package signer

import (
	"github.com/nervosnetwork/ckb-sdk-go/crypto/secp256k1"
	"github.com/nervosnetwork/ckb-sdk-go/types"
)

type AnyCanPaySigner struct {
}

func (s *AnyCanPaySigner) SignTransaction(tx *types.Transaction, group *ScriptGroup, ctx *Context) (bool, error) {
	matched, err := IsAnyCanPayMatched(ctx.Key, group.Script.Args)
	if err != nil {
		return false, err
	}
	if matched {
		i0 := group.InputIndices[0]
		signature, err := SignTransaction(tx, uint32ArrayToIntArray(group.InputIndices), tx.Witnesses[i0], ctx.Key)
		if err != nil {
			return false, err
		}
		witnessArgs, err := types.DeserializeWitnessArgs(tx.Witnesses[i0])
		if err != nil {
			return false, err
		}
		witnessArgs.Lock = signature
		tx.Witnesses[i0] = witnessArgs.Serialize()
		return true, nil
	} else {
		return false, nil
	}
}

func IsAnyCanPayMatched(key *secp256k1.Secp256k1Key, scriptArgs []byte) (bool, error) {
	return IsSingleSigMatched(key, scriptArgs[:20])
}
