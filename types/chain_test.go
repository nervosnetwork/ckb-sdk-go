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
	assert.Equal(t, uint(0x2), v.Index)
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
	assert.Equal(t, uint64(0x4fe230), v.Since)
	assert.NotNil(t, v.PreviousOutput)

	jsonText2, _ := json.Marshal(v)
	AssertJsonEqual(t, jsonText1, jsonText2)
}

func TestJsonCellOutput(t *testing.T) {
	jsonText1 := []byte(`
{
    "capacity": "0x9502f9000",
    "lock": {
        "args": "0xa897829e60ee4e3fb0e4abe65549ec4a5ddafad7",
        "code_hash": "0x9bd7e06f3ecf4be0f2fcd2188b23f1b9fcc88e5d4b65a8637b17723bbda3cce8",
        "hash_type": "type"
    },
    "type": {
        "args": "0x02",
        "code_hash": "0x554cff969f3148e3c620749384004e9692e67c429f621554d139b505a281c7b8",
        "hash_type": "type"
    }
}`)
	var v CellOutput
	json.Unmarshal(jsonText1, &v)
	assert.Equal(t, uint64(0x9502f9000), v.Capacity)
	assert.Equal(t, HexToHash("0x9bd7e06f3ecf4be0f2fcd2188b23f1b9fcc88e5d4b65a8637b17723bbda3cce8"), v.Lock.CodeHash)
	assert.Equal(t, HexToHash("0x554cff969f3148e3c620749384004e9692e67c429f621554d139b505a281c7b8"), v.Type.CodeHash)

	jsonText2, _ := json.Marshal(v)
	AssertJsonEqual(t, jsonText1, jsonText2)
}

func TestJsonCellDep(t *testing.T) {
	jsonText1 := []byte(`
{
    "dep_type": "code",
    "out_point": {
        "index": "0x2",
        "tx_hash": "0x8f8c79eb6671709633fe6a46de93c0fedc9c1b8a6527a18d3983879542635c9f"
    }
}`)
	var v CellDep
	json.Unmarshal(jsonText1, &v)
	assert.Equal(t, DepTypeCode, v.DepType)
	assert.NotNil(t, v.OutPoint)

	jsonText2, _ := json.Marshal(v)
	AssertJsonEqual(t, jsonText1, jsonText2)
}

func TestJsonTransaction(t *testing.T) {
	jsonText1 := []byte(`
{
    "cell_deps": [
        {
            "dep_type": "dep_group",
            "out_point": {
                "index": "0x1",
                "tx_hash": "0xf8de3bb47d055cdf460d93a2a6e1b05f7432f9777c8c474abf4eec1d4aee5d37"
            }
        }
    ],
    "hash": "0xb19806b3ccc091a19d929d0237e6dc6e9b128a468b5b33c121c1bc59ad87877a",
    "header_deps": [],
    "inputs": [
        {
            "previous_output": {
                "index": "0x0",
                "tx_hash": "0x0dff101e716d77507bddc5ca189dc24c80e0fb8c269775b988b3cdd64e4f3395"
            },
            "since": "0x0"
        }
    ],
    "outputs": [
        {
            "capacity": "0xbaa315500",
            "lock": {
                "args": "0x4049ed9cec8a0d39c7a1e899f0dacb8a8c28ad14",
                "code_hash": "0x9bd7e06f3ecf4be0f2fcd2188b23f1b9fcc88e5d4b65a8637b17723bbda3cce8",
                "hash_type": "type"
            },
            "type": null
        },
        {
            "capacity": "0xdd2a73b8bf",
            "lock": {
                "args": "0xbc9818d8a149cfc0cd0323386c46ba07920a037f",
                "code_hash": "0x5c5069eb0857efc65e1bca0c07df34c31663b3622fd3876c876320fc9634e2a8",
                "hash_type": "type"
            },
            "type": null
        }
    ],
    "outputs_data": ["0x", "0x"],
    "version": "0x0",
    "witnesses": [
        "0xc200000010000000c2000000c2000000ae000000000002027336b0ba900684cb3cb00f0d46d4f64c0994a5625724c1e3925a5206944d753a6f3edaedf977d77f75ef2bf584ab0f400063964d5cddb3443fb5f11cbf00eedd76c64205f6c2d2ce342582871a010af6560bc6f559222852ffc44d3c9db9ae76092d843a05e39c0000ae2adec03512e320c2f0c087ec1d366c5fb43f7862fd1a7693284d356fbf56196e8f8ccd5cabe21bf3f0b2763d0c4f02c79af0d9993572eb3b752b09b08b6b1f00"
    ]
}`)
	var v Transaction
	json.Unmarshal(jsonText1, &v)

	jsonText2, _ := json.Marshal(v)
	AssertJsonEqual(t, jsonText1, jsonText2)
}

func TestJsonEpoch(t *testing.T) {
	jsonText1 := []byte(`
{
    "compact_target": "0x1d5f396f",
    "length": "0x356",
    "number": "0x100",
    "start_number": "0x2b445"
}`)
	var v Epoch
	json.Unmarshal(jsonText1, &v)
	assert.Equal(t, uint64(0x1d5f396f), v.CompactTarget)
	assert.Equal(t, uint64(0x356), v.Length)
	assert.Equal(t, uint64(0x100), v.Number)
	assert.Equal(t, uint64(0x2b445), v.StartNumber)

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
