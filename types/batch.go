package types

// refer BatchElem from go-ethereum
type BatchTransactionItem struct {
	Hash   Hash
	Result *TransactionWithStatus
	Error  error
}
