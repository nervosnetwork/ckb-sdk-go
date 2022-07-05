package model

import "github.com/nervosnetwork/ckb-sdk-go/types"

type AssetType string

const (
	AssetTypeCKB AssetType = "CKB"
	AssetTypeUDT AssetType = "UDT"
)

type AssetInfo struct {
	AssetType AssetType  `json:"asset_type"`
	UdtHash   types.Hash `json:"udt_hash"`
}

func NewCkbAsset() *AssetInfo {
	return &AssetInfo{
		AssetType: AssetTypeCKB,
		UdtHash:   types.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
	}
}

func NewUdtAsset(udtHash types.Hash) *AssetInfo {
	return &AssetInfo{
		AssetType: AssetTypeUDT,
		UdtHash:   udtHash,
	}
}
