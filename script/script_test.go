package script

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMultisigConfigEncode(t *testing.T) {
	m := &MultisigConfig{
		Version:    0,
		FirstN:     0,
		Threshold:  2,
		KeysHashes: getKeysHashes(),
	}
	encoded := m.Encode()
	assert.Equal(t, common.FromHex("0x000002029b41c025515b00c24e2e2042df7b221af5c1891fe732dcd15b7618eb1d7a11e6a68e4579b5be0114"), encoded)
	hash := m.Hash160()
	assert.Equal(t, common.FromHex("0x35ed7b939b4ac9cb447b82340fd8f26d344f7a62"), hash)
}

func TestMultisigConfigDecode(t *testing.T) {
	bytes := common.FromHex("0x000002029b41c025515b00c24e2e2042df7b221af5c1891fe732dcd15b7618eb1d7a11e6a68e4579b5be0114")
	m, err := DecodeToMultisigConfig(bytes)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, byte(0), m.FirstN)
	assert.Equal(t, byte(2), m.Threshold)
	assert.Equal(t, getKeysHashes(), m.KeysHashes)

	bytes = common.FromHex("0x000002039b41c025515b00c24e2e2042df7b221af5c1891fe732dcd15b7618eb1d7a11e6a68e4579b5be0114")
	_, err = DecodeToMultisigConfig(bytes)
	assert.Error(t, err)

	bytes = common.FromHex("0x000002029b41c025515b00c24e2e2042df7b221af5c1891f")
	_, err = DecodeToMultisigConfig(bytes)
	assert.Error(t, err)
}

func getKeysHashes() [][20]byte {
	keysHashes := make([][20]byte, 2)
	copy(keysHashes[0][:], common.FromHex("0x9b41c025515b00c24e2e2042df7b221af5c1891f"))
	copy(keysHashes[1][:], common.FromHex("0xe732dcd15b7618eb1d7a11e6a68e4579b5be0114"))
	return keysHashes
}
