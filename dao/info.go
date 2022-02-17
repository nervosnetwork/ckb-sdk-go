package dao

import (
	"context"
	"encoding/binary"
	"github.com/nervosnetwork/ckb-sdk-go/rpc"
	"github.com/nervosnetwork/ckb-sdk-go/types"
	"github.com/pkg/errors"
	"math/big"
)

type DaoHelper struct {
	Client rpc.Client
}

type DaoDepositCellInfo struct {
	Outpoint                 types.OutPoint
	withdrawBlockHash        types.Hash
	withdrawBlockNumber      uint64
	DepositCapacity          uint64
	Compensation             uint64
	NextClaimableBlock       uint64
	NextClaimableEpochNumber uint32
}

// GetDaoDepositCellInfo Get information for dao cell deposited as outpoint and withdrawn in block of withdrawBlockHash
func (c *DaoHelper) GetDaoDepositCellInfo(outpoint *types.OutPoint, withdrawBlockHash types.Hash) (DaoDepositCellInfo, error) {
	withdrawBlock, err := c.Client.GetBlock(context.Background(), withdrawBlockHash)
	if err != nil {
		return DaoDepositCellInfo{}, err
	}
	return c.getDaoDepositCellInfo(outpoint, withdrawBlock)
}

// GetDaoDepositCellInfoWithWithdrawOutpoint Get information for dao cell deposited as outpoint and withdrawn in block where the withdrawCellOutPoint is
func (c *DaoHelper) GetDaoDepositCellInfoWithWithdrawOutpoint(outpoint *types.OutPoint, withdrawCellOutPoint *types.OutPoint) (DaoDepositCellInfo, error) {
	withdrawTransaction, err := c.Client.GetTransaction(context.Background(), withdrawCellOutPoint.TxHash)
	if err != nil {
		return DaoDepositCellInfo{}, err
	}
	return c.GetDaoDepositCellInfo(outpoint, *withdrawTransaction.TxStatus.BlockHash)
}

// GetDaoDepositCellInfoByNow DAO cell's information for dao cell deposited as outpoint and withdrawn in tip block
func (c *DaoHelper) GetDaoDepositCellInfoByNow(outpoint *types.OutPoint) (DaoDepositCellInfo, error) {
	tipBlockNumber, err := c.Client.GetTipBlockNumber(context.Background())
	if err != nil {
		return DaoDepositCellInfo{}, err
	}

	tipBlock, err := c.Client.GetBlockByNumber(context.Background(), tipBlockNumber)
	if err != nil {
		return DaoDepositCellInfo{}, err
	}

	return c.getDaoDepositCellInfo(outpoint, tipBlock)
}

// GetDaoDepositCellInfo Get DAO cell's information for dao cell deposited as outpoint and withdrawn in withdrawBlock
func (c *DaoHelper) getDaoDepositCellInfo(outpoint *types.OutPoint, withdrawBlock *types.Block) (DaoDepositCellInfo, error) {
	cellInfo := DaoDepositCellInfo{}
	cellInfo.Outpoint = *outpoint
	cellInfo.withdrawBlockHash = withdrawBlock.Header.Hash
	cellInfo.withdrawBlockNumber = withdrawBlock.Header.Number

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
	depositAr := new(big.Int).SetUint64(extractArFromDaoData(depositBlock.Header.Dao))
	withdrawAr := new(big.Int).SetUint64(extractArFromDaoData(withdrawBlock.Header.Dao))
	compensation := new(big.Int)
	compensation.Mul(freeCapacity, withdrawAr).Div(compensation, depositAr).Sub(compensation, freeCapacity)
	cellInfo.Compensation = compensation.Uint64()

	cellInfo.DepositCapacity = totalCapacity

	epochLength, blockIndexInEpoch, epochNumber := ResolveEpoch(withdrawBlock.Header.Epoch)
	cellInfo.NextClaimableBlock = withdrawBlock.Header.Number + uint64(epochLength-blockIndexInEpoch)
	cellInfo.NextClaimableEpochNumber = epochNumber + 1

	return cellInfo, nil
}

func extractArFromDaoData(headerDao types.Hash) uint64 {
	ar := headerDao[8:16]
	return binary.LittleEndian.Uint64(ar)
}

// ResolveEpoch return (epochLength, blockIndexInEpoch, epochNumber)
func ResolveEpoch(epoch uint64) (uint16, uint16, uint32) {
	epochBinary := make([]byte, 8)
	binary.BigEndian.PutUint64(epochBinary, epoch)

	epochLength := binary.BigEndian.Uint16(epochBinary[1:3])
	blockIndexInEpoch := binary.BigEndian.Uint16(epochBinary[3:5])
	epochNumberBinary := append(make([]byte, 1), epochBinary[5:8]...)
	epochNumber := binary.BigEndian.Uint32(epochNumberBinary)

	return epochLength, blockIndexInEpoch, epochNumber
}
