package resp

type HashAlgorithm int

const (
	Blake2b HashAlgorithm = iota
)

type SignAlgorithm int

const (
	Secp256k1 SignAlgorithm = iota
)

type SignatureLocation struct {
	Index  int
	Offset int
}

type SignatureInfo struct {
	Algorithm SignAlgorithm
	Address   string
}
type SignatureAction struct {
	SignatureLocation   *SignatureLocation `json:"signature_location"`
	SignatureInfo       *SignatureInfo     `json:"signature_info"`
	HashAlgorithm       *HashAlgorithm     `json:"hash_algorithm"`
	OtherIndexesInGroup []int              `json:"other_indexes_in_group"`
}
