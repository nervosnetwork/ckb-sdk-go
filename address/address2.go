package address

import (
	"errors"
	"github.com/nervosnetwork/ckb-sdk-go/types"
)

type Address struct {
	Script  types.Script
	Network types.Network
}

func Decode(s string) (*Address, error) {
	return nil, nil
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
