package handler

import (
	"fmt"
	"github.com/nervosnetwork/ckb-sdk-go/v2/collector"
	"github.com/nervosnetwork/ckb-sdk-go/v2/systemscript"
	"github.com/nervosnetwork/ckb-sdk-go/v2/transaction"
	"github.com/nervosnetwork/ckb-sdk-go/v2/transaction/signer"
	"github.com/nervosnetwork/ckb-sdk-go/v2/transaction/signer/omnilock"
	"github.com/nervosnetwork/ckb-sdk-go/v2/types"
	"reflect"
)

type OmnilockScriptHandler struct {
	SingleSignCellDep *types.CellDep
	MultiSignCellDep  *types.CellDep
	CellDep           *types.CellDep
	CodeHash          types.Hash
}

func NewOmnilockScriptHandler(network types.Network) *OmnilockScriptHandler {
	switch network {
	case types.NetworkTest:
		return &OmnilockScriptHandler{
			SingleSignCellDep: &types.CellDep{
				OutPoint: &types.OutPoint{
					TxHash: types.HexToHash("0xf8de3bb47d055cdf460d93a2a6e1b05f7432f9777c8c474abf4eec1d4aee5d37"),
					Index:  0,
				},
				DepType: types.DepTypeDepGroup,
			},
			MultiSignCellDep: &types.CellDep{
				OutPoint: &types.OutPoint{
					TxHash: types.HexToHash("0xf8de3bb47d055cdf460d93a2a6e1b05f7432f9777c8c474abf4eec1d4aee5d37"),
					Index:  1,
				},
				DepType: types.DepTypeDepGroup,
			},
			CellDep: &types.CellDep{
				OutPoint: systemscript.GetInfo(network, systemscript.Omnilock).OutPoint,
				DepType:  systemscript.GetInfo(network, systemscript.Omnilock).DepType,
			},
			CodeHash: systemscript.GetCodeHash(network, systemscript.Omnilock),
		}
	case types.NetworkMain:
		return &OmnilockScriptHandler{
			SingleSignCellDep: &types.CellDep{
				OutPoint: &types.OutPoint{
					TxHash: types.HexToHash("0x71a7ba8fc96349fea0ed3a5c47992e3b4084b031a42264a018e0072e8172e46c"),
					Index:  0,
				},
				DepType: types.DepTypeDepGroup,
			},
			MultiSignCellDep: &types.CellDep{
				OutPoint: &types.OutPoint{
					TxHash: types.HexToHash("0x71a7ba8fc96349fea0ed3a5c47992e3b4084b031a42264a018e0072e8172e46c"),
					Index:  1,
				},
				DepType: types.DepTypeDepGroup,
			},
			CellDep: &types.CellDep{
				OutPoint: systemscript.GetInfo(network, systemscript.Omnilock).OutPoint,
				DepType:  systemscript.GetInfo(network, systemscript.Omnilock).DepType,
			},
			CodeHash: systemscript.GetCodeHash(network, systemscript.Omnilock),
		}
	default:
		return nil
	}
}

func (o *OmnilockScriptHandler) BuildTransaction(builder collector.TransactionBuilder, group *transaction.ScriptGroup, context interface{}) (bool, error) {
	if group == nil || !o.isMatched(group.Script) {
		return false, nil
	}
	if config, ok := context.(*signer.OmnilockConfiguration); ok {
		builder.AddCellDep(o.CellDep)

		switch config.Mode {
		case signer.OmnolockModeAuth:
			return o.buildTransactionForAuthMode(builder, group, config)
		case signer.OmnolockModeAdministrator:
			return o.buildTransactionForAdministratorMode(builder, group, config)
		default:
			return false, fmt.Errorf("unknown Omnilock mode %d", config.Mode)
		}
	} else {
		return false, nil
	}
}

func (o *OmnilockScriptHandler) buildTransactionForAuthMode(builder collector.TransactionBuilder, group *transaction.ScriptGroup, configuration *signer.OmnilockConfiguration) (bool, error) {
	omnilockWitnessLock := new(omnilock.OmnilockWitnessLock)
	switch configuration.Args.Authentication.Flag {
	case omnilock.AuthFlagCKBSecp256k1Blake160:
		builder.AddCellDep(o.SingleSignCellDep)
		omnilockWitnessLock.Signature = make([]byte, 65)
	case omnilock.AuthFlagEthereum:
		return false, fmt.Errorf("unsupported flag Ethereum")
	case omnilock.AuthFlagEOS:
		return false, fmt.Errorf("unsupported flag EOS")
	case omnilock.AuthFlagTRON:
		return false, fmt.Errorf("unsupported flag TRON")
	case omnilock.AuthFlagBitcoin:
		return false, fmt.Errorf("unsupported flag Bitcoin")
	case omnilock.AuthFlagDogcoin:
		return false, fmt.Errorf("unsupported flag Dogecoin")
	case omnilock.AuthFlagCKBMultiSig:
		builder.AddCellDep(o.MultiSignCellDep)
		omnilockWitnessLock.Signature = configuration.MultisigConfig.WitnessEmptyPlaceholderInLock()
	case omnilock.AuthFlagLockScriptHash:
	case omnilock.AuthFlagExec:
		return false, fmt.Errorf("unsupported flag Exec")
	case omnilock.AuthFlagDynamicLinking:
		return false, fmt.Errorf("unsupported flag Dynamic Linking")
	default:
		return false, fmt.Errorf("unknown auth flag %d", configuration.Args.Authentication.Flag)
	}
	builder.SetWitness(uint(group.InputIndices[0]), types.WitnessTypeLock, omnilockWitnessLock.SerializeAsPlaceholder())
	return true, nil
}

func (o *OmnilockScriptHandler) buildTransactionForAdministratorMode(builder collector.TransactionBuilder, group *transaction.ScriptGroup, configuration *signer.OmnilockConfiguration) (bool, error) {
	return false, fmt.Errorf("unsupported")
}

func (o *OmnilockScriptHandler) isMatched(script *types.Script) bool {
	if script == nil {
		return false
	}
	return reflect.DeepEqual(script.CodeHash, o.CodeHash)
}
