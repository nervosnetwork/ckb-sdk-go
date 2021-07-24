package mercury

import (
	"context"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/nervosnetwork/ckb-sdk-go/indexer"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/action"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/resp"
	"github.com/pkg/errors"
)

type MercuryApi interface {
	indexer.Client
	GetBalance(payload *model.GetBalancePayload) (*resp.GetBalanceResponse, error)
	BuildTransferTransaction(payload *model.TransferPayload) (*resp.TransferCompletionResponse, error)
	BuildWalletCreationTransaction(payload *model.CreateWalletPayload) (*resp.TransferCompletionResponse, error)
	RegisterAddresses(normalAddresses []string) ([]string, error)
}

type DefaultMercuryApi struct {
	indexer indexer.Client
	c       *rpc.Client
}

func (cli *DefaultMercuryApi) RegisterAddresses(normalAddresses []string) ([]string, error) {
	var scriptHash []string
	err := cli.c.Call(&scriptHash, "register_addresses", normalAddresses)
	if err != nil {
		return scriptHash, err
	}

	return scriptHash, err
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
			if item.To.Action == action.Lend_by_from || item.To.Action == action.Pay_by_to {
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

func (cli *DefaultMercuryApi) BuildWalletCreationTransaction(payload *model.CreateWalletPayload) (*resp.TransferCompletionResponse, error) {
	var resp resp.TransferCompletionResponse
	err := cli.c.Call(&resp, "build_wallet_creation_transaction", payload)
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
