package collector

import (
	"context"
	"github.com/nervosnetwork/ckb-sdk-go/address"
	"github.com/nervosnetwork/ckb-sdk-go/indexer"
	"github.com/nervosnetwork/ckb-sdk-go/lightclient"
	"github.com/nervosnetwork/ckb-sdk-go/rpc"
	"github.com/nervosnetwork/ckb-sdk-go/types"
)

type CellIterator interface {
	HasNext() bool
	Next() *types.TransactionInput
}

type LiveCellsGetter interface {
	GetCells(searchKey *indexer.SearchKey, order indexer.SearchOrder, limit uint64, afterCursor string) (*indexer.LiveCells, error)
}

type CkbLiveCellGetter struct {
	Client  rpc.Client
	Context context.Context
}

func (c *CkbLiveCellGetter) GetCells(searchKey *indexer.SearchKey, order indexer.SearchOrder, limit uint64, afterCursor string) (*indexer.LiveCells, error) {
	ctx := c.Context
	if ctx == nil {
		ctx = context.Background()
	}
	return c.Client.GetCells(ctx, searchKey, order, limit, afterCursor)
}

type LightClientLiveCellGetter struct {
	Client  lightclient.Client
	Context context.Context
}

func (c *LightClientLiveCellGetter) GetCells(searchKey *indexer.SearchKey, order indexer.SearchOrder, limit uint64, afterCursor string) (*indexer.LiveCells, error) {
	ctx := c.Context
	if ctx == nil {
		ctx = context.Background()
	}
	return c.Client.GetCells(ctx, searchKey, order, limit, afterCursor)
}

func newLiveCellIterator(getter LiveCellsGetter, key *indexer.SearchKey) CellIterator {
	return &LiveCellIterator{
		LiveCellGetter: getter,
		SearchKey:      key,
		SearchOrder:    indexer.SearchOrderAsc,
		Limit:          100,
		afterCursor:    "",
		cells:          nil,
		index:          0,
	}
}

func newLiveCellIteratorFromAddress(getter LiveCellsGetter, addr string) (CellIterator, error) {
	a, err := address.Decode(addr)
	if err != nil {
		return nil, err
	}
	searchKey := &indexer.SearchKey{
		Script:     a.Script,
		ScriptType: types.ScriptTypeLock,
		Filter:     nil,
		WithData:   true,
	}
	return newLiveCellIterator(getter, searchKey), nil
}

func NewLiveCellIterator(client rpc.Client, key *indexer.SearchKey) CellIterator {
	return newLiveCellIterator(&CkbLiveCellGetter{Client: client}, key)
}

func NewLiveCellIteratorFromAddress(client rpc.Client, addr string) (CellIterator, error) {
	return newLiveCellIteratorFromAddress(&CkbLiveCellGetter{Client: client}, addr)
}

func NewLiveCellIteratorByLightClient(client lightclient.Client, key *indexer.SearchKey) CellIterator {
	return newLiveCellIterator(&LightClientLiveCellGetter{Client: client}, key)
}

func NewLiveCellIteratorByLightClientFromAddress(client lightclient.Client, addr string) (CellIterator, error) {
	return newLiveCellIteratorFromAddress(&LightClientLiveCellGetter{Client: client}, addr)
}

type LiveCellIterator struct {
	LiveCellGetter LiveCellsGetter
	SearchKey      *indexer.SearchKey
	SearchOrder    indexer.SearchOrder
	Limit          uint64
	afterCursor    string
	cells          []*types.TransactionInput
	index          int
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
	liveCells, err := r.LiveCellGetter.GetCells(r.SearchKey, r.SearchOrder, r.Limit, r.afterCursor)
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
