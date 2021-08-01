package test

import (
	"context"
	"fmt"
	E "github.com/ethereum/go-ethereum/rpc"
	"github.com/nervosnetwork/ckb-sdk-go/api"
	"github.com/nervosnetwork/ckb-sdk-go/indexer"
	"github.com/nervosnetwork/ckb-sdk-go/mercury"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/example/constant"
	"github.com/nervosnetwork/ckb-sdk-go/rpc"
	"github.com/stretchr/testify/assert"
	"testing"
)

const RPC_ADDRESS = "http://127.0.0.1:8116"

func TestUseCkbAlone(t *testing.T) {
	client, err := rpc.Dial(RPC_ADDRESS)
	assert.Nil(t, err)

	number, err := client.GetTipBlockNumber(context.Background())
	assert.Nil(t, err)

	fmt.Println(number)
}

func TestUseIndexerAlone(t *testing.T) {
	client, err := indexer.Dial(RPC_ADDRESS)
	assert.Nil(t, err)

	number, err := client.GetTip(context.Background())
	assert.Nil(t, err)

	fmt.Printf("block info: %+v", number)

}

func TestUseMercuryAlone(t *testing.T) {
	client, err := mercury.Dial(RPC_ADDRESS)
	assert.Nil(t, err)

	number, err := client.RegisterAddresses([]string{constant.TEST_ADDRESS3})
	assert.Nil(t, err)

	fmt.Printf("block info: %+v", number)

}

func TestUseCkbAndIndexer(t *testing.T) {

	dial, err := E.Dial(RPC_ADDRESS)
	assert.Nil(t, err)

	indexerClient := indexer.NewClient(dial)
	ckbClient := rpc.NewClientWithIndexer(dial, indexerClient)

	// rpc using ckb
	number1, err := ckbClient.GetTipBlockNumber(context.Background())
	assert.Nil(t, err)
	fmt.Println(number1)

	// rpc using indexer
	number2, err := ckbClient.GetTip(context.Background())
	assert.Nil(t, err)
	fmt.Printf("block info: %+v", number2)

}

func TestUseCkbApiAlone(t *testing.T) {
	api, err := api.NewCkbApi(RPC_ADDRESS)
	assert.Nil(t, err)

	// rpc using ckb
	number1, err := api.GetTipBlockNumber(context.Background())
	assert.Nil(t, err)
	fmt.Println(number1)

	// rpc using indexer
	number2, err := api.GetTip(context.Background())
	assert.Nil(t, err)
	fmt.Printf("block info: %+v\n", number2)

	// rpc using mercury
	scriptHashes, err := api.RegisterAddresses([]string{constant.TEST_ADDRESS3})
	assert.Nil(t, err)
	fmt.Println(scriptHashes)
}
