package utils

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/nervosnetwork/ckb-sdk-go/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChequeCellArgs(t *testing.T) {
	senderLock := &types.Script{
		CodeHash: types.HexToHash("0x9bd7e06f3ecf4be0f2fcd2188b23f1b9fcc88e5d4b65a8637b17723bbda3cce8"),
		HashType: "type",
		Args:     common.FromHex("0xedcda9513fa030ce4308e29245a22c022d0443bb"),
	}
	receiverLock := &types.Script{
		CodeHash: types.HexToHash("0x9bd7e06f3ecf4be0f2fcd2188b23f1b9fcc88e5d4b65a8637b17723bbda3cce8"),
		HashType: "type",
		Args:     common.FromHex("0xedcda9513fa030ce4308e29245a22c022d0443bb"),
	}

	senderLockHash, err := senderLock.Hash()
	if err != nil {
		t.Fatal(err)
	}
	receiverLockHash, err := receiverLock.Hash()
	if err != nil {
		t.Fatal(err)
	}
	expectedArgs := append(receiverLockHash.Bytes()[0:20], senderLockHash.Bytes()[0:20]...)
	actualArgs, err := ChequeCellArgs(senderLock, receiverLock)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, expectedArgs, actualArgs)
}
