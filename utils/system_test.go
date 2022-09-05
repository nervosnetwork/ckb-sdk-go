package utils

import (
	"github.com/nervosnetwork/ckb-sdk-go/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetSystemScriptInfo(t *testing.T) {
	s := GetSystemScriptInfo(types.NetworkMain, types.BuiltinScriptSecp256k1Blake160SighashAll)
	assert.NotNil(t, s)
	s = GetSystemScriptInfo(types.NetworkMain, types.BuiltinScriptSecp256k1Blake160MultisigAll)
	assert.NotNil(t, s)
}
