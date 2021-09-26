package common

import (
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
