package model

type GetGenericBlockPayload struct {
	BlockNum  uint64 `json:"block_num,omitempty"`
	BlockHash string `json:"block_hash,omitempty"`
}

type getGenericBlockPayloadBuilder struct {
	blockNum  uint64
	blockHash string
}

func (builder *getGenericBlockPayloadBuilder) AddBlockNumber(blockNumber uint64) {
	builder.blockNum = blockNumber
}

func (builder *getGenericBlockPayloadBuilder) AddBlockHash(blockHash string) {
	builder.blockHash = blockHash
}

func (builder *getGenericBlockPayloadBuilder) Build() (*GetGenericBlockPayload, error) {
	return &GetGenericBlockPayload{
		builder.blockNum,
		builder.blockHash,
	}, nil
}

func NewGetGenericBlockPayloadBuilder() *getGenericBlockPayloadBuilder {
	return &getGenericBlockPayloadBuilder{}
}
