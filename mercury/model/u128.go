package model

import (
	"encoding/json"
	"math/big"
)

type U128 struct {
	big.Int
}

func (u128 *U128) UnmarshalJSON(bytes []byte) error {
	var amount string
	json.Unmarshal(bytes, &amount)
	u128.SetString(amount, 0)

	return nil
}
