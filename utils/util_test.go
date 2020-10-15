package utils

import (
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

func TestGenerateSudtAmount(t *testing.T) {
	amount := big.NewInt(10000000)
	data := GenerateSudtAmount(amount)
	expectedData := []byte{0x80, 0x96, 0x98, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}

	assert.Equal(t, expectedData, data)
}

func TestParseSudtAmountReturnError(t *testing.T) {
	data := []byte{0x80, 0x96}
	_, err := ParseSudtAmount(data)

	assert.Error(t, err)
}

func TestParseSudtAmountWith16BytesData(t *testing.T) {
	data := []byte{0x80, 0xC3, 0xC9, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	amount, err := ParseSudtAmount(data)
	assert.NoError(t, err)
	assert.True(t, big.NewInt(30000000).Cmp(amount) == 0)
}