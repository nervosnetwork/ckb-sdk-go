package builder

import (
	"github.com/nervosnetwork/ckb-sdk-go/types"
)

type UnsignedTxBuilder interface {
	NewTransaction()
	BuildVersion()
	BuildHeaderDeps()
	BuildCellDeps()
	BuildOutputsAndOutputsData() error
	BuildInputsAndWitnesses() error
	UpdateChangeOutput() error
	GetResult() (*types.Transaction, map[string][]int)
}

type Director struct {
	builder UnsignedTxBuilder
}

func (d *Director) SetBuilder(builder UnsignedTxBuilder) {
	d.builder = builder
}

func (d *Director) Generate() (*types.Transaction, map[string][]int, error) {
	d.builder.NewTransaction()
	d.builder.BuildVersion()
	d.builder.BuildHeaderDeps()
	d.builder.BuildCellDeps()
	err := d.builder.BuildOutputsAndOutputsData()
	if err != nil {
		return nil, nil, err
	}
	err = d.builder.BuildInputsAndWitnesses()
	if err != nil {
		return nil, nil, err
	}
	err = d.builder.UpdateChangeOutput()
	tx, groups := d.builder.GetResult()

	return tx, groups, err
}
