package transaction

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	s "github.com/nervosnetwork/ckb-sdk-go/crypto/secp256k1"
	"github.com/nervosnetwork/ckb-sdk-go/mocks"
	"github.com/nervosnetwork/ckb-sdk-go/types"
	"github.com/nervosnetwork/ckb-sdk-go/utils"
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
					CodeHash: systemScripts.SecpSingleSigCell.CellHash,
					HashType: systemScripts.SecpSingleSigCell.HashType,
					Args:     common.FromHex("0x6d217d56cd45d321f64d3f38a2b21a41dbb23aa0"),
				},
			},
			{
				Capacity: 499999999000,
				Lock: &types.Script{
					CodeHash: systemScripts.SecpSingleSigCell.CellHash,
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
