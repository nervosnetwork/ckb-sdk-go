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

var _ UnsignedTxBuilder = (*SudtTransferUnsignedTxBuilder)(nil)

type SudtTransferUnsignedTxBuilder struct {
	CkbChanger     *types.Script
	SudtChanger    *types.Script
	Senders        []*types.Script
	ReceiverInfo   []types.ReceiverInfo
	CkbIterator    collector.CellCollectionIterator
	SUDTIterators  []collector.CellCollectionIterator
	SystemScripts  *utils.SystemScripts
	TransferAmount *big.Int
	UUID           string
	FeeRate        uint64

	tx                    *types.Transaction
	result                *collector.LiveCellCollectResult
	ckbChangeOutputIndex  *collector.ChangeOutputIndex
	sUDTChangeOutputIndex *collector.ChangeOutputIndex
	groups                [][]int
}

func (s *SudtTransferUnsignedTxBuilder) NewTransaction() {
	s.tx = &types.Transaction{}
}

func (s *SudtTransferUnsignedTxBuilder) BuildVersion() {
	s.tx.Version = 0
}

func (s *SudtTransferUnsignedTxBuilder) BuildHeaderDeps() {
	s.tx.HeaderDeps = []types.Hash{}
}

func (s *SudtTransferUnsignedTxBuilder) BuildCellDeps() {
	s.tx.CellDeps = []*types.CellDep{
		{
			OutPoint: s.SystemScripts.SecpSingleSigCell.OutPoint,
			DepType:  types.DepTypeDepGroup,
		},
		{
			OutPoint: s.SystemScripts.SUDTCell.OutPoint,
			DepType:  s.SystemScripts.SUDTCell.DepType,
		},
	}
}

func (s *SudtTransferUnsignedTxBuilder) BuildOutputsAndOutputsData() error {
	udtType := &types.Script{
		CodeHash: s.SystemScripts.SUDTCell.CellHash,
		HashType: s.SystemScripts.SUDTCell.HashType,
		Args:     common.FromHex(s.UUID),
	}
	// set receivers sudt output
	for _, r := range s.ReceiverInfo {
		s.tx.Outputs = append(s.tx.Outputs, &types.CellOutput{
			Capacity: udtCellCapacity,
			Lock: &types.Script{
				CodeHash: r.Receiver.CodeHash,
				HashType: r.Receiver.HashType,
				Args:     r.Receiver.Args,
			},
			Type: udtType,
		})
		s.tx.OutputsData = append(s.tx.OutputsData, utils.GenerateSudtAmount(r.Amount))
	}

	// set ckb change output
	s.tx.Outputs = append(s.tx.Outputs, &types.CellOutput{
		Capacity: 0,
		Lock:     s.CkbChanger,
	})
	s.tx.OutputsData = append(s.tx.OutputsData, []byte{})
	// set ckb change output index
	s.ckbChangeOutputIndex = &collector.ChangeOutputIndex{Value: len(s.tx.Outputs) - 1}

	// set sudt change output
	s.tx.Outputs = append(s.tx.Outputs, &types.CellOutput{
		Capacity: udtCellCapacity,
		Lock:     s.SudtChanger,
		Type:     udtType,
	})
	s.tx.OutputsData = append(s.tx.OutputsData, sudtDataPlaceHolder)
	// set sudt change output index
	s.sUDTChangeOutputIndex = &collector.ChangeOutputIndex{Value: len(s.tx.Outputs) - 1}

	return nil
}

func (s *SudtTransferUnsignedTxBuilder) BuildInputsAndWitnesses() error {
	if s.TransferAmount == nil {
		return errors.New("transfer amount is required")
	}
	// collect sudt cells first
	err := s.collectSUDTCells()
	if err != nil {
		return err
	}

	// then collect ckb cells
	err = s.collectCkbCells()
	if err != nil {
		return err
	}
	return nil
}

func (s *SudtTransferUnsignedTxBuilder) UpdateChangeOutput() error {
	// update sudt change output first
	totalAmount := s.result.Options["totalAmount"].(*big.Int)
	if totalAmount.Cmp(s.TransferAmount) > 0 && bytes.Compare(s.tx.OutputsData[s.sUDTChangeOutputIndex.Value], sudtDataPlaceHolder) == 0 {
		s.tx.OutputsData[s.sUDTChangeOutputIndex.Value] = utils.GenerateSudtAmount(big.NewInt(0).Sub(totalAmount, s.TransferAmount))
	}
	if totalAmount.Cmp(s.TransferAmount) == 0 {
		s.tx.Outputs = utils.RemoveCellOutput(s.tx.Outputs, s.sUDTChangeOutputIndex.Value)
		s.tx.OutputsData = utils.RemoveCellOutputData(s.tx.OutputsData, s.sUDTChangeOutputIndex.Value)
	}

	// then update ckb change output
	fee, err := transaction.CalculateTransactionFee(s.tx, s.FeeRate)
	if err != nil {
		return err
	}
	changeCapacity := s.result.Capacity - s.tx.OutputsCapacity() - fee
	s.tx.Outputs[s.ckbChangeOutputIndex.Value].Capacity = changeCapacity
	err = s.generateGroups()
	if err != nil {
		return err
	}

	return nil
}

func (s *SudtTransferUnsignedTxBuilder) GetResult() (*types.Transaction, [][]int) {
	return s.tx, s.groups
}

func (s *SudtTransferUnsignedTxBuilder) collectCkbCells() error {
	for s.CkbIterator.HasNext() {
		liveCell, err := s.CkbIterator.CurrentItem()
		if err != nil {
			return err
		}
		s.result.Capacity += liveCell.Output.Capacity
		s.result.LiveCells = append(s.result.LiveCells, liveCell)
		input := &types.CellInput{
			Since: 0,
			PreviousOutput: &types.OutPoint{
				TxHash: liveCell.OutPoint.TxHash,
				Index:  liveCell.OutPoint.Index,
			},
		}
		s.tx.Inputs = append(s.tx.Inputs, input)
		s.tx.Witnesses = append(s.tx.Witnesses, []byte{})
		ok, err := s.isCkbEnough()
		if err != nil {
			return err
		}
		if ok {
			return nil
		}
		err = s.CkbIterator.Next()
		if err != nil {
			return err
		}
	}
	return errors.New("insufficient ckb balance")
}

func (s *SudtTransferUnsignedTxBuilder) collectSUDTCells() error {
	s.result = &collector.LiveCellCollectResult{}
	for _, iterator := range s.SUDTIterators {
		for iterator.HasNext() {
			liveCell, err := iterator.CurrentItem()
			if err != nil {
				return err
			}
			s.result.Capacity += liveCell.Output.Capacity
			s.result.LiveCells = append(s.result.LiveCells, liveCell)
			// init totalAmount
			if _, ok := s.result.Options["totalAmount"]; !ok {
				s.result.Options = make(map[string]interface{})
				s.result.Options["totalAmount"] = big.NewInt(0)
			}
			amount, err := utils.ParseSudtAmount(liveCell.OutputData)
			if err != nil {
				return errors.WithMessage(err, "sudt amount parse error")
			}
			totalAmount := s.result.Options["totalAmount"].(*big.Int)
			s.result.Options["totalAmount"] = big.NewInt(0).Add(totalAmount, amount)
			input := &types.CellInput{
				Since: 0,
				PreviousOutput: &types.OutPoint{
					TxHash: liveCell.OutPoint.TxHash,
					Index:  liveCell.OutPoint.Index,
				},
			}
			s.tx.Inputs = append(s.tx.Inputs, input)
			s.tx.Witnesses = append(s.tx.Witnesses, []byte{})
			if len(s.tx.Witnesses[0]) == 0 {
				s.tx.Witnesses[0] = transaction.EmptyWitnessArgPlaceholder
			}
			// stop collect
			if s.isSUDTEnough() {
				return nil
			}
			err = iterator.Next()
			if err != nil {
				return err
			}
		}
	}
	return errors.New("insufficient sudt balance")
}

func (s *SudtTransferUnsignedTxBuilder) isSUDTEnough() bool {
	totalAmount := s.result.Options["totalAmount"].(*big.Int)
	if totalAmount.Cmp(s.TransferAmount) >= 0 {
		return true
	}
	return false
}

func (s *SudtTransferUnsignedTxBuilder) isCkbEnough() (bool, error) {
	inputsCapacity := big.NewInt(0).SetUint64(s.result.Capacity)
	outputsCapacity := big.NewInt(0).SetUint64(s.tx.OutputsCapacity())
	changeCapacity := big.NewInt(0).Sub(inputsCapacity, outputsCapacity)
	if changeCapacity.Cmp(big.NewInt(0)) > 0 {
		fee, err := transaction.CalculateTransactionFee(s.tx, s.FeeRate)
		if err != nil {
			return false, err
		}
		changeCapacity = big.NewInt(0).Sub(changeCapacity, big.NewInt(0).SetUint64(fee))
		changeOutput := s.tx.Outputs[s.ckbChangeOutputIndex.Value]
		changeOutputData := s.tx.OutputsData[s.ckbChangeOutputIndex.Value]

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

func (s *SudtTransferUnsignedTxBuilder) generateGroups() error {
	groupInfo := make(map[string][]int)
	for _, sender := range s.Senders {
		senderLockHash, err := sender.Hash()
		if err != nil {
			return err
		}
		groupInfo[senderLockHash.String()] = []int{}
	}
	for i, liveCell := range s.result.LiveCells {
		lockHash, err := liveCell.Output.Lock.Hash()
		if err != nil {
			return err
		}
		key := lockHash.String()
		if v, ok := groupInfo[key]; ok {
			v = append(v, i)
			groupInfo[key] = v
		}
	}
	var groups [][]int
	for _, group := range groupInfo {
		groups = append(groups, group)
	}
	s.groups = groups
	return nil
}
