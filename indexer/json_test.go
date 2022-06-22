package indexer

import (
	"encoding/json"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJsonLiveCell(t *testing.T) {
	jsonText := []byte(`
{
    "block_number": "0x55e6c8",
    "out_point": {
        "index": "0x0",
        "tx_hash": "0x287554d155a9b9e30a1a6fd9e5d9e41afee612b0c8996f0073afb7f2894025f9"
    },
    "output": {
        "capacity": "0xba43b7400",
        "lock": {
            "args": "0x4049ed9cec8a0d39c7a1e899f0dacb8a8c28ad14",
            "code_hash": "0x9bd7e06f3ecf4be0f2fcd2188b23f1b9fcc88e5d4b65a8637b17723bbda3cce8",
            "hash_type": "type"
        },
        "type": {
            "args": "0x",
            "code_hash": "0x82d76d1b75fe2fd9a27dfbaa65a039221a380d76c926f378d3f81cf3e7e13f2e",
            "hash_type": "type"
        }
    },
    "output_data": "0x0fe5550000000000",
    "tx_index": "0x1"
}`)
	var v LiveCell
	json.Unmarshal(jsonText, &v)
	assert.Equal(t, uint64(0x55e6c8), v.BlockNumber)
	assert.Equal(t, ethcommon.FromHex("0x0fe5550000000000"), v.OutputData)
	assert.Equal(t, uint(0x1), v.TxIndex)
	assert.NotNil(t, v.OutPoint)
	assert.NotNil(t, v.Output)
}
