package req

import (
	"errors"
	"github.com/ethereum/go-ethereum/common/hexutil"
	address2 "github.com/nervosnetwork/ckb-sdk-go/address"
	"github.com/nervosnetwork/ckb-sdk-go/types"
)

const (
	IdentityFlagsCkb byte = 0x00
)

func NewAddressItem(addr string) (*Item, error) {
	return &Item{
		ItemAddress,
		addr,
	}, nil
}

func NewIdentityItemByCkb(publicKeyHash string) (*Item, error) {
	hash, err := hexutil.Decode(publicKeyHash)
	if err != nil {
		return nil, err
	}
	identity, err := toIdentity(IdentityFlagsCkb, hash)
	if err != nil {
		return nil, err
	}
	return &Item{ItemIdentity, identity}, nil
}

func NewIdentityItemByAddress(address string) (*Item, error) {
	// TODO: check address type
	parse, err := address2.Parse(address)
	if err != nil {
		return nil, err
	}
	identity, err := toIdentity(IdentityFlagsCkb, parse.Script.Args)
	if err != nil {
		return nil, err
	}

	return &Item{ItemIdentity, identity}, nil
}

func toIdentity(flag byte, content []byte) (string, error) {
	if len(content) != 20 {
		return "", errors.New("identity content should be 20 bytes length")
	}
	identity := append([]byte{flag}, content...)
	return hexutil.Encode(identity), nil
}

func NewOutPointItem(txHash types.Hash, index uint) *Item {
	outPoint := types.OutPoint{TxHash: txHash, Index: index}
	return &Item{
		ItemOutPoint,
		outPoint,
	}
}

type Item struct {
	Type  ItemType    `json:"type"`
	Value interface{} `json:"value"`
}

type ItemType string

const (
	ItemAddress  ItemType = "Address"
	ItemIdentity          = "Identity"
	ItemOutPoint          = "OutPoint"
)
