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

// Cheque object
type Cheque struct {
	Sender        *types.Script
	Receiver      *types.Script
	UUID          string
	Amount        *big.Int
	FeeRate       uint64
	tx            *types.Transaction
	systemScripts *utils.SystemScripts
}

// NewCheque returns a new Cheque object
func NewCheque(senderAddr, receiverAddr, uuid, amount string, feeRate uint64, systemScripts *utils.SystemScripts) (*Cheque, error) {
	parsedSenderAddr, err := address.ValidateChequeAddress(senderAddr, systemScripts)
	if err != nil {
		return nil, err
	}
	parsedReceiverAddr, err := address.ValidateChequeAddress(receiverAddr, systemScripts)
	if err != nil {
		return nil, err
	}
	n, b := big.NewInt(0).SetString(amount, 10)
	if !b {
		return nil, errors.WithMessage(err, "invalid amount")
	}

	return &Cheque{
		Sender:        parsedSenderAddr.Script,
		Receiver:      parsedReceiverAddr.Script,
		UUID:          uuid,
		Amount:        n,
		FeeRate:       feeRate,
		systemScripts: systemScripts,
	}, nil
}

// GenerateIssuingChequeUnsignedTx generate an unsigned transaction for issuing a cheque cell
func (c *Cheque) GenerateIssuingChequeUnsignedTx(client rpc.Client) (*types.Transaction, error) {
	// collect udt cells
	udtType := &types.Script{
		CodeHash: c.systemScripts.SUDTCell.CellHash,
		HashType: c.systemScripts.SUDTCell.HashType,
		Args:     common.FromHex(c.UUID),
	}
	searchKey := &indexer.SearchKey{
		Script:     c.Sender,
		ScriptType: indexer.ScriptTypeLock,
	}

	// sudt Iterator
	sudtCollector := collector.NewLiveCellCollector(client, searchKey, indexer.SearchOrderAsc, indexer.SearchLimit, "")
	sudtCollector.TypeScript = udtType
	sudtIterator, err := sudtCollector.Iterator()
	if err != nil {
		return nil, fmt.Errorf("collect sudt cells error: %v", err)
	}
	// ckb Iterator
	ckbCollector := collector.NewLiveCellCollector(client, searchKey, indexer.SearchOrderAsc, indexer.SearchLimit, "")
	ckbCollector.EmptyData = true
	ckbIterator, err := ckbCollector.Iterator()
	if err != nil {
		return nil, fmt.Errorf("collect ckb cells error: %v", err)
	}

	director := builder.Director{}
	txBuilder := &builder.IssuingChequeUnsignedTxBuilder{
		Sender:         c.Sender,
		Receiver:       c.Receiver,
		FeeRate:        c.FeeRate,
		CkbIterator:    ckbIterator,
		SUDTIterator:   sudtIterator,
		SystemScripts:  c.systemScripts,
		TransferAmount: c.Amount,
		UUID:           c.UUID,
	}
	director.SetBuilder(txBuilder)
	tx, _, err := director.Generate()
	c.tx = tx

	return tx, err
}

// SignTx sign an unsigned issuing cheque transaction and return an signed transaction
func (c *Cheque) SignTx(key crypto.Key) (*types.Transaction, error) {
	err := transaction.SingleSegmentSignTransaction(c.tx, 0, len(c.tx.Witnesses), transaction.EmptyWitnessArg, key)
	if err != nil {
		return nil, fmt.Errorf("sign transaction error: %v", err)
	}
	return c.tx, nil
}

// Send can send a tx to tx pool
func (c *Cheque) Send(client rpc.Client) (*types.Hash, error) {
	return client.SendTransaction(context.Background(), c.tx)
}
