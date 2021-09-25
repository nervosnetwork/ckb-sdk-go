package model

type GetBlockInfoPayload struct {
	BlockNumber uint64 `json:"block_number,omitempty"`
	BlockHash   string `json:"block_hash,omitempty"`
}

type getBlockInfoPayloadBuilder struct {
	blockNumber uint64
	blockHash   string
}

func (builder *getBlockInfoPayloadBuilder) AddBlockNumber(blockNumber uint64) {
	builder.blockNumber = blockNumber
}

func (builder *getBlockInfoPayloadBuilder) AddBlockHash(blockHash string) {
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
