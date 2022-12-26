package systemscript

import (
	"github.com/nervosnetwork/ckb-sdk-go/v2/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetSystemScriptInfo(t *testing.T) {
	script := GetInfo(types.NetworkMain, Secp256k1Blake160SighashAll)
	assert.NotNil(t, script)
	script = GetInfo(types.NetworkMain, Secp256k1Blake160MultisigAll)
	assert.NotNil(t, script)
	script = GetInfo(types.NetworkTest, Secp256k1Blake160MultisigAll)
	assert.NotNil(t, script)
}
