package model

type GetBalancePayload struct {
	UdtHashes []interface{} `json:"udt_hashes"`
	BlockNum  uint          `json:"block_num,omitempty"`
	Address   QueryAddress  `json:"address"`
}

type getBalancePayloadBuilder struct {
	UdtHashes []interface{}
	BlockNum  uint
	Address   QueryAddress
}

type QueryAddress interface {
	GetAddress() string
}

// Only addresses in secp256k1 format are available, and the balance contains the balance of addresses in other formats.
type KeyAddress struct {
	KeyAddress string
}

func (addr *KeyAddress) GetAddress() string {
	return addr.KeyAddress
}

// Only the balance of the address in the corresponding format is available.
// For example, the secp256k1 address will only query the balance of the secp256k1 format, and will not contain the balance of the remaining formats.
type NormalAddress struct {
	NormalAddress string
}

func (addr *NormalAddress) GetAddress() string {
	return addr.NormalAddress
}

func (builder *getBalancePayloadBuilder) AddUdtHash(udtHash string) {
	if builder.UdtHashes[0] == nil {
		builder.UdtHashes = builder.UdtHashes[1:]
	}
	builder.UdtHashes = append(builder.UdtHashes, udtHash)
}

func (builder *getBalancePayloadBuilder) AddBlockNum(blockNum uint) {
	builder.BlockNum = blockNum
}

func (builder *getBalancePayloadBuilder) AddAddress(addr string) {
	builder.Address = &KeyAddress{KeyAddress: addr}
}

func (builder *getBalancePayloadBuilder) AddKeyAddress(addr *KeyAddress) {
	builder.Address = addr
}

func (builder *getBalancePayloadBuilder) AddNormalAddressAddress(addr *NormalAddress) {
	builder.Address = addr
}

func (builder *getBalancePayloadBuilder) AllBalance() {
	//var udtHashes []interface{}
	builder.UdtHashes = make([]interface{}, 0)
}

func (builder *getBalancePayloadBuilder) Build() *GetBalancePayload {

	return &GetBalancePayload{
		builder.UdtHashes,
		builder.BlockNum,
		builder.Address,
	}
}

func NewGetBalancePayloadBuilder() *getBalancePayloadBuilder {
	udtHashes := make([]interface{}, 1)
	return &getBalancePayloadBuilder{
		UdtHashes: udtHashes,
	}

}
