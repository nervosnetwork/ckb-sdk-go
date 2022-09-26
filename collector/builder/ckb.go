package builder

import (
	"errors"
	address2 "github.com/nervosnetwork/ckb-sdk-go/address"
	"github.com/nervosnetwork/ckb-sdk-go/collector"
	"github.com/nervosnetwork/ckb-sdk-go/collector/handler"
	"github.com/nervosnetwork/ckb-sdk-go/transaction"
	"github.com/nervosnetwork/ckb-sdk-go/types"
)

type CkbTransactionBuilder struct {
	SimpleTransactionBuilder
	FeeRate uint

	iterator               collector.CellIterator
	transactionInputs      []*types.TransactionInput // customized inputs
	transactionInputsIndex int

	changeOutputIndex int
	reward            uint64
}

func NewCkbTransactionBuilder(network types.Network, iterator collector.CellIterator) *CkbTransactionBuilder {
	return &CkbTransactionBuilder{
		SimpleTransactionBuilder: *NewSimpleTransactionBuilder(network),
		FeeRate:                  1000,

		iterator:          iterator,
		changeOutputIndex: -1,
	}
}

func (r *CkbTransactionBuilder) AddChangeOutputByAddress(addr string) error {
	if r.changeOutputIndex != -1 {
		return errors.New("change output has been set")
	}
	err := r.AddOutputByAddress(addr, 0)
	if err == nil {
		r.changeOutputIndex = len(r.Outputs) - 1
	}
	return err
}

func (r *CkbTransactionBuilder) AddDaoDepositOutputByAddress(addr string, capacity uint64) error {
	a, err := address2.Decode(addr)
	if err != nil {
		return err
	}
	output := &types.CellOutput{
		Capacity: capacity,
		Lock:     a.Script,
		Type:     handler.DaoScript,
	}
	data := handler.DaoDepositOutputData
	r.AddOutput(output, data)
	return nil
}

func getOrPutScriptGroup(m map[types.Hash]*transaction.ScriptGroup, script *types.Script, scriptType types.ScriptType) (*transaction.ScriptGroup, error) {
	if script == nil {
		return nil, nil
	}
	hash := script.Hash()
	if m[hash] == nil {
		m[hash] = &transaction.ScriptGroup{
			Script:    script,
			GroupType: scriptType,
		}
	}
	return m[hash], nil
}

func executeHandlers(s *SimpleTransactionBuilder, group *transaction.ScriptGroup, contexts ...interface{}) error {
	if len(contexts) == 0 {
		contexts = append(contexts, nil)
	}
	for _, v := range s.ScriptHandlers {
		for _, c := range contexts {
			if _, err := v.BuildTransaction(s, group, c); err != nil {
				return err
			}
		}
	}
	return nil
}

func (r *CkbTransactionBuilder) Build(contexts ...interface{}) (*transaction.TransactionWithScriptGroups, error) {
	var (
		err             error
		script          *types.Script
		group           *transaction.ScriptGroup
		m               = make(map[types.Hash]*transaction.ScriptGroup)
		outputsCapacity = uint64(0)
	)
	for i := 0; i < len(r.Outputs); i++ {
		outputsCapacity += r.Outputs[i].Capacity
		script = r.Outputs[i].Type
		if script != nil {
			if group, err = getOrPutScriptGroup(m, script, types.ScriptTypeType); err != nil {
				return nil, err
			}
			group.OutputIndices = append(group.OutputIndices, uint32(i))
			if err := executeHandlers(&r.SimpleTransactionBuilder, group, contexts...); err != nil {
				return nil, err
			}
		}
	}

	var (
		enoughCapacity = false
		inputsCapacity = uint64(0)
		i              = -1
	)
	for {
		cell := r.getNextCell()
		if cell == nil {
			break // break when can't find cell
		}
		r.AddInput(&types.CellInput{
			Since:          0,
			PreviousOutput: cell.OutPoint,
		})
		i += 1

		// process input's LOCK
		script = cell.Output.Lock
		if script != nil {
			if group, err = getOrPutScriptGroup(m, script, types.ScriptTypeLock); err != nil {
				return nil, err
			}
			group.InputIndices = append(group.InputIndices, uint32(i))
			if err := executeHandlers(&r.SimpleTransactionBuilder, group, contexts...); err != nil {
				return nil, err
			}
		}

		// process input's TYPE
		script = cell.Output.Type
		if script != nil {
			if group, err = getOrPutScriptGroup(m, script, types.ScriptTypeType); err != nil {
				return nil, err
			}
			group.InputIndices = append(group.InputIndices, uint32(i))
			if err := executeHandlers(&r.SimpleTransactionBuilder, group, contexts...); err != nil {
				return nil, err
			}
		}

		inputsCapacity += cell.Output.Capacity
		tx := r.BuildTransaction().TxView
		// check if there is enough capacity for output capacity and change
		fee := tx.CalculateFee(uint64(r.FeeRate))
		if (inputsCapacity + r.reward) < (outputsCapacity + fee) {
			continue
		}
		changeCapacity := inputsCapacity + r.reward - outputsCapacity - fee
		changeOutput := r.Outputs[r.changeOutputIndex]
		changeOutputData := r.OutputsData[r.changeOutputIndex]
		if changeCapacity >= changeOutput.OccupiedCapacity(changeOutputData) {
			changeOutput.Capacity = changeCapacity
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

func (r *CkbTransactionBuilder) getNextCell() *types.TransactionInput {
	// consume customized inputs at first
	if r.transactionInputsIndex < len(r.transactionInputs) {
		t := r.transactionInputs[r.transactionInputsIndex]
		r.transactionInputsIndex += 1
		return t
	}

	for {
		if !r.iterator.HasNext() {
			return nil
		}
		cell := r.iterator.Next()
		// filter cell with non-nil type
		if cell.Output.Type == nil {
			return cell
		}
	}
}
