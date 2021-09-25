package types

import (
	"encoding/json"
	"math/big"

)
type TransactionInfo struct {
	TxHash     string               `json:"tx_hash"`
	Records    []*Record            `json:"records"`
	Fee        int64                `json:"fee"`
	Burn       []*BurnInfo          `json:"burn"`
}

type Record struct {
	Id string `json:"id"`
	AddressOrLockHash AddressOrLockHash `json:"address_or_lock_hash"`
	Amount big.Int `json:"amount"`
	Occupied uint64 `json:"occupied"`
	AssetInfo AssetInfo `json:"asset_info"`
	Status RecordStatus `json:"status"`
	Extra *ExtraFilter `json:"extra"`
	BlockNumber BlockNumber `json:"block_number"`
	EpochNumber []byte `json:"epoch_number"`
}

func (r *Record) UnmarshalJSON(data []byte) error {
	type Alias Record
	aux := &struct {
		AddressOrLockHash map[string]string `json:"address_or_lock_hash"`
		Status map[string]BlockNumber `json:"status"`
		*Alias
	}{
		Alias: (*Alias)(r),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	for k, v := range aux.AddressOrLockHash {
		switch k {
			case "Address":
				r.AddressOrLockHash = &Address{v}
			case "LockHash":
				r.AddressOrLockHash = &LockHash{v}
		}
		break
	}
	for k, v := range aux.Status {
		switch k {
			case "Claimable":
				r.Status = &Claimable{v}
			case "Fixed":
				r.Status = &Fixed{v}
		}
		break
	}
	return nil
}

type BurnInfo struct {
	UdtHash   string    `json:"udt_hash"`
	Amount big.Int `json:"amount"`
}

type AddressOrLockHash interface {
	GetAddress() string
}
type Address struct {
	Address string
}

func (addr *Address) GetAddress() string {
	return addr.Address
}

type LockHash struct {
	LockHash string
}
func (addr *LockHash) GetAddress() string {
	return addr.LockHash
}

type RecordStatus interface {
	GetBlockNumber() BlockNumber
}

type Claimable struct {
	BlockNumber BlockNumber
}
func (status *Claimable) GetBlockNumber() BlockNumber {
	return status.BlockNumber
}

type Fixed struct {
	BlockNumber BlockNumber
}
func (status *Fixed) GetBlockNumber() BlockNumber {
	return status.BlockNumber
}


type ExtraFilter string

const (
	Dao ExtraFilter = "Dao"
	CellBase ExtraFilter = "CellBase"
)
