package types

import (
	"encoding/binary"
)

const u32Size uint = 4

func SerializeUint(n uint) []byte {
	b := make([]byte, u32Size)
	binary.LittleEndian.PutUint32(b, uint32(n))

	return b
}

func SerializeUint64(n uint64) []byte {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, n)

	return b
}
