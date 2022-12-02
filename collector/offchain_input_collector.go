package collector

import (
	"container/list"
	"github.com/nervosnetwork/ckb-sdk-go/rpc"
	"github.com/nervosnetwork/ckb-sdk-go/types"
)

type OffChainInputCollector struct {
	Client            rpc.Client
	blockNumberOffset uint64
	usedLiveCells     list.List
	offChainLiveCells list.List
}

type OutPointWithBlockNumber struct {
	*types.OutPoint
	blockNumber uint64
}

type TransactionInputWithBlockNumber struct {
	types.TransactionInput
	blockNumber uint64
}

func (c *OffChainInputCollector) setBlockNumberOffset(blockNumberOffset uint64) {
	c.blockNumberOffset = blockNumberOffset
}

func (c *OffChainInputCollector) applyOffChainTransaction(tipBlockNumber uint64, transaction types.Transaction) {
	transactionHash := transaction.ComputeHash()
	var next *list.Element
	for o := c.usedLiveCells.Front(); o != nil; o = next {
		next = o.Next()
		blockNumber := o.Value.(OutPointWithBlockNumber).blockNumber
		if tipBlockNumber >= blockNumber && tipBlockNumber-blockNumber <= c.blockNumberOffset {
			// keeps
		} else {
			c.usedLiveCells.Remove(o)
		}
	}
	next = nil
	for o := c.offChainLiveCells.Front(); o != nil; o = next {
		next = o.Next()
		if tipBlockNumber >= o.Value.(TransactionInputWithBlockNumber).blockNumber && tipBlockNumber-o.Value.(TransactionInputWithBlockNumber).blockNumber <= c.blockNumberOffset {

		} else {
			c.offChainLiveCells.Remove(o)
		}
	}

	for _, tx_input := range transaction.Inputs {
		c.usedLiveCells.PushBack(OutPointWithBlockNumber{tx_input.PreviousOutput, tipBlockNumber})
		next = nil
		for cell := c.offChainLiveCells.Front(); cell != nil; cell = next {
			next = cell.Next()
			if next != nil && tx_input.PreviousOutput == next.Value.(TransactionInputWithBlockNumber).OutPoint {
				c.offChainLiveCells.Remove(cell)
			}
		}
	}

	for i, txOutput := range transaction.Outputs {
		c.offChainLiveCells.PushBack(TransactionInputWithBlockNumber{
			TransactionInput: types.TransactionInput{
				OutPoint: &types.OutPoint{
					TxHash: transactionHash, Index: uint32(i),
				},
				Output:     txOutput,
				OutputData: transaction.OutputsData[i],
			},
			blockNumber: tipBlockNumber,
		})
	}

}
