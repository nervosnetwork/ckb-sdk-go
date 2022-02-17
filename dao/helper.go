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
	ClaimableEpoch      struct {
		Numerator, Denominator uint64
	}
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

	withdrawEpochLength, withdrawBlockIndexInEpoch, withdrawEpochNumber := ResolveEpoch(withdrawBlockHeader.Epoch)
	depositEpochLength, depositBlockIndexInEpoch, depositEpochNumber := ResolveEpoch(depositBlock.Header.Epoch)
	withdrawEpoch := float64(withdrawEpochNumber) + float64(withdrawBlockIndexInEpoch)/float64(withdrawEpochLength)
	depositEpoch := float64(depositEpochNumber) + float64(depositBlockIndexInEpoch)/float64(depositEpochLength)

	// claimableEpoch = depositEpoch + round_up( (withdrawEpoch - depositEpoch) / 180) * 180
	//                = (depositEpochNumber + depositBlockIndexInEpoch / depositEpochLength) + epochDistance
	//                = ( (depositEpochNumber + epochDistance) * depositEpochLength + depositBlockIndexInEpoch ) / depositEpochLength
	epochDistance := uint64(math.Ceil((withdrawEpoch-depositEpoch)/180) * 180)
	cellInfo.ClaimableEpoch = struct{ Numerator, Denominator uint64 }{
		Numerator:   (uint64(depositEpochNumber)+epochDistance)*uint64(depositEpochLength) + uint64(depositBlockIndexInEpoch+1), // block index in epoch starts with 0 but proportion should start with 1
		Denominator: uint64(depositEpochLength),
	}

	return cellInfo, nil
}

func extractArFromDaoData(headerDao *types.Hash) uint64 {
	ar := headerDao[8:16]
	return binary.LittleEndian.Uint64(ar)
}

// ResolveEpoch resolve epoch to (epochLength, blockIndexInEpoch, epochNumber)
func ResolveEpoch(epoch uint64) (uint16, uint16, uint32) {
	epochBinary := make([]byte, 8)
	binary.BigEndian.PutUint64(epochBinary, epoch)

	epochLength := binary.BigEndian.Uint16(epochBinary[1:3])
	blockIndexInEpoch := binary.BigEndian.Uint16(epochBinary[3:5])
	epochNumberBinary := append(make([]byte, 1), epochBinary[5:8]...)
	epochNumber := binary.BigEndian.Uint32(epochNumberBinary)

	return epochLength, blockIndexInEpoch, epochNumber
}
