package common

type AssetType string

const (
	CKB AssetType = "CKB"
	UDT           = "UDT"
)

type AssetInfo struct {
	AssetType AssetType `json:"asset_type"`
	UdtHash   string    `json:"udt_hash"`
}

func NewCkbAsset() *AssetInfo {
	return &AssetInfo{
		AssetType: CKB,
		UdtHash:   "0x0000000000000000000000000000000000000000000000000000000000000000",
	}
}

func NewUdtAsset(udtHash string) *AssetInfo {
	return &AssetInfo{
		AssetType: UDT,
		UdtHash:   udtHash,
	}
}
