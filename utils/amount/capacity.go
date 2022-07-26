package amount

import "math"

func CkbWithDecimalToShannon(amount float64) uint64 {
	return uint64(amount * float64(exponent(10, 8)))
}

func CkbToShannon(amount uint64) uint64 {
	return amount * exponent(10, 8)
}

func exponent(a, n uint64) uint64 {
	return uint64(math.Pow(float64(a), float64(n)))
}
