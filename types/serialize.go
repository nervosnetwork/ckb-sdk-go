package types

import (
	"errors"
	"fmt"
	"github.com/nervosnetwork/ckb-sdk-go/v2/types/molecule"
)

func SerializeUint32(n uint32) []byte {
	return PackUint32(n).AsSlice()
}

func SerializeUint64(n uint64) []byte {
	return PackUint64(n).AsSlice()
}

func (h Hash) Serialize() []byte {
	return h.Pack().AsSlice()
}

func (t ScriptHashType) Serialize() []byte {
	return t.Pack().AsSlice()
}

func (t DepType) Serialize() []byte {
	return t.Pack().AsSlice()
}

func (r *Script) Serialize() []byte {
	return r.Pack().AsSlice()
}

func (r *OutPoint) Serialize() []byte {
	return r.Pack().AsSlice()
}

func (r *CellInput) Serialize() []byte {
	return r.Pack().AsSlice()
}

func (r *CellOutput) Serialize() []byte {
	return r.Pack().AsSlice()
}

func (d *CellDep) Serialize() []byte {
	return d.Pack().AsSlice()
}

func (t *Transaction) SerializeWithoutWitnesses() []byte {
	return t.PackToRawTransaction().AsSlice()
}

func (t *Transaction) Serialize() []byte {
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
	case HashTypeData2:
		return 0x04, nil
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
	case 0x04:
		return HashTypeData2, nil
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
