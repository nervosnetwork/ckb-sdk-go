package model

import (
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/common"
)

type GetBalancePayload struct {
	AssetInfos     []*common.AssetInfo `json:"asset_infos"`
	TipBlockNumber uint64              `json:"tip_block_number,omitempty"`
	Item           interface{}         `json:"item"`
}

type getBalancePayloadBuilder struct {
	item           interface{}
	assetInfos     []*common.AssetInfo
	TipBlockNumber uint64
}

func (builder *getBalancePayloadBuilder) AddAssetInfo(info *common.AssetInfo) {
	builder.assetInfos = append(builder.assetInfos, info)
}

func (builder *getBalancePayloadBuilder) AddItem(item interface{}) {
	builder.item = item
}

func (builder *getBalancePayloadBuilder) AddTipBlockNumber(tipBlockNumber uint64) {
	builder.TipBlockNumber = tipBlockNumber
}

func (builder *getBalancePayloadBuilder) Build() *GetBalancePayload {

	payload := &GetBalancePayload{
		Item:           builder.item,
		TipBlockNumber: builder.TipBlockNumber,
		AssetInfos:     builder.assetInfos,
	}

	return payload
}

func NewGetBalancePayloadBuilder() *getBalancePayloadBuilder {
	return &getBalancePayloadBuilder{
		assetInfos: make([]*common.AssetInfo, 0),
	}
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
