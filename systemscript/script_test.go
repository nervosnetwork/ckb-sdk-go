package systemscript

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/nervosnetwork/ckb-sdk-go/types"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

func TestMultisigConfigEncode(t *testing.T) {
	config := &MultisigConfig{
		Version:    0,
		FirstN:     0,
		Threshold:  2,
		KeysHashes: getKeysHashes(),
	}
	encoded := config.Encode()
	assert.Equal(t, common.FromHex("0x000002029b41c025515b00c24e2e2042df7b221af5c1891fe732dcd15b7618eb1d7a11e6a68e4579b5be0114"), encoded)
	hash := config.Hash160()
	assert.Equal(t, common.FromHex("0x35ed7b939b4ac9cb447b82340fd8f26d344f7a62"), hash)
}

func TestMultisigConfigDecode(t *testing.T) {
	bytes := common.FromHex("0x000002029b41c025515b00c24e2e2042df7b221af5c1891fe732dcd15b7618eb1d7a11e6a68e4579b5be0114")
	config, err := DecodeToMultisigConfig(bytes)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, byte(0), config.FirstN)
	assert.Equal(t, byte(2), config.Threshold)
	assert.Equal(t, getKeysHashes(), config.KeysHashes)

	bytes = common.FromHex("0x000002039b41c025515b00c24e2e2042df7b221af5c1891fe732dcd15b7618eb1d7a11e6a68e4579b5be0114")
	_, err = DecodeToMultisigConfig(bytes)
	assert.Error(t, err)

	bytes = common.FromHex("0x000002029b41c025515b00c24e2e2042df7b221af5c1891f")
	_, err = DecodeToMultisigConfig(bytes)
	assert.Error(t, err)
}

func getKeysHashes() [][20]byte {
	keysHashes := make([][20]byte, 2)
	copy(keysHashes[0][:], common.FromHex("0x9b41c025515b00c24e2e2042df7b221af5c1891f"))
	copy(keysHashes[1][:], common.FromHex("0xe732dcd15b7618eb1d7a11e6a68e4579b5be0114"))
	return keysHashes
}

func TestDecodeSudtAmount(t *testing.T) {
	data := hexutil.MustDecode("0x0010a5d4e80000000000000000000000")
	amount, err := DecodeSudtAmount(data)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, big.NewInt(1000000000000), amount)
}

func TestEncodeSudtAmount(t *testing.T) {
	data := EncodeSudtAmount(big.NewInt(1000000000000))
	assert.Equal(t, hexutil.MustDecode("0x0010a5d4e80000000000000000000000"), data)
}

func TestChequeArgs(t *testing.T) {
	senderLock := &types.Script{
		CodeHash: types.HexToHash("0x9bd7e06f3ecf4be0f2fcd2188b23f1b9fcc88e5d4b65a8637b17723bbda3cce8"),
		HashType: "type",
		Args:     common.FromHex("0xedcda9513fa030ce4308e29245a22c022d0443bb"),
	}
	receiverLock := &types.Script{
		CodeHash: types.HexToHash("0x9bd7e06f3ecf4be0f2fcd2188b23f1b9fcc88e5d4b65a8637b17723bbda3cce8"),
		HashType: "type",
		Args:     common.FromHex("0xedcda9513fa030ce4308e29245a22c022d0212ab"),
	}

	senderLockHash := senderLock.Hash()
	receiverLockHash := receiverLock.Hash()
	expectedArgs := append(receiverLockHash.Bytes()[0:20], senderLockHash.Bytes()[0:20]...)
	actualArgs := ChequeArgs(senderLock, receiverLock)
	assert.Equal(t, expectedArgs, actualArgs)
}
