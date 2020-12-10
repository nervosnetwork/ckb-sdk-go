package utils

import "github.com/nervosnetwork/ckb-sdk-go/types"

func ChequeCellArgs(senderLock, receiverLock *types.Script) ([]byte, error) {
	senderLockHash, err := senderLock.Hash()
	if err != nil {
		return []byte{}, err
	}
	return append(receiverLock.Args, senderLockHash.Bytes()[0:20]...), nil
}
