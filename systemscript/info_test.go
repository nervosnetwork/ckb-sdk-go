package systemscript

import (
	"github.com/nervosnetwork/ckb-sdk-go/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetSystemScriptInfo(t *testing.T) {
	s := GetInfo(types.NetworkMain, Secp256k1Blake160SighashAll)
	assert.NotNil(t, s)
	s = GetInfo(types.NetworkMain, Secp256k1Blake160MultisigAll)
	assert.NotNil(t, s)
	s = GetInfo(types.NetworkTest, Secp256k1Blake160MultisigAll)
	assert.NotNil(t, s)
}
