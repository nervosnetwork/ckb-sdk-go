package lightclient

import (
	"context"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/nervosnetwork/ckb-sdk-go/v2/indexer"
	"github.com/nervosnetwork/ckb-sdk-go/v2/types"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var c, _ = DialContext(context.Background(), "http://localhost:9000")
var ctx = context.Background()

var scriptForTest = &types.Script{
	CodeHash: types.HexToHash("0x9bd7e06f3ecf4be0f2fcd2188b23f1b9fcc88e5d4b65a8637b17723bbda3cce8"),
	HashType: types.HashTypeType,
	Args:     ethcommon.FromHex("0x4049ed9cec8a0d39c7a1e899f0dacb8a8c28ad14"),
}

// Remove when we have light client node
func skipCI(t *testing.T) {
	if os.Getenv("CI") != "" {
		t.Skip("Skipping as we do not have light client node")
	}
}

func TestSetScripts(t *testing.T) {
	scriptDetail := ScriptDetail{
		// ckt1qzda0cr08m85hc8jlnfp3zer7xulejywt49kt2rr0vthywaa50xwsq2qf8keemy2p5uu0g0gn8cd4ju23s5269qk8rg4r
		Script:      scriptForTest,
		ScriptType:  types.ScriptTypeLock,
		BlockNumber: 7033100,
	}
	err := c.SetScripts(context.Background(), []*ScriptDetail{&scriptDetail})
	assert.NoError(t, err)
}

func TestGetScripts(t *testing.T) {
	scriptDetails, err := c.GetScripts(ctx)
	assert.NoError(t, err)
	assert.NotEmpty(t, scriptDetails)
	assert.NotEmpty(t, scriptDetails[0].Script)
	assert.NotEmpty(t, scriptDetails[0].ScriptType)
}

func TestTipHeader(t *testing.T) {
	skipCI(t)
	header, err := c.GetTipHeader(ctx)
	assert.NoError(t, err)
	assert.NotEmpty(t, header)
}

func TestGetGenesisBlock(t *testing.T) {
	skipCI(t)
	block, err := c.GetGenesisBlock(ctx)
	assert.NoError(t, err)
	assert.NotEmpty(t, block)
	assert.NotEmpty(t, block.Transactions)
	assert.NotEmpty(t, block.Header)
}

func TestGetHeader(t *testing.T) {
	skipCI(t)
	header, err := c.GetHeader(ctx,
		types.HexToHash("0x86487ca41db5141bb750a0b5dbea8b87c0b3a05dda1c1e587ca9f7ccae3b4ad5"))
	assert.NoError(t, err)
	assert.NotEmpty(t, header)
}

func TestGetTransaction(t *testing.T) {
	skipCI(t)
	txWitHeader, err := c.GetTransaction(ctx,
		types.HexToHash("0x151d4d450c9e3bccf4b47d1ba6942d4e9c8c0eeeb7b9f708df827c164f035aa8"))
	assert.NoError(t, err)
	assert.NotEmpty(t, txWitHeader.Transaction)
	assert.NotEmpty(t, txWitHeader.TxStatus)
}

func TestFetchHeader(t *testing.T) {
	skipCI(t)
	fetchedHeader, err := c.FetchHeader(ctx,
		types.HexToHash("0xcb5eae958e3ea24b0486a393133aa33d51224ffaab3c4819350095b3446e4f70"))
	assert.NoError(t, err)
	assert.NotEmpty(t, fetchedHeader.Status)
	assert.NotEmpty(t, *fetchedHeader.Data)
}

func TestFetchTransaction(t *testing.T) {
	skipCI(t)
	fetchedTransaction, err := c.FetchTransaction(ctx,
		types.HexToHash("0x716e211698d3d9499aae7903867c744b67b539beeceddad330e73d1b6b617aef"))
	assert.NoError(t, err)
	assert.NotEmpty(t, fetchedTransaction.Status)
}

func TestGetCells(t *testing.T) {
	skipCI(t)
	s := &indexer.SearchKey{
		Script:     scriptForTest,
		ScriptType: types.ScriptTypeLock,
	}
	resp, err := c.GetCells(context.Background(), s, indexer.SearchOrderAsc, 10, "")
	assert.NoError(t, err)
	assert.NotEmpty(t, resp)
	assert.NotEmpty(t, resp.Objects[0].BlockNumber)
	assert.NotEmpty(t, resp.Objects[0].OutPoint)
	assert.NotEmpty(t, resp.Objects[0].Output)
}

func TestGetTransactions(t *testing.T) {
	skipCI(t)
	s := &indexer.SearchKey{
		Script:     scriptForTest,
		ScriptType: types.ScriptTypeLock,
	}
	resp, err := c.GetTransactions(context.Background(), s, indexer.SearchOrderAsc, 10, "")
	assert.NoError(t, err)
	assert.NotEmpty(t, resp)
	assert.NotEmpty(t, resp.Objects[0].BlockNumber)
	assert.NotEmpty(t, resp.Objects[0].IoType)
	assert.NotEmpty(t, resp.Objects[0].Transaction)
}

func TestGetTransactionsGrouped(t *testing.T) {
	skipCI(t)
	s := &indexer.SearchKey{
		Script:     scriptForTest,
		ScriptType: types.ScriptTypeLock,
	}
	resp, err := c.GetTransactionsGrouped(context.Background(), s, indexer.SearchOrderAsc, 10, "")
	assert.NoError(t, err)
	assert.NotEqual(t, 0, len(resp.Objects))
	assert.NotEqual(t, 0, resp.Objects[0].BlockNumber)
	assert.NotEqual(t, 0, resp.Objects[0].Transaction)
	assert.NotEmpty(t, resp.Objects[0].Cells[0])
	assert.NotEmpty(t, resp.Objects[0].Cells[0].IoType)
}

func TestGetCellsCapacity(t *testing.T) {
	skipCI(t)
	s := &indexer.SearchKey{
		Script:     scriptForTest,
		ScriptType: types.ScriptTypeLock,
	}
	resp, err := c.GetCellsCapacity(context.Background(), s)
	assert.NoError(t, err)
	assert.NotEmpty(t, resp.BlockNumber)
	assert.NotEmpty(t, resp.BlockHash)
	assert.NotEmpty(t, resp.Capacity)
}

func TestGetPeers(t *testing.T) {
	peers, err := c.GetPeers(ctx)
	if err != nil {
		t.Fatal(err)
	}
	assert.True(t, len(peers) > 0)
	assert.True(t, len(peers[0].Addresses) > 0)
	assert.True(t, len(peers[0].Protocols) > 0)
}

func TestClient_LocalNodeInfo(t *testing.T) {
	nodeInfo, err := c.LocalNodeInfo(ctx)
	if err != nil {
		t.Fatal(err)
	}
	assert.True(t, len(nodeInfo.Addresses) > 0)
	assert.True(t, len(nodeInfo.Protocols) > 0)
	assert.True(t, len(nodeInfo.Protocols[0].SupportVersions) > 0)
}
