package builder

import (
	"errors"
	address2 "github.com/nervosnetwork/ckb-sdk-go/address"
	"github.com/nervosnetwork/ckb-sdk-go/collector"
	"github.com/nervosnetwork/ckb-sdk-go/systemscript"
	"github.com/nervosnetwork/ckb-sdk-go/transaction"
	"github.com/nervosnetwork/ckb-sdk-go/types"
	"math/big"
	"reflect"
)

type SudtTransactionType uint

const (
	SudtTransactionTypeIssue SudtTransactionType = iota
	SudtTransactionTypeTransfer
)

type SudtTransactionBuilder struct {
	SimpleTransactionBuilder
	FeeRate  uint
	SudtType *types.Script

	iterator          collector.CellIterator
	changeOutputIndex int
	transactionType   SudtTransactionType
}

func NewSudtTransactionBuilderFromSudtArgs(network types.Network, iterator collector.CellIterator,
	transactionType SudtTransactionType, sudtArgs []byte) *SudtTransactionBuilder {
	codeHash := systemscript.GetCodeHash(network, systemscript.Sudt)
	s := &SudtTransactionBuilder{
		SimpleTransactionBuilder: *NewSimpleTransactionBuilder(network),
		FeeRate:                  1000,
		SudtType: &types.Script{
			CodeHash: codeHash,
			HashType: types.HashTypeType,
			Args:     sudtArgs,
		},

		iterator:          iterator,
		changeOutputIndex: -1,
		transactionType:   transactionType,
	}
	return s
}

func NewSudtTransactionBuilderFromSudtOwnerAddress(network types.Network, iterator collector.CellIterator,
	transactionType SudtTransactionType, sudtOwnerAddress string) (*SudtTransactionBuilder, error) {

	addr, err := address2.Decode(sudtOwnerAddress)
	if err != nil {
		return nil, err
	}
	sudtArgs := addr.Script.Hash()
	return NewSudtTransactionBuilderFromSudtArgs(network, iterator, transactionType, sudtArgs.Bytes()), nil
}

func (r *SudtTransactionBuilder) AddSudtOutputByAddress(addr string, sudtAmount *big.Int) (int, error) {
	a, err := address2.Decode(addr)
	if err != nil {
		return 0, err
	}
	output := &types.CellOutput{
		Capacity: 0,
		Lock:     a.Script,
		Type:     r.SudtType,
	}
	data := systemscript.EncodeSudtAmount(sudtAmount)
	output.Capacity = output.OccupiedCapacity(data)
	return r.AddOutput(output, data), nil
}

func (r *SudtTransactionBuilder) AddSudtOutputWithCapacityByAddress(addr string, capacity uint64, sudtAmount *big.Int) (int, error) {
	a, err := address2.Decode(addr)
	if err != nil {
		return 0, err
	}
	output := &types.CellOutput{
		Capacity: capacity,
		Lock:     a.Script,
		Type:     r.SudtType,
	}
	data := systemscript.EncodeSudtAmount(sudtAmount)
	return r.AddOutput(output, data), nil
}

func (r *SudtTransactionBuilder) AddChangeOutputByAddress(addr string) error {
	if r.changeOutputIndex != -1 {
		return errors.New("change output has been set")
	}
	err := r.AddOutputByAddress(addr, 0)
	if err == nil {
		r.changeOutputIndex = len(r.Outputs) - 1
	}
	return err
}

func (r *SudtTransactionBuilder) Build(contexts ...interface{}) (*transaction.TransactionWithScriptGroups, error) {
	if r.SudtType == nil {
		return nil, errors.New("sudt type is not set")
	}

	// If transaction type is SudtTransactionTypeTransfer, we need the change output to receive SUDT
	if r.transactionType == SudtTransactionTypeTransfer {
		r.Outputs[r.changeOutputIndex].Type = r.SudtType
		r.OutputsData[r.changeOutputIndex] = systemscript.EncodeSudtAmount(big.NewInt(0))
	}

	var (
		err              error
		s                *types.Script
		g                *transaction.ScriptGroup
		m                = make(map[types.Hash]*transaction.ScriptGroup)
		outputsCapacity  = uint64(0)
		outputSudtAmount = big.NewInt(0)
	)
	for i := 0; i < len(r.Outputs); i++ {
		outputsCapacity += r.Outputs[i].Capacity
		data := r.OutputsData[i]
		if err := addSudtAmount(outputSudtAmount, data); err != nil {
			return nil, err
		}
		s = r.Outputs[i].Type
		if s != nil {
			if g, err = getOrPutScriptGroup(m, s, types.ScriptTypeType); err != nil {
				return nil, err
			}
			g.OutputIndices = append(g.OutputIndices, uint32(i))
			if err := executeHandlers(&r.SimpleTransactionBuilder, g, contexts); err != nil {
				return nil, err
			}
		}
	}

	var (
		enoughCapacity   = false
		inputsCapacity   = uint64(0)
		inputsSudtAmount = big.NewInt(0)
		i                = -1
	)
	for {
		cell := r.getNextCell() // only get SUDT cell
		if cell == nil {
			break // break when can't find cell
		}
		r.AddInput(&types.CellInput{
			Since:          0,
			PreviousOutput: cell.OutPoint,
		})
		i += 1

		// process input's LOCK
		s = cell.Output.Lock
		if s != nil {
			if g, err = getOrPutScriptGroup(m, s, types.ScriptTypeLock); err != nil {
				return nil, err
			}
			g.InputIndices = append(g.InputIndices, uint32(i))
			if err := executeHandlers(&r.SimpleTransactionBuilder, g, contexts...); err != nil {
				return nil, err
			}
		}

		// process input's TYPE
		s = cell.Output.Type
		if s != nil {
			if g, err = getOrPutScriptGroup(m, s, types.ScriptTypeType); err != nil {
				return nil, err
			}
			g.InputIndices = append(g.InputIndices, uint32(i))
			if err := executeHandlers(&r.SimpleTransactionBuilder, g, contexts...); err != nil {
				return nil, err
			}
		}

		inputsCapacity += cell.Output.Capacity
		if err := addSudtAmount(inputsSudtAmount, cell.OutputData); err != nil {
			return nil, err
		}
		// continue to iterator if no enough SUDT amount in transfer mod
		if r.transactionType == SudtTransactionTypeTransfer && inputsSudtAmount.Cmp(outputSudtAmount) < 0 {
			continue
		}

		tx := r.BuildTransaction().TxView
		// check if there is enough capacity for output capacity and change
		fee := tx.CalculateFee(uint64(r.FeeRate))
		if inputsCapacity < (outputsCapacity + fee) {
			continue
		}
		changeCapacity := inputsCapacity - outputsCapacity - fee
		changeOutput := r.Outputs[r.changeOutputIndex]
		changeOutputData := r.OutputsData[r.changeOutputIndex]
		if changeCapacity >= changeOutput.OccupiedCapacity(changeOutputData) {
			changeOutput.Capacity = changeCapacity
			if r.transactionType == SudtTransactionTypeTransfer {
				diff := big.NewInt(0)
				diff.Sub(inputsSudtAmount, outputSudtAmount)
				r.OutputsData[r.changeOutputIndex] = systemscript.EncodeSudtAmount(diff)
			}
			enoughCapacity = true
			break
		}
	}
	if !enoughCapacity {
		return nil, errors.New("no enough capacity")
	}
	r.scriptGroups = make([]*transaction.ScriptGroup, 0)
	for _, g := range m {
		r.scriptGroups = append(r.scriptGroups, g)
	}
	return r.BuildTransaction(), nil
}

func (r *SudtTransactionBuilder) getNextCell() *types.TransactionInput {
	for {
		if !r.iterator.HasNext() {
			return nil
		}
		cell := r.iterator.Next()
		// filter cell that has SUDT type
		if reflect.DeepEqual(cell.Output.Type, r.SudtType) {
			return cell
		}
	}
}

func addSudtAmount(a *big.Int, b []byte) error {
	if len(b) == 0 {
		return nil
	}
	amount, err := systemscript.DecodeSudtAmount(b)
	if err != nil {
		return err
	}
	a.Add(a, amount)
	return nil
}
