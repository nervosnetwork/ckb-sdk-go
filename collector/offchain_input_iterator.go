package collector

import (
	"container/list"
	"github.com/nervosnetwork/ckb-sdk-go/indexer"
	"github.com/nervosnetwork/ckb-sdk-go/types"
)

type OffChainInputIterator struct {
	iterator                    LiveCellIterator
	collector                   OffChainInputCollector
	consumeOffChainCellsFirstly bool
	isCurrentFromOffChain       bool
	current                     *types.TransactionInput
}

func newOffChainInputIterator(iterator LiveCellIterator, collector OffChainInputCollector, consumeOffChainCellsFirstly bool) CellIterator {
	return &OffChainInputIterator{
		iterator,
		collector,
		consumeOffChainCellsFirstly,
		true,
		nil,
	}
}

func (r *OffChainInputIterator) HasNext() bool {
	return r.iterator.HasNext()
}

func (r *OffChainInputIterator) Next() *types.TransactionInput {
	r.current = r.iterator.cells[r.iterator.index]
	if r.current != nil {
		if !r.isCurrentFromOffChain {
			r.iterator.index++
		}
		input := r.current
		r.current = nil
		return input
	} else {
		return nil
	}
}

func (r *OffChainInputIterator) consumeNextOffChainCell() *types.TransactionInput {
	var next *list.Element
	for it := r.collector.offChainLiveCells.Front(); it != nil; it = next {
		next = it.Next()
		if next != nil {
			if r.isTransactionInputForSearchKey(next.Value.(TransactionInputWithBlockNumber), r.iterator.SearchKey) {
				r.collector.offChainLiveCells.Remove(it)
				var result = next.Value.(TransactionInputWithBlockNumber)
				return &result.TransactionInput
			}
		}
	}
	return nil
}

func (r *OffChainInputIterator) update() bool {
	r.current = r.iterator.cells[r.iterator.index]
	if r.isCurrentFromOffChain && r.current != nil {
		return false
	}

	if r.consumeOffChainCellsFirstly {
		r.current = r.consumeNextOffChainCell()
		if r.current != nil {
			r.isCurrentFromOffChain = true
			return true
		}

	}

	r.isCurrentFromOffChain = false
	r.iterator.update()
	r.current = r.iterator.cells[r.iterator.index]
	if r.current == nil && !r.consumeOffChainCellsFirstly {
		r.current = r.consumeNextOffChainCell()
		if r.current != nil {
			r.isCurrentFromOffChain = true
		}
	}
	return true
}

func (r *OffChainInputIterator) isTransactionInputForSearchKey(transactionInputWithBlockNumber TransactionInputWithBlockNumber, searchKey *indexer.SearchKey) bool {
	if searchKey == nil {
		return false
	}

	cellOutput := transactionInputWithBlockNumber.Output
	cellOuputData := transactionInputWithBlockNumber.OutputData
	switch searchKey.ScriptType {
	case types.ScriptTypeLock:
		if cellOutput.Lock != searchKey.Script {
			return false
		}
		break
	case types.ScriptTypeType:
		if cellOutput.Type != searchKey.Script {
			return false
		}
		break
	}
	filter := searchKey.Filter
	if filter != nil {
		if filter.Script != nil {
			switch searchKey.ScriptType {
			case "lock":
				if cellOutput.Type != filter.Script {
					return false
				}
				break
			case "type":
				if cellOutput.Lock != filter.Script {
					return false
				}
			}

			if filter.OutputCapacityRange != nil {
				if cellOutput.Capacity < filter.OutputCapacityRange[0] ||
					cellOutput.Capacity >= filter.OutputCapacityRange[1] {
					return false
				}
			}
			if filter.BlockRange != nil {
				if transactionInputWithBlockNumber.blockNumber < filter.BlockRange[0] ||
					transactionInputWithBlockNumber.blockNumber >= filter.BlockRange[1] {
					return false
				}
			}

			if filter.OutputDataLenRange != nil {
				length := uint64(len(cellOuputData))
				if length < filter.OutputDataLenRange[0] ||
					length >= filter.OutputDataLenRange[1] {
					return false
				}
			}
		}
	}
	return true
}

func (r *OffChainInputIterator) GetLiveCells(searchKey *indexer.SearchKey, order indexer.SearchOrder, limit uint64, afterCursor string) (*indexer.LiveCells, error) {
	return r.iterator.LiveCellGetter.GetCells(searchKey, order, limit, afterCursor)
}
