package types

import (
	"encoding/hex"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestScriptOccupiedCapacity(t *testing.T) {
	s := &Script{
		CodeHash: HexToHash("0x68d5438ac952d2f584abf879527946a537e82c7f3c1cbf6d8ebf9767437d8e88"),
		HashType: HashTypeType,
		Args:     hexutil.MustDecode("0x36c329ed630d6ce750712a477543672adab57f4c"),
	}
	assert.Equal(t, uint64(5300000000), s.OccupiedCapacity())
}

func TestCellOutputOccupiedCapacityOnlyLock(t *testing.T) {
	o := CellOutput{
		Capacity: 100000000000,
		Lock: &Script{
			CodeHash: HexToHash("0x68d5438ac952d2f584abf879527946a537e82c7f3c1cbf6d8ebf9767437d8e88"),
			HashType: HashTypeType,
			Args:     hexutil.MustDecode("0x59a27ef3ba84f061517d13f42cf44ed020610061"),
		},
		Type: nil,
	}
	assert.Equal(t, uint64(6100000000), o.OccupiedCapacity([]byte{}))
}

func TestCellOutputOccupiedCapacityWithLockTypeAndData(t *testing.T) {
	args, _ := hex.DecodeString("3954acece65096bfa81258983ddb83915fc56bd8")
	s := &Script{
		CodeHash: HexToHash("0x9bd7e06f3ecf4be0f2fcd2188b23f1b9fcc88e5d4b65a8637b17723bbda3cce8"),
		HashType: ScriptHashType("type"),
		Args:     args,
	}
	tArgs, _ := hex.DecodeString("32e555f3ff8e135cece1351a6a2971518392c1e30375c1e006ad0ce8eac07947")
	ts := &Script{
		CodeHash: HexToHash("0x9bd7e06f3ecf4be0f2fcd2188b23f1b9fcc88e5d4b65a8637b17723bbda3cce8"),
		HashType: ScriptHashType("type"),
		Args:     tArgs,
	}
	o := CellOutput{
		Capacity: 100000000000,
		Lock:     s,
		Type:     ts,
	}
	data, _ := hex.DecodeString("a0860100000000000000000000000000")
	assert.Equal(t, uint64(14200000000), o.OccupiedCapacity(data))
}
