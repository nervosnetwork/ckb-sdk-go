package numeric

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCapacity(t *testing.T) {
	assert.True(t, NewCapacity(1234500000000) == NewCapacityFromCKBytes(12345.0))

	assert.Equal(t, uint64(1234500000000), NewCapacity(1234500000000).Shannon())
	assert.Equal(t, 12345.0, NewCapacity(1234500000000).CKBytes())
	assert.Equal(t, uint64(1234500000000), NewCapacityFromCKBytes(12345.0).Shannon())
	assert.Equal(t, 12345.0, NewCapacityFromCKBytes(12345.0).CKBytes())

	assert.Equal(t, uint64(12345000000), NewCapacityFromCKBytes(123.45).Shannon())
	assert.Equal(t, 123.45, NewCapacity(12345000000).CKBytes())

	assert.Equal(t, uint64(12000000), NewCapacityFromCKBytes(0.12).Shannon())
	assert.Equal(t, 0.12, NewCapacity(12000000).CKBytes())

	assert.Equal(t, uint64(0), NewCapacityFromCKBytes(0).Shannon())
	assert.Equal(t, 0.0, NewCapacity(0).CKBytes())
}
