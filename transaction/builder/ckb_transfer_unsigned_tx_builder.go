package builder

import (
	"github.com/nervosnetwork/ckb-sdk-go/collector"
	"github.com/nervosnetwork/ckb-sdk-go/transaction"
	"github.com/nervosnetwork/ckb-sdk-go/types"
	"github.com/nervosnetwork/ckb-sdk-go/utils"
	"github.com/pkg/errors"
	"math"
)

var _ UnsignedTxBuilder = (*CkbTransferUnsignedTxBuilder)(nil)

type CkbTransferUnsignedTxBuilder struct {
	To               *types.Script
	From             *types.Script
	FeeRate          uint64
	Iterator         collector.CellCollectionIterator
	TransferAll      bool
	SystemScripts    *utils.SystemScripts
	TransferCapacity uint64

	tx                   *types.Transaction
	result               collector.LiveCellCollectResult
	ckbChangeOutputIndex *collector.ChangeOutputIndex
}

func (b *CkbTransferUnsignedTxBuilder) NewTransaction() {
	b.tx = new(types.Transaction)
}

func (b *CkbTransferUnsignedTxBuilder) BuildVersion() {
	b.tx.Version = 0
}

func (b *CkbTransferUnsignedTxBuilder) BuildHeaderDeps() {
	b.tx.HeaderDeps = []types.Hash{}
}

func (b *CkbTransferUnsignedTxBuilder) BuildCellDeps() {
	b.tx.CellDeps = []*types.CellDep{
		{
			OutPoint: b.SystemScripts.SecpSingleSigCell.OutPoint,
			DepType:  types.DepTypeDepGroup,
		},
	}
}

func (b *CkbTransferUnsignedTxBuilder) BuildOutputsAndOutputsData() error {
	// set transfer output
	b.tx.Outputs = append(b.tx.Outputs, &types.CellOutput{
		Capacity: b.TransferCapacity,
		Lock:     b.To,
	})
	b.tx.OutputsData = [][]byte{{}}
	// set change output
	if !b.TransferAll {
		b.tx.Outputs = append(b.tx.Outputs, &types.CellOutput{
			Capacity: 0,
			Lock:     b.From,
		})
		b.tx.OutputsData = append(b.tx.OutputsData, []byte{})
		// set change output index
		b.ckbChangeOutputIndex = &collector.ChangeOutputIndex{Value: 1}
	}
	return nil
}

func (b *CkbTransferUnsignedTxBuilder) BuildInputsAndWitnesses() error {
	for b.Iterator.HasNext() {
		liveCell, err := b.Iterator.CurrentItem()
		if err != nil {
			return err
		}
		b.result.Capacity += liveCell.Output.Capacity
		b.result.LiveCells = append(b.result.LiveCells, liveCell)
		input := &types.CellInput{
			Since: 0,
			PreviousOutput: &types.OutPoint{
				TxHash: liveCell.OutPoint.TxHash,
				Index:  liveCell.OutPoint.Index,
			},
		}
		b.tx.Inputs = append(b.tx.Inputs, input)
		b.tx.Witnesses = append(b.tx.Witnesses, []byte{})
		if len(b.tx.Witnesses[0]) == 0 {
			b.tx.Witnesses[0] = transaction.EmptyWitnessArgPlaceholder
		}
		ok, err := b.isEnough()
		if err != nil {
			return err
		}
		if ok {
			return nil
		}
		err = b.Iterator.Next()
		if err != nil {
			return err
		}
	}

	return errors.New("insufficient ckb balance")
}

func (b *CkbTransferUnsignedTxBuilder) UpdateChangeOutput() error {
	if !b.TransferAll {
		fee, err := transaction.CalculateTransactionFee(b.tx, b.FeeRate)
		if err != nil {
			return err
		}
		changeCapacity := b.result.Capacity - b.tx.OutputsCapacity() - fee
		b.tx.Outputs[b.ckbChangeOutputIndex.Value].Capacity = changeCapacity
	}
	return nil
}

func (b *CkbTransferUnsignedTxBuilder) GetResult() (*types.Transaction, [][]int) {
	return b.tx, nil
}

func (b *CkbTransferUnsignedTxBuilder) isEnough() (bool, error) {
	changeCapacity := b.result.Capacity - b.tx.OutputsCapacity()
	if changeCapacity > 0 {
		fee, err := transaction.CalculateTransactionFee(b.tx, b.FeeRate)
		if err != nil {
			return false, err
		}
		changeCapacity -= fee
		if !b.TransferAll {
			changeOutput := b.tx.Outputs[b.ckbChangeOutputIndex.Value]
			changeOutputData := b.tx.OutputsData[b.ckbChangeOutputIndex.Value]
			changeOutputCapacity := changeOutput.OccupiedCapacity(changeOutputData) * uint64(math.Pow10(8))
			if changeCapacity >= changeOutputCapacity {
				return true, nil
			} else {
				return false, nil
			}
		} else {
			// check whether the handling fee is sufficient
			if changeCapacity > 0 {
				return true, nil
			} else {
				return false, nil
			}
		}
	} else {
		return false, nil
	}
}
