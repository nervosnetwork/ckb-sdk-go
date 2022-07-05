package signer

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/nervosnetwork/ckb-sdk-go/crypto/secp256k1"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsMatch(t *testing.T) {
	key, _ := secp256k1.HexToKey("9d8ca87d75d150692211fa62b0d30de4d1ee6c530d5678b40b8cedacf0750d0f")
	args := common.FromHex("af0b41c627807fbddcee75afa174d5a7e5135ebd")
	actual, err := IsMatch(key, args)
	assert.Equal(t, true, actual)
	assert.Nil(t, nil, err)

	key, _ = secp256k1.HexToKey("9d8ca87d75d150692211fa62b0d30de4d1ee6c530d5678b40b8cedacf0750d0f")
	args = common.FromHex("0450340178ae277261a838c89f9ccb76a190ed4b")
	actual, err = IsMatch(key, args)
	assert.Equal(t, false, actual)
	assert.Nil(t, err)

	actual, err = IsMatch(nil, args)
	assert.Equal(t, false, actual)
	assert.NotNil(t, err)

	actual, err = IsMatch(key, nil)
	assert.Equal(t, false, actual)
	assert.NotNil(t, err)
}
