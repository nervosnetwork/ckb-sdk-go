package resp

import (
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/common"
)

type TransactionInfoWrapper struct {
	TransactionInfo TransactionInfo `json:"TransactionInfo"`
}

type TransactionViewWrapper struct {
	TransactionView common.TransactionWithStatus `json:"TransactionView"`
}
