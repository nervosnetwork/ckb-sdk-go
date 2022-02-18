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
)

const chequeScriptArgsLength = 40

// ClaimCheque object
type ClaimCheque struct {
	Receiver      *types.Script
	UUID          string
	FeeRate       uint64
	tx            *types.Transaction
	systemScripts *utils.SystemScripts
	groups        map[string][]int
}

// NewClaimCheque returns a new ClaimCheque object
func NewClaimCheque(receiverAddr, uuid string, feeRate uint64, systemScripts *utils.SystemScripts) (*ClaimCheque, error) {
	parsedReceiverAddr, err := address.ValidateChequeAddress(receiverAddr, systemScripts)
	if err != nil {
		return nil, err
	}
	return &ClaimCheque{
		Receiver:      parsedReceiverAddr.Script,
		UUID:          uuid,
		FeeRate:       feeRate,
		systemScripts: systemScripts,
	}, nil
}

// GenerateClaimChequeUnsignedTx generate an unsigned transaction for claim cheque cells
func (c *ClaimCheque) GenerateClaimChequeUnsignedTx(client rpc.Client) (*types.Transaction, error) {
	// collect udt cells
	udtType := &types.Script{
		CodeHash: c.systemScripts.SUDTCell.CellHash,
		HashType: c.systemScripts.SUDTCell.HashType,
		Args:     common.FromHex(c.UUID),
	}
	receiverLockHash, err := c.Receiver.Hash()
	if err != nil {
		return nil, err
	}
	chequeSearchKey := &indexer.SearchKey{
		Script: &types.Script{
			CodeHash: c.systemScripts.ChequeCell.CellHash,
			HashType: c.systemScripts.ChequeCell.HashType,
			Args:     receiverLockHash.Bytes()[0:20],
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
		SystemScripts:  c.systemScripts,
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
		err := transaction.SingleSignTransaction(c.tx, group, transaction.Secp256k1EmptyWitnessArg, key)
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
