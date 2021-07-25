package mercury

import (
	"context"
	"errors"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/nervosnetwork/ckb-sdk-go/indexer"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/resp"
)

type MercuryApi interface {
	indexer.Client
	GetBalance(payload *model.GetBalancePayload) (*resp.GetBalanceResponse, error)
	BuildTransferTransaction(payload *model.TransferPayload) (*resp.TransferCompletionResponse, error)
	BuildAssetAccountCreationTransaction(payload *model.CreateAssetAccountPayload) (*resp.TransferCompletionResponse, error)
	BuildAssetCollectionTransaction(payload *model.CollectAssetPayload) (*resp.TransferCompletionResponse, error)
	RegisterAddresses(normalAddresses []string) ([]string, error)
	GetGenericTransaction(txHash string) (*resp.GetGenericTransactionResponse, error)
	GetGenericBlock(payload *model.GetGenericBlockPayload) (*resp.GenericBlockResponse, error)
	QueryGenericTransactions(payload *model.QueryGenericTransactionsPayload) (*resp.QueryGenericTransactionsResponse, error)
}
type DefaultMercuryApi struct {
	indexer indexer.Client
	c       *rpc.Client
}

func (cli *DefaultMercuryApi) QueryGenericTransactions(payload *model.QueryGenericTransactionsPayload) (*resp.QueryGenericTransactionsResponse, error) {
	var queryGenericTransactionsResponse resp.QueryGenericTransactionsResponse
	err := cli.c.Call(&queryGenericTransactionsResponse, "query_generic_transactions", payload)
	if err != nil {
		return &queryGenericTransactionsResponse, err
	}

	return &queryGenericTransactionsResponse, err
}

func (cli *DefaultMercuryApi) BuildAssetCollectionTransaction(payload *model.CollectAssetPayload) (*resp.TransferCompletionResponse, error) {
	var resp resp.TransferCompletionResponse

	err := cli.c.Call(&resp, "build_asset_collection_transaction", payload)
	if err != nil {
		return &resp, err
	}

	return &resp, err
}

func (cli *DefaultMercuryApi) RegisterAddresses(normalAddresses []string) ([]string, error) {
	var scriptHash []string
	err := cli.c.Call(&scriptHash, "register_addresses", normalAddresses)
	if err != nil {
		return scriptHash, err
	}

	return scriptHash, err
}

func (cli *DefaultMercuryApi) GetGenericBlock(payload *model.GetGenericBlockPayload) (*resp.GenericBlockResponse, error) {
	var block resp.GenericBlockResponse
	err := cli.c.Call(&block, "get_generic_block", payload)
	if err != nil {
		return &block, err
	}

	return &block, err
}

func (cli *DefaultMercuryApi) GetGenericTransaction(txHash string) (*resp.GetGenericTransactionResponse, error) {
	var tx resp.GetGenericTransactionResponse
	err := cli.c.Call(&tx, "get_generic_transaction", txHash)
	if err != nil {
		return &tx, err
	}

	return &tx, err
}

func (cli *DefaultMercuryApi) GetBalance(payload *model.GetBalancePayload) (*resp.GetBalanceResponse, error) {
	var balance resp.GetBalanceResponse
	err := cli.c.Call(&balance, "get_balance", payload)
	if err != nil {
		return &balance, err
	}

	return &balance, err
}

func (cli *DefaultMercuryApi) GetCells(ctx context.Context, searchKey *indexer.SearchKey, order indexer.SearchOrder, limit uint64, afterCursor string) (*indexer.LiveCells, error) {
	return cli.indexer.GetCells(ctx, searchKey, order, limit, afterCursor)
}

func (cli *DefaultMercuryApi) GetTransactions(ctx context.Context, searchKey *indexer.SearchKey, order indexer.SearchOrder, limit uint64, afterCursor string) (*indexer.Transactions, error) {
	return cli.indexer.GetTransactions(ctx, searchKey, order, limit, afterCursor)
}

func (cli *DefaultMercuryApi) GetTip(ctx context.Context) (*indexer.TipHeader, error) {
	return cli.indexer.GetTip(ctx)
}

func (cli *DefaultMercuryApi) GetCellsCapacity(ctx context.Context, searchKey *indexer.SearchKey) (*indexer.Capacity, error) {
	return cli.indexer.GetCellsCapacity(ctx, searchKey)
}

func (cli *DefaultMercuryApi) Close() {
	cli.indexer.Close()
}

func (cli *DefaultMercuryApi) BuildTransferTransaction(payload *model.TransferPayload) (*resp.TransferCompletionResponse, error) {
	var resp resp.TransferCompletionResponse
	if payload.UdtHash == "" {
		for _, item := range payload.Items {
			if !item.To.IsPayBayFrom() || !item.To.IsPayBayFrom() {
				return &resp, errors.New("The transaction does not support ckb")
			}
		}
	}

	err := cli.c.Call(&resp, "build_transfer_transaction", payload)
	if err != nil {
		return &resp, err
	}

	return &resp, err
}

func (cli *DefaultMercuryApi) BuildAssetAccountCreationTransaction(payload *model.CreateAssetAccountPayload) (*resp.TransferCompletionResponse, error) {
	var resp resp.TransferCompletionResponse
	err := cli.c.Call(&resp, "build_asset_account_creation_transaction", payload)
	if err != nil {
		return &resp, err
	}

	return &resp, err
}

func NewMercuryApi(address string) (MercuryApi, error) {
	dial, err := rpc.Dial(address)
	client, err := indexer.Dial(address)
	if err != nil {
		return nil, err
	}

	return &DefaultMercuryApi{
		indexer: client,
		c:       dial}, err
}
