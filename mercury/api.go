package mercury

import (
	"context"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/nervosnetwork/ckb-sdk-go/indexer"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/resp"
	C "github.com/nervosnetwork/ckb-sdk-go/rpc"
	"github.com/nervosnetwork/ckb-sdk-go/types"
)

type MercuryApi interface {
	C.Client
	Client
}

type DefaultMercuryApi struct {
	ckb     C.Client
	mercury Client
}

func (cli *DefaultMercuryApi) GetTransactionProof(ctx context.Context, txHashes []string, blockHash *types.Hash) (*types.TransactionProof, error) {
	return cli.ckb.GetTransactionProof(ctx, txHashes, blockHash)
}

func (cli *DefaultMercuryApi) GetBalance(payload *model.GetBalancePayload) (*resp.GetBalanceResponse, error) {
	return cli.mercury.GetBalance(payload)
}

func (cli *DefaultMercuryApi) BuildTransferTransaction(payload *model.TransferPayload) (*resp.TransferCompletionResponse, error) {
	return cli.mercury.BuildTransferTransaction(payload)
}

func (cli *DefaultMercuryApi) BuildAssetAccountCreationTransaction(payload *model.CreateAssetAccountPayload) (*resp.TransferCompletionResponse, error) {
	return cli.mercury.BuildAssetAccountCreationTransaction(payload)
}

func (cli *DefaultMercuryApi) BuildAssetCollectionTransaction(payload *model.CollectAssetPayload) (*resp.TransferCompletionResponse, error) {
	return cli.mercury.BuildAssetCollectionTransaction(payload)
}

func (cli *DefaultMercuryApi) RegisterAddresses(normalAddresses []string) ([]string, error) {
	return cli.mercury.RegisterAddresses(normalAddresses)
}

func (cli *DefaultMercuryApi) GetGenericTransaction(txHash string) (*resp.GetGenericTransactionResponse, error) {
	return cli.mercury.GetGenericTransaction(txHash)
}

func (cli *DefaultMercuryApi) GetGenericBlock(payload *model.GetGenericBlockPayload) (*resp.GenericBlockResponse, error) {
	return cli.mercury.GetGenericBlock(payload)
}

func (cli *DefaultMercuryApi) QueryGenericTransactions(payload *model.QueryGenericTransactionsPayload) (*resp.QueryGenericTransactionsResponse, error) {
	return cli.mercury.QueryGenericTransactions(payload)
}

func (cli *DefaultMercuryApi) GetTipBlockNumber(ctx context.Context) (uint64, error) {
	return cli.ckb.GetTipBlockNumber(ctx)
}

func (cli *DefaultMercuryApi) GetTipHeader(ctx context.Context) (*types.Header, error) {
	return cli.ckb.GetTipHeader(ctx)
}

func (cli *DefaultMercuryApi) GetCurrentEpoch(ctx context.Context) (*types.Epoch, error) {
	return cli.ckb.GetCurrentEpoch(ctx)
}

func (cli *DefaultMercuryApi) GetEpochByNumber(ctx context.Context, number uint64) (*types.Epoch, error) {
	return cli.ckb.GetEpochByNumber(ctx, number)
}

func (cli *DefaultMercuryApi) GetBlockHash(ctx context.Context, number uint64) (*types.Hash, error) {
	return cli.ckb.GetBlockHash(ctx, number)
}

func (cli *DefaultMercuryApi) GetBlock(ctx context.Context, hash types.Hash) (*types.Block, error) {
	return cli.ckb.GetBlock(ctx, hash)
}

func (cli *DefaultMercuryApi) GetHeader(ctx context.Context, hash types.Hash) (*types.Header, error) {
	return cli.ckb.GetHeader(ctx, hash)
}

func (cli *DefaultMercuryApi) GetHeaderByNumber(ctx context.Context, number uint64) (*types.Header, error) {
	return cli.ckb.GetHeaderByNumber(ctx, number)
}

func (cli *DefaultMercuryApi) GetLiveCell(ctx context.Context, outPoint *types.OutPoint, withData bool) (*types.CellWithStatus, error) {
	return cli.ckb.GetLiveCell(ctx, outPoint, withData)
}

func (cli *DefaultMercuryApi) GetTransaction(ctx context.Context, hash types.Hash) (*types.TransactionWithStatus, error) {
	return cli.ckb.GetTransaction(ctx, hash)
}

func (cli *DefaultMercuryApi) GetBlockEconomicState(ctx context.Context, hash types.Hash) (*types.BlockEconomicState, error) {
	return cli.ckb.GetBlockEconomicState(ctx, hash)
}

func (cli *DefaultMercuryApi) GetBlockByNumber(ctx context.Context, number uint64) (*types.Block, error) {
	return cli.ckb.GetBlockByNumber(ctx, number)
}

func (cli *DefaultMercuryApi) GetConsensus(ctx context.Context) (*types.Consensus, error) {
	return cli.ckb.GetConsensus(ctx)
}

func (cli *DefaultMercuryApi) DryRunTransaction(ctx context.Context, transaction *types.Transaction) (*types.DryRunTransactionResult, error) {
	return cli.ckb.DryRunTransaction(ctx, transaction)
}

func (cli *DefaultMercuryApi) CalculateDaoMaximumWithdraw(ctx context.Context, point *types.OutPoint, hash types.Hash) (uint64, error) {
	return cli.ckb.CalculateDaoMaximumWithdraw(ctx, point, hash)
}

func (cli *DefaultMercuryApi) EstimateFeeRate(ctx context.Context, blocks uint64) (*types.EstimateFeeRateResult, error) {
	return cli.ckb.EstimateFeeRate(ctx, blocks)
}

func (cli *DefaultMercuryApi) LocalNodeInfo(ctx context.Context) (*types.Node, error) {
	return cli.ckb.LocalNodeInfo(ctx)
}

func (cli *DefaultMercuryApi) GetPeers(ctx context.Context) ([]*types.Node, error) {
	return cli.ckb.GetPeers(ctx)
}

func (cli *DefaultMercuryApi) GetBannedAddresses(ctx context.Context) ([]*types.BannedAddress, error) {
	return cli.ckb.GetBannedAddresses(ctx)
}

func (cli *DefaultMercuryApi) SetBan(ctx context.Context, address string, command string, banTime uint64, absolute bool, reason string) error {
	return cli.ckb.SetBan(ctx, address, command, banTime, absolute, reason)
}

func (cli *DefaultMercuryApi) SendTransaction(ctx context.Context, tx *types.Transaction) (*types.Hash, error) {
	return cli.ckb.SendTransaction(ctx, tx)
}

func (cli *DefaultMercuryApi) SendTransactionNoneValidation(ctx context.Context, tx *types.Transaction) (*types.Hash, error) {
	return cli.ckb.SendTransactionNoneValidation(ctx, tx)
}

func (cli *DefaultMercuryApi) TxPoolInfo(ctx context.Context) (*types.TxPoolInfo, error) {
	return cli.ckb.TxPoolInfo(ctx)
}

func (cli *DefaultMercuryApi) GetBlockchainInfo(ctx context.Context) (*types.BlockchainInfo, error) {
	return cli.ckb.GetBlockchainInfo(ctx)
}

func (cli *DefaultMercuryApi) BatchTransactions(ctx context.Context, batch []types.BatchTransactionItem) error {
	return cli.ckb.BatchTransactions(ctx, batch)
}

func (cli *DefaultMercuryApi) BatchLiveCells(ctx context.Context, batch []types.BatchLiveCellItem) error {
	return cli.ckb.BatchLiveCells(ctx, batch)
}

func (cli *DefaultMercuryApi) CallContext(ctx context.Context, result interface{}, method string, args ...interface{}) error {
	return cli.ckb.CallContext(ctx, result, method, args)
}

func (cli *DefaultMercuryApi) GetCells(ctx context.Context, searchKey *indexer.SearchKey, order indexer.SearchOrder, limit uint64, afterCursor string) (*indexer.LiveCells, error) {
	return cli.ckb.GetCells(ctx, searchKey, order, limit, afterCursor)
}

func (cli *DefaultMercuryApi) GetTransactions(ctx context.Context, searchKey *indexer.SearchKey, order indexer.SearchOrder, limit uint64, afterCursor string) (*indexer.Transactions, error) {
	return cli.ckb.GetTransactions(ctx, searchKey, order, limit, afterCursor)
}

func (cli *DefaultMercuryApi) GetTip(ctx context.Context) (*indexer.TipHeader, error) {
	return cli.ckb.GetTip(ctx)
}

func (cli *DefaultMercuryApi) GetCellsCapacity(ctx context.Context, searchKey *indexer.SearchKey) (*indexer.Capacity, error) {
	return cli.ckb.GetCellsCapacity(ctx, searchKey)
}

func (cli *DefaultMercuryApi) Close() {
	cli.ckb.Close()
}
func NewMercuryApi(address string) (MercuryApi, error) {
	dial, err := rpc.Dial(address)
	if err != nil {
		return nil, err
	}

	indexerClient := indexer.NewClient(dial)
	mercuryClient := newClient(dial)
	ckbClient := C.NewClientWithIndexer(dial, indexerClient)

	return &DefaultMercuryApi{
		ckb:     ckbClient,
		mercury: mercuryClient}, err
}
