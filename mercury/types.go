package mercury

import (
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/resp"
	"github.com/nervosnetwork/ckb-sdk-go/types"
)

// -----------------------------------------------------------------------------------------------------------------------------------------------

type rpcTransactionInfoWithStatusResponse struct {
	Transaction     rpcTransactionInfoResponse `json:"transaction"`
	Status          types.TransactionStatus    `json:"status"`
	BlockHash       string                     `json:"block_hash"`
	BlockNumber     uint64                     `json:"block_number"`
	ConfirmedNumber uint64                     `json:"confirmed_number"`
}

type rpcTransactionInfoResponse struct {
	TxHash     string               `json:"tx_hash"`
	Operations []*rpcRecordResponse `json:"operations"`
}

type rpcRecordResponse struct {
	Id            uint          `json:"id"`
	KeyAddress    string        `json:"key_address"`
	NormalAddress string        `json:"normal_address"`
	Amount        rpcAmountResp `json:"amount"`
}

type rpcAmountResp struct {
	Value   string      `json:"value"`
	UdtHash string      `json:"udt_hash"`
	Status  interface{} `json:"status"`
}

func toTransactionInfoWithStatusResponse(tx *rpcTransactionInfoWithStatusResponse) (*resp.GetTransactionInfoResponse, error) {
	result := &resp.GetTransactionInfoResponse{
		// Status:          tx.Status,
		// BlockHash:       tx.BlockHash,
		// BlockNumber:     tx.BlockNumber,
		// ConfirmedNumber: tx.ConfirmedNumber,
		// Transaction:     transactionInfoResponse,
	}

	return result, nil
}

// -----------------------------------------------------------------------------------------------------------------------------------------------
type rpcBlockInfoResponse struct {
	BlockNumber     uint64                        `json:"block_number"`
	BlockHash       string                        `json:"block_hash"`
	ParentBlockHash string                        `json:"parent_block_hash"`
	Timestamp       uint64                        `json:"timestamp"`
	Transactions    []*rpcTransactionInfoResponse `json:"transactions"`
}
