package indexer

import (
	"context"
	"encoding/json"
	"fmt"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/nervosnetwork/ckb-sdk-go/v2/types"
	"github.com/stretchr/testify/assert"
	"math"
	"runtime/debug"
	"testing"
)

var c, _ = Dial("https://testnet.ckb.dev/")

func TestGetTip(t *testing.T) {
	// TODO: fix all deprecated RPC caused tests
	t.Skip("Skipping testing")
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
		ScriptType: types.ScriptTypeLock,
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
		ScriptType: types.ScriptTypeLock,
	}
	resp, err := c.GetCells(context.Background(), s, SearchOrderAsc, 10, "")
	checkError(t, err)
	assert.Equal(t, 10, len(resp.Objects))

	// Check response when `WithData` == true in request
	s = &SearchKey{
		Script: &types.Script{
			// https://pudge.explorer.nervos.org/address/ckt1qzda0cr08m85hc8jlnfp3zer7xulejywt49kt2rr0vthywaa50xwsqgxc8z84suk20xzx8337sckkkjfqvzk2ysq48gzc
			CodeHash: types.HexToHash("0x9bd7e06f3ecf4be0f2fcd2188b23f1b9fcc88e5d4b65a8637b17723bbda3cce8"),
			HashType: types.HashTypeType,
			Args:     ethcommon.FromHex("0x06c1c47ac39653cc231e31f4316b5a4903056512"),
		},
		ScriptType: types.ScriptTypeLock,
		WithData:   true,
	}
	resp, err = c.GetCells(context.Background(), s, SearchOrderAsc, 10, "")
	checkError(t, err)
	assert.Equal(t, ethcommon.FromHex("0x0000000000000000"), resp.Objects[0].OutputData)
}

func TestGetCellsMaxLimit(t *testing.T) {
	s := &SearchKey{
		Script: &types.Script{
			CodeHash: types.HexToHash("0x9bd7e06f3ecf4be0f2fcd2188b23f1b9fcc88e5d4b65a8637b17723bbda3cce8"),
			HashType: types.HashTypeType,
			Args:     ethcommon.FromHex("0x4049ed9cec8a0d39c7a1e899f0dacb8a8c28ad14"),
		},
		ScriptType: types.ScriptTypeLock,
	}
	resp, err := c.GetCells(context.Background(), s, SearchOrderAsc, math.MaxUint32, "")
	checkError(t, err)
	assert.Equal(t, 34, len(resp.Objects))

	// Check response when `WithData` == true in request
	s = &SearchKey{
		Script: &types.Script{
			// https://pudge.explorer.nervos.org/address/ckt1qzda0cr08m85hc8jlnfp3zer7xulejywt49kt2rr0vthywaa50xwsqgxc8z84suk20xzx8337sckkkjfqvzk2ysq48gzc
			CodeHash: types.HexToHash("0x9bd7e06f3ecf4be0f2fcd2188b23f1b9fcc88e5d4b65a8637b17723bbda3cce8"),
			HashType: types.HashTypeType,
			Args:     ethcommon.FromHex("0x06c1c47ac39653cc231e31f4316b5a4903056512"),
		},
		ScriptType: types.ScriptTypeLock,
		WithData:   true,
	}
	resp, err = c.GetCells(context.Background(), s, SearchOrderAsc, 10, "")
	checkError(t, err)
	assert.Equal(t, ethcommon.FromHex("0x0000000000000000"), resp.Objects[0].OutputData)
}

func TestGetTransactions(t *testing.T) {
	s := &SearchKey{
		Script: &types.Script{
			CodeHash: types.HexToHash("0x58c5f491aba6d61678b7cf7edf4910b1f5e00ec0cde2f42e0abb4fd9aff25a63"),
			HashType: types.HashTypeType,
			Args:     ethcommon.FromHex("0xe53f35ccf63bb37a3bb0ac3b7f89808077a78eae"),
		},
		ScriptType: types.ScriptTypeLock,
	}
	resp, err := c.GetTransactions(context.Background(), s, SearchOrderAsc, 10, "")
	checkError(t, err)
	assert.True(t, len(resp.Objects) >= 1)
	assert.NotEqual(t, 0, resp.Objects[0].BlockNumber)
	assert.NotEqual(t, "", resp.Objects[0].IoType)
}

func TestGetTransactionsGrouped(t *testing.T) {
	s := &SearchKey{
		Script: &types.Script{
			CodeHash: types.HexToHash("0x58c5f491aba6d61678b7cf7edf4910b1f5e00ec0cde2f42e0abb4fd9aff25a63"),
			HashType: types.HashTypeType,
			Args:     ethcommon.FromHex("0xe53f35ccf63bb37a3bb0ac3b7f89808077a78eae"),
		},
		ScriptType: types.ScriptTypeLock,
	}
	resp, err := c.GetTransactionsGrouped(context.Background(), s, SearchOrderAsc, 10, "")
	checkError(t, err)
	assert.Equal(t, 10, len(resp.Objects))
	assert.NotEqual(t, 0, resp.Objects[0].BlockNumber)
	assert.NotEmpty(t, resp.Objects[0].Cells[0])
	assert.NotEmpty(t, resp.Objects[0].Cells[0].IoType)
	assert.NotEmpty(t, resp.Objects[0].Cells[0].IoIndex)
}

func TestFilter(t *testing.T) {
	lockScript := &types.Script{
		CodeHash: types.HexToHash("0x58c5f491aba6d61678b7cf7edf4910b1f5e00ec0cde2f42e0abb4fd9aff25a63"),
		HashType: types.HashTypeType,
		Args:     ethcommon.FromHex("0xe53f35ccf63bb37a3bb0ac3b7f89808077a78eae"),
	}
	filter := &Filter{
		Script: &types.Script{
			CodeHash: types.HexToHash("0x58c5f491aba6d61678b7cf7edf4910b1f5e00ec0cde2f42e0abb4fd9aff25a63"),
			HashType: types.HashTypeType,
			Args:     lockScript.Hash().Bytes()[:20],
		},
	}
	s := &SearchKey{
		Script:     lockScript,
		ScriptType: types.ScriptTypeLock,
		Filter:     filter,
	}
	jsonF, _ := json.Marshal(filter)
	fmt.Println(string(jsonF))
	_, err := c.GetTransactions(context.Background(), s, SearchOrderAsc, 1000, "")
	checkError(t, err)
}

func checkError(t *testing.T, err error) {
	if err != nil {
		t.Error(err, string(debug.Stack()))
	}
}
