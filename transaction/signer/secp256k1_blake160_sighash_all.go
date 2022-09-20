package signer

import (
	"bytes"
	"errors"
	"github.com/nervosnetwork/ckb-sdk-go/crypto/blake2b"
	"github.com/nervosnetwork/ckb-sdk-go/crypto/secp256k1"
	"github.com/nervosnetwork/ckb-sdk-go/transaction"
	"github.com/nervosnetwork/ckb-sdk-go/types"
)

type Secp256k1Blake160SighashAllSigner struct {
}

func uint32ArrayToIntArray(uint32Array []uint32) []int {
	var intArray []int
	for _, i := range uint32Array {
		intArray = append(intArray, int(i))
	}
	return intArray
}

func (s *Secp256k1Blake160SighashAllSigner) SignTransaction(tx *types.Transaction, group *transaction.ScriptGroup, ctx *transaction.Context) (bool, error) {
	matched, err := IsSingleSigMatched(ctx.Key, group.Script.Args)
	if err != nil {
		return false, err
	}
	if matched {
		i0 := group.InputIndices[0]
		signature, err := transaction.SignTransaction(tx, uint32ArrayToIntArray(group.InputIndices), tx.Witnesses[i0], ctx.Key)
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

func IsSingleSigMatched(key *secp256k1.Secp256k1Key, scriptArgs []byte) (bool, error) {
	if key == nil || scriptArgs == nil {
		return false, errors.New("key or scriptArgs is nil")
	}
	hash := blake2b.Blake160(key.PubKey())
	return bytes.Equal(scriptArgs, hash), nil
}
