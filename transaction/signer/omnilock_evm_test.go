package signer

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	s "github.com/nervosnetwork/ckb-sdk-go/v2/crypto/secp256k1"
	"github.com/nervosnetwork/ckb-sdk-go/v2/transaction"
	"github.com/nervosnetwork/ckb-sdk-go/v2/transaction/signer/omnilock"
	"github.com/nervosnetwork/ckb-sdk-go/v2/types"
	"github.com/stretchr/testify/assert"
)

func TestSignEVMTransaction(t *testing.T) {
	// https://testnet.explorer.nervos.org/transaction/0xdb015b26c18d3509c5c71d5a7f948fecb1ff4ffc1c6a7fe7062056d668512878
	tx := &types.Transaction{
		Version: 0,
		CellDeps: []*types.CellDep{
			{
				OutPoint: &types.OutPoint{
					TxHash: types.HexToHash("0xec18bf0d857c981c3d1f4e17999b9b90c484b303378e94de1a57b0872f5d4602"),
					Index:  0,
				},
				DepType: types.DepTypeCode,
			},
			{
				OutPoint: &types.OutPoint{
					TxHash: types.HexToHash("0xf8de3bb47d055cdf460d93a2a6e1b05f7432f9777c8c474abf4eec1d4aee5d37"),
					Index:  0,
				},
				DepType: types.DepTypeDepGroup,
			},
		},
		Inputs: []*types.CellInput{
			{
				Since: 0,
				PreviousOutput: &types.OutPoint{
					TxHash: types.HexToHash("0x69cddf7e884db14629676b0f21e03a30711e2bbf5fd21339b47eed8a74e968c4"),
					Index:  1,
				},
			},
			{
				Since: 0,
				PreviousOutput: &types.OutPoint{
					TxHash: types.HexToHash("0x9908f4d4a8664279e33e3c8104c4e1a2228be29e407bca7e31a679ac7235921f"),
					Index:  0,
				},
			},
		},
		Outputs: []*types.CellOutput{
			{
				Capacity: 6300000000,
				Lock: &types.Script{
					CodeHash: types.HexToHash("0xf329effd1c475a2978453c8600e1eaf0bc2087ee093c3ee64cc96ec6847752cb"),
					HashType: types.HashTypeType,
					Args:     common.FromHex("0x122edf94a6324817fa29e9ae83217d08527268924300"),
				},
			},
			{
				Capacity: 6300000000,
				Lock: &types.Script{
					CodeHash: types.HexToHash("0xf329effd1c475a2978453c8600e1eaf0bc2087ee093c3ee64cc96ec6847752cb"),
					HashType: types.HashTypeType,
					Args:     common.FromHex("0x127db9408ff1e32e7c58db6d795acb40d3e332241f00"),
				},
			},
		},
		OutputsData: [][]byte{
			common.FromHex("0x"),
			common.FromHex("0x"),
		},
		Witnesses: [][]byte{
			common.FromHex("0x690000001000000069000000690000005500000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"),
			common.FromHex("0x"),
		},
	}
	scriptGroup := &transaction.ScriptGroup{
		Script: &types.Script{
			CodeHash: types.HexToHash("0xf329effd1c475a2978453c8600e1eaf0bc2087ee093c3ee64cc96ec6847752cb"),
			HashType: types.HashTypeType,
			Args:     common.FromHex("0x127db9408ff1e32e7c58db6d795acb40d3e332241f00"),
		},
		GroupType:    types.ScriptTypeLock,
		InputIndices: []uint32{0, 1},
	}
	key, err := s.HexToKey("0x9002712e307e466e1c7ed1b00e142256b763b3167de32a0f6090e2bb26065bed")
	if err != nil {
		t.Error(err)
	}
	privKey := key.PrivateKey
	pubKey := privKey.PublicKey
	pubBytes := crypto.FromECDSAPub(&pubKey)[1:]
	ethAddress := crypto.Keccak256(pubBytes)[12:]

	var authContent [20]byte
	copy(authContent[:], ethAddress)

	authentication := &omnilock.Authentication{
		Flag:        omnilock.AuthFlagEVM,
		AuthContent: authContent,
	}
	config := &OmnilockConfiguration{
		Args: &omnilock.OmnilockArgs{
			Authentication: authentication,
			OmniConfig:     &omnilock.OmniConfig{},
		},
		Mode: OmnolockModeAuth,
	}
	witnesses := tx.Witnesses
	witnessArgs, err := types.DeserializeWitnessArgs(tx.Witnesses[0])
	assert.NoError(t, err)

	omniWitness, err := signForAuthMode(tx, scriptGroup, key, config)
	assert.NoError(t, err)
	assert.NotNil(t, omniWitness)

	witnessArgs.Lock = omniWitness.Serialize()
	tx.Witnesses[0] = witnessArgs.Serialize()
	expectedSignature := common.FromHex("0x690000001000000069000000690000005500000055000000100000005500000055000000410000005792dda2dff6e5fae61241fec7d3300ece2a215a8d8839da480305c1b76880247dd2e163d2d684c8f09d0cf9154666749d01a5362e8341277f2afaa8f7b7d9b400")
	assert.Equal(t, expectedSignature, witnesses[0])
}
