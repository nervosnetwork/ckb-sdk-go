package builder

import (
	"github.com/nervosnetwork/ckb-sdk-go/collector"
	"github.com/nervosnetwork/ckb-sdk-go/transaction"
	"github.com/nervosnetwork/ckb-sdk-go/types"
	"github.com/nervosnetwork/ckb-sdk-go/utils"
	"github.com/pkg/errors"
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

func (c *CkbTransferUnsignedTxBuilder) NewTransaction() {
	c.tx = new(types.Transaction)
}

func (c *CkbTransferUnsignedTxBuilder) BuildVersion() {
	c.tx.Version = 0
}

func (c *CkbTransferUnsignedTxBuilder) BuildHeaderDeps() {
	c.tx.HeaderDeps = []types.Hash{}
}

func (c *CkbTransferUnsignedTxBuilder) BuildCellDeps() {
	c.tx.CellDeps = []*types.CellDep{
		{
			OutPoint: c.SystemScripts.SecpSingleSigCell.OutPoint,
			DepType:  types.DepTypeDepGroup,
		},
	}
}

func (c *CkbTransferUnsignedTxBuilder) BuildOutputsAndOutputsData() error {
	// set transfer output
	c.tx.Outputs = append(c.tx.Outputs, &types.CellOutput{
		Capacity: c.TransferCapacity,
		Lock:     c.To,
	})
	c.tx.OutputsData = [][]byte{{}}
	// set change output
	if !c.TransferAll {
		c.tx.Outputs = append(c.tx.Outputs, &types.CellOutput{
			Capacity: 0,
			Lock:     c.From,
		})
		c.tx.OutputsData = append(c.tx.OutputsData, []byte{})
		// set change output index
		c.ckbChangeOutputIndex = &collector.ChangeOutputIndex{Value: 1}
	}
	return nil
}

func (c *CkbTransferUnsignedTxBuilder) BuildInputsAndWitnesses() error {
	for ; c.Iterator.HasNext(); c.Iterator.Next() {
		liveCell, err := c.Iterator.CurrentItem()
		if err != nil {
			return err
		}
		c.result.LiveCells = append(c.result.LiveCells, liveCell)
		c.result.Capacity = liveCell.Output.Capacity
		input := &types.CellInput{
			Since: 0,
			PreviousOutput: &types.OutPoint{
				TxHash: liveCell.OutPoint.TxHash,
				Index:  liveCell.OutPoint.Index,
			},
		}
		c.tx.Inputs = append(c.tx.Inputs, input)
		c.tx.Witnesses = append(c.tx.Witnesses, []byte{})
		if len(c.tx.Witnesses[0]) == 0 {
			c.tx.Witnesses[0] = transaction.EmptyWitnessArgPlaceholder
		}
		ok, err := c.isEnough()
		if err != nil {
			return err
		}
		if ok {
			return nil
		}
	}

	return errors.New("insufficient ckb balance")
}

func (c *CkbTransferUnsignedTxBuilder) UpdateChangeOutput() error {
	if !c.TransferAll {
		fee, err := transaction.CalculateTransactionFee(c.tx, c.FeeRate)
		if err != nil {
			return err
		}
		changeCapacity := c.result.Capacity - c.tx.OutputsCapacity() - fee
		c.tx.Outputs[c.ckbChangeOutputIndex.Value].Capacity = changeCapacity
	}
	return nil
}

func (c *CkbTransferUnsignedTxBuilder) GetResult() *types.Transaction {
	return c.tx
}

func (c *CkbTransferUnsignedTxBuilder) isEnough() (bool, error) {
	changeCapacity := c.result.Capacity - c.tx.OutputsCapacity()
	if changeCapacity > 0 {
		fee, err := transaction.CalculateTransactionFee(c.tx, c.FeeRate)
		if err != nil {
			return false, err
		}
		changeCapacity -= fee
		if !c.TransferAll {
			changeOutput := c.tx.Outputs[c.ckbChangeOutputIndex.Value]
			changeOutputData := c.tx.OutputsData[c.ckbChangeOutputIndex.Value]
			changeOutputCapacity := changeOutput.OccupiedCapacity(changeOutputData)
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
