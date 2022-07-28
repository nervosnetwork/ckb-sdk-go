package example

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/nervosnetwork/ckb-sdk-go/collector"
	builder2 "github.com/nervosnetwork/ckb-sdk-go/collector/builder"
	"github.com/nervosnetwork/ckb-sdk-go/indexer"
	"github.com/nervosnetwork/ckb-sdk-go/rpc"
	"github.com/nervosnetwork/ckb-sdk-go/transaction/signer"
	"github.com/nervosnetwork/ckb-sdk-go/types"
	"math/big"
)

func SendCkbExample() error {
	sender := "ckt1qzda0cr08m85hc8jlnfp3zer7xulejywt49kt2rr0vthywaa50xwsq2qf8keemy2p5uu0g0gn8cd4ju23s5269qk8rg4r"
	receiver := "ckt1qzda0cr08m85hc8jlnfp3zer7xulejywt49kt2rr0vthywaa50xwsq2qf8keemy2p5uu0g0gn8cd4ju23s5269qk8rg4r"
	network := types.NetworkTest
	indexerClient, err := indexer.Dial("https://testnet.ckb.dev/indexer")
	if err != nil {
		return err
	}
	iterator, err := collector.NewLiveCellIteratorFromAddress(indexerClient, sender)
	if err != nil {
		return err
	}

	builder := builder2.NewCkbTransactionBuilder(network, iterator)
	builder.FeeRate = 1000
	if err := builder.AddOutputByAddress(receiver, 50100000000); err != nil {
		return err
	}
	builder.AddChangeOutputByAddress(sender)
	txWithGroups, err := builder.Build()
	if err != nil {
		return err
	}

	txSigner := signer.GetTransactionSignerInstance(network)
	_, err = txSigner.SignTransactionByPrivateKeys(txWithGroups, "0x6c9ed03816e3111e49384b8d180174ad08e29feb1393ea1b51cef1c505d4e36a")
	if err != nil {
		return err
	}

	ckbClient, err := rpc.Dial("https://testnet.ckb.dev")
	hash, err := ckbClient.SendTransaction(context.Background(), txWithGroups.TxView)
	if err != nil {
		return err
	}
	fmt.Println("transaction hash: " + hexutil.Encode(hash.Bytes()))
	return nil
}

func IssueSudtExample() error {
	sender := "ckt1qzda0cr08m85hc8jlnfp3zer7xulejywt49kt2rr0vthywaa50xwsqdamwzrffgc54ef48493nfd2sd0h4cjnxg4850up"
	network := types.NetworkTest
	indexerClient, err := indexer.Dial("https://testnet.ckb.dev/indexer")
	if err != nil {
		return err
	}
	iterator, err := collector.NewLiveCellIteratorFromAddress(indexerClient, sender)
	if err != nil {
		return err
	}

	builder, err := builder2.NewSudtTransactionBuilderFromSudtOwnerAddress(network, iterator, builder2.SudtTransactionTypeIssue, sender)
	if err != nil {
		return err
	}
	builder.FeeRate = 1000
	_, err = builder.AddSudtOutputByAddress(sender, big.NewInt(99))
	if err != nil {
		return err
	}
	builder.AddChangeOutputByAddress(sender)
	txWithGroups, err := builder.Build()
	if err != nil {
		return err
	}

	txSigner := signer.GetTransactionSignerInstance(network)
	_, err = txSigner.SignTransactionByPrivateKeys(txWithGroups, "0x0c982052ffd4af5f3bbf232301dcddf468009161fc48ba1426e3ce0929fb59f8")
	if err != nil {
		return err
	}

	ckbClient, err := rpc.Dial("https://testnet.ckb.dev")
	hash, err := ckbClient.SendTransaction(context.Background(), txWithGroups.TxView)
	if err != nil {
		return err
	}
	fmt.Println("transaction hash: " + hexutil.Encode(hash.Bytes()))
	return nil
}

func SendSudtExample() error {
	sender := "ckt1qzda0cr08m85hc8jlnfp3zer7xulejywt49kt2rr0vthywaa50xwsqdamwzrffgc54ef48493nfd2sd0h4cjnxg4850up"
	receiver := "ckt1qzda0cr08m85hc8jlnfp3zer7xulejywt49kt2rr0vthywaa50xwsqd0pdquvfuq077aemn447shf4d8u5f4a0glzz2g4"
	sudtArgs := hexutil.MustDecode("0x9d2dab815b9158b2344827749d769fd66e2d3ebdfca32e5628ba0454651851f5")

	network := types.NetworkTest
	indexerClient, err := indexer.Dial("https://testnet.ckb.dev/indexer")
	if err != nil {
		return err
	}
	iterator, err := collector.NewLiveCellIteratorFromAddress(indexerClient, sender)
	if err != nil {
		return err
	}

	builder, err := builder2.NewSudtTransactionBuilderFromSudtArgs(network, iterator, builder2.SudtTransactionTypeTransfer, sudtArgs)
	if err != nil {
		return err
	}
	builder.FeeRate = 1000
	_, err = builder.AddSudtOutputByAddress(receiver, big.NewInt(1))
	if err != nil {
		return err
	}
	builder.AddChangeOutputByAddress(sender)
	txWithGroups, err := builder.Build()
	if err != nil {
		return err
	}

	txSigner := signer.GetTransactionSignerInstance(network)
	_, err = txSigner.SignTransactionByPrivateKeys(txWithGroups, "0x0c982052ffd4af5f3bbf232301dcddf468009161fc48ba1426e3ce0929fb59f8")
	if err != nil {
		return err
	}

	ckbClient, err := rpc.Dial("https://testnet.ckb.dev")
	hash, err := ckbClient.SendTransaction(context.Background(), txWithGroups.TxView)
	if err != nil {
		return err
	}
	fmt.Println("transaction hash: " + hexutil.Encode(hash.Bytes()))
	return nil
}

func DepositDaoExample() error {
	sender := "ckt1qzda0cr08m85hc8jlnfp3zer7xulejywt49kt2rr0vthywaa50xwsq2qf8keemy2p5uu0g0gn8cd4ju23s5269qk8rg4r"
	network := types.NetworkTest
	indexerClient, err := indexer.Dial("https://testnet.ckb.dev/indexer")
	if err != nil {
		return err
	}
	iterator, err := collector.NewLiveCellIteratorFromAddress(indexerClient, sender)
	if err != nil {
		return err
	}

	builder := builder2.NewCkbTransactionBuilder(network, iterator)
	builder.FeeRate = 1000
	if err := builder.AddDaoDepositOutputByAddress(sender, 50100000000); err != nil {
		return err
	}

	builder.AddChangeOutputByAddress(sender)
	txWithGroups, err := builder.Build()
	if err != nil {
		return err
	}

	txSigner := signer.GetTransactionSignerInstance(network)
	_, err = txSigner.SignTransactionByPrivateKeys(txWithGroups, "0x6c9ed03816e3111e49384b8d180174ad08e29feb1393ea1b51cef1c505d4e36a")
	if err != nil {
		return err
	}

	ckbClient, err := rpc.Dial("https://testnet.ckb.dev")
	hash, err := ckbClient.SendTransaction(context.Background(), txWithGroups.TxView)
	if err != nil {
		return err
	}
	fmt.Println("transaction hash: " + hexutil.Encode(hash.Bytes()))
	return nil
}
