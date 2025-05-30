package types

import (
	"encoding/json"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

func (r *ScriptType) UnmarshalJSON(input []byte) error {
	var jsonObj string
	if err := json.Unmarshal(input, &jsonObj); err != nil {
		return err
	}
	switch strings.ToLower(jsonObj) {
	case strings.ToLower(string(ScriptTypeLock)):
		*r = ScriptTypeLock
	case strings.ToLower(string(ScriptTypeType)):
		*r = ScriptTypeType
	default:
		return fmt.Errorf("can't unmarshal json from unknown script type %s", input)
	}
	return nil
}

type jsonEpoch struct {
	CompactTarget hexutil.Uint64 `json:"compact_target"`
	Length        hexutil.Uint64 `json:"length"`
	Number        hexutil.Uint64 `json:"number"`
	StartNumber   hexutil.Uint64 `json:"start_number"`
}

func (r Epoch) MarshalJSON() ([]byte, error) {
	jsonObj := &jsonEpoch{
		CompactTarget: hexutil.Uint64(r.CompactTarget),
		Length:        hexutil.Uint64(r.Length),
		Number:        hexutil.Uint64(r.Number),
		StartNumber:   hexutil.Uint64(r.StartNumber),
	}
	return json.Marshal(jsonObj)
}

func (r *Epoch) UnmarshalJSON(input []byte) error {
	var jsonObj jsonEpoch
	if err := json.Unmarshal(input, &jsonObj); err != nil {
		return err
	}
	*r = Epoch{
		CompactTarget: uint64(jsonObj.CompactTarget),
		Length:        uint64(jsonObj.Length),
		Number:        uint64(jsonObj.Number),
		StartNumber:   uint64(jsonObj.StartNumber),
	}
	return nil
}

type headerAlias Header
type jsonHeader struct {
	headerAlias
	CompactTarget hexutil.Uint   `json:"compact_target"`
	Epoch         hexutil.Uint64 `json:"epoch"`
	Nonce         *hexutil.Big   `json:"nonce"`
	Number        hexutil.Uint64 `json:"number"`
	Timestamp     hexutil.Uint64 `json:"timestamp"`
	Version       hexutil.Uint   `json:"version"`
}

func (r Header) MarshalJSON() ([]byte, error) {
	jsonObj := &jsonHeader{
		headerAlias:   headerAlias(r),
		CompactTarget: hexutil.Uint(r.CompactTarget),
		Epoch:         hexutil.Uint64(r.Epoch),
		Nonce:         (*hexutil.Big)(r.Nonce),
		Number:        hexutil.Uint64(r.Number),
		Timestamp:     hexutil.Uint64(r.Timestamp),
		Version:       hexutil.Uint(r.Version),
	}
	return json.Marshal(jsonObj)
}

func (r *Header) UnmarshalJSON(input []byte) error {
	var jsonObj jsonHeader
	if err := json.Unmarshal(input, &jsonObj); err != nil {
		return err
	}
	*r = Header{
		CompactTarget:    uint32(jsonObj.CompactTarget),
		Dao:              jsonObj.Dao,
		Epoch:            uint64(jsonObj.Epoch),
		Hash:             jsonObj.Hash,
		Nonce:            (*big.Int)(jsonObj.Nonce),
		Number:           uint64(jsonObj.Number),
		ParentHash:       jsonObj.ParentHash,
		ProposalsHash:    jsonObj.ProposalsHash,
		Timestamp:        uint64(jsonObj.Timestamp),
		TransactionsRoot: jsonObj.TransactionsRoot,
		ExtraHash:        jsonObj.ExtraHash,
		Version:          uint32(jsonObj.Version),
	}
	return nil
}

type outPointAlias OutPoint
type jsonOutPoint struct {
	outPointAlias
	Index hexutil.Uint `json:"index"`
}

func (r OutPoint) MarshalJSON() ([]byte, error) {
	jsonObj := &jsonOutPoint{
		outPointAlias: outPointAlias(r),
		Index:         hexutil.Uint(r.Index),
	}
	return json.Marshal(jsonObj)
}

func (r *OutPoint) UnmarshalJSON(input []byte) error {
	var jsonObj jsonOutPoint
	if err := json.Unmarshal(input, &jsonObj); err != nil {
		return err
	}
	*r = OutPoint{
		TxHash: jsonObj.TxHash,
		Index:  uint32(jsonObj.Index),
	}
	return nil
}

type scriptAlias Script
type jsonScript struct {
	scriptAlias
	Args hexutil.Bytes `json:"args"`
}

func (r Script) MarshalJSON() ([]byte, error) {
	jsonObj := &jsonScript{
		scriptAlias: scriptAlias(r),
		Args:        r.Args,
	}
	return json.Marshal(jsonObj)
}

func (r *Script) UnmarshalJSON(input []byte) error {
	var jsonObj jsonScript
	if err := json.Unmarshal(input, &jsonObj); err != nil {
		return err
	}
	*r = Script{
		CodeHash: jsonObj.CodeHash,
		HashType: jsonObj.HashType,
		Args:     jsonObj.Args,
	}
	return nil
}

type cellInputAlias CellInput
type jsonCellInput struct {
	cellInputAlias
	Since hexutil.Uint64 `json:"since"`
}

func (r CellInput) MarshalJSON() ([]byte, error) {
	jsonObj := &jsonCellInput{
		cellInputAlias: cellInputAlias(r),
		Since:          hexutil.Uint64(r.Since),
	}
	return json.Marshal(jsonObj)
}

func (r *CellInput) UnmarshalJSON(input []byte) error {
	var jsonObj jsonCellInput
	if err := json.Unmarshal(input, &jsonObj); err != nil {
		return err
	}
	*r = CellInput{
		Since:          uint64(jsonObj.Since),
		PreviousOutput: jsonObj.PreviousOutput,
	}
	return nil
}

type cellOutputAlias CellOutput
type jsonCellOutput struct {
	cellOutputAlias
	Capacity hexutil.Uint64 `json:"capacity"`
}

func (r CellOutput) MarshalJSON() ([]byte, error) {
	jsonObj := &jsonCellOutput{
		cellOutputAlias: cellOutputAlias(r),
		Capacity:        hexutil.Uint64(r.Capacity),
	}
	return json.Marshal(jsonObj)
}

func (r *CellOutput) UnmarshalJSON(input []byte) error {
	var jsonObj jsonCellOutput
	if err := json.Unmarshal(input, &jsonObj); err != nil {
		return err
	}
	*r = CellOutput{
		Capacity: uint64(jsonObj.Capacity),
		Lock:     jsonObj.Lock,
		Type:     jsonObj.Type,
	}
	return nil
}

type jsonTransaction struct {
	Version     hexutil.Uint    `json:"version"`
	CellDeps    []*CellDep      `json:"cell_deps"`
	HeaderDeps  []Hash          `json:"header_deps"`
	Inputs      []*CellInput    `json:"inputs"`
	Outputs     []*CellOutput   `json:"outputs"`
	OutputsData []hexutil.Bytes `json:"outputs_data"`
	Witnesses   []hexutil.Bytes `json:"witnesses"`
}

func (t Transaction) MarshalJSON() ([]byte, error) {
	toBytesArray := func(bytes [][]byte) []hexutil.Bytes {
		result := make([]hexutil.Bytes, len(bytes))
		for i, data := range bytes {
			result[i] = data
		}
		return result
	}
	jsonObj := &jsonTransaction{
		Version:     hexutil.Uint(t.Version),
		CellDeps:    t.CellDeps,
		HeaderDeps:  t.HeaderDeps,
		Inputs:      t.Inputs,
		Outputs:     t.Outputs,
		OutputsData: toBytesArray(t.OutputsData),
		Witnesses:   toBytesArray(t.Witnesses),
	}
	if jsonObj.HeaderDeps == nil {
		jsonObj.HeaderDeps = make([]Hash, 0)
	}
	return json.Marshal(jsonObj)
}

func (t *Transaction) UnmarshalJSON(input []byte) error {
	var jsonObj struct {
		jsonTransaction
		Hash Hash `json:"hash"`
	}
	if err := json.Unmarshal(input, &jsonObj); err != nil {
		return err
	}
	toBytesArray := func(byteArray []hexutil.Bytes) [][]byte {
		result := make([][]byte, len(byteArray))
		for i, data := range byteArray {
			result[i] = data
		}
		return result
	}
	*t = Transaction{
		Version:     uint32(jsonObj.Version),
		Hash:        jsonObj.Hash,
		CellDeps:    jsonObj.CellDeps,
		HeaderDeps:  jsonObj.HeaderDeps,
		Inputs:      jsonObj.Inputs,
		Outputs:     jsonObj.Outputs,
		OutputsData: toBytesArray(jsonObj.OutputsData),
		Witnesses:   toBytesArray(jsonObj.Witnesses),
	}
	return nil
}

func (r *TransactionStatus) UnmarshalJSON(input []byte) error {
	var jsonObj string
	if err := json.Unmarshal(input, &jsonObj); err != nil {
		return err
	}
	switch strings.ToLower(jsonObj) {
	case "":
		*r = ""
	case strings.ToLower(string(TransactionStatusPending)):
		*r = TransactionStatusPending
	case strings.ToLower(string(TransactionStatusProposed)):
		*r = TransactionStatusProposed
	case strings.ToLower(string(TransactionStatusCommitted)):
		*r = TransactionStatusCommitted
	case strings.ToLower(string(TransactionStatusUnknown)):
		*r = TransactionStatusUnknown
	case strings.ToLower(string(TransactionStatusRejected)):
		*r = TransactionStatusRejected
	default:
		return fmt.Errorf("can't unmarshal json from unknown transaction status value %s", jsonObj)
	}
	return nil
}

type jsonCellData struct {
	Content hexutil.Bytes `json:"content"`
	Hash    Hash          `json:"hash"`
}

func (r CellData) MarshalJSON() ([]byte, error) {
	jsonObj := &jsonCellData{
		Content: r.Content,
		Hash:    r.Hash,
	}
	return json.Marshal(jsonObj)
}

func (r *CellData) UnmarshalJSON(input []byte) error {
	var jsonObj jsonCellData
	if err := json.Unmarshal(input, &jsonObj); err != nil {
		return err
	}
	*r = CellData{
		Content: jsonObj.Content,
		Hash:    jsonObj.Hash,
	}
	return nil
}

type jsonRationalU256 struct {
	Denom hexutil.Big `json:"denom"`
	Numer hexutil.Big `json:"numer"`
}

type jsonHardForkFeature struct {
	Rfc         string          `json:"rfc"`
	EpochNumber *hexutil.Uint64 `json:"epoch_number,omitempty"`
}

type consensusAlias Consensus
type jsonConsensus struct {
	consensusAlias
	InitialPrimaryEpochReward hexutil.Uint64   `json:"initial_primary_epoch_reward"`
	SecondaryEpochReward      hexutil.Uint64   `json:"secondary_epoch_reward"`
	MaxUnclesNum              hexutil.Uint64   `json:"max_uncles_num"`
	OrphanRateTarget          jsonRationalU256 `json:"orphan_rate_target"`
	EpochDurationTarget       hexutil.Uint64   `json:"epoch_duration_target"`
	TxProposalWindow          struct {
		Closest  hexutil.Uint64 `json:"closest"`
		Farthest hexutil.Uint64 `json:"farthest"`
	} `json:"tx_proposal_window"`
	ProposerRewardRatio               jsonRationalU256       `json:"proposer_reward_ratio"`
	CellbaseMaturity                  hexutil.Uint64         `json:"cellbase_maturity"`
	MedianTimeBlockCount              hexutil.Uint64         `json:"median_time_block_count"`
	MaxBlockCycles                    hexutil.Uint64         `json:"max_block_cycles"`
	MaxBlockBytes                     hexutil.Uint64         `json:"max_block_bytes"`
	BlockVersion                      hexutil.Uint           `json:"block_version"`
	TxVersion                         hexutil.Uint           `json:"tx_version"`
	MaxBlockProposalsLimit            hexutil.Uint64         `json:"max_block_proposals_limit"`
	PrimaryEpochRewardHalvingInterval hexutil.Uint64         `json:"primary_epoch_reward_halving_interval"`
	PermanentDifficultyInDummy        bool                   `json:"permanent_difficulty_in_dummy"`
	HardforkFeatures                  []*jsonHardForkFeature `json:"hardfork_features"`
}

func (r *Consensus) UnmarshalJSON(input []byte) error {
	var jsonObj jsonConsensus
	if err := json.Unmarshal(input, &jsonObj); err != nil {
		return err
	}
	toHardForkFeatureArray := func(a []*jsonHardForkFeature) HardForkFeatures {
		result := make(map[string]*HardForkFeature)
		for _, data := range a {
			result[data.Rfc] = &HardForkFeature{
				Rfc:         data.Rfc,
				EpochNumber: (*uint64)(data.EpochNumber),
			}
		}
		return result
	}
	*r = Consensus{
		Id:                                   jsonObj.Id,
		GenesisHash:                          jsonObj.GenesisHash,
		DaoTypeHash:                          jsonObj.DaoTypeHash,
		Secp256k1Blake160SighashAllTypeHash:  jsonObj.Secp256k1Blake160SighashAllTypeHash,
		Secp256k1Blake160MultisigAllTypeHash: jsonObj.Secp256k1Blake160MultisigAllTypeHash,
		InitialPrimaryEpochReward:            uint64(jsonObj.InitialPrimaryEpochReward),
		SecondaryEpochReward:                 uint64(jsonObj.SecondaryEpochReward),
		MaxUnclesNum:                         uint64(jsonObj.MaxUnclesNum),
		OrphanRateTarget: RationalU256{
			Denom: (*big.Int)(&jsonObj.OrphanRateTarget.Denom),
			Numer: (*big.Int)(&jsonObj.OrphanRateTarget.Numer),
		},
		EpochDurationTarget: uint64(jsonObj.EpochDurationTarget),
		TxProposalWindow: ProposalWindow{
			Closest:  uint64(jsonObj.TxProposalWindow.Closest),
			Farthest: uint64(jsonObj.TxProposalWindow.Farthest),
		},
		ProposerRewardRatio: RationalU256{
			Denom: (*big.Int)(&jsonObj.ProposerRewardRatio.Denom),
			Numer: (*big.Int)(&jsonObj.ProposerRewardRatio.Numer),
		},
		CellbaseMaturity:                  uint64(jsonObj.CellbaseMaturity),
		MedianTimeBlockCount:              uint64(jsonObj.MedianTimeBlockCount),
		MaxBlockCycles:                    uint64(jsonObj.MaxBlockCycles),
		MaxBlockBytes:                     uint64(jsonObj.MaxBlockBytes),
		BlockVersion:                      uint32(jsonObj.BlockVersion),
		TxVersion:                         uint32(jsonObj.TxVersion),
		TypeIdCodeHash:                    jsonObj.TypeIdCodeHash,
		MaxBlockProposalsLimit:            uint64(jsonObj.MaxBlockProposalsLimit),
		PrimaryEpochRewardHalvingInterval: uint64(jsonObj.PrimaryEpochRewardHalvingInterval),
		PermanentDifficultyInDummy:        jsonObj.PermanentDifficultyInDummy,
		HardforkFeatures:                  toHardForkFeatureArray(jsonObj.HardforkFeatures),
	}
	return nil
}

type jsonSyncState struct {
	Ibd                     bool           `json:"ibd"`
	BestKnownBlockNumber    hexutil.Uint64 `json:"best_known_block_number"`
	BestKnownBlockTimestamp hexutil.Uint64 `json:"best_known_block_timestamp"`
	OrphanBlocksCount       hexutil.Uint64 `json:"orphan_blocks_count"`
	InflightBlocksCount     hexutil.Uint64 `json:"inflight_blocks_count"`
	FastTime                hexutil.Uint64 `json:"fast_time"`
	LowTime                 hexutil.Uint64 `json:"low_time"`
	NormalTime              hexutil.Uint64 `json:"normal_time"`
	TipHash                 Hash           `json:"tip_hash"`
	TipNumber               hexutil.Uint64 `json:"tip_number"`
	UnverifiedTipHash       Hash           `json:"unverified_tip_hash"`
	UnverifiedTipNumber     hexutil.Uint64 `json:"unverified_tip_number"`
}

func (t *SyncState) UnmarshalJSON(input []byte) error {
	var jsonObj jsonSyncState
	if err := json.Unmarshal(input, &jsonObj); err != nil {
		return err
	}
	*t = SyncState{
		Ibd:                     jsonObj.Ibd,
		BestKnownBlockNumber:    uint64(jsonObj.BestKnownBlockNumber),
		BestKnownBlockTimestamp: uint64(jsonObj.BestKnownBlockTimestamp),
		OrphanBlocksCount:       uint64(jsonObj.OrphanBlocksCount),
		InflightBlocksCount:     uint64(jsonObj.InflightBlocksCount),
		FastTime:                uint64(jsonObj.FastTime),
		LowTime:                 uint64(jsonObj.LowTime),
		NormalTime:              uint64(jsonObj.NormalTime),
		TipHash:                 jsonObj.TipHash,
		TipNumber:               uint64(jsonObj.TipNumber),
		UnverifiedTipHash:       jsonObj.UnverifiedTipHash,
		UnverifiedTipNumber:     uint64(jsonObj.UnverifiedTipNumber),
	}
	return nil
}

type jsonProof struct {
	Indices []hexutil.Uint `json:"indices"`
	Lemmas  []Hash         `json:"lemmas"`
}

func (r Proof) MarshalJSON() ([]byte, error) {
	indices := make([]hexutil.Uint, len(r.Indices))
	for i, v := range r.Indices {
		indices[i] = hexutil.Uint(v)
	}
	jsonObj := &jsonProof{
		Indices: indices,
		Lemmas:  r.Lemmas,
	}
	return json.Marshal(jsonObj)
}

func (r *Proof) UnmarshalJSON(input []byte) error {
	var jsonObj jsonProof
	if err := json.Unmarshal(input, &jsonObj); err != nil {
		return err
	}
	indices := make([]uint, len(jsonObj.Indices))
	for i, v := range jsonObj.Indices {
		indices[i] = uint(v)
	}
	*r = Proof{
		Indices: indices,
		Lemmas:  jsonObj.Lemmas,
	}
	return nil
}

func (r *RemoteNodeProtocol) UnmarshalJSON(input []byte) error {
	var jsonObj struct {
		ID      hexutil.Uint64 `json:"id"`
		Version string         `json:"version"`
	}

	if err := json.Unmarshal(input, &jsonObj); err != nil {
		return err
	}
	*r = RemoteNodeProtocol{
		ID:      uint64(jsonObj.ID),
		Version: jsonObj.Version,
	}
	return nil
}

func (r *PeerSyncState) UnmarshalJSON(input []byte) error {
	type PeerSyncStateAlias PeerSyncState
	var jsonObj struct {
		PeerSyncStateAlias
		BestKnownHeaderNumber  *hexutil.Uint64 `json:"best_known_header_number,omitempty"`
		LastCommonHeaderNumber *hexutil.Uint64 `json:"last_common_header_number,omitempty"`
		UnknownHeaderListSize  hexutil.Uint64  `json:"unknown_header_list_size"`
		InflightCount          hexutil.Uint64  `json:"inflight_count"`
		CanFetchCount          hexutil.Uint64  `json:"can_fetch_count"`
	}
	if err := json.Unmarshal(input, &jsonObj); err != nil {
		return err
	}
	*r = PeerSyncState{
		BestKnownHeaderHash:    jsonObj.BestKnownHeaderHash,
		BestKnownHeaderNumber:  (*uint64)(jsonObj.BestKnownHeaderNumber),
		LastCommonHeaderHash:   jsonObj.LastCommonHeaderHash,
		LastCommonHeaderNumber: (*uint64)(jsonObj.LastCommonHeaderNumber),
		UnknownHeaderListSize:  uint64(jsonObj.UnknownHeaderListSize),
		InflightCount:          uint64(jsonObj.InflightCount),
		CanFetchCount:          uint64(jsonObj.CanFetchCount),
	}
	return nil
}

func (r *NodeAddress) UnmarshalJSON(input []byte) error {
	var jsonObj struct {
		Address string         `json:"address"`
		Score   hexutil.Uint64 `json:"score"`
	}
	if err := json.Unmarshal(input, &jsonObj); err != nil {
		return err
	}
	*r = NodeAddress{
		Address: jsonObj.Address,
		Score:   uint64(jsonObj.Score),
	}
	return nil
}

func (r *RemoteNode) UnmarshalJSON(input []byte) error {
	type RemoteAlias RemoteNode
	var jsonObj struct {
		RemoteAlias
		ConnectedDuration hexutil.Uint64  `json:"connected_duration"`
		LastPingDuration  *hexutil.Uint64 `json:"last_ping_duration,omitempty"`
	}
	if err := json.Unmarshal(input, &jsonObj); err != nil {
		return err
	}
	*r = RemoteNode{
		Version:           jsonObj.Version,
		NodeID:            jsonObj.NodeID,
		Addresses:         jsonObj.Addresses,
		IsOutbound:        jsonObj.IsOutbound,
		ConnectedDuration: uint64(jsonObj.ConnectedDuration),
		LastPingDuration:  (*uint64)(jsonObj.LastPingDuration),
		SyncState:         jsonObj.SyncState,
		Protocols:         jsonObj.Protocols,
	}
	return nil
}

func (r *LocalNodeProtocol) UnmarshalJSON(input []byte) error {
	type LocalNodeProtocolAlias LocalNodeProtocol
	var jsonObj struct {
		LocalNodeProtocolAlias
		Id hexutil.Uint64 `json:"id"`
	}
	if err := json.Unmarshal(input, &jsonObj); err != nil {
		return err
	}
	*r = LocalNodeProtocol{
		Id:              uint64(jsonObj.Id),
		Name:            jsonObj.Name,
		SupportVersions: jsonObj.SupportVersions,
	}
	return nil
}

func (r *LocalNode) UnmarshalJSON(input []byte) error {
	type LocalNodeAlias LocalNode
	var jsonObj struct {
		LocalNodeAlias
		Connections hexutil.Uint64 `json:"connections"`
	}
	if err := json.Unmarshal(input, &jsonObj); err != nil {
		return err
	}
	*r = LocalNode{
		Version:     jsonObj.Version,
		NodeId:      jsonObj.NodeId,
		Active:      jsonObj.Active,
		Addresses:   jsonObj.Addresses,
		Protocols:   jsonObj.Protocols,
		Connections: uint64(jsonObj.Connections),
	}
	return nil
}

func (r *BlockEconomicState) UnmarshalJSON(input []byte) error {
	var jsonObj struct {
		Issuance struct {
			Primary   hexutil.Uint64 `json:"primary"`
			Secondary hexutil.Uint64 `json:"secondary"`
		} `json:"issuance"`
		MinerReward struct {
			Primary   hexutil.Uint64 `json:"primary"`
			Secondary hexutil.Uint64 `json:"secondary"`
			Committed hexutil.Uint64 `json:"committed"`
			Proposal  hexutil.Uint64 `json:"proposal"`
		} `json:"miner_reward"`
		TxsFee      hexutil.Uint64 `json:"txs_fee"`
		FinalizedAt Hash           `json:"finalized_at"`
	}
	if err := json.Unmarshal(input, &jsonObj); err != nil {
		return err
	}
	*r = BlockEconomicState{
		Issuance: BlockIssuance{
			Primary:   uint64(jsonObj.Issuance.Primary),
			Secondary: uint64(jsonObj.Issuance.Secondary),
		},
		MinerReward: MinerReward{
			Primary:   uint64(jsonObj.MinerReward.Primary),
			Secondary: uint64(jsonObj.MinerReward.Secondary),
			Committed: uint64(jsonObj.MinerReward.Committed),
			Proposal:  uint64(jsonObj.MinerReward.Proposal),
		},
		TxsFee:      uint64(jsonObj.TxsFee),
		FinalizedAt: jsonObj.FinalizedAt,
	}
	return nil
}

func (r *BlockchainInfo) UnmarshalJSON(input []byte) error {
	var jsonObj struct {
		Alerts []*struct {
			Id          hexutil.Uint   `json:"id"`
			Message     string         `json:"message"`
			NoticeUntil hexutil.Uint64 `json:"notice_until"`
			Priority    hexutil.Uint   `json:"priority"`
		} `json:"alerts"`
		Chain                  string         `json:"chain"`
		Difficulty             hexutil.Big    `json:"difficulty"`
		Epoch                  hexutil.Uint64 `json:"epoch"`
		IsInitialBlockDownload bool           `json:"is_initial_block_download"`
		MedianTime             hexutil.Uint64 `json:"median_time"`
	}
	if err := json.Unmarshal(input, &jsonObj); err != nil {
		return err
	}

	alerts := make([]*AlertMessage, len(jsonObj.Alerts))
	for i := 0; i < len(jsonObj.Alerts); i++ {
		alerts[i] = &AlertMessage{
			Id:          uint32(jsonObj.Alerts[i].Id),
			Message:     jsonObj.Alerts[i].Message,
			NoticeUntil: uint64(jsonObj.Alerts[i].NoticeUntil),
			Priority:    uint32(jsonObj.Alerts[i].Priority),
		}
	}

	*r = BlockchainInfo{
		Alerts:                 alerts,
		Chain:                  jsonObj.Chain,
		Difficulty:             (*big.Int)(&jsonObj.Difficulty),
		Epoch:                  uint64(jsonObj.Epoch),
		IsInitialBlockDownload: jsonObj.IsInitialBlockDownload,
		MedianTime:             uint64(jsonObj.MedianTime),
	}
	return nil
}

func (r *TxPoolInfo) UnmarshalJSON(input []byte) error {
	var jsonObj struct {
		TipHash          Hash           `json:"tip_hash"`
		TipNumber        hexutil.Uint64 `json:"tip_number"`
		Pending          hexutil.Uint64 `json:"pending"`
		Proposed         hexutil.Uint64 `json:"proposed"`
		Orphan           hexutil.Uint64 `json:"orphan"`
		TotalTxSize      hexutil.Uint64 `json:"total_tx_size"`
		TotalTxCycles    hexutil.Uint64 `json:"total_tx_cycles"`
		MinFeeRate       hexutil.Uint64 `json:"min_fee_rate"`
		LastTxsUpdatedAt hexutil.Uint64 `json:"last_txs_updated_at"`
	}
	if err := json.Unmarshal(input, &jsonObj); err != nil {
		return err
	}
	*r = TxPoolInfo{
		TipHash:          jsonObj.TipHash,
		TipNumber:        uint64(jsonObj.TipNumber),
		Pending:          uint64(jsonObj.Pending),
		Proposed:         uint64(jsonObj.Proposed),
		Orphan:           uint64(jsonObj.Orphan),
		TotalTxSize:      uint64(jsonObj.TotalTxSize),
		TotalTxCycles:    uint64(jsonObj.TotalTxCycles),
		MinFeeRate:       uint64(jsonObj.MinFeeRate),
		LastTxsUpdatedAt: uint64(jsonObj.LastTxsUpdatedAt),
	}
	return nil
}

func (r *BannedAddress) UnmarshalJSON(input []byte) error {
	var jsonObj struct {
		Address   string         `json:"address"`
		BanReason string         `json:"ban_reason"`
		BanUntil  hexutil.Uint64 `json:"ban_until"`
		CreatedAt hexutil.Uint64 `json:"created_at"`
	}
	if err := json.Unmarshal(input, &jsonObj); err != nil {
		return err
	}
	*r = BannedAddress{
		Address:   jsonObj.Address,
		BanReason: jsonObj.BanReason,
		BanUntil:  uint64(jsonObj.BanUntil),
		CreatedAt: uint64(jsonObj.CreatedAt),
	}
	return nil
}

func (r *DryRunTransactionResult) UnmarshalJSON(input []byte) error {
	var result struct {
		Cycles hexutil.Uint64 `json:"cycles"`
	}
	if err := json.Unmarshal(input, &result); err != nil {
		return err
	}
	*r = DryRunTransactionResult{
		Cycles: uint64(result.Cycles),
	}
	return nil
}

func (r *EstimateCycles) UnmarshalJSON(input []byte) error {
	var result struct {
		Cycles hexutil.Uint64 `json:"cycles"`
	}
	if err := json.Unmarshal(input, &result); err != nil {
		return err
	}
	*r = EstimateCycles{
		Cycles: uint64(result.Cycles),
	}
	return nil
}

type JsonFeeRateStatics struct {
	Mean   hexutil.Uint64 `json:"mean"`
	Median hexutil.Uint64 `json:"median"`
}

func (r *FeeRateStatics) UnmarshalJSON(input []byte) error {
	var result JsonFeeRateStatics
	if err := json.Unmarshal(input, &result); err != nil {
		return err
	}

	*r = FeeRateStatics{
		Mean:   uint64(result.Mean),
		Median: uint64(result.Median),
	}
	return nil
}

func (r *FeeRateStatics) MarshalJSON() ([]byte, error) {
	jsonObj := &JsonFeeRateStatics{
		Mean:   hexutil.Uint64(r.Mean),
		Median: hexutil.Uint64(r.Median),
	}
	return json.Marshal(jsonObj)
}

type jsonTxStatus struct {
	Status      TransactionStatus `json:"status"`
	BlockHash   *Hash             `json:"block_hash"`
	BlockNumber *hexutil.Uint64   `json:"block_number"`
	TxIndex     *hexutil.Uint     `json:"tx_index"`
	Reason      *string           `json:"reason"`
}

func (r *TxStatus) MarshalJSON() ([]byte, error) {
	jsonObj := &jsonTxStatus{
		Status:    r.Status,
		BlockHash: r.BlockHash,
		Reason:    r.Reason,
	}

	if r.BlockNumber != nil {
		jsonObj.BlockNumber = (*hexutil.Uint64)(r.BlockNumber)
	}
	if r.TxIndex != nil {
		jsonObj.TxIndex = (*hexutil.Uint)(r.TxIndex)
	}
	return json.Marshal(jsonObj)
}

func (r *TxStatus) UnmarshalJSON(input []byte) error {
	var result jsonTxStatus
	if err := json.Unmarshal(input, &result); err != nil {
		return err
	}

	*r = TxStatus{
		Status:    result.Status,
		BlockHash: result.BlockHash,
		Reason:    result.Reason,
	}

	if result.BlockNumber != nil {
		r.BlockNumber = (*uint64)(result.BlockNumber)
	}

	if result.TxIndex != nil {
		r.TxIndex = (*uint)(result.TxIndex)
	}

	return nil
}

type jsonTransactionWithStatus struct {
	Transaction     *Transaction    `json:"transaction"`
	Cycles          *hexutil.Uint64 `json:"cycles"`
	TimeAddedToPool *hexutil.Uint64 `json:"time_added_to_pool"`
	TxStatus        *TxStatus       `json:"tx_status"`
}

func (r *TransactionWithStatus) MarshalJSON() ([]byte, error) {
	jsonObj := &jsonTransactionWithStatus{
		Transaction: r.Transaction,
		TxStatus:    r.TxStatus,
	}

	if r.Cycles != nil {
		jsonObj.Cycles = (*hexutil.Uint64)(r.Cycles)
	}

	if r.TimeAddedToPool != nil {
		jsonObj.TimeAddedToPool = (*hexutil.Uint64)(r.TimeAddedToPool)
	}

	return json.Marshal(jsonObj)
}

func (r *TransactionWithStatus) UnmarshalJSON(input []byte) error {
	var result jsonTransactionWithStatus
	if err := json.Unmarshal(input, &result); err != nil {
		return err
	}

	*r = TransactionWithStatus{
		Transaction: result.Transaction,
		TxStatus:    result.TxStatus,
	}

	if result.Cycles != nil {
		r.Cycles = (*uint64)(result.Cycles)
	}

	if result.TimeAddedToPool != nil {
		r.TimeAddedToPool = (*uint64)(result.TimeAddedToPool)
	}

	return nil
}

func (r *PackedBlock) UnmarshalJSON(input []byte) error {
	if err := json.Unmarshal(input, &r.Block); err != nil {
		return err
	}
	return nil
}

type jsonCellbaseTemplate struct {
	Hash   Hash            `json:"hash"`
	Cycles *hexutil.Uint64 `json:"cycles"`
	Data   Transaction     `json:"data"`
}

var _ json.Marshaler = new(CellbaseTemplate)
var _ json.Unmarshaler = new(CellbaseTemplate)

func (r *CellbaseTemplate) MarshalJSON() ([]byte, error) {
	jsonObj := &jsonCellbaseTemplate{
		Hash: r.Hash,
		Data: r.Data,
	}
	if r.Cycles != nil {
		jsonObj.Cycles = (*hexutil.Uint64)(r.Cycles)
	}
	return json.Marshal(jsonObj)
}

func (r *CellbaseTemplate) UnmarshalJSON(input []byte) error {
	var jsonObj jsonCellbaseTemplate
	if err := json.Unmarshal(input, &jsonObj); err != nil {
		return err
	}
	*r = CellbaseTemplate{
		Hash: jsonObj.Hash,
		Data: jsonObj.Data,
	}
	if jsonObj.Cycles != nil {
		r.Cycles = (*uint64)(jsonObj.Cycles)
	}
	return nil
}

type jsonBlockTemplate struct {
	Version          hexutil.Uint          `json:"version"`
	CompactTarget    hexutil.Uint          `json:"compact_target"`
	CurrentTime      hexutil.Uint64        `json:"current_time"`
	Number           hexutil.Uint64        `json:"number"`
	Epoch            hexutil.Uint64        `json:"epoch"`
	ParentHash       Hash                  `json:"parent_hash"`
	CyclesLimit      hexutil.Uint64        `json:"cycles_limit"`
	BytesLimit       uint64                `json:"bytes_limit"`
	UnclesCountLimit hexutil.Uint64        `json:"uncles_count_limit"`
	Uncles           []UncleTemplate       `json:"uncles"`
	Transactions     []TransactionTemplate `json:"transactions"`
	Proposals        []string              `json:"proposals"`
	Cellbase         CellbaseTemplate      `json:"cellbase"`
	WorkId           hexutil.Uint64        `json:"work_id"`
	Dao              Hash                  `json:"dao"`
	Extension        *json.RawMessage      `json:"extension"`
}

var _ json.Marshaler = new(BlockTemplate)
var _ json.Unmarshaler = new(BlockTemplate)

func (r *BlockTemplate) MarshalJSON() ([]byte, error) {
	jsonObj := &jsonBlockTemplate{
		Version:          hexutil.Uint(r.Version),
		CompactTarget:    hexutil.Uint(r.CompactTarget),
		CurrentTime:      hexutil.Uint64(r.CurrentTime),
		Number:           hexutil.Uint64(r.Number),
		Epoch:            hexutil.Uint64(r.Epoch),
		ParentHash:       r.ParentHash,
		CyclesLimit:      hexutil.Uint64(r.CyclesLimit),
		BytesLimit:       r.BytesLimit,
		UnclesCountLimit: hexutil.Uint64(r.UnclesCountLimit),
		Uncles:           r.Uncles,
		Transactions:     r.Transactions,
		Proposals:        r.Proposals,
		Cellbase:         r.Cellbase,
		WorkId:           hexutil.Uint64(r.WorkId),
		Dao:              r.Dao,
	}
	if r.Extension != nil {
		jsonObj.Extension = r.Extension
	}
	return json.Marshal(jsonObj)
}

func (r *BlockTemplate) UnmarshalJSON(input []byte) error {
	var jsonObj jsonBlockTemplate
	if err := json.Unmarshal(input, &jsonObj); err != nil {
		return err
	}
	*r = BlockTemplate{
		Version:          uint32(jsonObj.Version),
		CompactTarget:    uint32(jsonObj.CompactTarget),
		CurrentTime:      uint64(jsonObj.CurrentTime),
		Number:           uint64(jsonObj.Number),
		Epoch:            uint64(jsonObj.Epoch),
		ParentHash:       jsonObj.ParentHash,
		CyclesLimit:      uint64(jsonObj.CyclesLimit),
		BytesLimit:       jsonObj.BytesLimit,
		UnclesCountLimit: uint64(jsonObj.UnclesCountLimit),
		Uncles:           jsonObj.Uncles,
		Transactions:     jsonObj.Transactions,
		Proposals:        jsonObj.Proposals,
		Cellbase:         jsonObj.Cellbase,
		WorkId:           uint64(jsonObj.WorkId),
		Dao:              jsonObj.Dao,
	}
	if jsonObj.Extension != nil {
		r.Extension = jsonObj.Extension
	}
	return nil
}

type jsonAncestorsScoreSortKey struct {
	AncestorsFee    hexutil.Uint64 `json:"ancestors_fee"`
	AncestorsWeight hexutil.Uint64 `json:"ancestors_weight"`
	Fee             hexutil.Uint64 `json:"fee"`
	Weight          hexutil.Uint64 `json:"weight"`
}

var _ json.Marshaler = new(AncestorsScoreSortKey)
var _ json.Unmarshaler = new(AncestorsScoreSortKey)

func (r AncestorsScoreSortKey) MarshalJSON() ([]byte, error) {
	jsonObj := &jsonAncestorsScoreSortKey{
		AncestorsFee:    hexutil.Uint64(r.AncestorsFee),
		AncestorsWeight: hexutil.Uint64(r.AncestorsWeight),
		Fee:             hexutil.Uint64(r.Fee),
		Weight:          hexutil.Uint64(r.Weight),
	}
	return json.Marshal(jsonObj)
}

func (r *AncestorsScoreSortKey) UnmarshalJSON(input []byte) error {
	var jsonObj jsonAncestorsScoreSortKey
	if err := json.Unmarshal(input, &jsonObj); err != nil {
		return err
	}
	*r = AncestorsScoreSortKey{
		AncestorsFee:    uint64(jsonObj.AncestorsFee),
		AncestorsWeight: uint64(jsonObj.AncestorsWeight),
		Fee:             uint64(jsonObj.Fee),
		Weight:          uint64(jsonObj.Weight),
	}
	return nil
}

type jsonPoolTxDetailInfo struct {
	AncestorsCount   hexutil.Uint64        `json:"ancestors_count"`
	DescendantsCount hexutil.Uint64        `json:"descendants_count"`
	EntryStatus      string                `json:"entry_status"`
	PendingCount     hexutil.Uint64        `json:"pending_count"`
	ProposedCount    hexutil.Uint64        `json:"proposed_count"`
	RankInPending    hexutil.Uint64        `json:"rank_in_pending"`
	ScoreSortKey     AncestorsScoreSortKey `json:"score_sortkey"`
	Timestamp        hexutil.Uint64        `json:"timestamp"`
}

var _ json.Marshaler = new(PoolTxDetailInfo)
var _ json.Unmarshaler = new(PoolTxDetailInfo)

func (r PoolTxDetailInfo) MarshalJSON() ([]byte, error) {
	jsonObj := &jsonPoolTxDetailInfo{
		AncestorsCount:   hexutil.Uint64(r.AncestorsCount),
		DescendantsCount: hexutil.Uint64(r.DescendantsCount),
		EntryStatus:      r.EntryStatus,
		PendingCount:     hexutil.Uint64(r.PendingCount),
		ProposedCount:    hexutil.Uint64(r.ProposedCount),
		RankInPending:    hexutil.Uint64(r.RankInPending),
		ScoreSortKey:     r.ScoreSortKey,
		Timestamp:        hexutil.Uint64(r.Timestamp),
	}
	return json.Marshal(jsonObj)
}

func (r *PoolTxDetailInfo) UnmarshalJSON(input []byte) error {
	var jsonObj jsonPoolTxDetailInfo
	if err := json.Unmarshal(input, &jsonObj); err != nil {
		return err
	}
	*r = PoolTxDetailInfo{
		AncestorsCount:   uint64(jsonObj.AncestorsCount),
		DescendantsCount: uint64(jsonObj.DescendantsCount),
		EntryStatus:      jsonObj.EntryStatus,
		PendingCount:     uint64(jsonObj.PendingCount),
		ProposedCount:    uint64(jsonObj.ProposedCount),
		RankInPending:    uint64(jsonObj.RankInPending),
		ScoreSortKey:     jsonObj.ScoreSortKey,
		Timestamp:        uint64(jsonObj.Timestamp),
	}
	return nil
}

type jsonAlert struct {
	Id          hexutil.Uint      `json:"id"`
	Cancel      hexutil.Uint      `json:"cancel"`
	MinVersion  *string           `json:"min_version"`
	MaxVersion  *string           `json:"max_version"`
	Priority    hexutil.Uint      `json:"priority"`
	NoticeUntil hexutil.Uint64    `json:"notice_until"`
	Message     string            `json:"message"`
	Signatures  []json.RawMessage `json:"signatures"`
}

var _ json.Marshaler = new(Alert)
var _ json.Unmarshaler = new(Alert)

func (r Alert) MarshalJSON() ([]byte, error) {
	jsonObj := &jsonAlert{
		Id:          hexutil.Uint(r.Id),
		Cancel:      hexutil.Uint(r.Cancel),
		Priority:    hexutil.Uint(r.Priority),
		NoticeUntil: hexutil.Uint64(r.NoticeUntil),
		Message:     r.Message,
		Signatures:  r.Signatures,
	}
	if r.MinVersion != nil {
		jsonObj.MinVersion = r.MinVersion
	}
	if r.MaxVersion != nil {
		jsonObj.MaxVersion = r.MaxVersion
	}
	return json.Marshal(jsonObj)
}

func (r *Alert) UnmarshalJSON(input []byte) error {
	var jsonObj jsonAlert
	if err := json.Unmarshal(input, &jsonObj); err != nil {
		return err
	}
	*r = Alert{
		Id:          uint32(jsonObj.Id),
		Cancel:      uint32(jsonObj.Cancel),
		Priority:    uint32(jsonObj.Priority),
		NoticeUntil: uint64(jsonObj.NoticeUntil),
		Message:     jsonObj.Message,
		Signatures:  jsonObj.Signatures,
	}
	if jsonObj.MinVersion != nil {
		r.MinVersion = jsonObj.MinVersion
	}
	if jsonObj.MaxVersion != nil {
		r.MaxVersion = jsonObj.MaxVersion
	}
	return nil
}

type jsonAlertMessage struct {
	Id          hexutil.Uint   `json:"id"`
	Message     string         `json:"message"`
	NoticeUntil hexutil.Uint64 `json:"notice_until"`
	Priority    hexutil.Uint   `json:"priority"`
}

func (r AlertMessage) MarshalJSON() ([]byte, error) {
	jsonObj := &jsonAlertMessage{
		Id:          hexutil.Uint(r.Id),
		Message:     r.Message,
		NoticeUntil: hexutil.Uint64(r.NoticeUntil),
		Priority:    hexutil.Uint(r.Priority),
	}
	return json.Marshal(jsonObj)
}

func (r *AlertMessage) UnmarshalJSON(input []byte) error {
	var jsonObj jsonAlertMessage
	if err := json.Unmarshal(input, &jsonObj); err != nil {
		return err
	}
	*r = AlertMessage{
		Id:          uint32(jsonObj.Id),
		Message:     jsonObj.Message,
		NoticeUntil: uint64(jsonObj.NoticeUntil),
		Priority:    uint32(jsonObj.Priority),
	}
	return nil
}

type jsonDeploymentInfo struct {
	Bit                hexutil.Uint     `json:"bit"`
	Start              hexutil.Uint64   `json:"start"`
	Timeout            hexutil.Uint64   `json:"timeout"`
	MinActivationEpoch hexutil.Uint64   `json:"min_activation_epoch"`
	Period             hexutil.Uint64   `json:"period"`
	Threshold          jsonRationalU256 `json:"threshold"`
	Since              hexutil.Uint64   `json:"since"`
	State              DeploymentState  `json:"state"`
}

var _ json.Marshaler = new(DeploymentInfo)
var _ json.Unmarshaler = new(DeploymentInfo)

func (r DeploymentInfo) MarshalJSON() ([]byte, error) {
	jsonObj := &jsonDeploymentInfo{
		Bit:                hexutil.Uint(r.Bit),
		Start:              hexutil.Uint64(r.Start),
		Timeout:            hexutil.Uint64(r.Timeout),
		MinActivationEpoch: hexutil.Uint64(r.MinActivationEpoch),
		Period:             hexutil.Uint64(r.Period),
		Threshold:          r.Threshold,
		Since:              hexutil.Uint64(r.Since),
		State:              r.State,
	}
	return json.Marshal(jsonObj)
}
func (r *DeploymentInfo) UnmarshalJSON(input []byte) error {
	var jsonObj jsonDeploymentInfo
	if err := json.Unmarshal(input, &jsonObj); err != nil {
		return err
	}
	*r = DeploymentInfo{
		Bit:                uint8(jsonObj.Bit),
		Start:              uint64(jsonObj.Start),
		Timeout:            uint64(jsonObj.Timeout),
		MinActivationEpoch: uint64(jsonObj.MinActivationEpoch),
		Period:             uint64(jsonObj.Period),
		Threshold:          jsonObj.Threshold,
		Since:              uint64(jsonObj.Since),
		State:              jsonObj.State,
	}
	return nil
}

type jsonDeploymentsInfo struct {
	Hash        Hash                      `json:"hash"`
	Epoch       hexutil.Uint64            `json:"epoch"`
	Deployments map[string]DeploymentInfo `json:"deployments"`
}

var _ json.Marshaler = new(DeploymentsInfo)
var _ json.Unmarshaler = new(DeploymentsInfo)

func (r DeploymentsInfo) MarshalJSON() ([]byte, error) {
	deployments := make(map[string]DeploymentInfo)
	for k, v := range r.Deployments {
		deployments[k.String()] = v
	}
	jsonObj := &jsonDeploymentsInfo{
		Hash:        r.Hash,
		Epoch:       hexutil.Uint64(r.Epoch),
		Deployments: deployments,
	}
	return json.Marshal(jsonObj)
}
func (r *DeploymentsInfo) UnmarshalJSON(input []byte) error {
	var jsonObj jsonDeploymentsInfo
	if err := json.Unmarshal(input, &jsonObj); err != nil {
		return err
	}
	deployments := make(map[DeploymentPos]DeploymentInfo)
	for k, v := range jsonObj.Deployments {
		pos, err := DeploymentPosFromString(k)
		if err != nil {
			return err
		}
		deployments[pos] = v
	}
	*r = DeploymentsInfo{
		Hash:        jsonObj.Hash,
		Epoch:       uint64(jsonObj.Epoch),
		Deployments: deployments,
	}
	return nil
}

type jsonTxPoolEntry struct {
	Cycles          hexutil.Uint64 `json:"cycles"`
	Size            hexutil.Uint64 `json:"size"`
	Fee             hexutil.Uint64 `json:"fee"`
	AncestorsSize   hexutil.Uint64 `json:"ancestors_size"`
	AncestorsCycles hexutil.Uint64 `json:"ancestors_cycles"`
	AncestorsCount  hexutil.Uint64 `json:"ancestors_count"`
	Timestamp       hexutil.Uint64 `json:"timestamp"`
}

var _ json.Marshaler = new(TxPoolEntry)
var _ json.Unmarshaler = new(TxPoolEntry)

func (r TxPoolEntry) MarshalJSON() ([]byte, error) {
	jsonObj := &jsonTxPoolEntry{
		Cycles:          hexutil.Uint64(r.Cycles),
		Size:            hexutil.Uint64(r.Size),
		Fee:             hexutil.Uint64(r.Fee),
		AncestorsSize:   hexutil.Uint64(r.AncestorsSize),
		AncestorsCycles: hexutil.Uint64(r.AncestorsCycles),
		AncestorsCount:  hexutil.Uint64(r.AncestorsCount),
		Timestamp:       hexutil.Uint64(r.Timestamp),
	}
	return json.Marshal(jsonObj)
}
func (r *TxPoolEntry) UnmarshalJSON(input []byte) error {
	var jsonObj jsonTxPoolEntry
	if err := json.Unmarshal(input, &jsonObj); err != nil {
		return err
	}
	*r = TxPoolEntry{
		Cycles:          uint64(jsonObj.Cycles),
		Size:            uint64(jsonObj.Size),
		Fee:             uint64(jsonObj.Fee),
		AncestorsSize:   uint64(jsonObj.AncestorsSize),
		AncestorsCycles: uint64(jsonObj.AncestorsCycles),
		AncestorsCount:  uint64(jsonObj.AncestorsCount),
		Timestamp:       uint64(jsonObj.Timestamp),
	}
	return nil
}

type jsonEntryCompleted struct {
	Cycles hexutil.Uint64 `json:"cycles"`
	Fee    hexutil.Uint64 `json:"fee"`
}

var _ json.Marshaler = new(EntryCompleted)
var _ json.Unmarshaler = new(EntryCompleted)

func (r *EntryCompleted) UnmarshalJSON(input []byte) error {
	var result jsonEntryCompleted
	if err := json.Unmarshal(input, &result); err != nil {
		return err
	}
	*r = EntryCompleted{
		Cycles: uint64(result.Cycles),
		Fee:    uint64(result.Fee),
	}
	return nil
}

func (r *EntryCompleted) MarshalJSON() ([]byte, error) {
	jsonObj := &jsonEntryCompleted{
		Cycles: hexutil.Uint64(r.Cycles),
		Fee:    hexutil.Uint64(r.Fee),
	}
	return json.Marshal(jsonObj)
}
