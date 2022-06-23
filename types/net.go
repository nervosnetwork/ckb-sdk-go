package types

type NodeAddress struct {
	Address string `json:"address"`
	Score   uint64 `json:"score"`
}

type LocalNode struct {
	Version     string               `json:"version"`
	NodeId      string               `json:"node_id"`
	Active      bool                 `json:"active"`
	Addresses   []*NodeAddress       `json:"addresses"`
	Protocols   []*LocalNodeProtocol `json:"protocols"`
	Connections uint64               `json:"connections"`
}

type LocalNodeProtocol struct {
	Id              uint64   `json:"id"`
	Name            string   `json:"name"`
	SupportVersions []string `json:"support_versions"`
}

type RemoteNode struct {
	Version           string                `json:"version"`
	NodeID            string                `json:"node_id"`
	Addresses         []*NodeAddress        `json:"addresses"`
	IsOutbound        bool                  `json:"is_outbound"`
	ConnectedDuration uint64                `json:"connected_duration"`
	LastPingDuration  uint64                `json:"last_ping_duration,omitempty"`
	SyncState         *PeerSyncState        `json:"sync_state,omitempty"`
	Protocols         []*RemoteNodeProtocol `json:"protocols"`
}

type RemoteNodeProtocol struct {
	ID      uint64 `json:"id"`
	Version string `json:"version"`
}

type PeerSyncState struct {
	BestKnownHeaderHash    Hash   `json:"best_known_header_hash,omitempty"`
	BestKnownHeaderNumber  uint64 `json:"best_known_header_number,omitempty"`
	LastCommonHeaderHash   Hash   `json:"last_common_header_hash,omitempty"`
	LastCommonHeaderNumber uint64 `json:"last_common_header_number,omitempty"`
	UnknownHeaderListSize  uint64 `json:"unknown_header_list_size"`
	InflightCount          uint64 `json:"inflight_count"`
	CanFetchCount          uint64 `json:"can_fetch_count"`
}

type BannedAddress struct {
	Address   string `json:"address"`
	BanReason string `json:"ban_reason"`
	BanUntil  uint64 `json:"ban_until"`
	CreatedAt uint64 `json:"created_at"`
}
