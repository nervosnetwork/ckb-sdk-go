package utils

import (
	"context"
	"errors"
	"github.com/nervosnetwork/ckb-sdk-go/rpc"
	"github.com/nervosnetwork/ckb-sdk-go/types"
	"math/big"
	"strconv"
	"strings"
)

func ParseSudtAmount(outputData []byte) (*big.Int, error) {
	if len(outputData) == 0 {
		return big.NewInt(0), nil
	}
	tmpData := make([]byte, len(outputData))
	copy(tmpData, outputData)
	if len(tmpData) < 16 {
		return nil, errors.New("invalid sUDT amount")
	}
	b := tmpData[0:16]
	b = reverse(b)

	return big.NewInt(0).SetBytes(b), nil
}

func GenerateSudtAmount(amount *big.Int) []byte {
	b := amount.Bytes()
	b = reverse(b)
	if len(b) < 16 {
		for i := len(b); i < 16; i++ {
			b = append(b, 0)
		}
	}

	return b
}

func reverse(b []byte) []byte {
	for i := 0; i < len(b)/2; i++ {
		b[i], b[len(b)-i-1] = b[len(b)-i-1], b[i]
	}
	return b
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

// GetMaxMatureBlockNumber return max mature block number
func GetMaxMatureBlockNumber(client rpc.Client, ctx context.Context) (uint64, error) {
	var cellbaseMaturity *types.EpochParams
	cellbaseMaturity, err := getCellbaseMaturity(client, ctx, cellbaseMaturity)
	if err != nil {
		return 0, err
	}
	tipHeader, err := client.GetTipHeader(ctx)
	if err != nil {
		return 0, err
	}
	tipEpoch := types.ParseEpoch(tipHeader.Epoch)
	if tipEpoch.Number < cellbaseMaturity.Number {
		return 0, errors.New("there are no mature live cells")
	} else {
		maxMatureEpoch, err := client.GetEpochByNumber(ctx, tipEpoch.Number-cellbaseMaturity.Number)
		if err != nil {
			return 0, err
		}
		number, err := calcMaxMatureBlockNumber(tipEpoch, maxMatureEpoch.StartNumber, maxMatureEpoch.Length, cellbaseMaturity)
		if err != nil {
			return 0, err
		}
		return number, nil
	}
}

func getCellbaseMaturity(client rpc.Client, ctx context.Context, cellbaseMaturity *types.EpochParams) (*types.EpochParams, error) {
	major, minor, _, err := ParseNodeVersion(client, ctx)
	if err != nil {
		return nil, err
	}
	if major > 0 || minor >= 39 {
		consensus, err := client.GetConsensus(ctx)
		if err != nil {
			return nil, err
		}
		cellbaseMaturity = types.ParseEpoch(consensus.CellbaseMaturity)
	} else {
		cellbaseMaturity = &types.EpochParams{
			Length: 1,
			Index:  0,
			Number: 4,
		}
	}
	return cellbaseMaturity, nil
}

// startNumber is maxMatureEpoch.StartNumber, length is maxMatureEpoch.Length
func calcMaxMatureBlockNumber(tipEpoch *types.EpochParams, startNumber, length uint64, cellbaseMaturity *types.EpochParams) (uint64, error) {
	a := isTipEpochLessThanCellbaseMaturity(tipEpoch, cellbaseMaturity)
	if a {
		return 0, nil
	} else {
		tipEpochRN, cellbaseMaturityRN := epochRationalNum(tipEpoch, cellbaseMaturity)
		epochDelta := rationalNumber{
			numer: tipEpochRN.numer - cellbaseMaturityRN.numer,
			denom: tipEpochRN.denom,
		}
		iNum := epochDelta.numer / epochDelta.denom
		index := (epochDelta.numer - iNum*epochDelta.denom) * length / epochDelta.denom
		blockNumber := index + startNumber

		return blockNumber, nil
	}
}

type rationalNumber struct {
	// Numerator.
	numer uint64
	// Denominator.
	denom uint64
}

func epochRationalNum(tipEpoch, cellbaseMaturity *types.EpochParams) (rationalNumber, rationalNumber) {
	if tipEpoch.Length != cellbaseMaturity.Length {
		tipEpochMN := rationalNumber{
			numer: tipEpoch.Number*tipEpoch.Length + tipEpoch.Index,
			denom: tipEpoch.Length,
		}
		cellbaseMaturityMN := rationalNumber{
			numer: (cellbaseMaturity.Number*cellbaseMaturity.Length + cellbaseMaturity.Index) * tipEpoch.Length,
			denom: cellbaseMaturity.Length * cellbaseMaturity.Length * tipEpoch.Length,
		}
		return tipEpochMN, cellbaseMaturityMN
	} else {
		tipEpochMN := rationalNumber{
			numer: tipEpoch.Number*tipEpoch.Length + tipEpoch.Index,
			denom: tipEpoch.Length,
		}
		cellbaseMaturityMN := rationalNumber{
			numer: cellbaseMaturity.Number*cellbaseMaturity.Length + cellbaseMaturity.Index,
			denom: cellbaseMaturity.Length,
		}
		return tipEpochMN, cellbaseMaturityMN
	}
}

func isTipEpochLessThanCellbaseMaturity(tipEpoch *types.EpochParams, cellbaseMaturity *types.EpochParams) bool {
	tipEpochRN, cellbaseMaturityRN := epochRationalNum(tipEpoch, cellbaseMaturity)
	if tipEpochRN.numer-cellbaseMaturityRN.numer < 0 {
		return true
	}
	return false
}

// ParseNodeVersion return ckb node version number
func ParseNodeVersion(client rpc.Client, ctx context.Context) (int, int, int, error) {
	nodeInfo, err := client.LocalNodeInfo(ctx)
	if err != nil {
		return 0, 0, 0, err
	}
	nodeVersion := strings.Split(nodeInfo.Version, " (")
	versionArr := strings.Split(nodeVersion[0], ".")
	major, err := strconv.Atoi(versionArr[0])
	minor, err := strconv.Atoi(versionArr[1])
	patch, err := strconv.Atoi(versionArr[2])
	return major, minor, patch, nil
}
