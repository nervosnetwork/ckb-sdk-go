package types

import (
	"github.com/nervosnetwork/ckb-sdk-go/molecule"
)

func DeserializeWitnessArgs(in []byte) (*WitnessArgs, error) {
	m, err := molecule.WitnessArgsFromSlice(in, false)
	if err != nil {
		return nil, err
	}
	return UnpackWitnessArgs(m), nil
}
