package crypto

// Key key pair
type Key interface {
	Bytes() []byte
	Sign(data []byte) ([]byte, error)
}
