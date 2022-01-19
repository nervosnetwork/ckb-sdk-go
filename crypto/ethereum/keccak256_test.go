package ethereum

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDigest(t *testing.T) {
	bytes := []byte("1")
	digest, _ := Keccak256(bytes)
	assert.Equal(t, "c89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc6", hex.EncodeToString(digest))

	bytes = []byte("1ddeo39%7")
	digest, _ = Keccak256(bytes)
	assert.Equal(t, "488471036d1f58480c9bf3e29154c833dc234daea981b09cd35f632fec20a8ec", hex.EncodeToString(digest))
}
