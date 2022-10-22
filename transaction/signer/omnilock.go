package signer

import (
	"bytes"
	"fmt"
	"github.com/nervosnetwork/ckb-sdk-go/crypto"
	"github.com/nervosnetwork/ckb-sdk-go/crypto/blake2b"
	"github.com/nervosnetwork/ckb-sdk-go/crypto/secp256k1"
	"github.com/nervosnetwork/ckb-sdk-go/transaction"
	"github.com/nervosnetwork/ckb-sdk-go/transaction/signer/omnilock"
	"github.com/nervosnetwork/ckb-sdk-go/types"
)

type OmnilockSigner struct {
}

func (s *OmnilockSigner) SignTransaction(transaction *types.Transaction, group *transaction.ScriptGroup, ctx *transaction.Context) (bool, error) {
	config, ok := ctx.Payload.(*OmnilockConfiguration)
	if !ok {
		return false, nil
	}
	if !bytes.Equal(group.Script.Args, config.Args.Encode()) {
		return false, nil
	}
	index := group.InputIndices[0]
	witnesses := transaction.Witnesses
	witnessArgs, err := types.DeserializeWitnessArgs(witnesses[index])
	if err != nil {
		return false, err
	}
	var omnilockWitnessLock *omnilock.OmnilockWitnessLock
	switch config.Mode {
	case OmnolockModeAuth:
		omnilockWitnessLock, err = signForAuthMode(transaction, group, ctx.Key, config)
	case OmnolockModeAdministrator:
		omnilockWitnessLock, err = signForAdministratorMode(transaction, group, ctx.Key, config)
	default:
		return false, fmt.Errorf("unknown Omnilock mode %d", config.Mode)
	}
	if err != nil {
		return false, err
	}
	if omnilockWitnessLock != nil {
		witnessArgs.Lock = omnilockWitnessLock.Serialize()
		witnesses[index] = witnessArgs.Serialize()
		return true, nil
	} else {
		return false, nil
	}
}

func signForAuthMode(tx *types.Transaction, group *transaction.ScriptGroup, key crypto.Key, config *OmnilockConfiguration) (*omnilock.OmnilockWitnessLock, error) {
	authArgs := group.Script.Args[1:21]
	firstIndex := group.InputIndices[0]
	firstWitness := tx.Witnesses[firstIndex]

	witnessArgs, err := types.DeserializeWitnessArgs(firstWitness)
	if err != nil {
		return nil, err
	}
	omnilockWitnessLock := new(omnilock.OmnilockWitnessLock)
	switch config.Args.Authentication.Flag {
	case omnilock.AuthFlagCKBSecp256k1Blake160:
		hash := blake2b.Blake160(key.(*secp256k1.Secp256k1Key).PubKey())
		if !bytes.Equal(hash, authArgs) {
			return nil, nil
		}
		omnilockWitnessLock.Signature = make([]byte, 65)
		witnessArgs.Lock = make([]byte, len(omnilockWitnessLock.Serialize()))
		witnessPlaceholder := witnessArgs.Serialize()
		signature, err := SignTransaction(tx, uint32ArrayToIntArray(group.InputIndices), witnessPlaceholder, key)
		if err != nil {
			return nil, err
		}
		omnilockWitnessLock.Signature = signature
	case omnilock.AuthFlagEthereum:
		return nil, fmt.Errorf("unsupported flag Ethereum")
	case omnilock.AuthFlagEOS:
		return nil, fmt.Errorf("unsupported flag EOS")
	case omnilock.AuthFlagTRON:
		return nil, fmt.Errorf("unsupported flag TRON")
	case omnilock.AuthFlagBitcoin:
		return nil, fmt.Errorf("unsupported flag Bitcoin")
	case omnilock.AuthFlagDogcoin:
		return nil, fmt.Errorf("unsupported flag Dogecoin")
	case omnilock.AuthFlagCKBMultiSig:
		// TODO
	case omnilock.AuthFlagLockScriptHash:
	case omnilock.AuthFlagExec:
		return nil, fmt.Errorf("unsupported flag Exec")
	case omnilock.AuthFlagDynamicLinking:
		return nil, fmt.Errorf("unsupported flag Dynamic Linking")
	default:
		return nil, fmt.Errorf("unknown auth flag %d", config.Args.Authentication.Flag)
	}
	return omnilockWitnessLock, nil
}

func signForAdministratorMode(tx *types.Transaction, group *transaction.ScriptGroup, key crypto.Key, config *OmnilockConfiguration) (*omnilock.OmnilockWitnessLock, error) {
	return nil, nil
}

type OmnilockConfiguration struct {
	Args *omnilock.OmnilockArgs
	Mode OmnilockMode
}

type OmnilockMode uint

const (
	OmnolockModeAuth          OmnilockMode = 0
	OmnolockModeAdministrator OmnilockMode = 1
)
