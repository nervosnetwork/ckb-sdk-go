package signer

import (
	"github.com/nervosnetwork/ckb-sdk-go/v3/systemscript"
	"github.com/nervosnetwork/ckb-sdk-go/v3/types"
)

func init() {
	networks := []types.Network{types.NetworkMain, types.NetworkTest}
	for _, network := range networks {
		instance := GetTransactionSignerInstance(network)
		instance.RegisterLockSigner(
			systemscript.GetCodeHash(network, systemscript.Secp256k1Blake160SighashAll), &Secp256k1Blake160SighashAllSigner{})
		instance.RegisterLockSigner(
			systemscript.GetCodeHash(network, systemscript.Secp256k1Blake160MultisigAllLegacy), &Secp256k1Blake160MultisigAllSigner{})
		instance.RegisterLockSigner(
			systemscript.GetCodeHash(network, systemscript.Secp256k1Blake160MultisigAllV2), &Secp256k1Blake160MultisigAllSigner{})
		instance.RegisterLockSigner(
			systemscript.GetCodeHash(network, systemscript.AnyoneCanPay), &AnyCanPaySigner{})
		instance.RegisterLockSigner(
			systemscript.GetCodeHash(network, systemscript.PwLock), &PWLockSigner{})
		instance.RegisterLockSigner(
			systemscript.GetCodeHash(network, systemscript.Omnilock), &OmnilockSigner{})
	}
}
