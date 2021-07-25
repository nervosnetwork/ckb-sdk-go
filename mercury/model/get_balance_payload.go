package model

import (
	"encoding/hex"
	"github.com/nervosnetwork/ckb-sdk-go/address"
	"github.com/nervosnetwork/ckb-sdk-go/crypto/bech32"
)

type GetBalancePayload struct {
	UdtHashes []interface{} `json:"udt_hashes"`
	BlockNum  uint          `json:"block_num,omitempty"`
	Address   QueryAddress  `json:"address"`
}

type getBalancePayloadBuilder struct {
	UdtHashes []interface{}
	BlockNum  uint
	Address   string
}

type QueryAddress interface {
	GetAddress() string
}

type KeyAddress struct {
	KeyAddress string
}

func (addr *KeyAddress) GetAddress() string {
	return addr.KeyAddress
}

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

	builder.Address = addr
}

func (builder *getBalancePayloadBuilder) AllBalance() {
	//var udtHashes []interface{}
	builder.UdtHashes = make([]interface{}, 0)
}

func (builder *getBalancePayloadBuilder) Build() (*GetBalancePayload, error) {
	address, err := getQueryAddressByAddress(builder.Address)
	return &GetBalancePayload{
		builder.UdtHashes,
		builder.BlockNum,
		address,
	}, err
}

func getQueryAddressByAddress(addr string) (QueryAddress, error) {
	_, decoded, err := bech32.Decode(addr)
	if err != nil {
		return nil, err
	}
	data, err := bech32.ConvertBits(decoded, 5, 8, false)
	if err != nil {
		return nil, err
	}
	payload := hex.EncodeToString(data)
	if address.CodeHashIndexSingleSig == payload[2:4] {
		return &KeyAddress{addr}, nil
	} else {
		return &NormalAddress{addr}, nil
	}

}

func NewGetBalancePayloadBuilder() *getBalancePayloadBuilder {
	udtHashes := make([]interface{}, 1)
	return &getBalancePayloadBuilder{
		UdtHashes: udtHashes,
	}

}
