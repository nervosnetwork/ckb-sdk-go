package address

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/nervosnetwork/ckb-sdk-go/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateShortAddress(t *testing.T) {
	shortAddress, _ := GenerateShortAddress(Mainnet)
	fmt.Println(shortAddress)
}

func TestGenerateAcpAddress(t *testing.T) {
	address := "ckt1qzda0cr08m85hc8jlnfp3zer7xulejywt49kt2rr0vthywaa50xwsqg958atl2zdh8jn3ch8lc72nt0cf864ecqdxm9zf"
	acpAddress, err := GenerateAcpAddress(address)
	assert.Nil(t, err)
	assert.Equal(t, "ckt1qq6pngwqn6e9vlm92th84rk0l4jp2h8lurchjmnwv8kq3rt5psf4vqg958atl2zdh8jn3ch8lc72nt0cf864ecqz4aphl", acpAddress)
}

func TestGenerateChequeAddress(t *testing.T) {
	senderAddress := "ckt1qzda0cr08m85hc8jlnfp3zer7xulejywt49kt2rr0vthywaa50xwsqd0pdquvfuq077aemn447shf4d8u5f4a0glzz2g4"
	receiverAddress := "ckt1qzda0cr08m85hc8jlnfp3zer7xulejywt49kt2rr0vthywaa50xwsqg958atl2zdh8jn3ch8lc72nt0cf864ecqdxm9zf"
	acpAddress, err := GenerateChequeAddress(senderAddress, receiverAddress)
	assert.Nil(t, err)
	assert.Equal(t, "ckt1qpsdtuu7lnjqn3v8ew02xkwwlh4dv5x2z28shkwt8p2nfruccux4kq2je6sm0zczgrepc8y547zvuu6zpshfvvjh7h2ln2w035d2lnh32ylk5ydmjq5ypwq24ftzt", acpAddress)
}

func TestBech32mTypeFullMainnetAddressGenerate(t *testing.T) {
	script := &types.Script{
		CodeHash: types.HexToHash("0x9bd7e06f3ecf4be0f2fcd2188b23f1b9fcc88e5d4b65a8637b17723bbda3cce8"),
		HashType: types.HashTypeType,
		Args:     common.FromHex("0xb39bbc0b3673c7d36450bc14cfcdad2d559c6c64"),
	}

	address, err := ConvertScriptToBech32mFullAddress(Mainnet, script)
	println(address)
	if err != nil {
		return
	}

	assert.Equal(t,
		"ckb1qzda0cr08m85hc8jlnfp3zer7xulejywt49kt2rr0vthywaa50xwsqdnnw7qkdnnclfkg59uzn8umtfd2kwxceqxwquc4",
		address)

	t.Run("parse full address", func(t *testing.T) {
		mnAddress, err := Parse(address)

		assert.Nil(t, err)
		assert.Equal(t, Mainnet, mnAddress.Mode)
		assert.Equal(t, FullBech32, mnAddress.Type)
		assert.Equal(t, script.CodeHash, mnAddress.Script.CodeHash)
		assert.Equal(t, script.HashType, mnAddress.Script.HashType)
		assert.Equal(t, script.Args, mnAddress.Script.Args)
	})

}

func TestBech32mDataFullMainnetAddressGenerate(t *testing.T) {
	script := &types.Script{
		CodeHash: types.HexToHash("0xa656f172b6b45c245307aeb5a7a37a176f002f6f22e92582c58bf7ba362e4176"),
		HashType: types.HashTypeData,
		Args:     common.FromHex("0x36c329ed630d6ce750712a477543672adab57f4c"),
	}

	address, err := ConvertScriptToBech32mFullAddress(Mainnet, script)
	println(address)
	if err != nil {
		return
	}

	assert.Equal(t,
		"ckb1qzn9dutjk669cfznq7httfar0gtk7qp0du3wjfvzck9l0w3k9eqhvqpkcv576ccddnn4quf2ga65xee2m26h7nqdcg257",
		address)

	t.Run("parse full address", func(t *testing.T) {
		mnAddress, err := Parse(address)

		assert.Nil(t, err)
		assert.Equal(t, Mainnet, mnAddress.Mode)
		assert.Equal(t, FullBech32, mnAddress.Type)
		assert.Equal(t, script.CodeHash, mnAddress.Script.CodeHash)
		assert.Equal(t, script.HashType, mnAddress.Script.HashType)
		assert.Equal(t, script.Args, mnAddress.Script.Args)
	})
}

func TestBech32mData1FullMainnetAddressGenerate(t *testing.T) {
	script := &types.Script{
		CodeHash: types.HexToHash("0xa656f172b6b45c245307aeb5a7a37a176f002f6f22e92582c58bf7ba362e4176"),
		HashType: types.HashTypeData1,
		Args:     common.FromHex("0x36c329ed630d6ce750712a477543672adab57f4c"),
	}

	address, err := ConvertScriptToBech32mFullAddress(Mainnet, script)
	println(address)
	if err != nil {
		return
	}

	assert.Equal(t,
		"ckb1qzn9dutjk669cfznq7httfar0gtk7qp0du3wjfvzck9l0w3k9eqhvq3kcv576ccddnn4quf2ga65xee2m26h7nqe5e7m2",
		address)

	t.Run("parse full address", func(t *testing.T) {
		mnAddress, err := Parse(address)

		assert.Nil(t, err)
		assert.Equal(t, Mainnet, mnAddress.Mode)
		assert.Equal(t, FullBech32, mnAddress.Type)
		assert.Equal(t, script.CodeHash, mnAddress.Script.CodeHash)
		assert.Equal(t, script.HashType, mnAddress.Script.HashType)
		assert.Equal(t, script.Args, mnAddress.Script.Args)
	})

}
