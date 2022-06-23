package model

import "github.com/nervosnetwork/ckb-sdk-go/types"

type GetBlockInfoPayload struct {
	BlockNumber uint64     `json:"block_number,omitempty"`
	BlockHash   types.Hash `json:"block_hash,omitempty"`
}

type getBlockInfoPayloadBuilder struct {
	blockNumber uint64
	blockHash   types.Hash
}

func (builder *getBlockInfoPayloadBuilder) AddBlockNumber(blockNumber uint64) {
	builder.blockNumber = blockNumber
}

func (builder *getBlockInfoPayloadBuilder) AddBlockHash(blockHash string) {
	builder.blockHash = types.HexToHash(blockHash)
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
