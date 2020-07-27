package rpc

import (
	"context"
	"encoding/json"
	"errors"
	"math/big"

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

	// GetCellsByLockHash returns the information about cells collection by the hash of lock script.
	GetCellsByLockHash(ctx context.Context, hash types.Hash, from uint64, to uint64) ([]*types.Cell, error)

	// GetLiveCell returns the information about a cell by out_point if it is live.
	// If second with_data argument set to true, will return cell data and data_hash if it is live.
	GetLiveCell(ctx context.Context, outPoint *types.OutPoint, withData bool) (*types.CellWithStatus, error)

	// GetTransaction returns the information about a transaction requested by transaction hash.
	GetTransaction(ctx context.Context, hash types.Hash) (*types.TransactionWithStatus, error)

	// GetCellbaseOutputCapacityDetails returns each component of the created CKB in this block's cellbase,
	// which is issued to a block N - 1 - ProposalWindow.farthest, where this block's height is N.
	GetCellbaseOutputCapacityDetails(ctx context.Context, hash types.Hash) (*types.BlockReward, error)

	// GetBlockByNumber get block by number
	GetBlockByNumber(ctx context.Context, number uint64) (*types.Block, error)

	////// Experiment
	// DryRunTransaction dry run transaction and return the execution cycles.
	// This method will not check the transaction validity,
	// but only run the lock script and type script and then return the execution cycles.
	// Used to debug transaction scripts and query how many cycles the scripts consume.
	DryRunTransaction(ctx context.Context, transaction *types.Transaction) (*types.DryRunTransactionResult, error)

	// CalculateDaoMaximumWithdraw calculate the maximum withdraw one can get, given a referenced DAO cell, and a withdraw block hash.
	CalculateDaoMaximumWithdraw(ctx context.Context, point *types.OutPoint, hash types.Hash) (uint64, error)

	// EstimateFeeRate Estimate a fee rate (capacity/KB) for a transaction that to be committed in expect blocks.
	EstimateFeeRate(ctx context.Context, blocks uint64) (*types.EstimateFeeRateResult, error)

	////// Indexer
	// IndexLockHash create index for live cells and transactions by the hash of lock script.
	IndexLockHash(ctx context.Context, lockHash types.Hash, indexFrom uint64) (*types.LockHashIndexState, error)

	// GetLockHashIndexStates Get lock hash index states.
	GetLockHashIndexStates(ctx context.Context) ([]*types.LockHashIndexState, error)

	// GetLiveCellsByLockHash returns the live cells collection by the hash of lock script.
	GetLiveCellsByLockHash(ctx context.Context, lockHash types.Hash, page uint, per uint, reverseOrder bool) ([]*types.LiveCell, error)

	// GetTransactionsByLockHash returns the transactions collection by the hash of lock script.
	// Returns empty array when the lock_hash has not been indexed yet.
	GetTransactionsByLockHash(ctx context.Context, lockHash types.Hash, page uint, per uint, reverseOrder bool) ([]*types.CellTransaction, error)

	// DeindexLockHash Remove index for live cells and transactions by the hash of lock script.
	DeindexLockHash(ctx context.Context, lockHash types.Hash) error

	////// Net
	// LocalNodeInfo returns the local node information.
	LocalNodeInfo(ctx context.Context) (*types.Node, error)

	// GetPeers returns the connected peers information.
	GetPeers(ctx context.Context) ([]*types.Node, error)

	// GetBannedAddresses returns all banned IPs/Subnets.
	GetBannedAddresses(ctx context.Context) ([]*types.BannedAddress, error)

	// SetBan insert or delete an IP/Subnet from the banned list
	SetBan(ctx context.Context, address string, command string, banTime uint64, absolute bool, reason string) error

	////// Pool
	// SendTransaction send new transaction into transaction pool.
	SendTransaction(ctx context.Context, tx *types.Transaction) (*types.Hash, error)

	// SendTransactionNoneValidation send new transaction into transaction pool skipping outputs validation.
	SendTransactionNoneValidation(ctx context.Context, tx *types.Transaction) (*types.Hash, error)

	// TxPoolInfo return the transaction pool information
	TxPoolInfo(ctx context.Context) (*types.TxPoolInfo, error)

	////// Stats
	// GetBlockchainInfo return state info of blockchain
	GetBlockchainInfo(ctx context.Context) (*types.BlockchainInfo, error)

	////// Batch
	BatchTransactions(ctx context.Context, batch []types.BatchTransactionItem) error

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
	var result header
	err := cli.c.CallContext(ctx, &result, "get_tip_header")
	if err != nil {
		return nil, err
	}
	return toHeader(result), err
}

func (cli *client) GetCurrentEpoch(ctx context.Context) (*types.Epoch, error) {
	var result epoch
	err := cli.c.CallContext(ctx, &result, "get_current_epoch")
	if err != nil {
		return nil, err
	}
	return &types.Epoch{
		CompactTarget: uint64(result.CompactTarget),
		Length:        uint64(result.Length),
		Number:        uint64(result.Number),
		StartNumber:   uint64(result.StartNumber),
	}, err
}

func (cli *client) GetEpochByNumber(ctx context.Context, number uint64) (*types.Epoch, error) {
	var result epoch
	err := cli.c.CallContext(ctx, &result, "get_epoch_by_number", hexutil.Uint64(number))
	if err != nil {
		return nil, err
	}
	return &types.Epoch{
		CompactTarget: uint64(result.CompactTarget),
		Length:        uint64(result.Length),
		Number:        uint64(result.Number),
		StartNumber:   uint64(result.StartNumber),
	}, err
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
	var raw json.RawMessage

	err := cli.c.CallContext(ctx, &raw, "get_block", hash)
	if err != nil {
		return nil, err
	} else if len(raw) == 0 {
		return nil, NotFound
	}

	var block block
	if err := json.Unmarshal(raw, &block); err != nil {
		return nil, err
	}

	return &types.Block{
		Header:       toHeader(block.Header),
		Proposals:    block.Proposals,
		Transactions: toTransactions(block.Transactions),
		Uncles:       toUncles(block.Uncles),
	}, nil
}

func (cli *client) GetHeader(ctx context.Context, hash types.Hash) (*types.Header, error) {
	var result header
	err := cli.c.CallContext(ctx, &result, "get_header", hash)
	if err != nil {
		return nil, err
	}
	return toHeader(result), err
}

func (cli *client) GetHeaderByNumber(ctx context.Context, number uint64) (*types.Header, error) {
	var result header
	err := cli.c.CallContext(ctx, &result, "get_header_by_number", hexutil.Uint64(number))
	if err != nil {
		return nil, err
	}
	return toHeader(result), err
}

func (cli *client) GetCellsByLockHash(ctx context.Context, hash types.Hash, from uint64, to uint64) ([]*types.Cell, error) {
	var result []cell
	err := cli.c.CallContext(ctx, &result, "get_cells_by_lock_hash", hash, hexutil.Uint64(from), hexutil.Uint64(to))
	if err != nil {
		return nil, err
	}
	return toCells(result), err
}

func (cli *client) GetLiveCell(ctx context.Context, point *types.OutPoint, withData bool) (*types.CellWithStatus, error) {
	var result cellWithStatus
	err := cli.c.CallContext(ctx, &result, "get_live_cell", outPoint{
		TxHash: point.TxHash,
		Index:  hexutil.Uint(point.Index),
	}, true)
	if err != nil {
		return nil, err
	}
	return toCellWithStatus(result), err
}

func (cli *client) GetTransaction(ctx context.Context, hash types.Hash) (*types.TransactionWithStatus, error) {
	var result transactionWithStatus
	err := cli.c.CallContext(ctx, &result, "get_transaction", hash)
	if err != nil {
		return nil, err
	}
	return &types.TransactionWithStatus{
		Transaction: toTransaction(result.Transaction),
		TxStatus: &types.TxStatus{
			BlockHash: result.TxStatus.BlockHash,
			Status:    result.TxStatus.Status,
		},
	}, err
}

func (cli *client) GetCellbaseOutputCapacityDetails(ctx context.Context, hash types.Hash) (*types.BlockReward, error) {
	var result blockReward
	err := cli.c.CallContext(ctx, &result, "get_cellbase_output_capacity_details", hash)
	if err != nil {
		return nil, err
	}

	return &types.BlockReward{
		Primary:        (*big.Int)(&result.Primary),
		ProposalReward: (*big.Int)(&result.ProposalReward),
		Secondary:      (*big.Int)(&result.Secondary),
		Total:          (*big.Int)(&result.Total),
		TxFee:          (*big.Int)(&result.TxFee),
	}, err
}

func (cli *client) GetBlockByNumber(ctx context.Context, number uint64) (*types.Block, error) {
	var raw json.RawMessage

	err := cli.c.CallContext(ctx, &raw, "get_block_by_number", hexutil.Uint64(number))
	if err != nil {
		return nil, err
	} else if len(raw) == 0 {
		return nil, NotFound
	}

	var block block
	if err := json.Unmarshal(raw, &block); err != nil {
		return nil, err
	}

	return &types.Block{
		Header:       toHeader(block.Header),
		Proposals:    block.Proposals,
		Transactions: toTransactions(block.Transactions),
		Uncles:       toUncles(block.Uncles),
	}, nil
}

func (cli *client) DryRunTransaction(ctx context.Context, transaction *types.Transaction) (*types.DryRunTransactionResult, error) {
	var result dryRunTransactionResult
	err := cli.c.CallContext(ctx, &result, "dry_run_transaction", fromTransaction(transaction))
	if err != nil {
		return nil, err
	}

	return &types.DryRunTransactionResult{
		Cycles: uint64(result.Cycles),
	}, err
}

func (cli *client) CalculateDaoMaximumWithdraw(ctx context.Context, point *types.OutPoint, hash types.Hash) (uint64, error) {
	var result hexutil.Uint64
	err := cli.c.CallContext(ctx, &result, "calculate_dao_maximum_withdraw", outPoint{TxHash: point.TxHash, Index: hexutil.Uint(point.Index)}, hash)
	if err != nil {
		return 0, err
	}

	return uint64(result), err
}

func (cli *client) EstimateFeeRate(ctx context.Context, blocks uint64) (*types.EstimateFeeRateResult, error) {
	var result estimateFeeRateResult

	err := cli.c.CallContext(ctx, &result, "estimate_fee_rate", hexutil.Uint64(blocks))
	if err != nil {
		return nil, err
	}

	return &types.EstimateFeeRateResult{
		FeeRate: uint64(result.FeeRate),
	}, err
}

func (cli *client) IndexLockHash(ctx context.Context, lockHash types.Hash, indexFrom uint64) (*types.LockHashIndexState, error) {
	var result lockHashIndexState

	err := cli.c.CallContext(ctx, &result, "index_lock_hash", lockHash, hexutil.Uint64(indexFrom))
	if err != nil {
		return nil, err
	}

	return &types.LockHashIndexState{
		BlockHash:   result.BlockHash,
		BlockNumber: uint64(result.BlockNumber),
		LockHash:    result.LockHash,
	}, err
}

func (cli *client) GetLockHashIndexStates(ctx context.Context) ([]*types.LockHashIndexState, error) {
	var result []lockHashIndexState

	err := cli.c.CallContext(ctx, &result, "get_lock_hash_index_states")
	if err != nil {
		return nil, err
	}

	ret := make([]*types.LockHashIndexState, len(result))
	for i := 0; i < len(result); i++ {
		state := result[i]
		ret[i] = &types.LockHashIndexState{
			BlockHash:   state.BlockHash,
			BlockNumber: uint64(state.BlockNumber),
			LockHash:    state.LockHash,
		}
	}

	return ret, err
}

func (cli *client) GetLiveCellsByLockHash(ctx context.Context, lockHash types.Hash, page uint, per uint, reverseOrder bool) ([]*types.LiveCell, error) {
	var result []liveCell

	err := cli.c.CallContext(ctx, &result, "get_live_cells_by_lock_hash", lockHash, hexutil.Uint(page), hexutil.Uint(per), reverseOrder)
	if err != nil {
		return nil, err
	}

	ret := make([]*types.LiveCell, len(result))

	for i := 0; i < len(result); i++ {
		cell := result[i]
		ret[i] = &types.LiveCell{
			CellOutput: &types.CellOutput{
				Capacity: uint64(cell.CellOutput.Capacity),
				Lock: &types.Script{
					CodeHash: cell.CellOutput.Lock.CodeHash,
					HashType: cell.CellOutput.Lock.HashType,
					Args:     cell.CellOutput.Lock.Args,
				},
			},
			CreatedBy: &types.TransactionPoint{
				BlockNumber: uint64(cell.CreatedBy.BlockNumber),
				Index:       uint(cell.CreatedBy.Index),
				TxHash:      cell.CreatedBy.TxHash,
			},
		}
		if cell.CellOutput.Type != nil {
			ret[i].CellOutput.Type = &types.Script{
				CodeHash: cell.CellOutput.Type.CodeHash,
				HashType: cell.CellOutput.Type.HashType,
				Args:     cell.CellOutput.Type.Args,
			}
		}
	}

	return ret, err
}

func (cli *client) GetTransactionsByLockHash(ctx context.Context, lockHash types.Hash, page uint, per uint, reverseOrder bool) ([]*types.CellTransaction, error) {
	var result []cellTransaction

	err := cli.c.CallContext(ctx, &result, "get_transactions_by_lock_hash", lockHash, hexutil.Uint(page), hexutil.Uint(per), reverseOrder)
	if err != nil {
		return nil, err
	}

	ret := make([]*types.CellTransaction, len(result))

	for i := 0; i < len(result); i++ {
		tx := result[i]
		ret[i] = &types.CellTransaction{
			CreatedBy: &types.TransactionPoint{
				BlockNumber: uint64(tx.CreatedBy.BlockNumber),
				Index:       uint(tx.CreatedBy.Index),
				TxHash:      tx.CreatedBy.TxHash,
			},
		}
		if tx.ConsumedBy != nil {
			ret[i].ConsumedBy = &types.TransactionPoint{
				BlockNumber: uint64(tx.ConsumedBy.BlockNumber),
				Index:       uint(tx.ConsumedBy.Index),
				TxHash:      tx.ConsumedBy.TxHash,
			}
		}
	}
	return ret, err
}

func (cli *client) DeindexLockHash(ctx context.Context, lockHash types.Hash) error {
	return cli.c.CallContext(ctx, nil, "deindex_lock_hash", lockHash)
}

func (cli *client) LocalNodeInfo(ctx context.Context) (*types.Node, error) {
	var result node

	err := cli.c.CallContext(ctx, &result, "local_node_info")
	if err != nil {
		return nil, err
	}

	return toNode(result), err
}

func (cli *client) GetPeers(ctx context.Context) ([]*types.Node, error) {
	var result []node

	err := cli.c.CallContext(ctx, &result, "get_peers")
	if err != nil {
		return nil, err
	}

	ret := make([]*types.Node, len(result))
	for i := 0; i < len(result); i++ {
		ret[i] = toNode(result[i])
	}

	return ret, err
}

func (cli *client) GetBannedAddresses(ctx context.Context) ([]*types.BannedAddress, error) {
	var result []bannedAddress

	err := cli.c.CallContext(ctx, &result, "get_banned_addresses")
	if err != nil {
		return nil, err
	}

	ret := make([]*types.BannedAddress, len(result))
	for i := 0; i < len(result); i++ {
		ret[i] = &types.BannedAddress{
			Address:   result[i].Address,
			BanReason: result[i].BanReason,
			BanUntil:  uint64(result[i].BanUntil),
			CreatedAt: uint64(result[i].CreatedAt),
		}
	}

	return ret, err
}

func (cli *client) SetBan(ctx context.Context, address string, command string, banTime uint64, absolute bool, reason string) error {
	return cli.c.CallContext(ctx, nil, "set_ban", address, command, hexutil.Uint64(banTime), absolute, reason)
}

func (cli *client) SendTransaction(ctx context.Context, tx *types.Transaction) (*types.Hash, error) {
	var result types.Hash

	err := cli.c.CallContext(ctx, &result, "send_transaction", fromTransaction(tx))
	if err != nil {
		return nil, err
	}

	return &result, err
}

func (cli *client) SendTransactionNoneValidation(ctx context.Context, tx *types.Transaction) (*types.Hash, error) {
	var result types.Hash

	err := cli.c.CallContext(ctx, &result, "send_transaction", fromTransaction(tx), "passthrough")
	if err != nil {
		return nil, err
	}

	return &result, err
}

func (cli *client) TxPoolInfo(ctx context.Context) (*types.TxPoolInfo, error) {
	var result txPoolInfo

	err := cli.c.CallContext(ctx, &result, "tx_pool_info")
	if err != nil {
		return nil, err
	}

	return &types.TxPoolInfo{
		LastTxsUpdatedAt: uint64(result.LastTxsUpdatedAt),
		Orphan:           uint64(result.Orphan),
		Pending:          uint64(result.Pending),
		Proposed:         uint64(result.Proposed),
		TotalTxCycles:    uint64(result.TotalTxCycles),
		TotalTxSize:      uint64(result.TotalTxSize),
	}, err
}

func (cli *client) GetBlockchainInfo(ctx context.Context) (*types.BlockchainInfo, error) {
	var result blockchainInfo

	err := cli.c.CallContext(ctx, &result, "get_blockchain_info")

	if err != nil {
		return nil, err
	}

	ret := &types.BlockchainInfo{
		Chain:                  result.Chain,
		Difficulty:             (*big.Int)(&result.Difficulty),
		Epoch:                  uint64(result.Epoch),
		IsInitialBlockDownload: result.IsInitialBlockDownload,
		MedianTime:             uint64(result.MedianTime),
	}

	ret.Alerts = make([]*types.AlertMessage, len(result.Alerts))
	for i := 0; i < len(result.Alerts); i++ {
		ret.Alerts[i] = &types.AlertMessage{
			Id:          result.Alerts[i].Id,
			Message:     result.Alerts[i].Message,
			NoticeUntil: uint64(result.Alerts[i].NoticeUntil),
			Priority:    result.Alerts[i].Priority,
		}
	}

	return ret, err
}

func (cli *client) BatchTransactions(ctx context.Context, batch []types.BatchTransactionItem) error {
	req := make([]rpc.BatchElem, len(batch))

	for i, item := range batch {
		args := make([]interface{}, 1)
		args[0] = item.Hash
		req[i] = rpc.BatchElem{
			Method: "get_transaction",
			Result: &transactionWithStatus{},
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
			batch[i].Result = &types.TransactionWithStatus{
				Transaction: toTransaction(item.Result.(*transactionWithStatus).Transaction),
				TxStatus: &types.TxStatus{
					BlockHash: item.Result.(*transactionWithStatus).TxStatus.BlockHash,
					Status:    item.Result.(*transactionWithStatus).TxStatus.Status,
				},
			}
		}
	}

	return nil
}
