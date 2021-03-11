package payment

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/google/go-cmp/cmp"
	"github.com/nervosnetwork/ckb-sdk-go/address"
	"github.com/nervosnetwork/ckb-sdk-go/indexer"
	"github.com/nervosnetwork/ckb-sdk-go/mocks"
	"github.com/nervosnetwork/ckb-sdk-go/transaction"
	"github.com/nervosnetwork/ckb-sdk-go/types"
	"github.com/nervosnetwork/ckb-sdk-go/utils"
	"github.com/stretchr/testify/assert"
	"math"
	"math/big"
	"testing"
)

func TestIssuingCheque(t *testing.T) {
	uuid := "0x683574c1275eb5cfe6f8745faa375b08bf773223fd8d2b4db28dbd90a27f1586"
	tests := []struct {
		name              string
		senderAddr        string
		receiverAddr      string
		uuid              string
		amount            string
		searchKey         *indexer.SearchKey
		searchOrder       indexer.SearchOrder
		limit             uint64
		cursor            string
		expectedLiveCells *indexer.LiveCells
		expectedTx        *types.Transaction
		chain             string
		err               error
		feeRate           uint64
	}{
		{
			"short payload address normal on mainnet",
			"ckt1qyqvgpevpyh45a7a4t0l5n7apqduw7y9y99qpyrsd5",
			"ckt1qyqrd7cglncpwfzn73qwhed5mvjnrl8v6nvq2cpmd8",
			uuid,
			"10000000000",
			genSearchKey("ckt1qyqvgpevpyh45a7a4t0l5n7apqduw7y9y99qpyrsd5"),
			indexer.SearchOrderAsc,
			indexer.SearchLimit,
			"",
			mockLiveCellsForIssuingCheque(getLock("ckt1qyqvgpevpyh45a7a4t0l5n7apqduw7y9y99qpyrsd5"), &types.Script{
				CodeHash: types.HexToHash("0x5e7a36a77e68eecc013dfa2fe6a23f3b6c344b04005808694ae6dd45eea4cfd5"),
				HashType: types.HashTypeType,
				Args:     common.FromHex(uuid),
			}, "20000000000"),
			mockUnsignedIssuingChequeTx("ckt1qyqvgpevpyh45a7a4t0l5n7apqduw7y9y99qpyrsd5", "ckt1qyqrd7cglncpwfzn73qwhed5mvjnrl8v6nvq2cpmd8", "20000000000", "10000000000", "ckb", &types.Script{
				CodeHash: types.HexToHash("0x5e7a36a77e68eecc013dfa2fe6a23f3b6c344b04005808694ae6dd45eea4cfd5"),
				HashType: types.HashTypeType,
				Args:     common.FromHex(uuid),
			}, 1000),
			"ckb",
			nil,
			1000,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockClient := &mocks.Client{}
			mockClient.On("GetCells", context.Background(), test.searchKey, test.searchOrder, test.limit, test.cursor).Return(test.expectedLiveCells, nil)
			mockClient.On("GetBlockchainInfo", context.Background()).Return(&types.BlockchainInfo{Chain: test.chain}, nil)
			systemScripts, _ := utils.NewSystemScripts(mockClient)
			c, err := NewCheque(test.senderAddr, test.receiverAddr, test.uuid, test.amount, test.feeRate, systemScripts)
			if err != nil {
				t.Fatal(err)
			}
			tx, err := c.GenerateIssuingChequeUnsignedTx(mockClient)
			if err != nil {
				t.Fatal(err)
			}
			if !compareTransaction(test.expectedTx, tx) {
				t.Fatalf("want %+v but got %+v", test.expectedTx, tx)
			}
			assert.Equal(t, test.err, err)
		})
	}
}

func compareTransaction(expectedTx, actualTx *types.Transaction) bool {
	return cmp.Equal(expectedTx, actualTx)
}

func genSearchKey(addr string) *indexer.SearchKey {
	return &indexer.SearchKey{
		Script:     getLock(addr),
		ScriptType: indexer.ScriptTypeLock,
	}
}

func getLock(addr string) *types.Script {
	parsedAddr, _ := address.Parse(addr)
	return parsedAddr.Script
}

func mockUnsignedIssuingChequeTx(senderAddr, receiverAddr, totalAmount, transferAmount string, chain string, udtType *types.Script, feeRate uint64) *types.Transaction {
	mockClient := &mocks.Client{}
	mockClient.On("GetBlockchainInfo", context.Background()).Return(&types.BlockchainInfo{Chain: chain}, nil)
	systemScripts, _ := utils.NewSystemScripts(mockClient)
	number, _ := big.NewInt(0).SetString(transferAmount, 10)
	totalN, _ := big.NewInt(0).SetString(totalAmount, 10)
	amountBytes := utils.GenerateSudtAmount(number)
	tx := transaction.NewSecp256k1SingleSigTx(systemScripts)
	tx.CellDeps = append(tx.CellDeps, &types.CellDep{
		OutPoint: systemScripts.SUDTCell.OutPoint,
		DepType:  systemScripts.SUDTCell.DepType,
	})
	args, err := utils.ChequeCellArgs(getLock(senderAddr), getLock(receiverAddr))
	if err != nil {
		return nil
	}
	tx.Outputs = append(tx.Outputs, &types.CellOutput{
		Capacity: uint64(162 * math.Pow10(8)),
		Lock: &types.Script{
			CodeHash: systemScripts.ChequeCell.CellHash,
			HashType: systemScripts.ChequeCell.HashType,
			Args:     args,
		},
		Type: udtType,
	}, &types.CellOutput{
		Capacity: 0,
		Lock:     getLock(senderAddr),
	}, &types.CellOutput{
		Capacity: uint64(142 * math.Pow10(8)),
		Lock:     getLock(senderAddr),
		Type:     udtType,
	})
	remainingAmountBytes := utils.GenerateSudtAmount(big.NewInt(0).Sub(totalN, number))
	tx.OutputsData = append(tx.OutputsData, amountBytes, []byte{}, remainingAmountBytes)
	tx.Inputs = append(tx.Inputs, &types.CellInput{
		Since: 0,
		PreviousOutput: &types.OutPoint{
			TxHash: types.HexToHash("0x8aa76892d65a1e9964b093e71bdb89b53d65f68d5c001d2199edd6c79db7f7ad"),
			Index:  0,
		},
	}, &types.CellInput{
		Since: 0,
		PreviousOutput: &types.OutPoint{
			TxHash: types.HexToHash("0x7d0ecdb8bad4064788b67dfafe71757e7caa2ad2cbe5597a02df95f8792bdb21"),
			Index:  1,
		},
	})
	tx.Witnesses = append(tx.Witnesses, transaction.EmptyWitnessArgPlaceholder, []byte{})
	fee, err := transaction.CalculateTransactionFee(tx, feeRate)
	changeCapacity := uint64((1142-162-142)*math.Pow10(8)) - fee
	tx.Outputs[1].Capacity = changeCapacity
	if err != nil {
		return nil
	}
	return tx
}

func mockLiveCellsForIssuingCheque(lock *types.Script, udtType *types.Script, amount string) *indexer.LiveCells {
	number, _ := big.NewInt(0).SetString(amount, 10)
	amountBytes := utils.GenerateSudtAmount(number)
	return &indexer.LiveCells{
		LastCursor: "",
		Objects: []*indexer.LiveCell{
			{
				BlockNumber: 1000,
				OutPoint: &types.OutPoint{
					TxHash: types.HexToHash("0x8aa76892d65a1e9964b093e71bdb89b53d65f68d5c001d2199edd6c79db7f7ad"),
					Index:  0,
				},
				Output: &types.CellOutput{
					Capacity: uint64(142 * math.Pow10(8)),
					Lock:     lock,
					Type:     udtType,
				},
				OutputData: amountBytes,
				TxIndex:    0,
			},
			{
				BlockNumber: 1000,
				OutPoint: &types.OutPoint{
					TxHash: types.HexToHash("0x7d0ecdb8bad4064788b67dfafe71757e7caa2ad2cbe5597a02df95f8792bdb21"),
					Index:  1,
				},
				Output: &types.CellOutput{
					Capacity: uint64(1000 * math.Pow10(8)),
					Lock:     lock,
				},
				TxIndex: 1,
			},
		},
	}
}
