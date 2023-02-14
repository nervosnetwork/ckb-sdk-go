package lightclient

import (
	"context"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/nervosnetwork/ckb-sdk-go/v2/indexer"
	"github.com/nervosnetwork/ckb-sdk-go/v2/mocking"
	"github.com/nervosnetwork/ckb-sdk-go/v2/types"
)

type Client interface {
	SetScripts(ctx context.Context, scriptDetails []*ScriptDetail) error
	GetScripts(ctx context.Context) ([]*ScriptDetail, error)
	SendTransaction(ctx context.Context, tx *types.Transaction) (*types.Hash, error)
	GetTipHeader(ctx context.Context) (*types.Header, error)
	GetGenesisBlock(ctx context.Context) (*types.Block, error)
	GetHeader(ctx context.Context, hash types.Hash) (*types.Header, error)
	GetTransaction(ctx context.Context, hash types.Hash) (*TransactionStatus, error)
	FetchHeader(ctx context.Context, hash types.Hash) (*FetchedHeader, error)
	FetchTransaction(ctx context.Context, hash types.Hash) (*FetchedTransaction, error)
	GetPeers(ctx context.Context) ([]*types.RemoteNode, error)
	LocalNodeInfo(ctx context.Context) (*types.LocalNode, error)
	GetCells(ctx context.Context, searchKey *indexer.SearchKey, order indexer.SearchOrder, limit uint64, afterCursor string) (*indexer.LiveCells, error)
	GetTransactions(ctx context.Context, searchKey *indexer.SearchKey, order indexer.SearchOrder, limit uint64, afterCursor string) (*TxsWithCell, error)
	GetTransactionsGrouped(ctx context.Context, searchKey *indexer.SearchKey, order indexer.SearchOrder, limit uint64, afterCursor string) (*TxsWithCells, error)
	GetCellsCapacity(ctx context.Context, searchKey *indexer.SearchKey) (*indexer.Capacity, error)
	CallContext(ctx context.Context, result interface{}, method string, args ...interface{}) error
	Close()
	GetRawClient() types.GenericRPCClient
}

type client struct {
	c types.GenericRPCClient
}

func (cli *client) GetRawClient() types.GenericRPCClient {
	return cli.c
}

func (cli *client) SetScripts(ctx context.Context, scriptDetails []*ScriptDetail) error {
	err := cli.GetRawClient().CallContext(ctx, nil, "set_scripts", scriptDetails)
	return err
}

func (cli *client) GetScripts(ctx context.Context) ([]*ScriptDetail, error) {
	var result []*ScriptDetail
	err := cli.GetRawClient().CallContext(ctx, &result, "get_scripts")
	if err != nil {
		return nil, err
	}
	return result, err
}

func (cli *client) SendTransaction(ctx context.Context, tx *types.Transaction) (*types.Hash, error) {
	var result types.Hash
	err := cli.GetRawClient().CallContext(ctx, &result, "send_transaction", *tx)
	if err != nil {
		return nil, err
	}
	return &result, err
}

func (cli *client) GetTipHeader(ctx context.Context) (*types.Header, error) {
	var result types.Header
	err := cli.GetRawClient().CallContext(ctx, &result, "get_tip_header")
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (cli *client) GetGenesisBlock(ctx context.Context) (*types.Block, error) {
	var result types.Block
	err := cli.GetRawClient().CallContext(ctx, &result, "get_genesis_block")
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (cli *client) GetHeader(ctx context.Context, hash types.Hash) (*types.Header, error) {
	var result types.Header
	err := cli.GetRawClient().CallContext(ctx, &result, "get_header", hash)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (cli *client) GetTransaction(ctx context.Context, hash types.Hash) (*TransactionStatus, error) {
	var result TransactionStatus
	err := cli.GetRawClient().CallContext(ctx, &result, "get_transaction", hash)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (cli *client) FetchHeader(ctx context.Context, hash types.Hash) (*FetchedHeader, error) {
	var result FetchedHeader
	err := cli.GetRawClient().CallContext(ctx, &result, "fetch_header", hash)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (cli *client) FetchTransaction(ctx context.Context, hash types.Hash) (*FetchedTransaction, error) {
	var result FetchedTransaction
	err := cli.GetRawClient().CallContext(ctx, &result, "fetch_transaction", hash)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (cli *client) GetPeers(ctx context.Context) ([]*types.RemoteNode, error) {
	var result []*types.RemoteNode

	err := cli.GetRawClient().CallContext(ctx, &result, "get_peers")
	if err != nil {
		return nil, err
	}

	return result, err
}

func (cli *client) LocalNodeInfo(ctx context.Context) (*types.LocalNode, error) {
	var result types.LocalNode

	err := cli.GetRawClient().CallContext(ctx, &result, "local_node_info")
	if err != nil {
		return nil, err
	}

	return &result, err
}

func (cli *client) GetCells(ctx context.Context, searchKey *indexer.SearchKey, order indexer.SearchOrder, limit uint64, afterCursor string) (*indexer.LiveCells, error) {
	var (
		result indexer.LiveCells
		err    error
	)
	if afterCursor == "" {
		err = cli.GetRawClient().CallContext(ctx, &result, "get_cells", searchKey, order, hexutil.Uint64(limit))
	} else {
		err = cli.GetRawClient().CallContext(ctx, &result, "get_cells", searchKey, order, hexutil.Uint64(limit), afterCursor)
	}
	if err != nil {
		return nil, err
	}
	return &result, err

}

func (cli *client) GetTransactions(ctx context.Context, searchKey *indexer.SearchKey, order indexer.SearchOrder, limit uint64, afterCursor string) (*TxsWithCell, error) {
	var (
		result TxsWithCell
		err    error
	)
	if afterCursor == "" {
		err = cli.GetRawClient().CallContext(ctx, &result, "get_transactions", searchKey, order, hexutil.Uint64(limit))
	} else {
		err = cli.GetRawClient().CallContext(ctx, &result, "get_transactions", searchKey, order, hexutil.Uint64(limit), afterCursor)
	}
	if err != nil {
		return nil, err
	}
	return &result, err
}

func (cli *client) GetTransactionsGrouped(ctx context.Context, searchKey *indexer.SearchKey, order indexer.SearchOrder, limit uint64, afterCursor string) (*TxsWithCells, error) {
	payload := &struct {
		indexer.SearchKey
		GroupByTransaction bool `json:"group_by_transaction"`
	}{
		SearchKey:          *searchKey,
		GroupByTransaction: true,
	}
	var result TxsWithCells
	var err error
	if afterCursor == "" {
		err = cli.GetRawClient().CallContext(ctx, &result, "get_transactions", payload, order, hexutil.Uint64(limit))
	} else {
		err = cli.GetRawClient().CallContext(ctx, &result, "get_transactions", payload, order, hexutil.Uint64(limit), afterCursor)
	}
	if err != nil {
		return nil, err
	}
	return &result, err
}

func (cli *client) GetCellsCapacity(ctx context.Context, searchKey *indexer.SearchKey) (*indexer.Capacity, error) {
	var result indexer.Capacity
	err := cli.GetRawClient().CallContext(ctx, &result, "get_cells_capacity", searchKey)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (cli *client) CallContext(ctx context.Context, result interface{}, method string, args ...interface{}) error {
	err := cli.GetRawClient().CallContext(ctx, result, method, args...)
	if err != nil {
		return err
	}
	return nil
}

func (cli *client) Close() {
	cli.GetRawClient().Close()
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

func DialMockContext(ctx context.Context, url string) (Client, error) {
	c, err := mocking.DialContext(ctx, url)
	if err != nil {
		return nil, err
	}
	return NewMockingClient(c), err
}

func NewClient(c *rpc.Client) Client {
	return &client{c}
}

func NewMockingClient(c *mocking.MockClient) Client {
	return &client{c}
}
