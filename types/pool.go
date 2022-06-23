package types

type TxPoolInfo struct {
	TipHash          Hash   `json:"tip_hash"`
	TipNumber        uint64 `json:"tip_number"`
	Pending          uint64 `json:"pending"`
	Proposed         uint64 `json:"proposed"`
	Orphan           uint64 `json:"orphan"`
	TotalTxSize      uint64 `json:"total_tx_size"`
	TotalTxCycles    uint64 `json:"total_tx_cycles"`
	MinFeeRate       uint64 `json:"min_fee_rate"`
	LastTxsUpdatedAt uint64 `json:"last_txs_updated_at"`
}
