package resp

import (
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/common"
	"math/big"
)

type Balance struct {
	Ownership string            `json:"ownership"`
	AssetInfo *common.AssetInfo `json:"asset_info"`
	Free      *big.Int          `json:"free"`
	Occupied  *big.Int          `json:"occupied"`
	Frozen    *big.Int          `json:"frozen"`
}

type GetBalanceResponse struct {
	Balances       []*Balance `json:"balances"`
	TipBlockNumber uint64     `json:"tip_block_number"`
}
