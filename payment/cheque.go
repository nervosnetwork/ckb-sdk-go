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
	"github.com/nervosnetwork/ckb-sdk-go/types"
	"github.com/nervosnetwork/ckb-sdk-go/utils"
	"github.com/pkg/errors"
	"math"
	"math/big"
)

var (
	chequeCellCapacity = uint64(162 * math.Pow10(8))
	udtCellCapacity    = uint64(142 * math.Pow10(8))
)

type Cheque struct {
	Sender   *types.Script
	Receiver *types.Script
	UUID     string
	Amount   *big.Int
	FeeRate  uint64
	tx       *types.Transaction
}

func NewCheque(senderAddr, receiverAddr, uuid, amount string, feeRate uint64) (*Cheque, error) {
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
	return &Cheque{
		Sender:   parsedSenderAddr.Script,
		Receiver: parsedReceiverAddr.Script,
		UUID:     uuid,
		Amount:   n,
		FeeRate:  feeRate,
	}, nil
}

// GenerateIssueChequeTx generate an unsigned transaction for issuing a cheque cell
func (c *Cheque) GenerateIssueChequeTx(client rpc.Client) (*types.Transaction, error) {
	systemScripts, err := utils.NewSystemScripts(client)
	if err != nil {
		return nil, err
	}
	udtType := &types.Script{
		CodeHash: systemScripts.SUDTCell.CellHash,
		HashType: systemScripts.SUDTCell.HashType,
		Args:     common.FromHex(c.UUID),
	}

	tx := transaction.NewSecp256k1SingleSigTx(systemScripts)
	// set sudt and cheque scripts cell deps
	tx.CellDeps = append(tx.CellDeps, &types.CellDep{
		OutPoint: systemScripts.SUDTCell.OutPoint,
		DepType:  systemScripts.SUDTCell.DepType,
	}, &types.CellDep{
		OutPoint: systemScripts.ChequeCell.OutPoint,
		DepType:  systemScripts.ChequeCell.DepType,
	})

	// cheque output
	chequeCellArgs, err := utils.ChequeCellArgs(c.Sender, c.Receiver)
	if err != nil {
		return nil, err
	}
	tx.Outputs = append(tx.Outputs, &types.CellOutput{
		Capacity: chequeCellCapacity,
		Lock: &types.Script{
			CodeHash: systemScripts.ChequeCell.CellHash,
			HashType: systemScripts.ChequeCell.HashType,
			Args:     chequeCellArgs,
		},
		Type: udtType,
	})
	tx.OutputsData = append(tx.OutputsData, utils.GenerateSudtAmount(c.Amount))

	// ckb change output
	tx.Outputs = append(tx.Outputs, &types.CellOutput{
		Capacity: 0,
		Lock:     c.Sender,
	})
	tx.OutputsData = append(tx.OutputsData, []byte{})

	// sudt change output
	tx.Outputs = append(tx.Outputs, &types.CellOutput{
		Capacity: udtCellCapacity,
		Lock:     c.Sender,
		Type:     udtType,
	})
	tx.OutputsData = append(tx.OutputsData, utils.GenerateSudtAmount(big.NewInt(0)))

	// collect udt cells
	searchKey := &indexer.SearchKey{
		Script:     c.Sender,
		ScriptType: indexer.ScriptTypeLock,
	}
	processor := collector.NewUDTLiveCellProcessor(c.Amount)
	processor.CkbChangeOutputIndex = &collector.ChangeOutputIndex{Value: 1}
	processor.SUDTChangeOutputIndex = &collector.ChangeOutputIndex{Value: 2}
	processor.TypeScript = udtType
	processor.Tx = tx
	processor.FeeRate = c.FeeRate
	nCollector := utils.NewLiveCellCollector(client, searchKey, indexer.SearchOrderAsc, indexer.SearchLimit, "", processor)
	cells, err := nCollector.Collect()
	if err != nil {
		return nil, fmt.Errorf("collect cell error: %v", err)
	}
	totalAmount := cells.Options["totalAmount"].(*big.Int)
	if totalAmount.Cmp(c.Amount) < 0 {
		return nil, errors.New("insufficient udt balance")
	}
	c.tx = tx

	return tx, nil
}

// SignIssueChequeTx sign an unsigned issuing cheque transaction and return an signed transaction
func (c *Cheque) SignIssueChequeTx(key crypto.Key) (*types.Transaction, error) {
	err := transaction.SingleSegmentSignTransaction(c.tx, 0, len(c.tx.Witnesses), transaction.EmptyWitnessArg, key)
	if err != nil {
		return nil, fmt.Errorf("sign transaction error: %v", err)
	}
	return c.tx, nil
}

func (c *Cheque) Send(client rpc.Client) (*types.Hash, error) {
	return client.SendTransaction(context.Background(), c.tx)
}
