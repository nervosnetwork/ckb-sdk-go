package resp

type MercurySyncState struct {
	State    SyncState `json:"type"`
	SyncInfo struct {
		Current  uint64 `json:"current"`
		Target   uint64 `json:"target"`
		Progress uint64 `json:"progress"`
	} `json:"sync_info,omitempty"`
}

type SyncState string

const (
	ReadOnly            SyncState = "ReadOnly"
	Serial              SyncState = "Serial"
	ParallelFirstStage  SyncState = "ParallelFirstStage"
	ParallelSecondStage SyncState = "ParallelSecondStage"
)
