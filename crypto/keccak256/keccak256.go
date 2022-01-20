package keccak256

import (
	"github.com/ethereum/go-ethereum/crypto"
)

var ckbHashPersonalization = []byte("ckb-default-hash")

func Keccak256(data []byte) ([]byte, error) {
	hash := crypto.Keccak256Hash(data)
	return hash[:], nil
}
