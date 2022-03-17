package resp

import (
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/common"
)

type Balance struct {
	Ownership *common.Ownership `json:"ownership"`
	AssetInfo *common.AssetInfo `json:"asset_info"`
	Free      *model.U128       `json:"free"`
	Occupied  *model.U128       `json:"occupied"`
	Frozen    *model.U128       `json:"frozen"`
	Claimable *model.U128       `json:"claimable"`
}

type GetBalanceResponse struct {
	Balances       []*Balance `json:"balances"`
	TipBlockNumber uint64     `json:"tip_block_number"`
}
