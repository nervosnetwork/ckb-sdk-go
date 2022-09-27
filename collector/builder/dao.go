package builder

import (
	"bytes"
	"context"
	"encoding/binary"
	"errors"
	"github.com/nervosnetwork/ckb-sdk-go/v2/address"
	"github.com/nervosnetwork/ckb-sdk-go/v2/collector"
	"github.com/nervosnetwork/ckb-sdk-go/v2/collector/handler"
	"github.com/nervosnetwork/ckb-sdk-go/v2/rpc"
	"github.com/nervosnetwork/ckb-sdk-go/v2/types"
	"math/big"
)

type DaoTransactionType uint

const (
	DaoTransactionTypeWithdraw DaoTransactionType = iota
	DaoTransactionTypeClaim
)

type DaoTransactionBuilder struct {
	CkbTransactionBuilder
	client              rpc.Client
	transactionType     DaoTransactionType
	depositBlockNumber  uint64
	depositCellCapacity uint64
}

func NewDaoTransactionBuilder(network types.Network, iterator collector.CellIterator, daoOutPoint *types.OutPoint, client rpc.Client) (*DaoTransactionBuilder, error) {
	cellWithStatus, err := client.GetLiveCell(context.Background(), daoOutPoint, true)
	if err != nil {
		return nil, err
	}
	input := &types.TransactionInput{
		OutPoint:   daoOutPoint,
		Output:     cellWithStatus.Cell.Output,
		OutputData: cellWithStatus.Cell.Data.Content,
	}
	transactionType, err := getTransactionType(input.OutputData)
	if err != nil {
		return nil, err
	}

	depositBlockNumber := uint64(0)
	depositCellCapacity := uint64(0)
	reward := uint64(0)
	if transactionType == DaoTransactionTypeWithdraw {
		txWithStatus, err := client.GetTransaction(context.Background(), daoOutPoint.TxHash)
		if err != nil {
			return nil, err
		}
		header, err := client.GetHeader(context.Background(), txWithStatus.TxStatus.BlockHash)
		if err != nil {
			return nil, err
		}
		depositBlockNumber = header.Number
		depositCellCapacity = txWithStatus.Transaction.Outputs[daoOutPoint.Index].Capacity
	} else if transactionType == DaoTransactionTypeClaim {
		reward, err = getDaoReward(daoOutPoint, client)
		if err != nil {
			return nil, err
		}
	}

	b := &DaoTransactionBuilder{
		CkbTransactionBuilder: CkbTransactionBuilder{
			SimpleTransactionBuilder: *NewSimpleTransactionBuilder(network),
			FeeRate:                  1000,
			Network:                  network,
			iterator:                 iterator,
			transactionInputs:        []*types.TransactionInput{input}, // add dao inputs
			transactionInputsIndex:   0,
			changeOutputIndex:        -1,
			reward:                   reward,
		},
		client:              client,
		transactionType:     transactionType,
		depositBlockNumber:  depositBlockNumber,
		depositCellCapacity: depositCellCapacity,
	}
	return b, nil
}

func getTransactionType(outputData []byte) (DaoTransactionType, error) {
	if len(outputData) != 8 {
		return 0, errors.New("dao cell's output data length should be 8")
	}
	if bytes.Equal(outputData, handler.DaoDepositOutputData) {
		return DaoTransactionTypeWithdraw, nil
	} else {
		return DaoTransactionTypeClaim, nil
	}
}

func getDaoReward(withdrawOutPoint *types.OutPoint, client rpc.Client) (uint64, error) {
	txWithStatus, err := client.GetTransaction(context.Background(), withdrawOutPoint.TxHash)
	if err != nil {
		return 0, err
	}
	withdrawTx := txWithStatus.Transaction
	withdrawBlockHash := txWithStatus.TxStatus.BlockHash
	var (
		depositCell      *types.CellOutput = nil
		depositCellData  []byte
		depositBlockHash types.Hash
	)
	for i := 0; i < len(withdrawTx.Inputs); i++ {
		outPoint := withdrawTx.Inputs[i].PreviousOutput
		txWithStatus, err := client.GetTransaction(context.Background(), outPoint.TxHash)
		if err != nil {
			return 0, err
		}
		tx := txWithStatus.Transaction
		output := tx.Outputs[outPoint.Index]
		data := tx.OutputsData[outPoint.Index]
		if handler.IsDepositCell(output, data) {
			depositCell = output
			depositCellData = data
			depositBlockHash = txWithStatus.TxStatus.BlockHash
			break
		}
	}
	if depositCell == nil {
		return 0, errors.New("can't find deposit cell")
	}
	depositBlockHeader, err := client.GetHeader(context.Background(), depositBlockHash)
	if err != nil {
		return 0, err
	}
	withdrawBlockHeader, err := client.GetHeader(context.Background(), withdrawBlockHash)
	if err != nil {
		return 0, err
	}
	occupiedCapacity := depositCell.OccupiedCapacity(depositCellData)
	daoMaximumWithdraw := calculateDaoMaximumWithdraw(depositBlockHeader, withdrawBlockHeader, depositCell, occupiedCapacity)
	daoReward := daoMaximumWithdraw - depositCell.Capacity
	return daoReward, nil
}

func calculateDaoMaximumWithdraw(depositBlockHeader, withdrawBlockHeader *types.Header, output *types.CellOutput, occupiedCapacity uint64) uint64 {
	depositAr := extractAr(depositBlockHeader.Dao)
	withdrawAr := extractAr(withdrawBlockHeader.Dao)

	maximumWithdraw := new(big.Int).SetUint64(output.Capacity)
	maximumWithdraw.Sub(maximumWithdraw, new(big.Int).SetUint64(occupiedCapacity))
	maximumWithdraw.Mul(maximumWithdraw, new(big.Int).SetUint64(withdrawAr))
	maximumWithdraw.Div(maximumWithdraw, new(big.Int).SetUint64(depositAr))
	maximumWithdraw.Add(maximumWithdraw, new(big.Int).SetUint64(occupiedCapacity))

	return maximumWithdraw.Uint64()
}

func extractAr(dao types.Hash) uint64 {
	return binary.LittleEndian.Uint64(dao[8:16])
}

func (r *DaoTransactionBuilder) AddWithdrawOutput(addr string) error {
	if r.depositBlockNumber == 0 {
		return errors.New("deposit block number not initialized")
	}
	a, err := address.Decode(addr)
	if err != nil {
		return err
	}
	output := &types.CellOutput{
		Capacity: r.depositCellCapacity,
		Lock:     a.Script,
		Type:     handler.DaoScript,
	}
	data := types.SerializeUint64(r.depositBlockNumber)
	r.AddOutput(output, data)
	return nil
}
