package rpc

import (
	"context"
	"encoding/json"
	"fmt"
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
