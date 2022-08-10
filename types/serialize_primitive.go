package types

const u32Size uint = 4

func SerializeUint32(n uint32) []byte {
	return PackUint32(n).AsSlice()
}

func SerializeUint64(n uint64) []byte {
	return PackUint64(n).AsSlice()
}
