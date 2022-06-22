package indexer

import (
	"context"
	"github.com/nervosnetwork/ckb-sdk-go/types"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rpc"
)

const SearchLimit uint64 = 1000

type Client interface {
	// GetCells returns the live cells collection by the lock or type script.
	GetCells(ctx context.Context, searchKey *SearchKey, order SearchOrder, limit uint64, afterCursor string) (*LiveCells, error)

	// GetTransactions returns the transactions collection by the lock or type script.
	GetTransactions(ctx context.Context, searchKey *SearchKey, order SearchOrder, limit uint64, afterCursor string) (*Transactions, error)

	//GetTip returns the latest height processed by indexer
	GetTip(ctx context.Context) (*TipHeader, error)

	//GetCellsCapacity returns the live cells capacity by the lock or type script.
	GetCellsCapacity(ctx context.Context, searchKey *SearchKey) (*Capacity, error)

	// Close close client
	Close()
}

type client struct {
	c *rpc.Client
}

func Dial(url string) (Client, error) {
	return DialContext(context.Background(), url)
}

func DialContext(ctx context.Context, url string) (Client, error) {
	c, err := rpc.DialContext(ctx, url)
	if err != nil {
		return nil, err
	}
	return NewClient(c), nil
}

func NewClient(c *rpc.Client) Client {
	return &client{c}
}

func (cli *client) Close() {
	cli.c.Close()
}

func (cli *client) GetCells(ctx context.Context, searchKey *SearchKey, order SearchOrder, limit uint64, afterCursor string) (*LiveCells, error) {
	var result LiveCells
	var err error
	if afterCursor == "" {
		err = cli.c.CallContext(ctx, &result, "get_cells", searchKey, order, hexutil.Uint64(limit))
	} else {
		err = cli.c.CallContext(ctx, &result, "get_cells", searchKey, order, hexutil.Uint64(limit), afterCursor)
	}
	if err != nil {
		return nil, err
	}
	return &result, err
}

func (cli *client) GetTransactions(ctx context.Context, searchKey *SearchKey, order SearchOrder, limit uint64, afterCursor string) (*Transactions, error) {
	var result Transactions
	var err error
	if afterCursor == "" {
		err = cli.c.CallContext(ctx, &result, "get_transactions", searchKey, order, hexutil.Uint64(limit))
	} else {
		err = cli.c.CallContext(ctx, &result, "get_transactions", searchKey, order, hexutil.Uint64(limit), afterCursor)
	}
	if err != nil {
		return nil, err
	}
	return &result, err
}

func (cli *client) GetTip(ctx context.Context) (*TipHeader, error) {
	var result struct {
		BlockHash   types.Hash     `json:"block_hash"`
		BlockNumber hexutil.Uint64 `json:"block_number"`
	}
	err := cli.c.CallContext(ctx, &result, "get_tip")
	if err != nil {
		return nil, err
	}
	return &TipHeader{
		BlockHash:   result.BlockHash,
		BlockNumber: uint64(result.BlockNumber),
	}, nil
}

func (cli *client) GetCellsCapacity(ctx context.Context, searchKey *SearchKey) (*Capacity, error) {
	var result struct {
		Capacity    hexutil.Uint64 `json:"capacity"`
		BlockHash   types.Hash     `json:"block_hash"`
		BlockNumber hexutil.Uint64 `json:"block_number"`
	}
	err := cli.c.CallContext(ctx, &result, "get_cells_capacity", searchKey)
	if err != nil {
		return nil, err
	}
	return &Capacity{
		Capacity:    uint64(result.Capacity),
		BlockHash:   result.BlockHash,
		BlockNumber: uint64(result.BlockNumber),
	}, nil
}
