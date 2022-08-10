package types

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeserializeWitnessArgs(t *testing.T) {
	witnessArgsBytes := common.FromHex("0x2700000010000000180000002000000004000000123400000400000056780000030000009a0000")

	witnessArgs, err := DeserializeWitnessArgs(witnessArgsBytes)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, []byte{0x12, 0x34, 0x00, 0x00}, witnessArgs.Lock)
	assert.Equal(t, []byte{0x56, 0x78, 0x00, 0x00}, witnessArgs.InputType)
	assert.Equal(t, []byte{0x9a, 0x00, 0x00}, witnessArgs.OutputType)

	witnessArgsBytes = common.FromHex("0x10000000100000001000000010000000")
	witnessArgs, err = DeserializeWitnessArgs(witnessArgsBytes)
	if err != nil {
		t.Error(err)
	}
	assert.Nil(t, witnessArgs.Lock)
	assert.Nil(t, witnessArgs.InputType)
	assert.Nil(t, witnessArgs.OutputType)
}
