package address

import (
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/nervosnetwork/ckb-sdk-go/crypto/blake2b"
	"github.com/nervosnetwork/ckb-sdk-go/crypto/secp256k1"
	"github.com/nervosnetwork/ckb-sdk-go/transaction/signer"
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
	return parsedSenderAddr.Script.CodeHash == systemScripts.SecpSingleSigCell.CodeHash &&
		parsedSenderAddr.Script.HashType == systemScripts.SecpSingleSigCell.HashType &&
		len(parsedSenderAddr.Script.Args) == 20
}

// GenerateSecp256k1Blake160MultisigScript generate scep256k1 multisig script.
// It can accept public key (in compressed format, 33 bytes each) array or public key hash (20 bytes) array, and
// return error if giving none of them.
func GenerateSecp256k1Blake160MultisigScript(requireN, threshold int, publicKeysOrHashes [][]byte) (*types.Script, []byte, error) {
	multisigScript := signer.MultisigScript{
		Version:    0,
		FirstN:     byte(requireN),
		Threshold:  byte(threshold),
		KeysHashes: [][20]byte{},
	}

	isPublicKeyHash := len(publicKeysOrHashes[0]) == 20
	if isPublicKeyHash {
		for _, publicKeyHash := range publicKeysOrHashes {
			if len(publicKeyHash) != 20 {
				return nil, nil, errors.New("public key hash length must be 20 bytes")
			}
			if err := multisigScript.AddKeyHashBySlice(publicKeyHash); err != nil {
				return nil, nil, err
			}
		}
	} else {
		for _, publicKey := range publicKeysOrHashes {
			if len(publicKey) != 33 {
				return nil, nil, errors.New("public key (compressed) length must be 33 bytes")
			}
			publicKeyHash, err := blake2b.Blake160(publicKey)
			if err != nil {
				return nil, nil, err
			}
			if err := multisigScript.AddKeyHashBySlice(publicKeyHash); err != nil {
				return nil, nil, err
			}
		}
	}

	args, err := multisigScript.ComputeHash()
	if err != nil {
		return nil, nil, err
	}

	// secp256k1_blake160_multisig_all share the same code hash in network main and test
	codeHash := types.GetCodeHash(types.BuiltinScriptSecp256k1Blake160MultisigAll, types.NetworkTest)
	return &types.Script{
		CodeHash: codeHash,
		HashType: types.HashTypeType,
		Args:     args,
	}, multisigScript.Encode(), nil
}
