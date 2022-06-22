package types

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
)

type jsonEpoch struct {
	CompactTarget hexutil.Uint64 `json:"compact_target"`
	Length        hexutil.Uint64 `json:"length"`
	Number        hexutil.Uint64 `json:"number"`
	StartNumber   hexutil.Uint64 `json:"start_number"`
}

func (r Epoch) MarshalJSON() ([]byte, error) {
	jsonObj := &jsonEpoch{
		CompactTarget: hexutil.Uint64(r.CompactTarget),
		Length:        hexutil.Uint64(r.Length),
		Number:        hexutil.Uint64(r.Number),
		StartNumber:   hexutil.Uint64(r.StartNumber),
	}
	return json.Marshal(jsonObj)
}

func (r *Epoch) UnmarshalJSON(input []byte) error {
	var jsonObj jsonEpoch
	err := json.Unmarshal(input, &jsonObj)
	if err != nil {
		return err
	}
	*r = Epoch{
		CompactTarget: uint64(jsonObj.CompactTarget),
		Length:        uint64(jsonObj.Length),
		Number:        uint64(jsonObj.Number),
		StartNumber:   uint64(jsonObj.StartNumber),
	}
	return nil
}

type headerAlias Header
type jsonHeader struct {
	headerAlias
	CompactTarget hexutil.Uint   `json:"compact_target"`
	Epoch         hexutil.Uint64 `json:"epoch"`
	Nonce         *hexutil.Big   `json:"nonce"`
	Number        hexutil.Uint64 `json:"number"`
	Timestamp     hexutil.Uint64 `json:"timestamp"`
	Version       hexutil.Uint   `json:"version"`
}

func (r Header) MarshalJSON() ([]byte, error) {
	jsonObj := &jsonHeader{
		headerAlias:   headerAlias(r),
		CompactTarget: hexutil.Uint(r.CompactTarget),
		Epoch:         hexutil.Uint64(r.Epoch),
		Nonce:         (*hexutil.Big)(r.Nonce),
		Number:        hexutil.Uint64(r.Number),
		Timestamp:     hexutil.Uint64(r.Timestamp),
		Version:       hexutil.Uint(r.Version),
	}
	return json.Marshal(jsonObj)
}

func (r *Header) UnmarshalJSON(input []byte) error {
	var jsonObj jsonHeader
	err := json.Unmarshal(input, &jsonObj)
	if err != nil {
		return err
	}
	*r = Header{
		CompactTarget:    uint(jsonObj.CompactTarget),
		Dao:              jsonObj.Dao,
		Epoch:            uint64(jsonObj.Epoch),
		Hash:             jsonObj.Hash,
		Nonce:            (*big.Int)(jsonObj.Nonce),
		Number:           uint64(jsonObj.Number),
		ParentHash:       jsonObj.ParentHash,
		ProposalsHash:    jsonObj.ProposalsHash,
		Timestamp:        uint64(jsonObj.Timestamp),
		TransactionsRoot: jsonObj.TransactionsRoot,
		ExtraHash:        jsonObj.ExtraHash,
		Version:          uint(jsonObj.Version),
	}
	return nil
}

type outPointAlias OutPoint
type jsonOutPoint struct {
	outPointAlias
	Index hexutil.Uint `json:"index"`
}

func (r OutPoint) MarshalJSON() ([]byte, error) {
	jsonObj := &jsonOutPoint{
		outPointAlias: outPointAlias(r),
		Index:         hexutil.Uint(r.Index),
	}
	return json.Marshal(jsonObj)
}

func (r *OutPoint) UnmarshalJSON(input []byte) error {
	var jsonObj jsonOutPoint
	err := json.Unmarshal(input, &jsonObj)
	if err != nil {
		return err
	}
	*r = OutPoint{
		TxHash: jsonObj.TxHash,
		Index:  uint(jsonObj.Index),
	}
	return nil
}

type scriptAlias Script
type jsonScript struct {
	scriptAlias
	Args hexutil.Bytes `json:"args"`
}

func (r Script) MarshalJSON() ([]byte, error) {
	jsonObj := &jsonScript{
		scriptAlias: scriptAlias(r),
		Args:        r.Args,
	}
	return json.Marshal(jsonObj)
}

func (r *Script) UnmarshalJSON(input []byte) error {
	var jsonObj jsonScript
	err := json.Unmarshal(input, &jsonObj)
	if err != nil {
		return err
	}
	*r = Script{
		CodeHash: jsonObj.CodeHash,
		HashType: jsonObj.HashType,
		Args:     jsonObj.Args,
	}
	return nil
}

type cellInputAlias CellInput
type jsonCellInput struct {
	cellInputAlias
	Since hexutil.Uint64 `json:"since"`
}

func (r CellInput) MarshalJSON() ([]byte, error) {
	jsonObj := &jsonCellInput{
		cellInputAlias: cellInputAlias(r),
		Since:          hexutil.Uint64(r.Since),
	}
	return json.Marshal(jsonObj)
}

func (r *CellInput) UnmarshalJSON(input []byte) error {
	var jsonObj jsonCellInput
	err := json.Unmarshal(input, &jsonObj)
	if err != nil {
		return err
	}
	*r = CellInput{
		Since:          uint64(jsonObj.Since),
		PreviousOutput: jsonObj.PreviousOutput,
	}
	return nil
}

type cellOutputAlias CellOutput
type jsonCellOutput struct {
	cellOutputAlias
	Capacity hexutil.Uint64 `json:"capacity"`
}

func (r CellOutput) MarshalJSON() ([]byte, error) {
	jsonObj := &jsonCellOutput{
		cellOutputAlias: cellOutputAlias(r),
		Capacity:        hexutil.Uint64(r.Capacity),
	}
	return json.Marshal(jsonObj)
}

func (r *CellOutput) UnmarshalJSON(input []byte) error {
	var jsonObj jsonCellOutput
	err := json.Unmarshal(input, &jsonObj)
	if err != nil {
		return err
	}
	*r = CellOutput{
		Capacity: uint64(jsonObj.Capacity),
		Lock:     jsonObj.Lock,
		Type:     jsonObj.Type,
	}
	return nil
}

type transactionAlias Transaction
type jsonTransaction struct {
	transactionAlias
	Version     hexutil.Uint    `json:"version"`
	OutputsData []hexutil.Bytes `json:"outputs_data"`
	Witnesses   []hexutil.Bytes `json:"witnesses"`
}

func (t Transaction) MarshalJSON() ([]byte, error) {
	toBytes := func(bytes [][]byte) []hexutil.Bytes {
		result := make([]hexutil.Bytes, len(bytes))
		for i, data := range bytes {
			result[i] = data
		}
		return result
	}
	jsonObj := &jsonTransaction{
		transactionAlias: transactionAlias(t),
		Version:          hexutil.Uint(t.Version),
		OutputsData:      toBytes(t.OutputsData),
		Witnesses:        toBytes(t.Witnesses),
	}
	return json.Marshal(jsonObj)
}

func (t *Transaction) UnmarshalJSON(input []byte) error {
	var jsonObj jsonTransaction
	err := json.Unmarshal(input, &jsonObj)
	if err != nil {
		return err
	}
	toByteArray := func(byteArray []hexutil.Bytes) [][]byte {
		result := make([][]byte, len(byteArray))
		for i, data := range byteArray {
			result[i] = data
		}
		return result
	}
	*t = Transaction{
		Version:     uint(jsonObj.Version),
		Hash:        jsonObj.Hash,
		CellDeps:    jsonObj.CellDeps,
		HeaderDeps:  jsonObj.HeaderDeps,
		Inputs:      jsonObj.Inputs,
		Outputs:     jsonObj.Outputs,
		OutputsData: toByteArray(jsonObj.OutputsData),
		Witnesses:   toByteArray(jsonObj.Witnesses),
	}
	return nil
}