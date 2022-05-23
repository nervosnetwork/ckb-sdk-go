package address

import (
	"encoding/hex"
	"github.com/nervosnetwork/ckb-sdk-go/crypto/blake2b"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateSecp256k1MultisigScript(t *testing.T) {
	var publicKeys [][]byte

	key, err := hex.DecodeString("032edb83018b57ddeb9bcc7287c5cc5da57e6e0289d31c9e98cb361e88678d6288")
	if err != nil {
		assert.Error(t, err)
	}
	publicKeys = append(publicKeys, key)

	key, err = hex.DecodeString("033aeb3fdbfaac72e9e34c55884a401ee87115302c146dd9e314677d826375dc8f")
	if err != nil {
		assert.Error(t, err)
	}
	publicKeys = append(publicKeys, key)

	key, err = hex.DecodeString("029a685b8206550ea1b600e347f18fd6115bffe582089d3567bec7eba57d04df01")
	if err != nil {
		assert.Error(t, err)
	}
	publicKeys = append(publicKeys, key)

	script, _, err := GenerateSecp256k1MultisigScript(0, 2, publicKeys)
	if err != nil {
		assert.Error(t, err)
	}

	address, err := ConvertScriptToAddress(Testnet, script)
	if err != nil {
		assert.Error(t, err)
	}
	assert.Equal(t, "ckt1qpw9q60tppt7l3j7r09qcp7lxnp3vcanvgha8pmvsa3jplykxn32sq0sfnkgf0ph76pkzwld9ujzex4pkeuwnlsdc5tqu", address)
}

func TestGenerateSecp256k1MultisigScriptByHash(t *testing.T) {
	publicKeysHex := []string {
		"032edb83018b57ddeb9bcc7287c5cc5da57e6e0289d31c9e98cb361e88678d6288",
		"033aeb3fdbfaac72e9e34c55884a401ee87115302c146dd9e314677d826375dc8f",
		"029a685b8206550ea1b600e347f18fd6115bffe582089d3567bec7eba57d04df01",
	}
	var publicKeysHash [][]byte
	for _, publicKeyHex := range publicKeysHex {
		key, err := hex.DecodeString(publicKeyHex)
		if err != nil {
			assert.Error(t, err)
		}
		hash, err := blake2b.Blake160(key)
		if err != nil {
			assert.Error(t, err)
		}
		publicKeysHash = append(publicKeysHash, hash)
	}
	script, _, err := GenerateSecp256k1MultisigScript(0, 2, publicKeysHash)
	if err != nil {
		assert.Error(t, err)
	}

	address, err := ConvertScriptToAddress(Testnet, script)
	if err != nil {
		assert.Error(t, err)
	}
	assert.Equal(t, "ckt1qpw9q60tppt7l3j7r09qcp7lxnp3vcanvgha8pmvsa3jplykxn32sq0sfnkgf0ph76pkzwld9ujzex4pkeuwnlsdc5tqu", address)
}
