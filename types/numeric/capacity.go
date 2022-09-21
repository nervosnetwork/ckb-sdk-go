package numeric

type Capacity uint64

const scale = 100000000

func NewCapacity(shannon uint64) Capacity {
	return Capacity(shannon)
}

func NewCapacityFromCKBytes(ckBytes float64) Capacity {
	return Capacity(ckBytes * scale)
}

func (c Capacity) Shannon() uint64 {
	return uint64(c)
}

func (c Capacity) CKBytes() float64 {
	return float64(c) / scale
}
