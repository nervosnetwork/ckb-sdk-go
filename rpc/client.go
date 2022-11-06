package rpc

import (
	"context"
	"errors"
	"github.com/nervosnetwork/ckb-sdk-go/indexer"
	"reflect"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rpc"

	"github.com/nervosnetwork/ckb-sdk-go/types"
)

var (
	NotFound = errors.New("not found")
)

// Client for the Nervos RPC API.
type Client interface {
	////// Chain
	// GetTipBlockNumber returns the number of blocks in the longest blockchain.
	GetTipBlockNumber(ctx context.Context) (uint64, error)

	// GetTipHeader returns the information about the tip header of the longest.
	GetTipHeader(ctx context.Context) (*types.Header, error)

	// GetCurrentEpoch returns the information about the current epoch.
	GetCurrentEpoch(ctx context.Context) (*types.Epoch, error)

	// GetEpochByNumber return the information corresponding the given epoch number.
	GetEpochByNumber(ctx context.Context, number uint64) (*types.Epoch, error)

	// GetBlockHash returns the hash of a block in the best-block-chain by block number; block of No.0 is the genesis block.
	GetBlockHash(ctx context.Context, number uint64) (*types.Hash, error)

	// GetBlock returns the information about a block by hash.
	GetBlock(ctx context.Context, hash types.Hash) (*types.Block, error)

	// GetHeader returns the information about a block header by hash.
	GetHeader(ctx context.Context, hash types.Hash) (*types.Header, error)

	// GetHeaderByNumber returns the information about a block header by block number.
	GetHeaderByNumber(ctx context.Context, number uint64) (*types.Header, error)

	// GetLiveCell returns the information about a cell by out_point if it is live.
	// If second with_data argument set to true, will return cell data and data_hash if it is live.
	GetLiveCell(ctx context.Context, outPoint *types.OutPoint, withData bool) (*types.CellWithStatus, error)

	// GetTransaction returns the information about a transaction requested by transaction hash.
	GetTransaction(ctx context.Context, hash types.Hash) (*types.TransactionWithStatus, error)

	// GetBlockEconomicState return block economic state, It includes the rewards details and when it is finalized.
	GetBlockEconomicState(ctx context.Context, hash types.Hash) (*types.BlockEconomicState, error)

	// GetTransactionProof Returns a Merkle proof that transactions are included in a block.
	GetTransactionProof(ctx context.Context, txHashes []string, blockHash *types.Hash) (*types.TransactionProof, error)

	//VerifyTransactionProof verifies that a proof points to transactions in a block, returning the transaction hashes it commits to.
	VerifyTransactionProof(ctx context.Context, proof *types.TransactionProof) ([]*types.Hash, error)

	// GetBlockByNumber get block by number
	GetBlockByNumber(ctx context.Context, number uint64) (*types.Block, error)

	// GetForkBlock The RPC returns a fork block or null. When the RPC returns a block, the block hash must equal to the parameter block_hash.
	GetForkBlock(ctx context.Context, blockHash types.Hash) (*types.Block, error)

	// GetConsensus Return various consensus parameters.
	GetConsensus(ctx context.Context) (*types.Consensus, error)

	// GetBlockMedianTime When the given block hash is not on the current canonical chain, this RPC returns null;
	// otherwise returns the median time of the consecutive 37 blocks where the given block_hash has the highest height.
	// Note that the given block is included in the median time. The included block number range is [MAX(block - 36, 0), block].
	GetBlockMedianTime(ctx context.Context, blockHash types.Hash) (uint64, error)

	////// Experiment
	// DryRunTransaction dry run transaction and return the execution cycles.
	// This method will not check the transaction validity,
	// but only run the lock script and type script and then return the execution cycles.
	// Used to debug transaction scripts and query how many cycles the scripts consume.
	// Deprecated
	DryRunTransaction(ctx context.Context, transaction *types.Transaction) (*types.DryRunTransactionResult, error)

	EstimateCycles(ctx context.Context, transaction *types.Transaction) (*types.EstimateCycles, error)

	// CalculateDaoMaximumWithdraw calculate the maximum withdraw one can get, given a referenced DAO cell, and a withdraw block hash.
	CalculateDaoMaximumWithdraw(ctx context.Context, point *types.OutPoint, hash types.Hash) (uint64, error)

	////// Net
	// LocalNodeInfo returns the local node information.
	LocalNodeInfo(ctx context.Context) (*types.LocalNode, error)

	// GetPeers returns the connected peers information.
	GetPeers(ctx context.Context) ([]*types.RemoteNode, error)

	// GetBannedAddresses returns all banned IPs/Subnets.
	GetBannedAddresses(ctx context.Context) ([]*types.BannedAddress, error)

	// ClearBannedAddress returns all banned IPs/Subnets.
	ClearBannedAddresses(ctx context.Context) error

	// SetBan insert or delete an IP/Subnet from the banned list
	SetBan(ctx context.Context, address string, command string, banTime uint64, absolute bool, reason string) error

	// SyncState returns chain synchronization state of this node.
	SyncState(ctx context.Context) (*types.SyncState, error)

	// SetNetworkActive state - true to enable networking, false to disable
	SetNetworkActive(ctx context.Context, state bool) error

	// AddNode Attempts to add a node to the peers list and try connecting to it
	AddNode(ctx context.Context, peerId, address string) error

	// RemoveNode Attempts to remove a node from the peers list and try disconnecting from it.
	RemoveNode(ctx context.Context, peerId string) error

	// PingPeers Requests that a ping is sent to all connected peers, to measure ping time.
	PingPeers(ctx context.Context) error

	////// Pool
	// SendTransaction send new transaction into transaction pool.
	SendTransaction(ctx context.Context, tx *types.Transaction) (*types.Hash, error)

	// TxPoolInfo return the transaction pool information
	TxPoolInfo(ctx context.Context) (*types.TxPoolInfo, error)

	// GetRawTxPool Returns all transaction ids in tx pool as a json array of string transaction ids.
	GetRawTxPool(ctx context.Context) (*types.RawTxPool, error)

	// ClearTxPool Removes all transactions from the transaction pool.
	ClearTxPool(ctx context.Context) error

	////// Stats
	// GetBlockchainInfo return state info of blockchain
	GetBlockchainInfo(ctx context.Context) (*types.BlockchainInfo, error)

	////// Batch
	BatchTransactions(ctx context.Context, batch []types.BatchTransactionItem) error

	// Batch Live cells
	BatchLiveCells(ctx context.Context, batch []types.BatchLiveCellItem) error

	// GetCells returns the live cells collection by the lock or type script.
	GetCells(ctx context.Context, searchKey *indexer.SearchKey, order indexer.SearchOrder, limit uint64, afterCursor string) (*indexer.LiveCells, error)

	// GetTransactions returns the transactions collection by the lock or type script.
	GetTransactions(ctx context.Context, searchKey *indexer.SearchKey, order indexer.SearchOrder, limit uint64, afterCursor string) (*indexer.TxsWithCell, error)

	// GetTransactionsGrouped returns the grouped transactions collection by the lock or type script.
	GetTransactionsGrouped(ctx context.Context, searchKey *indexer.SearchKey, order indexer.SearchOrder, limit uint64, afterCursor string) (*indexer.TxsWithCells, error)

	//GetTip returns the latest height processed by indexer
	GetIndexerTip(ctx context.Context) (*indexer.TipHeader, error)

	//GetCellsCapacity returns the live cells capacity by the lock or type script.
	GetCellsCapacity(ctx context.Context, searchKey *indexer.SearchKey) (*indexer.Capacity, error)

	// Close close client
	Close()

	CallContext(ctx context.Context, result interface{}, method string, args ...interface{}) error
}

type client struct {
	c *rpc.Client
}

func (cli *client) CallContext(ctx context.Context, result interface{}, method string, args ...interface{}) error {
	err := cli.c.CallContext(ctx, result, method, args...)
	if err != nil {
		return err
	}
	return nil
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

// Chain RPC

func (cli *client) GetTipBlockNumber(ctx context.Context) (uint64, error) {
	var num hexutil.Uint64
	err := cli.c.CallContext(ctx, &num, "get_tip_block_number")
	if err != nil {
		return 0, err
	}
	return uint64(num), err
}

func (cli *client) GetTipHeader(ctx context.Context) (*types.Header, error) {
	var result types.Header
	err := cli.c.CallContext(ctx, &result, "get_tip_header")
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (cli *client) GetCurrentEpoch(ctx context.Context) (*types.Epoch, error) {
	var result types.Epoch
	err := cli.c.CallContext(ctx, &result, "get_current_epoch")
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (cli *client) GetEpochByNumber(ctx context.Context, number uint64) (*types.Epoch, error) {
	var result types.Epoch
	err := cli.c.CallContext(ctx, &result, "get_epoch_by_number", hexutil.Uint64(number))
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (cli *client) GetBlockHash(ctx context.Context, number uint64) (*types.Hash, error) {
	var result types.Hash

	err := cli.c.CallContext(ctx, &result, "get_block_hash", hexutil.Uint64(number))
	if err != nil {
		return nil, err
	}

	return &result, err
}

func (cli *client) GetBlock(ctx context.Context, hash types.Hash) (*types.Block, error) {
	var result types.Block
	err := cli.c.CallContext(ctx, &result, "get_block", hash)
	if err != nil {
		return nil, err
	}
	if (reflect.DeepEqual(result, types.Block{})) {
		return nil, NotFound
	}
	return &result, nil
}

func (cli *client) GetHeader(ctx context.Context, hash types.Hash) (*types.Header, error) {
	var result types.Header
	err := cli.c.CallContext(ctx, &result, "get_header", hash)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (cli *client) GetHeaderByNumber(ctx context.Context, number uint64) (*types.Header, error) {
	var result types.Header
	err := cli.c.CallContext(ctx, &result, "get_header_by_number", hexutil.Uint64(number))
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (cli *client) GetTransactionProof(ctx context.Context, txHashes []string, blockHash *types.Hash) (*types.TransactionProof, error) {
	var transactionProof types.TransactionProof
	err := cli.c.CallContext(ctx, &transactionProof, "get_transaction_proof", txHashes, blockHash)
	if err != nil {
		return nil, err
	}

	return &transactionProof, err
}

func (cli *client) VerifyTransactionProof(ctx context.Context, proof *types.TransactionProof) ([]*types.Hash, error) {
	var result []*types.Hash
	err := cli.c.CallContext(ctx, &result, "verify_transaction_proof", *proof)
	if err != nil {
		return result, err
	}

	return result, err
}

func (cli *client) GetLiveCell(ctx context.Context, point *types.OutPoint, withData bool) (*types.CellWithStatus, error) {
	var result types.CellWithStatus
	err := cli.c.CallContext(ctx, &result, "get_live_cell", *point, withData)
	if err != nil {
		return nil, err
	}
	return &result, err
}

func (cli *client) GetTransaction(ctx context.Context, hash types.Hash) (*types.TransactionWithStatus, error) {
	var result types.TransactionWithStatus
	err := cli.c.CallContext(ctx, &result, "get_transaction", hash)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (cli *client) GetBlockByNumber(ctx context.Context, number uint64) (*types.Block, error) {
	var result types.Block
	err := cli.c.CallContext(ctx, &result, "get_block_by_number", hexutil.Uint64(number))
	if err != nil {
		return nil, err
	}
	if (reflect.DeepEqual(result, types.Block{})) {
		return nil, NotFound
	}
	return &result, nil
}

func (cli *client) GetForkBlock(ctx context.Context, blockHash types.Hash) (*types.Block, error) {
	var block types.Block
	err := cli.c.CallContext(ctx, &block, "get_fork_block", blockHash)
	if err != nil {
		return nil, nil
	}

	if block.Header.Hash.String() == "0x0000000000000000000000000000000000000000000000000000000000000000" {
		return nil, nil
	}
	return &block, nil
}

func (cli *client) DryRunTransaction(ctx context.Context, transaction *types.Transaction) (*types.DryRunTransactionResult, error) {
	var result types.DryRunTransactionResult
	err := cli.c.CallContext(ctx, &result, "dry_run_transaction", *transaction)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (cli *client) EstimateCycles(ctx context.Context, transaction *types.Transaction) (*types.EstimateCycles, error) {
	var result types.EstimateCycles
	err := cli.c.CallContext(ctx, &result, "estimate_cycles", *transaction)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (cli *client) CalculateDaoMaximumWithdraw(ctx context.Context, point *types.OutPoint, hash types.Hash) (uint64, error) {
	var result hexutil.Uint64
	err := cli.c.CallContext(ctx, &result, "calculate_dao_maximum_withdraw", *point, hash)
	if err != nil {
		return 0, err
	}

	return uint64(result), err
}

func (cli *client) GetConsensus(ctx context.Context) (*types.Consensus, error) {
	var result types.Consensus
	err := cli.c.CallContext(ctx, &result, "get_consensus")
	if err != nil {
		return nil, nil
	}
	return &result, nil
}

func (cli *client) GetBlockMedianTime(ctx context.Context, blockHash types.Hash) (uint64, error) {
	var result hexutil.Uint64
	err := cli.c.CallContext(ctx, &result, "get_block_median_time", blockHash)
	if err != nil {
		return uint64(result), nil
	}
	return uint64(result), nil
}

func (cli *client) LocalNodeInfo(ctx context.Context) (*types.LocalNode, error) {
	var result types.LocalNode

	err := cli.c.CallContext(ctx, &result, "local_node_info")
	if err != nil {
		return nil, err
	}

	return &result, err
}

func (cli *client) GetPeers(ctx context.Context) ([]*types.RemoteNode, error) {
	var result []*types.RemoteNode

	err := cli.c.CallContext(ctx, &result, "get_peers")
	if err != nil {
		return nil, err
	}

	return result, err
}

func (cli *client) GetBannedAddresses(ctx context.Context) ([]*types.BannedAddress, error) {
	var result []*types.BannedAddress
	err := cli.c.CallContext(ctx, &result, "get_banned_addresses")
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (cli *client) ClearBannedAddresses(ctx context.Context) error {
	return cli.c.CallContext(ctx, nil, "clear_banned_addresses")
}

func (cli *client) SetBan(ctx context.Context, address string, command string, banTime uint64, absolute bool, reason string) error {
	return cli.c.CallContext(ctx, nil, "set_ban", address, command, hexutil.Uint64(banTime), absolute, reason)
}

func (cli *client) SyncState(ctx context.Context) (*types.SyncState, error) {
	var result types.SyncState
	err := cli.c.CallContext(ctx, &result, "sync_state")
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (cli *client) SetNetworkActive(ctx context.Context, state bool) error {
	err := cli.c.CallContext(ctx, nil, "set_network_active", state)
	if err != nil {
		return err
	}
	return err
}

func (cli *client) AddNode(ctx context.Context, peerId, address string) error {
	err := cli.c.CallContext(ctx, nil, "add_node", peerId, address)
	if err != nil {
		return err
	}
	return err
}

func (cli *client) RemoveNode(ctx context.Context, peerId string) error {
	err := cli.c.CallContext(ctx, nil, "remove_node", peerId)
	if err != nil {
		return err
	}
	return err
}

func (cli *client) PingPeers(ctx context.Context) error {
	err := cli.c.CallContext(ctx, nil, "ping_peers")
	if err != nil {
		return err
	}
	return err
}

func (cli *client) SendTransaction(ctx context.Context, tx *types.Transaction) (*types.Hash, error) {
	var result types.Hash

	err := cli.c.CallContext(ctx, &result, "send_transaction", *tx, "passthrough")
	if err != nil {
		return nil, err
	}

	return &result, err
}

func (cli *client) TxPoolInfo(ctx context.Context) (*types.TxPoolInfo, error) {
	var result types.TxPoolInfo
	err := cli.c.CallContext(ctx, &result, "tx_pool_info")
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (cli *client) GetRawTxPool(ctx context.Context) (*types.RawTxPool, error) {
	var txPool types.RawTxPool
	err := cli.c.CallContext(ctx, &txPool, "get_raw_tx_pool")
	if err != nil {
		return nil, err
	}

	return &txPool, err
}

func (cli *client) ClearTxPool(ctx context.Context) error {
	return cli.c.CallContext(ctx, nil, "clear_tx_pool")
}

func (cli *client) GetBlockchainInfo(ctx context.Context) (*types.BlockchainInfo, error) {
	var result types.BlockchainInfo
	err := cli.c.CallContext(ctx, &result, "get_blockchain_info")
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (cli *client) BatchTransactions(ctx context.Context, batch []types.BatchTransactionItem) error {
	req := make([]rpc.BatchElem, len(batch))

	for i, item := range batch {
		args := make([]interface{}, 1)
		args[0] = item.Hash
		req[i] = rpc.BatchElem{
			Method: "get_transaction",
			Result: &types.TransactionWithStatus{},
			Args:   args,
		}
	}

	err := cli.c.BatchCallContext(ctx, req)
	if err != nil {
		return err
	}

	for i, item := range req {
		batch[i].Error = item.Error
		if batch[i].Error == nil {
			batch[i].Result = item.Result.(*types.TransactionWithStatus)
		}
	}

	return nil
}

func (cli *client) BatchLiveCells(ctx context.Context, batch []types.BatchLiveCellItem) error {
	req := make([]rpc.BatchElem, len(batch))

	for i, item := range batch {
		args := make([]interface{}, 2)
		args[0] = item.OutPoint
		args[1] = item.WithData
		req[i] = rpc.BatchElem{
			Method: "get_live_cell",
			Result: &types.CellWithStatus{},
			Args:   args,
		}
	}

	err := cli.c.BatchCallContext(ctx, req)
	if err != nil {
		return err
	}

	for i, item := range req {
		batch[i].Error = item.Error
		if batch[i].Error == nil {
			batch[i].Result = item.Result.(*types.CellWithStatus)
		}
	}
	return nil
}

func (cli *client) GetIndexerTip(ctx context.Context) (*indexer.TipHeader, error) {
	var result indexer.TipHeader
	err := cli.c.CallContext(ctx, &result, "get_indexer_tip")
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (cli *client) GetCellsCapacity(ctx context.Context, searchKey *indexer.SearchKey) (*indexer.Capacity, error) {
	var result indexer.Capacity
	err := cli.c.CallContext(ctx, &result, "get_cells_capacity", searchKey)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (cli *client) GetCells(ctx context.Context, searchKey *indexer.SearchKey, order indexer.SearchOrder, limit uint64, afterCursor string) (*indexer.LiveCells, error) {
	var (
		result indexer.LiveCells
		err    error
	)
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

func (cli *client) GetTransactions(ctx context.Context, searchKey *indexer.SearchKey, order indexer.SearchOrder, limit uint64, afterCursor string) (*indexer.TxsWithCell, error) {
	var (
		result indexer.TxsWithCell
		err    error
	)
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

func (cli *client) GetTransactionsGrouped(ctx context.Context, searchKey *indexer.SearchKey, order indexer.SearchOrder, limit uint64, afterCursor string) (*indexer.TxsWithCells, error) {
	payload := &struct {
		indexer.SearchKey
		GroupByTransaction bool `json:"group_by_transaction"`
	}{
		SearchKey:          *searchKey,
		GroupByTransaction: true,
	}
	var result indexer.TxsWithCells
	var err error
	if afterCursor == "" {
		err = cli.c.CallContext(ctx, &result, "get_transactions", payload, order, hexutil.Uint64(limit))
	} else {
		err = cli.c.CallContext(ctx, &result, "get_transactions", payload, order, hexutil.Uint64(limit), afterCursor)
	}
	if err != nil {
		return nil, err
	}
	return &result, err
}

func (cli *client) GetBlockEconomicState(ctx context.Context, blockHash types.Hash) (*types.BlockEconomicState, error) {
	var result types.BlockEconomicState
	err := cli.c.CallContext(ctx, &result, "get_block_economic_state", blockHash)
	if err != nil {
		return nil, err
	}

	// if FinalizedAt is equal to "0x0000000000000000000000000000000000000000000000000000000000000000" means block economic state is empty
	if result.FinalizedAt == types.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000") {
		return nil, nil
	}
	return &result, nil
}
