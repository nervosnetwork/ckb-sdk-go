package signer

import (
	"github.com/nervosnetwork/ckb-sdk-go/script"
	"github.com/nervosnetwork/ckb-sdk-go/types"
)

func init() {
	ns := []types.Network{types.NetworkMain, types.NetworkTest}
	for _, network := range ns {
		instance := GetTransactionSignerInstance(network)
		instance.RegisterLockSigner(
			script.GetCodeHash(network, script.SystemScriptSecp256k1Blake160SighashAll), &Secp256k1Blake160SighashAllSigner{})
		instance.RegisterLockSigner(
			script.GetCodeHash(network, script.SystemScriptSecp256k1Blake160MultisigAll), &Secp256k1Blake160MultisigAllSigner{})
		instance.RegisterLockSigner(
			script.GetCodeHash(network, script.SystemScriptAnyoneCanPay), &AnyCanPaySigner{})
		instance.RegisterLockSigner(
			script.GetCodeHash(network, script.SystemScriptPwLock), &PWLockSigner{})
	}
}
