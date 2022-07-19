package signer

import (
	"github.com/nervosnetwork/ckb-sdk-go/crypto/secp256k1"
	"github.com/nervosnetwork/ckb-sdk-go/transaction"
	"github.com/nervosnetwork/ckb-sdk-go/types"
)

type AnyCanPaySigner struct {
}

func (s *AnyCanPaySigner) SignTransaction(transaction *types.Transaction, group *transaction.ScriptGroup, ctx *transaction.Context) (bool, error) {
	key := ctx.Key
	matched, err := IsAnyCanPayMatched(key, group.Script.Args)
	if err != nil {
		return false, err
	}
	if matched {
		return SingleSignTransaction(transaction, group, key)
	} else {
		return false, nil
	}
}

func IsAnyCanPayMatched(key *secp256k1.Secp256k1Key, scriptArgs []byte) (bool, error) {
	return IsSingleSigMatched(key, scriptArgs[:20])
}
