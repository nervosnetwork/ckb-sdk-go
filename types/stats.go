package types

import (
	"encoding/json"
	"fmt"
	"math/big"
)

//	pub struct Alert {
//	    /// The identifier of the alert. Clients use id to filter duplicated alerts.
//	    pub id: AlertId,
//	    /// Cancel a previous sent alert.
//	    pub cancel: AlertId,
//	    /// Optionally set the minimal version of the target clients.
//	    ///
//	    /// See [Semantic Version](https://semver.org/) about how to specify a version.
//	    pub min_version: Option<String>,
//	    /// Optionally set the maximal version of the target clients.
//	    ///
//	    /// See [Semantic Version](https://semver.org/) about how to specify a version.
//	    pub max_version: Option<String>,
//	    /// Alerts are sorted by priority, highest first.
//	    pub priority: AlertPriority,
//	    /// The alert is expired after this timestamp.
//	    pub notice_until: Timestamp,
//	    /// Alert message.
//	    pub message: String,
//	    /// The list of required signatures.
//	    pub signatures: Vec<JsonBytes>,
//	}
type Alert struct {
	Id          uint32            `json:"id"`
	Cancel      uint32            `json:"cancel"`
	MinVersion  *string           `json:"min_version"`
	MaxVersion  *string           `json:"max_version"`
	Priority    uint32            `json:"priority"`
	NoticeUntil uint64            `json:"notice_until"`
	Message     string            `json:"message"`
	Signatures  []json.RawMessage `json:"signatures"`
}

type AlertMessage struct {
	Id          uint32 `json:"id"`
	Message     string `json:"message"`
	NoticeUntil uint64 `json:"notice_until"`
	Priority    uint32 `json:"priority"`
}

type BlockchainInfo struct {
	Alerts                 []*AlertMessage `json:"alerts"`
	Chain                  string          `json:"chain"`
	Difficulty             *big.Int        `json:"difficulty"`
	Epoch                  uint64          `json:"epoch"`
	IsInitialBlockDownload bool            `json:"is_initial_block_download"`
	MedianTime             uint64          `json:"median_time"`
}

// DeploymentState represents the possible states of a deployment.
type DeploymentState int

const (
	// Defined is the first state that each softfork starts.
	Defined DeploymentState = iota
	// Started is the state for epochs past the `start` epoch.
	Started
	// LockedIn is the state for epochs after the first epoch period with STARTED epochs of which at least `threshold` has the associated bit set in `version`.
	LockedIn
	// Active is the state for all epochs after the LOCKED_IN epoch.
	Active
	// Failed is the state for epochs past the `timeout_epoch`, if LOCKED_IN was not reached.
	Failed
)

// DeploymentPos represents the possible positions for deployments.
type DeploymentPos int

const (
	// Testdummy represents a dummy deployment.
	Testdummy DeploymentPos = iota
	// LightClient represents the light client protocol deployment.
	LightClient
)

// convert DeploymentPos to string
func (d DeploymentPos) String() string {
	switch d {
	case Testdummy:
		return "Testdummy"
	case LightClient:
		return "LightClient"
	}
	return "Unknown"
}

// convert string to DeploymentPos
func DeploymentPosFromString(s string) (DeploymentPos, error) {
	switch s {
	case "Testdummy":
		return Testdummy, nil
	case "LightClient":
		return LightClient, nil
	}
	return -1, fmt.Errorf("invalid DeploymentPos: %s", s)
}

// DeploymentInfo represents information about a deployment.
type DeploymentInfo struct {
	Bit                uint8
	Start              uint64
	Timeout            uint64
	MinActivationEpoch uint64
	Period             uint64
	Threshold          jsonRationalU256
	Since              uint64
	State              DeploymentState
}

// DeploymentsInfo represents information about multiple deployments.
type DeploymentsInfo struct {
	Hash        Hash
	Epoch       uint64
	Deployments map[DeploymentPos]DeploymentInfo
}
