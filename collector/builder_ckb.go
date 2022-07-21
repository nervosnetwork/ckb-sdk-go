package collector

import (
	"errors"
	"github.com/nervosnetwork/ckb-sdk-go/transaction"
	"github.com/nervosnetwork/ckb-sdk-go/types"
)

type CkbTransactionBuilder struct {
	SimpleTransactionBuilder
	FeeRate uint
	Network types.Network

	iterator          CellIterator
	changeOutputIndex int
	reward            uint64
}

func NewCkbTransactionBuilder(network types.Network, iterator CellIterator) *CkbTransactionBuilder {
	return &CkbTransactionBuilder{
		SimpleTransactionBuilder: *NewSimpleTransactionBuilder(network),
		FeeRate:                  1000,
		Network:                  network,

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

func getOrPutScriptGroup(m map[types.Hash]*transaction.ScriptGroup, script *types.Script, scriptType transaction.ScriptType) (*transaction.ScriptGroup, error) {
	if script == nil {
		return nil, nil
	}
	hash, err := script.Hash()
	if err != nil {
		return nil, err
	}
	if m[hash] == nil {
		m[hash] = &transaction.ScriptGroup{
			Script:    script,
			GroupType: scriptType,
		}
	}
	return m[hash], nil

}

func (r *CkbTransactionBuilder) executeHandlers(group *transaction.ScriptGroup, contexts ...interface{}) error {
	for _, v := range r.ScriptHandlers {
		for _, c := range contexts {
			if _, err := v.BuildTransaction(r, group, c); err != nil {
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
			if group, err = getOrPutScriptGroup(m, script, transaction.ScriptTypeType); err != nil {
				return nil, err
			}
			group.OutputIndices = append(group.OutputIndices, uint32(i))
			if err := r.executeHandlers(group, contexts); err != nil {
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
		if !r.iterator.HasNext() {
			break
		}
		cell := r.iterator.Next()
		r.AddInput(&types.CellInput{
			Since:          0,
			PreviousOutput: cell.OutPoint,
		})
		i += 1

		// process input's LOCK
		script = cell.Output.Lock
		if script != nil {
			if group, err = getOrPutScriptGroup(m, script, transaction.ScriptTypeLock); err != nil {
				return nil, err
			}
			group.InputIndices = append(group.InputIndices, uint32(i))
			if err := r.executeHandlers(group, contexts); err != nil {
				return nil, err
			}
		}

		// process input's TYPE
		script = cell.Output.Type
		if script != nil {
			if group, err = getOrPutScriptGroup(m, script, transaction.ScriptTypeType); err != nil {
				return nil, err
			}
			group.InputIndices = append(group.InputIndices, uint32(i))
			if err := r.executeHandlers(group, contexts); err != nil {
				return nil, err
			}
		}

		inputsCapacity += cell.Output.Capacity
		tx := r.BuildTransaction().TxView
		// check if there is enough capacity for output capacity and change
		fee, err := transaction.CalculateTransactionFee(tx, uint64(r.FeeRate))
		if err != nil {
			return nil, err
		}
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