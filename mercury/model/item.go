package model

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
		ItemTypeAddress,
		addr,
	}, nil
}

func NewIdentityItemByPublicKeyHash(publicKeyHash string) (*Item, error) {
	hash, err := hexutil.Decode(publicKeyHash)
	if err != nil {
		return nil, err
	}
	identity, err := toIdentity(IdentityFlagsCkb, hash[:20])
	if err != nil {
		return nil, err
	}
	return &Item{ItemTypeIdentity, identity}, nil
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
	return &Item{ItemTypeIdentity, identity}, nil
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

	return &Item{ItemTypeIdentity, identity}, nil
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
		ItemTypeOutPoint,
		outPoint,
	}
}

type Item struct {
	Type  ItemType    `json:"type"`
	Value interface{} `json:"value"`
}

type ItemType string

const (
	ItemTypeAddress  ItemType = "Address"
	ItemTypeIdentity          = "Identity"
	ItemTypeOutPoint          = "OutPoint"
)