package numeric

import "testing"

func TestSinceFromAbsoluteBlockNumber(t *testing.T) {
	type args struct {
		blockNumber uint64
	}
	tests := []struct {
		name string
		args args
		want Since
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
			if got := NewSinceFromAbsoluteBlockNumber(tt.args.blockNumber); got != tt.want {
				t.Errorf("NewSinceFromAbsoluteBlockNumber() = %v, want %v", got, tt.want)
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
		want Since
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
			if got := NewSinceFromAbsoluteEpochNumber(tt.args.epochNumber); got != tt.want {
				t.Errorf("NewSinceFromAbsoluteEpochNumber() = %v, want %v", got, tt.want)
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
		want Since
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
			if got := NewSinceFromAbsoluteTimestamp(tt.args.timestamp); got != tt.want {
				t.Errorf("NewSinceFromAbsoluteTimestamp() = %v, want %v", got, tt.want)
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
		want Since
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
			if got := NewSinceFromRelativeBlockNumber(tt.args.blockNumber); got != tt.want {
				t.Errorf("NewSinceFromRelativeBlockNumber() = %v, want %v", got, tt.want)
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
		want Since
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
			if got := NewSinceFromRelativeEpochNumber(tt.args.epochNumber); got != tt.want {
				t.Errorf("NewSinceFromRelativeEpochNumber() = %v, want %v", got, tt.want)
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
		want Since
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
			if got := NewSinceFromRelativeTimestamp(tt.args.timestamp); got != tt.want {
				t.Errorf("NewSinceFromRelativeTimestamp() = %v, want %v", got, tt.want)
			}
		})
	}
}
