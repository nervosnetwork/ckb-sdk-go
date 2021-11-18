package common

type Ownership struct {
	Type  OwnershipType `json:"type"`
	Value string        `json:"value"`
}

type OwnershipType string

const (
	OwnershipAddress  OwnershipType = "Address"
	OwnershipLockhash               = "LockHash"
)
