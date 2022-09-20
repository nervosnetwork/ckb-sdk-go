package transaction

import (
	"github.com/nervosnetwork/ckb-sdk-go/types"
)

func CalculateTransactionFee(tx *types.Transaction, feeRate uint64) uint64 {
	txSize := tx.SizeInBlock()
	fee := txSize * feeRate / 1000
	if fee*1000 < txSize*feeRate {
		fee += 1
	}
	return fee
}
