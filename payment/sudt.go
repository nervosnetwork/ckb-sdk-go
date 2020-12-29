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

type Sudt struct {
	Sender   *types.Script
	Receiver *types.Script
	UUID     string
	Amount   *big.Int
	FeeRate  uint64

	tx            *types.Transaction
	systemScripts *utils.SystemScripts
}

// NewSudt returns a new NewSudt object
func NewSudt(senderAddr, receiverAddr, uuid, amount string, feeRate uint64, systemScripts *utils.SystemScripts) (*Sudt, error) {
	parsedSenderAddr, err := address.Parse(senderAddr)
	if err != nil {
		return nil, err
	}
	parsedReceiverAddr, err := address.Parse(receiverAddr)
	if err != nil {
		return nil, err
	}
	n, b := big.NewInt(0).SetString(amount, 10)
	if !b {
		return nil, errors.WithMessage(err, "invalid amount")
	}

	return &Sudt{
		Sender:        parsedSenderAddr.Script,
		Receiver:      parsedReceiverAddr.Script,
		UUID:          uuid,
		Amount:        n,
		FeeRate:       feeRate,
		systemScripts: systemScripts,
	}, nil
}

// GenerateTransferSudtUnsignedTx generate an unsigned transaction for transfer sudt
func (s *Sudt) GenerateTransferSudtUnsignedTx(client rpc.Client) (*types.Transaction, error) {
	// udt type script
	udtType := &types.Script{
		CodeHash: s.systemScripts.SUDTCell.CellHash,
		HashType: s.systemScripts.SUDTCell.HashType,
		Args:     common.FromHex(s.UUID),
	}
	searchKey := &indexer.SearchKey{
		Script:     s.Sender,
		ScriptType: indexer.ScriptTypeLock,
	}
	// sudt Iterator
	sudtCollector := collector.NewLiveCellCollector(client, searchKey, indexer.SearchOrderAsc, indexer.SearchLimit, "")
	sudtCollector.TypeScript = udtType
	sudtIterator, err := sudtCollector.Iterator()
	if err != nil {
		return nil, errors.Errorf("collect sudt cells error: %v", err)
	}
	// ckb Iterator
	ckbCollector := collector.NewLiveCellCollector(client, searchKey, indexer.SearchOrderAsc, indexer.SearchLimit, "")
	ckbCollector.EmptyData = true
	ckbIterator, err := ckbCollector.Iterator()
	if err != nil {
		return nil, errors.Errorf("collect sudt cells error: %v", err)
	}
	director := builder.Director{}
	txBuilder := &builder.SudtTransferUnsignedTxBuilder{
		Sender:         s.Sender,
		Receiver:       s.Receiver,
		FeeRate:        s.FeeRate,
		CkbIterator:    ckbIterator,
		SUDTIterator:   sudtIterator,
		SystemScripts:  s.systemScripts,
		TransferAmount: s.Amount,
		UUID:           s.UUID,
	}

	director.SetBuilder(txBuilder)
	tx, _, err := director.Generate()
	s.tx = tx

	return tx, err
}

// SignTx sign an unsigned sudt transfer transaction and return an signed transaction
func (s *Sudt) SignTx(key crypto.Key) (*types.Transaction, error) {
	err := transaction.SingleSegmentSignTransaction(s.tx, 0, len(s.tx.Witnesses), transaction.EmptyWitnessArg, key)
	if err != nil {
		return nil, fmt.Errorf("sign transaction error: %v", err)
	}
	return s.tx, nil
}

// Send can send a tx to tx pool
func (s *Sudt) Send(client rpc.Client) (*types.Hash, error) {
	return client.SendTransaction(context.Background(), s.tx)
}
