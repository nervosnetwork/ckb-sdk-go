package resp

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/common"
)

func (r *Balance) UnmarshalJSON(input []byte) error {
	var jsonObj struct {
		Ownership string            `json:"ownership"`
		AssetInfo *common.AssetInfo `json:"asset_info"`
		Free      *hexutil.Big      `json:"free"`
		Occupied  *hexutil.Big      `json:"occupied"`
		Frozen    *hexutil.Big      `json:"frozen"`
	}

	if err := json.Unmarshal(input, &jsonObj); err != nil {
		return err
	}
	*r = Balance{
		Ownership: r.Ownership,
		AssetInfo: r.AssetInfo,
		Free:      r.Free,
		Occupied:  r.Occupied,
		Frozen:    r.Frozen,
	}
	return nil
}

func (r *GetBalanceResponse) UnmarshalJSON(input []byte) error {
	var jsonObj struct {
		Balances       []*Balance     `json:"balances"`
		TipBlockNumber hexutil.Uint64 `json:"tip_block_number"`
	}

	if err := json.Unmarshal(input, &jsonObj); err != nil {
		return err
	}
	*r = GetBalanceResponse{
		Balances:       r.Balances,
		TipBlockNumber: r.TipBlockNumber,
	}
	return nil
}
