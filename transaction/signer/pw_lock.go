package signer

import (
	"bytes"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/nervosnetwork/ckb-sdk-go/crypto/secp256k1"
	"github.com/nervosnetwork/ckb-sdk-go/transaction"
	"github.com/nervosnetwork/ckb-sdk-go/types"
	"strconv"
)

type PWLockSigner struct {
}

func (s *PWLockSigner) SignTransaction(transaction *types.Transaction, group *transaction.ScriptGroup, ctx *transaction.Context) (bool, error) {
	key := ctx.Key
	matched := IsPWLockMatched(key, group.Script.Args)
	if matched {
		return PWLockSignTransaction(transaction, group, key)
	} else {
		return false, nil
	}
}

func PWLockSignTransaction(tx *types.Transaction, group *transaction.ScriptGroup, key *secp256k1.Secp256k1Key) (bool, error) {
	txHash := tx.ComputeHash()
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
	msg := crypto.Keccak256(data)
	prefix := []byte("\u0019Ethereum Signed Message:\n" + strconv.Itoa(len(msg)))
	msg = append(prefix, msg...)
	msg = crypto.Keccak256(msg)
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

func IsPWLockMatched(key *secp256k1.Secp256k1Key, scriptArgs []byte) bool {
	if key == nil || scriptArgs == nil {
		return false
	}
	encoded := key.PubKeyUncompressed()
	hash := crypto.Keccak256(encoded[1:])
	ethAddress := hash[len(hash)-20:]
	return bytes.Equal(scriptArgs, ethAddress)
}
