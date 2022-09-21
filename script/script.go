package script

import (
	"errors"
	"github.com/nervosnetwork/ckb-sdk-go/crypto/blake2b"
	"github.com/nervosnetwork/ckb-sdk-go/types"
)

type MultisigConfig struct {
	Version    byte
	FirstN     byte
	Threshold  byte
	KeysHashes [][20]byte
}

func NewMultisigConfig(firstN byte, threshold byte) *MultisigConfig {
	return &MultisigConfig{
		Version:    0,
		FirstN:     firstN,
		Threshold:  threshold,
		KeysHashes: make([][20]byte, 0),
	}
}

// AddKeyHash adds key hash, and panic if keyHash is shorter than 20 bytes.
func (r *MultisigConfig) AddKeyHash(keyHash []byte) {
	var h [20]byte
	copy(h[:], keyHash[:20])
	r.KeysHashes = append(r.KeysHashes, h)
}

func (r *MultisigConfig) Encode() []byte {
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

func DecodeToMultisigConfig(in []byte) (*MultisigConfig, error) {
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
	m := &MultisigConfig{
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

func (r *MultisigConfig) WitnessPlaceholder(originalWitness []byte) ([]byte, error) {
	var (
		witnessArgs *types.WitnessArgs
		err         error
	)
	if len(originalWitness) == 0 {
		witnessArgs = &types.WitnessArgs{}
	} else {
		if witnessArgs, err = types.DeserializeWitnessArgs(originalWitness); err != nil {
			return nil, err
		}
	}
	witnessArgs.Lock = r.WitnessPlaceholderInLock()
	b := witnessArgs.Serialize()
	return b, nil
}

func (r *MultisigConfig) WitnessPlaceholderInLock() []byte {
	header := r.Encode()
	b := make([]byte, len(header)+65*int(r.Threshold))
	copy(b[:len(header)], header)
	return b
}

func (r *MultisigConfig) Hash160() []byte {
	return blake2b.Blake160(r.Encode()[:])
}
