package model

import "github.com/nervosnetwork/ckb-sdk-go/types"

type GetBlockInfoPayload struct {
	BlockNumber uint64     `json:"block_number,omitempty"`
	BlockHash   types.Hash `json:"block_hash,omitempty"`
}
