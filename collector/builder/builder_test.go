package builder

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/nervosnetwork/ckb-sdk-go/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

type mockIterator struct {
	Cells []*types.TransactionInput
	index int
}

func (m *mockIterator) HasNext() bool {
	return m.index < len(m.Cells)
}

func (m *mockIterator) Next() *types.TransactionInput {
	current := m.Cells[m.index]
	m.index += 1
	return current
}

var (
	lock = &types.Script{
		CodeHash: types.HexToHash("0x9bd7e06f3ecf4be0f2fcd2188b23f1b9fcc88e5d4b65a8637b17723bbda3cce8"),
		HashType: types.HashTypeType,
		Args:     common.FromHex("0xeac21ac6d373414aaa9ba34c469f805d48b62f86"),
	}
)

func getMockIterator() *mockIterator {
	return &mockIterator{
		Cells: []*types.TransactionInput{
			{
				OutPoint: &types.OutPoint{
					TxHash: types.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
					Index:  0,
				},
				Output: &types.CellOutput{
					Capacity: 100000000000,
					Lock:     lock,
				},
				OutputData: []byte{},
			},
			{
				OutPoint: &types.OutPoint{
					TxHash: types.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
					Index:  1,
				},
				Output: &types.CellOutput{
					Capacity: 10000000000,
					Lock:     lock,
				},
				OutputData: []byte{},
			},
		},
	}
}

func TestCkbTransactionBuilderSingleInput(t *testing.T) {
	iterator := getMockIterator()
	builder := NewCkbTransactionBuilder(types.NetworkTest, iterator)
	builder.FeeRate = 1000
	builder.AddOutputByAddress("ckt1qzda0cr08m85hc8jlnfp3zer7xulejywt49kt2rr0vthywaa50xwsq2qf8keemy2p5uu0g0gn8cd4ju23s5269qk8rg4r", 50100000000)
	err := builder.AddChangeOutputByAddress("ckt1qzda0cr08m85hc8jlnfp3zer7xulejywt49kt2rr0vthywaa50xwsq02cgdvd5mng9924xarf3rflqzafzmzlpsuhh83c")
	if err != nil {
		t.Error(err)
	}
	tx, err := builder.Build()
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, 1, len(tx.TxView.Inputs))
	assert.Equal(t, 1, len(tx.ScriptGroups))
	assert.Equal(t, 2, len(tx.TxView.Outputs))
	assert.Equal(t, *lock, *tx.ScriptGroups[0].Script)
	fee := 100000000000 - tx.TxView.Outputs[0].Capacity - tx.TxView.Outputs[1].Capacity
	assert.Equal(t, uint64(464), fee)
}

func TestCkbTransactionBuilderMultipleInputs(t *testing.T) {
	iterator := getMockIterator()
	builder := NewCkbTransactionBuilder(types.NetworkTest, iterator)
	builder.FeeRate = 1000
	builder.AddOutputByAddress("ckt1qzda0cr08m85hc8jlnfp3zer7xulejywt49kt2rr0vthywaa50xwsq2qf8keemy2p5uu0g0gn8cd4ju23s5269qk8rg4r", 100000000000)
	err := builder.AddChangeOutputByAddress("ckt1qzda0cr08m85hc8jlnfp3zer7xulejywt49kt2rr0vthywaa50xwsq02cgdvd5mng9924xarf3rflqzafzmzlpsuhh83c")
	if err != nil {
		t.Error(err)
	}
	tx, err := builder.Build()
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, 2, len(tx.TxView.Inputs))
	assert.Equal(t, 1, len(tx.ScriptGroups))
	assert.Equal(t, 2, len(tx.TxView.Outputs))
	assert.Equal(t, *lock, *tx.ScriptGroups[0].Script)
	fee := 110000000000 - tx.TxView.Outputs[0].Capacity - tx.TxView.Outputs[1].Capacity
	assert.Equal(t, uint64(516), fee)
}