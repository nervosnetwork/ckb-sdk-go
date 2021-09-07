package model

import (
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/common"
)

type GetBalancePayload struct {
	UdtHashes []interface{} `json:"udt_hashes"`
	BlockNum  uint          `json:"block_num,omitempty"`
	Address   QueryAddress  `json:"address"`
}

type getBalancePayloadBuilder struct {
	Address    QueryAddress
	assetInfos []*common.AssetInfo
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

func (builder *getBalancePayloadBuilder) AddAssetInfo(info *common.AssetInfo) {
	builder.assetInfos = append(builder.assetInfos, info)
}

func (builder *getBalancePayloadBuilder) AddAddress(address string) {
	builder.Address = &KeyAddress{address}
}

func (builder *getBalancePayloadBuilder) Build() *GetBalancePayload {

	payload := &GetBalancePayload{
		Address: builder.Address,
	}
	if len(builder.assetInfos) == 0 {
		payload.UdtHashes = make([]interface{}, 0)
		return payload
	}

	for _, info := range builder.assetInfos {
		if info.AssetType == common.Ckb {
			payload.UdtHashes = append(payload.UdtHashes, nil)
		} else {
			payload.UdtHashes = append(payload.UdtHashes, info.UdtHash)
		}

	}

	return payload
}

func NewGetBalancePayloadBuilder() *getBalancePayloadBuilder {
	return &getBalancePayloadBuilder{}
}
