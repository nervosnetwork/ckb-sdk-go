package model

import (
	"github.com/nervosnetwork/ckb-sdk-go/crypto"
	"github.com/nervosnetwork/ckb-sdk-go/types"
)

type TransferCompletionResponse struct {
	TxView       *types.Transaction `json:"tx_view"`
	ScriptGroups []*ScriptGroup     `json:"script_groups"`
}

type ScriptGroup struct {
	Script        types.Script `json:"script"`
	GroupType     GroupType    `json:"group_type"`
	InputIndices  []uint32     `json:"input_indices"`
	OutputIndices []uint32     `json:"output_indices"`
}

type GroupType string

const (
	GroupTypeLock GroupType = "Lock"
	GroupTypeType GroupType = "Type"
)

func SignTransaction(transaction *types.Transaction, scriptGroup *ScriptGroup, privateKey crypto.Key) error {
	if isPWLock(scriptGroup) {
		return ethereumPersonalKeccakSign(transaction, scriptGroup, privateKey)
	} else {
		return secp256Blake2bSign(transaction, scriptGroup, privateKey)
	}
}

func secp256Blake2bSign(transaction *types.Transaction, scriptGroup *ScriptGroup, privateKey crypto.Key) error {
	// TODO: remove
	//witnessBytes := scriptGroup.GetWitness()
	//groupWitnesses := scriptGroup.GetGroupWitnesses()
	//
	//txHash, err := transaction.ComputeHash()
	//if err != nil {
	//	return err
	//}
	//
	//length := make([]byte, 8)
	//binary.LittleEndian.PutUint64(length, uint64(len(witnessBytes)))
	//
	//message := txHash.Bytes()
	//message = append(message, length...)
	//message = append(message, witnessBytes...)
	//
	//for i := 1; i < len(groupWitnesses); i++ {
	//	witnessBytes := groupWitnesses[i]
	//	length := make([]byte, 8)
	//	binary.LittleEndian.PutUint64(length, uint64(len(witnessBytes)))
	//	message = append(message, length...)
	//	message = append(message, witnessBytes...)
	//}
	//
	//hash, err := blake2b.Blake256(message)
	//if err != nil {
	//	return err
	//}
	//
	//signature, err := privateKey.Sign(hash)
	//if err != nil {
	//	return err
	//}
	//
	//newWitness := scriptGroup.GetWitness()
	//offset := scriptGroup.GetOffSet()
	//for i := 0; i < len(signature); i++ {
	//	newWitness[i+offset] = signature[i]
	//}
	//transaction.Witnesses[scriptGroup.GetWitnessIndex()] = newWitness
	return nil
}

func ethereumPersonalKeccakSign(transaction *types.Transaction, scriptGroup *ScriptGroup, privateKey crypto.Key) error {
	// TODO: remove
	//witnessBytes := scriptGroup.GetWitness()
	//groupWitnesses := scriptGroup.GetGroupWitnesses()
	//
	//txHash, err := transaction.ComputeHash()
	//if err != nil {
	//	return err
	//}
	//
	//length := make([]byte, 8)
	//binary.LittleEndian.PutUint64(length, uint64(len(witnessBytes)))
	//
	//message := txHash.Bytes()
	//message = append(message, length...)
	//message = append(message, witnessBytes...)
	//
	//for i := 1; i < len(groupWitnesses); i++ {
	//	witnessBytes := groupWitnesses[i]
	//	length := make([]byte, 8)
	//	binary.LittleEndian.PutUint64(length, uint64(len(witnessBytes)))
	//	message = append(message, length...)
	//	message = append(message, witnessBytes...)
	//}
	//
	//hash, err := keccak256.Keccak256(message)
	//if err != nil {
	//	return err
	//}
	//
	//prefix := []byte("\u0019Ethereum Signed Message:\n" + strconv.Itoa(len(hash)))
	//
	//message = append(prefix, hash...)
	//hash, err = keccak256.Keccak256(message)
	//
	//if err != nil {
	//	return err
	//}
	//
	//signature, err := privateKey.Sign(hash)
	//if err != nil {
	//	return err
	//}
	//
	//newWitness := scriptGroup.GetWitness()
	//offset := scriptGroup.GetOffSet()
	//for i := 0; i < len(signature); i++ {
	//	newWitness[i+offset] = signature[i]
	//}
	//transaction.Witnesses[scriptGroup.GetWitnessIndex()] = newWitness
	return nil
}

func isPWLock(g *ScriptGroup) bool {
	//if Keccak256 == *g.Action.HashAlgorithm && EthereumPersonal == g.Action.SignatureInfo.Algorithm {
	//	return true
	//}
	return false
}
