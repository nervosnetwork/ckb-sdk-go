package secp256k1

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
	"math/big"
)

var (
	secp256k1N, _  = new(big.Int).SetString("fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364141", 16)
	secp256k1halfN = new(big.Int).Div(secp256k1N, big.NewInt(2))
)

type Secp256k1Key struct {
	PrivateKey *ecdsa.PrivateKey
}

func (k *Secp256k1Key) Bytes() []byte {
	return math.PaddedBigBytes(k.PrivateKey.D, k.PrivateKey.Params().BitSize/8)
}

func (k *Secp256k1Key) Sign(data []byte) ([]byte, error) {
	seckey := k.Bytes()

	defer func(bytes []byte) {
		for i := range bytes {
			bytes[i] = 0
		}
	}(seckey)

	return secp256k1.Sign(data, seckey)
}

func (k *Secp256k1Key) PubKey() []byte {
	pub := &k.PrivateKey.PublicKey
	if pub == nil || pub.X == nil || pub.Y == nil {
		return nil
	}

	return secp256k1.CompressPubkey(pub.X, pub.Y)
}

func (k *Secp256k1Key) PubKeyUncompressed() []byte {
	pub := &k.PrivateKey.PublicKey
	if pub == nil || pub.X == nil || pub.Y == nil {
		return nil
	}
	return elliptic.Marshal(pub.Curve, pub.X, pub.Y)
}

func RandomNew() (*Secp256k1Key, error) {
	randBytes := make([]byte, 64)
	_, err := rand.Read(randBytes)
	if err != nil {
		return nil, errors.New("key generation: could not read from random source: " + err.Error())
	}
	reader := bytes.NewReader(randBytes)
	priv, err := ecdsa.GenerateKey(secp256k1.S256(), reader)
	if err != nil {
		return nil, errors.New("key generation: ecdsa.GenerateKey failed: " + err.Error())
	}

	return &Secp256k1Key{PrivateKey: priv}, nil
}

func HexToKey(hexKey string) (*Secp256k1Key, error) {
	if has0xPrefix(hexKey) {
		hexKey = hexKey[2:]
	}
	b, err := hex.DecodeString(hexKey)
	if err != nil {
		return nil, errors.New("invalid hex string")
	}
	return ToKey(b)
}

func has0xPrefix(input string) bool {
	return len(input) >= 2 && input[0] == '0' && (input[1] == 'x' || input[1] == 'X')
}

func ToKey(d []byte) (*Secp256k1Key, error) {
	return toKey(d, true)
}

func toKey(d []byte, strict bool) (*Secp256k1Key, error) {
	priv := new(ecdsa.PrivateKey)
	priv.PublicKey.Curve = secp256k1.S256()
	if strict && 8*len(d) != priv.Params().BitSize {
		return nil, fmt.Errorf("invalid length, need %d bits", priv.Params().BitSize)
	}
	priv.D = new(big.Int).SetBytes(d)

	// The priv.D must < N
	if priv.D.Cmp(secp256k1N) >= 0 {
		return nil, errors.New("invalid private key, >=N")
	}
	// The priv.D must not be zero or negative.
	if priv.D.Sign() <= 0 {
		return nil, errors.New("invalid private key, zero or negative")
	}

	priv.PublicKey.X, priv.PublicKey.Y = priv.PublicKey.Curve.ScalarBaseMult(d)
	if priv.PublicKey.X == nil {
		return nil, errors.New("invalid private key")
	}
	return &Secp256k1Key{PrivateKey: priv}, nil
}
