package types

import "math/big"
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
	Asset_info AssetInfo `json:"asset_info"`
	Status RecordStatus `json:"status"`
	Extra *ExtraFilter `json:"extra"`
	BlockNumber uint64 `json:"block_number"`
	EpochNumber []byte `json:"epoch_number"`
}

type BurnInfo struct {
	UdtHash   string    `json:"udt_hash"`
	Amount big.Int `json:"amount"`
}

type AddressOrLockHash interface {
	GetAddress() string
}

type LockHash struct {
	LockHash string `json:"LockHash"`
}
func (addr *LockHash) GetAddress() string {
	return addr.LockHash
}

type RecordStatus interface {
	GetBlockNumber() uint64
}

type Claimable struct {
	BlockNumber uint64 `json:"Claimable"`
}
func (status *Claimable) GetBlockNumber() uint64 {
	return status.BlockNumber
}

type Fixed struct {
	BlockNumber uint64 `json:"Fixed"`
}
func (status *Fixed) GetBlockNumber() uint64 {
	return status.BlockNumber
}


type ExtraFilter string

const (
	Dao ExtraFilter = "Dao"
	CellBase ExtraFilter = "CellBase"
)
