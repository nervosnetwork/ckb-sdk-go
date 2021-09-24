package resp

import (
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/common"
	"github.com/nervosnetwork/ckb-sdk-go/types"
)

type GetTransactionInfoResponse struct {
	Transaction  *common.TransactionInfo `json:"transaction"`
	Status       types.TransactionStatus `json:"status"`
	RejectReason *uint8                  `json:"reject_reason"`
}
