package types

import (
	"encoding/hex"
	"encoding/json"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJsonScript(t *testing.T) {
	jsonText1 := []byte(`
{
    "args": "0xa897829e60ee4e3fb0e4abe65549ec4a5ddafad7",
    "code_hash": "0x9bd7e06f3ecf4be0f2fcd2188b23f1b9fcc88e5d4b65a8637b17723bbda3cce8",
    "hash_type": "type"
}`)
	var v Script
	json.Unmarshal(jsonText1, &v)
	assert.Equal(t, ethcommon.FromHex("0xa897829e60ee4e3fb0e4abe65549ec4a5ddafad7"), v.Args)
	assert.Equal(t, HexToHash("0x9bd7e06f3ecf4be0f2fcd2188b23f1b9fcc88e5d4b65a8637b17723bbda3cce8"), v.CodeHash)
	assert.Equal(t, HashTypeType, v.HashType)

	jsonText2, _ := json.Marshal(v)
	AssertJsonEqual(t, jsonText1, jsonText2)
}

func TestJsonOutPoint(t *testing.T) {
	jsonText1 := []byte(`
{
    "index": "0x2",
    "tx_hash": "0xf8de3bb47d055cdf460d93a2a6e1b05f7432f9777c8c474abf4eec1d4aee5d37"
}`)
	var v OutPoint
	json.Unmarshal(jsonText1, &v)
	assert.Equal(t, uint(2), v.Index)
	assert.Equal(t, HexToHash("0xf8de3bb47d055cdf460d93a2a6e1b05f7432f9777c8c474abf4eec1d4aee5d37"), v.TxHash)

	jsonText2, _ := json.Marshal(v)
	AssertJsonEqual(t, jsonText1, jsonText2)
}

func TestJsonCellInput(t *testing.T) {
	jsonText1 := []byte(`
{
    "previous_output": {
        "index": "0xffffffff",
        "tx_hash": "0x0000000000000000000000000000000000000000000000000000000000000000"
    },
    "since": "0x4fe230"
}`)
	var v CellInput
	json.Unmarshal(jsonText1, &v)
	assert.Equal(t, uint64(5235248), v.Since)
	assert.NotNil(t, v.PreviousOutput)

	jsonText2, _ := json.Marshal(v)
	AssertJsonEqual(t, jsonText1, jsonText2)
}


func AssertJsonEqual(t *testing.T, t1, t2 []byte) {
	m1 := map[string]interface{}{}
	m2 := map[string]interface{}{}
	json.Unmarshal(t1, &m1)
	json.Unmarshal(t2, &m2)
	assert.Equal(t, m2, m1)
}

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
