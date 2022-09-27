package systemscript

import (
	"github.com/nervosnetwork/ckb-sdk-go/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetSystemScriptInfo(t *testing.T) {
	s := SystemScriptInfo(types.NetworkMain, SystemScriptSecp256k1Blake160SighashAll)
	assert.NotNil(t, s)
	s = SystemScriptInfo(types.NetworkMain, SystemScriptSecp256k1Blake160MultisigAll)
	assert.NotNil(t, s)
	s = SystemScriptInfo(types.NetworkTest, SystemScriptSecp256k1Blake160MultisigAll)
	assert.NotNil(t, s)
}
