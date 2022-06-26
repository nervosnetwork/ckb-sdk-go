package model

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/common"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/req"
	"github.com/nervosnetwork/ckb-sdk-go/types"
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

func (r GetBlockInfoPayload) MarshalJSON() ([]byte, error) {
	jsonObj := &struct {
		BlockNumber hexutil.Uint64 `json:"block_number,omitempty"`
		BlockHash   types.Hash     `json:"block_hash,omitempty"`
	}{
		BlockNumber: hexutil.Uint64(r.BlockNumber),
		BlockHash:   r.BlockHash,
	}
	return json.Marshal(jsonObj)
}

func (r BlockRange) MarshalJSON() ([]byte, error) {
	jsonObj := &struct {
		From hexutil.Uint64 `json:"from"`
		To   hexutil.Uint64 `json:"to"`
	}{
		From: hexutil.Uint64(r.From),
		To:   hexutil.Uint64(r.To),
	}
	return json.Marshal(jsonObj)
}

func (r PaginationRequest) MarshalJSON() ([]byte, error) {
	jsonObj := &struct {
		Cursor      hexutil.Uint64 `json:"cursor,omitempty"`
		Order       Order          `json:"order"`
		Limit       hexutil.Uint64 `json:"limit,omitempty"`
		ReturnCount bool           `json:"return_count"`
	}{
		Cursor:      hexutil.Uint64(r.Cursor),
		Order:       r.Order,
		Limit:       hexutil.Uint64(r.Limit),
		ReturnCount: r.ReturnCount,
	}
	return json.Marshal(jsonObj)
}
