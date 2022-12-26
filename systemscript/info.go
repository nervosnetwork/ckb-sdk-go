package systemscript

import (
	"github.com/nervosnetwork/ckb-sdk-go/v2/types"
)

type Info struct {
	CodeHash types.Hash
	HashType types.ScriptHashType
	OutPoint *types.OutPoint
	DepType  types.DepType
}

type SystemScript uint

const (
	Secp256k1Blake160SighashAll SystemScript = iota
	Secp256k1Blake160MultisigAll
	AnyoneCanPay
	Dao
	Sudt
	Cheque
	PwLock
	Omnilock
)

var mainnetContracts = make(map[SystemScript]*Info)
var testnetContracts = make(map[SystemScript]*Info)

func init() {
	initMainnetSystemScript()
	initTestnetSystemScript()
}

func initMainnetSystemScript() {
	mainnetContracts[Secp256k1Blake160SighashAll] = &Info{
		CodeHash: types.HexToHash("0x9bd7e06f3ecf4be0f2fcd2188b23f1b9fcc88e5d4b65a8637b17723bbda3cce8"),
		HashType: types.HashTypeType,
		OutPoint: &types.OutPoint{
			TxHash: types.HexToHash("0x71a7ba8fc96349fea0ed3a5c47992e3b4084b031a42264a018e0072e8172e46c"),
			Index:  0,
		},
		DepType: types.DepTypeDepGroup,
	}
	mainnetContracts[Secp256k1Blake160MultisigAll] = &Info{
		CodeHash: types.HexToHash("0x5c5069eb0857efc65e1bca0c07df34c31663b3622fd3876c876320fc9634e2a8"),
		HashType: types.HashTypeType,
		OutPoint: &types.OutPoint{
			TxHash: types.HexToHash("0x71a7ba8fc96349fea0ed3a5c47992e3b4084b031a42264a018e0072e8172e46c"),
			Index:  1,
		},
		DepType: types.DepTypeDepGroup,
	}
	mainnetContracts[Dao] = &Info{
		CodeHash: types.HexToHash("0x82d76d1b75fe2fd9a27dfbaa65a039221a380d76c926f378d3f81cf3e7e13f2e"),
		HashType: types.HashTypeType,
		OutPoint: &types.OutPoint{
			TxHash: types.HexToHash("0xe2fb199810d49a4d8beec56718ba2593b665db9d52299a0f9e6e75416d73ff5c"),
			Index:  2,
		},
		DepType: types.DepTypeCode,
	}
	mainnetContracts[Sudt] = &Info{
		CodeHash: types.HexToHash("0x5e7a36a77e68eecc013dfa2fe6a23f3b6c344b04005808694ae6dd45eea4cfd5"),
		HashType: types.HashTypeType,
		OutPoint: &types.OutPoint{
			TxHash: types.HexToHash("0xc7813f6a415144643970c2e88e0bb6ca6a8edc5dd7c1022746f628284a9936d5"),
			Index:  0,
		},
		DepType: types.DepTypeCode,
	}
	mainnetContracts[Cheque] = &Info{
		CodeHash: types.HexToHash("0xe4d4ecc6e5f9a059bf2f7a82cca292083aebc0c421566a52484fe2ec51a9fb0c"),
		HashType: types.HashTypeType,
		OutPoint: &types.OutPoint{
			TxHash: types.HexToHash("0x04632cc459459cf5c9d384b43dee3e36f542a464bdd4127be7d6618ac6f8d268"),
			Index:  0,
		},
		DepType: types.DepTypeDepGroup,
	}
	mainnetContracts[AnyoneCanPay] = &Info{
		CodeHash: types.HexToHash("0xd369597ff47f29fbc0d47d2e3775370d1250b85140c670e4718af712983a2354"),
		HashType: types.HashTypeType,
		OutPoint: &types.OutPoint{
			TxHash: types.HexToHash("0x4153a2014952d7cac45f285ce9a7c5c0c0e1b21f2d378b82ac1433cb11c25c4d"),
			Index:  0,
		},
		DepType: types.DepTypeDepGroup,
	}
	mainnetContracts[PwLock] = &Info{
		CodeHash: types.HexToHash("0xbf43c3602455798c1a61a596e0d95278864c552fafe231c063b3fabf97a8febc"),
		HashType: types.HashTypeType,
		OutPoint: &types.OutPoint{
			TxHash: types.HexToHash("0x1d60cb8f4666e039f418ea94730b1a8c5aa0bf2f7781474406387462924d15d4"),
			Index:  0,
		},
		DepType: types.DepTypeCode,
	}
	mainnetContracts[Omnilock] = &Info{
		CodeHash: types.HexToHash("0x9b819793a64463aed77c615d6cb226eea5487ccfc0783043a587254cda2b6f26"),
		HashType: types.HashTypeType,
		OutPoint: &types.OutPoint{
			TxHash: types.HexToHash("0xdfdb40f5d229536915f2d5403c66047e162e25dedd70a79ef5164356e1facdc8"),
			Index:  0,
		},
		DepType: types.DepTypeCode,
	}
}

func initTestnetSystemScript() {
	testnetContracts[Secp256k1Blake160SighashAll] = &Info{
		CodeHash: types.HexToHash("0x9bd7e06f3ecf4be0f2fcd2188b23f1b9fcc88e5d4b65a8637b17723bbda3cce8"),
		HashType: types.HashTypeType,
		OutPoint: &types.OutPoint{
			TxHash: types.HexToHash("0xf8de3bb47d055cdf460d93a2a6e1b05f7432f9777c8c474abf4eec1d4aee5d37"),
			Index:  0,
		},
		DepType: types.DepTypeDepGroup,
	}
	testnetContracts[Secp256k1Blake160MultisigAll] = &Info{
		CodeHash: types.HexToHash("0x5c5069eb0857efc65e1bca0c07df34c31663b3622fd3876c876320fc9634e2a8"),
		HashType: types.HashTypeType,
		OutPoint: &types.OutPoint{
			TxHash: types.HexToHash("0xf8de3bb47d055cdf460d93a2a6e1b05f7432f9777c8c474abf4eec1d4aee5d37"),
			Index:  1,
		},
		DepType: types.DepTypeDepGroup,
	}
	testnetContracts[Dao] = &Info{
		CodeHash: types.HexToHash("0x82d76d1b75fe2fd9a27dfbaa65a039221a380d76c926f378d3f81cf3e7e13f2e"),
		HashType: types.HashTypeType,
		OutPoint: &types.OutPoint{
			TxHash: types.HexToHash("0x8f8c79eb6671709633fe6a46de93c0fedc9c1b8a6527a18d3983879542635c9f"),
			Index:  2,
		},
		DepType: types.DepTypeCode,
	}
	testnetContracts[Sudt] = &Info{
		CodeHash: types.HexToHash("0xc5e5dcf215925f7ef4dfaf5f4b4f105bc321c02776d6e7d52a1db3fcd9d011a4"),
		HashType: types.HashTypeType,
		OutPoint: &types.OutPoint{
			TxHash: types.HexToHash("0xe12877ebd2c3c364dc46c5c992bcfaf4fee33fa13eebdf82c591fc9825aab769"),
			Index:  0,
		},
		DepType: types.DepTypeCode,
	}
	testnetContracts[Cheque] = &Info{
		CodeHash: types.HexToHash("0x60d5f39efce409c587cb9ea359cefdead650ca128f0bd9cb3855348f98c70d5b"),
		HashType: types.HashTypeType,
		OutPoint: &types.OutPoint{
			TxHash: types.HexToHash("0x7f96858be0a9d584b4a9ea190e0420835156a6010a5fde15ffcdc9d9c721ccab"),
			Index:  0,
		},
		DepType: types.DepTypeDepGroup,
	}
	testnetContracts[AnyoneCanPay] = &Info{
		CodeHash: types.HexToHash("0x3419a1c09eb2567f6552ee7a8ecffd64155cffe0f1796e6e61ec088d740c1356"),
		HashType: types.HashTypeType,
		OutPoint: &types.OutPoint{
			TxHash: types.HexToHash("0xec26b0f85ed839ece5f11c4c4e837ec359f5adc4420410f6453b1f6b60fb96a6"),
			Index:  0,
		},
		DepType: types.DepTypeDepGroup,
	}
	testnetContracts[PwLock] = &Info{
		CodeHash: types.HexToHash("0x58c5f491aba6d61678b7cf7edf4910b1f5e00ec0cde2f42e0abb4fd9aff25a63"),
		HashType: types.HashTypeType,
		OutPoint: &types.OutPoint{
			TxHash: types.HexToHash("0x57a62003daeab9d54aa29b944fc3b451213a5ebdf2e232216a3cfed0dde61b38"),
			Index:  0,
		},
		DepType: types.DepTypeCode,
	}
	testnetContracts[Omnilock] = &Info{
		CodeHash: types.HexToHash("0xf329effd1c475a2978453c8600e1eaf0bc2087ee093c3ee64cc96ec6847752cb"),
		HashType: types.HashTypeType,
		OutPoint: &types.OutPoint{
			TxHash: types.HexToHash("0x27b62d8be8ed80b9f56ee0fe41355becdb6f6a40aeba82d3900434f43b1c8b60"),
			Index:  0,
		},
		DepType: types.DepTypeCode,
	}
}

func GetInfo(network types.Network, script SystemScript) *Info {
	switch network {
	case types.NetworkMain:
		return mainnetContracts[script]
	case types.NetworkTest:
		return testnetContracts[script]
	default:
		return nil
	}
}

func GetCodeHash(network types.Network, script SystemScript) types.Hash {
	return GetInfo(network, script).CodeHash
}

func NewScript(script SystemScript, args []byte, network types.Network) *types.Script {
	systemScriptInfo := GetInfo(network, script)
	return &types.Script{
		CodeHash: systemScriptInfo.CodeHash,
		HashType: systemScriptInfo.HashType,
		Args:     args,
	}
}
