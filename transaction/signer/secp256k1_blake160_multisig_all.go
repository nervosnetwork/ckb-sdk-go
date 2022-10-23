package signer

import (
	"bytes"
	"errors"
	"github.com/nervosnetwork/ckb-sdk-go/crypto/secp256k1"
	"github.com/nervosnetwork/ckb-sdk-go/systemscript"
	"github.com/nervosnetwork/ckb-sdk-go/transaction"
	"github.com/nervosnetwork/ckb-sdk-go/types"
)

type Secp256k1Blake160MultisigAllSigner struct {
}

func (s *Secp256k1Blake160MultisigAllSigner) SignTransaction(transaction *types.Transaction, group *transaction.ScriptGroup, ctx *transaction.Context) (bool, error) {
	var config *systemscript.MultisigConfig
	switch ctx.Payload.(type) {
	case systemscript.MultisigConfig:
		mm := ctx.Payload.(systemscript.MultisigConfig)
		config = &mm
	case *systemscript.MultisigConfig:
		config = ctx.Payload.(*systemscript.MultisigConfig)
	default:
		return false, nil
	}
	matched, err := IsMultiSigMatched(ctx.Key, config, group.Script.Args)
	if err != nil {
		return false, err
	}
	if matched {
		return MultiSignTransaction(transaction, uint32ArrayToIntArray(group.InputIndices), ctx.Key, config)
	} else {
		return false, nil
	}
}

func MultiSignTransaction(tx *types.Transaction, group []int, key *secp256k1.Secp256k1Key, config *systemscript.MultisigConfig) (bool, error) {
	var err error
	i0 := group[0]
	witnessPlaceholder, err := config.WitnessPlaceholder(tx.Witnesses[i0])
	if err != nil {
		return false, nil
	}
	signature, err := SignTransaction(tx, group, witnessPlaceholder, key)
	if err != nil {
		return false, err
	}
	witnessArgs, err := types.DeserializeWitnessArgs(tx.Witnesses[i0])
	if err != nil {
		return false, err
	}
	witnessArgs.Lock = setMultisigSignature(witnessArgs.Lock, signature, config)
	return true, nil
}

func setMultisigSignature(signatures []byte, signature []byte, multisigConfig *systemscript.MultisigConfig) []byte {
	offset := len(multisigConfig.Encode())
	for i := 0; i < int(multisigConfig.Threshold); i++ {
		if isEmptyByteSlice(signatures, offset, 65) {
			copy(signatures[offset:offset+65], signature[:])
			break
		}
		offset += 65
	}
	return signatures
}

func isEmptyByteSlice(lock []byte, offset int, length int) bool {
	for i := offset; i < offset+length; i++ {
		if lock[i] != 0 {
			return false
		}
	}
	return true
}

func IsMultiSigMatched(key *secp256k1.Secp256k1Key, config *systemscript.MultisigConfig, scriptArgs []byte) (bool, error) {
	if key == nil || scriptArgs == nil {
		return false, errors.New("key or scriptArgs is nil")
	}
	hash := config.Hash160()
	return bytes.Equal(scriptArgs, hash), nil
}
