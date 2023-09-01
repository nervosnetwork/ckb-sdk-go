package transaction

import (
	"github.com/pkg/errors"

	"github.com/nervosnetwork/ckb-sdk-go/v2/crypto/secp256k1"
)

// Context is user provided information for
// `signer.TransactionSigner.SignTransaction`. In turn the context is passed to `signer.ScriptSigner.SignTransaction`.
//
// See more in github.com/nervosnetwork/ckb-sdk-go/v2/transaction/signer
type Context struct {
	Key     *secp256k1.Secp256k1Key
	Payload interface{}
}

func NewContext(ecPrivateKey string) (*Context, error) {
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
