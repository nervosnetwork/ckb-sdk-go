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

func NewOffChainInputCollector(Client rpc.Client) *OffChainInputCollector {
	return &OffChainInputCollector{
		Client:            Client,
		blockNumberOffset: 13,
	}
}

func (c *OffChainInputCollector) setBlockNumberOffset(blockNumberOffset uint64) {
	c.blockNumberOffset = blockNumberOffset
}

func (c *OffChainInputCollector) ApplyOffChainTransaction(tipBlockNumber uint64, transaction types.Transaction) {
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
		blockNumber := o.Value.(TransactionInputWithBlockNumber).blockNumber
		if tipBlockNumber >= blockNumber && tipBlockNumber-blockNumber <= c.blockNumberOffset {

		} else {
			c.offChainLiveCells.Remove(o)
		}
	}

	for _, tx_input := range transaction.Inputs {
		consumedOutpoint := tx_input.PreviousOutput
		c.usedLiveCells.PushBack(OutPointWithBlockNumber{consumedOutpoint, tipBlockNumber})
		next = c.offChainLiveCells.Front()
		for cell := next; cell != nil; cell = next {
			if next != nil {
				outpoint := next.Value.(TransactionInputWithBlockNumber).OutPoint
				if tx_input.PreviousOutput == outpoint {
					c.offChainLiveCells.Remove(cell)
				}
			}
			next = cell.Next()
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
