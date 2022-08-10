package types

import (
	"github.com/nervosnetwork/ckb-sdk-go/crypto/blake2b"
	"github.com/nervosnetwork/ckb-sdk-go/utils/amount"
	"math/big"
)

type ScriptHashType string
type DepType string
type TransactionStatus string

const (
	HashTypeData  ScriptHashType = "data"
	HashTypeData1 ScriptHashType = "data1"
	HashTypeType  ScriptHashType = "type"

	DepTypeCode     DepType = "code"
	DepTypeDepGroup DepType = "dep_group"

	TransactionStatusPending   TransactionStatus = "pending"
	TransactionStatusProposed  TransactionStatus = "proposed"
	TransactionStatusCommitted TransactionStatus = "committed"
	TransactionStatusUnknown   TransactionStatus = "unknown"
	TransactionStatusRejected  TransactionStatus = "rejected"
)

type Epoch struct {
	CompactTarget uint64 `json:"compact_target"`
	Length        uint64 `json:"length"`
	Number        uint64 `json:"number"`
	StartNumber   uint64 `json:"start_number"`
}

type Header struct {
	CompactTarget    uint     `json:"compact_target"`
	Dao              Hash     `json:"dao"`
	Epoch            uint64   `json:"epoch"`
	Hash             Hash     `json:"hash"`
	Nonce            *big.Int `json:"nonce"`
	Number           uint64   `json:"number"`
	ParentHash       Hash     `json:"parent_hash"`
	ProposalsHash    Hash     `json:"proposals_hash"`
	Timestamp        uint64   `json:"timestamp"`
	TransactionsRoot Hash     `json:"transactions_root"`
	ExtraHash        Hash     `json:"extra_hash"`
	Version          uint     `json:"version"`
}

type OutPoint struct {
	TxHash Hash `json:"tx_hash"`
	Index  uint `json:"index"`
}

type CellDep struct {
	OutPoint *OutPoint `json:"out_point"`
	DepType  DepType   `json:"dep_type"`
}

type Script struct {
	CodeHash Hash           `json:"code_hash"`
	HashType ScriptHashType `json:"hash_type"`
	Args     []byte         `json:"args"`
}

func (r *Script) OccupiedCapacity() uint64 {
	ckBytes := uint64(len(r.Args)) + uint64(len(r.CodeHash.Bytes())) + 1
	return amount.CkbToShannon(ckBytes)
}

func (r *Script) Hash() (Hash, error) {
	data := r.Serialize()
	hash, err := blake2b.Blake256(data)
	if err != nil {
		return Hash{}, err
	}

	return BytesToHash(hash), nil
}

func (r *Script) Equals(obj *Script) bool {
	if obj == nil {
		return false
	}

	sh, _ := r.Hash()
	oh, _ := obj.Hash()
	return sh.String() == oh.String()
}

type CellInput struct {
	Since          uint64    `json:"since"`
	PreviousOutput *OutPoint `json:"previous_output"`
}

type CellOutput struct {
	Capacity uint64  `json:"capacity"`
	Lock     *Script `json:"lock"`
	Type     *Script `json:"type"`
}

func (r CellOutput) OccupiedCapacity(outputData []byte) uint64 {
	occupiedCapacity := amount.CkbToShannon(8 + uint64(len(outputData)))
	occupiedCapacity += r.Lock.OccupiedCapacity()
	if r.Type != nil {
		occupiedCapacity += r.Type.OccupiedCapacity()
	}
	return occupiedCapacity
}

type Transaction struct {
	Version     uint          `json:"version"`
	Hash        Hash          `json:"hash"`
	CellDeps    []*CellDep    `json:"cell_deps"`
	HeaderDeps  []Hash        `json:"header_deps"`
	Inputs      []*CellInput  `json:"inputs"`
	Outputs     []*CellOutput `json:"outputs"`
	OutputsData [][]byte      `json:"outputs_data"`
	Witnesses   [][]byte      `json:"witnesses"`
}

func (t *Transaction) ComputeHash() (Hash, error) {
	data := t.SerializeWithoutWitnesses()

	hash, err := blake2b.Blake256(data)
	if err != nil {
		return Hash{}, err
	}

	return BytesToHash(hash), nil
}

func (t *Transaction) SizeInBlock() uint64 {
	b := t.Serialize()
	size := uint64(len(b)) + 4 // add header size
	return size
}

func (t *Transaction) OutputsCapacity() (totalCapacity uint64) {
	for _, output := range t.Outputs {
		totalCapacity += output.Capacity
	}
	return
}

type WitnessArgs struct {
	Lock       []byte `json:"lock"`
	InputType  []byte `json:"input_type"`
	OutputType []byte `json:"output_type"`
}

type UncleBlock struct {
	Header    *Header  `json:"header"`
	Proposals []string `json:"proposals"`
}

type Block struct {
	Header       *Header        `json:"header"`
	Proposals    []string       `json:"proposals"`
	Transactions []*Transaction `json:"transactions"`
	Uncles       []*UncleBlock  `json:"uncles"`
}

type TransactionInput struct {
	OutPoint   *OutPoint   `json:"out_point"`
	Output     *CellOutput `json:"output"`
	OutputData []byte      `json:"output_data"`
}

type Cell struct {
	BlockHash     Hash      `json:"block_hash"`
	Capacity      uint64    `json:"capacity"`
	Lock          *Script   `json:"lock"`
	OutPoint      *OutPoint `json:"out_point"`
	Type          *Script   `json:"type"`
	Cellbase      bool      `json:"cellbase,omitempty"`
	OutputDataLen uint64    `json:"output_data_len,omitempty"`
}

type CellData struct {
	Content []byte `json:"content"`
	Hash    Hash   `json:"hash"`
}

type CellInfo struct {
	Data   *CellData   `json:"data"`
	Output *CellOutput `json:"output"`
}

type CellWithStatus struct {
	Cell   *CellInfo `json:"cell"`
	Status string    `json:"status"`
}

type TxStatus struct {
	BlockHash Hash              `json:"block_hash"`
	Status    TransactionStatus `json:"status"`
}

type TransactionWithStatus struct {
	Transaction *Transaction `json:"transaction"`
	TxStatus    *TxStatus    `json:"tx_status"`
}

type BlockReward struct {
	Primary        *big.Int `json:"primary"`
	ProposalReward *big.Int `json:"proposal_reward"`
	Secondary      *big.Int `json:"secondary"`
	Total          *big.Int `json:"total"`
	TxFee          *big.Int `json:"tx_fee"`
}

type BlockEconomicState struct {
	Issuance    BlockIssuance `json:"issuance"`
	MinerReward MinerReward   `json:"miner_reward"`
	TxsFee      uint64        `json:"txs_fee"`
	FinalizedAt Hash          `json:"finalized_at"`
}

type BlockIssuance struct {
	Primary   uint64 `json:"primary"`
	Secondary uint64 `json:"secondary"`
}

type MinerReward struct {
	Primary   uint64 `json:"primary"`
	Secondary uint64 `json:"secondary"`
	Committed uint64 `json:"committed"`
	Proposal  uint64 `json:"proposal"`
}

type RationalU256 struct {
	Denom *big.Int `json:"denom"`
	Numer *big.Int `json:"numer"`
}

type ProposalWindow struct {
	Closest  uint64 `json:"closest"`
	Farthest uint64 `json:"farthest"`
}

type Consensus struct {
	Id                                   string             `json:"ID"`
	GenesisHash                          Hash               `json:"genesis_hash"`
	DaoTypeHash                          Hash               `json:"dao_type_hash"`
	Secp256k1Blake160SighashAllTypeHash  Hash               `json:"secp256k1_blake160_sighash_all_type_hash"`
	Secp256k1Blake160MultisigAllTypeHash Hash               `json:"secp256k1_blake160_multisig_all_type_hash"`
	InitialPrimaryEpochReward            uint64             `json:"initial_primary_epoch_reward"`
	SecondaryEpochReward                 uint64             `json:"secondary_epoch_reward"`
	MaxUnclesNum                         uint64             `json:"max_uncles_num"`
	OrphanRateTarget                     RationalU256       `json:"orphan_rate_target"`
	EpochDurationTarget                  uint64             `json:"epoch_duration_target"`
	TxProposalWindow                     ProposalWindow     `json:"tx_proposal_window"`
	ProposerRewardRatio                  RationalU256       `json:"proposer_reward_ratio"`
	CellbaseMaturity                     uint64             `json:"cellbase_maturity"`
	MedianTimeBlockCount                 uint64             `json:"median_time_block_count"`
	MaxBlockCycles                       uint64             `json:"max_block_cycles"`
	MaxBlockBytes                        uint64             `json:"max_block_bytes"`
	BlockVersion                         uint               `json:"block_version"`
	TxVersion                            uint               `json:"tx_version"`
	TypeIdCodeHash                       Hash               `json:"type_id_code_hash"`
	MaxBlockProposalsLimit               uint64             `json:"max_block_proposals_limit"`
	PrimaryEpochRewardHalvingInterval    uint64             `json:"primary_epoch_reward_halving_interval"`
	PermanentDifficultyInDummy           bool               `json:"permanent_difficulty_in_dummy"`
	HardforkFeatures                     []*HardForkFeature `json:"hardfork_features"`
}

type HardForkFeature struct {
	Rfc         string `json:"rfc"`
	EpochNumber uint64 `json:"epoch_number,omitempty"`
}
