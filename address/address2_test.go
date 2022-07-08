package address

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/nervosnetwork/ckb-sdk-go/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

var singleSigScript = types.Script{
	CodeHash: types.HexToHash("0x9bd7e06f3ecf4be0f2fcd2188b23f1b9fcc88e5d4b65a8637b17723bbda3cce8"),
	HashType: types.HashTypeType,
	Args:     common.FromHex("0xb39bbc0b3673c7d36450bc14cfcdad2d559c6c64"),
}

var multiSigScript = types.Script{
	CodeHash: types.HexToHash("0x5c5069eb0857efc65e1bca0c07df34c31663b3622fd3876c876320fc9634e2a8"),
	HashType: types.HashTypeType,
	Args:     common.FromHex("0x4fb2be2e5d0c1a3b8694f832350a33c1685d477a"),
}

var acpScript = types.Script{
	CodeHash: types.HexToHash("0xd369597ff47f29fbc0d47d2e3775370d1250b85140c670e4718af712983a2354"),
	HashType: types.HashTypeType,
	Args:     common.FromHex("bd07d9f32bce34d27152a6a0391d324f79aab854"),
}

var singleSigScriptTypeData = types.Script{
	CodeHash: types.HexToHash("9bd7e06f3ecf4be0f2fcd2188b23f1b9fcc88e5d4b65a8637b17723bbda3cce8"),
	HashType: types.HashTypeData,
	Args:     common.FromHex("b39bbc0b3673c7d36450bc14cfcdad2d559c6c64"),
}

func TestDecode(t *testing.T) {
	// short format
	testDecode(t, "ckb1qyqt8xaupvm8837nv3gtc9x0ekkj64vud3jqfwyw5v", &Address{singleSigScript, types.NetworkMain})
	// long bech32 format
	testDecode(t, "ckb1qjda0cr08m85hc8jlnfp3zer7xulejywt49kt2rr0vthywaa50xw3vumhs9nvu786dj9p0q5elx66t24n3kxgj53qks",
		&Address{singleSigScript, types.NetworkMain})
	// long bech32m format
	testDecode(t, "ckb1qzda0cr08m85hc8jlnfp3zer7xulejywt49kt2rr0vthywaa50xwsqdnnw7qkdnnclfkg59uzn8umtfd2kwxceqxwquc4",
		&Address{singleSigScript, types.NetworkMain})

	// Multisig
	testDecode(t, "ckb1qyq5lv479ewscx3ms620sv34pgeuz6zagaaqklhtgg", &Address{multiSigScript, types.NetworkMain})
	testDecode(t, "ckt1qyq5lv479ewscx3ms620sv34pgeuz6zagaaqt6f5y5", &Address{multiSigScript, types.NetworkTest})

	// Any can pay
	testDecode(t, "ckb1qypt6p7e7v4uudxjw9f2dgper5ey77d2hp2qxz4u4u", &Address{acpScript, types.NetworkMain})

	// hashType DATA
	testDecode(t, "ckb1qzda0cr08m85hc8jlnfp3zer7xulejywt49kt2rr0vthywaa50xwsq9nnw7qkdnnclfkg59uzn8umtfd2kwxceqvguktl",
		&Address{singleSigScriptTypeData, types.NetworkMain})
	testDecode(t, "ckt1qzda0cr08m85hc8jlnfp3zer7xulejywt49kt2rr0vthywaa50xwsq9nnw7qkdnnclfkg59uzn8umtfd2kwxceqz6hep8",
		&Address{singleSigScriptTypeData, types.NetworkTest})
}

func testDecode(t *testing.T, encoded string, address *Address) {
	a, err := Decode(encoded)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, address, a)
}

func TestEncode(t *testing.T) {
	var (
		address Address
		encoded string
		err     error
	)
	address = Address{singleSigScript, types.NetworkMain}
	if encoded, err = address.EncodeShort(); err != nil {
		t.Error(err)
	}
	assert.Equal(t, "ckb1qyqt8xaupvm8837nv3gtc9x0ekkj64vud3jqfwyw5v", encoded)
	if encoded, err = address.EncodeFullBech32(); err != nil {
		t.Error(err)
	}
	assert.Equal(t, "ckb1qjda0cr08m85hc8jlnfp3zer7xulejywt49kt2rr0vthywaa50xw3vumhs9nvu786dj9p0q5elx66t24n3kxgj53qks", encoded)
	if encoded, err = address.EncodeFullBech32m(); err != nil {
		t.Error(err)
	}
	assert.Equal(t, "ckb1qzda0cr08m85hc8jlnfp3zer7xulejywt49kt2rr0vthywaa50xwsqdnnw7qkdnnclfkg59uzn8umtfd2kwxceqxwquc4", encoded)

	// Multisig
	address = Address{multiSigScript, types.NetworkMain}
	if encoded, err = address.EncodeShort(); err != nil {
		t.Error(err)
	}
	assert.Equal(t, "ckb1qyq5lv479ewscx3ms620sv34pgeuz6zagaaqklhtgg", encoded)
	address = Address{multiSigScript, types.NetworkTest}
	if encoded, err = address.EncodeShort(); err != nil {
		t.Error(err)
	}
	assert.Equal(t, "ckt1qyq5lv479ewscx3ms620sv34pgeuz6zagaaqt6f5y5", encoded)

	// anyone can pay
	address = Address{acpScript, types.NetworkMain}
	if encoded, err = address.EncodeShort(); err != nil {
		t.Error(err)
	}
	assert.Equal(t, "ckb1qypt6p7e7v4uudxjw9f2dgper5ey77d2hp2qxz4u4u", encoded)

	// hashType DATA
	address = Address{singleSigScriptTypeData, types.NetworkMain}
	if encoded, err = address.EncodeFullBech32m(); err != nil {
		t.Error(err)
	}
	assert.Equal(t, "ckb1qzda0cr08m85hc8jlnfp3zer7xulejywt49kt2rr0vthywaa50xwsq9nnw7qkdnnclfkg59uzn8umtfd2kwxceqvguktl", encoded)
	address = Address{singleSigScriptTypeData, types.NetworkTest}
	if encoded, err = address.EncodeFullBech32m(); err != nil {
		t.Error(err)
	}
	assert.Equal(t, "ckt1qzda0cr08m85hc8jlnfp3zer7xulejywt49kt2rr0vthywaa50xwsq9nnw7qkdnnclfkg59uzn8umtfd2kwxceqz6hep8", encoded)
}
