package transaction

import (
	"github.com/nervosnetwork/ckb-sdk-go/crypto/secp256k1"
	"github.com/pkg/errors"
	"strings"
)

type Context struct {
	Key     *secp256k1.Secp256k1Key
	Payload interface{}
}

func NewContext(ecPrivateKey string) (*Context, error) {
	if strings.HasPrefix(ecPrivateKey, "0x") {
		ecPrivateKey = ecPrivateKey[2:]
	}
	key, err := secp256k1.HexToKey(ecPrivateKey)
	if err != nil {
		return nil, errors.WithMessage(err, ecPrivateKey)
	}
	return &Context{
		Key:     key,
		Payload: nil,
	}, nil
}

func NewContextWithPayload(ecPrivateKey string, payload interface{}) (*Context, error) {
	context, err := NewContext(ecPrivateKey)
	if err != nil {
		return nil, err
	}
	context.Payload = payload
	return context, nil
}

type Contexts []*Context

func NewContexts() Contexts {
	c := make([]*Context, 0)
	return c
}

func (r *Contexts) AddByPrivateKeys(ecPrivateKeys ...string) error {
	for _, key := range ecPrivateKeys {
		context, err := NewContext(key)
		if err != nil {
			return err
		}
		r.Add(context)
	}
	return nil
}

func (r *Contexts) Add(context *Context) bool {
	if context == nil {
		return false
	}
	*r = append(*r, context)
	return true
}
