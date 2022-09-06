package amount

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCkbToShannon(t *testing.T) {
	assert.Equal(t, uint64(234300000000), CkbToShannon(2343))
}
