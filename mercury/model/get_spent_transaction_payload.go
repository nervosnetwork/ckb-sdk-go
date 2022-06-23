package model

import (
	"github.com/nervosnetwork/ckb-sdk-go/types"
)

type GetSpentTransactionPayload struct {
	OutPoint      types.OutPoint `json:"outpoint"`
	StructureType StructureType  `json:"structure_type"`
}

type StructureType string

const (
	Native      StructureType = "Native"
	DoubleEntry StructureType = "DoubleEntry"
)
