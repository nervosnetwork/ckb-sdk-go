package builder

import (
	"errors"
	"github.com/nervosnetwork/ckb-sdk-go/address"
	"github.com/nervosnetwork/ckb-sdk-go/collector"
	"github.com/nervosnetwork/ckb-sdk-go/collector/handler"
	"github.com/nervosnetwork/ckb-sdk-go/transaction"
	"github.com/nervosnetwork/ckb-sdk-go/types"
	"reflect"
	"strconv"
)

type SimpleTransactionBuilder struct {
	Version     uint
	CellDeps    []*types.CellDep
	HeaderDeps  []types.Hash
	Inputs      []*types.CellInput
	Outputs     []*types.CellOutput
	OutputsData [][]byte
	Witnesses   [][]byte

	scriptGroups   []*transaction.ScriptGroup
	ScriptHandlers []collector.ScriptHandler
}

func NewSimpleTransactionBuilder(network types.Network) *SimpleTransactionBuilder {
	if network == types.NetworkMain || network == types.NetworkTest {
		s := SimpleTransactionBuilder{}
		s.Register(handler.NewSecp256k1Blake160SighashAllScriptHandler(network))
		s.Register(handler.NewSecp256k1Blake160MultisigAllScriptHandler(network))
		s.Register(handler.NewSudtScriptHandler(network))
		s.Register(handler.NewDaoScriptHandler(network))
		return &s
	} else {
		return nil
	}
}

func (r *SimpleTransactionBuilder) Register(handler collector.ScriptHandler) {
	r.ScriptHandlers = append(r.ScriptHandlers, handler)
}

func (r *SimpleTransactionBuilder) SetVersion(version uint) {
	r.Version = version
}

func (r *SimpleTransactionBuilder) AddHeaderDep(headerDep types.Hash) int {
	for i, v := range r.HeaderDeps {
		if reflect.DeepEqual(v, headerDep) {
			return i
		}
	}
	r.HeaderDeps = append(r.HeaderDeps, headerDep)
	return len(r.HeaderDeps) - 1
}

func (r *SimpleTransactionBuilder) AddCellDep(cellDep *types.CellDep) int {
	for i, v := range r.CellDeps {
		if reflect.DeepEqual(v, cellDep) {
			return i
		}
	}
	r.CellDeps = append(r.CellDeps, cellDep)
	return len(r.CellDeps) - 1
}

func (r *SimpleTransactionBuilder) AddInput(input *types.CellInput) int {
	r.Inputs = append(r.Inputs, input)
	r.Witnesses = append(r.Witnesses, []byte{})
	return len(r.Inputs) - 1
}

func (r *SimpleTransactionBuilder) SetSince(index uint, since uint64) error {
	if index >= uint(len(r.Inputs)) {
		return errors.New("index " + strconv.Itoa(int(index)) + " out of range")
	}
	r.Inputs[index].Since = since
	return nil
}

func (r *SimpleTransactionBuilder) AddOutput(output *types.CellOutput, data []byte) int {
	r.Outputs = append(r.Outputs, output)
	r.OutputsData = append(r.OutputsData, data)
	return len(r.Outputs) - 1
}

func (r *SimpleTransactionBuilder) AddOutputByAddress(addr string, capacity uint64) error {
	a, err := address.Decode(addr)
	if err != nil {
		return err
	}
	output := &types.CellOutput{
		Capacity: capacity,
		Lock:     a.Script,
		Type:     nil,
	}
	r.AddOutput(output, []byte{})
	return nil
}

func (r *SimpleTransactionBuilder) SetOutputData(index uint, data []byte) error {
	if index >= uint(len(r.OutputsData)) {
		return errors.New("index " + strconv.Itoa(int(index)) + " out of range")
	}
	r.OutputsData[index] = data
	return nil
}

func (r *SimpleTransactionBuilder) SetWitness(index uint, witnessType types.WitnessType, data []byte) error {
	if index >= uint(len(r.Witnesses)) {
		return errors.New("index " + strconv.Itoa(int(index)) + " out of range")
	}
	var wArgs *types.WitnessArgs
	var err error
	w := r.Witnesses[index]
	if len(w) == 0 {
		wArgs = &types.WitnessArgs{}
	} else if wArgs, err = types.DeserializeWitnessArgs(w); err != nil {
		return err
	}

	switch witnessType {
	case types.WitnessTypeLock:
		wArgs.Lock = data
	case types.WitnessTypeInputType:
		wArgs.InputType = data
	case types.WitnessTypeOutputType:
		wArgs.OutputType = data
	default:
		return errors.New("unknown data type " + strconv.Itoa(int(witnessType)))
	}
	w = wArgs.Serialize()
	r.Witnesses[index] = w
	return nil
}

func (r *SimpleTransactionBuilder) AddScriptGroup(group *transaction.ScriptGroup) int {
	r.scriptGroups = append(r.scriptGroups, group)
	return len(r.scriptGroups) - 1
}

func (r *SimpleTransactionBuilder) Build(contexts ...interface{}) (*transaction.TransactionWithScriptGroups, error) {
	return r.BuildTransaction(), nil
}

func (r *SimpleTransactionBuilder) BuildTransaction() *transaction.TransactionWithScriptGroups {
	return &transaction.TransactionWithScriptGroups{
		TxView: &types.Transaction{
			Version:     r.Version,
			CellDeps:    r.CellDeps,
			HeaderDeps:  r.HeaderDeps,
			Inputs:      r.Inputs,
			Outputs:     r.Outputs,
			OutputsData: r.OutputsData,
			Witnesses:   r.Witnesses,
		},
		ScriptGroups: r.scriptGroups,
	}
}
