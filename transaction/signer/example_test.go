package signer_test

import (
	"context"
	"encoding/binary"
	"reflect"

	"github.com/nervosnetwork/ckb-sdk-go/v2/rpc"
	"github.com/nervosnetwork/ckb-sdk-go/v2/transaction"
	"github.com/nervosnetwork/ckb-sdk-go/v2/transaction/signer"
	"github.com/nervosnetwork/ckb-sdk-go/v2/types"
)

type CapacityDiffContext struct {
	rpc rpc.Client
	ctx context.Context
}

func (ctx CapacityDiffContext) getInputCell(outPoint *types.OutPoint) (*types.CellOutput, error) {
	cellWithStatus, err := ctx.rpc.GetLiveCell(ctx.ctx, outPoint, false)
	if err != nil {
		return nil, err
	}

	return cellWithStatus.Cell.Output, nil
}

type CapacityDiffScriptSigner struct{}

func (s *CapacityDiffScriptSigner) SignTransaction(tx *types.Transaction, group *transaction.ScriptGroup, ctx *transaction.Context) (bool, error) {
	scriptContext, ok := ctx.Payload.(CapacityDiffContext)
	if !ok {
		return false, nil
	}

	total := int64(0)
	for _, i := range group.InputIndices {
		inputCell, err := scriptContext.getInputCell(tx.Inputs[i].PreviousOutput)
		if err != nil {
			return false, nil
		}
		total -= int64(inputCell.Capacity)
	}
	for _, output := range tx.Outputs {
		if reflect.DeepEqual(output.Lock, group.Script) {
			total += int64(output.Capacity)
		}
	}

	// The specification https://go.dev/ref/spec#Numeric_types says integres in
	// Go are repsented using two's complementation. So we can just cast it to
	// uin64 and get the little endian bytes.
	witness := make([]byte, 8)
	binary.LittleEndian.PutUint64(witness, uint64(total))

	witnessIndex := group.InputIndices[0]
	witnessArgs, err := types.DeserializeWitnessArgs(tx.Witnesses[witnessIndex])
	if err != nil {
		return false, err
	}
	witnessArgs.Lock = witness
	tx.Witnesses[witnessIndex] = witnessArgs.Serialize()

	return true, nil
}

// This example demonstrates how to use a custom script CapacityDiff
// (https://github.com/doitian/ckb-sdk-examples-capacity-diff).
//
// CapacityDiff verifies the witness matches the capacity difference.
//
//   - The script loads the witness for the first input in the script group using the WitnessArgs layout.
//   - The total input capacity is the sum of all the input cells in the script group.
//   - The total output capacity is the sum of all the output cells having the same lock script as the script group.
//   - The capacity difference is a 64-bit signed integer which equals to total output capacity minus total input capacity.
//   - The witness is encoded using two's complement and little endian.
func ExampleScriptSigner() {
	signer := signer.NewTransactionSigner()
	signer.RegisterSigner(
		types.HexToHash("0x6283a479a3cf5d4276cd93594de9f1827ab9b55c7b05b3d28e4c2e0a696cfefd"),
		types.ScriptTypeType,
		&CapacityDiffScriptSigner{},
	)
}
