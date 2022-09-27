package signer

import (
	"bytes"
	"errors"
	"github.com/nervosnetwork/ckb-sdk-go/v2/crypto/blake2b"
	"github.com/nervosnetwork/ckb-sdk-go/v2/crypto/secp256k1"
	"github.com/nervosnetwork/ckb-sdk-go/v2/transaction"
	"github.com/nervosnetwork/ckb-sdk-go/v2/types"
)

type Secp256k1Blake160SighashAllSigner struct {
}

func (s *Secp256k1Blake160SighashAllSigner) SignTransaction(transaction *types.Transaction, group *transaction.ScriptGroup, ctx *transaction.Context) (bool, error) {
	key := ctx.Key
	matched, err := IsSingleSigMatched(key, group.Script.Args)
	if err != nil {
		return false, err
	}
	if matched {
		return SingleSignTransaction(transaction, group, key)
	} else {
		return false, nil
	}
}

func SingleSignTransaction(tx *types.Transaction, group *transaction.ScriptGroup, key *secp256k1.Secp256k1Key) (bool, error) {
	txHash, err := tx.ComputeHash()
	if err != nil {
		return false, err
	}
	data := txHash.Bytes()
	for _, v := range group.InputIndices {
		witness := tx.Witnesses[v]
		data = append(data, types.SerializeUint64(uint64(len(witness)))...)
		data = append(data, witness...)
	}
	for i := len(tx.Inputs); i < len(tx.Witnesses); i++ {
		witness := tx.Witnesses[i]
		data = append(data, types.SerializeUint64(uint64(len(witness)))...)
		data = append(data, witness...)
	}
	msg, err := blake2b.Blake256(data)
	if err != nil {
		return false, err
	}
	signature, err := key.Sign(msg)
	if err != nil {
		return false, err
	}
	i := group.InputIndices[0]
	w := tx.Witnesses[i]
	witnessArgs, err := types.DeserializeWitnessArgs(w)
	if err != nil {
		return false, err
	}
	witnessArgs.Lock = signature
	tx.Witnesses[i] = witnessArgs.Serialize()
	return true, nil
}

func IsSingleSigMatched(key *secp256k1.Secp256k1Key, scriptArgs []byte) (bool, error) {
	if key == nil || scriptArgs == nil {
		return false, errors.New("key or scriptArgs is nil")
	}
	hash, err := blake2b.Blake160(key.PubKey())
	if err != nil {
		return false, err
	}
	return bytes.Equal(scriptArgs, hash), nil
}
