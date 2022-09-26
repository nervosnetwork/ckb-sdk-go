package handler

import (
	"bytes"
	"context"
	"errors"
	"github.com/nervosnetwork/ckb-sdk-go/collector"
	"github.com/nervosnetwork/ckb-sdk-go/rpc"
	"github.com/nervosnetwork/ckb-sdk-go/script"
	"github.com/nervosnetwork/ckb-sdk-go/transaction"
	"github.com/nervosnetwork/ckb-sdk-go/types"
	"reflect"
)

const DaoLockPeriodEpochs = 180

var (
	DaoDepositOutputData = []byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}
	DaoScript            = &types.Script{
		CodeHash: script.GetCodeHash(types.NetworkMain, script.SystemScriptDao),
		HashType: types.HashTypeType,
		Args:     []byte{},
	}
)

type DaoScriptHandler struct {
	CellDep  *types.CellDep
	CodeHash types.Hash
}

func NewDaoScriptHandler(network types.Network) *DaoScriptHandler {
	var txHash types.Hash
	if network == types.NetworkMain {
		txHash = types.HexToHash("0xe2fb199810d49a4d8beec56718ba2593b665db9d52299a0f9e6e75416d73ff5c")
	} else if network == types.NetworkTest {
		txHash = types.HexToHash("0x8f8c79eb6671709633fe6a46de93c0fedc9c1b8a6527a18d3983879542635c9f")
	} else {
		return nil
	}

	return &DaoScriptHandler{
		CellDep: &types.CellDep{
			OutPoint: &types.OutPoint{
				TxHash: txHash,
				Index:  2,
			},
			DepType: types.DepTypeCode,
		},
		CodeHash: script.GetCodeHash(network, script.SystemScriptDao),
	}
}

func (r *DaoScriptHandler) isMatched(script *types.Script) bool {
	if script == nil {
		return false
	}
	return reflect.DeepEqual(script.CodeHash, r.CodeHash)
}

func IsDepositCell(output *types.CellOutput, outputData []byte) bool {
	return reflect.DeepEqual(DaoScript, output.Type) &&
		bytes.Equal(DaoDepositOutputData, outputData)
}

func (r *DaoScriptHandler) BuildTransaction(builder collector.TransactionBuilder, group *transaction.ScriptGroup, context interface{}) (bool, error) {
	if group == nil || !r.isMatched(group.Script) {
		return false, nil
	}

	builder.AddCellDep(r.CellDep)

	var ok bool
	switch context.(type) {
	case ClaimInfo, *ClaimInfo:
		var claimInfo *ClaimInfo
		if claimInfo, ok = context.(*ClaimInfo); !ok {
			v, _ := context.(ClaimInfo)
			claimInfo = &v
		}
		index := group.InputIndices[0]
		depositHeaderDepIndex := builder.AddHeaderDep(claimInfo.DepositBlockHeader.Hash)
		builder.AddHeaderDep(claimInfo.WithdrawBlockHeader.Hash)
		inputType := types.SerializeUint64(uint64(depositHeaderDepIndex))
		builder.SetWitness(uint(index), types.WitnessTypeInputType, inputType)
		builder.SetSince(uint(index), claimInfo.CalculateDaoMinimumSince())
	case WithdrawInfo, *WithdrawInfo:
		var withdrawInfo *WithdrawInfo
		if withdrawInfo, ok = context.(*WithdrawInfo); !ok {
			v, _ := context.(WithdrawInfo)
			withdrawInfo = &v
		}
		builder.AddHeaderDep(withdrawInfo.DepositBlockHash)
	default:
	}
	return true, nil
}

type ClaimInfo struct {
	DepositBlockHeader  *types.Header
	WithdrawBlockHeader *types.Header
	WithdrawOutpoint    *types.OutPoint
}

func NewClaimInfo(client rpc.Client, withdrawOutpoint *types.OutPoint) (*ClaimInfo, error) {
	txWithStatus, err := client.GetTransaction(context.Background(), withdrawOutpoint.TxHash)
	if err != nil {
		return nil, err
	}
	withdrawTx := txWithStatus.Transaction
	withdrawBlockHash := txWithStatus.TxStatus.BlockHash
	var depositBlockHash types.Hash
	for i := 0; i < len(withdrawTx.Inputs); i++ {
		outPoint := withdrawTx.Inputs[i].PreviousOutput
		txWithStatus, err := client.GetTransaction(context.Background(), outPoint.TxHash)
		if err != nil {
			return nil, err
		}
		tx := txWithStatus.Transaction
		index := outPoint.Index
		if IsDepositCell(tx.Outputs[index], tx.OutputsData[index]) {
			depositBlockHash = txWithStatus.TxStatus.BlockHash
			break
		}
	}
	if reflect.DeepEqual(depositBlockHash, types.Hash{}) {
		return nil, errors.New("can't find deposit cell")
	}
	depositBlockHeader, err := client.GetHeader(context.Background(), depositBlockHash)
	if err != nil {
		return nil, err
	}
	withdrawBlockHeader, err := client.GetHeader(context.Background(), withdrawBlockHash)
	if err != nil {
		return nil, err
	}
	return &ClaimInfo{
		DepositBlockHeader:  depositBlockHeader,
		WithdrawBlockHeader: withdrawBlockHeader,
		WithdrawOutpoint:    withdrawOutpoint,
	}, nil
}

func (r *ClaimInfo) CalculateDaoMinimumSince() uint64 {
	return calculateDaoMinimumSince(r.DepositBlockHeader, r.WithdrawBlockHeader)
}

func calculateDaoMinimumSince(depositBlockHeader *types.Header, withdrawBlockHeader *types.Header) uint64 {
	depositEpoch := types.ParseEpoch(depositBlockHeader.Epoch)
	withdrawEpoch := types.ParseEpoch(withdrawBlockHeader.Epoch)

	// calculate since
	withdrawFraction := withdrawEpoch.Index * depositEpoch.Length
	depositFraction := depositEpoch.Index * withdrawEpoch.Length
	depositedEpochs := withdrawEpoch.Number - depositEpoch.Number
	if withdrawFraction > depositFraction {
		depositedEpochs += 1
	}
	lockEpochs := (depositedEpochs + (DaoLockPeriodEpochs - 1)) / DaoLockPeriodEpochs * DaoLockPeriodEpochs

	minimumSinceEpochNumber := depositEpoch.Number + lockEpochs
	minimumSinceEpochIndex := depositEpoch.Index
	minimumSinceEpochLength := depositEpoch.Length
	minimumSince := &types.EpochParams{
		Length: minimumSinceEpochLength,
		Index:  minimumSinceEpochIndex,
		Number: minimumSinceEpochNumber,
	}
	return minimumSince.Uint64()
}

type WithdrawInfo struct {
	DepositOutPoint    *types.OutPoint
	DepositBlockNumber uint64
	DepositBlockHash   types.Hash
}

func NewWithdrawInfo(client rpc.Client, depositOutPoint *types.OutPoint) (*WithdrawInfo, error) {
	txWithStatus, err := client.GetTransaction(context.Background(), depositOutPoint.TxHash)
	if err != nil {
		return nil, err
	}
	depositBlockHash := txWithStatus.TxStatus.BlockHash
	header, err := client.GetHeader(context.Background(), depositBlockHash)
	if err != nil {
		return nil, err
	}
	return &WithdrawInfo{
		DepositOutPoint:    depositOutPoint,
		DepositBlockNumber: header.Number,
		DepositBlockHash:   depositBlockHash,
	}, nil
}
