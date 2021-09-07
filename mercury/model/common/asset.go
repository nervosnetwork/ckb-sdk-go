package common

type AssetType string

const (
	Ckb AssetType = "Ckb"
	Udt           = "Udt"
)

type AssetInfo struct {
	AssetType AssetType
	UdtHash   string
}

func NewCkbAsset() *AssetInfo {
	return &AssetInfo{
		AssetType: Ckb,
	}
}

func NewUdtAsset(udtHash string) *AssetInfo {
	return &AssetInfo{
		AssetType: Udt,
		UdtHash:   udtHash,
	}
}
