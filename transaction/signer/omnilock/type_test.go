package omnilock

import (
	"github.com/nervosnetwork/ckb-sdk-go/v2/address"
	"github.com/nervosnetwork/ckb-sdk-go/v2/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewOmnilockArgsFromAddress(t *testing.T) {
	_, err := NewOmnilockArgsFromAddress("ckt1qrejnmlar3r452tcg57gvq8patctcgy8acync0hxfnyka35ywafvkqgqgpy7m88v3gxnn3apazvlpkkt32xz3tg5qq3kzjf3")
	assert.NoError(t, err)

	// test invalid hash type
	a, _ := address.Decode("ckt1qrejnmlar3r452tcg57gvq8patctcgy8acync0hxfnyka35ywafvkqgqgpy7m88v3gxnn3apazvlpkkt32xz3tg5qq3kzjf3")
	a.Script.HashType = types.HashTypeData
	encoded, _ := a.Encode()
	_, err = NewOmnilockArgsFromAddress(encoded)
	assert.Error(t, err)

	// test invalid code hash
	_, err = NewOmnilockArgsFromAddress("ckt1qzda0cr08m85hc8jlnfp3zer7xulejywt49kt2rr0vthywaa50xwsq2qf8keemy2p5uu0g0gn8cd4ju23s5269qk8rg4r")
	assert.Error(t, err)
}
