package signer

import (
	"github.com/nervosnetwork/ckb-sdk-go/systemscript"
	"github.com/nervosnetwork/ckb-sdk-go/types"
)

func init() {
	networks := []types.Network{types.NetworkMain, types.NetworkTest}
	for _, network := range networks {
		instance := GetTransactionSignerInstance(network)
		instance.RegisterLockSigner(
			systemscript.GetCodeHash(network, systemscript.Secp256k1Blake160SighashAll), &Secp256k1Blake160SighashAllSigner{})
		instance.RegisterLockSigner(
			systemscript.GetCodeHash(network, systemscript.Secp256k1Blake160MultisigAll), &Secp256k1Blake160MultisigAllSigner{})
		instance.RegisterLockSigner(
			systemscript.GetCodeHash(network, systemscript.AnyoneCanPay), &AnyCanPaySigner{})
		instance.RegisterLockSigner(
			systemscript.GetCodeHash(network, systemscript.PwLock), &PWLockSigner{})
	}
}
