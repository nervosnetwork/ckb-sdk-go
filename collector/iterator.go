package collector

import (
	"context"
	"github.com/nervosnetwork/ckb-sdk-go/address"
	"github.com/nervosnetwork/ckb-sdk-go/indexer"
	"github.com/nervosnetwork/ckb-sdk-go/types"
)

type CellIterator interface {
	HasNext() bool
	Next() *types.TransactionInput
}

func NewLiveCellIterator(client indexer.Client, key *indexer.SearchKey) CellIterator {
	return &LiveCellIterator{
		Client:      client,
		SearchKey:   key,
		SearchOrder: indexer.SearchOrderAsc,
		Limit:       100,
		afterCursor: "",
		cells:       nil,
		index:       0,
	}
}

func NewLiveCellIteratorFromAddress(client indexer.Client, addr string) (CellIterator, error) {
	a, err := address.Decode(addr)
	if err != nil {
		return nil, err
	}
	searchKey := &indexer.SearchKey{
		Script:     a.Script,
		ScriptType: indexer.ScriptTypeLock,
		Filter:     nil,
		WithData:   true,
	}
	return NewLiveCellIterator(client, searchKey), nil
}

type LiveCellIterator struct {
	Client      indexer.Client
	SearchKey   *indexer.SearchKey
	SearchOrder indexer.SearchOrder
	Limit       uint64
	afterCursor string
	cells       []*types.TransactionInput
	index       int
}

func (r *LiveCellIterator) HasNext() bool {
	r.update()
	return r.index < len(r.cells)
}

func (r *LiveCellIterator) Next() *types.TransactionInput {
	current := r.cells[r.index]
	r.index++
	return current
}

func (r *LiveCellIterator) update() bool {
	if r.index >= 0 && r.index < len(r.cells) {
		return false
	}
	liveCells, err := r.Client.GetCells(context.Background(), r.SearchKey, r.SearchOrder, r.Limit, r.afterCursor)
	if err != nil {
		return false
	}
	r.cells = make([]*types.TransactionInput, 0)
	for _, c := range liveCells.Objects {
		i := &types.TransactionInput{
			OutPoint:   c.OutPoint,
			Output:     c.Output,
			OutputData: c.OutputData,
		}
		r.cells = append(r.cells, i)
	}
	r.afterCursor = liveCells.LastCursor
	r.index = 0
	return true
}