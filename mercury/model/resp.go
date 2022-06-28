package model

import (
	"github.com/nervosnetwork/ckb-sdk-go/types"
	"math/big"
)

type IoType string
type DaoStateType = string
type DBDriver string
type NetworkType string
type SyncState string
type AccountType string

const (
	IoTypeInput          IoType       = "Input"
	IoTypeOutput         IoType       = "Output"
	DaoStateTypeDeposit  DaoStateType = "Deposit"
	DaoStateTypeWithdraw DaoStateType = "Withdraw"

	AccountTypeAcp    AccountType = "Acp"
	AccountTypePwLock AccountType = "PwLock"

	DBDriverPostgreSQL DBDriver = "PostgreSQL"
	DBDriverMySQL      DBDriver = "MySQL"
	DBDriverSQLite     DBDriver = "SQLite"

	NetworkTypeMainnet NetworkType = "Mainnet"
	NetworkTypeTestnet NetworkType = "Testnet"
	NetworkTypeStaging NetworkType = "Staging"
	NetworkTypeDev     NetworkType = "Dev"

	SyncStateReadOnly            SyncState = "ReadOnly"
	SyncStateSerial              SyncState = "Serial"
	SyncStateParallelFirstStage  SyncState = "ParallelFirstStage"
	SyncStateParallelSecondStage SyncState = "ParallelSecondStage"
)

type Balance struct {
	Ownership string     `json:"ownership"`
	AssetInfo *AssetInfo `json:"asset_info"`
	Free      *big.Int   `json:"free"`
	Occupied  *big.Int   `json:"occupied"`
	Frozen    *big.Int   `json:"frozen"`
}

type GetBalanceResponse struct {
	Balances       []*Balance `json:"balances"`
	TipBlockNumber uint64     `json:"tip_block_number"`
}

type DBInfo struct {
	Version   string   `json:"version"`
	DB        DBDriver `json:"db"`
	ConnSize  uint32   `json:"conn_size"`
	CenterId  int64    `json:"center_id"`
	MachineId int64    `json:"machine_id"`
}

type AccountInfo struct {
	AccountNumber  uint32      `json:"account_number"`
	AccountAddress string      `json:"account_address"`
	AccountType    AccountType `json:"account_type"`
}

type BlockInfo struct {
	BlockNumber  uint64             `json:"block_number"`
	BlockHash    types.Hash         `json:"block_hash"`
	ParentHash   types.Hash         `json:"parent_hash"`
	Timestamp    uint64             `json:"timestamp"`
	Transactions []*TransactionInfo `json:"transactions"`
}

type TransactionInfoWrapper struct {
	Type  TransactionType  `json:"type"`
	Value *TransactionInfo `json:"value"`
}

type TransactionViewWrapper struct {
	Type  TransactionType            `json:"type"`
	Value *TransactionWithRichStatus `json:"value"`
}

type TransactionType string

const (
	TransactionTransactionView TransactionType = "TransactionWithRichStatus"
	TransactionTransactionInfo TransactionType = "TransactionInfo"
)

type GetTransactionInfoResponse struct {
	Transaction *TransactionInfo        `json:"transaction,omitempty"`
	Status      types.TransactionStatus `json:"status"`
}

type TransactionInfo struct {
	TxHash    types.Hash  `json:"tx_hash"`
	Records   []*Record   `json:"records"`
	Fee       uint64      `json:"fee"`
	Burn      []*BurnInfo `json:"burn"`
	Timestamp uint64      `json:"timestamp"`
}

type BurnInfo struct {
	UdtHash types.Hash `json:"udt_hash"`
	Amount  *big.Int   `json:"amount"`
}

type Record struct {
	OutPoint    *types.OutPoint `json:"out_point"`
	Ownership   string          `json:"ownership"`
	IoType      IoType          `json:"io_type"`
	Amount      *big.Int        `json:"amount"`
	Occupied    *big.Int        `json:"occupied"`
	AssetInfo   *AssetInfo      `json:"asset_info"`
	Extra       *ExtraFilter    `json:"extra,omitempty"`
	BlockNumber uint64          `json:"block_number"`
	EpochNumber uint64          `json:"epoch_number"`
}

type ExtraFilter struct {
	Type  ExtraFilterType `json:"type"`
	Value *DaoInfo        `json:"value,omitempty"`
}

type DaoInfo struct {
	State  DaoState `json:"state"`
	Reward uint64   `json:"reward"`
}

type DaoState struct {
	Type  DaoStateType `json:"type"`
	Value []uint64     `json:"value"`
}

type TransactionWithRichStatus struct {
	Transaction types.Transaction `json:"transaction,omitempty"`
	TxStatus    TxRichStatus      `json:"tx_status"`
}

type TxRichStatus struct {
	Status    types.TransactionStatus `json:"status"`
	BlockHash types.Hash              `json:"block_hash,omitempty"`
	Reason    string                  `json:"reason,omitempty"`
	Timestamp uint64                  `json:"timestamp,omitempty"`
}

type PaginationResponseTransactionView struct {
	Response   []*TransactionViewWrapper `json:"response"`
	Count      uint64                    `json:"count,omitempty"`
	NextCursor uint64                    `json:"next_cursor,omitempty"`
}

type PaginationResponseTransactionInfo struct {
	Response   []*TransactionInfoWrapper `json:"response"`
	Count      uint64                    `json:"count,omitempty"`
	NextCursor uint64                    `json:"next_cursor,omitempty"`
}

type MercuryInfo struct {
	MercuryVersion    string      `json:"mercury_version"`
	CkbNodeVersion    string      `json:"ckb_node_version"`
	NetworkType       NetworkType `json:"network_type"`
	EnabledExtensions []Extension `json:"enabled_extensions"`
}

type Extension struct {
	Name     string          `json:"name"`
	Scripts  []types.Script  `json:"scripts"`
	CellDeps []types.CellDep `json:"cell_deps"`
}

type MercurySyncState struct {
	State    SyncState `json:"type"`
	SyncInfo struct {
		Current  uint64 `json:"current"`
		Target   uint64 `json:"target"`
		Progress uint64 `json:"progress"`
	} `json:"sync_info,omitempty"`
}
