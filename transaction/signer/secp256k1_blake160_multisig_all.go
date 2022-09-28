package signer

import (
	"bytes"
	"errors"
	"github.com/nervosnetwork/ckb-sdk-go/crypto/secp256k1"
	"github.com/nervosnetwork/ckb-sdk-go/systemscript"
	"github.com/nervosnetwork/ckb-sdk-go/transaction"
	"github.com/nervosnetwork/ckb-sdk-go/types"
	"reflect"
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
	if tx.Witnesses[i0], err = setSignatureToWitness(tx.Witnesses[i0], signature, config); err != nil {
		return false, err
	}
	return true, nil
}

func setSignatureToWitness(witness []byte, signature []byte, config *systemscript.MultisigConfig) ([]byte, error) {
	witnessArgs, err := types.DeserializeWitnessArgs(witness)
	if err != nil {
		return nil, err
	}
	lock := witnessArgs.Lock
	pos := len(config.Encode())
	emptySignature := [65]byte{}
	for i := 0; i < int(config.Threshold); i++ {
		if reflect.DeepEqual(emptySignature[:], lock[pos:pos+65]) {
			copy(lock[pos:pos+65], signature[:])
			break
		}
		pos += 65
	}
	witnessArgs.Lock = lock
	w := witnessArgs.Serialize()
	return w, err
}

func IsMultiSigMatched(key *secp256k1.Secp256k1Key, config *systemscript.MultisigConfig, scriptArgs []byte) (bool, error) {
	if key == nil || scriptArgs == nil {
		return false, errors.New("key or scriptArgs is nil")
	}
	hash := config.Hash160()
	return bytes.Equal(scriptArgs, hash), nil
}
