package address

import (
	"encoding/binary"
	"errors"

	"github.com/nervosnetwork/ckb-sdk-go/crypto/blake2b"
	"github.com/nervosnetwork/ckb-sdk-go/transaction"
	"github.com/nervosnetwork/ckb-sdk-go/types"
)

// GenerateSecp256k1MultisigScript generate scep256k1 multisig script.
// It can accept public key (in compressed format, 33 bytes each) array or public key hash (20 bytes) array, and
// return error if giving none of them.
func GenerateSecp256k1MultisigScript(requireN, threshold int, publicKeysOrHashes [][]byte) (*types.Script, []byte, error) {
	if requireN < 0 || requireN > 255 {
		return nil, nil, errors.New("requireN must ranging from 0 to 255")
	}
	if threshold < 0 || threshold > 255 {
		return nil, nil, errors.New("requireN must ranging from 0 to 255")
	}
	if len(publicKeysOrHashes) > 255 {
		return nil, nil, errors.New("public keys size must be less than 256")
	}
	if len(publicKeysOrHashes) < requireN || len(publicKeysOrHashes) < threshold {
		return nil, nil, errors.New("public keys error")
	}

	isPublicKeyHash := len(publicKeysOrHashes[0]) == 20
	var publicKeysHash [][]byte
	if isPublicKeyHash {
		for _, publicKeyHash := range publicKeysOrHashes {
			if len(publicKeyHash) != 20 {
				return nil, nil, errors.New("public key hash length must be 20 bytes")
			}
			publicKeysHash = append(publicKeysHash, publicKeyHash)
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
			publicKeysHash = append(publicKeysHash, publicKeyHash)
		}
	}

	var data []byte
	data = append(data, 0)

	b := make([]byte, 2)
	binary.LittleEndian.PutUint16(b, uint16(requireN))
	data = append(data, b[:1]...)

	b = make([]byte, 2)
	binary.LittleEndian.PutUint16(b, uint16(threshold))
	data = append(data, b[:1]...)

	b = make([]byte, 2)
	binary.LittleEndian.PutUint16(b, uint16(len(publicKeysHash)))
	data = append(data, b[:1]...)

	for _, hash := range publicKeysHash {
		data = append(data, hash...)
	}

	args, err := blake2b.Blake160(data)
	if err != nil {
		return nil, nil, err
	}

	return &types.Script{
		CodeHash: types.HexToHash(transaction.SECP256K1_BLAKE160_MULTISIG_ALL_TYPE_HASH),
		HashType: types.HashTypeType,
		Args:     args,
	}, data, nil
}
