package test

import (
	"encoding/json"
	"fmt"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/example/constant"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model"
	"strings"
	"testing"
)

func TestGetGenericBlockWithBlockNumber(t *testing.T) {
	builder := model.NewGetGenericBlockPayloadBuilder()
	builder.AddBlockNumber(2172093)

	payload, err := builder.Build()
	if err != nil {
		t.Error(err)
	}

	json2, err := json.Marshal(payload)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(json2))

	block, err := constant.GetMercuryApiInstance().GetBlockInfo(payload)
	if err != nil {
		t.Error(err)
	}

	json, err := json.Marshal(block)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(json))
}

func TestGetGenericBlockWithBlockHash(t *testing.T) {
	builder := model.NewGetGenericBlockPayloadBuilder()
	builder.AddBlockHash("0xee8adba356105149cb9dc1cb0d09430a6bd01182868787ace587961c0d64e742")

	payload, err := builder.Build()
	if err != nil {
		t.Error(err)
	}

	block, err := constant.GetMercuryApiInstance().GetBlockInfo(payload)
	if err != nil {
		t.Error(err)
	}

	json, err := json.Marshal(block)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(json))

}

func TestGetGenericBlockWithBlockHashAndBlockNumber(t *testing.T) {
	builder := model.NewGetGenericBlockPayloadBuilder()
	builder.AddBlockNumber(2172093)
	builder.AddBlockHash("0xee8adba356105149cb9dc1cb0d09430a6bd01182868787ace587961c0d64e742")

	payload, err := builder.Build()
	if err != nil {
		t.Error(err)
	}

	block, err := constant.GetMercuryApiInstance().GetBlockInfo(payload)
	if err != nil {
		t.Error(err)
	}

	json, err := json.Marshal(block)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(json))
}

func TestTipGenericBlock(t *testing.T) {
	builder := model.NewGetGenericBlockPayloadBuilder()

	payload, err := builder.Build()
	if err != nil {
		t.Error(err)
	}

	block, err := constant.GetMercuryApiInstance().GetBlockInfo(payload)
	if err != nil {
		t.Error(err)
	}

	json, err := json.Marshal(block)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(json))
}

func TestBlockHashAndBlockNumberDoNotMatch(t *testing.T) {
	builder := model.NewGetGenericBlockPayloadBuilder()
	builder.AddBlockNumber(2172092)
	builder.AddBlockHash("0xee8adba356105149cb9dc1cb0d09430a6bd01182868787ace587961c0d64e742")

	payload, err := builder.Build()
	if err != nil {
		t.Error(err)
	}

	_, getErr := constant.GetMercuryApiInstance().GetBlockInfo(payload)
	if getErr != nil && getErr.Error() != "block number and hash mismatch" {
		t.Error(err)
	}
}

func TestCannotFind(t *testing.T) {
	builder := model.NewGetGenericBlockPayloadBuilder()
	builder.AddBlockHash("0xee8adba356105149cb9dc1cb0d09430a6bd01182868787ace587961c0d64e741")

	payload, err := builder.Build()
	if err != nil {
		t.Error(err)
	}

	_, getErr := constant.GetMercuryApiInstance().GetBlockInfo(payload)
	if getErr != nil && !strings.HasPrefix(getErr.Error(), "Cannot get block by hash H256") {
		t.Error(err)
	}
}

func TestWrongHeight(t *testing.T) {
	builder := model.NewGetGenericBlockPayloadBuilder()
	builder.AddBlockNumber(217209233)

	payload, err := builder.Build()
	if err != nil {
		t.Error(err)
	}
	// error: invalid block number
	_, getErr := constant.GetMercuryApiInstance().GetBlockInfo(payload)
	if getErr != nil && getErr.Error() != "invalid block number" {
		t.Error(err)
	}
}
