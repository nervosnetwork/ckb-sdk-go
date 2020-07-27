package address

import (
	"encoding/hex"
	"errors"
	"strings"

	"github.com/ethereum/go-ethereum/common"

	"github.com/nervosnetwork/ckb-sdk-go/crypto/bech32"
	"github.com/nervosnetwork/ckb-sdk-go/transaction"
	"github.com/nervosnetwork/ckb-sdk-go/types"
)

type Mode string
type Type string

const (
	Mainnet Mode = "ckb"
	Testnet Mode = "ckt"

	TypeFull  Type = "Full"
	TypeShort Type = "Short"

	SHORT_FORMAT                 = "01"
	FULL_DATA_FORMAT             = "02"
	FULL_TYPE_FORMAT             = "04"
	CODE_HASH_INDEX_SINGLESIG    = "00"
	CODE_HASH_INDEX_MULTISIG_SIG = "01"
)

type ParsedAddress struct {
	Mode   Mode
	Type   Type
	Script *types.Script
}

func Generate(mode Mode, script *types.Script) (string, error) {
	if script.HashType == types.HashTypeType && len(script.Args) == 20 {
		if transaction.SECP256K1_BLAKE160_SIGHASH_ALL_TYPE_HASH == script.CodeHash.String() {
			// generate_short_payload_singlesig_address
			payload := SHORT_FORMAT + CODE_HASH_INDEX_SINGLESIG + hex.EncodeToString(script.Args)
			data, err := bech32.ConvertBits(common.FromHex(payload), 8, 5, true)
			if err != nil {
				return "", err
			}
			return bech32.Encode((string)(mode), data)
		} else if transaction.SECP256K1_BLAKE160_MULTISIG_ALL_TYPE_HASH == script.CodeHash.String() {
			// generate_short_payload_multisig_address
			payload := SHORT_FORMAT + CODE_HASH_INDEX_MULTISIG_SIG + hex.EncodeToString(script.Args)
			data, err := bech32.ConvertBits(common.FromHex(payload), 8, 5, true)
			if err != nil {
				return "", err
			}
			return bech32.Encode((string)(mode), data)
		}
	}

	hashType := FULL_TYPE_FORMAT
	if script.HashType == types.HashTypeData {
		hashType = FULL_DATA_FORMAT
	}

	return generateFullPayloadAddress(hashType, mode, script)
}

func generateFullPayloadAddress(hashType string, mode Mode, script *types.Script) (string, error) {
	payload := hashType + hex.EncodeToString(script.CodeHash.Bytes()) + hex.EncodeToString(script.Args)
	data, err := bech32.ConvertBits(common.FromHex(payload), 8, 5, true)
	if err != nil {
		return "", err
	}
	return bech32.Encode((string)(mode), data)
}

func Parse(address string) (*ParsedAddress, error) {
	hrp, decoded, err := bech32.Decode(address)
	if err != nil {
		return nil, err
	}
	data, err := bech32.ConvertBits(decoded, 5, 8, false)
	if err != nil {
		return nil, err
	}
	payload := hex.EncodeToString(data)

	var addressType Type
	var script types.Script
	if strings.HasPrefix(payload, "01") {
		addressType = TypeShort
		if CODE_HASH_INDEX_SINGLESIG == payload[2:4] {
			script = types.Script{
				CodeHash: types.HexToHash(transaction.SECP256K1_BLAKE160_SIGHASH_ALL_TYPE_HASH),
				HashType: types.HashTypeType,
				Args:     common.Hex2Bytes(payload[4:]),
			}
		} else {
			script = types.Script{
				CodeHash: types.HexToHash(transaction.SECP256K1_BLAKE160_MULTISIG_ALL_TYPE_HASH),
				HashType: types.HashTypeType,
				Args:     common.Hex2Bytes(payload[4:]),
			}
		}
	} else if strings.HasPrefix(payload, "02") {
		addressType = TypeFull
		script = types.Script{
			CodeHash: types.HexToHash(payload[2:66]),
			HashType: types.HashTypeData,
			Args:     common.Hex2Bytes(payload[66:]),
		}
	} else if strings.HasPrefix(payload, "04") {
		addressType = TypeFull
		script = types.Script{
			CodeHash: types.HexToHash(payload[2:66]),
			HashType: types.HashTypeType,
			Args:     common.Hex2Bytes(payload[66:]),
		}
	} else {
		return nil, errors.New("address type error:" + payload[:2])
	}

	result := &ParsedAddress{
		Mode:   Mode(hrp),
		Type:   addressType,
		Script: &script,
	}
	return result, nil
}
