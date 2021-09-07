package model

type GetBlockInfoPayload struct {
	BlockNum  uint64 `json:"block_num,omitempty"`
	BlockHash string `json:"block_hash,omitempty"`
}

type getBlockInfoPayloadBuilder struct {
	blockNum  uint64
	blockHash string
}

func (builder *getBlockInfoPayloadBuilder) AddBlockNumber(blockNumber uint64) {
	builder.blockNum = blockNumber
}

func (builder *getBlockInfoPayloadBuilder) AddBlockHash(blockHash string) {
	builder.blockHash = blockHash
}

func (builder *getBlockInfoPayloadBuilder) Build() (*GetBlockInfoPayload, error) {
	return &GetBlockInfoPayload{
		builder.blockNum,
		builder.blockHash,
	}, nil
}

func NewGetGenericBlockPayloadBuilder() *getBlockInfoPayloadBuilder {
	return &getBlockInfoPayloadBuilder{}
}
