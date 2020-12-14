package utils

import "testing"

func TestSinceFromAbsoluteBlockNumber(t *testing.T) {
	type args struct {
		blockNumber uint64
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{
			"absolute block number",
			args{
				blockNumber: 12345,
			},
			12345,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SinceFromAbsoluteBlockNumber(tt.args.blockNumber); got != tt.want {
				t.Errorf("SinceFromAbsoluteBlockNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSinceFromAbsoluteEpochNumber(t *testing.T) {
	type args struct {
		epochNumber uint64
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{
			"absolute epoch number",
			args{
				epochNumber: 1024,
			},
			2305843009213694976,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SinceFromAbsoluteEpochNumber(tt.args.epochNumber); got != tt.want {
				t.Errorf("SinceFromAbsoluteEpochNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSinceFromAbsoluteTimestamp(t *testing.T) {
	type args struct {
		timestamp uint64
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{
			"absolute timestamp",
			args{
				timestamp: 1585699200,
			},
			0x400000005e83d980,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SinceFromAbsoluteTimestamp(tt.args.timestamp); got != tt.want {
				t.Errorf("SinceFromAbsoluteTimestamp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSinceFromRelativeBlockNumber(t *testing.T) {
	type args struct {
		blockNumber uint64
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{
			"relative block number",
			args{
				blockNumber: 100,
			},
			9223372036854775908,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SinceFromRelativeBlockNumber(tt.args.blockNumber); got != tt.want {
				t.Errorf("SinceFromRelativeBlockNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSinceFromRelativeEpochNumber(t *testing.T) {
	type args struct {
		epochNumber uint64
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{
			"relative epoch number",
			args{
				epochNumber: 6,
			},
			11529215046068469766,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SinceFromRelativeEpochNumber(tt.args.epochNumber); got != tt.want {
				t.Errorf("SinceFromRelativeEpochNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSinceFromRelativeTimestamp(t *testing.T) {
	type args struct {
		timestamp uint64
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{
			"relative timestamp",
			args{
				timestamp: 14 * 24 * 60 * 60,
			},
			0xc000000000127500,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SinceFromRelativeTimestamp(tt.args.timestamp); got != tt.want {
				t.Errorf("SinceFromRelativeTimestamp() = %v, want %v", got, tt.want)
			}
		})
	}
}
