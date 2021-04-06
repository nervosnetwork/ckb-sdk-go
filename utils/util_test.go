package utils

import (
	"github.com/nervosnetwork/ckb-sdk-go/types"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

func TestGenerateSudtAmount(t *testing.T) {
	amount := big.NewInt(10000000)
	data := GenerateSudtAmount(amount)
	expectedData := []byte{0x80, 0x96, 0x98, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}

	assert.Equal(t, expectedData, data)
}

func TestParseSudtAmountReturnError(t *testing.T) {
	data := []byte{0x80, 0x96}
	_, err := ParseSudtAmount(data)

	assert.Error(t, err)
}

func TestParseSudtAmountWith16BytesData(t *testing.T) {
	data := []byte{0x80, 0xC3, 0xC9, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	amount, err := ParseSudtAmount(data)
	assert.NoError(t, err)
	assert.True(t, big.NewInt(30000000).Cmp(amount) == 0)
}

func Test_calcMaxMatureBlockNumber(t *testing.T) {
	cellbaseMaturity := &types.EpochParams{
		Length: 1,
		Index:  0,
		Number: 4,
	}
	type args struct {
		tipEpoch         *types.EpochParams
		startNumber      uint64
		length           uint64
		cellbaseMaturity *types.EpochParams
	}
	tests := []struct {
		name    string
		args    args
		want    uint64
		wantErr bool
	}{
		{
			name: "at 3 epochs",
			args: args{
				tipEpoch: &types.EpochParams{
					Length: 1800,
					Index:  86,
					Number: 3,
				},
				startNumber:      0,
				length:           3,
				cellbaseMaturity: cellbaseMaturity,
			},
			want:    0,
			wantErr: false,
		},
		{
			name: "at 4 epochs and has index",
			args: args{
				tipEpoch: &types.EpochParams{
					Length: 1800,
					Index:  86,
					Number: 4,
				},
				startNumber:      0,
				length:           1000,
				cellbaseMaturity: cellbaseMaturity,
			},
			want:    47,
			wantErr: false,
		},
		{
			name: "at 4 epochs without index",
			args: args{
				tipEpoch: &types.EpochParams{
					Length: 1800,
					Index:  0,
					Number: 4,
				},
				startNumber:      0,
				length:           1000,
				cellbaseMaturity: cellbaseMaturity,
			},
			want:    0,
			wantErr: false,
		},
		{
			name: "at 5 epochs",
			args: args{
				tipEpoch: &types.EpochParams{
					Length: 1800,
					Index:  900,
					Number: 5,
				},
				startNumber:      2000,
				length:           1000,
				cellbaseMaturity: cellbaseMaturity,
			},
			want:    2500,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := calcMaxMatureBlockNumber(tt.args.tipEpoch, tt.args.startNumber, tt.args.length, tt.args.cellbaseMaturity)
			if (err != nil) != tt.wantErr {
				t.Errorf("calcMaxMatureBlockNumber() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("calcMaxMatureBlockNumber() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseNodeVersion(t *testing.T) {
	type args struct {
		nodeVersion string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		want1   int
		want2   int
		wantErr bool
	}{
		{
			name:    "0.39.1",
			args:    args{"0.39.1 (7f5d486 2021-01-06)"},
			want:    0,
			want1:   39,
			want2:   1,
			wantErr: false,
		},
		{
			name:    "0.39.6",
			args:    args{"0.39.6 (7f5d486 2021-01-06)"},
			want:    0,
			want1:   39,
			want2:   6,
			wantErr: false,
		},
		{
			name:    "1.40.1",
			args:    args{"1.40.1 (7f5d486 2021-01-06)"},
			want:    1,
			want1:   40,
			want2:   1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2, err := ParseNodeVersion(tt.args.nodeVersion)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseNodeVersion() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseNodeVersion() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ParseNodeVersion() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("ParseNodeVersion() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}
