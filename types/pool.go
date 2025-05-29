package types

import "encoding/json"

// /// The uncle block template of the new block for miners.
// #[derive(Clone, Default, Serialize, Deserialize, PartialEq, Eq, Hash, Debug, JsonSchema)]
// pub struct UncleTemplate {
//     /// The uncle block hash.
//     pub hash: H256,
//     /// Whether miners must include this uncle in the submit block.
//     pub required: bool,
//     /// The proposals of the uncle block.
//     ///
//     /// Miners must keep this unchanged when including this uncle in the new block.
//     pub proposals: Vec<ProposalShortId>,
//     /// The header of the uncle block.
//     ///
//     /// Miners must keep this unchanged when including this uncle in the new block.
//     pub header: Header,
// }

type UncleTemplate struct {
	Hash      Hash     `json:"hash"`
	Required  bool     `json:"required"`
	Proposals []string `json:"proposals"`
	Header    *Header  `json:"header"`
}

type BlockTemplate struct {
	Version          uint32                `json:"version"`
	CompactTarget    uint32                `json:"compact_target"`
	CurrentTime      uint64                `json:"current_time"`
	Number           uint64                `json:"number"`
	Epoch            uint64                `json:"epoch"`
	ParentHash       Hash                  `json:"parent_hash"`
	CyclesLimit      uint64                `json:"cycles_limit"`
	BytesLimit       uint64                `json:"bytes_limit"`
	UnclesCountLimit uint64                `json:"uncles_count_limit"`
	Uncles           []UncleTemplate       `json:"uncles"`
	Transactions     []TransactionTemplate `json:"transactions"`
	Proposals        []string              `json:"proposals"`
	Cellbase         CellbaseTemplate      `json:"cellbase"`
	WorkId           uint64                `json:"work_id"`
	Dao              Hash                  `json:"dao"`
	Extension        *json.RawMessage      `json:"extension"`
}

type CellbaseTemplate struct {
	Hash   Hash        `json:"hash"`
	Cycles *uint64     `json:"cycles"`
	Data   Transaction `json:"data"`
}

type TransactionTemplate struct {
	Hash     Hash        `json:"hash"`
	Required bool        `json:"required"`
	Cycles   *uint64     `json:"cycles"`
	Depends  *[]uint64   `json:"depends"`
	Data     Transaction `json:"data"`
}

type TxPoolInfo struct {
	LastTxsUpdatedAt uint64 `json:"last_txs_updated_at"`
	MaxTxPoolSize    uint64 `json:"max_tx_pool_size"`
	MinFeeRate       uint64 `json:"min_fee_rate"`
	MinRbfRate       uint64 `json:"min_fee_rate"`
	Orphan           uint64 `json:"orphan"`
	Pending          uint64 `json:"pending"`
	Proposed         uint64 `json:"proposed"`
	TipHash          Hash   `json:"tip_hash"`
	TipNumber        uint64 `json:"tip_number"`
	TotalTxCycles    uint64 `json:"total_tx_cycles"`
	TotalTxSize      uint64 `json:"total_tx_size"`
	TxSizeLimit      uint64 `json:"tx_size_limit"`
	VerifyQueueSize  uint64 `json:"verify_queue_size"`
}

type RawTxPool struct {
	Pending  []Hash `json:"pending"`
	Proposed []Hash `json:"proposed"`
}

type AncestorsScoreSortKey struct {
	AncestorsFee    uint64 `json:"ancestors_fee"`
	AncestorsWeight uint64 `json:"ancestors_weight"`
	Fee             uint64 `json:"fee"`
	Weight          uint64 `json:"weight"`
}

type PoolTxDetailInfo struct {
	AncestorsCount   uint64                `json:"ancestors_count"`
	DescendantsCount uint64                `json:"descendants_count"`
	EntryStatus      string                `json:"entry_status"`
	PendingCount     uint64                `json:"pending_count"`
	ProposedCount    uint64                `json:"proposed_count"`
	RankInPending    uint64                `json:"rank_in_pending"`
	ScoreSortKey     AncestorsScoreSortKey `json:"score_sortkey"`
	Timestamp        uint64                `json:"timestamp"`
}

type EntryCompleted struct {
	// Cached tx cycles
	Cycles uint64 `json:"cycles"`
	// Cached tx fee
	Fee uint64 `json:"fee"`
}
