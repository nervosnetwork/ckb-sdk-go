package address

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/nervosnetwork/ckb-sdk-go/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

// test fixture comes from: https://github.com/rev-chaos/ckb-address-demo
func TestShort(t *testing.T) {
	// sighash
	s := generateScript(
		"0x9bd7e06f3ecf4be0f2fcd2188b23f1b9fcc88e5d4b65a8637b17723bbda3cce8",
		"b39bbc0b3673c7d36450bc14cfcdad2d559c6c64",
		types.HashTypeType)
	testShort(t, s, types.NetworkMain, "ckb1qyqt8xaupvm8837nv3gtc9x0ekkj64vud3jqfwyw5v")
	testShort(t, s, types.NetworkTest, "ckt1qyqt8xaupvm8837nv3gtc9x0ekkj64vud3jq5t63cs")

	// multisig
	s = generateScript(
		"0x5c5069eb0857efc65e1bca0c07df34c31663b3622fd3876c876320fc9634e2a8",
		"4fb2be2e5d0c1a3b8694f832350a33c1685d477a",
		types.HashTypeType)
	testShort(t, s, types.NetworkMain, "ckb1qyq5lv479ewscx3ms620sv34pgeuz6zagaaqklhtgg")
	testShort(t, s, types.NetworkTest, "ckt1qyq5lv479ewscx3ms620sv34pgeuz6zagaaqt6f5y5")

	// acp mainnet
	s = generateScript(
		"0xd369597ff47f29fbc0d47d2e3775370d1250b85140c670e4718af712983a2354",
		"bd07d9f32bce34d27152a6a0391d324f79aab854",
		types.HashTypeType)
	testShort(t, s, types.NetworkMain, "ckb1qypt6p7e7v4uudxjw9f2dgper5ey77d2hp2qxz4u4u")
	// acp testnet
	s = generateScript(
		"0x3419a1c09eb2567f6552ee7a8ecffd64155cffe0f1796e6e61ec088d740c1356",
		"bd07d9f32bce34d27152a6a0391d324f79aab854",
		types.HashTypeType)
	testShort(t, s, types.NetworkTest, "ckt1qypt6p7e7v4uudxjw9f2dgper5ey77d2hp2qm8treq")
}

func testShort(t *testing.T, script *types.Script, network types.Network, encoded string) {
	a := &Address{
		Script:  script,
		Network: network,
	}
	result, err := a.EncodeShort()
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, encoded, result)
	addr, err := Decode(encoded)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, a, addr)
}

func TestFullBech32(t *testing.T) {
	s := generateScript(
		"9bd7e06f3ecf4be0f2fcd2188b23f1b9fcc88e5d4b65a8637b17723bbda3cce8",
		"b39bbc0b3673c7d36450bc14cfcdad2d559c6c64",
		types.HashTypeData)
	testFullBech32(t, s, types.NetworkMain, "ckb1q2da0cr08m85hc8jlnfp3zer7xulejywt49kt2rr0vthywaa50xw3vumhs9nvu786dj9p0q5elx66t24n3kxgdwd2q8")
	testFullBech32(t, s, types.NetworkTest, "ckt1q2da0cr08m85hc8jlnfp3zer7xulejywt49kt2rr0vthywaa50xw3vumhs9nvu786dj9p0q5elx66t24n3kxgqd588c")

	s.HashType = types.HashTypeType
	testFullBech32(t, s, types.NetworkMain, "ckb1qjda0cr08m85hc8jlnfp3zer7xulejywt49kt2rr0vthywaa50xw3vumhs9nvu786dj9p0q5elx66t24n3kxgj53qks")
	testFullBech32(t, s, types.NetworkTest, "ckt1qjda0cr08m85hc8jlnfp3zer7xulejywt49kt2rr0vthywaa50xw3vumhs9nvu786dj9p0q5elx66t24n3kxglhgd30")
}

func testFullBech32(t *testing.T, script *types.Script, network types.Network, encoded string) {
	a := &Address{
		Script:  script,
		Network: network,
	}
	result, err := a.EncodeFullBech32()
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, encoded, result)
	addr, err := Decode(encoded)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, a, addr)
}

func TestFullBech32m(t *testing.T) {
	s := generateScript(
		"9bd7e06f3ecf4be0f2fcd2188b23f1b9fcc88e5d4b65a8637b17723bbda3cce8",
		"b39bbc0b3673c7d36450bc14cfcdad2d559c6c64",
		types.HashTypeData)
	testFullBech32m(t, s, types.NetworkMain, "ckb1qzda0cr08m85hc8jlnfp3zer7xulejywt49kt2rr0vthywaa50xwsq9nnw7qkdnnclfkg59uzn8umtfd2kwxceqvguktl")
	testFullBech32m(t, s, types.NetworkTest, "ckt1qzda0cr08m85hc8jlnfp3zer7xulejywt49kt2rr0vthywaa50xwsq9nnw7qkdnnclfkg59uzn8umtfd2kwxceqz6hep8")

	s.HashType = types.HashTypeType
	testFullBech32m(t, s, types.NetworkMain, "ckb1qzda0cr08m85hc8jlnfp3zer7xulejywt49kt2rr0vthywaa50xwsqdnnw7qkdnnclfkg59uzn8umtfd2kwxceqxwquc4")
	testFullBech32m(t, s, types.NetworkTest, "ckt1qzda0cr08m85hc8jlnfp3zer7xulejywt49kt2rr0vthywaa50xwsqdnnw7qkdnnclfkg59uzn8umtfd2kwxceqgutnjd")

	s.HashType = types.HashTypeData1
	testFullBech32m(t, s, types.NetworkMain, "ckb1qzda0cr08m85hc8jlnfp3zer7xulejywt49kt2rr0vthywaa50xwsq4nnw7qkdnnclfkg59uzn8umtfd2kwxceqcydzyt")
	testFullBech32m(t, s, types.NetworkTest, "ckt1qzda0cr08m85hc8jlnfp3zer7xulejywt49kt2rr0vthywaa50xwsq4nnw7qkdnnclfkg59uzn8umtfd2kwxceqkkxdwn")
}

func testFullBech32m(t *testing.T, script *types.Script, network types.Network, encoded string) {
	a := &Address{
		Script:  script,
		Network: network,
	}
	result, err := a.EncodeFullBech32m()
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, encoded, result)
	addr, err := Decode(encoded)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, a, addr)
}

func generateScript(codeHash string, args string, hashType types.ScriptHashType) *types.Script {
	return &types.Script{
		CodeHash: types.HexToHash(codeHash),
		HashType: hashType,
		Args:     common.FromHex(args),
	}
}

// These invalid addresses come form https://github.com/nervosnetwork/ckb-sdk-rust/pull/7/files
func TestInvalidDecode(t *testing.T) {
	var err error
	_, err = Decode("ckb1qyqylv479ewscx3ms620sv34pgeuz6zagaaqh0knz7")
	assert.NotNil(t, err)
	_, err = Decode("ckb1qyqylv479ewscx3ms620sv34pgeuz6zagaarxdzvx03")
	assert.NotNil(t, err)
	_, err = Decode("ckb1qyg5lv479ewscx3ms620sv34pgeuz6zagaaqajch0c")
	assert.NotNil(t, err)
	_, err = Decode("ckb1q2da0cr08m85hc8jlnfp3zer7xulejywt49kt2rr0vthywaa50xwsnajhch96rq68wrqn2tmhm")
	assert.NotNil(t, err)
	_, err = Decode("ckb1qzda0cr08m85hc8jlnfp3zer7xulejywt49kt2rr0vthywaa50xwsq20k2lzuhgvrgacv4tmr88")
	assert.NotNil(t, err)
	_, err = Decode("ckb1qzda0cr08m85hc8jlnfp3zer7xulejywt49kt2rr0vthywaa50xwsqz0k2lzuhgvrgacvhcym08")
	assert.NotNil(t, err)
	_, err = Decode("ckb1qzda0cr08m85hc8jlnfp3zer7xulejywt49kt2rr0vthywaa50xwsqj0k2lzuhgvrgacvnhnzl8")
	assert.NotNil(t, err)
}