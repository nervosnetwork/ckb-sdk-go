package dao

import (
	"errors"

	"github.com/nervosnetwork/ckb-sdk-go/types"
	"github.com/nervosnetwork/ckb-sdk-go/utils"
)

type Deposit struct {
	Transaction *types.Transaction
}

func NewDeposit(scripts *utils.SystemScripts, isMultisig bool) *Deposit {
	var baseDep *types.CellDep
	if isMultisig {
		baseDep = &types.CellDep{
			OutPoint: scripts.SecpMultiSigCell.OutPoint,
			DepType:  types.DepTypeDepGroup,
		}
	} else {
		baseDep = &types.CellDep{
			OutPoint: scripts.SecpSingleSigCell.OutPoint,
			DepType:  types.DepTypeDepGroup,
		}
	}

	tx := &types.Transaction{
		Version:    0,
		HeaderDeps: []types.Hash{},
		CellDeps: []*types.CellDep{
			baseDep,
			{
				OutPoint: scripts.DaoCell.OutPoint,
				DepType:  types.DepTypeCode,
			},
		},
	}

	return &Deposit{
		Transaction: tx,
	}
}

func (d *Deposit) AddDaoOutput(scripts *utils.SystemScripts, lock *types.Script, amount uint64) error {
	if d.Transaction == nil {
		return errors.New("must init transaction first")
	}
	d.Transaction.Outputs = append(d.Transaction.Outputs, &types.CellOutput{
		Capacity: amount,
		Lock:     lock,
		Type: &types.Script{
			CodeHash: scripts.DaoCell.CodeHash,
			HashType: types.HashTypeType,
			Args:     []byte{},
		},
	})
	d.Transaction.OutputsData = append(d.Transaction.OutputsData, make([]byte, 8))

	return nil
}

func (d *Deposit) AddOutput(lock *types.Script, amount uint64) error {
	if d.Transaction == nil {
		return errors.New("must init transaction first")
	}
	d.Transaction.Outputs = append(d.Transaction.Outputs, &types.CellOutput{
		Capacity: amount,
		Lock:     lock,
	})
	d.Transaction.OutputsData = append(d.Transaction.OutputsData, []byte{})

	return nil
}
