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
	"math/big"
)

// WithdrawCheque object
type WithdrawCheque struct {
	Sender   *types.Script
	Receiver *types.Script
	UUID     string
	FeeRate  uint64
	Amount   *big.Int
	tx       *types.Transaction
	groups   [][]int
}

// NewWithdrawCheque returns a new WithdrawCheque object
func NewWithdrawCheque(senderAddr, receiverAddr, uuid, amount string, feeRate uint64) (*WithdrawCheque, error) {
	parsedSenderAddr, err := address.Parse(senderAddr)
	if err != nil {
		return nil, errors.WithMessage(err, "invalid sender address")
	}
	parsedReceiverAddr, err := address.Parse(receiverAddr)
	if err != nil {
		return nil, errors.WithMessage(err, "invalid receiver address")
	}
	n, b := big.NewInt(0).SetString(amount, 10)
	if !b {
		return nil, errors.WithMessage(err, "invalid amount")
	}

	return &WithdrawCheque{
		Sender:   parsedSenderAddr.Script,
		Receiver: parsedReceiverAddr.Script,
		UUID:     uuid,
		Amount:   n,
		FeeRate:  feeRate,
	}, nil
}

// GenerateWithdrawChequeUnsignedTx generate an unsigned transaction for withdraw the cheque cell
func (c *WithdrawCheque) GenerateWithdrawChequeUnsignedTx(client rpc.Client, systemScripts *utils.SystemScripts) (*types.Transaction, error) {
	// collect udt cells
	udtType := &types.Script{
		CodeHash: systemScripts.SUDTCell.CellHash,
		HashType: systemScripts.SUDTCell.HashType,
		Args:     common.FromHex(c.UUID),
	}
	chequeCellArgs, err := utils.ChequeCellArgs(c.Sender, c.Receiver)
	if err != nil {
		return nil, err
	}

	chequeSearchKey := &indexer.SearchKey{
		Script: &types.Script{
			CodeHash: systemScripts.ChequeCell.CellHash,
			HashType: systemScripts.ChequeCell.HashType,
			Args:     chequeCellArgs,
		},
		ScriptType: indexer.ScriptTypeLock,
	}
	// cheque cell iterator
	chequeCollector := collector.NewLiveCellCollector(client, chequeSearchKey, indexer.SearchOrderAsc, indexer.SearchLimit, "")
	chequeCollector.TypeScript = udtType
	chequeIterator, err := chequeCollector.Iterator()
	if err != nil {
		return nil, fmt.Errorf("collect sudt cells error: %v", err)
	}

	ckbSearchKey := &indexer.SearchKey{
		Script:     c.Sender,
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
	txBuilder := &builder.WithdrawChequesUnsignedTxBuilder{
		Sender:         c.Sender,
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

// SignTx sign an unsigned withdraw cheque transaction and return an signed transaction
func (c *WithdrawCheque) SignTx(key crypto.Key) (*types.Transaction, error) {
	for _, group := range c.groups {
		err := transaction.SingleSignTransaction(c.tx, group, transaction.EmptyWitnessArg, key)
		if err != nil {
			return nil, fmt.Errorf("sign transaction error: %v", err)
		}
	}

	return c.tx, nil
}

// Send can send a tx to tx pool
func (c *WithdrawCheque) Send(client rpc.Client) (*types.Hash, error) {
	return client.SendTransaction(context.Background(), c.tx)
}
