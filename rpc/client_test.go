package rpc

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	mock_rpc "github.com/nervosnetwork/ckb-sdk-go/test/mock/rpc"
)

func TestGetTipBlockNumber(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mc := mock_rpc.NewMockClient(ctrl)

	mc.
		EXPECT().
		GetTipBlockNumber(gomock.Any()).
		Return(uint64(100), nil).
		AnyTimes()

	num, err := mc.GetTipBlockNumber(context.Background())

	assert.Nil(t, err, "get tip block number error")
	assert.Equal(t, uint64(100), num)
}
