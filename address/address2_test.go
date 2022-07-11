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

var scriptWithHashTypeData = types.Script{
	CodeHash: types.HexToHash("0x709f3fda12f561cfacf92273c57a98fede188a3f1a59b1f888d113f9cce08649"),
	HashType: types.HashTypeData,
	Args:     common.FromHex("0xb73961e46d9eb118d3de1d1e8f30b3af7bbf3160"),
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
	testDecode(t, "ckb1qfcf7076zt6krnavly3883t6nrlduxy28ud9nv0c3rg387wvuzryndeev8jxm843rrfau8g73uct8tmmhuckqy57acj",
		&Address{scriptWithHashTypeData, types.NetworkMain})
	testDecode(t, "ckb1qpcf7076zt6krnavly3883t6nrlduxy28ud9nv0c3rg387wvuzryjq9h89s7gmv7kyvd8hsar68npva00wlnzcqgh76tz",
		&Address{scriptWithHashTypeData, types.NetworkMain})
}

func testDecode(t *testing.T, encoded string, address *Address) {
	a, err := Decode(encoded)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, address, a)
}

func TestEncode(t *testing.T) {
	testEncode(t, "ckb1qyqt8xaupvm8837nv3gtc9x0ekkj64vud3jqfwyw5v", singleSigScript, types.NetworkMain, Address.EncodeShort)
	testEncode(t, "ckb1qjda0cr08m85hc8jlnfp3zer7xulejywt49kt2rr0vthywaa50xw3vumhs9nvu786dj9p0q5elx66t24n3kxgj53qks", singleSigScript, types.NetworkMain, Address.EncodeFullBech32)
	testEncode(t, "ckb1qzda0cr08m85hc8jlnfp3zer7xulejywt49kt2rr0vthywaa50xwsqdnnw7qkdnnclfkg59uzn8umtfd2kwxceqxwquc4", singleSigScript, types.NetworkMain, Address.EncodeFullBech32m)

	// Multisig
	testEncode(t, "ckb1qyq5lv479ewscx3ms620sv34pgeuz6zagaaqklhtgg", multiSigScript, types.NetworkMain, Address.EncodeShort)
	testEncode(t, "ckt1qyq5lv479ewscx3ms620sv34pgeuz6zagaaqt6f5y5", multiSigScript, types.NetworkTest, Address.EncodeShort)

	// anyone can pay
	testEncode(t, "ckb1qypt6p7e7v4uudxjw9f2dgper5ey77d2hp2qxz4u4u", acpScript, types.NetworkMain, Address.EncodeShort)

	// hashType DATA
	testEncode(t, "ckb1qfcf7076zt6krnavly3883t6nrlduxy28ud9nv0c3rg387wvuzryndeev8jxm843rrfau8g73uct8tmmhuckqy57acj",
		scriptWithHashTypeData, types.NetworkMain, Address.EncodeFullBech32)
	testEncode(t, "ckb1qpcf7076zt6krnavly3883t6nrlduxy28ud9nv0c3rg387wvuzryjq9h89s7gmv7kyvd8hsar68npva00wlnzcqgh76tz",
		scriptWithHashTypeData, types.NetworkMain, Address.EncodeFullBech32m)
}

func testEncode(t *testing.T, expected string, script types.Script, network types.Network, f func(Address) (string, error)) {
	address := Address{script, network}
	encoded, err := f(address)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, expected, encoded)
}

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
