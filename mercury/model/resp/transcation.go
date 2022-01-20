package resp

import (
	"encoding/binary"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/nervosnetwork/ckb-sdk-go/crypto"
	"github.com/nervosnetwork/ckb-sdk-go/crypto/blake2b"
	"github.com/nervosnetwork/ckb-sdk-go/crypto/ethereum"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/common"
	"github.com/nervosnetwork/ckb-sdk-go/types"
	"strconv"
)

type TransferCompletionResponse struct {
	TxView           *transactionResp   `json:"tx_view"`
	SignatureActions []*SignatureAction `json:"signature_actions"`
}

type ScriptGroup struct {
	Action          *SignatureAction
	Transaction     *transactionResp
	OriginalWitness []byte
}

func (self *TransferCompletionResponse) GetTransaction() *types.Transaction {
	return toTransaction(self.TxView)
}

func (self *TransferCompletionResponse) GetScriptGroup() []*ScriptGroup {
	scriptGroups := make([]*ScriptGroup, len(self.SignatureActions))
	for i, v := range self.SignatureActions {
		scriptGroups[i] = NewScriptGroup(v, self.TxView)
	}
	return scriptGroups
}

func toTransaction(tx *transactionResp) *types.Transaction {
	return &types.Transaction{
		Version:     uint(tx.Version),
		Hash:        tx.Hash,
		CellDeps:    toCellDeps(tx.CellDeps),
		HeaderDeps:  tx.HeaderDeps,
		Inputs:      toInputs(tx.Inputs),
		Outputs:     toOutputs(tx.Outputs),
		OutputsData: toBytesArray(tx.OutputsData),
		Witnesses:   toBytesArray(tx.Witnesses),
	}
}

func toCellDeps(deps []common.CellDep) []*types.CellDep {
	result := make([]*types.CellDep, len(deps))
	for i := 0; i < len(deps); i++ {
		dep := deps[i]
		result[i] = &types.CellDep{
			OutPoint: &types.OutPoint{
				TxHash: dep.OutPoint.TxHash,
				Index:  uint(dep.OutPoint.Index),
			},
			DepType: dep.DepType,
		}
	}
	return result
}

func toInputs(inputs []common.CellInput) []*types.CellInput {
	result := make([]*types.CellInput, len(inputs))
	for i := 0; i < len(inputs); i++ {
		input := inputs[i]
		result[i] = &types.CellInput{
			Since: uint64(input.Since),
			PreviousOutput: &types.OutPoint{
				TxHash: input.PreviousOutput.TxHash,
				Index:  uint(input.PreviousOutput.Index),
			},
		}
	}
	return result
}

func toOutputs(outputs []common.CellOutput) []*types.CellOutput {
	result := make([]*types.CellOutput, len(outputs))
	for i := 0; i < len(outputs); i++ {
		output := outputs[i]
		result[i] = &types.CellOutput{
			Capacity: uint64(output.Capacity),
			Lock: &types.Script{
				CodeHash: output.Lock.CodeHash,
				HashType: output.Lock.HashType,
				Args:     output.Lock.Args,
			},
		}
		if output.Type != nil {
			result[i].Type = &types.Script{
				CodeHash: output.Type.CodeHash,
				HashType: output.Type.HashType,
				Args:     output.Type.Args,
			}
		}
	}
	return result
}

func toBytesArray(bytes []hexutil.Bytes) [][]byte {
	result := make([][]byte, len(bytes))
	for i, data := range bytes {
		result[i] = data
	}
	return result
}

func NewScriptGroup(action *SignatureAction, transaction *transactionResp) *ScriptGroup {
	return &ScriptGroup{
		Action:          action,
		Transaction:     transaction,
		OriginalWitness: transaction.Witnesses[action.SignatureLocation.Index],
	}
}

func (g *ScriptGroup) GetOffSet() int {
	return g.Action.SignatureLocation.Offset
}

func (g *ScriptGroup) GetWitness() []byte {
	return g.OriginalWitness
}

func (g *ScriptGroup) GetWitnessIndex() int {
	return g.Action.SignatureLocation.Index
}

func (g *ScriptGroup) GetAddress() string {
	return g.Action.SignatureInfo.Address
}

func (g *ScriptGroup) GetGroupWitnesses() [][]byte {
	var groupWitnesses [][]byte
	groupWitnesses = append(groupWitnesses, g.OriginalWitness)
	for _, v := range g.Action.OtherIndexesInGroup {
		groupWitnesses = append(groupWitnesses, g.Transaction.Witnesses[v])
	}
	return groupWitnesses
}

func SignTransaction(transaction *types.Transaction, scriptGroup *ScriptGroup, privateKey crypto.Key) error {
	if isPWLock(scriptGroup) {
		return ethereumPersonalKeccakSign(transaction, scriptGroup, privateKey)
	} else {
		return secp256Blake2bSign(transaction, scriptGroup, privateKey)
	}
}

func secp256Blake2bSign(transaction *types.Transaction, scriptGroup *ScriptGroup, privateKey crypto.Key) error {
	witnessBytes := scriptGroup.GetWitness()
	groupWitnesses := scriptGroup.GetGroupWitnesses()

	txHash, err := transaction.ComputeHash()
	if err != nil {
		return err
	}

	length := make([]byte, 8)
	binary.LittleEndian.PutUint64(length, uint64(len(witnessBytes)))

	message := txHash.Bytes()
	message = append(message, length...)
	message = append(message, witnessBytes...)

	for i := 1; i < len(groupWitnesses); i++ {
		witnessBytes := groupWitnesses[i]
		length := make([]byte, 8)
		binary.LittleEndian.PutUint64(length, uint64(len(witnessBytes)))
		message = append(message, length...)
		message = append(message, witnessBytes...)
	}

	hash, err := blake2b.Blake256(message)
	if err != nil {
		return err
	}

	signature, err := privateKey.Sign(hash)
	if err != nil {
		return err
	}

	newWitness := scriptGroup.GetWitness()
	offset := scriptGroup.GetOffSet()
	for i := 0; i < len(signature); i++ {
		newWitness[i+offset] = signature[i]
	}
	transaction.Witnesses[scriptGroup.GetWitnessIndex()] = newWitness
	return nil
}

func ethereumPersonalKeccakSign(transaction *types.Transaction, scriptGroup *ScriptGroup, privateKey crypto.Key) error {
	witnessBytes := scriptGroup.GetWitness()
	groupWitnesses := scriptGroup.GetGroupWitnesses()

	txHash, err := transaction.ComputeHash()
	if err != nil {
		return err
	}

	length := make([]byte, 8)
	binary.LittleEndian.PutUint64(length, uint64(len(witnessBytes)))

	message := txHash.Bytes()
	message = append(message, length...)
	message = append(message, witnessBytes...)

	for i := 1; i < len(groupWitnesses); i++ {
		witnessBytes := groupWitnesses[i]
		length := make([]byte, 8)
		binary.LittleEndian.PutUint64(length, uint64(len(witnessBytes)))
		message = append(message, length...)
		message = append(message, witnessBytes...)
	}

	hash, err := ethereum.Keccak256(message)
	if err != nil {
		return err
	}

	prefix := []byte("\u0019Ethereum Signed Message:\n" + strconv.Itoa(len(hash)))

	message = append(prefix, hash...)
	hash, err = ethereum.Keccak256(message)

	if err != nil {
		return err
	}

	signature, err := privateKey.Sign(hash)
	if err != nil {
		return err
	}

	newWitness := scriptGroup.GetWitness()
	offset := scriptGroup.GetOffSet()
	for i := 0; i < len(signature); i++ {
		newWitness[i+offset] = signature[i]
	}
	transaction.Witnesses[scriptGroup.GetWitnessIndex()] = newWitness
	return nil
}

func isPWLock(g *ScriptGroup) bool {
	if Keccak256 == *g.Action.HashAlgorithm && EthereumPersonal == g.Action.SignatureInfo.Algorithm {
		return true
	}
	return false
}
