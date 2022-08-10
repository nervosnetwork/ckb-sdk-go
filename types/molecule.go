package types

import (
	"encoding/binary"
	"github.com/nervosnetwork/ckb-sdk-go/molecule"
)

func (r *WitnessArgs) Pack() *molecule.WitnessArgs {
	builder := molecule.NewWitnessArgsBuilder()
	builder.Lock(*PackBytesToOpt(r.Lock))
	builder.InputType(*PackBytesToOpt(r.InputType))
	builder.OutputType(*PackBytesToOpt(r.OutputType))
	b := builder.Build()
	return &b
}

func (r *Transaction) PackToRawTransaction() *molecule.RawTransaction {
	builder := molecule.NewRawTransactionBuilder()
	builder.Version(*PackUint32(uint32(r.Version)))

	cellDepsVecBuilder := molecule.NewCellDepVecBuilder()
	for _, v := range r.CellDeps {
		cellDepsVecBuilder.Push(*v.Pack())
	}
	builder.CellDeps(cellDepsVecBuilder.Build())

	byte32VecBuilder := molecule.NewByte32VecBuilder()
	for _, v := range r.HeaderDeps {
		byte32VecBuilder.Push(*v.Pack())
	}
	builder.HeaderDeps(byte32VecBuilder.Build())

	inputVecBuilder := molecule.NewCellInputVecBuilder()
	for _, v := range r.Inputs {
		inputVecBuilder.Push(*v.Pack())
	}
	builder.Inputs(inputVecBuilder.Build())

	outputVecBuilder := molecule.NewCellOutputVecBuilder()
	for _, v := range r.Outputs {
		outputVecBuilder.Push(*v.Pack())
	}
	builder.Outputs(outputVecBuilder.Build())

	builder.OutputsData(*PackBytesVec(r.OutputsData))
	b := builder.Build()
	return &b
}

func (r *Transaction) Pack() *molecule.Transaction {
	builder := molecule.NewTransactionBuilder()
	builder.Raw(*r.PackToRawTransaction())
	builder.Witnesses(*PackBytesVec(r.Witnesses))
	b := builder.Build()
	return &b
}

func (r *Script) Pack() *molecule.Script {
	builder := molecule.NewScriptBuilder()
	builder.CodeHash(*r.CodeHash.Pack())
	builder.HashType(*r.HashType.Pack())
	builder.Args(*PackBytes(r.Args))
	b := builder.Build()
	return &b
}

func (t DepType) Pack() *molecule.Byte {
	var b byte
	switch t {
	case DepTypeCode:
		b = 0x00
	case DepTypeDepGroup:
		b = 0x01
	default:
		panic("invalid dep type")
	}
	return PackByte(b)
}

func (t *CellDep) Pack() *molecule.CellDep {
	builder := molecule.NewCellDepBuilder()
	builder.OutPoint(*t.OutPoint.Pack())
	builder.DepType(*t.DepType.Pack())
	b := builder.Build()
	return &b
}

func (r *OutPoint) Pack() *molecule.OutPoint {
	builder := molecule.NewOutPointBuilder()
	builder.TxHash(*r.TxHash.Pack())
	builder.Index(*PackUint32(uint32(r.Index)))
	b := builder.Build()
	return &b
}

func (r *CellInput) Pack() *molecule.CellInput {
	builder := molecule.NewCellInputBuilder()
	builder.PreviousOutput(*r.PreviousOutput.Pack())
	builder.Since(*PackUint64(r.Since))
	b := builder.Build()
	return &b
}

func (r *CellOutput) Pack() *molecule.CellOutput {
	builder := molecule.NewCellOutputBuilder()
	builder.Capacity(*PackUint64(r.Capacity))
	builder.Lock(*r.Lock.Pack())
	builder.Type(*packScriptToOpt(r.Type))
	b := builder.Build()
	return &b
}

func packScriptToOpt(r *Script) *molecule.ScriptOpt {
	builder := molecule.NewScriptOptBuilder()
	if r != nil {
		builder.Set(*r.Pack())
	}
	v := builder.Build()
	return &v
}

func UnpackBytes(v *molecule.Bytes) []byte {
	return v.AsSlice()
}

func PackBytes(v []byte) *molecule.Bytes {
	return molecule.BytesFromSliceUnchecked(v)
}

func PackBytesVec(v [][]byte) *molecule.BytesVec {
	builder := molecule.NewBytesVecBuilder()
	for _, v := range v {
		builder.Push(*PackBytes(v))
	}
	b := builder.Build()
	return &b
}

func PackBytesToOpt(v []byte) *molecule.BytesOpt {
	builder := molecule.NewBytesOptBuilder()
	if v != nil {
		builder.Set(*PackBytes(v))
	}
	b := builder.Build()
	return &b
}

func PackByte(v byte) *molecule.Byte {
	return molecule.ByteFromSliceUnchecked([]byte{v})
}

func PackUint64(v uint64) *molecule.Uint64 {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, v)
	return molecule.Uint64FromSliceUnchecked(b)
}

func PackUint32(v uint32) *molecule.Uint32 {
	b := make([]byte, 4)
	binary.LittleEndian.PutUint32(b, v)
	return molecule.Uint32FromSliceUnchecked(b)
}

func UnpackHash(v *molecule.Byte32) Hash {
	return BytesToHash(v.AsSlice())
}

func (h *Hash) Pack() *molecule.Byte32 {
	return molecule.Byte32FromSliceUnchecked(h.Bytes())
}

func (t ScriptHashType) Pack() *molecule.Byte {
	var b byte
	switch t {
	case HashTypeData:
		b = 0x00
	case HashTypeType:
		b = 0x01
	case HashTypeData1:
		b = 0x02
	default:
		panic("invalid script hash type")
	}
	return PackByte(b)
}
