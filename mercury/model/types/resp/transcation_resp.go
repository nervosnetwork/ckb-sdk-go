package resp

import (
	"github.com/ethereum/go-ethereum/common/hexutil"
	. "github.com/nervosnetwork/ckb-sdk-go/mercury/model/types"
	"github.com/nervosnetwork/ckb-sdk-go/types"
)

type transactionResp struct {
	Version     hexutil.Uint    `json:"version"`
	Hash        types.Hash         `json:"hash"`
	CellDeps    []CellDep    `json:"cell_deps"`
	HeaderDeps  []types.Hash       `json:"header_deps"`
	Inputs      []CellInput  `json:"inputs"`
	Outputs     []CellOutput `json:"outputs"`
	OutputsData []hexutil.Bytes    `json:"outputs_data"`
	Witnesses   []hexutil.Bytes `json:"witnesses"`
}
