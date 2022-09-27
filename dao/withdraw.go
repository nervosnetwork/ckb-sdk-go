package dao

import (
	"context"
	"errors"
	"fmt"

	"github.com/nervosnetwork/ckb-sdk-go/v2/rpc"
	"github.com/nervosnetwork/ckb-sdk-go/v2/types"
	"github.com/nervosnetwork/ckb-sdk-go/v2/utils"
)

type WithdrawPhase1 struct {
	Transaction *types.Transaction
}

func NewWithdrawPhase1(scripts *utils.SystemScripts, isMultisig bool) *WithdrawPhase1 {
	var baseDep *types.CellDep
	if isMultisig {
		baseDep = &types.CellDep{
			OutPoint: scripts.SecpMultiSigCell.OutPoint,
			DepType:  types.DepTypeDepGroup,
		}
	} else {
		baseDep = &types.CellDep{
			OutPoint: scripts.SecpSingleSigCell.OutPoint,
			DepType:  types.DepTypeDepGroup,
		}
	}

	tx := &types.Transaction{
		Version:    0,
		HeaderDeps: []types.Hash{},
		CellDeps: []*types.CellDep{
			baseDep,
			{
				OutPoint: scripts.DaoCell.OutPoint,
				DepType:  types.DepTypeCode,
			},
		},
	}

	return &WithdrawPhase1{
		Transaction: tx,
	}
}

func (w *WithdrawPhase1) AddDaoDepositTick(client rpc.Client, cell *types.Cell) (int, error) {
	header, err := client.GetHeader(context.Background(), cell.BlockHash)
	if err != nil {
		return 0, fmt.Errorf("get block header from address %s error: %v", cell.BlockHash.String(), err)
	}

	w.Transaction.HeaderDeps = append(w.Transaction.HeaderDeps, cell.BlockHash)
	w.Transaction.Inputs = append(w.Transaction.Inputs, &types.CellInput{
		Since: 0,
		PreviousOutput: &types.OutPoint{
			TxHash: cell.OutPoint.TxHash,
			Index:  cell.OutPoint.Index,
		},
	})
	w.Transaction.Witnesses = append(w.Transaction.Witnesses, []byte{})
	w.Transaction.Outputs = append(w.Transaction.Outputs, &types.CellOutput{
		Capacity: cell.Capacity,
		Lock:     cell.Lock,
		Type:     cell.Type,
	})
	w.Transaction.OutputsData = append(w.Transaction.OutputsData, types.SerializeUint64(header.Number))
	return len(w.Transaction.Inputs) - 1, nil
}

func (w *WithdrawPhase1) AddOutput(lock *types.Script, amount uint64) error {
	if w.Transaction == nil {
		return errors.New("must init transaction first")
	}
	w.Transaction.Outputs = append(w.Transaction.Outputs, &types.CellOutput{
		Capacity: amount,
		Lock:     lock,
	})
	w.Transaction.OutputsData = append(w.Transaction.OutputsData, []byte{})

	return nil
}

type WithdrawPhase2 struct {
	Transaction *types.Transaction
}

func NewWithdrawPhase2(scripts *utils.SystemScripts, isMultisig bool) *WithdrawPhase2 {
	var baseDep *types.CellDep
	if isMultisig {
		baseDep = &types.CellDep{
			OutPoint: scripts.SecpMultiSigCell.OutPoint,
			DepType:  types.DepTypeDepGroup,
		}
	} else {
		baseDep = &types.CellDep{
			OutPoint: scripts.SecpSingleSigCell.OutPoint,
			DepType:  types.DepTypeDepGroup,
		}
	}

	tx := &types.Transaction{
		Version:    0,
		HeaderDeps: []types.Hash{},
		CellDeps: []*types.CellDep{
			baseDep,
			{
				OutPoint: scripts.DaoCell.OutPoint,
				DepType:  types.DepTypeCode,
			},
		},
	}

	return &WithdrawPhase2{
		Transaction: tx,
	}
}

func (w *WithdrawPhase2) AddDaoWithdrawTick(client rpc.Client, depositCell *types.Cell, withdrawCell *types.Cell, fee uint64) (int, *types.WitnessArgs, error) {
	headerDeposit, err := client.GetHeader(context.Background(), depositCell.BlockHash)
	if err != nil {
		return 0, nil, fmt.Errorf("get block header from address %s error: %v", depositCell.BlockHash.String(), err)
	}

	headerWithdraw, err := client.GetHeader(context.Background(), withdrawCell.BlockHash)
	if err != nil {
		return 0, nil, fmt.Errorf("get block header from address %s error: %v", withdrawCell.BlockHash.String(), err)
	}

	withdrawEpoch := types.ParseEpoch(headerWithdraw.Epoch)
	depositEpoch := types.ParseEpoch(headerDeposit.Epoch)

	withdrawFraction := withdrawEpoch.Index * depositEpoch.Length
	depositFraction := depositEpoch.Index * withdrawEpoch.Length
	depositedEpochs := withdrawEpoch.Number - depositEpoch.Number
	if withdrawFraction > depositFraction {
		depositedEpochs += 1
	}
	lockEpochs := (depositedEpochs + 179) / 180 * 180
	minimalSinceEpochNumber := depositEpoch.Number + lockEpochs
	minimalSinceEpochIndex := depositEpoch.Index
	minimalSinceEpochLength := depositEpoch.Length

	minimalSince := &types.EpochParams{
		Length: minimalSinceEpochLength,
		Index:  minimalSinceEpochIndex,
		Number: minimalSinceEpochNumber,
	}
	capacity, err := client.CalculateDaoMaximumWithdraw(context.Background(), &types.OutPoint{
		TxHash: depositCell.OutPoint.TxHash,
		Index:  depositCell.OutPoint.Index,
	}, headerWithdraw.Hash)
	if err != nil {
		return 0, nil, fmt.Errorf("calculate Dao maximum withdraw error: %v", err)
	}

	if fee > capacity {
		return 0, nil, fmt.Errorf("the fee(%d) is too big that withdraw(%d) is not enough", fee, capacity)
	}

	w.Transaction.HeaderDeps = append(w.Transaction.HeaderDeps, depositCell.BlockHash)
	w.Transaction.HeaderDeps = append(w.Transaction.HeaderDeps, withdrawCell.BlockHash)

	w.Transaction.Inputs = append(w.Transaction.Inputs, &types.CellInput{
		Since: minimalSince.Uint64(),
		PreviousOutput: &types.OutPoint{
			TxHash: withdrawCell.OutPoint.TxHash,
			Index:  withdrawCell.OutPoint.Index,
		},
	})
	w.Transaction.Witnesses = append(w.Transaction.Witnesses, []byte{})
	w.Transaction.Outputs = append(w.Transaction.Outputs, &types.CellOutput{
		Capacity: capacity - fee,
		Lock:     withdrawCell.Lock,
	})
	w.Transaction.OutputsData = append(w.Transaction.OutputsData, []byte{})

	return len(w.Transaction.Inputs) - 1, &types.WitnessArgs{
		Lock:       make([]byte, 65),
		InputType:  types.SerializeUint64(uint64(len(w.Transaction.HeaderDeps) - 2)),
		OutputType: nil,
	}, nil
}

func (w *WithdrawPhase2) AddOutput(lock *types.Script, amount uint64) error {
	if w.Transaction == nil {
		return errors.New("must init transaction first")
	}
	w.Transaction.Outputs = append(w.Transaction.Outputs, &types.CellOutput{
		Capacity: amount,
		Lock:     lock,
	})
	w.Transaction.OutputsData = append(w.Transaction.OutputsData, []byte{})

	return nil
}
