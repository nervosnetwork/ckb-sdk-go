package resp

import (
	. "github.com/nervosnetwork/ckb-sdk-go/mercury/model/types"
	"github.com/nervosnetwork/ckb-sdk-go/types"
)

type GetTransactionInfoResponse struct {
	Transaction  *TransactionInfo  `json:"transaction"`
	Status       types.TransactionStatus `json:"status"`
	RejectReason *uint8                  `json:"reject_reason"`
}
