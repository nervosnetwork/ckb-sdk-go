package resp

import (
	"math/big"

	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/common"
)

type Balance struct {
	Address   string `json:"address_or_lock_hash"`
	AssetInfo *common.AssetInfo `json:"asset_info"`
	Free      *big.Int `json:"free"`
	Occupied  *big.Int `json:"occupied"`
	Freezed   *big.Int `json:"freezed"`
	Claimable *big.Int `json:"claimable"`
}

type GetBalanceResponse struct {
	Balances []*Balance `json:"balances"`
	TipBlockNumber uint64 `json:"tip_block_number"`
}
