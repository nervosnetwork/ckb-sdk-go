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

type SystemScriptCell struct {
	CellHash types.Hash
	OutPoint *types.OutPoint
}

type SystemScripts struct {
	SecpSingleSigCell *SystemScriptCell
	SecpMultiSigCell  *SystemScriptCell
	DaoCell           *SystemScriptCell
}

func NewSystemScripts(client rpc.Client) (*SystemScripts, error) {
	genesis, err := client.GetBlockByNumber(context.Background(), 0)
	if err != nil {
		return nil, err
	}

	secpHash, err := genesis.Transactions[0].Outputs[1].Type.Hash()
	if err != nil {
		return nil, err
	}
	multiSigHash, err := genesis.Transactions[0].Outputs[4].Type.Hash()
	if err != nil {
		return nil, err
	}
	daoHash, err := genesis.Transactions[0].Outputs[2].Type.Hash()
	if err != nil {
		return nil, err
	}

	return &SystemScripts{
		SecpSingleSigCell: &SystemScriptCell{
			CellHash: secpHash,
			OutPoint: &types.OutPoint{
				TxHash: genesis.Transactions[1].Hash,
				Index:  0,
			},
		},
		SecpMultiSigCell: &SystemScriptCell{
			CellHash: multiSigHash,
			OutPoint: &types.OutPoint{
				TxHash: genesis.Transactions[1].Hash,
				Index:  1,
			},
		},
		DaoCell: &SystemScriptCell{
			CellHash: daoHash,
			OutPoint: &types.OutPoint{
				TxHash: genesis.Transactions[0].Hash,
				Index:  2,
			},
		},
	}, nil
}
