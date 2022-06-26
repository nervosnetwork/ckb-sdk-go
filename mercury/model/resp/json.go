package resp

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/common"
	"github.com/nervosnetwork/ckb-sdk-go/types"
	"math/big"
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
		Ownership: jsonObj.Ownership,
		AssetInfo: jsonObj.AssetInfo,
		Free:      (*big.Int)(jsonObj.Free),
		Occupied:  (*big.Int)(jsonObj.Occupied),
		Frozen:    (*big.Int)(jsonObj.Frozen),
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
		Balances:       jsonObj.Balances,
		TipBlockNumber: uint64(jsonObj.TipBlockNumber),
	}
	return nil
}

func (r *BurnInfo) UnmarshalJSON(input []byte) error {
	var jsonObj struct {
		UdtHash types.Hash   `json:"udt_hash"`
		Amount  *hexutil.Big `json:"amount"`
	}
	if err := json.Unmarshal(input, &jsonObj); err != nil {
		return err
	}
	*r = BurnInfo{
		UdtHash: jsonObj.UdtHash,
		Amount:  (*big.Int)(jsonObj.Amount),
	}
	return nil
}

func (r *Record) UnmarshalJSON(input []byte) error {
	type recordAlias Record
	var jsonObj struct {
		recordAlias
		Amount      *hexutil.Big   `json:"amount"`
		Occupied    *hexutil.Big   `json:"occupied"`
		BlockNumber hexutil.Uint64 `json:"block_number"`
		EpochNumber hexutil.Uint64 `json:"epoch_number"`
	}
	if err := json.Unmarshal(input, &jsonObj); err != nil {
		return err
	}
	*r = Record{
		OutPoint:    jsonObj.OutPoint,
		Ownership:   jsonObj.Ownership,
		IoType:      jsonObj.IoType,
		Amount:      (*big.Int)(jsonObj.Amount),
		Occupied:    (*big.Int)(jsonObj.Occupied),
		AssetInfo:   jsonObj.AssetInfo,
		Extra:       jsonObj.Extra,
		BlockNumber: uint64(jsonObj.BlockNumber),
		EpochNumber: uint64(jsonObj.EpochNumber),
	}
	return nil
}

func (r *TransactionInfo) UnmarshalJSON(input []byte) error {
	type transactionInfoAlias TransactionInfo
	var jsonObj struct {
		transactionInfoAlias
		Fee       hexutil.Uint64 `json:"fee"`
		Timestamp hexutil.Uint64 `json:"timestamp"`
	}

	if err := json.Unmarshal(input, &jsonObj); err != nil {
		return err
	}
	*r = TransactionInfo{
		TxHash:    jsonObj.TxHash,
		Records:   jsonObj.Records,
		Fee:       uint64(jsonObj.Fee),
		Burn:      jsonObj.Burn,
		Timestamp: uint64(jsonObj.Timestamp),
	}
	return nil
}

func (r *PaginationResponseTransactionView) UnmarshalJSON(input []byte) error {
	var jsonObj struct {
		Response   []*TransactionViewWrapper `json:"response"`
		Count      hexutil.Uint64            `json:"count,omitempty"`
		NextCursor hexutil.Uint64            `json:"next_cursor,omitempty"`
	}

	if err := json.Unmarshal(input, &jsonObj); err != nil {
		return err
	}
	*r = PaginationResponseTransactionView{
		Response:   jsonObj.Response,
		Count:      uint64(jsonObj.Count),
		NextCursor: uint64(jsonObj.NextCursor),
	}
	return nil
}

func (r *PaginationResponseTransactionInfo) UnmarshalJSON(input []byte) error {
	var jsonObj struct {
		Response   []*TransactionInfoWrapper `json:"response"`
		Count      hexutil.Uint64            `json:"count,omitempty"`
		NextCursor hexutil.Uint64            `json:"next_cursor,omitempty"`
	}

	if err := json.Unmarshal(input, &jsonObj); err != nil {
		return err
	}
	*r = PaginationResponseTransactionInfo{
		Response:   jsonObj.Response,
		Count:      uint64(jsonObj.Count),
		NextCursor: uint64(jsonObj.NextCursor),
	}
	return nil
}
