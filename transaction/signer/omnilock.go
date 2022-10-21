package signer

import (
	"bytes"
	"fmt"
	"github.com/nervosnetwork/ckb-sdk-go/crypto"
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
	if bytes.Equal(group.Script.Args, config.Args.Encode()) {
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
	return nil, nil
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
