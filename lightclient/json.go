package lightclient

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/nervosnetwork/ckb-sdk-go/types"
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
