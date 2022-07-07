package indexer

import (
	"context"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/nervosnetwork/ckb-sdk-go/types"
	"github.com/stretchr/testify/assert"
	"runtime/debug"
	"testing"
)

var c, _ = Dial("https://testnet.ckb.dev/indexer")

func TestGetTip(t *testing.T) {
	resp, err := c.GetTip(context.Background())
	checkError(t, err)
	assert.NotEqual(t, 0, resp.BlockNumber)
	assert.NotEqual(t, types.Hash{}, resp.BlockHash)
}

func TestGetCellsCapacity(t *testing.T) {
	s := &SearchKey{
		Script: &types.Script{
			CodeHash: types.HexToHash("0x9bd7e06f3ecf4be0f2fcd2188b23f1b9fcc88e5d4b65a8637b17723bbda3cce8"),
			HashType: types.HashTypeType,
			Args:     ethcommon.FromHex("0x4049ed9cec8a0d39c7a1e899f0dacb8a8c28ad14"),
		},
		ScriptType: ScriptTypeLock,
	}
	resp, err := c.GetCellsCapacity(context.Background(), s)
	checkError(t, err)
	assert.NotEqual(t, uint64(0x0), resp.BlockNumber)
	assert.NotEqual(t, uint64(0), resp.Capacity)
}

func TestGetCells(t *testing.T) {
	s := &SearchKey{
		Script: &types.Script{
			CodeHash: types.HexToHash("0x9bd7e06f3ecf4be0f2fcd2188b23f1b9fcc88e5d4b65a8637b17723bbda3cce8"),
			HashType: types.HashTypeType,
			Args:     ethcommon.FromHex("0x4049ed9cec8a0d39c7a1e899f0dacb8a8c28ad14"),
		},
		ScriptType: ScriptTypeLock,
	}
	resp, err := c.GetCells(context.Background(), s, SearchOrderAsc, 10, "")
	checkError(t, err)
	assert.Equal(t, 10, len(resp.Objects))
}

func TestGetTransactions(t *testing.T) {
	s := &SearchKey{
		Script: &types.Script{
			CodeHash: types.HexToHash("0x58c5f491aba6d61678b7cf7edf4910b1f5e00ec0cde2f42e0abb4fd9aff25a63"),
			HashType: types.HashTypeType,
			Args:     ethcommon.FromHex("0xe53f35ccf63bb37a3bb0ac3b7f89808077a78eae"),
		},
		ScriptType: ScriptTypeLock,
	}
	resp, err := c.GetTransactions(context.Background(), s, SearchOrderAsc, 10, "")
	checkError(t, err)
	assert.True(t, len(resp.Objects) >= 1)
	assert.NotEqual(t, 0, resp.Objects[0].BlockNumber)
	assert.NotEqual(t, "", resp.Objects[0].IoType)
}

func checkError(t *testing.T, err error) {
	if err != nil {
		t.Error(err, string(debug.Stack()))
	}
}