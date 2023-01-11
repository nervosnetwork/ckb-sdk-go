package collector

import (
	"container/list"
	"github.com/nervosnetwork/ckb-sdk-go/v2/indexer"
	"github.com/nervosnetwork/ckb-sdk-go/v2/rpc"
	"github.com/nervosnetwork/ckb-sdk-go/v2/types"
)

type OffChainInputIterator struct {
	Iterator                    *LiveCellIterator
	Collector                   *OffChainInputCollector
	ConsumeOffChainCellsFirstly bool
	isCurrentFromOffChain       bool
	current                     *types.TransactionInput
}

func NewOffChainInputIterator(iterator CellIterator, collector *OffChainInputCollector, consumeOffChainCellsFirstly bool) *OffChainInputIterator {
	return &OffChainInputIterator{
		Iterator:                    iterator.(*LiveCellIterator),
		Collector:                   collector,
		ConsumeOffChainCellsFirstly: consumeOffChainCellsFirstly,
		isCurrentFromOffChain:       true,
	}
}

func NewOffChainInputIteratorFromAddress(client rpc.Client, addr string, collector *OffChainInputCollector, consumeOffChainCellsFirstly bool) (*OffChainInputIterator, error) {
	iterator, err := NewLiveCellIteratorFromAddress(client, addr)
	if err != nil {
		return nil, err
	}
	return &OffChainInputIterator{
		Iterator:                    iterator.(*LiveCellIterator),
		Collector:                   collector,
		ConsumeOffChainCellsFirstly: consumeOffChainCellsFirstly,
		isCurrentFromOffChain:       true,
	}, nil
}

func (r *OffChainInputIterator) HasNext() bool {
	return r.Iterator.HasNext()
}

func (r *OffChainInputIterator) Next() *types.TransactionInput {
	r.update()
	r.current = r.Iterator.cells[r.Iterator.index]
	if r.current != nil {
		if !r.isCurrentFromOffChain {
			r.Iterator.index++
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
	for it := r.Collector.offChainLiveCells.Front(); it != nil; it = next {
		if it != nil {
			if r.isTransactionInputForSearchKey(next.Value.(TransactionInputWithBlockNumber), r.Iterator.SearchKey) {
				r.Collector.offChainLiveCells.Remove(it)
				var result = next.Value.(TransactionInputWithBlockNumber)
				return &result.TransactionInput
			}
			next = it.Next()
		}
	}
	return nil
}

func (r *OffChainInputIterator) update() bool {
	//r.current = r.Iterator.cells[r.Iterator.index]
	if r.isCurrentFromOffChain && r.current != nil {
		return false
	}

	if r.ConsumeOffChainCellsFirstly {
		r.current = r.consumeNextOffChainCell()
		if r.current != nil {
			r.isCurrentFromOffChain = true
			return true
		}

	}

	r.isCurrentFromOffChain = false
	r.Iterator.update()
	r.current = r.Iterator.cells[r.Iterator.index]
	if r.current == nil && !r.ConsumeOffChainCellsFirstly {
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
	cellOutputData := transactionInputWithBlockNumber.OutputData
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
				length := uint64(len(cellOutputData))
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
	return r.Iterator.LiveCellGetter.GetCells(searchKey, order, limit, afterCursor)
}
