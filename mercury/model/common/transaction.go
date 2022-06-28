package common

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/nervosnetwork/ckb-sdk-go/types"
)

type CellDep struct {
	OutPoint OutPoint      `json:"out_point"`
	DepType  types.DepType `json:"dep_type"`
}

type OutPoint struct {
	TxHash types.Hash   `json:"tx_hash"`
	Index  hexutil.Uint `json:"index"`
}

type CellInput struct {
	Since          hexutil.Uint64 `json:"since"`
	PreviousOutput OutPoint       `json:"previous_output"`
}

type CellOutput struct {
	Capacity hexutil.Uint64 `json:"capacity"`
	Lock     *Script        `json:"lock"`
	Type     *Script        `json:"type"`
}

type Script struct {
	CodeHash types.Hash           `json:"code_hash"`
	HashType types.ScriptHashType `json:"hash_type"`
	Args     hexutil.Bytes        `json:"args"`
}

type Transaction struct {
	Version     hexutil.Uint    `json:"version"`
	Hash        types.Hash      `json:"hash"`
	CellDeps    []CellDep       `json:"cell_deps"`
	HeaderDeps  []types.Hash    `json:"header_deps"`
	Inputs      []CellInput     `json:"inputs"`
	Outputs     []CellOutput    `json:"outputs"`
	OutputsData []hexutil.Bytes `json:"outputs_data"`
	Witnesses   []hexutil.Bytes `json:"witnesses"`
}

type TransactionWithRichStatus struct {
	Transaction types.Transaction `json:"transaction,omitempty"`
	TxStatus    TxRichStatus      `json:"tx_status"`
}

type TxRichStatus struct {
	Status    types.TransactionStatus `json:"status"`
	BlockHash types.Hash              `json:"block_hash,omitempty"`
	Reason    string                  `json:"reason,omitempty"`
	Timestamp uint64                  `json:"timestamp,omitempty"`
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
