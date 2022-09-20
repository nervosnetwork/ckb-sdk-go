package utils

import "github.com/nervosnetwork/ckb-sdk-go/types"

func ChequeCellArgs(senderLock, receiverLock *types.Script) []byte {
	senderLockHash := senderLock.Hash()
	receiverLockHash := receiverLock.Hash()
	return append(receiverLockHash.Bytes()[0:20], senderLockHash.Bytes()[0:20]...)
}
