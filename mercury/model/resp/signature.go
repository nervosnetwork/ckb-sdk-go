package resp

type HashAlgorithm string

const (
	Blake2b HashAlgorithm = "Blake2b"
)

type SignAlgorithm string

const (
	Secp256k1 SignAlgorithm = "Secp256k1"
)

type SignatureLocation struct {
	Index  int `json:"index"`
	Offset int `json:"offset"`
}

type SignatureInfo struct {
	Algorithm SignAlgorithm `json:"algorithm"`
	Address   string        `json:"address"`
}
type SignatureAction struct {
	SignatureLocation   *SignatureLocation `json:"signature_location"`
	SignatureInfo       *SignatureInfo     `json:"signature_info"`
	HashAlgorithm       *HashAlgorithm     `json:"hash_algorithm"`
	OtherIndexesInGroup []int              `json:"other_indexes_in_group"`
}
