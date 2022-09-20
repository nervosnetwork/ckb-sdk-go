package types

import (
	"encoding/json"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

func AssertJsonEqual(t *testing.T, t1, t2 []byte) {
	assert.JSONEq(t, string(t1), string(t2))
}

func TestJsonScript(t *testing.T) {
	jsonText1 := []byte(`
{
    "args": "0xa897829e60ee4e3fb0e4abe65549ec4a5ddafad7",
    "code_hash": "0x9bd7e06f3ecf4be0f2fcd2188b23f1b9fcc88e5d4b65a8637b17723bbda3cce8",
    "hash_type": "type"
}`)
	var v Script
	json.Unmarshal(jsonText1, &v)
	assert.Equal(t, ethcommon.FromHex("0xa897829e60ee4e3fb0e4abe65549ec4a5ddafad7"), v.Args)
	assert.Equal(t, HexToHash("0x9bd7e06f3ecf4be0f2fcd2188b23f1b9fcc88e5d4b65a8637b17723bbda3cce8"), v.CodeHash)
	assert.Equal(t, HashTypeType, v.HashType)

	jsonText2, _ := json.Marshal(v)
	AssertJsonEqual(t, jsonText1, jsonText2)
}

func TestJsonOutPoint(t *testing.T) {
	jsonText1 := []byte(`
{
    "index": "0x2",
    "tx_hash": "0xf8de3bb47d055cdf460d93a2a6e1b05f7432f9777c8c474abf4eec1d4aee5d37"
}`)
	var v OutPoint
	json.Unmarshal(jsonText1, &v)
	assert.Equal(t, uint32(0x2), v.Index)
	assert.Equal(t, HexToHash("0xf8de3bb47d055cdf460d93a2a6e1b05f7432f9777c8c474abf4eec1d4aee5d37"), v.TxHash)

	jsonText2, _ := json.Marshal(v)
	AssertJsonEqual(t, jsonText1, jsonText2)
}

func TestJsonCellInput(t *testing.T) {
	jsonText1 := []byte(`
{
    "previous_output": {
        "index": "0xffffffff",
        "tx_hash": "0x0000000000000000000000000000000000000000000000000000000000000000"
    },
    "since": "0x4fe230"
}`)
	var v CellInput
	json.Unmarshal(jsonText1, &v)
	assert.Equal(t, uint64(0x4fe230), v.Since)
	assert.NotNil(t, v.PreviousOutput)

	jsonText2, _ := json.Marshal(v)
	AssertJsonEqual(t, jsonText1, jsonText2)
}

func TestJsonCellOutput(t *testing.T) {
	jsonText1 := []byte(`
{
    "capacity": "0x9502f9000",
    "lock": {
        "args": "0xa897829e60ee4e3fb0e4abe65549ec4a5ddafad7",
        "code_hash": "0x9bd7e06f3ecf4be0f2fcd2188b23f1b9fcc88e5d4b65a8637b17723bbda3cce8",
        "hash_type": "type"
    },
    "type": {
        "args": "0x02",
        "code_hash": "0x554cff969f3148e3c620749384004e9692e67c429f621554d139b505a281c7b8",
        "hash_type": "type"
    }
}`)
	var v CellOutput
	json.Unmarshal(jsonText1, &v)
	assert.Equal(t, uint64(0x9502f9000), v.Capacity)
	assert.Equal(t, HexToHash("0x9bd7e06f3ecf4be0f2fcd2188b23f1b9fcc88e5d4b65a8637b17723bbda3cce8"), v.Lock.CodeHash)
	assert.Equal(t, HexToHash("0x554cff969f3148e3c620749384004e9692e67c429f621554d139b505a281c7b8"), v.Type.CodeHash)

	jsonText2, _ := json.Marshal(v)
	AssertJsonEqual(t, jsonText1, jsonText2)
}

func TestJsonCellDep(t *testing.T) {
	jsonText1 := []byte(`
{
    "dep_type": "code",
    "out_point": {
        "index": "0x2",
        "tx_hash": "0x8f8c79eb6671709633fe6a46de93c0fedc9c1b8a6527a18d3983879542635c9f"
    }
}`)
	var v CellDep
	json.Unmarshal(jsonText1, &v)
	assert.Equal(t, DepTypeCode, v.DepType)
	assert.NotNil(t, v.OutPoint)

	jsonText2, _ := json.Marshal(v)
	AssertJsonEqual(t, jsonText1, jsonText2)
}

func TestJsonTransactionWithStatus(t *testing.T) {
	jsonText1 := []byte(`
{
    "transaction": {
        "cell_deps": [
            {
                "dep_type": "dep_group",
                "out_point": {
                    "index": "0x1",
                    "tx_hash": "0xf8de3bb47d055cdf460d93a2a6e1b05f7432f9777c8c474abf4eec1d4aee5d37"
                }
            }
        ],
        "header_deps": [],
        "inputs": [
            {
                "previous_output": {
                    "index": "0x0",
                    "tx_hash": "0x0dff101e716d77507bddc5ca189dc24c80e0fb8c269775b988b3cdd64e4f3395"
                },
                "since": "0x0"
            }
        ],
        "outputs": [
            {
                "capacity": "0xbaa315500",
                "lock": {
                    "args": "0x4049ed9cec8a0d39c7a1e899f0dacb8a8c28ad14",
                    "code_hash": "0x9bd7e06f3ecf4be0f2fcd2188b23f1b9fcc88e5d4b65a8637b17723bbda3cce8",
                    "hash_type": "type"
                },
                "type": null
            },
            {
                "capacity": "0xdd2a73b8bf",
                "lock": {
                    "args": "0xbc9818d8a149cfc0cd0323386c46ba07920a037f",
                    "code_hash": "0x5c5069eb0857efc65e1bca0c07df34c31663b3622fd3876c876320fc9634e2a8",
                    "hash_type": "type"
                },
                "type": null
            }
        ],
        "outputs_data": ["0x", "0x"],
        "version": "0x0",
        "witnesses": [
            "0xc200000010000000c2000000c2000000ae000000000002027336b0ba900684cb3cb00f0d46d4f64c0994a5625724c1e3925a5206944d753a6f3edaedf977d77f75ef2bf584ab0f400063964d5cddb3443fb5f11cbf00eedd76c64205f6c2d2ce342582871a010af6560bc6f559222852ffc44d3c9db9ae76092d843a05e39c0000ae2adec03512e320c2f0c087ec1d366c5fb43f7862fd1a7693284d356fbf56196e8f8ccd5cabe21bf3f0b2763d0c4f02c79af0d9993572eb3b752b09b08b6b1f00"
        ]
    },
    "tx_status": {
        "block_hash": "0xe1ed2d2282aad742a95abe51c21d50b1c19e194f21fbd1ed2516f82bd042579a",
        "status": "committed"
    }
}`)
	var v TransactionWithStatus
	json.Unmarshal(jsonText1, &v)

	jsonText2, _ := json.Marshal(v)
	AssertJsonEqual(t, jsonText1, jsonText2)
}

func TestTransactionHashNotToMarshal(t *testing.T) {
	jsonText1 := []byte(`
{
    "version": "0x0",
    "cell_deps": null,
    "header_deps": [],
    "inputs": null,
    "outputs": null,
    "outputs_data": [],
    "witnesses": []
}`)
	v := &Transaction{
		Version:     0,
		Hash:        HexToHash("0xae02c44fb5b78b4b1bfc6097d89e0563da323e316ed0551091912d3ddf3f5a19"),
		CellDeps:    nil,
		HeaderDeps:  nil,
		Inputs:      nil,
		Outputs:     nil,
		OutputsData: nil,
		Witnesses:   nil,
	}

	jsonText2, _ := json.Marshal(v)
	AssertJsonEqual(t, jsonText1, jsonText2)
}

func TestTransactionHashToUnmarshal(t *testing.T) {
	jsonText := []byte(`
{
    "version": "0x0",
    "cell_deps": [],
    "header_deps": [],
    "inputs": [],
    "outputs": [],
    "outputs_data": [],
    "witnesses": [],
	"hash": "0xae02c44fb5b78b4b1bfc6097d89e0563da323e316ed0551091912d3ddf3f5a19"
}`)

	var v Transaction
	json.Unmarshal(jsonText, &v)
	assert.Equal(t, HexToHash("0xae02c44fb5b78b4b1bfc6097d89e0563da323e316ed0551091912d3ddf3f5a19"), v.Hash)
}

func TestJsonEpoch(t *testing.T) {
	jsonText1 := []byte(`
{
    "compact_target": "0x1d5f396f",
    "length": "0x356",
    "number": "0x100",
    "start_number": "0x2b445"
}`)
	var v Epoch
	json.Unmarshal(jsonText1, &v)
	assert.Equal(t, uint64(0x1d5f396f), v.CompactTarget)
	assert.Equal(t, uint64(0x356), v.Length)
	assert.Equal(t, uint64(0x100), v.Number)
	assert.Equal(t, uint64(0x2b445), v.StartNumber)

	jsonText2, _ := json.Marshal(v)
	AssertJsonEqual(t, jsonText1, jsonText2)
}

func TestJsonHeader(t *testing.T) {
	jsonText1 := []byte(`
{
    "compact_target": "0x1d43106d",
    "dao": "0x0e6beebedbb7962fb1389bfef5b32300a47716f7b5ae3200005910b7600e0507",
    "epoch": "0x28c0033000111",
    "extra_hash": "0x0000000000000000000000000000000000000000000000000000000000000000",
    "hash": "0x9f2b44451708cd7dcf671613cf30409b7b2f94dc32a35babb7cdca085a8062e7",
    "nonce": "0xae986fa353b387f912f1b181439f26fe",
    "number": "0x2e60b",
    "parent_hash": "0xf45e0ba01bce37a285b3b649ee59fc3dfbe115ead2c2367cb96ba0ea97f3e8a1",
    "proposals_hash": "0x0000000000000000000000000000000000000000000000000000000000000000",
    "timestamp": "0x1732486bcfe",
    "transactions_root": "0xb73f9303351a7bd0f81ae8cbda665ace579be0f801bdbed8b52904e768b45f46",
    "version": "0x0"
}`)
	var v Header
	json.Unmarshal(jsonText1, &v)

	jsonText2, _ := json.Marshal(v)
	AssertJsonEqual(t, jsonText1, jsonText2)
}

func TestJsonCellWithStatus(t *testing.T) {
	jsonText1 := []byte(`
{
    "cell": {
        "data": {
            "content": "0xf868560000000000",
            "hash": "0x8933d7a3cb3f30a589b766ff8ac1314989f4909354c6688f89275f690d306c67"
        },
        "output": {
            "capacity": "0xbdfd63e00",
            "lock": {
                "args": "0x4049ed9cec8a0d39c7a1e899f0dacb8a8c28ad14",
                "code_hash": "0x9bd7e06f3ecf4be0f2fcd2188b23f1b9fcc88e5d4b65a8637b17723bbda3cce8",
                "hash_type": "type"
            },
            "type": {
                "args": "0x",
                "code_hash": "0x82d76d1b75fe2fd9a27dfbaa65a039221a380d76c926f378d3f81cf3e7e13f2e",
                "hash_type": "type"
            }
        }
    },
    "status": "live"
}`)
	var v CellWithStatus
	json.Unmarshal(jsonText1, &v)

	jsonText2, _ := json.Marshal(v)
	AssertJsonEqual(t, jsonText1, jsonText2)
}

func TestJsonConsensus(t *testing.T) {
	jsonText := []byte(`
{
    "block_version": "0x0",
    "cellbase_maturity": "0x10000000004",
    "dao_type_hash": "0x82d76d1b75fe2fd9a27dfbaa65a039221a380d76c926f378d3f81cf3e7e13f2e",
    "epoch_duration_target": "0x3840",
    "genesis_hash": "0x10639e0895502b5688a6be8cf69460d76541bfa4821629d86d62ba0aae3f9606",
    "hardfork_features": [
        { "epoch_number": "0xc29", "rfc": "0028" },
        { "epoch_number": "0xc29", "rfc": "0029" },
        { "epoch_number": "0xc29", "rfc": "0030" },
        { "epoch_number": "0xc29", "rfc": "0031" },
        { "epoch_number": "0xc29", "rfc": "0032" },
        { "epoch_number": "0xc29", "rfc": "0036" },
        { "epoch_number": "0xc29", "rfc": "0038" }
    ],
    "id": "ckb_testnet",
    "initial_primary_epoch_reward": "0xae6c73c3e070",
    "max_block_bytes": "0x91c08",
    "max_block_cycles": "0xd09dc300",
    "max_block_proposals_limit": "0x5dc",
    "max_uncles_num": "0x2",
    "median_time_block_count": "0x25",
    "orphan_rate_target": { "denom": "0x28", "numer": "0x1" },
    "permanent_difficulty_in_dummy": false,
    "primary_epoch_reward_halving_interval": "0x2238",
    "proposer_reward_ratio": { "denom": "0xa", "numer": "0x4" },
    "secondary_epoch_reward": "0x37d0c8e28542",
    "secp256k1_blake160_multisig_all_type_hash": "0x5c5069eb0857efc65e1bca0c07df34c31663b3622fd3876c876320fc9634e2a8",
    "secp256k1_blake160_sighash_all_type_hash": "0x9bd7e06f3ecf4be0f2fcd2188b23f1b9fcc88e5d4b65a8637b17723bbda3cce8",
    "tx_proposal_window": { "closest": "0x2", "farthest": "0xa" },
    "tx_version": "0x0",
    "type_id_code_hash": "0x00000000000000000000000000000000000000000000000000545950455f4944"
}`)
	var v Consensus
	json.Unmarshal(jsonText, &v)
	assert.Equal(t, "ckb_testnet", v.Id)
	assert.Equal(t, HexToHash("0x10639e0895502b5688a6be8cf69460d76541bfa4821629d86d62ba0aae3f9606"), v.GenesisHash)
	assert.Equal(t, HexToHash("0x82d76d1b75fe2fd9a27dfbaa65a039221a380d76c926f378d3f81cf3e7e13f2e"), v.DaoTypeHash)
	assert.Equal(t, HexToHash("0x9bd7e06f3ecf4be0f2fcd2188b23f1b9fcc88e5d4b65a8637b17723bbda3cce8"), v.Secp256k1Blake160SighashAllTypeHash)
	assert.Equal(t, HexToHash("0x5c5069eb0857efc65e1bca0c07df34c31663b3622fd3876c876320fc9634e2a8"), v.Secp256k1Blake160MultisigAllTypeHash)
	assert.Equal(t, uint64(0xae6c73c3e070), v.InitialPrimaryEpochReward)
	assert.Equal(t, uint64(0x37d0c8e28542), v.SecondaryEpochReward)
	assert.Equal(t, uint64(0x2), v.MaxUnclesNum)
	assert.Equal(t, big.NewInt(0x28), v.OrphanRateTarget.Denom)
	assert.Equal(t, big.NewInt(0x1), v.OrphanRateTarget.Numer)
	assert.Equal(t, uint64(0x3840), v.EpochDurationTarget)
	assert.Equal(t, uint64(0x2), v.TxProposalWindow.Closest)
	assert.Equal(t, uint64(0xa), v.TxProposalWindow.Farthest)
	assert.Equal(t, big.NewInt(0xa), v.ProposerRewardRatio.Denom)
	assert.Equal(t, big.NewInt(0x4), v.ProposerRewardRatio.Numer)
	assert.Equal(t, uint64(0x10000000004), v.CellbaseMaturity)
	assert.Equal(t, uint64(0x25), v.MedianTimeBlockCount)
	assert.Equal(t, uint64(0xd09dc300), v.MaxBlockCycles)
	assert.Equal(t, uint64(0x91c08), v.MaxBlockBytes)
	assert.Equal(t, uint32(0x0), v.BlockVersion)
	assert.Equal(t, uint32(0x0), v.TxVersion)
	assert.Equal(t, HexToHash("0x00000000000000000000000000000000000000000000000000545950455f4944"), v.TypeIdCodeHash)
	assert.Equal(t, uint64(0x5dc), v.MaxBlockProposalsLimit)
	assert.Equal(t, uint64(0x2238), v.PrimaryEpochRewardHalvingInterval)
	assert.Equal(t, false, v.PermanentDifficultyInDummy)
	assert.Equal(t, 7, len(v.HardforkFeatures))
}

func TestJsonSyncState(t *testing.T) {
	jsonText := []byte(`
{
    "best_known_block_number": "0x5829fe",
    "best_known_block_timestamp": "0x1818919dfd6",
    "fast_time": "0x1d11",
    "ibd": false,
    "inflight_blocks_count": "0x0",
    "low_time": "0x9a4f",
    "normal_time": "0x9877",
    "orphan_blocks_count": "0x0"
}`)
	var v SyncState
	json.Unmarshal(jsonText, &v)
	assert.Equal(t, false, v.Ibd)
	assert.Equal(t, uint64(0x5829fe), v.BestKnownBlockNumber)
	assert.Equal(t, uint64(0x1818919dfd6), v.BestKnownBlockTimestamp)
	assert.Equal(t, uint64(0x1d11), v.FastTime)
	assert.Equal(t, uint64(0x0), v.InflightBlocksCount)
	assert.Equal(t, uint64(0x9a4f), v.LowTime)
	assert.Equal(t, uint64(0x9877), v.NormalTime)
	assert.Equal(t, uint64(0x0), v.OrphanBlocksCount)
}

func TestJsonBlock(t *testing.T) {
	jsonText1 := []byte(`
{
    "header": {
        "compact_target": "0x1e015555",
        "dao": "0x046d9f215d59a12eb612ba52fd8623008ffc0768330c000000906f0260fcfe06",
        "epoch": "0x3e80100000000",
        "extra_hash": "0x0000000000000000000000000000000000000000000000000000000000000000",
        "hash": "0x9584cfe1b317037028e487c46aebcfa6266c09d7f3ed7598d0011dbd6f612408",
        "nonce": "0x5289f79360d80b3233f667c39cd03c9d",
        "number": "0x100",
        "parent_hash": "0x946c36de76de83cd517d78e43b09c5df18c8ec1d66868306171e14b80d0f714b",
        "proposals_hash": "0x0000000000000000000000000000000000000000000000000000000000000000",
        "timestamp": "0x1723bae4a66",
        "transactions_root": "0xae66b0df7bdcc194d509aeb908a734d0c0795b87e360b5e074ec6754a64fcb5d",
        "version": "0x0"
    },
    "proposals": [],
    "transactions": [
        {
            "cell_deps": [],
            "header_deps": [],
            "inputs": [
                {
                    "previous_output": {
                        "index": "0xffffffff",
                        "tx_hash": "0x0000000000000000000000000000000000000000000000000000000000000000"
                    },
                    "since": "0x100"
                }
            ],
            "outputs": [
                {
                    "capacity": "0x2ecbd5c40a",
                    "lock": {
                        "args": "0xda648442dbb7347e467d1d09da13e5cd3a0ef0e1",
                        "code_hash": "0x9bd7e06f3ecf4be0f2fcd2188b23f1b9fcc88e5d4b65a8637b17723bbda3cce8",
                        "hash_type": "type"
                    },
                    "type": null
                }
            ],
            "outputs_data": ["0x"],
            "version": "0x0",
            "witnesses": [
                "0x5d0000000c00000055000000490000001000000030000000310000009bd7e06f3ecf4be0f2fcd2188b23f1b9fcc88e5d4b65a8637b17723bbda3cce80114000000da648442dbb7347e467d1d09da13e5cd3a0ef0e104000000deadbeef"
            ]
        }
    ],
    "uncles": []
}`)
	var v Block
	json.Unmarshal(jsonText1, &v)

	jsonText2, _ := json.Marshal(v)
	AssertJsonEqual(t, jsonText1, jsonText2)
}

func TestJsonTransactionProof(t *testing.T) {
	jsonText1 := []byte(`
{
    "block_hash": "0x151a28416008bd1f6ee7472e29db1641e626acfae97d8d53389e4184b359d82d",
    "proof": {
        "indices": ["0x4"],
        "lemmas": [
            "0xf268a2da588c5603295faf8dbae71efece10d829295587fc7be2852583839419",
            "0xd41ec2a757311c56ae635e0bcb9d5caf2d5c5293b6685ba2e28757cea1b7b6e9"
        ]
    },
    "witnesses_root": "0x06beb892d12c795d84e325c257750d79de29ee7e4fb583058ca5d8b1073e230d"
}`)
	var v TransactionProof
	json.Unmarshal(jsonText1, &v)

	jsonText2, _ := json.Marshal(v)
	AssertJsonEqual(t, jsonText1, jsonText2)
}

func TestJsonRemoteNode(t *testing.T) {
	jsonText := []byte(`
{
    "addresses": [
        {
            "address": "/ip4/47.74.66.72/tcp/8111/p2p/QmPhgweKm2ciYq52LjtEDmKFqHxGcg2WQ8RLCayRRycanD",
            "score": "0x64"
        },
        {
            "address": "/ip4/47.74.66.72/tcp/8111/p2p/QmPhgweKm2ciYq52LjtEDmKFqHxGcg2WQ8RLCayRRycanD",
            "score": "0x64"
        }
    ],
    "connected_duration": "0x909ae7b6",
    "is_outbound": true,
    "last_ping_duration": "0x1a0",
    "node_id": "QmPhgweKm2ciYq52LjtEDmKFqHxGcg2WQ8RLCayRRycanD",
    "protocols": [
        { "id": "0x1", "version": "2" },
        { "id": "0x2", "version": "2" },
        { "id": "0x67", "version": "2" },
        { "id": "0x66", "version": "2" },
        { "id": "0x6e", "version": "2" },
        { "id": "0x64", "version": "2" },
        { "id": "0x0", "version": "2" },
        { "id": "0x4", "version": "2" }
    ],
    "sync_state": {
        "best_known_header_hash": "0x1201e4a20d3cddc682173f892bea13127d6de3e00719a038d16a660968be067e",
        "best_known_header_number": "0x583019",
        "can_fetch_count": "0x10",
        "inflight_count": "0x0",
        "last_common_header_hash": "0x1201e4a20d3cddc682173f892bea13127d6de3e00719a038d16a660968be067e",
        "last_common_header_number": "0x583019",
        "unknown_header_list_size": "0x0"
    },
    "version": "0.103.0 (e77138e 2022-04-11)"
}`)
	var v RemoteNode
	json.Unmarshal(jsonText, &v)
	assert.Equal(t, 2, len(v.Addresses))
	assert.Equal(t, "/ip4/47.74.66.72/tcp/8111/p2p/QmPhgweKm2ciYq52LjtEDmKFqHxGcg2WQ8RLCayRRycanD", v.Addresses[0].Address)
	assert.Equal(t, uint64(0x64), v.Addresses[0].Score)
	assert.Equal(t, uint64(0x909ae7b6), v.ConnectedDuration)
	assert.Equal(t, uint64(0x1a0), v.LastPingDuration)
	assert.Equal(t, "QmPhgweKm2ciYq52LjtEDmKFqHxGcg2WQ8RLCayRRycanD", v.NodeID)
	assert.Equal(t, 8, len(v.Protocols))
	assert.Equal(t, uint64(0x67), v.Protocols[2].ID)
	assert.Equal(t, "2", v.Protocols[2].Version)
	assert.Equal(t, HexToHash("0x1201e4a20d3cddc682173f892bea13127d6de3e00719a038d16a660968be067e"), v.SyncState.BestKnownHeaderHash)
	assert.Equal(t, uint64(0x583019), v.SyncState.BestKnownHeaderNumber)
	assert.Equal(t, uint64(0x10), v.SyncState.CanFetchCount)
	assert.Equal(t, uint64(0x0), v.SyncState.InflightCount)
	assert.Equal(t, HexToHash("0x1201e4a20d3cddc682173f892bea13127d6de3e00719a038d16a660968be067e"), v.SyncState.LastCommonHeaderHash)
	assert.Equal(t, uint64(0x583019), v.SyncState.LastCommonHeaderNumber)
	assert.Equal(t, uint64(0x0), v.SyncState.UnknownHeaderListSize)
}

func TestJsonLocalNode(t *testing.T) {
	jsonText1 := []byte(`
{
    "active": true,
    "addresses": [{ "address": "/ip4/0.0.0.0/tcp/8115", "score": "0x1" }],
    "connections": "0x8",
    "node_id": "Qmeubaw22HAiGh236uRuT17KPK8Jjfi8zv9Sz6VBjxgfqn",
    "protocols": [
        {
            "id": "0x64",
            "name": "/ckb/syn",
            "support_versions": ["1", "2"]
        },
        { "id": "0x67", "name": "/ckb/relay", "support_versions": ["2"] },
        { "id": "0x65", "name": "/ckb/rel", "support_versions": ["1"] },
        {
            "id": "0x66",
            "name": "/ckb/tim",
            "support_versions": ["1", "2"]
        },
        {
            "id": "0x6e",
            "name": "/ckb/alt",
            "support_versions": ["1", "2"]
        },
        {
            "id": "0x2",
            "name": "/ckb/identify",
            "support_versions": ["0.0.1", "2"]
        },
        {
            "id": "0x0",
            "name": "/ckb/ping",
            "support_versions": ["0.0.1", "2"]
        },
        {
            "id": "0x1",
            "name": "/ckb/discovery",
            "support_versions": ["0.0.1", "2"]
        },
        {
            "id": "0x3",
            "name": "/ckb/flr",
            "support_versions": ["0.0.1", "2"]
        },
        {
            "id": "0x4",
            "name": "/ckb/disconnectmsg",
            "support_versions": ["0.0.1", "2"]
        }
    ],
    "version": "0.103.0 (e77138e 2022-04-11)"
}`)
	var v LocalNode
	json.Unmarshal(jsonText1, &v)
	assert.Equal(t, true, v.Active)
	assert.Equal(t, 1, len(v.Addresses))
	assert.Equal(t, uint64(0x8), v.Connections)
	assert.Equal(t, "Qmeubaw22HAiGh236uRuT17KPK8Jjfi8zv9Sz6VBjxgfqn", v.NodeId)
	assert.Equal(t, 10, len(v.Protocols))
	assert.Equal(t, uint64(0x64), v.Protocols[0].Id)
	assert.Equal(t, "/ckb/syn", v.Protocols[0].Name)
	assert.Equal(t, []string{"1", "2"}, v.Protocols[0].SupportVersions)
	assert.Equal(t, "0.103.0 (e77138e 2022-04-11)", v.Version)
}

func TestJsonBlockEconomicState(t *testing.T) {
	jsonText := []byte(`
{
    "finalized_at": "0xb2b45d98c93bfcf95a3edaabc4ddd784b208923448fc474339b0f016ab8e6b42",
    "issuance": { "primary": "0x18ce922bca", "secondary": "0x7f02ec655" },
    "miner_reward": {
        "committed": "0x0",
        "primary": "0x18ce922bca",
        "proposal": "0x190",
        "secondary": "0x109c18998"
    },
    "txs_fee": "0x100"
}`)
	var v BlockEconomicState
	json.Unmarshal(jsonText, &v)
	assert.Equal(t, HexToHash("0xb2b45d98c93bfcf95a3edaabc4ddd784b208923448fc474339b0f016ab8e6b42"), v.FinalizedAt)
	assert.Equal(t, uint64(0x18ce922bca), v.Issuance.Primary)
	assert.Equal(t, uint64(0x7f02ec655), v.Issuance.Secondary)
	assert.Equal(t, uint64(0x18ce922bca), v.MinerReward.Primary)
	assert.Equal(t, uint64(0x190), v.MinerReward.Proposal)
	assert.Equal(t, uint64(0x109c18998), v.MinerReward.Secondary)
	assert.Equal(t, uint64(0x100), v.TxsFee)
}

func TestJsonBlockchainInfo(t *testing.T) {
	jsonText := []byte(`
{
    "alerts": [
        {
            "id": "0x2a",
            "message": "An example alert message!",
            "notice_until": "0x24bcca57c00",
            "priority": "0x1"
        }
    ],
    "chain": "ckb",
    "difficulty": "0x1f4003",
    "epoch": "0x7080018000001",
    "is_initial_block_download": true,
    "median_time": "0x5cd2b105"
}`)
	var v BlockchainInfo
	json.Unmarshal(jsonText, &v)
	assert.Equal(t, 1, len(v.Alerts))
	assert.Equal(t, uint32(0x2a), v.Alerts[0].Id)
	assert.Equal(t, "An example alert message!", v.Alerts[0].Message)
	assert.Equal(t, "ckb", v.Chain)
	assert.Equal(t, big.NewInt(0x1f4003), v.Difficulty)
	assert.Equal(t, uint64(0x7080018000001), v.Epoch)
	assert.Equal(t, true, v.IsInitialBlockDownload)
	assert.Equal(t, uint64(0x5cd2b105), v.MedianTime)
}

func TestJsonTxPoolInfo(t *testing.T) {
	jsonText := []byte(`
{
    "last_txs_updated_at": "0x0",
    "min_fee_rate": "0x0",
    "orphan": "0x0",
    "pending": "0x1",
    "proposed": "0x0",
    "tip_hash": "0xa5f5c85987a15de25661e5a214f2c1449cd803f071acc7999820f25246471f40",
    "tip_number": "0x400",
    "total_tx_cycles": "0x219",
    "total_tx_size": "0x112"
}`)
	var v TxPoolInfo
	json.Unmarshal(jsonText, &v)
	assert.Equal(t, HexToHash("0xa5f5c85987a15de25661e5a214f2c1449cd803f071acc7999820f25246471f40"), v.TipHash)
	assert.Equal(t, uint64(0x0), v.LastTxsUpdatedAt)
	assert.Equal(t, uint64(0x0), v.MinFeeRate)
	assert.Equal(t, uint64(0x1), v.Pending)
	assert.Equal(t, uint64(0x400), v.TipNumber)
	assert.Equal(t, uint64(0x219), v.TotalTxCycles)
	assert.Equal(t, uint64(0x112), v.TotalTxSize)
}

func TestJsonBannedAddress(t *testing.T) {
	jsonText := []byte(`
{
    "address": "192.168.0.2/32",
    "ban_reason": "",
    "ban_until": "0x1ac89236180",
    "created_at": "0x16bde533338"
}`)
	var v BannedAddress
	json.Unmarshal(jsonText, &v)
	assert.Equal(t, "192.168.0.2/32", v.Address)
	assert.Equal(t, uint64(0x1ac89236180), v.BanUntil)
	assert.Equal(t, uint64(0x16bde533338), v.CreatedAt)
}
