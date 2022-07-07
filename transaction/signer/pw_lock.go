package signer

import (
	"bytes"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/nervosnetwork/ckb-sdk-go/crypto/secp256k1"
	"github.com/nervosnetwork/ckb-sdk-go/transaction"
	"github.com/nervosnetwork/ckb-sdk-go/types"
)

type PWLockSigner struct {
}

func (s PWLockSigner) SignTransaction(transaction *types.Transaction, group *transaction.ScriptGroup, ctx *transaction.Context) (bool, error) {
	key := ctx.Key
	matched, err := IsSingleSigMatched(key, group.Script.Args)
	if err != nil {
		return false, err
	}
	if matched {
		// TODO
		return false, nil
	} else {
		return false, nil
	}
}

func IsPWLockMatched(key *secp256k1.Secp256k1Key, scriptArgs []byte) bool {
	if key == nil || scriptArgs == nil {
		return false
	}
	encoded := key.PubKeyUncompressed()
	hash := crypto.Keccak256(encoded[1:])
	ethAddress := hash[len(hash)-20:]
	return bytes.Equal(scriptArgs, ethAddress)
}
