package signer

import (
	"bytes"
	"errors"
	"github.com/nervosnetwork/ckb-sdk-go/crypto/blake2b"
	"github.com/nervosnetwork/ckb-sdk-go/crypto/secp256k1"
	"github.com/nervosnetwork/ckb-sdk-go/transaction"
	"github.com/nervosnetwork/ckb-sdk-go/types"
	"reflect"
)

type Secp256k1Blake160MultisigAllSigner struct {
}

func (s *Secp256k1Blake160MultisigAllSigner) SignTransaction(transaction *types.Transaction, group *transaction.ScriptGroup, ctx *transaction.Context) (bool, error) {
	var m *MultisigScript
	switch ctx.Payload.(type) {
	case MultisigScript:
		mm := ctx.Payload.(MultisigScript)
		m = &mm
	case *MultisigScript:
		m = ctx.Payload.(*MultisigScript)
	default:
		return false, nil
	}
	matched, err := IsMultiSigMatched(ctx.Key, m, group.Script.Args)
	if err != nil {
		return false, err
	}
	if matched {
		return MultiSignTransaction(transaction, uint32ArrayToIntArray(group.InputIndices), ctx.Key, m)
	} else {
		return false, nil
	}
}

func MultiSignTransaction(tx *types.Transaction, group []int, key *secp256k1.Secp256k1Key, m *MultisigScript) (bool, error) {
	var err error
	i0 := group[0]
	witnessPlaceholder, err := m.WitnessPlaceholder(tx.Witnesses[i0])
	if err != nil {
		return false, nil
	}
	signature, err := SignTransaction(tx, group, witnessPlaceholder, key)
	if err != nil {
		return false, err
	}
	if tx.Witnesses[i0], err = setSignatureToWitness(tx.Witnesses[i0], signature, m); err != nil {
		return false, err
	}
	return true, nil
}

func setSignatureToWitness(witness []byte, signature []byte, m *MultisigScript) ([]byte, error) {
	witnessArgs, err := types.DeserializeWitnessArgs(witness)
	if err != nil {
		return nil, err
	}
	lock := witnessArgs.Lock
	pos := len(m.Encode())
	emptySignature := [65]byte{}
	for i := 0; i < int(m.Threshold); i++ {
		if reflect.DeepEqual(emptySignature[:], lock[pos:pos+65]) {
			copy(lock[pos:pos+65], signature[:])
			break
		}
		pos += 65
	}
	witnessArgs.Lock = lock
	w := witnessArgs.Serialize()
	return w, err
}

func IsMultiSigMatched(key *secp256k1.Secp256k1Key, multisigScript *MultisigScript, scriptArgs []byte) (bool, error) {
	if key == nil || scriptArgs == nil {
		return false, errors.New("key or scriptArgs is nil")
	}
	hash := multisigScript.ComputeHash160()
	return bytes.Equal(scriptArgs, hash), nil
}

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

func (r *MultisigScript) Encode() []byte {
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

func (r *MultisigScript) WitnessPlaceholder(originalWitness []byte) ([]byte, error) {
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

func (r *MultisigScript) WitnessPlaceholderInLock() []byte {
	header := r.Encode()
	b := make([]byte, len(header)+65*int(r.Threshold))
	copy(b[:len(header)], header)
	return b
}

func (r *MultisigScript) ComputeHash160() []byte {
	return blake2b.Blake160(r.Encode()[:])
}
