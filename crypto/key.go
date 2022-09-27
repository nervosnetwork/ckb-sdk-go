package crypto

// Key key pair
type Key interface {
	Bytes() []byte
	Sign(data []byte) ([]byte, error)
}

func ZeroBytes(bytes []byte) {
	for i := range bytes {
		bytes[i] = 0
	}
}
