package transaction

import (
	"github.com/ethereum/go-ethereum/common"
	s "github.com/nervosnetwork/ckb-sdk-go/crypto/secp256k1"
	"github.com/nervosnetwork/ckb-sdk-go/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSignTransaction(t *testing.T) {
	// https://pudge.explorer.nervos.org/transaction/0x150ab94cc3d35daf96d0d55a4efc420323adcc36662b2bdcab826e16ce38dd81
	tx := &types.Transaction{
		Version: 0,
		CellDeps: []*types.CellDep{
			{
				OutPoint: &types.OutPoint{
					TxHash: types.HexToHash("0xf8de3bb47d055cdf460d93a2a6e1b05f7432f9777c8c474abf4eec1d4aee5d37"),
					Index:  0,
				},
				DepType: types.DepTypeDepGroup,
			},
			{
				OutPoint: &types.OutPoint{
					TxHash: types.HexToHash("0xec26b0f85ed839ece5f11c4c4e837ec359f5adc4420410f6453b1f6b60fb96a6"),
					Index:  0,
				},
				DepType: types.DepTypeDepGroup,
			},
			{
				OutPoint: &types.OutPoint{
					TxHash: types.HexToHash("0xe12877ebd2c3c364dc46c5c992bcfaf4fee33fa13eebdf82c591fc9825aab769"),
					Index:  0,
				},
				DepType: types.DepTypeCode,
			},
		},
		Inputs: []*types.CellInput{
			{
				Since: 0,
				PreviousOutput: &types.OutPoint{
					TxHash: types.HexToHash("0xc43c8198c4ead3dce957cc3a3ab2ca6c8f4c23ad9d74cb083daefd5d2e4fba4e"),
					Index:  0,
				},
			},
			{
				Since: 0,
				PreviousOutput: &types.OutPoint{
					TxHash: types.HexToHash("0x469100c2149317341756e80f369c94ed2a84b58349ff41985819d49413377ae8"),
					Index:  0,
				},
			},
		},
		Outputs: []*types.CellOutput{
			{
				Capacity: 0xea46318821,
				Lock: &types.Script{
					CodeHash: types.HexToHash("0x3419a1c09eb2567f6552ee7a8ecffd64155cffe0f1796e6e61ec088d740c1356"),
					HashType: types.HashTypeType,
					Args:     common.FromHex("0xa3b8598e1d53e6c5e89e8acb6b4c34d3adb13f2b"),
				},
				Type: &types.Script{
					CodeHash: types.HexToHash("0xc5e5dcf215925f7ef4dfaf5f4b4f105bc321c02776d6e7d52a1db3fcd9d011a4"),
					HashType: types.HashTypeType,
					Args:     common.FromHex("0xc772f4d885ca6285d87d82b8edc1643df9f3ce63c40d0f81f2a38c147328d430"),
				},
			},
		},
		OutputsData: [][]byte{
			common.FromHex("0x00000000000000000000000000000000"),
		},
		Witnesses: [][]byte{
			common.FromHex("0x55000000100000005500000055000000410000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"),
			common.FromHex("0x10000000100000001000000010000000"),
		},
	}
	key, err := s.HexToKey("0x6fc935dad260867c749cf1ba6602d5f5ed7fb1131f1beb65be2d342e912eaafe")
	if err != nil {
		t.Error(err)
	}
	wa := &types.WitnessArgs{
		Lock: make([]byte, 65),
	}
	signature, err := SignTransaction(tx, []int{0, 1}, wa.Serialize(), key)
	if err != nil {
		t.Error(err)
	}
	expectedSignature := common.FromHex("ed0c2ec9523029ed21be22fce92ff158d4da25da0aebd050cdd4b04a9c980ccf5f76afc8d33fa890fcb231bde3eba46b2932d4aaecd4df559ecc3d268d90ef8c01")
	assert.Equal(t, expectedSignature, signature)
}

func TestMsgFromTxForMultiSig(t *testing.T) {
	// https://pudge.explorer.nervos.org/transaction/0x8b9027c407ee95f043b158b4bb5fe685b2e6159723b48712d91ec733b3068a5c
	tx := &types.Transaction{
		Version: 0,
		CellDeps: []*types.CellDep{
			{
				OutPoint: &types.OutPoint{
					TxHash: types.HexToHash("0xf8de3bb47d055cdf460d93a2a6e1b05f7432f9777c8c474abf4eec1d4aee5d37"),
					Index:  1,
				},
				DepType: types.DepTypeDepGroup,
			},
		},
		HeaderDeps: []types.Hash{},
		Inputs: []*types.CellInput{
			{
				Since: 0,
				PreviousOutput: &types.OutPoint{
					TxHash: types.HexToHash("0xb8e52009fb4dc0d63dd2a0547909bb1d66dff83e14645c70b25222c1e04ec593"),
					Index:  0,
				},
			},
		},
		Outputs: []*types.CellOutput{
			{
				Capacity: 10000000000,
				Lock: &types.Script{
					CodeHash: types.HexToHash("0x9bd7e06f3ecf4be0f2fcd2188b23f1b9fcc88e5d4b65a8637b17723bbda3cce8"),
					HashType: types.HashTypeType,
					Args:     common.FromHex("0x6cfd0e42b63a6fccf5eda9cef74d5fd0537fd55a"),
				},
			},
			{
				Capacity: 989900000000,
				Lock: &types.Script{
					CodeHash: types.HexToHash("0x5c5069eb0857efc65e1bca0c07df34c31663b3622fd3876c876320fc9634e2a8"),
					HashType: types.HashTypeType,
					Args:     common.FromHex("0x35ed7b939b4ac9cb447b82340fd8f26d344f7a62"),
				},
			},
		},
		OutputsData: [][]byte{{}, {}},
		Witnesses: [][]byte{
			common.FromHex("0x10000000100000001000000010000000"),
			[]byte{0x12, 0x34},
		},
	}
	multisigScript := common.FromHex("000002029b41c025515b00c24e2e2042df7b221af5c1891fe732dcd15b7618eb1d7a11e6a68e4579b5be0114")
	hash, _ := MsgFromTxForMultiSig(tx, []int{0}, multisigScript)
	key, _ := s.HexToKey("5271b0e474609ee280eb6ba07895718863a0eb8f114afd7217fa371fd48f6941")
	signature, _ := key.Sign(hash)
	expectedSignature := common.FromHex("bf990766e3efa8253c58330f4366bef09f49cbe4efa47b2b491541ad919c90c33ca6763c5780693efd0efea89c1645e8992520e8e551c1dec50fc41fac14b3a401")
	assert.Equal(t, expectedSignature, signature)
}
