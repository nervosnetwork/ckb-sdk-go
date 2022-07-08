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
		return decodeLongBech32M(data, network)
	case 0x01:
		return decodeShort(data, network)
	case 0x02, 0x04:
		return decodeLongBech32(data, network)
	default:
		fmt.Println("Unkown")
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
	fmt.Println("bech32m")
	if payload[0] != 0x00 {
		return nil, errors.New(fmt.Sprintf("invalid payload header 0x%d", payload[0]))
	}
	codeHash := types.BytesToHash(payload[1:33])
	// TODO: Extract function to convert byte to hashType
	var hashType types.ScriptHashType
	switch payload[33] {
	case 0x00:
		hashType = types.HashTypeData
	case 0x01:
		hashType = types.HashTypeType
	case 0x02:
		hashType = types.HashTypeData1
	default:
		return nil, errors.New("unknown script hash type")
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


func (a *Address) Encode() (string, error) {
	return a.EncodeFullBech32m()
}

func (a *Address) EncodeShort() (string, error) {
	return "", nil
}

func (a *Address) EncodeFullBech32() (string, error) {
	return "", nil
}

func (a *Address) EncodeFullBech32m() (string, error) {
	return "", nil
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
