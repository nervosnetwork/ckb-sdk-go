package lightclient

import (
	"context"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/nervosnetwork/ckb-sdk-go/v2/indexer"
	"github.com/nervosnetwork/ckb-sdk-go/v2/mocking"
	"github.com/nervosnetwork/ckb-sdk-go/v2/types"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var c, _ = DialMockContext(context.Background(), "http://localhost:9000") // We are using mocking client now, url is just a placeholder param and takes no effect to these tests
var ctx = context.Background()
var mockClient = interface{}(c.GetRawClient()).(*mocking.MockClient)
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
	mockClient.LoadMockingTestFromFile(t, "set_scripts", []*ScriptDetail{&scriptDetail})
	err := c.SetScripts(context.Background(), []*ScriptDetail{&scriptDetail})
	assert.NoError(t, err)
}

func TestGetScripts(t *testing.T) {
	mockClient.LoadMockingTestFromFile(t, "get_scripts")
	scriptDetails, err := c.GetScripts(ctx)
	assert.NoError(t, err)
	assert.NotEmpty(t, scriptDetails)
	assert.NotEmpty(t, scriptDetails[0].Script)
	assert.NotEmpty(t, scriptDetails[0].ScriptType)
}

func TestTipHeader(t *testing.T) {
	mockClient.LoadMockingTestFromFile(t, "get_tip_header")
	header, err := c.GetTipHeader(ctx)
	assert.NoError(t, err)
	assert.NotEmpty(t, header)
}

func TestGetGenesisBlock(t *testing.T) {
	mockClient.LoadMockingTestFromFile(t, "get_genesis_block")
	block, err := c.GetGenesisBlock(ctx)
	assert.NoError(t, err)
	assert.NotEmpty(t, block)
	assert.NotEmpty(t, block.Transactions)
	assert.NotEmpty(t, block.Header)
}

func TestGetHeader(t *testing.T) {
	mockClient.LoadMockingTestFromFile(t, "get_header", types.HexToHash("0x10639e0895502b5688a6be8cf69460d76541bfa4821629d86d62ba0aae3f9606"))
	header, err := c.GetHeader(ctx,
		types.HexToHash("0x10639e0895502b5688a6be8cf69460d76541bfa4821629d86d62ba0aae3f9606"))
	assert.NoError(t, err)
	assert.NotEmpty(t, header)
}

func TestGetTransaction(t *testing.T) {
	mockClient.LoadMockingTestFromFile(t, "get_transaction", types.HexToHash("0x8f8c79eb6671709633fe6a46de93c0fedc9c1b8a6527a18d3983879542635c9f"))
	txWitHeader, err := c.GetTransaction(ctx,
		types.HexToHash("0x8f8c79eb6671709633fe6a46de93c0fedc9c1b8a6527a18d3983879542635c9f"))
	assert.NoError(t, err)
	assert.NotEmpty(t, txWitHeader.Transaction)
	assert.NotEmpty(t, txWitHeader.TxStatus)
}

func TestFetchHeader(t *testing.T) {
	mockClient.LoadMockingTestFromFile(t, "fetch_header", types.HexToHash("0xcb5eae958e3ea24b0486a393133aa33d51224ffaab3c4819350095b3446e4f70"))
	fetchedHeader, err := c.FetchHeader(ctx,
		types.HexToHash("0xcb5eae958e3ea24b0486a393133aa33d51224ffaab3c4819350095b3446e4f70"))
	assert.NoError(t, err)
	assert.NotEmpty(t, fetchedHeader.Status)
	assert.NotEmpty(t, *fetchedHeader.Data)
}

func TestFetchTransaction(t *testing.T) {
	mockClient.LoadMockingTestFromFile(t, "fetch_transaction", types.HexToHash("0x716e211698d3d9499aae7903867c744b67b539beeceddad330e73d1b6b617aef"))
	fetchedTransaction, err := c.FetchTransaction(ctx,
		types.HexToHash("0x716e211698d3d9499aae7903867c744b67b539beeceddad330e73d1b6b617aef"))
	assert.NoError(t, err)
	assert.NotEmpty(t, fetchedTransaction.Status)
}

func TestGetCells(t *testing.T) {
	s := &indexer.SearchKey{
		Script:     scriptForTest,
		ScriptType: types.ScriptTypeLock,
	}
	mockClient.LoadMockingTestFromFile(t, "get_cells", s, indexer.SearchOrderAsc, hexutil.Uint64(10)) // this is a special cast, make it same with the actual call
	resp, err := c.GetCells(context.Background(), s, indexer.SearchOrderAsc, 10, "")
	assert.NoError(t, err)
	assert.NotEmpty(t, resp)
	assert.NotEmpty(t, resp.Objects[0].BlockNumber)
	assert.NotEmpty(t, resp.Objects[0].OutPoint)
	assert.NotEmpty(t, resp.Objects[0].Output)
}

func TestGetTransactions(t *testing.T) {
	s := &indexer.SearchKey{
		Script:     scriptForTest,
		ScriptType: types.ScriptTypeLock,
	}
	mockClient.LoadMockingTestFromFile(t, "get_transactions", s, indexer.SearchOrderAsc, hexutil.Uint64(10)) // this is a special cast, make it same with the actual call
	resp, err := c.GetTransactions(context.Background(), s, indexer.SearchOrderAsc, 10, "")
	assert.NoError(t, err)
	assert.NotEmpty(t, resp)
	assert.NotEmpty(t, resp.Objects[0].BlockNumber)
	assert.NotEmpty(t, resp.Objects[0].IoType)
	assert.NotEmpty(t, resp.Objects[0].Transaction)
}

func TestGetTransactionsGrouped(t *testing.T) {
	s := &indexer.SearchKey{
		Script:     scriptForTest,
		ScriptType: types.ScriptTypeLock,
	}
	payload := &struct {
		indexer.SearchKey
		GroupByTransaction bool `json:"group_by_transaction"`
	}{
		SearchKey:          *s,
		GroupByTransaction: true,
	}
	mockClient.LoadMockingTestFromFilePatched(t, "get_transactions_grouped", "get_transactions", payload, indexer.SearchOrderAsc, hexutil.Uint64(10)) // this is a special cast, make it same with the actual call
	resp, err := c.GetTransactionsGrouped(context.Background(), s, indexer.SearchOrderAsc, 10, "")
	assert.NoError(t, err)
	assert.NotEqual(t, 0, len(resp.Objects))
	assert.NotEqual(t, 0, resp.Objects[0].BlockNumber)
	assert.NotEqual(t, 0, resp.Objects[0].Transaction)
	assert.NotEmpty(t, resp.Objects[0].Cells[0])
	assert.NotEmpty(t, resp.Objects[0].Cells[0].IoType)
}

func TestGetCellsCapacity(t *testing.T) {
	s := &indexer.SearchKey{
		Script:     scriptForTest,
		ScriptType: types.ScriptTypeLock,
	}
	mockClient.LoadMockingTestFromFile(t, "get_cells_capacity", s)
	resp, err := c.GetCellsCapacity(context.Background(), s)
	assert.NoError(t, err)
	assert.NotEmpty(t, resp.BlockNumber)
	assert.NotEmpty(t, resp.BlockHash)
	assert.NotEmpty(t, resp.Capacity)
}

func TestGetPeers(t *testing.T) {
	mockClient.LoadMockingTestFromFile(t, "get_peers")
	peers, err := c.GetPeers(ctx)
	if err != nil {
		t.Fatal(err)
	}
	assert.True(t, len(peers) > 0)
	assert.True(t, len(peers[0].Addresses) > 0)
	assert.True(t, len(peers[0].Protocols) > 0)
}

func TestClient_LocalNodeInfo(t *testing.T) {
	mockClient.LoadMockingTestFromFile(t, "local_node_info")
	nodeInfo, err := c.LocalNodeInfo(ctx)
	if err != nil {
		t.Fatal(err)
	}
	assert.True(t, len(nodeInfo.Addresses) > 0)
	assert.True(t, len(nodeInfo.Protocols) > 0)
	assert.True(t, len(nodeInfo.Protocols[0].SupportVersions) > 0)
}
