package numeric

// define some useful const
// https://github.com/nervosnetwork/ckb/blob/35392279150fe4e61b7904516be91bda18c46f05/test/src/utils.rs#L24

type Since uint64

const (
	FlagSinceRelative    = 0x8000000000000000
	FlagSinceEpochNumber = 0x2000000000000000
	FlagSinceBlockNumber = 0x0
	FlagSinceTimestamp   = 0x4000000000000000
)

func NewSinceFromRelativeBlockNumber(blockNumber uint64) Since {
	return Since(FlagSinceRelative | FlagSinceBlockNumber | blockNumber)
}

func NewSinceFromAbsoluteBlockNumber(blockNumber uint64) Since {
	return Since(FlagSinceBlockNumber | blockNumber)
}

func NewSinceFromRelativeEpochNumber(epochNumber uint64) Since {
	return Since(FlagSinceRelative | FlagSinceEpochNumber | epochNumber)
}

func NewSinceFromAbsoluteEpochNumber(epochNumber uint64) Since {
	return Since(FlagSinceEpochNumber | epochNumber)
}

func NewSinceFromRelativeTimestamp(timestamp uint64) Since {
	return Since(FlagSinceRelative | FlagSinceTimestamp | timestamp)
}

func NewSinceFromAbsoluteTimestamp(timestamp uint64) Since {
	return Since(FlagSinceTimestamp | timestamp)
}
