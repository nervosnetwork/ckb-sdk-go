package collector

import (
	"github.com/nervosnetwork/ckb-sdk-go/v2/address"
	"github.com/nervosnetwork/ckb-sdk-go/v2/rpc"
	"github.com/nervosnetwork/ckb-sdk-go/v2/types"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestOffChainInputCollector(t *testing.T) {
	input := getRandomTransactionInput(300)
	client, err := rpc.Dial("https://testnet.ckb.dev")
	if err != nil {
		t.Error(err)
	}
	collector := NewOffChainInputCollector(client)
	collector.offChainLiveCells.PushBack(input)
	tx := types.Transaction{}
	tx.Inputs = append(tx.Inputs, &types.CellInput{
		PreviousOutput: input.OutPoint,
	})
	tx.Outputs = append(tx.Outputs, getRandomOutput())
	tx.OutputsData = append(tx.OutputsData, make([]byte, 0))
	tx.Outputs = append(tx.Outputs, getRandomOutput())
	tx.OutputsData = append(tx.OutputsData, make([]byte, 0))

	assert.Equal(t, 0, collector.usedLiveCells.Len())
	assert.Equal(t, 1, collector.offChainLiveCells.Len())

	collector.ApplyOffChainTransaction(500, tx)

	assert.Equal(t, 1, collector.usedLiveCells.Len())

	txInput1 := tx.Inputs[0]
	evaluateInput := false
	for cell := collector.usedLiveCells.Front(); cell != nil; cell = cell.Next() {
		if cell.Value.(OutPointWithBlockNumber).Index == txInput1.PreviousOutput.Index &&
			cell.Value.(OutPointWithBlockNumber).TxHash == txInput1.PreviousOutput.TxHash {
			evaluateInput = true
		}
	}
	assert.True(t, evaluateInput)

	output1 := tx.Outputs[0]
	assert.Equal(t, 2, collector.offChainLiveCells.Len())
	evaluateOutput := false
	for cell := collector.offChainLiveCells.Front(); cell != nil; cell = cell.Next() {
		if cell.Value.(TransactionInputWithBlockNumber).Output.Lock == output1.Lock && cell.Value.(TransactionInputWithBlockNumber).Output.Capacity == output1.Capacity {
			evaluateOutput = true
		}
	}
	assert.True(t, evaluateOutput)

	tx = types.Transaction{}
	tx.Inputs = append(tx.Inputs, &types.CellInput{
		PreviousOutput: getRandomTransactionInput(600).OutPoint,
	})
	tx.Outputs = append(tx.Outputs, getRandomOutput())
	tx.OutputsData = append(tx.OutputsData, make([]byte, 0))
	collector.ApplyOffChainTransaction(999, tx)
	// Because 999 > 500, so clear all usedLiveCells and offChainLiveCells at first.
	assert.Equal(t, 1, collector.usedLiveCells.Len())
	assert.Equal(t, 1, collector.offChainLiveCells.Len())

	tx = types.Transaction{}
	tx.Inputs = append(tx.Inputs,
		&types.CellInput{
			PreviousOutput: collector.offChainLiveCells.Front().Value.(TransactionInputWithBlockNumber).OutPoint,
		})
	tx.Outputs = append(tx.Outputs, getRandomOutput())
	tx.OutputsData = append(tx.OutputsData, make([]byte, 0))
	addr, _ := address.Decode("ckt1qrl2cyw7ulrxu48ysexpwus46r9md670h5h73cxjh3zmxsf4gt3d5qg2d5amjwfzgtqr2l72ulxw4k8c0dpga55qjzdlm749f9ffhpwl8zc422t2hvxmtlkk299l30k6xlgccjps9pe2sfhx5y3flvtlu56lu9u6pcqqqqqqqqvykxmu")
	tx.Outputs = append(tx.Outputs, &types.CellOutput{
		Capacity: 50000000000,
		Lock:     addr.Script,
	})
	tx.OutputsData = append(tx.OutputsData, make([]byte, 0))
	collector.ApplyOffChainTransaction(1000, tx)
	assert.Equal(t, 2, collector.usedLiveCells.Len())
	assert.Equal(t, 2, collector.offChainLiveCells.Len())
}

func getRandomTransactionInput(blockNumber uint64) TransactionInputWithBlockNumber {
	return TransactionInputWithBlockNumber{
		TransactionInput: types.TransactionInput{
			OutPoint:   getRandomOutPoint(),
			Output:     getRandomOutput(),
			OutputData: make([]byte, 0),
		},
		blockNumber: blockNumber,
	}
}

func getRandomOutPoint() *types.OutPoint {
	hash := make([]byte, 32)
	rand.Read(hash)
	return &types.OutPoint{
		TxHash: types.BytesToHash(hash),
		Index:  uint32(rand.Intn(100)),
	}
}

func getRandomOutput() *types.CellOutput {
	addr, _ := address.Decode("ckt1qzda0cr08m85hc8jlnfp3zer7xulejywt49kt2rr0vthywaa50xwsq03ewkvsva4cchhntydu648l7lyvn9w2cctnpask")
	return &types.CellOutput{
		Capacity: rand.Uint64(),
		Lock:     addr.Script,
	}
}
