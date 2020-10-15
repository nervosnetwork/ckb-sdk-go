package utils

import (
	"errors"
	"math/big"
)

func ParseSudtAmount(outputData []byte) (*big.Int, error) {
	if len(outputData) < 16 {
		return nil, errors.New("invalid sUDT amount")
	}
	b := outputData[0:16]
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