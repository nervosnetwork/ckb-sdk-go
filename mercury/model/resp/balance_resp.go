package resp

import "github.com/nervosnetwork/ckb-sdk-go/mercury/model/common"

type BalanceResp struct {
	Address   string
	AssetInfo *common.AssetInfo
	Free      string
	Claimable string
	Freezed   string
}

type GetBalanceResponse struct {
	Balances []*BalanceResp `json:"balances"`
}
