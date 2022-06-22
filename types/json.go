package types

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
)

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
	err := json.Unmarshal(input, &jsonObj)
	if err != nil {
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
	err := json.Unmarshal(input, &jsonObj)
	if err != nil {
		return err
	}
	*r = Header{
		CompactTarget:    uint(jsonObj.CompactTarget),
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
		Version:          uint(jsonObj.Version),
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
	err := json.Unmarshal(input, &jsonObj)
	if err != nil {
		return err
	}
	*r = OutPoint{
		TxHash: jsonObj.TxHash,
		Index:  uint(jsonObj.Index),
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
	err := json.Unmarshal(input, &jsonObj)
	if err != nil {
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
	err := json.Unmarshal(input, &jsonObj)
	if err != nil {
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
	err := json.Unmarshal(input, &jsonObj)
	if err != nil {
		return err
	}
	*r = CellOutput{
		Capacity: uint64(jsonObj.Capacity),
		Lock:     jsonObj.Lock,
		Type:     jsonObj.Type,
	}
	return nil
}

type transactionAlias Transaction
type jsonTransaction struct {
	transactionAlias
	Version     hexutil.Uint    `json:"version"`
	OutputsData []hexutil.Bytes `json:"outputs_data"`
	Witnesses   []hexutil.Bytes `json:"witnesses"`
}

func (t Transaction) MarshalJSON() ([]byte, error) {
	toBytes := func(bytes [][]byte) []hexutil.Bytes {
		result := make([]hexutil.Bytes, len(bytes))
		for i, data := range bytes {
			result[i] = data
		}
		return result
	}
	jsonObj := &jsonTransaction{
		transactionAlias: transactionAlias(t),
		Version:          hexutil.Uint(t.Version),
		OutputsData:      toBytes(t.OutputsData),
		Witnesses:        toBytes(t.Witnesses),
	}
	return json.Marshal(jsonObj)
}

func (t *Transaction) UnmarshalJSON(input []byte) error {
	var jsonObj jsonTransaction
	err := json.Unmarshal(input, &jsonObj)
	if err != nil {
		return err
	}
	toByteArray := func(byteArray []hexutil.Bytes) [][]byte {
		result := make([][]byte, len(byteArray))
		for i, data := range byteArray {
			result[i] = data
		}
		return result
	}
	*t = Transaction{
		Version:     uint(jsonObj.Version),
		Hash:        jsonObj.Hash,
		CellDeps:    jsonObj.CellDeps,
		HeaderDeps:  jsonObj.HeaderDeps,
		Inputs:      jsonObj.Inputs,
		Outputs:     jsonObj.Outputs,
		OutputsData: toByteArray(jsonObj.OutputsData),
		Witnesses:   toByteArray(jsonObj.Witnesses),
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
	err := json.Unmarshal(input, &jsonObj)
	if err != nil {
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
	Rfc         string         `json:"rfc"`
	EpochNumber hexutil.Uint64 `json:"epoch_number,omitempty"`
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
	err := json.Unmarshal(input, &jsonObj)
	if err != nil {
		return err
	}
	toHardForkFeatureArray := func(a []*jsonHardForkFeature) []*HardForkFeature {
		result := make([]*HardForkFeature, len(a))
		for i, data := range a {
			result[i] = &HardForkFeature{
				Rfc:         data.Rfc,
				EpochNumber: uint64(data.EpochNumber),
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
		BlockVersion:                      uint(jsonObj.BlockVersion),
		TxVersion:                         uint(jsonObj.TxVersion),
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
}

func (t *SyncState) UnmarshalJSON(input []byte) error {
	var jsonObj jsonSyncState
	err := json.Unmarshal(input, &jsonObj)
	if err != nil {
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
	err := json.Unmarshal(input, &jsonObj)
	if err != nil {
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

	err := json.Unmarshal(input, &jsonObj)
	if err != nil {
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
		BestKnownHeaderNumber  hexutil.Uint64 `json:"best_known_header_number,omitempty"`
		LastCommonHeaderNumber hexutil.Uint64 `json:"last_common_header_number,omitempty"`
		UnknownHeaderListSize  hexutil.Uint64 `json:"unknown_header_list_size"`
		InflightCount          hexutil.Uint64 `json:"inflight_count"`
		CanFetchCount          hexutil.Uint64 `json:"can_fetch_count"`
	}
	err := json.Unmarshal(input, &jsonObj)
	if err != nil {
		return err
	}
	*r = PeerSyncState{
		BestKnownHeaderHash:    jsonObj.BestKnownHeaderHash,
		BestKnownHeaderNumber:  uint64(jsonObj.BestKnownHeaderNumber),
		LastCommonHeaderHash:   jsonObj.LastCommonHeaderHash,
		LastCommonHeaderNumber: uint64(jsonObj.LastCommonHeaderNumber),
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
	err := json.Unmarshal(input, &jsonObj)
	if err != nil {
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
		ConnectedDuration hexutil.Uint64 `json:"connected_duration"`
		LastPingDuration  hexutil.Uint64 `json:"last_ping_duration,omitempty"`
	}
	err := json.Unmarshal(input, &jsonObj)
	if err != nil {
		return err
	}
	*r = RemoteNode{
		Version:           jsonObj.Version,
		NodeID:            jsonObj.NodeID,
		Addresses:         jsonObj.Addresses,
		IsOutbound:        jsonObj.IsOutbound,
		ConnectedDuration: uint64(jsonObj.ConnectedDuration),
		LastPingDuration:  uint64(jsonObj.LastPingDuration),
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
	err := json.Unmarshal(input, &jsonObj)
	if err != nil {
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
	err := json.Unmarshal(input, &jsonObj)
	if err != nil {
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
