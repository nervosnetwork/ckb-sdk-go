package transaction

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/nervosnetwork/ckb-sdk-go/crypto"
	"github.com/nervosnetwork/ckb-sdk-go/types"
)

type TransactionWithScriptGroups struct {
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

func (r *ScriptGroup) UnmarshalJSON(input []byte) error {
	var jsonObj struct {
		Script        types.Script   `json:"script"`
		GroupType     GroupType      `json:"group_type"`
		InputIndices  []hexutil.Uint `json:"input_indices"`
		OutputIndices []hexutil.Uint `json:"output_indices"`
	}
	if err := json.Unmarshal(input, &jsonObj); err != nil {
		return err
	}
	toUint32Array := func(a []hexutil.Uint) []uint32 {
		result := make([]uint32, len(a))
		for i, data := range a {
			result[i] = uint32(data)
		}
		return result
	}
	*r = ScriptGroup{
		Script:        jsonObj.Script,
		GroupType:     jsonObj.GroupType,
		InputIndices:  toUint32Array(jsonObj.InputIndices),
		OutputIndices: toUint32Array(jsonObj.OutputIndices),
	}
	return nil
}

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
