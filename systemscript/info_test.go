package systemscript

import (
	"github.com/nervosnetwork/ckb-sdk-go/v2/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetSystemScriptInfo(t *testing.T) {
	script := GetInfo(types.NetworkMain, Secp256k1Blake160SighashAll)
	assert.NotNil(t, script)
	script = GetInfo(types.NetworkMain, Secp256k1Blake160MultisigAllLegacy)
	assert.NotNil(t, script)
	script = GetInfo(types.NetworkMain, Secp256k1Blake160MultisigAllV2)
	assert.NotNil(t, script)
	script = GetInfo(types.NetworkTest, Secp256k1Blake160MultisigAllLegacy)
	assert.NotNil(t, script)
	script = GetInfo(types.NetworkTest, Secp256k1Blake160MultisigAllV2)
	assert.NotNil(t, script)
}
