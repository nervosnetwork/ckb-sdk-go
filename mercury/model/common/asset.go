package common

import "github.com/nervosnetwork/ckb-sdk-go/types"

type AssetType string

const (
	CKB AssetType = "CKB"
	UDT AssetType = "UDT"
)

type AssetInfo struct {
	AssetType AssetType  `json:"asset_type"`
	UdtHash   types.Hash `json:"udt_hash"`
}

func NewCkbAsset() *AssetInfo {
	return &AssetInfo{
		AssetType: CKB,
		UdtHash:   types.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
	}
}

func NewUdtAsset(udtHash string) *AssetInfo {
	return &AssetInfo{
		AssetType: UDT,
		UdtHash:   types.HexToHash(udtHash),
	}
}
