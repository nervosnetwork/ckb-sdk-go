package common

import (
	"encoding/json"
)

type AddressOrLockHash struct {
	AddressOrLockHash     string
	AddressOrLockHashType AddressOrLockHashType
}

type AddressOrLockHashType string

const (
	Address  AssetType = "Address"
	LockHash           = "LockHash"
)

func (address *AddressOrLockHash) UnmarshalJSON(bytes []byte) error {
	addressOrLockHash := make(map[string]string)
	json.Unmarshal(bytes, &addressOrLockHash)

	if _, ok := addressOrLockHash["Address"]; ok {
		address.AddressOrLockHash = addressOrLockHash["Address"]
		address.AddressOrLockHashType = AddressOrLockHashType(Address)
	} else {
		address.AddressOrLockHash = addressOrLockHash["LockHash"]
		address.AddressOrLockHashType = LockHash
	}

	return nil
}
