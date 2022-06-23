package resp

import "github.com/nervosnetwork/ckb-sdk-go/mercury/model/common"

type TransactionInfoWrapper struct {
	Type  TransactionType  `json:"type"`
	Value *TransactionInfo `json:"value"`
}

type TransactionViewWrapper struct {
	Type  TransactionType                   `json:"type"`
	Value *common.TransactionWithRichStatus `json:"value"`
}

type TransactionType string

const (
	TransactionTransactionView TransactionType = "TransactionWithRichStatus"
	TransactionTransactionInfo TransactionType = "TransactionInfo"
)
