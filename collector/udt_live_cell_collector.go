package collector

import (
	"bytes"
	"github.com/nervosnetwork/ckb-sdk-go/indexer"
	"github.com/nervosnetwork/ckb-sdk-go/transaction"
	"github.com/nervosnetwork/ckb-sdk-go/types"
	"github.com/nervosnetwork/ckb-sdk-go/utils"
	"github.com/pkg/errors"
	"math/big"
)

var DataPlaceHolder = make([]byte, 16)

type UDTLiveCellProcessor struct {
	Max                        *big.Int
	TypeScript                 *types.Script
	Tx                         *types.Transaction
	FeeRate                    uint64
	CkbChangeOutputIndex       *ChangeOutputIndex
	SUDTChangeOutputIndex      *ChangeOutputIndex
	shouldHaveSUDTChangeOutput bool
}

type ChangeOutputIndex struct {
	Value int
}

func outputsCapacity(outputs []*types.CellOutput) (totalCapacity uint64) {
	for _, output := range outputs {
		totalCapacity += output.Capacity
	}
	return
}

func isEnoughCapacity(rs *utils.LiveCellCollectResult, p *UDTLiveCellProcessor) (bool, error) {
	changeCapacity := rs.Capacity - outputsCapacity(p.Tx.Outputs)
	if changeCapacity > 0 {
		fee, err := transaction.CalculateTransactionFee(p.Tx, p.FeeRate)
		if err != nil {
			return false, err
		}
		changeCapacity -= fee
		if p.CkbChangeOutputIndex != nil {
			changeOutput := p.Tx.Outputs[p.CkbChangeOutputIndex.Value]
			changeOutputData := p.Tx.OutputsData[p.CkbChangeOutputIndex.Value]
			changeOutputCapacity := changeOutput.OccupiedCapacity(changeOutputData)
			if changeCapacity >= changeOutputCapacity {
				changeOutput.Capacity = changeCapacity
				return true, nil
			} else {
				return false, nil
			}
		} else {
			if changeCapacity > 0 {
				return false, errors.New("cannot find change output")
			} else {
				return true, nil
			}
		}
	} else {
		return false, nil
	}
}

func (p *UDTLiveCellProcessor) Process(liveCell *indexer.LiveCell, result *utils.LiveCellCollectResult) (bool, error) {
	if p.Max == nil {
		return false, nil
	}
	var totalAmount *big.Int
	if _, ok := result.Options["totalAmount"]; !ok {
		result.Options = make(map[string]interface{})
		zero := big.NewInt(0)
		result.Options["totalAmount"] = zero
		totalAmount = zero
	} else {
		totalAmount = result.Options["totalAmount"].(*big.Int)
	}
	if totalAmount.Cmp(p.Max) < 0 {
		cellType := liveCell.Output.Type
		if p.TypeScript != nil {
			if cellType != nil && !p.TypeScript.Equals(liveCell.Output.Type) {
				return false, nil
			}
		}
		amount, err := utils.ParseSudtAmount(liveCell.OutputData)
		if err != nil {
			return false, errors.WithMessage(err, "sudt amount parse error")
		}
		total, ok := result.Options["totalAmount"]
		if ok {
			totalAmount = big.NewInt(0).Add(total.(*big.Int), amount)
			result.Options["totalAmount"] = totalAmount
		} else {
			result.Options = make(map[string]interface{})
			result.Options["totalAmount"] = amount
		}
		if !p.shouldHaveSUDTChangeOutput {
			p.shouldHaveSUDTChangeOutput = true
		}
	}

	if p.shouldHaveSUDTChangeOutput && totalAmount.Cmp(p.Max) > 0 && bytes.Compare(p.Tx.OutputsData[p.SUDTChangeOutputIndex.Value], DataPlaceHolder) == 0 {
		p.Tx.OutputsData[p.SUDTChangeOutputIndex.Value] = utils.GenerateSudtAmount(big.NewInt(0).Sub(totalAmount, p.Max))
	}
	if p.shouldHaveSUDTChangeOutput && totalAmount.Cmp(p.Max) == 0 {
		p.shouldHaveSUDTChangeOutput = false
		p.Tx.Outputs = RemoveCellOutput(p.Tx.Outputs, p.SUDTChangeOutputIndex.Value)
		p.Tx.OutputsData = RemoveCellOutputData(p.Tx.OutputsData, p.SUDTChangeOutputIndex.Value)
	}
	result.Capacity = result.Capacity + liveCell.Output.Capacity
	result.LiveCells = append(result.LiveCells, liveCell)
	input := &types.CellInput{
		Since: 0,
		PreviousOutput: &types.OutPoint{
			TxHash: liveCell.OutPoint.TxHash,
			Index:  liveCell.OutPoint.Index,
		},
	}
	p.Tx.Inputs = append(p.Tx.Inputs, input)
	p.Tx.Witnesses = append(p.Tx.Witnesses, []byte{})
	if len(p.Tx.Witnesses[0]) == 0 {
		p.Tx.Witnesses[0] = transaction.EmptyWitnessArgPlaceholder
	}
	if totalAmount.Cmp(p.Max) >= 0 {
		ok, err := isEnoughCapacity(result, p)
		if err != nil {
			return false, err
		}
		if ok {
			return true, nil
		}
	}

	return false, nil
}

func RemoveCellOutput(cellOutputs []*types.CellOutput, index int) []*types.CellOutput {
	ret := make([]*types.CellOutput, 0)
	ret = append(ret, cellOutputs[:index]...)
	return append(ret, cellOutputs[index+1:]...)
}

func RemoveCellOutputData(cellOutputData [][]byte, index int) [][]byte {
	ret := make([][]byte, 0)
	ret = append(ret, cellOutputData[:index]...)
	return append(ret, cellOutputData[index+1:]...)
}

func NewUDTLiveCellProcessor(max *big.Int) *UDTLiveCellProcessor {
	return &UDTLiveCellProcessor{
		Max: max,
	}
}
