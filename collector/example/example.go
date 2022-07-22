package example

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/nervosnetwork/ckb-sdk-go/collector"
	"github.com/nervosnetwork/ckb-sdk-go/indexer"
	"github.com/nervosnetwork/ckb-sdk-go/rpc"
	"github.com/nervosnetwork/ckb-sdk-go/transaction/signer"
	"github.com/nervosnetwork/ckb-sdk-go/types"
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

	builder := collector.NewCkbTransactionBuilder(network, iterator)
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