package omnilock

import (
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/nervosnetwork/ckb-sdk-go/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOmnilockWitnessLockSerialize(t *testing.T) {
	witness := hexutil.MustDecode("0x690000001000000069000000690000005500000055000000100000005500000055000000410000003434ca813dc378de0146aac8e60431fb52114acb3cb639f2fb2a479e1f219223532540413a154f440e939ee888c29221c0e8d6fef39402cbeedb6155317b356200")

	witnessArgs, err := types.DeserializeWitnessArgs(witness)
	assert.NoError(t, err)
	omnilockWitnessLock, err := DeserializeOmnilockWitnessLock(witnessArgs.Lock)
	assert.NoError(t, err)

	expected := hexutil.MustDecode("0x55000000100000005500000055000000410000003434ca813dc378de0146aac8e60431fb52114acb3cb639f2fb2a479e1f219223532540413a154f440e939ee888c29221c0e8d6fef39402cbeedb6155317b356200")
	assert.Equal(t, expected, omnilockWitnessLock.Serialize())
}
