package address

import (
	"context"
	"github.com/nervosnetwork/ckb-sdk-go/mocks"
	"github.com/nervosnetwork/ckb-sdk-go/utils"
	"github.com/pkg/errors"
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"

	"github.com/nervosnetwork/ckb-sdk-go/types"
)

func TestGenerate(t *testing.T) {
	script := &types.Script{
		CodeHash: types.HexToHash("0x9bd7e06f3ecf4be0f2fcd2188b23f1b9fcc88e5d4b65a8637b17723bbda3cce8"),
		HashType: types.HashTypeType,
		Args:     common.FromHex("0xedcda9513fa030ce4308e29245a22c022d0443bb"),
	}

	mnAddress, err := ConvertScriptToShortAddress(Mainnet, script)

	assert.Nil(t, err)
	assert.Equal(t, "ckb1qyqwmndf2yl6qvxwgvyw9yj95gkqytgygwasshh9m8", mnAddress)

	tnAddress, err := ConvertScriptToShortAddress(Testnet, script)
	assert.Nil(t, err)
	assert.Equal(t, "ckt1qyqwmndf2yl6qvxwgvyw9yj95gkqytgygwasdjf6hm", tnAddress)

	t.Run("generate short payload acp address without minimum limit", func(t *testing.T) {
		mAcpLock := &types.Script{
			CodeHash: types.HexToHash(utils.AnyoneCanPayCodeHashOnLina),
			HashType: types.HashTypeType,
			Args:     common.FromHex("0x4fb2be2e5d0c1a3b8694f832350a33c1685d477a"),
		}

		mAddress, err := ConvertScriptToShortAddress(Mainnet, mAcpLock)
		assert.Nil(t, err)
		assert.Equal(t, "ckb1qypylv479ewscx3ms620sv34pgeuz6zagaaqvrugu7", mAddress)

		tAcpLock := &types.Script{
			CodeHash: types.HexToHash(utils.AnyoneCanPayCodeHashOnAggron),
			HashType: types.HashTypeType,
			Args:     common.FromHex("0x4fb2be2e5d0c1a3b8694f832350a33c1685d477a"),
		}
		tAddress, err := ConvertScriptToShortAddress(Testnet, tAcpLock)
		assert.Nil(t, err)
		assert.Equal(t, "ckt1qypylv479ewscx3ms620sv34pgeuz6zagaaq3xzhsz", tAddress)
	})

	t.Run("generate short payload acp address with ckb minimum limit", func(t *testing.T) {
		mAcpLock := &types.Script{
			CodeHash: types.HexToHash(utils.AnyoneCanPayCodeHashOnLina),
			HashType: types.HashTypeType,
			Args:     common.FromHex("0x4fb2be2e5d0c1a3b8694f832350a33c1685d477a0c"),
		}

		mAddress, err := ConvertScriptToShortAddress(Mainnet, mAcpLock)
		assert.Nil(t, err)
		assert.Equal(t, "ckb1qypylv479ewscx3ms620sv34pgeuz6zagaaqcehzz9g", mAddress)

		tAcpLock := &types.Script{
			CodeHash: types.HexToHash(utils.AnyoneCanPayCodeHashOnAggron),
			HashType: types.HashTypeType,
			Args:     common.FromHex("0x4fb2be2e5d0c1a3b8694f832350a33c1685d477a0c"),
		}
		tAddress, err := ConvertScriptToShortAddress(Testnet, tAcpLock)
		assert.Nil(t, err)
		assert.Equal(t, "ckt1qypylv479ewscx3ms620sv34pgeuz6zagaaqc9q8fqw", tAddress)
	})

	t.Run("generate short payload acp address with ckb and udt minimum limit", func(t *testing.T) {
		mAcpLock := &types.Script{
			CodeHash: types.HexToHash(utils.AnyoneCanPayCodeHashOnLina),
			HashType: types.HashTypeType,
			Args:     common.FromHex("0x4fb2be2e5d0c1a3b8694f832350a33c1685d477a0c01"),
		}

		mAddress, err := ConvertScriptToShortAddress(Mainnet, mAcpLock)
		assert.Nil(t, err)
		assert.Equal(t, "ckb1qypylv479ewscx3ms620sv34pgeuz6zagaaqcqgzc5xlw", mAddress)

		tAcpLock := &types.Script{
			CodeHash: types.HexToHash(utils.AnyoneCanPayCodeHashOnAggron),
			HashType: types.HashTypeType,
			Args:     common.FromHex("0x4fb2be2e5d0c1a3b8694f832350a33c1685d477a0c01"),
		}
		tAddress, err := ConvertScriptToShortAddress(Testnet, tAcpLock)
		assert.Nil(t, err)
		assert.Equal(t, "ckt1qypylv479ewscx3ms620sv34pgeuz6zagaaqcqgr072sz", tAddress)
	})

	t.Run("generate full payload address when acp lock args is more than 22 bytes", func(t *testing.T) {
		mAcpLock := &types.Script{
			CodeHash: types.HexToHash(utils.AnyoneCanPayCodeHashOnLina),
			HashType: types.HashTypeType,
			Args:     common.FromHex("0x4fb2be2e5d0c1a3b8694f832350a33c1685d477a0c0101"),
		}

		mAddress, err := ConvertScriptToFullAddress(FullTypeFormat, Mainnet, mAcpLock)
		assert.Nil(t, err)
		assert.Equal(t, "ckb1qnfkjktl73ljn77q637judm4xux3y59c29qvvu8ywx90wy5c8g34gnajhch96rq68wrff7pjx59r8stgt4rh5rqpqy532xj3", mAddress)

		tAcpLock := &types.Script{
			CodeHash: types.HexToHash(utils.AnyoneCanPayCodeHashOnAggron),
			HashType: types.HashTypeType,
			Args:     common.FromHex("0x4fb2be2e5d0c1a3b8694f832350a33c1685d477a0c0101"),
		}

		tAddress, err := ConvertScriptToFullAddress(FullTypeFormat, Testnet, tAcpLock)
		assert.Nil(t, err)
		assert.Equal(t, "ckt1qs6pngwqn6e9vlm92th84rk0l4jp2h8lurchjmnwv8kq3rt5psf4vnajhch96rq68wrff7pjx59r8stgt4rh5rqpqy2a9ak4", tAddress)
	})
}

func TestParse(t *testing.T) {
	script := &types.Script{
		CodeHash: types.HexToHash("0x9bd7e06f3ecf4be0f2fcd2188b23f1b9fcc88e5d4b65a8637b17723bbda3cce8"),
		HashType: types.HashTypeType,
		Args:     common.FromHex("0xedcda9513fa030ce4308e29245a22c022d0443bb"),
	}

	mnAddress, err := Parse("ckb1qyqwmndf2yl6qvxwgvyw9yj95gkqytgygwasshh9m8")

	assert.Nil(t, err)
	assert.Equal(t, Mainnet, mnAddress.Mode)
	assert.Equal(t, TypeShort, mnAddress.Type)
	assert.Equal(t, script.CodeHash, mnAddress.Script.CodeHash)
	assert.Equal(t, script.HashType, mnAddress.Script.HashType)
	assert.Equal(t, script.Args, mnAddress.Script.Args)

	t.Run("parse short payload acp address without minimum limit", func(t *testing.T) {
		mAcpLock := &types.Script{
			CodeHash: types.HexToHash(utils.AnyoneCanPayCodeHashOnLina),
			HashType: types.HashTypeType,
			Args:     common.FromHex("0x4fb2be2e5d0c1a3b8694f832350a33c1685d477a"),
		}

		mParsedAddress, err := Parse("ckb1qypylv479ewscx3ms620sv34pgeuz6zagaaqvrugu7")
		assert.Nil(t, err)
		assert.Equal(t, Mainnet, mParsedAddress.Mode)
		assert.Equal(t, TypeShort, mParsedAddress.Type)
		assert.Equal(t, mAcpLock.CodeHash, mParsedAddress.Script.CodeHash)
		assert.Equal(t, mAcpLock.HashType, mParsedAddress.Script.HashType)
		assert.Equal(t, mAcpLock.Args, mParsedAddress.Script.Args)

		tAcpLock := &types.Script{
			CodeHash: types.HexToHash(utils.AnyoneCanPayCodeHashOnAggron),
			HashType: types.HashTypeType,
			Args:     common.FromHex("0x4fb2be2e5d0c1a3b8694f832350a33c1685d477a"),
		}

		tParsedAddress, err := Parse("ckt1qypylv479ewscx3ms620sv34pgeuz6zagaaq3xzhsz")
		assert.Nil(t, err)
		assert.Equal(t, Testnet, tParsedAddress.Mode)
		assert.Equal(t, TypeShort, tParsedAddress.Type)
		assert.Equal(t, tAcpLock.CodeHash, tParsedAddress.Script.CodeHash)
		assert.Equal(t, tAcpLock.HashType, tParsedAddress.Script.HashType)
		assert.Equal(t, tAcpLock.Args, tParsedAddress.Script.Args)
	})

	t.Run("parse short payload acp address with ckb minimum limit", func(t *testing.T) {
		mAcpLock := &types.Script{
			CodeHash: types.HexToHash(utils.AnyoneCanPayCodeHashOnLina),
			HashType: types.HashTypeType,
			Args:     common.FromHex("0x4fb2be2e5d0c1a3b8694f832350a33c1685d477a0c"),
		}

		mParsedAddress, err := Parse("ckb1qypylv479ewscx3ms620sv34pgeuz6zagaaqcehzz9g")
		assert.Nil(t, err)
		assert.Equal(t, Mainnet, mParsedAddress.Mode)
		assert.Equal(t, TypeShort, mParsedAddress.Type)
		assert.Equal(t, mAcpLock.CodeHash, mParsedAddress.Script.CodeHash)
		assert.Equal(t, mAcpLock.HashType, mParsedAddress.Script.HashType)
		assert.Equal(t, mAcpLock.Args, mParsedAddress.Script.Args)

		tAcpLock := &types.Script{
			CodeHash: types.HexToHash(utils.AnyoneCanPayCodeHashOnAggron),
			HashType: types.HashTypeType,
			Args:     common.FromHex("0x4fb2be2e5d0c1a3b8694f832350a33c1685d477a0c"),
		}

		tParsedAddress, err := Parse("ckt1qypylv479ewscx3ms620sv34pgeuz6zagaaqc9q8fqw")
		assert.Nil(t, err)
		assert.Equal(t, Testnet, tParsedAddress.Mode)
		assert.Equal(t, TypeShort, tParsedAddress.Type)
		assert.Equal(t, tAcpLock.CodeHash, tParsedAddress.Script.CodeHash)
		assert.Equal(t, tAcpLock.HashType, tParsedAddress.Script.HashType)
		assert.Equal(t, tAcpLock.Args, tParsedAddress.Script.Args)
	})

	t.Run("parse short payload acp address with ckb minimum limit and udt minimum limit", func(t *testing.T) {
		mAcpLock := &types.Script{
			CodeHash: types.HexToHash(utils.AnyoneCanPayCodeHashOnLina),
			HashType: types.HashTypeType,
			Args:     common.FromHex("0x4fb2be2e5d0c1a3b8694f832350a33c1685d477a0c01"),
		}

		mParsedAddress, err := Parse("ckb1qypylv479ewscx3ms620sv34pgeuz6zagaaqcqgzc5xlw")
		assert.Nil(t, err)
		assert.Equal(t, Mainnet, mParsedAddress.Mode)
		assert.Equal(t, TypeShort, mParsedAddress.Type)
		assert.Equal(t, mAcpLock.CodeHash, mParsedAddress.Script.CodeHash)
		assert.Equal(t, mAcpLock.HashType, mParsedAddress.Script.HashType)
		assert.Equal(t, mAcpLock.Args, mParsedAddress.Script.Args)

		tAcpLock := &types.Script{
			CodeHash: types.HexToHash(utils.AnyoneCanPayCodeHashOnAggron),
			HashType: types.HashTypeType,
			Args:     common.FromHex("0x4fb2be2e5d0c1a3b8694f832350a33c1685d477a0c01"),
		}

		tParsedAddress, err := Parse("ckt1qypylv479ewscx3ms620sv34pgeuz6zagaaqcqgr072sz")
		assert.Nil(t, err)
		assert.Equal(t, Testnet, tParsedAddress.Mode)
		assert.Equal(t, TypeShort, tParsedAddress.Type)
		assert.Equal(t, tAcpLock.CodeHash, tParsedAddress.Script.CodeHash)
		assert.Equal(t, tAcpLock.HashType, tParsedAddress.Script.HashType)
		assert.Equal(t, tAcpLock.Args, tParsedAddress.Script.Args)
	})

	t.Run("parse full payload acp address with args more than 22 bytes", func(t *testing.T) {
		mAcpLock := &types.Script{
			CodeHash: types.HexToHash(utils.AnyoneCanPayCodeHashOnLina),
			HashType: types.HashTypeType,
			Args:     common.FromHex("0x4fb2be2e5d0c1a3b8694f832350a33c1685d477a0c0101"),
		}

		mParsedAddress, err := Parse("ckb1qnfkjktl73ljn77q637judm4xux3y59c29qvvu8ywx90wy5c8g34gnajhch96rq68wrff7pjx59r8stgt4rh5rqpqy532xj3")
		assert.Nil(t, err)
		assert.Equal(t, Mainnet, mParsedAddress.Mode)
		assert.Equal(t, TypeFull, mParsedAddress.Type)
		assert.Equal(t, mAcpLock.CodeHash, mParsedAddress.Script.CodeHash)
		assert.Equal(t, mAcpLock.HashType, mParsedAddress.Script.HashType)
		assert.Equal(t, mAcpLock.Args, mParsedAddress.Script.Args)

		tAcpLock := &types.Script{
			CodeHash: types.HexToHash(utils.AnyoneCanPayCodeHashOnAggron),
			HashType: types.HashTypeType,
			Args:     common.FromHex("0x4fb2be2e5d0c1a3b8694f832350a33c1685d477a0c0101"),
		}

		tParsedAddress, err := Parse("ckt1qs6pngwqn6e9vlm92th84rk0l4jp2h8lurchjmnwv8kq3rt5psf4vnajhch96rq68wrff7pjx59r8stgt4rh5rqpqy2a9ak4")
		assert.Nil(t, err)
		assert.Equal(t, Testnet, tParsedAddress.Mode)
		assert.Equal(t, TypeFull, tParsedAddress.Type)
		assert.Equal(t, tAcpLock.CodeHash, tParsedAddress.Script.CodeHash)
		assert.Equal(t, tAcpLock.HashType, tParsedAddress.Script.HashType)
		assert.Equal(t, tAcpLock.Args, tParsedAddress.Script.Args)
	})
}

func TestValidateChequeAddress(t *testing.T) {
	type args struct {
		addr string
	}
	tests := []struct {
		name    string
		args    args
		want    *ParsedAddress
		wantErr error
		chain   string
	}{
		{
			"invalid address for testnet",
			args{
				addr: "ckt1q3085d480e5wanqp8hazle4z8uakcdztqsq9szrfftnd630w5n8ath08sqwqw00mx3jv0v0stwqxhv4mhp8fjcjtac0",
			},
			nil,
			errors.Errorf("address %s is not an SECP256K1 short format address", "ckt1q3085d480e5wanqp8hazle4z8uakcdztqsq9szrfftnd630w5n8ath08sqwqw00mx3jv0v0stwqxhv4mhp8fjcjtac0"),
			"ckt",
		},
		{
			"invalid address for miannet",
			args{
				addr: "ckb1q3085d480e5wanqp8hazle4z8uakcdztqsq9szrfftnd630w5n8ath08sqwqw00mx3jv0v0stwqxhv4mhp8fj43jsls",
			},
			nil,
			errors.Errorf("address %s is not an SECP256K1 short format address", "ckb1q3085d480e5wanqp8hazle4z8uakcdztqsq9szrfftnd630w5n8ath08sqwqw00mx3jv0v0stwqxhv4mhp8fj43jsls"),
			"ckb",
		},
		{
			"valid address for testnet",
			args{
				addr: "ckt1qyqdmeuqrsrnm7e5vnrmruzmsp4m9wacf6vsmcwugu",
			},
			&ParsedAddress{
				Mode: Testnet,
				Type: TypeShort,
				Script: &types.Script{
					CodeHash: types.HexToHash("0x9bd7e06f3ecf4be0f2fcd2188b23f1b9fcc88e5d4b65a8637b17723bbda3cce8"),
					HashType: types.HashTypeType,
					Args:     common.FromHex("0xdde7801c073dfb3464c7b1f05b806bb2bbb84e99"),
				},
			},
			nil,
			"ckt",
		},
		{
			"valid address for miannet",
			args{
				addr: "ckb1qyqdmeuqrsrnm7e5vnrmruzmsp4m9wacf6vsxasryq",
			},
			&ParsedAddress{
				Mode: Mainnet,
				Type: TypeShort,
				Script: &types.Script{
					CodeHash: types.HexToHash("0x9bd7e06f3ecf4be0f2fcd2188b23f1b9fcc88e5d4b65a8637b17723bbda3cce8"),
					HashType: types.HashTypeType,
					Args:     common.FromHex("0xdde7801c073dfb3464c7b1f05b806bb2bbb84e99"),
				},
			},
			nil,
			"ckb",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockClient := &mocks.Client{}
			mockClient.On("GetBlockchainInfo", context.Background()).Return(&types.BlockchainInfo{Chain: tt.chain}, nil)
			systemScripts, _ := utils.NewSystemScripts(mockClient)
			got, err := ValidateChequeAddress(tt.args.addr, systemScripts)
			if (err != nil) && err.Error() != tt.wantErr.Error() {
				t.Errorf("ValidateAddress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ValidateAddress() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertShortAddressToBech32mFullAddress(t *testing.T) {
	shortAddress := "ckt1qyqxgp7za7dajm5wzjkye52asc8fxvvqy9eqlhp82g"
	bech32mFullAddress, err := ConvertToBech32mFullAddress(shortAddress)
	if err != nil {
		t.Errorf("Fail to convert deprecated address to bech32m full address. error = %v", err)
		return
	}
	assert.Equal(t, "ckt1qzda0cr08m85hc8jlnfp3zer7xulejywt49kt2rr0vthywaa50xwsqtyqlpwlx7ed68pftzv69wcvr5nxxqzzus2zxwa6", bech32mFullAddress)
}

func TestBech32mFullAddressToShortAddress(t *testing.T) {
	bech32mFullAddress := "ckt1qzda0cr08m85hc8jlnfp3zer7xulejywt49kt2rr0vthywaa50xwsqtyqlpwlx7ed68pftzv69wcvr5nxxqzzus2zxwa6"
	shortAddress, err := ConvertToShortAddress(bech32mFullAddress)
	if err != nil {
		t.Errorf("Fail to convert deprecated address to bech32m full address. error = %v", err)
		return
	}
	assert.Equal(t, "ckt1qyqxgp7za7dajm5wzjkye52asc8fxvvqy9eqlhp82g", shortAddress)
}

func TestConvertPublickeyToBech32mFullAddress(t *testing.T) {
	address, err := ConvertPublicToAddress(Mainnet, "0xb39bbc0b3673c7d36450bc14cfcdad2d559c6c64")
	if err != nil {
		t.Errorf("Fail to convert public key to bech32m full address. error = %v", err)
		return
	}
	assert.Equal(t, address, "ckb1qzda0cr08m85hc8jlnfp3zer7xulejywt49kt2rr0vthywaa50xwsqdnnw7qkdnnclfkg59uzn8umtfd2kwxceqxwquc4")
}
