package utils

import (
	"context"
	"github.com/nervosnetwork/ckb-sdk-go/mocks"
	"github.com/nervosnetwork/ckb-sdk-go/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSystemScripts(t *testing.T) {
	tests := []struct {
		name  string
		chain string
		want  *SystemScripts
	}{
		{
			"test NewSystemScripts for ckb mainnet",
			"ckb",
			&SystemScripts{
				SecpSingleSigCell: &SystemScriptCell{
					CellHash: types.HexToHash("0x9bd7e06f3ecf4be0f2fcd2188b23f1b9fcc88e5d4b65a8637b17723bbda3cce8"),
					OutPoint: &types.OutPoint{
						TxHash: types.HexToHash("0x71a7ba8fc96349fea0ed3a5c47992e3b4084b031a42264a018e0072e8172e46c"),
						Index:  0,
					},
				},
				SecpMultiSigCell: &SystemScriptCell{
					CellHash: types.HexToHash("0x5c5069eb0857efc65e1bca0c07df34c31663b3622fd3876c876320fc9634e2a8"),
					OutPoint: &types.OutPoint{
						TxHash: types.HexToHash("0x71a7ba8fc96349fea0ed3a5c47992e3b4084b031a42264a018e0072e8172e46c"),
						Index:  1,
					},
				},
				DaoCell: &SystemScriptCell{
					CellHash: types.HexToHash("0x82d76d1b75fe2fd9a27dfbaa65a039221a380d76c926f378d3f81cf3e7e13f2e"),
					OutPoint: &types.OutPoint{
						TxHash: types.HexToHash("0xe2fb199810d49a4d8beec56718ba2593b665db9d52299a0f9e6e75416d73ff5c"),
						Index:  2,
					},
				},
				ACPCell: &SystemScriptCell{
					CellHash: types.HexToHash("0xd369597ff47f29fbc0d47d2e3775370d1250b85140c670e4718af712983a2354"),
					OutPoint: &types.OutPoint{
						TxHash: types.HexToHash("0x4153a2014952d7cac45f285ce9a7c5c0c0e1b21f2d378b82ac1433cb11c25c4d"),
						Index:  0,
					},
				},
				SUDTCell: &SystemScriptCell{
					CellHash: types.HexToHash("0x5e7a36a77e68eecc013dfa2fe6a23f3b6c344b04005808694ae6dd45eea4cfd5"),
					OutPoint: &types.OutPoint{
						TxHash: types.HexToHash("0xc7813f6a415144643970c2e88e0bb6ca6a8edc5dd7c1022746f628284a9936d5"),
						Index:  0,
					},
				},
			},
		},
		{
			"test NewSystemScripts for ckb testnet",
			"ckb_testnet",
			&SystemScripts{
				SecpSingleSigCell: &SystemScriptCell{
					CellHash: types.HexToHash("0x9bd7e06f3ecf4be0f2fcd2188b23f1b9fcc88e5d4b65a8637b17723bbda3cce8"),
					OutPoint: &types.OutPoint{
						TxHash: types.HexToHash("0x71a7ba8fc96349fea0ed3a5c47992e3b4084b031a42264a018e0072e8172e46c"),
						Index:  0,
					},
				},
				SecpMultiSigCell: &SystemScriptCell{
					CellHash: types.HexToHash("0x5c5069eb0857efc65e1bca0c07df34c31663b3622fd3876c876320fc9634e2a8"),
					OutPoint: &types.OutPoint{
						TxHash: types.HexToHash("0x71a7ba8fc96349fea0ed3a5c47992e3b4084b031a42264a018e0072e8172e46c"),
						Index:  1,
					},
				},
				DaoCell: &SystemScriptCell{
					CellHash: types.HexToHash("0x82d76d1b75fe2fd9a27dfbaa65a039221a380d76c926f378d3f81cf3e7e13f2e"),
					OutPoint: &types.OutPoint{
						TxHash: types.HexToHash("0xe2fb199810d49a4d8beec56718ba2593b665db9d52299a0f9e6e75416d73ff5c"),
						Index:  2,
					},
				},
				ACPCell: &SystemScriptCell{
					CellHash: types.HexToHash("0x3419a1c09eb2567f6552ee7a8ecffd64155cffe0f1796e6e61ec088d740c1356"),
					OutPoint: &types.OutPoint{
						TxHash: types.HexToHash("0xec26b0f85ed839ece5f11c4c4e837ec359f5adc4420410f6453b1f6b60fb96a6"),
						Index:  0,
					},
				},
				SUDTCell: &SystemScriptCell{
					CellHash: types.HexToHash("0x48dbf59b4c7ee1547238021b4869bceedf4eea6b43772e5d66ef8865b6ae7212"),
					OutPoint: &types.OutPoint{
						TxHash: types.HexToHash("0xc1b2ae129fad7465aaa9acc9785f842ba3e6e8b8051d899defa89f5508a77958"),
						Index:  0,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockClient := &mocks.Client{}
			mockClient.On("GetBlockchainInfo", context.Background()).Return(&types.BlockchainInfo{Chain: tt.chain}, nil)
			got, _ := NewSystemScripts(mockClient)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestACPCell(t *testing.T) {
	type args struct {
		acpCell *SystemScriptCell
	}

	tests := []struct {
		name  string
		args  args
		chain string
		want  *SystemScriptCell
	}{
		{
			"set custom acp cell on mainnet",
			args{&SystemScriptCell{
				CellHash: types.HexToHash("0x683574c1275eb5cfe6f8745faa375b08bf773223fd8d2b4db28dbd90a27f1586"),
				OutPoint: &types.OutPoint{
					TxHash: types.HexToHash("0x7d0ecdb8bad4064788b67dfafe71757e7caa2ad2cbe5597a02df95f8792bdb21"),
					Index:  0,
				},
			}},
			"ckb",
			&SystemScriptCell{
				CellHash: types.HexToHash("0x683574c1275eb5cfe6f8745faa375b08bf773223fd8d2b4db28dbd90a27f1586"),
				OutPoint: &types.OutPoint{
					TxHash: types.HexToHash("0x7d0ecdb8bad4064788b67dfafe71757e7caa2ad2cbe5597a02df95f8792bdb21"),
					Index:  0,
				},
			},
		},
		{
			"set custom acp cell on testnet",
			args{&SystemScriptCell{
				CellHash: types.HexToHash("0x683574c1275eb5cfe6f8745faa375b08bf773223fd8d2b4db28dbd90a27f1586"),
				OutPoint: &types.OutPoint{
					TxHash: types.HexToHash("0x7d0ecdb8bad4064788b67dfafe71757e7caa2ad2cbe5597a02df95f8792bdb21"),
					Index:  0,
				},
			}},
			"ckb_testnet",
			&SystemScriptCell{
				CellHash: types.HexToHash("0x683574c1275eb5cfe6f8745faa375b08bf773223fd8d2b4db28dbd90a27f1586"),
				OutPoint: &types.OutPoint{
					TxHash: types.HexToHash("0x7d0ecdb8bad4064788b67dfafe71757e7caa2ad2cbe5597a02df95f8792bdb21"),
					Index:  0,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockClient := &mocks.Client{}
			mockClient.On("GetBlockchainInfo", context.Background()).Return(&types.BlockchainInfo{Chain: tt.chain}, nil)
			s, _ := NewSystemScripts(mockClient, ACPCell(tt.args.acpCell))
			assert.Equal(t, tt.want, s.ACPCell)
		})
	}
}

func TestSUDTCell(t *testing.T) {
	type args struct {
		acpCell *SystemScriptCell
	}

	tests := []struct {
		name  string
		args  args
		chain string
		want  *SystemScriptCell
	}{
		{
			"set custom sudt cell on mainnet",
			args{&SystemScriptCell{
				CellHash: types.HexToHash("0x683574c1275eb5cfe6f8745faa375b08bf773223fd8d2b4db28dbd90a27f1586"),
				OutPoint: &types.OutPoint{
					TxHash: types.HexToHash("0x7d0ecdb8bad4064788b67dfafe71757e7caa2ad2cbe5597a02df95f8792bdb21"),
					Index:  0,
				},
			}},
			"ckb",
			&SystemScriptCell{
				CellHash: types.HexToHash("0x683574c1275eb5cfe6f8745faa375b08bf773223fd8d2b4db28dbd90a27f1586"),
				OutPoint: &types.OutPoint{
					TxHash: types.HexToHash("0x7d0ecdb8bad4064788b67dfafe71757e7caa2ad2cbe5597a02df95f8792bdb21"),
					Index:  0,
				},
			},
		},
		{
			"set custom sudt cell on testnet",
			args{&SystemScriptCell{
				CellHash: types.HexToHash("0x683574c1275eb5cfe6f8745faa375b08bf773223fd8d2b4db28dbd90a27f1586"),
				OutPoint: &types.OutPoint{
					TxHash: types.HexToHash("0x7d0ecdb8bad4064788b67dfafe71757e7caa2ad2cbe5597a02df95f8792bdb21"),
					Index:  0,
				},
			}},
			"ckb_testnet",
			&SystemScriptCell{
				CellHash: types.HexToHash("0x683574c1275eb5cfe6f8745faa375b08bf773223fd8d2b4db28dbd90a27f1586"),
				OutPoint: &types.OutPoint{
					TxHash: types.HexToHash("0x7d0ecdb8bad4064788b67dfafe71757e7caa2ad2cbe5597a02df95f8792bdb21"),
					Index:  0,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockClient := &mocks.Client{}
			mockClient.On("GetBlockchainInfo", context.Background()).Return(&types.BlockchainInfo{Chain: tt.chain}, nil)
			s, _ := NewSystemScripts(mockClient, SUDTCell(tt.args.acpCell))
			assert.Equal(t, tt.want, s.SUDTCell)
		})
	}
}
