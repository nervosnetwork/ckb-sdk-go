package amount

import "math"

func CkbToShannon(amount uint64) uint64 {
	return amount * exponent(10, 8)
}

func exponent(a, n uint64) uint64 {
	return uint64(math.Pow(float64(a), float64(n)))
}
