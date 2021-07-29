package model

type QueryGenericTransactionsPayload struct {
	Address   QueryAddress  `json:"address"`
	UdtHashes []interface{} `json:"udt_hashes"`
	FromBlock uint64        `json:"from_block,omitempty"`
	ToBlock   uint64        `json:"to_block,omitempty"`
	Limit     uint64        `json:"limit,omitempty"`
	Offset    uint64        `json:"offset,omitempty"`
	Order     string        `json:"order,omitempty"`
}

type queryGenericTransactionsPayloadBuilder struct {
	Address   QueryAddress
	UdtHashes []interface{}
	FromBlock uint64
	ToBlock   uint64
	Limit     uint64
	Offset    uint64
	Order     string
}

func (builder *queryGenericTransactionsPayloadBuilder) AddKeyAddress(addr *KeyAddress) {
	builder.Address = addr
}

func (builder *queryGenericTransactionsPayloadBuilder) AddNormalAddress(addr *NormalAddress) {
	builder.Address = addr
}

func (builder *queryGenericTransactionsPayloadBuilder) AddUdtHash(udtHash string) {
	if builder.UdtHashes[0] == nil {
		builder.UdtHashes = builder.UdtHashes[1:]
	}
	builder.UdtHashes = append(builder.UdtHashes, udtHash)
}

func (builder *queryGenericTransactionsPayloadBuilder) AllTransactionType() {
	//var udtHashes []interface{}
	builder.UdtHashes = make([]interface{}, 0)
}

func (builder *queryGenericTransactionsPayloadBuilder) AddFromBlock(fromBlock uint64) {
	builder.FromBlock = fromBlock
}

func (builder *queryGenericTransactionsPayloadBuilder) AddToBlock(toBlock uint64) {
	builder.ToBlock = toBlock
}

func (builder *queryGenericTransactionsPayloadBuilder) AddLimit(limit uint64) {
	builder.Limit = limit
}

func (builder *queryGenericTransactionsPayloadBuilder) AddOffset(offset uint64) {
	builder.Offset = offset
}

func (builder *queryGenericTransactionsPayloadBuilder) AddOrder(order string) {
	builder.Order = order
}

func (builder *queryGenericTransactionsPayloadBuilder) Build() *QueryGenericTransactionsPayload {
	return &QueryGenericTransactionsPayload{
		Address:   builder.Address,
		UdtHashes: builder.UdtHashes,
		FromBlock: builder.FromBlock,
		ToBlock:   builder.ToBlock,
		Limit:     builder.Limit,
		Offset:    builder.Offset,
		Order:     builder.Order,
	}

}

func NewQueryGenericTransactionsPayloadBuilder() *queryGenericTransactionsPayloadBuilder {
	udtHashes := make([]interface{}, 1)
	return &queryGenericTransactionsPayloadBuilder{
		UdtHashes: udtHashes,
	}
}
