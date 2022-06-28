package resp

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model"
	"github.com/nervosnetwork/ckb-sdk-go/types"
	"math/big"
)

func (r *Balance) UnmarshalJSON(input []byte) error {
	var jsonObj struct {
		Ownership string           `json:"ownership"`
		AssetInfo *model.AssetInfo `json:"asset_info"`
		Free      *hexutil.Big     `json:"free"`
		Occupied  *hexutil.Big     `json:"occupied"`
		Frozen    *hexutil.Big     `json:"frozen"`
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

func (r *PaginationResponseTransactionView) UnmarshalJSON(input []byte) error {
	var jsonObj struct {
		Response   []*TransactionViewWrapper `json:"response"`
		Count      *hexutil.Uint64           `json:"count,omitempty"`
		NextCursor *hexutil.Uint64           `json:"next_cursor,omitempty"`
	}
	if err := json.Unmarshal(input, &jsonObj); err != nil {
		return err
	}
	*r = PaginationResponseTransactionView{
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

func (r *ScriptGroup) UnmarshalJSON(input []byte) error {
	var jsonObj struct {
		Script        types.Script   `json:"script"`
		GroupType     GroupType      `json:"group_type"`
		InputIndices  []hexutil.Uint `json:"input_indices"`
		OutputIndices []hexutil.Uint `json:"output_indices"`
	}
	if err := json.Unmarshal(input, &jsonObj); err != nil {
		return err
	}
	toUint32Array := func(a []hexutil.Uint) []uint32 {
		result := make([]uint32, len(a))
		for i, data := range a {
			result[i] = uint32(data)
		}
		return result
	}
	*r = ScriptGroup{
		Script:        jsonObj.Script,
		GroupType:     jsonObj.GroupType,
		InputIndices:  toUint32Array(jsonObj.InputIndices),
		OutputIndices: toUint32Array(jsonObj.OutputIndices),
	}
	return nil
}
