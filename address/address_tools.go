package address

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/nervosnetwork/ckb-sdk-go/crypto/bech32"
	"github.com/nervosnetwork/ckb-sdk-go/crypto/blake2b"
	"github.com/nervosnetwork/ckb-sdk-go/crypto/secp256k1"
	"github.com/nervosnetwork/ckb-sdk-go/transaction"
	"github.com/nervosnetwork/ckb-sdk-go/types"
	"github.com/pkg/errors"
)

const (
	MAINNET_ACP_CODE_HASH    = "0xd369597ff47f29fbc0d47d2e3775370d1250b85140c670e4718af712983a2354"
	TESTNET_ACP_CODE_HASH    = "0x3419a1c09eb2567f6552ee7a8ecffd64155cffe0f1796e6e61ec088d740c1356"
	MAINNET_CHEQUE_CODE_HASH = "0xe4d4ecc6e5f9a059bf2f7a82cca292083aebc0c421566a52484fe2ec51a9fb0c"
	TESTNET_CHEQUE_CODE_HASH = "0x60d5f39efce409c587cb9ea359cefdead650ca128f0bd9cb3855348f98c70d5b"
)

type AddressGenerateResult struct {
	Address    string
	LockArgs   string
	PrivateKey string
}

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

func GenerateAcpAddress(secp256k1Address string) (string, error) {
	addressScript, err := Parse(secp256k1Address)
	if err != nil {
		return "", err
	}

	script := &types.Script{
		CodeHash: types.HexToHash(getAcpCodeHash(addressScript.Mode)),
		HashType: types.HashTypeType,
		Args:     common.FromHex(hex.EncodeToString(addressScript.Script.Args)),
	}

	return Generate(addressScript.Mode, script)
}

func GenerateChequeAddress(senderAddress, receiverAddress string) (string, error) {
	senderScript, err := Parse(senderAddress)
	if err != nil {
		return "", err
	}
	receiverScript, err := Parse(receiverAddress)
	if err != nil {
		return "", err
	}

	if senderScript.Mode != receiverScript.Mode {
		return "", errors.New("The network type of senderAddress and receiverAddress must be the same")
	}

	senderScriptHash, err := senderScript.Script.Hash()
	if err != nil {
		return "", err
	}
	receiverScriptHash, err := receiverScript.Script.Hash()
	if err != nil {
		return "", err
	}

	s1 := senderScriptHash.String()[0:42]
	s2 := receiverScriptHash.String()[0:42]

	args := bytesCombine(common.FromHex(s2), common.FromHex(s1))
	pubKey := common.Bytes2Hex(args)
	fmt.Printf("pubKey: %s\n", pubKey)

	chequeLock := &types.Script{
		CodeHash: types.HexToHash(getChequeCodeHash(senderScript.Mode)),
		HashType: types.HashTypeType,
		Args:     common.FromHex(pubKey),
	}

	return Generate(senderScript.Mode, chequeLock)

}

func GenerateBech32mFullAddress(mode Mode, script *types.Script) (string, error) {

	hashType, err := types.SerializeHashType(script.HashType)
	if err != nil {
		return "", err
	}

	// Payload: type(00) | code hash | hash type | args
	payload := TYPE_FULL_WITH_BECH32M
	payload += script.CodeHash.Hex()[2:]
	payload += hashType

	payload += common.Bytes2Hex(script.Args)

	dataPart, err := bech32.ConvertBits(common.FromHex(payload), 8, 5, true)
	if err != nil {
		return "", err
	}
	return bech32.EncodeWithBech32m(string(mode), dataPart)
}

func getHashType(hashType types.ScriptHashType) string {
	if hashType == types.HashTypeType {
		return "01"
	} else {
		return "00"
	}
}

func getAcpCodeHash(mode Mode) string {
	if mode == Mainnet {
		return MAINNET_ACP_CODE_HASH
	} else {
		return TESTNET_ACP_CODE_HASH
	}
}

func getChequeCodeHash(mode Mode) string {
	if mode == Mainnet {
		return MAINNET_CHEQUE_CODE_HASH
	} else {
		return TESTNET_CHEQUE_CODE_HASH
	}
}

func bytesCombine(pBytes ...[]byte) []byte {
	return bytes.Join(pBytes, []byte(""))
}
