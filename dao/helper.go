package dao

import (
	"context"
	"encoding/binary"
	"github.com/nervosnetwork/ckb-sdk-go/rpc"
	"github.com/nervosnetwork/ckb-sdk-go/types"
	"github.com/pkg/errors"
	"math"
	"math/big"
)

type DaoHelper struct {
	Client rpc.Client
}

type DaoDepositCellInfo struct {
	Outpoint            types.OutPoint
	withdrawBlockHash   types.Hash
	withdrawBlockNumber uint64
	DepositCapacity     uint64
	Compensation        uint64
	EpochParams         types.EpochParams
}

// GetDaoDepositCellInfo Get information for DAO cell deposited as outpoint and withdrawn in block of withdrawBlockHash
func (c *DaoHelper) GetDaoDepositCellInfo(outpoint *types.OutPoint, withdrawBlockHash *types.Hash) (DaoDepositCellInfo, error) {
	withdrawBlock, err := c.Client.GetBlock(context.Background(), *withdrawBlockHash)
	if err != nil {
		return DaoDepositCellInfo{}, err
	}
	return c.getDaoDepositCellInfo(outpoint, withdrawBlock.Header)
}

// GetDaoDepositCellInfoWithWithdrawOutpoint Get information for DAO cell deposited as outpoint and withdrawn in block where the withdrawCellOutPoint is
func (c *DaoHelper) GetDaoDepositCellInfoWithWithdrawOutpoint(outpoint *types.OutPoint, withdrawCellOutPoint *types.OutPoint) (DaoDepositCellInfo, error) {
	withdrawTransaction, err := c.Client.GetTransaction(context.Background(), withdrawCellOutPoint.TxHash)
	if err != nil {
		return DaoDepositCellInfo{}, err
	}
	return c.GetDaoDepositCellInfo(outpoint, withdrawTransaction.TxStatus.BlockHash)
}

// GetDaoDepositCellInfoByNow DAO information for DAO cell deposited as outpoint and withdrawn in tip block
func (c *DaoHelper) GetDaoDepositCellInfoByNow(outpoint *types.OutPoint) (DaoDepositCellInfo, error) {
	tipBlockNumber, err := c.Client.GetTipBlockNumber(context.Background())
	if err != nil {
		return DaoDepositCellInfo{}, err
	}
	tipBlock, err := c.Client.GetBlockByNumber(context.Background(), tipBlockNumber)
	if err != nil {
		return DaoDepositCellInfo{}, err
	}

	return c.getDaoDepositCellInfo(outpoint, tipBlock.Header)
}

// getDaoDepositCellInfo Get information for DAO cell deposited as outpoint and withdrawn in withdrawBlock
func (c *DaoHelper) getDaoDepositCellInfo(outpoint *types.OutPoint, withdrawBlockHeader *types.Header) (DaoDepositCellInfo, error) {
	cellInfo := DaoDepositCellInfo{}
	cellInfo.Outpoint = *outpoint
	cellInfo.withdrawBlockHash = withdrawBlockHeader.Hash
	cellInfo.withdrawBlockNumber = withdrawBlockHeader.Number

	depositTransactionWithStatus, err := c.Client.GetTransaction(context.Background(), outpoint.TxHash)
	if err != nil {
		return cellInfo, err
	}
	depositBlock, err := c.Client.GetBlock(context.Background(), *depositTransactionWithStatus.TxStatus.BlockHash)
	if err != nil {
		return cellInfo, err
	}

	outpointData := depositTransactionWithStatus.Transaction.OutputsData[outpoint.Index]
	outpointCell := depositTransactionWithStatus.Transaction.Outputs[outpoint.Index]
	occupiedCapacity := outpointCell.OccupiedCapacity(outpointData) * 100000000
	totalCapacity := outpointCell.Capacity
	if totalCapacity < occupiedCapacity {
		return cellInfo, errors.New("Total capacity is less than occupied capacity")
	}

	freeCapacity := new(big.Int).SetUint64(totalCapacity - occupiedCapacity)
	depositAr := new(big.Int).SetUint64(extractArFromDaoData(&depositBlock.Header.Dao))
	withdrawAr := new(big.Int).SetUint64(extractArFromDaoData(&withdrawBlockHeader.Dao))
	compensation := new(big.Int)
	compensation.Mul(freeCapacity, withdrawAr).Div(compensation, depositAr).Sub(compensation, freeCapacity)
	cellInfo.Compensation = compensation.Uint64()
	cellInfo.DepositCapacity = totalCapacity

	//withdrawEpochLength, withdrawBlockIndexInEpoch, withdrawEpochNumber := ResolveEpoch(withdrawBlockHeader.Epoch)
	//depositEpochLength, depositBlockIndexInEpoch, depositEpochNumber := ResolveEpoch(depositBlock.Header.Epoch)

	withdrawEpochParams := types.ParseEpoch(withdrawBlockHeader.Epoch)
	depositEpochParams := types.ParseEpoch(depositBlock.Header.Epoch)

	withdrawEpoch := float64(withdrawEpochParams.Number) + float64(withdrawEpochParams.Index)/float64(withdrawEpochParams.Length)
	depositEpoch := float64(depositEpochParams.Number) + float64(depositEpochParams.Index)/float64(depositEpochParams.Length)
	epochDistance := uint64(math.Ceil((withdrawEpoch-depositEpoch)/180) * 180)

	cellInfo.EpochParams = types.EpochParams{
		Length: depositEpochParams.Length,
		Index:  depositEpochParams.Index,
		Number: depositEpochParams.Number + epochDistance,
	}

	return cellInfo, nil
}

func extractArFromDaoData(headerDao *types.Hash) uint64 {
	ar := headerDao[8:16]
	return binary.LittleEndian.Uint64(ar)
}
