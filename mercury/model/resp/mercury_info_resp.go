package resp

import "github.com/nervosnetwork/ckb-sdk-go/mercury/model/common"

type MercuryInfo struct {
    MercuryVersion string `json:"mercury_version"`
    CkbNodeVersion string `json:"ckb_node_version"`
    NetworkType NetworkType `json:"network_type"`
    EnabledExtensions []Extension `json:"enabled_extensions"`
}

type NetworkType string
const (
    Mainnet NetworkType = "Mainnet"
    Testnet NetworkType = "Testnet"
    Staging NetworkType = "Staging"
    Dev     NetworkType = "Dev"
)

type Extension struct {
	Name string `json:"name"`
	Scripts []common.Script `json:"scripts"`
	CellDeps []common.CellDep `json:"cell_deps"`
}
