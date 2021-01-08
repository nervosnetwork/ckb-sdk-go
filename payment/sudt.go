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
	Senders        []*types.Script
	ReceiverInfo   []types.ReceiverInfo
	CkbPayer       *types.Script
	CkbChanger     *types.Script
	SudtChanger    *types.Script
	UUID           string
	TransferAmount *big.Int
	FeeRate        uint64

	tx            *types.Transaction
	systemScripts *utils.SystemScripts
	groups        map[string][]int
}

// NewSudt returns a new NewSudt object
// receiverInfo's key is an address and its value is a sudt amount
func NewSudt(senderAddresses []string, receiverInfo map[string]string, ckbPayerAddress, ckbChangeAddress, sudtChangeAddress, uuid string, feeRate uint64, systemScripts *utils.SystemScripts) (*Sudt, error) {
	var senders []*types.Script
	var receivers []types.ReceiverInfo
	totalAmount := big.NewInt(0)
	for _, senderAddr := range senderAddresses {
		parsedSenderAddr, err := address.Parse(senderAddr)
		if err != nil {
			return nil, err
		}
		senders = append(senders, parsedSenderAddr.Script)
	}
	for receiverAddr, amount := range receiverInfo {
		parsedReceiverAddr, err := address.Parse(receiverAddr)
		if err != nil {
			return nil, err
		}
		n, b := big.NewInt(0).SetString(amount, 10)
		if !b {
			return nil, errors.New("invalid amount")
		}
		receivers = append(receivers, types.ReceiverInfo{
			Receiver: parsedReceiverAddr.Script,
			Amount:   n,
		})
		totalAmount = big.NewInt(0).Add(totalAmount, n)
	}
	parsedPayFeeAddr, err := address.Parse(ckbPayerAddress)
	if err != nil {
		return nil, err
	}
	parsedCkbChangeAddr, err := address.Parse(ckbChangeAddress)
	if err != nil {
		return nil, err
	}
	parsedSudtChangeAddr, err := address.Parse(sudtChangeAddress)
	if err != nil {
		return nil, err
	}

	return &Sudt{
		Senders:        senders,
		ReceiverInfo:   receivers,
		CkbPayer:       parsedPayFeeAddr.Script,
		CkbChanger:     parsedCkbChangeAddr.Script,
		SudtChanger:    parsedSudtChangeAddr.Script,
		UUID:           uuid,
		TransferAmount: totalAmount,
		FeeRate:        feeRate,

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
	var sudtIterators []collector.CellCollectionIterator
	// sudt Iterators
	for _, sender := range s.Senders {
		searchKey := &indexer.SearchKey{
			Script:     sender,
			ScriptType: indexer.ScriptTypeLock,
		}
		sudtCollector := collector.NewLiveCellCollector(client, searchKey, indexer.SearchOrderAsc, indexer.SearchLimit, "")
		sudtCollector.TypeScript = udtType
		sudtIterator, err := sudtCollector.Iterator()
		if err != nil {
			return nil, errors.Errorf("collect sudt cells error: %v", err)
		}
		sudtIterators = append(sudtIterators, sudtIterator)
	}

	// ckb Iterator
	searchKey := &indexer.SearchKey{
		Script:     s.CkbPayer,
		ScriptType: indexer.ScriptTypeLock,
	}
	ckbCollector := collector.NewLiveCellCollector(client, searchKey, indexer.SearchOrderAsc, indexer.SearchLimit, "")
	ckbCollector.EmptyData = true
	ckbIterator, err := ckbCollector.Iterator()
	if err != nil {
		return nil, errors.Errorf("collect sudt cells error: %v", err)
	}
	director := builder.Director{}
	txBuilder := &builder.SudtTransferUnsignedTxBuilder{
		CkbChanger:     s.CkbChanger,
		SudtChanger:    s.SudtChanger,
		Senders:        s.Senders,
		ReceiverInfo:   s.ReceiverInfo,
		FeeRate:        s.FeeRate,
		CkbIterator:    ckbIterator,
		SUDTIterators:  sudtIterators,
		SystemScripts:  s.systemScripts,
		TransferAmount: s.TransferAmount,
		UUID:           s.UUID,
	}

	director.SetBuilder(txBuilder)
	tx, groups, err := director.Generate()
	s.tx = tx
	s.groups = groups

	return tx, err
}

// SignTx sign an unsigned sudt transfer transaction and return an signed transaction
// The order of keys must be consistent with the order of locks in senders
func (s *Sudt) SignTx(keys map[string]crypto.Key) (*types.Transaction, error) {
	for lockHash, group := range s.groups {
		err := transaction.SingleSignTransaction(s.tx, group, transaction.EmptyWitnessArg, keys[lockHash])
		if err != nil {
			return nil, fmt.Errorf("sign transaction error: %v", err)
		}
	}

	return s.tx, nil
}

// Send can send a tx to tx pool
func (s *Sudt) Send(client rpc.Client) (*types.Hash, error) {
	return client.SendTransaction(context.Background(), s.tx)
}
