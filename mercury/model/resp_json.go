package model

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/nervosnetwork/ckb-sdk-go/types"
	"github.com/pkg/errors"
	"math/big"
	"reflect"
)

func (r *Balance) UnmarshalJSON(input []byte) error {
	var jsonObj struct {
		Ownership string       `json:"ownership"`
		AssetInfo *AssetInfo   `json:"asset_info"`
		Free      *hexutil.Big `json:"free"`
		Occupied  *hexutil.Big `json:"occupied"`
		Frozen    *hexutil.Big `json:"frozen"`
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

func (r *BlockInfo) UnmarshalJSON(input []byte) error {
	type blockInfoAlias BlockInfo
	var jsonObj struct {
		blockInfoAlias
		BlockNumber hexutil.Uint64 `json:"block_number"`
		Timestamp   hexutil.Uint64 `json:"timestamp"`
	}
	if err := json.Unmarshal(input, &jsonObj); err != nil {
		return err
	}
	*r = BlockInfo{
		BlockNumber:  uint64(jsonObj.BlockNumber),
		BlockHash:    jsonObj.BlockHash,
		ParentHash:   jsonObj.ParentHash,
		Timestamp:    uint64(jsonObj.Timestamp),
		Transactions: jsonObj.Transactions,
	}
	return nil
}

func (r *PaginationResponseTransactionWithRichStatus) UnmarshalJSON(input []byte) error {
	var jsonObj struct {
		Response   []*TransactionWithRichStatusWrapper `json:"response"`
		Count      *hexutil.Uint64                     `json:"count,omitempty"`
		NextCursor *hexutil.Uint64                     `json:"next_cursor,omitempty"`
	}
	if err := json.Unmarshal(input, &jsonObj); err != nil {
		return err
	}
	*r = PaginationResponseTransactionWithRichStatus{
		Response: jsonObj.Response,
	}
	if jsonObj.Count != nil {
		r.Count = uint64(*jsonObj.Count)
	}
	if jsonObj.NextCursor != nil {
		r.NextCursor = uint64(*jsonObj.NextCursor)
	}
	return nil
}

func (r *PaginationResponseTransactionInfo) UnmarshalJSON(input []byte) error {
	var jsonObj struct {
		Response   []*TransactionInfoWrapper `json:"response"`
		Count      *hexutil.Uint64           `json:"count,omitempty"`
		NextCursor *hexutil.Uint64           `json:"next_cursor,omitempty"`
	}
	if err := json.Unmarshal(input, &jsonObj); err != nil {
		return err
	}
	*r = PaginationResponseTransactionInfo{
		Response: jsonObj.Response,
	}
	if jsonObj.Count != nil {
		r.Count = uint64(*jsonObj.Count)
	}
	if jsonObj.NextCursor != nil {
		r.NextCursor = uint64(*jsonObj.NextCursor)
	}
	return nil
}

func (r *AccountInfo) UnmarshalJSON(input []byte) error {
	var jsonObj struct {
		AccountNumber  hexutil.Uint `json:"account_number"`
		AccountAddress string       `json:"account_address"`
		AccountType    AccountType  `json:"account_type"`
	}
	if err := json.Unmarshal(input, &jsonObj); err != nil {
		return err
	}
	*r = AccountInfo{
		AccountNumber:  uint32(jsonObj.AccountNumber),
		AccountAddress: jsonObj.AccountAddress,
		AccountType:    jsonObj.AccountType,
	}
	return nil
}

func (r *TxRichStatus) UnmarshalJSON(input []byte) error {
	type txRichStatusAlias TxRichStatus
	var jsonObj struct {
		txRichStatusAlias
		Timestamp hexutil.Uint64 `json:"timestamp,omitempty"`
	}
	if err := json.Unmarshal(input, &jsonObj); err != nil {
		return err
	}
	*r = TxRichStatus{
		Status:    jsonObj.Status,
		BlockHash: jsonObj.BlockHash,
		Reason:    jsonObj.Reason,
		Timestamp: uint64(jsonObj.Timestamp),
	}
	return nil
}

func (e *DaoState) UnmarshalJSON(input []byte) error {
	var data map[string]interface{}
	if err := json.Unmarshal(input, &data); err != nil {
		return err
	}
	switch reflect.ValueOf(data["value"]).Kind() {
	case reflect.String:
		var jsonObj struct {
			Type  DaoStateType   `json:"type"`
			Value hexutil.Uint64 `json:"value"`
		}
		if err := json.Unmarshal(input, &jsonObj); err != nil {
			return err
		}
		*e = DaoState{
			Type:  jsonObj.Type,
			Value: []uint64{uint64(jsonObj.Value)},
		}
	case reflect.Slice:
		var jsonObj struct {
			Type  DaoStateType     `json:"type"`
			Value []hexutil.Uint64 `json:"value"`
		}
		if err := json.Unmarshal(input, &jsonObj); err != nil {
			return err
		}
		*e = DaoState{
			Type:  jsonObj.Type,
			Value: []uint64{uint64(jsonObj.Value[0]), uint64(jsonObj.Value[1])},
		}
	default:
		return errors.New("invalid type while unmarshal DaoState")
	}
	return nil
}
