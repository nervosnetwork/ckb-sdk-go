package lightclient

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/nervosnetwork/ckb-sdk-go/v2/indexer"
	"github.com/nervosnetwork/ckb-sdk-go/v2/types"
)

type scriptDetailAlias ScriptDetail
type jsonScriptDetail struct {
	scriptDetailAlias
	BlockNumber hexutil.Uint64 `json:"block_number"`
}

func (r *ScriptDetail) UnmarshalJSON(input []byte) error {
	var jsonObj jsonScriptDetail
	if err := json.Unmarshal(input, &jsonObj); err != nil {
		return err
	}
	*r = ScriptDetail{
		Script:      jsonObj.Script,
		ScriptType:  jsonObj.ScriptType,
		BlockNumber: uint64(jsonObj.BlockNumber),
	}
	return nil
}

func (r ScriptDetail) MarshalJSON() ([]byte, error) {
	jsonObj := &jsonScriptDetail{
		scriptDetailAlias: scriptDetailAlias(r),
		BlockNumber:       hexutil.Uint64(r.BlockNumber),
	}
	return json.Marshal(jsonObj)
}

func (r *FetchedHeader) UnmarshalJSON(input []byte) error {
	var jsonObj struct {
		Status    FetchStatus    `json:"status"`
		Data      *types.Header  `json:"data"`
		FirstSent hexutil.Uint64 `json:"first_sent"`
		TimeStamp hexutil.Uint64 `json:"time_stamp"`
	}
	if err := json.Unmarshal(input, &jsonObj); err != nil {
		return err
	}
	*r = FetchedHeader{
		Status:    jsonObj.Status,
		Data:      jsonObj.Data,
		FirstSent: uint64(jsonObj.FirstSent),
		TimeStamp: uint64(jsonObj.TimeStamp),
	}
	return nil
}

func (r *FetchedTransaction) UnmarshalJSON(input []byte) error {
	var jsonObj struct {
		Status    FetchStatus            `json:"status"`
		Data      *TransactionWithHeader `json:"data"`
		FirstSent hexutil.Uint64         `json:"first_sent"`
		TimeStamp hexutil.Uint64         `json:"time_stamp"`
	}
	if err := json.Unmarshal(input, &jsonObj); err != nil {
		return err
	}
	*r = FetchedTransaction{
		Status:    jsonObj.Status,
		Data:      jsonObj.Data,
		FirstSent: uint64(jsonObj.FirstSent),
		TimeStamp: uint64(jsonObj.TimeStamp),
	}
	return nil
}

type jsonTxWithCell struct {
	BlockNumber hexutil.Uint64     `json:"block_number"`
	IoIndex     hexutil.Uint       `json:"io_index"`
	IoType      indexer.IoType     `json:"io_type"`
	Transaction *types.Transaction `json:"transaction"`
	TxIndex     hexutil.Uint       `json:"tx_index"`
}

func (r *TxWithCell) UnmarshalJSON(input []byte) error {
	var jsonObj jsonTxWithCell
	if err := json.Unmarshal(input, &jsonObj); err != nil {
		return err
	}
	*r = TxWithCell{
		BlockNumber: uint64(jsonObj.BlockNumber),
		IoIndex:     uint(jsonObj.IoIndex),
		IoType:      jsonObj.IoType,
		Transaction: jsonObj.Transaction,
		TxIndex:     uint(jsonObj.TxIndex),
	}
	return nil
}

type jsonTxWithCells struct {
	Transaction *types.Transaction `json:"transaction"`
	BlockNumber hexutil.Uint64     `json:"block_number"`
	TxIndex     hexutil.Uint       `json:"tx_index"`
	Cells       []*indexer.Cell    `json:"Cells"`
}

func (r *TxWithCells) UnmarshalJSON(input []byte) error {
	var jsonObj jsonTxWithCells
	if err := json.Unmarshal(input, &jsonObj); err != nil {
		return err
	}
	*r = TxWithCells{
		BlockNumber: uint64(jsonObj.BlockNumber),
		Transaction: jsonObj.Transaction,
		TxIndex:     uint(jsonObj.TxIndex),
		Cells:       jsonObj.Cells,
	}
	return nil
}
