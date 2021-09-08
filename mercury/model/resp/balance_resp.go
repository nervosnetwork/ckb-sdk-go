package resp

import (
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/common"
	"math/big"
)

type BalanceResp struct {
	Address   string
	AssetInfo *common.AssetInfo
	Free      *big.Int
	Claimable *big.Int
	Freezed   *big.Int
}

type GetBalanceResponse struct {
	Balances []*BalanceResp `json:"balances"`
}
