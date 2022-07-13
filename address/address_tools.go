package address

import (
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/nervosnetwork/ckb-sdk-go/crypto/blake2b"
	"github.com/nervosnetwork/ckb-sdk-go/crypto/secp256k1"
	"github.com/nervosnetwork/ckb-sdk-go/types"
	"github.com/nervosnetwork/ckb-sdk-go/utils"
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

func GenerateScriptSecp256K1Blake160SignhashAllByPublicKey(pubKey string) (*types.Script, error) {
	b := common.FromHex(pubKey)
	if len(b) != 33 {
		return nil, errors.New("only accept 33-byte compressed public key")
	}
	args, _ := blake2b.Blake160(b)
	return &types.Script{
		// The same code hash is shared by mainnet and testnet
		CodeHash: types.GetCodeHash(types.BuiltinScriptSecp256k1Blake160SighashAll, types.NetworkMain),
		HashType: types.HashTypeType,
		Args:     args,
	}, nil
}

func GenerateAddressSecp256K1Blake160SignhashAll(key *secp256k1.Secp256k1Key, network types.Network) *Address {
	script := GenerateScriptSecp256K1Blake160SignhashAll(key)
	return &Address{
		Script:  script,
		Network: network,
	}
}

func ValidateChequeAddress(addr string, systemScripts *utils.SystemScripts) (*Address, error) {
	address, err := Decode(addr)
	if err != nil {
		return nil, err
	}
	if isSecp256k1Lock(address, systemScripts) {
		return address, nil
	}
	return nil, errors.New(fmt.Sprintf("address %s is not an SECP256K1 short format address", addr))
}

func isSecp256k1Lock(parsedSenderAddr *Address, systemScripts *utils.SystemScripts) bool {
	return parsedSenderAddr.Script.CodeHash == systemScripts.SecpSingleSigCell.CellHash &&
		parsedSenderAddr.Script.HashType == systemScripts.SecpSingleSigCell.HashType &&
		len(parsedSenderAddr.Script.Args) == 20
}