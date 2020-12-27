package builder

import (
	"bytes"
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/nervosnetwork/ckb-sdk-go/collector"
	"github.com/nervosnetwork/ckb-sdk-go/indexer"
	"github.com/nervosnetwork/ckb-sdk-go/rpc"
	"github.com/nervosnetwork/ckb-sdk-go/transaction"
	"github.com/nervosnetwork/ckb-sdk-go/types"
	"github.com/nervosnetwork/ckb-sdk-go/utils"
	"github.com/pkg/errors"
	"math"
	"math/big"
)

var _ UnsignedTxBuilder = (*ClaimChequesUnsignedTxBuilder)(nil)

type refundInfo struct {
	lock     *types.Script
	capacity uint64
}

type ClaimChequesUnsignedTxBuilder struct {
	Receiver       *types.Script
	FeeRate        uint64
	CkbIterator    collector.CellCollectionIterator
	ChequeIterator collector.CellCollectionIterator
	SystemScripts  *utils.SystemScripts
	UUID           string
	Client         rpc.Client

	tx                    *types.Transaction
	result                *collector.LiveCellCollectResult
	ckbChangeOutputIndex  *collector.ChangeOutputIndex
	sUDTChangeOutputIndex *collector.ChangeOutputIndex
	groups                [][]int
}

func (b *ClaimChequesUnsignedTxBuilder) NewTransaction() {
	b.tx = &types.Transaction{}
}

func (b *ClaimChequesUnsignedTxBuilder) BuildVersion() {
	b.tx.Version = 0
}

func (b *ClaimChequesUnsignedTxBuilder) BuildHeaderDeps() {
	b.tx.HeaderDeps = []types.Hash{}
}

func (b *ClaimChequesUnsignedTxBuilder) BuildCellDeps() {
	b.tx.CellDeps = []*types.CellDep{
		{
			OutPoint: b.SystemScripts.SecpSingleSigCell.OutPoint,
			DepType:  types.DepTypeDepGroup,
		},
		{
			OutPoint: b.SystemScripts.SUDTCell.OutPoint,
			DepType:  b.SystemScripts.SUDTCell.DepType,
		},
		{
			OutPoint: b.SystemScripts.ChequeCell.OutPoint,
			DepType:  b.SystemScripts.ChequeCell.DepType,
		},
	}
}

func (b *ClaimChequesUnsignedTxBuilder) BuildOutputsAndOutputsData() error {
	udtType := &types.Script{
		CodeHash: b.SystemScripts.SUDTCell.CellHash,
		HashType: b.SystemScripts.SUDTCell.HashType,
		Args:     common.FromHex(b.UUID),
	}
	// set ckb change output
	b.tx.Outputs = append(b.tx.Outputs, &types.CellOutput{
		Capacity: 0,
		Lock:     b.Receiver,
	})
	b.tx.OutputsData = append(b.tx.OutputsData, []byte{})
	// set ckb change output index
	b.ckbChangeOutputIndex = &collector.ChangeOutputIndex{Value: 0}

	// set sudt claim output
	b.tx.Outputs = append(b.tx.Outputs, &types.CellOutput{
		Capacity: udtCellCapacity,
		Lock:     b.Receiver,
		Type:     udtType,
	})
	b.tx.OutputsData = append(b.tx.OutputsData, sudtDataPlaceHolder)
	// set ckb change output index
	b.sUDTChangeOutputIndex = &collector.ChangeOutputIndex{Value: 1}

	return nil
}

func (b *ClaimChequesUnsignedTxBuilder) BuildInputsAndWitnesses() error {
	// collect cheque cells first
	err := b.collectChequeCells()
	if err != nil {
		return err
	}
	// generate refund outputs
	rf := b.result.Options["refundInfo"].(map[string]*refundInfo)
	for _, info := range rf {
		b.tx.Outputs = append(b.tx.Outputs, &types.CellOutput{
			Capacity: info.capacity,
			Lock:     info.lock,
		})
		b.tx.OutputsData = append(b.tx.OutputsData, []byte{})
	}

	// then collect ckb cells
	err = b.collectCkbCells()
	if err != nil {
		return err
	}
	return nil
}

func (b *ClaimChequesUnsignedTxBuilder) collectCkbCells() error {
	lastChequeWitnessIndex := len(b.tx.Witnesses)
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
		if len(b.tx.Witnesses[lastChequeWitnessIndex]) == 0 {
			b.tx.Witnesses[lastChequeWitnessIndex] = transaction.EmptyWitnessArgPlaceholder
		}
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

func (b *ClaimChequesUnsignedTxBuilder) collectChequeCells() error {
	b.result = &collector.LiveCellCollectResult{}
	if !b.ChequeIterator.HasNext() {
		return errors.New("no cheque cells to claim")
	}
	for b.ChequeIterator.HasNext() {
		liveCell, err := b.ChequeIterator.CurrentItem()
		if err != nil {
			return err
		}
		b.result.Capacity += liveCell.Output.Capacity
		b.result.LiveCells = append(b.result.LiveCells, liveCell)
		// init totalAmount and refundInfo
		if _, ok := b.result.Options["totalAmount"]; !ok {
			b.result.Options = make(map[string]interface{})
			b.result.Options["totalAmount"] = big.NewInt(0)
			b.result.Options["refundInfo"] = make(map[string]*refundInfo)
		}
		// update sudt total Amount
		err = b.updateTotalAmount(err, liveCell)
		if err != nil {
			return err
		}
		// update input lock script capacity info
		err = b.updateRefundCapacityInfo(liveCell)
		if err != nil {
			return err
		}
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
		err = b.ChequeIterator.Next()
		if err != nil {
			return err
		}
	}
	return nil
}

func (b *ClaimChequesUnsignedTxBuilder) updateRefundCapacityInfo(liveCell *indexer.LiveCell) error {
	chequeOutput := liveCell.Output
	lockHash, err := chequeOutput.Lock.Hash()
	if err != nil {
		return err
	}
	senderLock, err := b.findSenderLock(liveCell, chequeOutput)
	if err != nil {
		return err
	}

	rf := b.result.Options["refundInfo"].(map[string]*refundInfo)
	if v, ok := rf[lockHash.String()]; ok {
		v.capacity += chequeOutput.Capacity
	} else {
		rf[lockHash.String()] = &refundInfo{
			lock:     senderLock,
			capacity: chequeOutput.Capacity,
		}
	}
	b.result.Options["refundInfo"] = rf
	return nil
}

func (b *ClaimChequesUnsignedTxBuilder) findSenderLock(liveCell *indexer.LiveCell, chequeOutput *types.CellOutput) (*types.Script, error) {
	previousTx, err := b.Client.GetTransaction(context.Background(), liveCell.OutPoint.TxHash)
	if err != nil {
		return nil, err
	}
	for _, input := range previousTx.Transaction.Inputs {
		txWithStatus, err := b.Client.GetTransaction(context.Background(), input.PreviousOutput.TxHash)
		if err != nil {
			return nil, err
		}
		lock := txWithStatus.Transaction.Outputs[input.PreviousOutput.Index].Lock
		lockHash, err := lock.Hash()
		if err != nil {
			return nil, err
		}
		if bytes.Compare(lockHash.Bytes()[0:20], chequeOutput.Lock.Args[20:]) == 0 {
			return lock, nil
		}
	}
	return nil, errors.New("sender lock not found")
}

func (b *ClaimChequesUnsignedTxBuilder) updateTotalAmount(err error, liveCell *indexer.LiveCell) error {
	amount, err := utils.ParseSudtAmount(liveCell.OutputData)
	if err != nil {
		return errors.WithMessage(err, "sudt amount parse error")
	}
	totalAmount := b.result.Options["totalAmount"].(*big.Int)
	b.result.Options["totalAmount"] = big.NewInt(0).Add(totalAmount, amount)
	return nil
}

func (b *ClaimChequesUnsignedTxBuilder) UpdateChangeOutput() error {
	// update sudt claim output
	totalAmount := b.result.Options["totalAmount"].(*big.Int)
	b.tx.OutputsData[b.sUDTChangeOutputIndex.Value] = utils.GenerateSudtAmount(totalAmount)

	// then update ckb change output
	fee, err := transaction.CalculateTransactionFee(b.tx, b.FeeRate)
	if err != nil {
		return err
	}
	changeCapacity := b.result.Capacity - b.tx.OutputsCapacity() - fee
	b.tx.Outputs[b.ckbChangeOutputIndex.Value].Capacity = changeCapacity
	err = b.generateGroups()
	if err != nil {
		return err
	}
	return nil
}

func (b *ClaimChequesUnsignedTxBuilder) GetResult() (*types.Transaction, [][]int) {
	return b.tx, b.groups
}

func (b *ClaimChequesUnsignedTxBuilder) isCkbEnough() (bool, error) {
	inputsCapacity := big.NewInt(0).SetUint64(b.result.Capacity)
	outputsCapacity := big.NewInt(0).SetUint64(b.tx.OutputsCapacity())
	changeCapacity := big.NewInt(0).Sub(inputsCapacity, outputsCapacity)

	if changeCapacity.Cmp(big.NewInt(0)) > 0 {
		fee, err := transaction.CalculateTransactionFee(b.tx, b.FeeRate)
		if err != nil {
			return false, err
		}
		changeCapacity = big.NewInt(0).Sub(changeCapacity, big.NewInt(0).SetUint64(fee))
		changeOutput := b.tx.Outputs[b.ckbChangeOutputIndex.Value]
		changeOutputData := b.tx.OutputsData[b.ckbChangeOutputIndex.Value]
		changeOutputCapacity := big.NewInt(0).SetUint64(changeOutput.OccupiedCapacity(changeOutputData) * uint64(math.Pow10(8)))
		if changeCapacity.Cmp(changeOutputCapacity) >= 0 {
			return true, nil
		} else {
			return false, nil
		}
	} else {
		return false, nil
	}
}

func (b *ClaimChequesUnsignedTxBuilder) generateGroups() error {
	groupInfo := make(map[string][]int)
	for i, liveCell := range b.result.LiveCells {
		lockHash, err := liveCell.Output.Lock.Hash()
		if err != nil {
			return err
		}
		key := lockHash.String()
		if v, ok := groupInfo[key]; ok {
			v = append(v, i)
			groupInfo[key] = v
		} else {
			groupInfo[key] = []int{i}
		}
	}
	var groups [][]int
	for _, group := range groupInfo {
		groups = append(groups, group)
	}
	b.groups = groups
	return nil
}
