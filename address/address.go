package address

import (
	"github.com/nervosnetwork/ckb-sdk-go/utils"
	"github.com/pkg/errors"
)

func ValidateChequeAddress(addr string, systemScripts *utils.SystemScripts) (*Address, error) {
	address, err := Decode(addr)
	if err != nil {
		return nil, err
	}
	if isSecp256k1Lock(address, systemScripts) {
		return address, nil
	}
	return nil, errors.Errorf("address %s is not an SECP256K1 short format address", addr)
}

func isSecp256k1Lock(parsedSenderAddr *Address, systemScripts *utils.SystemScripts) bool {
	return parsedSenderAddr.Script.CodeHash == systemScripts.SecpSingleSigCell.CellHash &&
		parsedSenderAddr.Script.HashType == systemScripts.SecpSingleSigCell.HashType &&
		len(parsedSenderAddr.Script.Args) == 20
}
