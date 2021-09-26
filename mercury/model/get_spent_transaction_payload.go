package model

import "github.com/nervosnetwork/ckb-sdk-go/mercury/model/common"

type GetSpentTransactionPayload struct {
	OutPoint      common.OutPoint `json:"outpoint"`
	StructureType StructureType   `json:"structure_type"`
}

type StructureType string

const (
	Native      StructureType = "Native"
	DoubleEntry StructureType = "DoubleEntry"
)
