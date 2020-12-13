package builder

import (
	"bytes"
	"github.com/ethereum/go-ethereum/common"
	"github.com/nervosnetwork/ckb-sdk-go/collector"
	"github.com/nervosnetwork/ckb-sdk-go/transaction"
	"github.com/nervosnetwork/ckb-sdk-go/types"
	"github.com/nervosnetwork/ckb-sdk-go/utils"
	"github.com/pkg/errors"
	"math"
	"math/big"
)

var _ UnsignedTxBuilder = (*IssuingChequeUnsignedTxBuilder)(nil)
var (
	chequeCellCapacity  = uint64(162 * math.Pow10(8))
	udtCellCapacity     = uint64(142 * math.Pow10(8))
	sudtDataPlaceHolder = make([]byte, 16)
)

type IssuingChequeUnsignedTxBuilder struct {
	Sender         *types.Script
	Receiver       *types.Script
	FeeRate        uint64
	CkbIterator    collector.CellCollectionIterator
	SUDTIterator   collector.CellCollectionIterator
	SystemScripts  *utils.SystemScripts
	TransferAmount *big.Int
	UUID           string

	tx                    *types.Transaction
	result                *collector.LiveCellCollectResult
	ckbChangeOutputIndex  *collector.ChangeOutputIndex
	sUDTChangeOutputIndex *collector.ChangeOutputIndex
}

func (b *IssuingChequeUnsignedTxBuilder) NewTransaction() {
	b.tx = &types.Transaction{}
}

func (b *IssuingChequeUnsignedTxBuilder) BuildVersion() {
	b.tx.Version = 0
}

func (b *IssuingChequeUnsignedTxBuilder) BuildHeaderDeps() {
	b.tx.HeaderDeps = []types.Hash{}
}

func (b *IssuingChequeUnsignedTxBuilder) BuildCellDeps() {
	b.tx.CellDeps = []*types.CellDep{
		{
			OutPoint: b.SystemScripts.SecpSingleSigCell.OutPoint,
			DepType:  types.DepTypeDepGroup,
		},
		{
			OutPoint: b.SystemScripts.SUDTCell.OutPoint,
			DepType:  b.SystemScripts.SUDTCell.DepType,
		},
	}
}

func (b *IssuingChequeUnsignedTxBuilder) BuildOutputsAndOutputsData() error {
	udtType := &types.Script{
		CodeHash: b.SystemScripts.SUDTCell.CellHash,
		HashType: b.SystemScripts.SUDTCell.HashType,
		Args:     common.FromHex(b.UUID),
	}
	chequeCellArgs, err := utils.ChequeCellArgs(b.Sender, b.Receiver)
	if err != nil {
		return err
	}
	// set cheque output
	b.tx.Outputs = append(b.tx.Outputs, &types.CellOutput{
		Capacity: chequeCellCapacity,
		Lock: &types.Script{
			CodeHash: b.SystemScripts.ChequeCell.CellHash,
			HashType: b.SystemScripts.ChequeCell.HashType,
			Args:     chequeCellArgs,
		},
		Type: udtType,
	})
	b.tx.OutputsData = append(b.tx.OutputsData, utils.GenerateSudtAmount(b.TransferAmount))

	// set ckb change output
	b.tx.Outputs = append(b.tx.Outputs, &types.CellOutput{
		Capacity: 0,
		Lock:     b.Sender,
	})
	b.tx.OutputsData = append(b.tx.OutputsData, []byte{})
	// set ckb change output index
	b.ckbChangeOutputIndex = &collector.ChangeOutputIndex{Value: 1}

	// set sudt change output
	b.tx.Outputs = append(b.tx.Outputs, &types.CellOutput{
		Capacity: udtCellCapacity,
		Lock:     b.Sender,
		Type:     udtType,
	})
	b.tx.OutputsData = append(b.tx.OutputsData, sudtDataPlaceHolder)
	// set ckb change output index
	b.sUDTChangeOutputIndex = &collector.ChangeOutputIndex{Value: 2}

	return nil
}

func (b *IssuingChequeUnsignedTxBuilder) BuildInputsAndWitnesses() error {
	if b.TransferAmount == nil {
		return errors.New("transfer amount is required")
	}
	// collect sudt cells first
	err := b.collectSUDTCells()

	if err != nil {
		return err
	}

	// then collect ckb cells
	err = b.collectCkbCells()
	if err != nil {
		return err
	}
	return nil
}

func (b *IssuingChequeUnsignedTxBuilder) collectCkbCells() error {
	for b.CkbIterator.HasNext() {
		liveCell, err := b.CkbIterator.CurrentItem()
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
		ok, err := b.isCkbEnough()
		if err != nil {
			return err
		}
		if ok {
			return nil
		}
		err = b.CkbIterator.Next()
		if err != nil {
			return err
		}
	}
	return errors.New("insufficient ckb balance")
}

func (b *IssuingChequeUnsignedTxBuilder) collectSUDTCells() error {
	b.result = &collector.LiveCellCollectResult{}
	for b.SUDTIterator.HasNext() {
		liveCell, err := b.SUDTIterator.CurrentItem()
		if err != nil {
			return err
		}
		b.result.Capacity += liveCell.Output.Capacity
		b.result.LiveCells = append(b.result.LiveCells, liveCell)
		// init totalAmount
		if _, ok := b.result.Options["totalAmount"]; !ok {
			b.result.Options = make(map[string]interface{})
			b.result.Options["totalAmount"] = big.NewInt(0)
		}
		amount, err := utils.ParseSudtAmount(liveCell.OutputData)
		if err != nil {
			return errors.WithMessage(err, "sudt amount parse error")
		}
		totalAmount := b.result.Options["totalAmount"].(*big.Int)
		b.result.Options["totalAmount"] = big.NewInt(0).Add(totalAmount, amount)
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
		// stop collect
		if b.isSUDTEnough() {
			return nil
		}
		err = b.SUDTIterator.Next()
		if err != nil {
			return err
		}
	}
	return errors.New("insufficient sudt balance")
}

func (b *IssuingChequeUnsignedTxBuilder) UpdateChangeOutput() error {
	// update sudt change output first
	totalAmount := b.result.Options["totalAmount"].(*big.Int)
	if totalAmount.Cmp(b.TransferAmount) > 0 && bytes.Compare(b.tx.OutputsData[b.sUDTChangeOutputIndex.Value], sudtDataPlaceHolder) == 0 {
		b.tx.OutputsData[b.sUDTChangeOutputIndex.Value] = utils.GenerateSudtAmount(big.NewInt(0).Sub(totalAmount, b.TransferAmount))
	}
	if totalAmount.Cmp(b.TransferAmount) == 0 {
		b.tx.Outputs = utils.RemoveCellOutput(b.tx.Outputs, b.sUDTChangeOutputIndex.Value)
		b.tx.OutputsData = utils.RemoveCellOutputData(b.tx.OutputsData, b.sUDTChangeOutputIndex.Value)
	}

	// then update ckb change output
	fee, err := transaction.CalculateTransactionFee(b.tx, b.FeeRate)
	if err != nil {
		return err
	}
	changeCapacity := b.result.Capacity - b.tx.OutputsCapacity() - fee
	b.tx.Outputs[b.ckbChangeOutputIndex.Value].Capacity = changeCapacity

	return nil
}

func (b *IssuingChequeUnsignedTxBuilder) GetResult() *types.Transaction {
	return b.tx
}

func (b *IssuingChequeUnsignedTxBuilder) isSUDTEnough() bool {
	totalAmount := b.result.Options["totalAmount"].(*big.Int)
	if totalAmount.Cmp(b.TransferAmount) >= 0 {
		return true
	}
	return false
}

func (b *IssuingChequeUnsignedTxBuilder) isCkbEnough() (bool, error) {
	changeCapacity := b.result.Capacity - b.tx.OutputsCapacity()
	if changeCapacity > 0 {
		fee, err := transaction.CalculateTransactionFee(b.tx, b.FeeRate)
		if err != nil {
			return false, err
		}
		changeCapacity -= fee
		changeOutput := b.tx.Outputs[b.ckbChangeOutputIndex.Value]
		changeOutputData := b.tx.OutputsData[b.ckbChangeOutputIndex.Value]
		changeOutputCapacity := changeOutput.OccupiedCapacity(changeOutputData)
		if changeCapacity >= changeOutputCapacity {
			return true, nil
		} else {
			return false, nil
		}
	} else {
		return false, nil
	}
}
