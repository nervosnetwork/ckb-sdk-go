package address

import (
	"encoding/hex"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/nervosnetwork/ckb-sdk-go/crypto/blake2b"
	"github.com/nervosnetwork/ckb-sdk-go/crypto/secp256k1"
	"github.com/nervosnetwork/ckb-sdk-go/transaction"
	"github.com/nervosnetwork/ckb-sdk-go/types"
)

func GenerateShortAddress(mode Mode) (*AddressGenerateResult, error) {

	key, err := secp256k1.RandomNew()
	if err != nil {
		return nil, err
	}

	pubKey, err := blake2b.Blake160(key.PubKey())
	if err != nil {
		return nil, err
	}

	script := &types.Script{
		CodeHash: types.HexToHash(transaction.SECP256K1_BLAKE160_SIGHASH_ALL_TYPE_HASH),
		HashType: types.HashTypeType,
		Args:     common.FromHex(hex.EncodeToString(pubKey)),
	}

	address, err := Generate(mode, script)
	if err != nil {
		return nil, err
	}

	return &AddressGenerateResult{
		Address:    address,
		LockArgs:   hexutil.Encode(pubKey),
		PrivateKey: hexutil.Encode(key.Bytes()),
	}, err

}

type AddressGenerateResult struct {
	Address    string
	LockArgs   string
	PrivateKey string
}
