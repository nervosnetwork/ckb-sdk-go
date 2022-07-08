package address

import (
	"errors"
	"fmt"
	"github.com/nervosnetwork/ckb-sdk-go/crypto/bech32"
	"github.com/nervosnetwork/ckb-sdk-go/types"
)

type Address struct {
	Script  types.Script
	Network types.Network
}

func Decode(s string) (*Address, error) {
	encoding, hrp, decoded, err := bech32.Decode(s)
	if err != nil {
		return nil, err
	}
	network, err := fromHrp(hrp)
	if err != nil {
		return nil, err
	}
	data, err := bech32.ConvertBits(decoded, 5, 8, false)
	if err != nil {
		return nil, err
	}
	switch data[0] {
	case 0x00:
		if encoding != bech32.BECH32M {
			return nil, errors.New("payload header 0x00 must have encoding bech32m")
		}
		return decodeLongBech32M(data, network)
	case 0x01:
		if encoding != bech32.BECH32 {
			return nil, errors.New("payload header 0x01 must have encoding bech32")
		}
		return decodeShort(data, network)
	case 0x02, 0x04:
		if encoding != bech32.BECH32 {
			return nil, errors.New("payload header 0x02 or 0x04 must have encoding bech32")
		}
		return decodeLongBech32(data, network)
	default:
		return nil, errors.New("unknown address format type")
	}
}

func decodeShort(payload []byte, network types.Network) (*Address, error) {
	codeHashIndex := payload[1]
	args := payload[2:]
	argsLen := len(args)
	var scriptType types.BuiltinScript
	switch codeHashIndex {
	case 0x00: // secp256k1_blake160_sighash_all
		if argsLen != 20 {
			return nil, errors.New(fmt.Sprintf("invalid args length %d", argsLen))
		}
		scriptType = types.BuiltinScriptSecp256k1Blake160SighashAll
	case 0x01: // secp256k1_blake160_multisig_all
		if argsLen != 20 {
			return nil, errors.New(fmt.Sprintf("invalid args length %d", argsLen))
		}
		scriptType = types.BuiltinScriptSecp256k1Blake160MultisigAll
	case 0x02: // any_can_pay
		if argsLen < 20 || argsLen > 22 {
			return nil, errors.New(fmt.Sprintf("invalid args length %d", argsLen))
		}
		scriptType = types.BuiltinScriptAnyoneCanPay
	default:
		return nil, errors.New("unknown code hash index")
	}
	codeHash := types.GetCodeHash(scriptType, network)
	return &Address{
		Script: types.Script{
			CodeHash: codeHash,
			HashType: types.HashTypeType,
			Args:     args,
		},
		Network: network,
	}, nil
}

func decodeLongBech32(payload []byte, network types.Network) (*Address, error) {
	var hashType types.ScriptHashType
	switch payload[0] {
	case 0x04:
		hashType = types.HashTypeType
	case 0x02:
		hashType = types.HashTypeData
	default:
		return nil, errors.New("unknown script hash type")
	}
	codeHash := types.BytesToHash(payload[1:33])
	args := payload[33:]
	return &Address{
		Script: types.Script{
			CodeHash: codeHash,
			HashType: hashType,
			Args:     args,
		},
		Network: network,
	}, nil
}

func decodeLongBech32M(payload []byte, network types.Network) (*Address, error) {
	if payload[0] != 0x00 {
		return nil, errors.New(fmt.Sprintf("invalid payload header 0x%d", payload[0]))
	}
	codeHash := types.BytesToHash(payload[1:33])
	hashType, err := types.DeserializeHashTypeByte(payload[33])
	if err != nil {
		return nil, err
	}
	args := payload[34:]
	return &Address{
		Script: types.Script{
			CodeHash: codeHash,
			HashType: hashType,
			Args:     args,
		},
		Network: network,
	}, nil
}

func (a Address) Encode() (string, error) {
	return a.EncodeFullBech32m()
}

func (a Address) EncodeShort() (string, error) {
	payload := make([]byte, 0)
	payload = append(payload, 0x01)
	if a.Script.CodeHash == types.GetCodeHash(types.BuiltinScriptSecp256k1Blake160SighashAll, a.Network) {
		payload = append(payload, 0x00)
	} else if a.Script.CodeHash == types.GetCodeHash(types.BuiltinScriptSecp256k1Blake160MultisigAll, a.Network) {
		payload = append(payload, 0x01)
	} else if a.Script.CodeHash == types.GetCodeHash(types.BuiltinScriptAnyoneCanPay, a.Network) {
		payload = append(payload, 0x02)
	} else {
		return "", errors.New("encoding to short address for given script is unsupported")
	}
	payload = append(payload, a.Script.Args...)
	payload, err := bech32.ConvertBits(payload, 8, 5, true)
	if err != nil {
		return "", err
	}
	hrp, err := toHrp(a.Network)
	if err != nil {
		return "", err
	}
	return bech32.Encode(hrp, payload)
}

func (a Address) EncodeFullBech32() (string, error) {
	payload := make([]byte, 0)
	if a.Script.HashType == types.HashTypeType {
		payload = append(payload, 0x04)
	} else if a.Script.HashType == types.HashTypeData {
		payload = append(payload, 0x02)
	} else {
		return "", errors.New(string("unknown hash type " + a.Script.HashType))
	}
	payload = append(payload, a.Script.CodeHash.Bytes()...)
	payload = append(payload, a.Script.Args...)
	payload, err := bech32.ConvertBits(payload, 8, 5, true)
	if err != nil {
		return "", err
	}
	hrp, err := toHrp(a.Network)
	if err != nil {
		return "", err
	}
	return bech32.Encode(hrp, payload)
}

func (a Address) EncodeFullBech32m() (string, error) {
	payload := make([]byte, 0)
	payload = append(payload, 0x00)
	payload = append(payload, a.Script.CodeHash.Bytes()...)
	hashType, err := types.SerializeHashTypeByte(a.Script.HashType)
	if err != nil {
		return "", err
	}
	payload = append(payload, hashType)
	payload = append(payload, a.Script.Args...)
	if payload, err = bech32.ConvertBits(payload, 8, 5, true); err != nil {
		return "", err
	}
	hrp, err := toHrp(a.Network)
	if err != nil {
		return "", err
	}
	return bech32.EncodeWithBech32m(hrp, payload)
}

func toHrp(network types.Network) (string, error) {
	switch network {
	case types.NetworkMain:
		return "ckb", nil
	case types.NetworkTest:
		return "ckt", nil
	default:
		return "", errors.New("unknown network")
	}
}

func fromHrp(hrp string) (types.Network, error) {
	switch hrp {
	case "ckb":
		return types.NetworkMain, nil
	case "ckt":
		return types.NetworkTest, nil
	default:
		return 0, errors.New("unknown network")
	}
}
