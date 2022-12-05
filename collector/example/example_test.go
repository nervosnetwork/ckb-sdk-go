package example

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOffChainTransaction(t *testing.T) {

	assert.Equal(t, nil, SendChainedTransactionExample())
}
