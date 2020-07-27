package bech32

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncode(t *testing.T) {
	bytes, err := hex.DecodeString("0004000b1e0f14121b090411031e121f0c08070716071e120f1016101b17080d1c1d0200")
	if err != nil {
		assert.Error(t, err)
	}
	address, err := Encode("ckb", bytes)
	assert.Equal(t, "ckb1qyqt705jmfy3r7jlvg88k87j0sksmhgduazqrr2qt2", address)
}

func TestDecode(t *testing.T) {
	hrp, decoded, err := Decode("ckb1qyqt705jmfy3r7jlvg88k87j0sksmhgduazqrr2qt2")
	if err != nil {
		assert.Error(t, err)
	}
	assert.Equal(t, "ckb", hrp)
	assert.Equal(t, "0004000b1e0f14121b090411031e121f0c08070716071e120f1016101b17080d1c1d0200", hex.EncodeToString(decoded))
}
