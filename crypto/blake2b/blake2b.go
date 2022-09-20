package blake2b

import (
	"github.com/minio/blake2b-simd"
)

var ckbHashPersonalization = []byte("ckb-default-hash")

func Blake160(data []byte) []byte {
	return Blake256(data)[:20]
}

func Blake256(data []byte) []byte {
	config := &blake2b.Config{
		Size:   32,
		Person: ckbHashPersonalization,
	}
	hash, _ := blake2b.New(config)
	hash.Write(data)
	return hash.Sum(nil)
}
