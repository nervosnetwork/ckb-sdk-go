package amount

import (
	"github.com/ethereum/go-ethereum/common/math"
	"math/big"
)

func CkbWithDecimalToShannon(amount float64) *big.Int {
	ckbAmount := big.NewFloat(amount)

	pow := new(big.Float)
	pow.SetInt(math.BigPow(10, 8))

	u, _ := ckbAmount.Mul(ckbAmount, pow).Uint64()
	result := new(big.Int)
	result.SetUint64(u)
	return result
}

func CkbToShannon(amount uint64) *big.Int {
	ckbAmount := new(big.Int)
	ckbAmount.SetUint64(amount)

	pow := math.BigPow(10, 8)

	result := ckbAmount.Mul(ckbAmount, pow)

	return result
}
