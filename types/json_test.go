package types

import (
	"encoding/json"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)
func AssertJsonEqual(t *testing.T, t1, t2 []byte) {
	m1 := map[string]interface{}{}
	m2 := map[string]interface{}{}
	json.Unmarshal(t1, &m1)
	json.Unmarshal(t2, &m2)
	assert.Equal(t, m2, m1)
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
	assert.Equal(t, uint(0x2), v.Index)
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

func TestJsonTransaction(t *testing.T) {
	jsonText1 := []byte(`
{
    "cell_deps": [
        {
            "dep_type": "dep_group",
            "out_point": {
                "index": "0x1",
                "tx_hash": "0xf8de3bb47d055cdf460d93a2a6e1b05f7432f9777c8c474abf4eec1d4aee5d37"
            }
        }
    ],
    "hash": "0xb19806b3ccc091a19d929d0237e6dc6e9b128a468b5b33c121c1bc59ad87877a",
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
}`)
	var v Transaction
	json.Unmarshal(jsonText1, &v)

	jsonText2, _ := json.Marshal(v)
	AssertJsonEqual(t, jsonText1, jsonText2)
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
	assert.Equal(t, uint(0x0), v.BlockVersion)
	assert.Equal(t, uint(0x0), v.TxVersion)
	assert.Equal(t, HexToHash("0x00000000000000000000000000000000000000000000000000545950455f4944"), v.TypeIdCodeHash)
	assert.Equal(t, uint64(0x5dc), v.MaxBlockProposalsLimit)
	assert.Equal(t, uint64(0x2238), v.PrimaryEpochRewardHalvingInterval)
	assert.Equal(t, false, v.PermanentDifficultyInDummy)
	assert.Equal(t, 7, len(v.HardforkFeatures))
}