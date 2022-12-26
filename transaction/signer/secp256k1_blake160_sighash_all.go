package signer

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/nervosnetwork/ckb-sdk-go/v2/crypto"
	"github.com/nervosnetwork/ckb-sdk-go/v2/crypto/blake2b"
	"github.com/nervosnetwork/ckb-sdk-go/v2/crypto/secp256k1"
	"github.com/nervosnetwork/ckb-sdk-go/v2/transaction"
	"github.com/nervosnetwork/ckb-sdk-go/v2/types"
)

type Secp256k1Blake160SighashAllSigner struct {
}

func uint32ArrayToIntArray(in []uint32) []int {
	var out []int
	for _, v := range in {
		out = append(out, int(v))
	}
	return out
}

func (s *Secp256k1Blake160SighashAllSigner) SignTransaction(tx *types.Transaction, group *transaction.ScriptGroup, ctx *transaction.Context) (bool, error) {
	matched, err := IsSingleSigMatched(ctx.Key, group.Script.Args)
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

func IsSingleSigMatched(key *secp256k1.Secp256k1Key, scriptArgs []byte) (bool, error) {
	if key == nil || scriptArgs == nil {
		return false, errors.New("key or scriptArgs is nil")
	}
	hash := blake2b.Blake160(key.PubKey())
	return bytes.Equal(scriptArgs, hash), nil
}

// SignTransaction signs transaction with index group and witness placeholder in secp256k1_blake160_sighash_all way
func SignTransaction(tx *types.Transaction, group []int, witnessPlaceholder []byte, key crypto.Key) ([]byte, error) {
	inputsLen := len(tx.Inputs)
	for i := 0; i < len(group); i++ {
		if i > 0 && group[i] <= group[i-1] {
			return nil, fmt.Errorf("group index is not in ascending order")
		}
		if group[i] > inputsLen {
			return nil, fmt.Errorf("group index %d is greater than input bytesLen %d", group[i], inputsLen)
		}
	}
	txHash := tx.ComputeHash()
	msg := txHash.Bytes()
	bytesLen := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytesLen, uint64(len(witnessPlaceholder)))
	msg = append(msg, bytesLen...)
	msg = append(msg, witnessPlaceholder...)

	var indexes []int
	for i := 1; i < len(group); i++ {
		indexes = append(indexes, group[i])
	}
	for i := inputsLen; i < len(tx.Witnesses); i++ {
		indexes = append(indexes, i)
	}
	for _, i := range indexes {
		bytes := tx.Witnesses[i]
		bytesLen := make([]byte, 8)
		binary.LittleEndian.PutUint64(bytesLen, uint64(len(bytes)))
		msg = append(msg, bytesLen...)
		msg = append(msg, bytes...)
	}

	msgHash := blake2b.Blake256(msg)
	signature, err := key.Sign(msgHash)
	if err != nil {
		return nil, err
	}
	return signature, nil
}
