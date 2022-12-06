package blake2b

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBlake160(t *testing.T) {
	in := "Simple one-hash test"
	good := "88ac9acbbb56403cb0fb7d45f158f0accbcd30da"
	assert.Equal(t, good, fmt.Sprintf("%x", Blake160([]byte(in))))
}

func TestBlake256(t *testing.T) {
	in := "Simple one-hash test"
	good := "88ac9acbbb56403cb0fb7d45f158f0accbcd30da7d756c4f811e24ae110285bb"
	assert.Equal(t, good, fmt.Sprintf("%x", Blake256([]byte(in))))
}
