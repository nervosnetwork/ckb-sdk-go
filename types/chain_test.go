package types

import (
	"encoding/hex"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestScriptOccupiedCapacity(t *testing.T) {
	args, _ := hex.DecodeString("3954acece65096bfa81258983ddb83915fc56bd8")
	s := &Script{
		CodeHash: HexToHash("0x9bd7e06f3ecf4be0f2fcd2188b23f1b9fcc88e5d4b65a8637b17723bbda3cce8"),
		HashType: ScriptHashType("type"),
		Args:     args,
	}
	expectedCapacity := uint64(len(s.CodeHash.Bytes())) + 1 + uint64(len(args))

	assert.Equal(t, expectedCapacity, s.OccupiedCapacity())
}

func TestCellOutputOccupiedCapacityOnlyLock(t *testing.T) {
	args, _ := hex.DecodeString("3954acece65096bfa81258983ddb83915fc56bd8")
	s := &Script{
		CodeHash: HexToHash("0x9bd7e06f3ecf4be0f2fcd2188b23f1b9fcc88e5d4b65a8637b17723bbda3cce8"),
		HashType: ScriptHashType("type"),
		Args:     args,
	}
	o := CellOutput{
		Capacity: 100000000000,
		Lock:     s,
		Type:     nil,
	}
	expectedCapacity := 8 + s.OccupiedCapacity()

	assert.Equal(t, expectedCapacity, o.OccupiedCapacity([]byte{}))
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
	expectedCapacity := 8 + s.OccupiedCapacity() + ts.OccupiedCapacity() + uint64(len(data))

	assert.Equal(t, expectedCapacity, o.OccupiedCapacity(data))
}
