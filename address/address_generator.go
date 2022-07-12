package address

import (
	"github.com/nervosnetwork/ckb-sdk-go/crypto/blake2b"
	"github.com/nervosnetwork/ckb-sdk-go/crypto/secp256k1"
	"github.com/nervosnetwork/ckb-sdk-go/types"
)

func GenerateScriptSecp256K1Blake160SignhashAll(key *secp256k1.Secp256k1Key) *types.Script {
	args, _ := blake2b.Blake160(key.PubKey())
	return &types.Script{
		// The same code hash is shared by mainnet and testnet
		CodeHash: types.GetCodeHash(types.BuiltinScriptSecp256k1Blake160SighashAll, types.NetworkMain),
		HashType: types.HashTypeType,
		Args:     args,
	}
}

func GenerateAddressSecp256K1Blake160SignhashAll(key *secp256k1.Secp256k1Key, network types.Network) *Address {
	script := GenerateScriptSecp256K1Blake160SignhashAll(key)
	return &Address{
		Script:  script,
		Network: network,
	}
}

