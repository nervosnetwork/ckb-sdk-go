package transaction

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	s "github.com/nervosnetwork/ckb-sdk-go/crypto/secp256k1"
	"github.com/nervosnetwork/ckb-sdk-go/mocks"
	"github.com/nervosnetwork/ckb-sdk-go/types"
	"github.com/nervosnetwork/ckb-sdk-go/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSingleSignTransaction(t *testing.T) {
	mockClient := &mocks.Client{}
	mockClient.On("GetBlockchainInfo", context.Background()).Return(&types.BlockchainInfo{Chain: "ckb"}, nil)
	systemScripts, _ := utils.NewSystemScripts(mockClient)
	tx := &types.Transaction{
		Version: 0,
		CellDeps: []*types.CellDep{
			{
				OutPoint: systemScripts.SecpSingleSigCell.OutPoint,
				DepType:  systemScripts.SecpSingleSigCell.DepType,
			},
		},
		HeaderDeps: []types.Hash{},
		Inputs: []*types.CellInput{
			{
				Since: 0,
				PreviousOutput: &types.OutPoint{
					TxHash: types.HexToHash("0x6e9d8aa79a592d66e800acf086eec792c07f0196930665cae21608ad0d429660"),
					Index:  0,
				},
			},
			{
				Since: 0,
				PreviousOutput: &types.OutPoint{
					TxHash: types.HexToHash("0xf19a86a2fb08c83e70e294febede18358896922ad4c94e6f8498828789cc6351"),
					Index:  0,
				},
			},
			{
				Since: 0,
				PreviousOutput: &types.OutPoint{
					TxHash: types.HexToHash("0x345501155651a4ff07e4cd5ecfb8fd936e2fbb9669f2346a651d6b430f9566e0"),
					Index:  0,
				},
			},
			{
				Since: 0,
				PreviousOutput: &types.OutPoint{
					TxHash: types.HexToHash("0x6e9d8aa79a592d66e800acf086eec792c07f0196930665cae21608ad0d429660"),
					Index:  1,
				},
			},
			{
				Since: 0,
				PreviousOutput: &types.OutPoint{
					TxHash: types.HexToHash("0x4f07bd7690edbb0ca4c97f6bb68c9728108750456417cd1390affafaeb10a667"),
					Index:  0,
				},
			},
		},
		Outputs: []*types.CellOutput{
			{
				Capacity: 1000000000000,
				Lock: &types.Script{
					CodeHash: systemScripts.SecpSingleSigCell.CodeHash,
					HashType: systemScripts.SecpSingleSigCell.HashType,
					Args:     common.FromHex("0x6d217d56cd45d321f64d3f38a2b21a41dbb23aa0"),
				},
			},
			{
				Capacity: 499999999000,
				Lock: &types.Script{
					CodeHash: systemScripts.SecpSingleSigCell.CodeHash,
					HashType: systemScripts.SecpSingleSigCell.HashType,
					Args:     common.FromHex("0x6d217d56cd45d321f64d3f38a2b21a41dbb23aa0"),
				},
			},
		},
		OutputsData: [][]byte{{}, {}},
		Witnesses:   [][]byte{Secp256k1EmptyWitnessArgPlaceholder, {}, Secp256k1EmptyWitnessArgPlaceholder, {}, Secp256k1EmptyWitnessArgPlaceholder, {0, 1, 0}, {1, 123, 4}},
	}
	groupA := []int{0, 1}
	keyA, err := s.HexToKey("0948fca7a59ec4d50390271458cd993ff1d95cd8228e50310978660760e56ac8")
	if err != nil {
		t.Error(err)
	}
	err = SingleSignTransaction(tx, groupA, Secp256k1EmptyWitnessArg, keyA)
	if err != nil {
		t.Error(err)
	}
	groupB := []int{2, 3}
	keyB, err := s.HexToKey("88345757accf0728a76292a845f02c52fde7b6bd60d520bf935b883b6cb5cfca")
	if err != nil {
		t.Error(err)
	}
	err = SingleSignTransaction(tx, groupB, Secp256k1EmptyWitnessArg, keyB)
	if err != nil {
		t.Error(err)
	}
	groupC := []int{4}
	keyC, err := s.HexToKey("7089b261ad6b710fc9ba719afbd36ba8409c7d9ad8d51cc141c7151561ecae8d")
	if err != nil {
		t.Error(err)
	}
	err = SingleSignTransaction(tx, groupC, Secp256k1EmptyWitnessArg, keyC)
	if err != nil {
		t.Error(err)
	}
}

func TestMsgFromTxForMultiSig(t *testing.T) {
	// https://pudge.explorer.nervos.org/transaction/0x8b9027c407ee95f043b158b4bb5fe685b2e6159723b48712d91ec733b3068a5c
	tx := &types.Transaction{
		Version: 0,
		CellDeps: []*types.CellDep{
			{
				OutPoint: &types.OutPoint{
					TxHash: types.HexToHash("0xf8de3bb47d055cdf460d93a2a6e1b05f7432f9777c8c474abf4eec1d4aee5d37"),
					Index:  1,
				},
				DepType: types.DepTypeDepGroup,
			},
		},
		HeaderDeps: []types.Hash{},
		Inputs: []*types.CellInput{
			{
				Since: 0,
				PreviousOutput: &types.OutPoint{
					TxHash: types.HexToHash("0xb8e52009fb4dc0d63dd2a0547909bb1d66dff83e14645c70b25222c1e04ec593"),
					Index:  0,
				},
			},
		},
		Outputs: []*types.CellOutput{
			{
				Capacity: 10000000000,
				Lock: &types.Script{
					CodeHash: types.HexToHash("0x9bd7e06f3ecf4be0f2fcd2188b23f1b9fcc88e5d4b65a8637b17723bbda3cce8"),
					HashType: types.HashTypeType,
					Args:     common.FromHex("0x6cfd0e42b63a6fccf5eda9cef74d5fd0537fd55a"),
				},
			},
			{
				Capacity: 989900000000,
				Lock: &types.Script{
					CodeHash: types.HexToHash("0x5c5069eb0857efc65e1bca0c07df34c31663b3622fd3876c876320fc9634e2a8"),
					HashType: types.HashTypeType,
					Args:     common.FromHex("0x35ed7b939b4ac9cb447b82340fd8f26d344f7a62"),
				},
			},
		},
		OutputsData: [][]byte{{}, {}},
		Witnesses: [][]byte{
			common.FromHex("0x10000000100000001000000010000000"),
			[]byte{0x12, 0x34},
		},
	}
	multisigScript := common.FromHex("000002029b41c025515b00c24e2e2042df7b221af5c1891fe732dcd15b7618eb1d7a11e6a68e4579b5be0114")
	hash, _ := MsgFromTxForMultiSig(tx, []int{0}, multisigScript)
	key, _ := s.HexToKey("5271b0e474609ee280eb6ba07895718863a0eb8f114afd7217fa371fd48f6941")
	signature, _ := key.Sign(hash)
	expectedSignature := common.FromHex("bf990766e3efa8253c58330f4366bef09f49cbe4efa47b2b491541ad919c90c33ca6763c5780693efd0efea89c1645e8992520e8e551c1dec50fc41fac14b3a401")
	assert.Equal(t, expectedSignature, signature)
}
