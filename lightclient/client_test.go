package lightclient

import (
	"context"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/nervosnetwork/ckb-sdk-go/indexer"
	"github.com/nervosnetwork/ckb-sdk-go/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

var c, _ = DialContext(context.Background(), "http://localhost:9000")
var ctx = context.Background()

var scriptForTest = &types.Script{
	CodeHash: types.HexToHash("0x9bd7e06f3ecf4be0f2fcd2188b23f1b9fcc88e5d4b65a8637b17723bbda3cce8"),
	HashType: types.HashTypeType,
	Args:     ethcommon.FromHex("0x4049ed9cec8a0d39c7a1e899f0dacb8a8c28ad14"),
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
	header, err := c.GetTipHeader(ctx)
	assert.NoError(t, err)
	assert.NotEmpty(t, header)
}

func TestGetGenesisBlock(t *testing.T) {
	block, err := c.GetGenesisBlock(ctx)
	assert.NoError(t, err)
	assert.NotEmpty(t, block)
	assert.NotEmpty(t, block.Transactions)
	assert.NotEmpty(t, block.Header)
}

func TestGetHeader(t *testing.T) {
	header, err := c.GetHeader(ctx,
		types.HexToHash("0xc78c65185c14e1b02d6457a06b4678bab7e15f194f49a840319b57c67d20053c"))
	assert.NoError(t, err)
	assert.NotEmpty(t, header)
}

func TestGetTransaction(t *testing.T) {
	txWitHeader, err := c.GetTransaction(ctx,
		types.HexToHash("0x151d4d450c9e3bccf4b47d1ba6942d4e9c8c0eeeb7b9f708df827c164f035aa8"))
	assert.NoError(t, err)
	assert.NotEmpty(t, txWitHeader.Transaction)
	assert.NotEmpty(t, txWitHeader.Header)
}

func TestFetchHeader(t *testing.T) {
	fetchedHeader, err := c.FetchHeader(ctx,
		types.HexToHash("0xe3cca69c9d00064b3bd5fa49d3f5b47d0653fb241b1996e6968b0100f10b56e3"))
	assert.NoError(t, err)
	assert.NotEmpty(t, fetchedHeader.Status)
}

func TestFetchTransaction(t *testing.T) {
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
	resp, err := c.GetTransactions(context.Background(), s, indexer.SearchOrderAsc, 10, "")
	assert.NoError(t, err)
	assert.NotEmpty(t, resp)
	assert.NotEmpty(t, resp.Objects[0].BlockNumber)
	assert.NotEmpty(t, resp.Objects[0].IoType)
}

func TestGetCellsCapacity(t *testing.T) {
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
