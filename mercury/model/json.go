package model

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/common"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/req"
)

func (r GetBalancePayload) MarshalJSON() ([]byte, error) {
	jsonObj := &struct {
		AssetInfos     []*common.AssetInfo `json:"asset_infos"`
		TipBlockNumber hexutil.Uint64      `json:"tip_block_number,omitempty"`
		Item           *req.Item           `json:"item"`
	}{
		AssetInfos:     r.AssetInfos,
		TipBlockNumber: hexutil.Uint64(r.TipBlockNumber),
		Item:           r.Item,
	}
	return json.Marshal(jsonObj)
}
