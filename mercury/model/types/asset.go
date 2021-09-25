package types

type AssetType string

const (
	CKB AssetType = "CKB"
	UDT AssetType = "UDT"
)

type AssetInfo struct {
	AssetType AssetType `json:"asset_type"`
	UdtHash   string    `json:"udt_hash"`
}

func NewCkbAsset() *AssetInfo {
	return &AssetInfo{
		AssetType: CKB,
	}
}

func NewUdtAsset(udtHash string) *AssetInfo {
	return &AssetInfo{
		AssetType: UDT,
		UdtHash:   udtHash,
	}
}
