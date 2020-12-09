package utils

import (
	"context"

	"github.com/nervosnetwork/ckb-sdk-go/rpc"
	"github.com/nervosnetwork/ckb-sdk-go/types"
)

const (
	AnyoneCanPayCodeHashOnLina   = "0xd369597ff47f29fbc0d47d2e3775370d1250b85140c670e4718af712983a2354"
	AnyoneCanPayCodeHashOnAggron = "0x3419a1c09eb2567f6552ee7a8ecffd64155cffe0f1796e6e61ec088d740c1356"
)

type Option func(*SystemScripts)
type SystemScriptCell struct {
	CellHash types.Hash
	OutPoint *types.OutPoint
}

type SystemScripts struct {
	SecpSingleSigCell *SystemScriptCell
	SecpMultiSigCell  *SystemScriptCell
	DaoCell           *SystemScriptCell
	ACPCell           *SystemScriptCell
	SUDTCell          *SystemScriptCell
}

func secpSingleSigCell() *SystemScriptCell {
	return &SystemScriptCell{
		CellHash: types.HexToHash("0x9bd7e06f3ecf4be0f2fcd2188b23f1b9fcc88e5d4b65a8637b17723bbda3cce8"),
		OutPoint: &types.OutPoint{
			TxHash: types.HexToHash("0x71a7ba8fc96349fea0ed3a5c47992e3b4084b031a42264a018e0072e8172e46c"),
			Index:  0,
		},
	}
}

func secpMultiSigCell() *SystemScriptCell {
	return &SystemScriptCell{
		CellHash: types.HexToHash("0x5c5069eb0857efc65e1bca0c07df34c31663b3622fd3876c876320fc9634e2a8"),
		OutPoint: &types.OutPoint{
			TxHash: types.HexToHash("0x71a7ba8fc96349fea0ed3a5c47992e3b4084b031a42264a018e0072e8172e46c"),
			Index:  1,
		},
	}
}

func daoCell() *SystemScriptCell {
	return &SystemScriptCell{
		CellHash: types.HexToHash("0x82d76d1b75fe2fd9a27dfbaa65a039221a380d76c926f378d3f81cf3e7e13f2e"),
		OutPoint: &types.OutPoint{
			TxHash: types.HexToHash("0xe2fb199810d49a4d8beec56718ba2593b665db9d52299a0f9e6e75416d73ff5c"),
			Index:  2,
		},
	}
}

func acpCell(chain string) *SystemScriptCell {
	if chain == "ckb" {
		return &SystemScriptCell{
			CellHash: types.HexToHash("0xd369597ff47f29fbc0d47d2e3775370d1250b85140c670e4718af712983a2354"),
			OutPoint: &types.OutPoint{
				TxHash: types.HexToHash("0x4153a2014952d7cac45f285ce9a7c5c0c0e1b21f2d378b82ac1433cb11c25c4d"),
				Index:  0,
			},
		}
	} else {
		return &SystemScriptCell{
			CellHash: types.HexToHash("0x3419a1c09eb2567f6552ee7a8ecffd64155cffe0f1796e6e61ec088d740c1356"),
			OutPoint: &types.OutPoint{
				TxHash: types.HexToHash("0xec26b0f85ed839ece5f11c4c4e837ec359f5adc4420410f6453b1f6b60fb96a6"),
				Index:  0,
			},
		}
	}
}

func sudtCell(chain string) *SystemScriptCell {
	if chain == "ckb" {
		return &SystemScriptCell{
			CellHash: types.HexToHash("0x5e7a36a77e68eecc013dfa2fe6a23f3b6c344b04005808694ae6dd45eea4cfd5"),
			OutPoint: &types.OutPoint{
				TxHash: types.HexToHash("0xc7813f6a415144643970c2e88e0bb6ca6a8edc5dd7c1022746f628284a9936d5"),
				Index:  0,
			},
		}
	} else {
		return &SystemScriptCell{
			CellHash: types.HexToHash("0x48dbf59b4c7ee1547238021b4869bceedf4eea6b43772e5d66ef8865b6ae7212"),
			OutPoint: &types.OutPoint{
				TxHash: types.HexToHash("0xc1b2ae129fad7465aaa9acc9785f842ba3e6e8b8051d899defa89f5508a77958"),
				Index:  0,
			},
		}
	}
}

// NewSystemScripts returns a SystemScripts object
func NewSystemScripts(client rpc.Client, options ...Option) (*SystemScripts, error) {
	info, err := client.GetBlockchainInfo(context.Background())
	if err != nil {
		return nil, err
	}
	scripts := &SystemScripts{
		SecpSingleSigCell: secpSingleSigCell(),
		SecpMultiSigCell:  secpMultiSigCell(),
		DaoCell:           daoCell(),
		ACPCell:           acpCell(info.Chain),
		SUDTCell:          sudtCell(info.Chain),
	}

	for _, option := range options {
		option(scripts)
	}

	return scripts, nil
}

// SecpSingleSigCell set a custom secp single sig cell to SystemScripts
func SecpSingleSigCell(secpSingleSigCell *SystemScriptCell) Option {
	return func(s *SystemScripts) {
		s.SecpSingleSigCell = secpSingleSigCell
	}
}

// SecpMultiSigCell set a custom secp mutisig cell to SystemScripts
func SecpMultiSigCell(secpMultiSigCell *SystemScriptCell) Option {
	return func(s *SystemScripts) {
		s.SecpMultiSigCell = secpMultiSigCell
	}
}

// DaoCell set a custom dao cell to SystemScripts
func DaoCell(daoCell *SystemScriptCell) Option {
	return func(s *SystemScripts) {
		s.DaoCell = daoCell
	}
}

// ACPCell set a custom acp system script cell to SystemScripts
func ACPCell(acpCell *SystemScriptCell) Option {
	return func(s *SystemScripts) {
		s.ACPCell = acpCell
	}
}

// SUDTCell set a custom sudt system script cell to SystemScripts
func SUDTCell(sudtCell *SystemScriptCell) Option {
	return func(s *SystemScripts) {
		s.SUDTCell = sudtCell
	}
}
