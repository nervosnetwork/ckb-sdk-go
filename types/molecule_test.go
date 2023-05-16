package types_test

import (
	"math/rand"
	"testing"

	"github.com/nervosnetwork/ckb-sdk-go/v2/types"
	"github.com/stretchr/testify/assert"
)

func TestScriptPacking(t *testing.T) {
	rng := rand.New(rand.NewSource(1337))
	for _, ty := range []types.ScriptHashType{types.HashTypeData, types.HashTypeType, types.HashTypeData1} {
		script := newScript(rng, ty)
		packed := types.UnpackScript(script.Pack())
		assert.Equal(t, script, packed, "script packing/unpacking should satisfy the round-trip property")
	}
}

func newScript(rng *rand.Rand, ht types.ScriptHashType) *types.Script {
	hash := [32]byte{}
	rng.Read(hash[:])
	argsLen := rng.Intn(4096)
	args := make([]byte, argsLen)
	rng.Read(args)
	return &types.Script{
		CodeHash: hash,
		HashType: ht,
		Args:     args,
	}
}
