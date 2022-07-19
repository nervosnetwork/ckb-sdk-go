package secp256k1

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPubKey(t *testing.T) {
	k, err := HexToKey("0xccb083b37aa346c5ce2e1f99a687a153baa04052f26db6ab3c26d6a4cc15c5f1")
	if err != nil {
		t.Error(err)
	}
	encoded := k.PubKeyUncompressed()
	assert.Equal(t, common.FromHex("0x04a0a7a7597b019828a1dda6ed52ab25181073ec3a9825d28b9abbb932fe1ec83dd117a8eef7649c25be5a591d08f80ffe7e9c14100ad1b58ac78afa606a576453"), encoded)
	encoded = k.PubKey()
	assert.Equal(t, common.FromHex("0x03a0a7a7597b019828a1dda6ed52ab25181073ec3a9825d28b9abbb932fe1ec83d"), encoded)
}