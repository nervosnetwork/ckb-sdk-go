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

func getApi() Client {
	api, _ := Dial("http://localhost:8114")
	return api
}
