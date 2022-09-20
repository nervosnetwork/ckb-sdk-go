package transaction

import (
	"encoding/binary"
	"fmt"
	"github.com/nervosnetwork/ckb-sdk-go/crypto"
	"github.com/nervosnetwork/ckb-sdk-go/crypto/blake2b"
	"github.com/nervosnetwork/ckb-sdk-go/types"
)

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

func CalculateTransactionFee(tx *types.Transaction, feeRate uint64) uint64 {
	txSize := tx.SizeInBlock()
	fee := txSize * feeRate / 1000
	if fee*1000 < txSize*feeRate {
		fee += 1
	}
	return fee
}
