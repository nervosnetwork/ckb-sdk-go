package req

import . "github.com/nervosnetwork/ckb-sdk-go/mercury/model/types"

type GetBlockInfoPayload struct {
	BlockNumber BlockNumber `json:"block_number,omitempty"`
	BlockHash   H256        `json:"block_hash,omitempty"`
}

type getBlockInfoPayloadBuilder struct {
	blockNumber BlockNumber
	blockHash   H256
}

func (builder *getBlockInfoPayloadBuilder) AddBlockNumber(blockNumber BlockNumber) {
	builder.blockNumber = blockNumber
}

func (builder *getBlockInfoPayloadBuilder) AddBlockHash(blockHash H256) {
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
