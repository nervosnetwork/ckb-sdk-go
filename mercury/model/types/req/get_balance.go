package req

import (
	. "github.com/nervosnetwork/ckb-sdk-go/mercury/model/types"
)

type GetBalancePayload struct {
	AssetInfos []*AssetInfo     `json:"asset_infos"`
	TipBlockNumber *BlockNumber `json:"block_numer,omitempty"`
	Item           Item               `json:"item"`
}

type Item interface {
	GetAddress() string
}

// Only addresses in secp256k1 format are available, and the balance contains the balance of addresses in other formats.
type Identity struct {
	Identity string `json:"identity"`
}

func (addr *Identity) GetAddress() string {
	return addr.Identity
}

// Only the balance of the address in the corresponding format is available.
// For example, the secp256k1 address will only query the balance of the secp256k1 format, and will not contain the balance of the remaining formats.
// see types.Address

type RecordId struct {
	RecordId string `json:"record_id"`
}
func (addr *RecordId) GetAddress() string {
	return addr.RecordId
}

type GetBalancePayloadBuilder struct {
	Item       Item
	assetInfos []*AssetInfo
	TipBlockNumber *BlockNumber
}

func (builder *GetBalancePayloadBuilder) AddAssetInfo(info *AssetInfo) {
	builder.assetInfos = append(builder.assetInfos, info)
}

func (builder *GetBalancePayloadBuilder) SetItemAsAddress(address string) {
	builder.Item = &Address{address}
}

func (builder *GetBalancePayloadBuilder) SetItem(item Item) {
	builder.Item = item
}

func (builder *GetBalancePayloadBuilder) SetTipBlockNumber(tipBlockNumber BlockNumber) {
	builder.TipBlockNumber = &tipBlockNumber
}

func (builder *GetBalancePayloadBuilder) Build() *GetBalancePayload {
	payload := &GetBalancePayload{
		AssetInfos: builder.assetInfos,
		Item: builder.Item,
		TipBlockNumber: builder.TipBlockNumber,
	}

	return payload
}

func NewGetBalancePayloadBuilder() *GetBalancePayloadBuilder {
	return &GetBalancePayloadBuilder{
		Item:           nil,
		assetInfos:     []*AssetInfo{},
		TipBlockNumber: nil,
	}
}
