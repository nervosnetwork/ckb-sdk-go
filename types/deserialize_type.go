package types

import (
	"encoding/binary"
	"errors"
)

func DeserializeWitnessArgs(in []byte) (*WitnessArgs, error) {
	length := binary.LittleEndian.Uint32(in[0:4])
	if length != uint32(len(in)) {
		return nil, errors.New("incorrect bytes length")
	}
	offsetLock := binary.LittleEndian.Uint32(in[4:8])
	offsetInputType := binary.LittleEndian.Uint32(in[8:12])
	offsetOutputType := binary.LittleEndian.Uint32(in[12:16])

	lockBytesLength := offsetInputType - offsetLock
	var lock []byte
	if lockBytesLength != 0 {
		lockLength := binary.LittleEndian.Uint32(in[offsetLock : offsetLock+4])
		if lockLength+4 != lockBytesLength {
			return nil, errors.New("incorrect lock bytes length")
		}
		lock = in[offsetLock+4 : offsetLock+4+lockLength]
	}

	inputTypeBytesLength := offsetOutputType - offsetInputType
	var inputType []byte
	if inputTypeBytesLength != 0 {
		inputTypeLength := binary.LittleEndian.Uint32(in[offsetInputType : offsetInputType+4])
		if inputTypeLength+4 != inputTypeBytesLength {
			return nil, errors.New("incorrect input type bytes length")
		}
		inputType = in[offsetInputType+4 : offsetInputType+4+inputTypeLength]
	}

	OutputTypeBytesLength := length - offsetOutputType
	var OutputType []byte
	if OutputTypeBytesLength != 0 {
		outputTypeLength := binary.LittleEndian.Uint32(in[offsetOutputType : offsetOutputType+4])
		if outputTypeLength+4 != OutputTypeBytesLength {
			return nil, errors.New("incorrect output type bytes length")
		}
		OutputType = in[offsetOutputType+4 : offsetOutputType+4+outputTypeLength]
	}

	return &WitnessArgs{
		Lock:       lock,
		InputType:  inputType,
		OutputType: OutputType,
	}, nil
}
