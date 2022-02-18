package payment

import (
	"context"
	"fmt"
	"github.com/nervosnetwork/ckb-sdk-go/collector"
	"github.com/nervosnetwork/ckb-sdk-go/indexer"
	"github.com/nervosnetwork/ckb-sdk-go/transaction/builder"

	"github.com/nervosnetwork/ckb-sdk-go/address"
	"github.com/nervosnetwork/ckb-sdk-go/crypto"
	"github.com/nervosnetwork/ckb-sdk-go/rpc"
	"github.com/nervosnetwork/ckb-sdk-go/transaction"
	"github.com/nervosnetwork/ckb-sdk-go/types"
	"github.com/nervosnetwork/ckb-sdk-go/utils"
)

type Payment struct {
	From        *types.Script
	To          *types.Script
	Amount      uint64
	FeeRate     uint64
	group       []int
	witnessArgs *types.WitnessArgs
	tx          *types.Transaction
}

// NewPayment returns a Payment object, amount's unit is shannon
func NewPayment(from, to string, amount, feeRate uint64) (*Payment, error) {
	fromAddress, err := address.Parse(from)
	if err != nil {
		return nil, fmt.Errorf("parse from address %s error: %v", from, err)
	}
	toAddress, err := address.Parse(to)
	if err != nil {
		return nil, fmt.Errorf("parse to address %s error: %v", to, err)
	}

	if fromAddress.Mode != toAddress.Mode {
		return nil, fmt.Errorf("from address and to address with diffrent network: %v:%v", fromAddress.Mode, toAddress.Mode)
	}

	return &Payment{
		From:    fromAddress.Script,
		To:      toAddress.Script,
		Amount:  amount,
		FeeRate: feeRate,
	}, nil
}

func (p *Payment) GenerateTx(client rpc.Client, systemScripts *utils.SystemScripts) (*types.Transaction, error) {
	return generateTxWithIndexer(client, p, systemScripts)
}

func generateTxWithIndexer(client rpc.Client, p *Payment, systemScripts *utils.SystemScripts) (*types.Transaction, error) {
	searchKey := &indexer.SearchKey{
		Script:     p.From,
		ScriptType: indexer.ScriptTypeLock,
	}
	c := collector.NewLiveCellCollector(client, searchKey, indexer.SearchOrderAsc, indexer.SearchLimit, "")
	c.EmptyData = true

	iterator, err := c.Iterator()
	if err != nil {
		return nil, fmt.Errorf("collect cell error: %v", err)
	}
	director := builder.Director{}
	txBuilder := &builder.CkbTransferUnsignedTxBuilder{
		From:             p.From,
		To:               p.To,
		FeeRate:          p.FeeRate,
		Iterator:         iterator,
		SystemScripts:    systemScripts,
		TransferCapacity: p.Amount,
	}
	director.SetBuilder(txBuilder)
	tx, _, err := director.Generate()
	p.tx = tx

	return tx, err
}

func (p *Payment) Sign(key crypto.Key) (*types.Transaction, error) {
	err := transaction.SingleSegmentSignTransaction(p.tx, 0, len(p.tx.Witnesses), transaction.Secp256k1EmptyWitnessArg, key)
	if err != nil {
		return nil, fmt.Errorf("sign transaction error: %v", err)
	}

	return p.tx, err
}

func (p *Payment) Send(client rpc.Client) (*types.Hash, error) {
	return client.SendTransaction(context.Background(), p.tx)
}
