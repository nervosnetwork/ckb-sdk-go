package dao

import (
	"context"
	"encoding/binary"
	"errors"
	"github.com/nervosnetwork/ckb-sdk-go/rpc"
	"github.com/nervosnetwork/ckb-sdk-go/types"
	"math/big"
)

type DaoHelper struct {
	Client rpc.Client
}

type DaoDepositCellInfo struct {
	Outpoint            types.OutPoint
	WithdrawBlockHash   types.Hash
	WithdrawBlockNumber uint64
	DepositCapacity     uint64
	Compensation        uint64
	UnlockableEpoch     types.EpochParams
}

// GetDaoDepositCellInfo Get information for DAO cell deposited as outpoint and withdrawn in block of withdrawBlockHash
func (c *DaoHelper) GetDaoDepositCellInfo(outpoint *types.OutPoint, withdrawBlockHash *types.Hash) (DaoDepositCellInfo, error) {
	blockHeader, err := c.Client.GetHeader(context.Background(), *withdrawBlockHash)
	if err != nil {
		return DaoDepositCellInfo{}, err
	}
	return c.getDaoDepositCellInfo(outpoint, blockHeader)
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
	tipBlockHeader, err := c.Client.GetTipHeader(context.Background())
	if err != nil {
		return DaoDepositCellInfo{}, err
	}
	return c.getDaoDepositCellInfo(outpoint, tipBlockHeader)
}

// getDaoDepositCellInfo Get information for DAO cell deposited as outpoint and withdrawn in withdrawBlock
func (c *DaoHelper) getDaoDepositCellInfo(outpoint *types.OutPoint, withdrawBlockHeader *types.Header) (DaoDepositCellInfo, error) {
	depositTransactionWithStatus, err := c.Client.GetTransaction(context.Background(), outpoint.TxHash)
	if err != nil {
		return DaoDepositCellInfo{}, err
	}
	depositBlockHeader, err := c.Client.GetHeader(context.Background(), *depositTransactionWithStatus.TxStatus.BlockHash)
	if err != nil {
		return DaoDepositCellInfo{}, err
	}

	if int(outpoint.Index) >= len(depositTransactionWithStatus.Transaction.Outputs) {
		return DaoDepositCellInfo{}, errors.New("index out of range of outputs in deposit transaction")
	}
	outpointCell := depositTransactionWithStatus.Transaction.Outputs[outpoint.Index]
	outpointData := depositTransactionWithStatus.Transaction.OutputsData[outpoint.Index]
	occupiedCapacity := outpointCell.OccupiedCapacity(outpointData) * 100000000
	totalCapacity := outpointCell.Capacity
	freeCapacity := new(big.Int).SetUint64(totalCapacity - occupiedCapacity)
	depositAr := new(big.Int).SetUint64(extractArFromDaoData(&depositBlockHeader.Dao))
	withdrawAr := new(big.Int).SetUint64(extractArFromDaoData(&withdrawBlockHeader.Dao))

	compensation := new(big.Int)
	compensation.Mul(freeCapacity, withdrawAr).Div(compensation, depositAr).Sub(compensation, freeCapacity)
	cellInfo := DaoDepositCellInfo{}
	cellInfo.Compensation = compensation.Uint64()
	cellInfo.DepositCapacity = totalCapacity

	withdrawEpochParams := types.ParseEpoch(withdrawBlockHeader.Epoch)
	depositEpochParams := types.ParseEpoch(depositBlockHeader.Epoch)
	// epochDistance = Ceil( (withdrawEpoch - depositEpoch ) / 180 ) * 180
	epochDistance := withdrawEpochParams.Number - depositEpochParams.Number
	if withdrawEpochParams.Index*depositEpochParams.Length > depositEpochParams.Index*withdrawEpochParams.Length {
		epochDistance += 1
	}
	epochDistance = (epochDistance + 179) / 180 * 180

	cellInfo.UnlockableEpoch = types.EpochParams{
		Length: depositEpochParams.Length,
		Index:  depositEpochParams.Index,
		Number: depositEpochParams.Number + epochDistance,
	}

	cellInfo.Outpoint = *outpoint
	cellInfo.WithdrawBlockHash = withdrawBlockHeader.Hash
	cellInfo.WithdrawBlockNumber = withdrawBlockHeader.Number

	return cellInfo, nil
}

func extractArFromDaoData(headerDao *types.Hash) uint64 {
	ar := headerDao[8:16]
	return binary.LittleEndian.Uint64(ar)
}
