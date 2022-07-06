package types

import (
	"reflect"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type Network uint
type BuiltinScript uint

const (
	HashLength = 32

	NetworkMain Network = iota
	NetworkTest

	BuiltinScriptSecp256k1Blake160SighashAll BuiltinScript = iota
	BuiltinScriptSecp256k1Blake160MultisigAll
	BuiltinScriptAnyoneCanPay
	BuiltinScriptDao
	BuiltinScriptSUDT
	BuiltinScriptCheque
	BuiltinScriptPWLock
)

//var CodeHashSecp256k1Blake160SighashAllCodeHash = CodeHash(HexToHash("0x9bd7e06f3ecf4be0f2fcd2188b23f1b9fcc88e5d4b65a8637b17723bbda3cce8"))

var (
	hashT = reflect.TypeOf(Hash{})
)

func GetCodeHash(script BuiltinScript, network Network) Hash {
	if network == NetworkMain {
		switch script {
		case BuiltinScriptSecp256k1Blake160SighashAll:
			return HexToHash("0x9bd7e06f3ecf4be0f2fcd2188b23f1b9fcc88e5d4b65a8637b17723bbda3cce8")
		case BuiltinScriptSecp256k1Blake160MultisigAll:
			return HexToHash("0x5c5069eb0857efc65e1bca0c07df34c31663b3622fd3876c876320fc9634e2a8")
		case BuiltinScriptAnyoneCanPay:
			return HexToHash("0xd369597ff47f29fbc0d47d2e3775370d1250b85140c670e4718af712983a2354")
		case BuiltinScriptDao:
			return HexToHash("0x82d76d1b75fe2fd9a27dfbaa65a039221a380d76c926f378d3f81cf3e7e13f2e")
		case BuiltinScriptSUDT:
			return HexToHash("0x5e7a36a77e68eecc013dfa2fe6a23f3b6c344b04005808694ae6dd45eea4cfd5")
		case BuiltinScriptCheque:
			return HexToHash("0xe4d4ecc6e5f9a059bf2f7a82cca292083aebc0c421566a52484fe2ec51a9fb0c")
		case BuiltinScriptPWLock:
			return HexToHash("0xbf43c3602455798c1a61a596e0d95278864c552fafe231c063b3fabf97a8febc")
		}
	} else if network == NetworkTest {
		switch script {
		case BuiltinScriptSecp256k1Blake160SighashAll:
			return HexToHash("0x9bd7e06f3ecf4be0f2fcd2188b23f1b9fcc88e5d4b65a8637b17723bbda3cce8")
		case BuiltinScriptSecp256k1Blake160MultisigAll:
			return HexToHash("0x5c5069eb0857efc65e1bca0c07df34c31663b3622fd3876c876320fc9634e2a8")
		case BuiltinScriptAnyoneCanPay:
			return HexToHash("0x3419a1c09eb2567f6552ee7a8ecffd64155cffe0f1796e6e61ec088d740c1356")
		case BuiltinScriptDao:
			return HexToHash("0x82d76d1b75fe2fd9a27dfbaa65a039221a380d76c926f378d3f81cf3e7e13f2e")
		case BuiltinScriptSUDT:
			return HexToHash("0xc5e5dcf215925f7ef4dfaf5f4b4f105bc321c02776d6e7d52a1db3fcd9d011a4")
		case BuiltinScriptCheque:
			return HexToHash("0x60d5f39efce409c587cb9ea359cefdead650ca128f0bd9cb3855348f98c70d5b")
		case BuiltinScriptPWLock:
			return HexToHash("0x58c5f491aba6d61678b7cf7edf4910b1f5e00ec0cde2f42e0abb4fd9aff25a63")
		}
	}
	return Hash{}
}

type Hash [HashLength]byte

func BytesToHash(b []byte) Hash {
	var h Hash
	h.SetBytes(b)
	return h
}

func HexToHash(s string) Hash {
	return BytesToHash(common.FromHex(s))
}

func (h *Hash) SetBytes(b []byte) {
	if len(b) > len(h) {
		b = b[len(b)-HashLength:]
	}

	copy(h[HashLength-len(b):], b)
}

func (h Hash) Bytes() []byte {
	return h[:]
}

func (h Hash) Hex() string {
	return hexutil.Encode(h[:])
}

func (h Hash) String() string {
	return h.Hex()
}

func (h *Hash) UnmarshalText(input []byte) error {
	return hexutil.UnmarshalFixedText("Hash", input, h[:])
}

func (h *Hash) UnmarshalJSON(input []byte) error {
	return hexutil.UnmarshalFixedJSON(hashT, input, h[:])
}

func (h Hash) MarshalText() ([]byte, error) {
	return hexutil.Bytes(h[:]).MarshalText()
}
