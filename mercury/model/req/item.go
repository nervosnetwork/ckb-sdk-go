package req

import (
	"bytes"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	address2 "github.com/nervosnetwork/ckb-sdk-go/address"
	"github.com/nervosnetwork/ckb-sdk-go/types"
	"unicode/utf8"
)

const (
	IDENTITY_FLAGS_CKB = "0x00"
)

type address struct {
	Address string `json:"Address"`
}

func NewAddressItem(addr string) (*address, error) {
	return &address{addr}, nil
}

type identity struct {
	Identity string `json:"Identity"`
}

func NewIdentityItemByCkb(pubKey string) (*identity, error) {
	return &identity{
		toIdentity(IDENTITY_FLAGS_CKB, common.FromHex(pubKey)),
	}, nil
}

func NewIdentityItemByAddress(address string) (*identity, error) {
	parse, err := address2.Parse(address)
	if err != nil {
		return nil, err
	}

	return &identity{
		toIdentity(IDENTITY_FLAGS_CKB, parse.Script.Args),
	}, nil

}

func toIdentity(flag string, pubKey []byte) string {
	byteArr := make([][]byte, 2)
	byteArr[0] = common.FromHex(flag)
	byteArr[1] = pubKey

	return hexutil.Encode(bytes.Join(byteArr, []byte("")))
}

type record struct {
	Record string `json:"Record"`
}

func NewRecordItemByScript(point *types.OutPoint, script *types.Script) (*record, error) {
	hash, err := script.Hash()
	if err != nil {
		return nil, err
	}

	byteArr := make([][]byte, 4)
	byteArr[0] = point.TxHash[:]
	byteArr[1] = intToByteArray(point.Index)
	byteArr[2] = common.FromHex("0x01")
	byteArr[3] = runesToUTF8Manual([]rune(hexutil.Encode(hash.Bytes())[2:42]))

	return &record{
		hexutil.Encode(bytes.Join(byteArr, []byte(""))),
	}, nil
}

func NewRecordItemByAddress(point *types.OutPoint, address string) (*record, error) {

	byteArr := make([][]byte, 4)
	byteArr[0] = point.TxHash[:]
	byteArr[1] = intToByteArray(point.Index)
	byteArr[2] = common.FromHex("0x00")
	byteArr[3] = []byte(address)

	return &record{
		hexutil.Encode(bytes.Join(byteArr, []byte(""))),
	}, nil
}

func intToByteArray(num uint) []byte {
	byteArr := make([]byte, 4)
	byteArr[3] = (byte)(num & 0xFF)
	byteArr[2] = (byte)(num & 0xFF00)
	byteArr[1] = (byte)(num & 0xFF0000)
	byteArr[0] = (byte)(num & 0xFF000000)

	return byteArr
}

func runesToUTF8Manual(rs []rune) []byte {
	size := 0
	for _, r := range rs {
		size += utf8.RuneLen(r)
	}

	bs := make([]byte, size)

	count := 0
	for _, r := range rs {
		count += utf8.EncodeRune(bs[count:], r)
	}

	return bs
}
