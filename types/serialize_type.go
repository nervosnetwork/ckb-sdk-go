package types

import (
	"errors"
	"fmt"
	"github.com/nervosnetwork/ckb-sdk-go/molecule"
)

func (h Hash) Serialize() []byte {
	return h.Pack().AsSlice()
}

func (t ScriptHashType) Serialize() []byte {
	return t.Pack().AsSlice()
}

// Serialize dep type
func (t DepType) Serialize() []byte {
	return t.Pack().AsSlice()
}

// Serialize script
func (r *Script) Serialize() []byte {
	return r.Pack().AsSlice()
}

// Serialize outpoint
func (r *OutPoint) Serialize() []byte {
	return r.Pack().AsSlice()
}

// Serialize cell input
func (r *CellInput) Serialize() []byte {
	return r.Pack().AsSlice()
}

// Serialize cell output
func (r *CellOutput) Serialize() []byte {
	return r.Pack().AsSlice()
}

// Serialize cell dep
func (d *CellDep) Serialize() []byte {
	return d.Pack().AsSlice()
}

// Serialize transaction
func (t *Transaction) Serialize() []byte {
	return t.PackToRawTransaction().AsSlice()
}

func (t *Transaction) SerializeWithWitness() []byte {
	return t.Pack().AsSlice()
}

func (w *WitnessArgs) Serialize() []byte {
	return w.Pack().AsSlice()
}

func SerializeHashTypeByte(hashType ScriptHashType) (byte, error) {
	switch hashType {
	case HashTypeData:
		return 0x00, nil
	case HashTypeType:
		return 0x01, nil
	case HashTypeData1:
		return 0x02, nil
	default:
		return 0, errors.New(string("unknown hash type " + hashType))
	}
}

func DeserializeHashTypeByte(hashType byte) (ScriptHashType, error) {
	switch hashType {
	case 0x00:
		return HashTypeData, nil
	case 0x01:
		return HashTypeType, nil
	case 0x02:
		return HashTypeData1, nil
	default:
		return "", errors.New(fmt.Sprintf("invalid script hash_type: %x", hashType))
	}
}

func DeserializeWitnessArgs(in []byte) (*WitnessArgs, error) {
	m, err := molecule.WitnessArgsFromSlice(in, false)
	if err != nil {
		return nil, err
	}
	return UnpackWitnessArgs(m), nil
}
