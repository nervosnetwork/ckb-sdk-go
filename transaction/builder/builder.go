package builder

import (
	"github.com/nervosnetwork/ckb-sdk-go/types"
)

type UnsignedTxBuilder interface {
	NewTransaction()
	BuildVersion()
	BuildHeaderDeps()
	BuildCellDeps()
	BuildOutputsAndOutputsData()
	BuildInputsAndWitnesses() error
	UpdateChangeOutput() error
	GetResult() *types.Transaction
}

type Director struct {
	builder UnsignedTxBuilder
}

func (d *Director) SetBuilder(builder UnsignedTxBuilder) {
	d.builder = builder
}

func (d *Director) Generate() (*types.Transaction, error) {
	d.builder.NewTransaction()
	d.builder.BuildVersion()
	d.builder.BuildHeaderDeps()
	d.builder.BuildCellDeps()
	d.builder.BuildOutputsAndOutputsData()
	err := d.builder.BuildInputsAndWitnesses()
	if err != nil {
		return nil, err
	}
	err = d.builder.UpdateChangeOutput()

	return d.builder.GetResult(), err
}
