package types

import "context"

type GenericRPCClient interface {
	CallContext(ctx context.Context, result interface{}, method string, args ...interface{}) error
	Close()
}
