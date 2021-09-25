package model

import "github.com/nervosnetwork/ckb-sdk-go/mercury/model/common"

type GetBlockInfoPayload struct {
	BlockNumber common.BlockNumber `json:"block_number,omitempty"`
	BlockHash common.H256 `json:"block_hash,omitempty"`
}

type getBlockInfoPayloadBuilder struct {
	blockNumber  common.BlockNumber
	blockHash common.H256
}

func (builder *getBlockInfoPayloadBuilder) AddBlockNumber(blockNumber common.BlockNumber) {
	builder.blockNumber = blockNumber
}

func (builder *getBlockInfoPayloadBuilder) AddBlockHash(blockHash common.H256) {
	builder.blockHash = blockHash
}

func (builder *getBlockInfoPayloadBuilder) Build() (*GetBlockInfoPayload, error) {
	return &GetBlockInfoPayload{
		builder.blockNumber,
		builder.blockHash,
	}, nil
}

func NewGetGenericBlockPayloadBuilder() *getBlockInfoPayloadBuilder {
	return &getBlockInfoPayloadBuilder{}
}
