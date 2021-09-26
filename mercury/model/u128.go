package model

import (
	"encoding/json"
	"math/big"
)

type U128 struct {
	big.Int
}

func NewU128WithString(value string) *U128 {
	u128 := new(U128)
	u128.SetString(value, 0)

	return u128
}

func NewU128WithU64(value uint64) *U128 {
	u128 := new(U128)
	u128.SetUint64(value)

	return u128
}

func NewU128WithBigInt(value *big.Int) *U128 {
	u128 := new(U128)
	u128.SetString(value.String(), 0)

	return u128
}

func (u128 *U128) MarshalJSON() ([]byte, error) {
	return json.Marshal(u128.String())
}

func (u128 *U128) UnmarshalJSON(bytes []byte) error {
	var amount string
	json.Unmarshal(bytes, &amount)
	u128.SetString(amount, 0)

	return nil
}
