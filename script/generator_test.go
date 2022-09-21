package script

import (
	"encoding/hex"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/nervosnetwork/ckb-sdk-go/crypto/blake2b"
	"github.com/nervosnetwork/ckb-sdk-go/crypto/secp256k1"
	"github.com/nervosnetwork/ckb-sdk-go/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSecp256K1Blake160SignhashAll(t *testing.T) {
	key, err := secp256k1.HexToKey("e79f3207ea4980b7fed79956d5934249ceac4751a4fae01a0f7c4a96884bc4e3")
	if err != nil {
		t.Error(err)
	}
	generated := Secp256K1Blake160SignhashAll(key)
	expected := &types.Script{
		CodeHash: types.HexToHash("0x9bd7e06f3ecf4be0f2fcd2188b23f1b9fcc88e5d4b65a8637b17723bbda3cce8"),
		HashType: types.HashTypeType,
		Args:     common.FromHex("0x36c329ed630d6ce750712a477543672adab57f4c"),
	}
	assert.Equal(t, expected, generated)

	generated, err = Secp256K1Blake160SignhashAllByPublicKey(common.FromHex("0x024a501efd328e062c8675f2365970728c859c592beeefd6be8ead3d901330bc01"))
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, expected, generated)
	// test public key hex without 0x
	generated, err = Secp256K1Blake160SignhashAllByPublicKey(common.FromHex("024a501efd328e062c8675f2365970728c859c592beeefd6be8ead3d901330bc01"))
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, expected, generated)
	// test invalid length
	_, err = Secp256K1Blake160SignhashAllByPublicKey(common.FromHex("0x024a501ef"))
	assert.NotNil(t, err)
	// test uncompressed public key
	_, err = Secp256K1Blake160SignhashAllByPublicKey(common.FromHex("0x044a501efd328e062c8675f2365970728c859c592beeefd6be8ead3d901330bc01d1868c7dabbf50e52ca7311e1263f917a8ced1d033e82dc2a68bed69397382f4"))
	assert.NotNil(t, err)
}

func TestSecp256k1Blake160Multisig(t *testing.T) {
	encodedPublicKeys := []string{
		"032edb83018b57ddeb9bcc7287c5cc5da57e6e0289d31c9e98cb361e88678d6288",
		"033aeb3fdbfaac72e9e34c55884a401ee87115302c146dd9e314677d826375dc8f",
		"029a685b8206550ea1b600e347f18fd6115bffe582089d3567bec7eba57d04df01",
	}
	multisigConfig := NewMultisigConfig(0, 2)
	for _, publicKeyHex := range encodedPublicKeys {
		key, err := hex.DecodeString(publicKeyHex)
		if err != nil {
			t.Error(t, err)
		}
		multisigConfig.AddKeyHash(blake2b.Blake256(key))
	}
	s, err := Secp256k1Blake160Multisig(multisigConfig)
	if err != nil {
		t.Error(t, err)
	}

	// ckt1qpw9q60tppt7l3j7r09qcp7lxnp3vcanvgha8pmvsa3jplykxn32sq0sfnkgf0ph76pkzwld9ujzex4pkeuwnlsdc5tqu
	assert.Equal(t, hexutil.MustDecode("0x5c5069eb0857efc65e1bca0c07df34c31663b3622fd3876c876320fc9634e2a8"), s.CodeHash.Bytes())
	assert.Equal(t, hexutil.MustDecode("0xf04cec84bc37f683613bed2f242c9aa1b678e9fe"), s.Args)
	assert.Equal(t, types.HashTypeType, s.HashType)
}
