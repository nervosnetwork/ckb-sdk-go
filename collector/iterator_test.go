package collector

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/nervosnetwork/ckb-sdk-go/indexer"
	"github.com/nervosnetwork/ckb-sdk-go/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLiveCellIterator(t *testing.T) {
	client, err := indexer.Dial("https://testnet.ckb.dev/indexer")
	if err != nil {
		t.Error(err)
	}
	i, err := NewLiveCellIteratorFromAddress(client, "ckt1qyqgrfqrklscqeutp3tlqhlcd8xrculgufqspwdp7m")
	if err != nil {
		t.Error(err)
	}
	count := 0
	for i.HasNext() {
		i.Next()
		count += 1
	}
	assert.Equal(t, 10, count)

	// Check outputData
	i, err = NewLiveCellIteratorFromAddress(client, "ckt1qzda0cr08m85hc8jlnfp3zer7xulejywt49kt2rr0vthywaa50xwsqgxc8z84suk20xzx8337sckkkjfqvzk2ysq48gzc")
	count = 0
	var v *types.TransactionInput
	for i.HasNext() {
		v = i.Next()
		count += 1
	}
	assert.Equal(t, 1, count)
	assert.Equal(t, common.FromHex("0x0000000000000000"), v.OutputData)
}