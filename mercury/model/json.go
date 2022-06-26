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

func (r BuildAdjustAccountPayload) MarshalJSON() ([]byte, error) {
	jsonObj := &struct {
		Item          *req.Item         `json:"item"`
		From          []*req.Item       `json:"from"`
		AssetInfo     *common.AssetInfo `json:"asset_info"`
		AccountNumber hexutil.Uint64    `json:"account_number,omitempty"`
		ExtraCKB      hexutil.Uint64    `json:"extra_ckb,omitempty"`
		FeeRate       hexutil.Uint64    `json:"fee_rate,omitempty"`
	}{
		Item:          r.Item,
		From:          r.From,
		AssetInfo:     r.AssetInfo,
		AccountNumber: hexutil.Uint64(r.AccountNumber),
		ExtraCKB:      hexutil.Uint64(r.ExtraCKB),
		FeeRate:       hexutil.Uint64(r.FeeRate),
	}
	return json.Marshal(jsonObj)
}

func (r SinceConfig) MarshalJSON() ([]byte, error) {
	jsonObj := &struct {
		Flag  SinceFlag      `json:"flag"`
		Type  SinceType      `json:"type_"`
		Value hexutil.Uint64 `json:"value"`
	}{
		Flag:  r.Flag,
		Type:  r.Type,
		Value: hexutil.Uint64(r.Value),
	}
	return json.Marshal(jsonObj)
}

func (r ToInfo) MarshalJSON() ([]byte, error) {
	jsonObj := &struct {
		Address string       `json:"address"`
		Amount  *hexutil.Big `json:"amount"`
	}{
		Address: r.Address,
		Amount:  (*hexutil.Big)(r.Amount),
	}
	return json.Marshal(jsonObj)
}

func (r TransferPayload) MarshalJSON() ([]byte, error) {
	type transferPayloadAlias TransferPayload
	jsonObj := &struct {
		transferPayloadAlias
		FeeRate hexutil.Uint64 `json:"fee_rate,omitempty"`
	}{
		transferPayloadAlias: transferPayloadAlias(r),
		FeeRate:              hexutil.Uint64(r.FeeRate),
	}
	return json.Marshal(jsonObj)
}

func (r SimpleTransferPayload) MarshalJSON() ([]byte, error) {
	type simpleTransferPayloadAlias SimpleTransferPayload
	jsonObj := &struct {
		simpleTransferPayloadAlias
		FeeRate hexutil.Uint64 `json:"fee_rate"`
	}{
		simpleTransferPayloadAlias: simpleTransferPayloadAlias(r),
		FeeRate:                    hexutil.Uint64(r.FeeRate),
	}
	return json.Marshal(jsonObj)
}

func (r DaoDepositPayload) MarshalJSON() ([]byte, error) {
	type daoDepositPayloadAlias DaoDepositPayload
	jsonObj := &struct {
		daoDepositPayloadAlias
		Amount  hexutil.Uint64 `json:"amount"`
		FeeRate hexutil.Uint64 `json:"fee_rate,omitempty"`
	}{
		daoDepositPayloadAlias: daoDepositPayloadAlias(r),
		Amount:                 hexutil.Uint64(r.Amount),
		FeeRate:                hexutil.Uint64(r.FeeRate),
	}
	return json.Marshal(jsonObj)
}

func (r DaoWithdrawPayload) MarshalJSON() ([]byte, error) {
	jsonObj := &struct {
		From    *req.Item `json:"from"`
		FeeRate uint64    `json:"fee_rate,omitempty"`
	}{
		From:    r.From,
		FeeRate: r.FeeRate,
	}
	return json.Marshal(jsonObj)
}

func (r DaoClaimPayload) MarshalJSON() ([]byte, error) {
	jsonObj := &struct {
		From    *req.Item      `json:"from"`
		To      string         `json:"to,omitempty"`
		FeeRate hexutil.Uint64 `json:"fee_rate,omitempty"`
	}{
		From:    r.From,
		To:      r.To,
		FeeRate: hexutil.Uint64(r.FeeRate),
	}
	return json.Marshal(jsonObj)
}
