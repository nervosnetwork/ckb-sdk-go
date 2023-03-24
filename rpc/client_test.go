package rpc

import (
	"context"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/nervosnetwork/ckb-sdk-go/v2/indexer"
	"github.com/nervosnetwork/ckb-sdk-go/v2/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

var testClient, _ = DialContext(context.Background(), "https://testnet.ckb.dev")
var ctx = context.Background()

func assertEqualHexBytes(t *testing.T, a string, b []byte) {
	a1, err := hexutil.Decode(a)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, a1, b)
}

func TestClient_GetBlockByNumber(t *testing.T) {
	block, err := testClient.GetBlockByNumber(ctx, 1)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 1, len(block.Transactions))
}

func TestClient_GetBlockByNumberWithCycles(t *testing.T) {
	block, err := testClient.GetBlockByNumberWithCycles(ctx, 1)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 1, len(block.Block.Transactions))
}

func TestClient_GetBlockHash(t *testing.T) {
	blockHash, err := testClient.GetBlockHash(ctx, 1)
	if err != nil {
		t.Fatal(err)
	}
	assertEqualHexBytes(t,
		"0xd5ac7cf8c34a975bf258a34f1c2507638487ab71aa4d10a9ec73704aa3abf9cd",
		blockHash.Bytes())
}

func TestClient_GetBlockEconomicState(t *testing.T) {
	blockEconomicState, err := testClient.GetBlockEconomicState(ctx,
		types.HexToHash("0xd5ac7cf8c34a975bf258a34f1c2507638487ab71aa4d10a9ec73704aa3abf9cd"))
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, uint64(9207601095), blockEconomicState.MinerReward.Secondary)
}

func TestClient_GetBlock(t *testing.T) {
	block, err := testClient.GetBlock(ctx,
		types.HexToHash("0xd5ac7cf8c34a975bf258a34f1c2507638487ab71aa4d10a9ec73704aa3abf9cd"))
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 1, len(block.Transactions))
	assert.NotNil(t, block.Header)
}

func TestClient_GetBlockVerbosity0(t *testing.T) {
	block, err := testClient.GetPackedBlock(ctx,
		types.HexToHash("0xd5ac7cf8c34a975bf258a34f1c2507638487ab71aa4d10a9ec73704aa3abf9cd"))
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 1, len(block.Transactions))
	assert.NotNil(t, block.Header)

	// verify equivalent
	block2, err := testClient.GetBlock(ctx,
		types.HexToHash("0xd5ac7cf8c34a975bf258a34f1c2507638487ab71aa4d10a9ec73704aa3abf9cd"))
	if err != nil {
		t.Fatal(err)
	}
	b1str := block.Header.Hash.String()
	b2str := block2.Header.Hash.String()
	assert.Equal(t, b1str, b2str)
}

func TestClient_GetBlockWithCycles(t *testing.T) {
	block, err := testClient.GetBlockWithCycles(ctx,
		types.HexToHash("0xd5ac7cf8c34a975bf258a34f1c2507638487ab71aa4d10a9ec73704aa3abf9cd"))
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 1, len(block.Block.Transactions))
	assert.NotNil(t, block.Block.Header)
}

func TestClient_GetTransaction(t *testing.T) {
	txView, err := testClient.GetTransaction(ctx,
		types.HexToHash("0x8277d74d33850581f8d843613ded0c2a1722dec0e87e748f45c115dfb14210f1"))
	assert.NoError(t, err)
	tx := txView.Transaction
	status := txView.TxStatus
	assert.Equal(t, 4, len(tx.CellDeps))
	assert.Equal(t, 1, len(tx.Inputs))
	assert.Equal(t, 3, len(tx.Outputs))
	assert.Equal(t, uint64(30000000000), tx.Outputs[0].Capacity)
	assert.Equal(t, types.TransactionStatusCommitted, status.Status)
	assert.NotNil(t, status.BlockHash)

	// NOTE: test commented because rejected tx will be removed after a expiry time by ckb
	// TODO: Adding mock RPC returns to make unit test more standalone
	//txView, err = testClient.GetTransaction(ctx,
	//	types.HexToHash("0xb2b8911aeac92de53fc3edc218cf979ae4752a7a67e698b0b1726db53126f31f"))
	//assert.NoError(t, err)
	//tx = txView.Transaction
	//status = txView.TxStatus
	//assert.Nil(t, tx)
	//assert.Equal(t, types.TransactionStatusRejected, status.Status)
	//assert.NotNil(t, status.Reason)
	//assert.Nil(t, status.Block)
}

func TestClient_GetTipHeader(t *testing.T) {
	header, err := testClient.GetTipHeader(ctx)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotEqual(t, uint64(0), header.Number)
	assert.NotEqual(t, uint(0), header.CompactTarget)
}

func TestClient_GetTipBlockNumber(t *testing.T) {
	blockNumber, err := testClient.GetTipBlockNumber(ctx)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotEqual(t, uint64(0), blockNumber)
}

func TestClient_GetCurrentEpoch(t *testing.T) {
	epoch, err := testClient.GetCurrentEpoch(ctx)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotEqual(t, uint64(0), epoch.Number)
	assert.NotEqual(t, uint64(0), epoch.CompactTarget)
}

func TestClient_GetEpochByNumber(t *testing.T) {
	epoch, err := testClient.GetEpochByNumber(ctx, 2)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, uint64(1500), epoch.StartNumber)
	assert.Equal(t, uint64(500945247), epoch.CompactTarget)
}

func TestClient_GetHeader(t *testing.T) {
	header, err := testClient.GetHeader(ctx,
		types.HexToHash("0xd5ac7cf8c34a975bf258a34f1c2507638487ab71aa4d10a9ec73704aa3abf9cd"))
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, uint64(1), header.Number)
	assert.Equal(t, uint64(1590137711584), header.Timestamp)
}

func TestClient_GetHeaderVerbosity0(t *testing.T) {
	header, err := testClient.GetPackedHeader(ctx,
		types.HexToHash("0xd5ac7cf8c34a975bf258a34f1c2507638487ab71aa4d10a9ec73704aa3abf9cd"))
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, uint64(1), header.Number)
	assert.Equal(t, uint64(1590137711584), header.Timestamp)
}

func TestClient_GetHeaderByNumber(t *testing.T) {
	header, err := testClient.GetHeaderByNumber(ctx, 1)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, uint64(1), header.Number)
	assert.Equal(t, uint64(1590137711584), header.Timestamp)
}

func TestClient_GetHeaderByNumberVerbosity0(t *testing.T) {
	header, err := testClient.GetPackedHeaderByNumber(ctx, 1)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, uint64(1), header.Number)
	assert.Equal(t, uint64(1590137711584), header.Timestamp)
}

func TestClient_GetConsensus(t *testing.T) {
	consensus, err := testClient.GetConsensus(ctx)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, uint64(3500000000), consensus.MaxBlockCycles)
}

func TestClient_GetBlockMedianTime(t *testing.T) {
	blockMedianTime, err := testClient.GetBlockMedianTime(ctx,
		types.HexToHash("0xd5ac7cf8c34a975bf258a34f1c2507638487ab71aa4d10a9ec73704aa3abf9cd"))
	if err != nil {
		t.Fatal(err)
	}
	assert.NotEqual(t, uint64(0), blockMedianTime)
}

func TestClient_GetTransactionProof(t *testing.T) {
	transactionProof, err := testClient.GetTransactionProof(ctx, []string{"0x8277d74d33850581f8d843613ded0c2a1722dec0e87e748f45c115dfb14210f1"}, nil)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, transactionProof.BlockHash)
	assert.Equal(t, 1, len(transactionProof.Proof.Indices))

	result, err := testClient.VerifyTransactionProof(ctx, transactionProof)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 1, len(result))
}

func TestClient_EstimateCycles(t *testing.T) {
	tx := &types.Transaction{
		Version:     0,
		Hash:        types.Hash{},
		CellDeps:    make([]*types.CellDep, 0),
		HeaderDeps:  make([]types.Hash, 0),
		Inputs:      make([]*types.CellInput, 0),
		Outputs:     make([]*types.CellOutput, 0),
		OutputsData: make([][]byte, 0),
		Witnesses:   make([][]byte, 0),
	}
	resp, err := testClient.EstimateCycles(context.Background(), tx)
	assert.Empty(t, err, err)
	assert.NotNil(t, resp)
}

func TestClient_LocalNodeInfo(t *testing.T) {
	nodeInfo, err := testClient.LocalNodeInfo(ctx)
	if err != nil {
		t.Fatal(err)
	}
	assert.True(t, len(nodeInfo.Addresses) > 0)
	assert.True(t, len(nodeInfo.Protocols) > 0)
	assert.True(t, len(nodeInfo.Protocols[0].SupportVersions) > 0)
}

func TestClient_GetPeers(t *testing.T) {
	peers, err := testClient.GetPeers(ctx)
	if err != nil {
		t.Fatal(err)
	}
	assert.True(t, len(peers) > 0)
	assert.True(t, len(peers[0].Addresses) > 0)
	assert.True(t, len(peers[0].Protocols) > 0)
}

func TestClient_SyncState(t *testing.T) {
	state, err := testClient.SyncState(ctx)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotEqual(t, uint64(0), state.BestKnownBlockNumber)
}

func TestClient_SetNetworkActive(t *testing.T) {
	// TODO: fix all deprecated RPC caused tests
	t.Skip("Skipping testing")
	err := testClient.SetNetworkActive(ctx, true)
	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_AddNode(t *testing.T) {
	// TODO: fix all deprecated RPC caused tests
	t.Skip("Skipping testing")
	err := testClient.AddNode(ctx, "QmUsZHPbjjzU627UZFt4k8j6ycEcNvXRnVGxCPKqwbAfQS", "/ip4/192.168.2.100/tcp/8114")
	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_RemoveNode(t *testing.T) {
	// TODO: fix all deprecated RPC caused tests
	t.Skip("Skipping testing")
	err := testClient.RemoveNode(ctx, "QmUsZHPbjjzU627UZFt4k8j6ycEcNvXRnVGxCPKqwbAfQS")
	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_SetBan(t *testing.T) {
	// TODO: fix all deprecated RPC caused tests
	t.Skip("Skipping testing")
	err := testClient.SetBan(ctx, "192.168.0.2", "insert", 1840546800000, true, "test set_ban rpc")
	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_GetBannedAddresses(t *testing.T) {
	bannedAddress, err := testClient.GetBannedAddresses(ctx)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, bannedAddress)
}

func TestClient_ClearBannedAddresses(t *testing.T) {
	// TODO: fix all deprecated RPC caused tests
	t.Skip("Skipping testing")
	err := testClient.ClearBannedAddresses(ctx)
	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_PingPeers(t *testing.T) {
	err := testClient.PingPeers(ctx)
	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_TxPoolInfo(t *testing.T) {
	txPoolInfo, err := testClient.TxPoolInfo(ctx)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotEqual(t, uint64(0), txPoolInfo.MinFeeRate)
	assert.NotEqual(t, types.Hash{}, txPoolInfo.TipHash)
}

func TestClient_ClearTxPool(t *testing.T) {
	// TODO: fix all deprecated RPC caused tests
	t.Skip("Skipping testing")
	err := testClient.ClearTxPool(ctx)
	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_GetRawTxPool(t *testing.T) {
	rawTxPool, err := testClient.GetRawTxPool(ctx)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, rawTxPool)
}

func TestClient_GetBlockchainInfo(t *testing.T) {
	blockchainInfo, err := testClient.GetBlockchainInfo(ctx)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, blockchainInfo)
}

func TestClient_GetLiveCell(t *testing.T) {
	outPoint := types.OutPoint{
		TxHash: types.HexToHash("0xf8de3bb47d055cdf460d93a2a6e1b05f7432f9777c8c474abf4eec1d4aee5d37"),
		Index:  0,
	}
	cellWithStatus, err := testClient.GetLiveCell(ctx, &outPoint, true)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, cellWithStatus)
}

func TestGetTip(t *testing.T) {
	resp, err := testClient.GetIndexerTip(context.Background())
	assert.NoError(t, err)
	assert.NotEqual(t, 0, resp.BlockNumber)
	assert.NotEqual(t, types.Hash{}, resp.BlockHash)
}

func TestGetCellsCapacity(t *testing.T) {
	s := &indexer.SearchKey{
		Script: &types.Script{
			CodeHash: types.HexToHash("0x9bd7e06f3ecf4be0f2fcd2188b23f1b9fcc88e5d4b65a8637b17723bbda3cce8"),
			HashType: types.HashTypeType,
			Args:     ethcommon.FromHex("0x4049ed9cec8a0d39c7a1e899f0dacb8a8c28ad14"),
		},
		ScriptType: types.ScriptTypeLock,
	}
	resp, err := testClient.GetCellsCapacity(context.Background(), s)
	assert.NoError(t, err)
	assert.NotEqual(t, uint64(0x0), resp.BlockNumber)
	assert.NotEqual(t, uint64(0), resp.Capacity)
}

func TestGetCells(t *testing.T) {
	s := &indexer.SearchKey{
		Script: &types.Script{
			CodeHash: types.HexToHash("0x9bd7e06f3ecf4be0f2fcd2188b23f1b9fcc88e5d4b65a8637b17723bbda3cce8"),
			HashType: types.HashTypeType,
			Args:     ethcommon.FromHex("0x4049ed9cec8a0d39c7a1e899f0dacb8a8c28ad14"),
		},
		ScriptType: types.ScriptTypeLock,
	}
	resp, err := testClient.GetCells(context.Background(), s, indexer.SearchOrderAsc, 10, "")
	assert.NoError(t, err)
	assert.Equal(t, 10, len(resp.Objects))

	// Check response when `WithData` == true in request
	s = &indexer.SearchKey{
		Script: &types.Script{
			// https://pudge.explorer.nervos.org/address/ckt1qzda0cr08m85hc8jlnfp3zer7xulejywt49kt2rr0vthywaa50xwsqgxc8z84suk20xzx8337sckkkjfqvzk2ysq48gzc
			CodeHash: types.HexToHash("0x9bd7e06f3ecf4be0f2fcd2188b23f1b9fcc88e5d4b65a8637b17723bbda3cce8"),
			HashType: types.HashTypeType,
			Args:     ethcommon.FromHex("0x06c1c47ac39653cc231e31f4316b5a4903056512"),
		},
		ScriptType: types.ScriptTypeLock,
		WithData:   true,
	}
	resp, err = testClient.GetCells(context.Background(), s, indexer.SearchOrderAsc, 10, "")
	assert.NoError(t, err)
	assert.Equal(t, ethcommon.FromHex("0x0000000000000000"), resp.Objects[0].OutputData)
}

func TestGetTransactions(t *testing.T) {
	s := &indexer.SearchKey{
		Script: &types.Script{
			CodeHash: types.HexToHash("0x58c5f491aba6d61678b7cf7edf4910b1f5e00ec0cde2f42e0abb4fd9aff25a63"),
			HashType: types.HashTypeType,
			Args:     ethcommon.FromHex("0xe53f35ccf63bb37a3bb0ac3b7f89808077a78eae"),
		},
		ScriptType: types.ScriptTypeLock,
	}
	resp, err := testClient.GetTransactions(context.Background(), s, indexer.SearchOrderAsc, 10, "")
	assert.NoError(t, err)
	assert.True(t, len(resp.Objects) >= 1)
	assert.NotEqual(t, 0, resp.Objects[0].BlockNumber)
	assert.NotEqual(t, "", resp.Objects[0].IoType)
}

func TestGetTransactionsGrouped(t *testing.T) {
	s := &indexer.SearchKey{
		Script: &types.Script{
			CodeHash: types.HexToHash("0x58c5f491aba6d61678b7cf7edf4910b1f5e00ec0cde2f42e0abb4fd9aff25a63"),
			HashType: types.HashTypeType,
			Args:     ethcommon.FromHex("0xe53f35ccf63bb37a3bb0ac3b7f89808077a78eae"),
		},
		ScriptType: types.ScriptTypeLock,
	}
	resp, err := testClient.GetTransactionsGrouped(context.Background(), s, indexer.SearchOrderAsc, 10, "")
	assert.NoError(t, err)
	assert.Equal(t, 10, len(resp.Objects))
	assert.NotEqual(t, 0, resp.Objects[0].BlockNumber)
	assert.NotEmpty(t, resp.Objects[0].Cells[0])
	assert.NotEmpty(t, resp.Objects[0].Cells[0].IoType)
	assert.NotEmpty(t, resp.Objects[0].Cells[0].IoIndex)
}

func TestClient_GetFeeRateStatics(t *testing.T) {
	statics, err := testClient.GetFeeRateStatics(context.Background(), nil)
	assert.NoError(t, err)
	assert.NotNil(t, statics)
	statics2, err := testClient.GetFeeRateStatics(context.Background(), 1)
	assert.NoError(t, err)
	assert.NotNil(t, statics2)
}

func TestClient_GetTransactions_PrefixMode(t *testing.T) {
	s := &indexer.SearchKey{
		Script: &types.Script{
			CodeHash: types.HexToHash("0x58c5f491aba6d61678b7cf7edf4910b1f5e00ec0cde2f42e0abb4fd9aff25a63"),
			HashType: types.HashTypeType,
			Args:     ethcommon.FromHex("0xe53f35ccf63bb37a3bb0ac3b7f89808077a78eae"[0:4]),
		},
		ScriptType:       types.ScriptTypeLock,
		ScriptSearchMode: types.ScriptSearchModePrefix,
	}
	resp, err := testClient.GetTransactions(context.Background(), s, indexer.SearchOrderAsc, 10, "")
	assert.NoError(t, err)
	assert.True(t, len(resp.Objects) >= 1)
	assert.NotEqual(t, 0, resp.Objects[0].BlockNumber)
	assert.NotEqual(t, "", resp.Objects[0].IoType)
}

func TestClient_GetTransactions_ExactMode(t *testing.T) {
	s1 := &indexer.SearchKey{
		Script: &types.Script{
			CodeHash: types.HexToHash("0x58c5f491aba6d61678b7cf7edf4910b1f5e00ec0cde2f42e0abb4fd9aff25a63"),
			HashType: types.HashTypeType,
			Args:     ethcommon.FromHex("0xe53f35ccf63bb37a3bb0ac3b7f89808077a78eae"[0:4]),
		},
		ScriptType:       types.ScriptTypeLock,
		ScriptSearchMode: types.ScriptSearchModeExact,
	}
	resp1, err := testClient.GetTransactions(context.Background(), s1, indexer.SearchOrderAsc, 10, "")
	assert.NoError(t, err)
	assert.Equal(t, 0, len(resp1.Objects))

	s2 := &indexer.SearchKey{
		Script: &types.Script{
			CodeHash: types.HexToHash("0x58c5f491aba6d61678b7cf7edf4910b1f5e00ec0cde2f42e0abb4fd9aff25a63"),
			HashType: types.HashTypeType,
			Args:     ethcommon.FromHex("0xe53f35ccf63bb37a3bb0ac3b7f89808077a78eae"),
		},
		ScriptType:       types.ScriptTypeLock,
		ScriptSearchMode: types.ScriptSearchModeExact,
	}
	resp2, err := testClient.GetTransactions(context.Background(), s2, indexer.SearchOrderAsc, 10, "")
	assert.NoError(t, err)
	assert.True(t, len(resp2.Objects) >= 1)
	assert.NotEqual(t, 0, resp2.Objects[0].BlockNumber)
	assert.NotEqual(t, "", resp2.Objects[0].IoType)
}
