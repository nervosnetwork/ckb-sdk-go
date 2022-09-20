package transaction

import (
	"encoding/binary"
	"fmt"
	"github.com/nervosnetwork/ckb-sdk-go/crypto"
	"github.com/nervosnetwork/ckb-sdk-go/crypto/blake2b"
	"github.com/nervosnetwork/ckb-sdk-go/types"
)

var (
	Secp256k1EmptyWitnessArg            = NewEmptyWitnessArg(65)
	Secp256k1EmptyWitnessArgPlaceholder = make([]byte, 85)
	Secp256k1SignaturePlaceholder       = make([]byte, 65)
)

func NewEmptyWitnessArg(LockScriptLength uint) *types.WitnessArgs {
	return &types.WitnessArgs{
		Lock:       make([]byte, LockScriptLength),
		InputType:  nil,
		OutputType: nil,
	}
}

// SignTransaction signs transaction with index group and witness placeholder in secp256k1_blake160_sighash_all way
func SignTransaction(tx *types.Transaction, group []int, witnessPlaceholder []byte, key crypto.Key) ([]byte, error) {
	inputsLen := len(tx.Inputs)
	for i := 0; i < len(group); i++ {
		if i > 0 && group[i] <= group[i-1] {
			return nil, fmt.Errorf("group index is not in ascending order")
		}
		if group[i] > inputsLen {
			return nil, fmt.Errorf("group index %d is greater than input bytesLen %d", group[i], inputsLen)
		}
	}
	txHash := tx.ComputeHash()
	msg := txHash.Bytes()
	bytesLen := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytesLen, uint64(len(witnessPlaceholder)))
	msg = append(msg, bytesLen...)
	msg = append(msg, witnessPlaceholder...)

	var indexes []int
	for i := 1; i < len(group); i++ {
		indexes = append(indexes, group[i])
	}
	for i := inputsLen; i < len(tx.Witnesses); i++ {
		indexes = append(indexes, i)
	}
	for _, i := range indexes {
		bytes := tx.Witnesses[i]
		bytesLen := make([]byte, 8)
		binary.LittleEndian.PutUint64(bytesLen, uint64(len(bytes)))
		msg = append(msg, bytesLen...)
		msg = append(msg, bytes...)
	}

	msgHash := blake2b.Blake256(msg)
	signature, err := key.Sign(msgHash)
	if err != nil {
		return nil, err
	}
	return signature, nil
}

func MsgFromTxForMultiSig(transaction *types.Transaction, group []int, multisigScript []byte) ([]byte, error) {
	var witnessArgs *types.WitnessArgs
	var err error
	originalWitness := transaction.Witnesses[group[0]]
	if originalWitness == nil || len(originalWitness) == 0 {
		witnessArgs = &types.WitnessArgs{}
	} else {
		witnessArgs, err = types.DeserializeWitnessArgs(originalWitness)
		if err != nil {
			return nil, err
		}
	}
	n := int(multisigScript[3])
	witnessArgsLock := append(multisigScript, make([]byte, n*len(Secp256k1SignaturePlaceholder))...)
	witnessArgs.Lock = witnessArgsLock
	data := witnessArgs.Serialize()
	length := make([]byte, 8)
	binary.LittleEndian.PutUint64(length, uint64(len(data)))

	hash := transaction.ComputeHash()

	message := append(hash.Bytes(), length...)
	message = append(message, data...)

	if len(group) > 1 {
		for i := 1; i < len(group); i++ {
			data = transaction.Witnesses[group[i]]
			length := make([]byte, 8)
			binary.LittleEndian.PutUint64(length, uint64(len(data)))
			message = append(message, length...)
			message = append(message, data...)
		}
	}
	// hash witnesses that are not in any input group
	for _, witness := range transaction.Witnesses[len(transaction.Inputs):] {
		length := make([]byte, 8)
		binary.LittleEndian.PutUint64(length, uint64(len(witness)))
		message = append(message, length...)
		message = append(message, witness...)
	}

	return blake2b.Blake256(message), nil
}

func MultiSignTransaction(transaction *types.Transaction, group []int, witnessArgs *types.WitnessArgs, serialize []byte, signatures ...[]byte) error {
	var signed []byte
	for _, sig := range signatures {
		signed = append(signed, sig...)
	}

	wa := &types.WitnessArgs{
		Lock:       append(serialize, signed...),
		InputType:  witnessArgs.InputType,
		OutputType: witnessArgs.OutputType,
	}
	wab := wa.Serialize()
	transaction.Witnesses[group[0]] = wab

	return nil
}

func SingleSegmentSignMessage(transaction *types.Transaction, start int, end int, witnessArgs *types.WitnessArgs) ([]byte, error) {
	data := witnessArgs.Serialize()
	length := make([]byte, 8)
	binary.LittleEndian.PutUint64(length, uint64(len(data)))

	hash := transaction.ComputeHash()

	message := append(hash.Bytes(), length...)
	message = append(message, data...)

	for i := start + 1; i < end; i++ {
		var data []byte
		length := make([]byte, 8)
		binary.LittleEndian.PutUint64(length, uint64(len(data)))
		message = append(message, length...)
		message = append(message, data...)
	}

	return blake2b.Blake256(message), nil
}

func SingleSegmentSignTransaction(transaction *types.Transaction, start int, end int, witnessArgs *types.WitnessArgs, key crypto.Key) error {
	message, err := SingleSegmentSignMessage(transaction, start, end, witnessArgs)
	if err != nil {
		return err
	}

	signed, err := key.Sign(message)
	if err != nil {
		return err
	}

	wa := &types.WitnessArgs{
		Lock:       signed,
		InputType:  witnessArgs.InputType,
		OutputType: witnessArgs.OutputType,
	}
	wab := wa.Serialize()
	transaction.Witnesses[start] = wab

	return nil
}

func CalculateTransactionFee(tx *types.Transaction, feeRate uint64) uint64 {
	txSize := tx.SizeInBlock()
	fee := txSize * feeRate / 1000
	if fee*1000 < txSize*feeRate {
		fee += 1
	}
	return fee
}
