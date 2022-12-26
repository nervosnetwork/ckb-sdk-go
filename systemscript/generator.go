package systemscript

import (
	"fmt"
	"github.com/nervosnetwork/ckb-sdk-go/v2/crypto/blake2b"
	"github.com/nervosnetwork/ckb-sdk-go/v2/crypto/secp256k1"
	"github.com/nervosnetwork/ckb-sdk-go/v2/types"
)

func Secp256K1Blake160SignhashAll(key *secp256k1.Secp256k1Key) *types.Script {
	args := blake2b.Blake160(key.PubKey())
	return &types.Script{
		// The same code hash is shared by mainnet and testnet
		CodeHash: GetCodeHash(types.NetworkMain, 0),
		HashType: types.HashTypeType,
		Args:     args,
	}
}

// Secp256K1Blake160SignhashAllByPublicKey generates scep256k1_blake160_sighash_all script with 33-byte compressed public key
func Secp256K1Blake160SignhashAllByPublicKey(compressedPubKey []byte) (*types.Script, error) {
	if len(compressedPubKey) != 33 {
		return nil, fmt.Errorf("only allow 33-byte compressed public key, but accept %d bytes", len(compressedPubKey))
	}
	args := blake2b.Blake160(compressedPubKey)
	return &types.Script{
		// The same code hash is shared by mainnet and testnet
		CodeHash: GetCodeHash(types.NetworkMain, Secp256k1Blake160SighashAll),
		HashType: types.HashTypeType,
		Args:     args,
	}, nil
}

// Secp256k1Blake160Multisig generates scep256k1_blake160_multisig script.
func Secp256k1Blake160Multisig(config *MultisigConfig) (*types.Script, error) {
	args := config.Hash160()
	// secp256k1_blake160_multisig_all share the same code hash in network main and test
	codeHash := GetCodeHash(types.NetworkTest, Secp256k1Blake160MultisigAll)
	return &types.Script{
		CodeHash: codeHash,
		HashType: types.HashTypeType,
		Args:     args,
	}, nil
}
