package rpc

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/nervosnetwork/ckb-sdk-go/types"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	mockRpc "github.com/nervosnetwork/ckb-sdk-go/test/mock/rpc"
)

func TestGetTipBlockNumber(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mc := mockRpc.NewMockClient(ctrl)

	mc.
		EXPECT().
		GetTipBlockNumber(gomock.Any()).
		Return(uint64(100), nil).
		AnyTimes()

	num, err := mc.GetTipBlockNumber(context.Background())

	assert.Nil(t, err, "get tip block number error")
	assert.Equal(t, uint64(100), num)
}

func TestSyncState(t *testing.T) {
	api, _ := Dial("http://localhost:8114")
	syncState, err := api.SyncState(context.Background())
	assert.Nil(t, err)

	json, err := json.Marshal(syncState)
	assert.Nil(t, err)
	fmt.Println(string(json))
}

func TestGetTransactionProof(t *testing.T) {
	api := getApi()

	proof, err := api.GetTransactionProof(context.Background(), []string{"0xc9ae96ff99b48e755ccdb350a69591ba80877be3d6c67ac9660bb9a0c52dc3d6"}, nil)
	assert.Nil(t, err)

	marshal, err := json.Marshal(proof)
	assert.Nil(t, err)
	fmt.Println(string(marshal))
}

func TestGetTransactionProofByBlockHash(t *testing.T) {
	api := getApi()
	hash := types.HexToHash("0x36038509b555c8acf360175b9bc4f67bd68be02b152f4a9d1131a424fffd8d23")
	proof, err := api.GetTransactionProof(context.Background(), []string{"0xc9ae96ff99b48e755ccdb350a69591ba80877be3d6c67ac9660bb9a0c52dc3d6"}, &hash)
	assert.Nil(t, err)

	marshal, err := json.Marshal(proof)
	assert.Nil(t, err)
	fmt.Println(string(marshal))
}

func TestVerifyTransactionProof(t *testing.T) {
	api := getApi()

	proof := &types.TransactionProof{
		Proof: &types.Proof{
			Indices: []uint{2},
			Lemmas:  []types.Hash{types.HexToHash("0x705d0774a1f870c1e92571e9db806bd85c0ac7f26015f3d6c7b822f7616c1fb4")},
		},
		BlockHash:     types.HexToHash("0x36038509b555c8acf360175b9bc4f67bd68be02b152f4a9d1131a424fffd8d23"),
		WitnessesRoot: types.HexToHash("0x56431856ad780db4cc1181c44b3fddf596380f1e21fb1c0b31db6deca2892c75"),
	}

	json, err := json.Marshal(proof)
	assert.Nil(t, err)
	fmt.Println(string(json))

	result, err := api.VerifyTransactionProof(context.Background(), proof)

	fmt.Println(result)
}

func TestSetNetworkActive(t *testing.T) {
	api := getApi()
	err := api.SetNetworkActive(context.Background(), true)
	assert.Nil(t, err)
}

func TestClearBannedAddresses(t *testing.T) {
	api := getApi()
	err := api.ClearBannedAddresses(context.Background())
	assert.Nil(t, err)
}

func TestAddNode(t *testing.T) {
	api := getApi()
	err := api.AddNode(context.Background(), "QmUsZHPbjjzU627UZFt4k8j6ycEcNvXRnVGxCPKqwbAfQS", "/ip4/192.168.2.100/tcp/8114")
	assert.Nil(t, err)
}

func TestRemoveNode(t *testing.T) {
	api := getApi()
	err := api.RemoveNode(context.Background(), "QmUsZHPbjjzU627UZFt4k8j6ycEcNvXRnVGxCPKqwbAfQS")
	assert.Nil(t, err)
}

func TestPingPeers(t *testing.T) {
	api := getApi()
	err := api.PingPeers(context.Background())
	assert.Nil(t, err)
}

func TestGetRawTxPool(t *testing.T) {
	api := getApi()
	pool, err := api.GetRawTxPool(context.Background())
	assert.Nil(t, err)

	json, err := json.Marshal(pool)
	assert.Nil(t, err)

	fmt.Println(len(pool.Pending))
	fmt.Println(len(pool.Proposed))

	fmt.Println(string(json))
}

func getApi() Client {
	api, _ := Dial("http://localhost:8114")
	return api
}
