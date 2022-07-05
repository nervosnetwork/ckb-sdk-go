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

func (s Secp256k1Blake160SighashAllSigner) SignTransaction(transaction *types.Transaction, group *transaction.ScriptGroup, ctx *transaction.Context) (bool, error) {
	key := ctx.Key
	matched, err := IsMatch(key, group.Script.Args)
	if err != nil {
		return false, err
	}
	if matched {
		return s.signTransaction(transaction, group, key)
	} else {
		return false, nil
	}
}

func (s Secp256k1Blake160SighashAllSigner) signTransaction(transaction *types.Transaction, group *transaction.ScriptGroup, key *secp256k1.Secp256k1Key) (bool, error) {
	return true, nil
}

func IsMatch(key *secp256k1.Secp256k1Key, scriptArgs []byte) (bool, error) {
	if key == nil || scriptArgs == nil {
		return false, errors.New("key or scriptArgs is nil")
	}
	hash, err := blake2b.Blake160(key.PubKey())
	if err != nil {
		return false, err
	}
	return bytes.Equal(scriptArgs, hash), nil
}
