package utils

import (
	"errors"
	"github.com/nervosnetwork/ckb-sdk-go/types"
	"math/big"
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
