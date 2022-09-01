package rpc

import (
	"context"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/nervosnetwork/ckb-sdk-go/types"
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

func TestClient_GetTransaction(t *testing.T) {
	txView, err := testClient.GetTransaction(ctx,
		types.HexToHash("0x8277d74d33850581f8d843613ded0c2a1722dec0e87e748f45c115dfb14210f1"))
	if err != nil {
		t.Fatal(err)
	}
	tx := txView.Transaction
	assert.Equal(t, 4, len(tx.CellDeps))
	assert.Equal(t, 1, len(tx.Inputs))
	assert.Equal(t, 3, len(tx.Outputs))
	assert.Equal(t, uint64(30000000000), tx.Outputs[0].Capacity)
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

func TestClient_GetHeaderByNumber(t *testing.T) {
	header, err := testClient.GetHeaderByNumber(ctx, 1)
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
	err := testClient.SetNetworkActive(ctx, true)
	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_AddNode(t *testing.T) {
	err := testClient.AddNode(ctx, "QmUsZHPbjjzU627UZFt4k8j6ycEcNvXRnVGxCPKqwbAfQS", "/ip4/192.168.2.100/tcp/8114")
	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_RemoveNode(t *testing.T) {
	err := testClient.RemoveNode(ctx, "QmUsZHPbjjzU627UZFt4k8j6ycEcNvXRnVGxCPKqwbAfQS")
	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_SetBan(t *testing.T) {
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
