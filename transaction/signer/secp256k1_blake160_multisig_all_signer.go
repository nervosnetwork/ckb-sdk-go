package signer

import (
	"errors"
	"github.com/nervosnetwork/ckb-sdk-go/crypto/blake2b"
)

type MultisigScript struct {
	Version    byte
	FirstN     byte
	Threshold  byte
	KeysHashes [][20]byte
}

func NewMultisigScript(firstN byte, threshold byte) *MultisigScript {
	return &MultisigScript{
		Version:    0,
		FirstN:     firstN,
		Threshold:  threshold,
		KeysHashes: make([][20]byte, 0),
	}
}

func (r *MultisigScript) AddKeyHash(keyHash [20]byte) {
	r.KeysHashes = append(r.KeysHashes, keyHash)
}

func (r *MultisigScript) AddKeyHashBySlice(keyHash []byte) error {
	if keyHash == nil {
		return errors.New("keyHash is nil")
	}
	if len(keyHash) != 20 {
		return errors.New("keyHash length should be 20-byte")
	}
	k := [20]byte{}
	copy(k[:], keyHash)
	r.AddKeyHash(k)
	return nil
}

func (r *MultisigScript) encode() []byte {
	out := make([]byte, 4)
	out[0] = r.Version
	out[1] = r.FirstN
	out[2] = r.Threshold
	out[3] = byte(len(r.KeysHashes))
	for _, b := range r.KeysHashes {
		out = append(out, b[:]...)
	}
	return out
}

func DecodeToMultisigScript(in []byte) (*MultisigScript, error) {
	l := len(in)
	if l < 24 {
		return nil, errors.New("bytes length should be greater than 24")
	}
	if (l-4)%4 != 0 {
		return nil, errors.New("invalid bytes length")
	}
	if l != int(in[3])*20+4 {
		return nil, errors.New("invalid public key list size")
	}
	m := &MultisigScript{
		Version:    in[0],
		FirstN:     in[1],
		Threshold:  in[2],
		KeysHashes: make([][20]byte, 0),
	}
	for i := 0; i < int(in[3]); i++ {
		var b [20]byte
		copy(b[:], in[4+i*20:4+i*20+20])
		m.KeysHashes = append(m.KeysHashes, b)
	}
	return m, nil
}

func (r *MultisigScript) ComputeHash() ([20]byte, error) {
	hash, err := blake2b.Blake160(r.encode()[:])
	if err != nil {
		return [20]byte{}, err
	}
	var arr [20]byte
	copy(arr[:], hash[:20])
	return arr, nil
}
