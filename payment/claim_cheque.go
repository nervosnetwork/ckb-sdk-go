package payment

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/nervosnetwork/ckb-sdk-go/address"
	"github.com/nervosnetwork/ckb-sdk-go/collector"
	"github.com/nervosnetwork/ckb-sdk-go/crypto"
	"github.com/nervosnetwork/ckb-sdk-go/indexer"
	"github.com/nervosnetwork/ckb-sdk-go/rpc"
	"github.com/nervosnetwork/ckb-sdk-go/transaction"
	"github.com/nervosnetwork/ckb-sdk-go/transaction/builder"
	"github.com/nervosnetwork/ckb-sdk-go/types"
	"github.com/nervosnetwork/ckb-sdk-go/utils"
	"github.com/pkg/errors"
)

const chequeScriptArgsLength = 40

// ClaimCheque object
type ClaimCheque struct {
	Receiver *types.Script
	UUID     string
	FeeRate  uint64
	tx       *types.Transaction
	groups   [][]int
}

// NewClaimCheque returns a new ClaimCheque object
func NewClaimCheque(receiverAddr, uuid string, feeRate uint64) (*ClaimCheque, error) {
	parsedReceiverAddr, err := address.Parse(receiverAddr)
	if err != nil {
		return nil, errors.WithMessage(err, "invalid receiver address")
	}
	return &ClaimCheque{
		Receiver: parsedReceiverAddr.Script,
		UUID:     uuid,
		FeeRate:  feeRate,
	}, nil
}

// GenerateClaimChequeUnsignedTx generate an unsigned transaction for claim cheque cells
func (c *ClaimCheque) GenerateClaimChequeUnsignedTx(client rpc.Client, systemScripts *utils.SystemScripts) (*types.Transaction, error) {
	// collect udt cells
	udtType := &types.Script{
		CodeHash: systemScripts.SUDTCell.CellHash,
		HashType: systemScripts.SUDTCell.HashType,
		Args:     common.FromHex(c.UUID),
	}
	chequeSearchKey := &indexer.SearchKey{
		Script: &types.Script{
			CodeHash: systemScripts.ChequeCell.CellHash,
			HashType: systemScripts.ChequeCell.HashType,
			Args:     c.Receiver.Args,
		},
		ScriptType: indexer.ScriptTypeLock,
		ArgsLen:    chequeScriptArgsLength,
	}
	// cheque cell iterator
	chequeCollector := collector.NewLiveCellCollector(client, chequeSearchKey, indexer.SearchOrderAsc, indexer.SearchLimit, "")
	chequeCollector.TypeScript = udtType
	chequeIterator, err := chequeCollector.Iterator()
	if err != nil {
		return nil, fmt.Errorf("collect sudt cells error: %v", err)
	}

	ckbSearchKey := &indexer.SearchKey{
		Script:     c.Receiver,
		ScriptType: indexer.ScriptTypeLock,
	}
	// ckb Iterator
	ckbCollector := collector.NewLiveCellCollector(client, ckbSearchKey, indexer.SearchOrderAsc, indexer.SearchLimit, "")
	ckbCollector.EmptyData = true
	ckbIterator, err := ckbCollector.Iterator()
	if err != nil {
		return nil, fmt.Errorf("collect ckb cells error: %v", err)
	}

	director := builder.Director{}
	txBuilder := &builder.ClaimChequesUnsignedTxBuilder{
		Receiver:       c.Receiver,
		FeeRate:        c.FeeRate,
		CkbIterator:    ckbIterator,
		ChequeIterator: chequeIterator,
		SystemScripts:  systemScripts,
		UUID:           c.UUID,
		Client:         client,
	}
	director.SetBuilder(txBuilder)
	tx, groups, err := director.Generate()
	if err != nil {
		return nil, err
	}
	c.tx = tx
	c.groups = groups

	return tx, err
}

// SignTx sign an unsigned claim cheque transaction and return an signed transaction
func (c *ClaimCheque) SignTx(key crypto.Key) (*types.Transaction, error) {
	for _, group := range c.groups {
		err := transaction.SingleSignTransaction(c.tx, group, transaction.EmptyWitnessArg, key)
		if err != nil {
			return nil, fmt.Errorf("sign transaction error: %v", err)
		}
	}

	return c.tx, nil
}

// Send can send a tx to tx pool
func (c *ClaimCheque) Send(client rpc.Client) (*types.Hash, error) {
	return client.SendTransaction(context.Background(), c.tx)
}
