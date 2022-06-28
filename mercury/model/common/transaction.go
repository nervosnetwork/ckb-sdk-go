package common

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/nervosnetwork/ckb-sdk-go/types"
)

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
