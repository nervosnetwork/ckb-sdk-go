package example

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	address2 "github.com/nervosnetwork/ckb-sdk-go/address"
	"github.com/nervosnetwork/ckb-sdk-go/collector"
	"github.com/nervosnetwork/ckb-sdk-go/collector/builder"
	"github.com/nervosnetwork/ckb-sdk-go/collector/handler"
	"github.com/nervosnetwork/ckb-sdk-go/indexer"
	"github.com/nervosnetwork/ckb-sdk-go/rpc"
	"github.com/nervosnetwork/ckb-sdk-go/systemscript"
	"github.com/nervosnetwork/ckb-sdk-go/transaction"
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

	// build transaction
	b := builder.NewCkbTransactionBuilder(network, iterator)
	b.FeeRate = 1000
	if err := b.AddOutputByAddress(receiver, 50100000000); err != nil {
		return err
	}
	b.AddChangeOutputByAddress(sender)
	txWithGroups, err := b.Build()
	if err != nil {
		return err
	}

	// sign transaction
	txSigner := signer.GetTransactionSignerInstance(network)
	_, err = txSigner.SignTransactionByPrivateKeys(txWithGroups, "0x6c9ed03816e3111e49384b8d180174ad08e29feb1393ea1b51cef1c505d4e36a")
	if err != nil {
		return err
	}

	// send transaction
	ckbClient, err := rpc.Dial("https://testnet.ckb.dev")
	hash, err := ckbClient.SendTransaction(context.Background(), txWithGroups.TxView)
	if err != nil {
		return err
	}
	fmt.Println("transaction hash: " + hexutil.Encode(hash.Bytes()))
	return nil
}

func SendCkbFromMultisigAddressExample() error {
	network := types.NetworkTest

	multisigConfig := systemscript.NewMultisigConfig(0, 2)
	multisigConfig.AddKeyHash(hexutil.MustDecode("0x7336b0ba900684cb3cb00f0d46d4f64c0994a562"))
	multisigConfig.AddKeyHash(hexutil.MustDecode("0x5724c1e3925a5206944d753a6f3edaedf977d77f"))

	args := multisigConfig.Hash160()
	// ckt1qpw9q60tppt7l3j7r09qcp7lxnp3vcanvgha8pmvsa3jplykxn32sqdunqvd3g2felqv6qer8pkydws8jg9qxlca0st5v
	sender, _ := address2.Address{
		Script: &types.Script{
			CodeHash: systemscript.GetCodeHash(network, systemscript.Secp256k1Blake160MultisigAll),
			HashType: types.HashTypeType,
			Args:     args,
		},
		Network: network,
	}.Encode()

	receiver := "ckt1qzda0cr08m85hc8jlnfp3zer7xulejywt49kt2rr0vthywaa50xwsq2qf8keemy2p5uu0g0gn8cd4ju23s5269qk8rg4r"
	indexerClient, err := indexer.Dial("https://testnet.ckb.dev/indexer")
	if err != nil {
		return err
	}
	iterator, err := collector.NewLiveCellIteratorFromAddress(indexerClient, sender)
	if err != nil {
		return err
	}

	// build transaction
	builder := builder.NewCkbTransactionBuilder(network, iterator)
	builder.FeeRate = 1000
	if err := builder.AddOutputByAddress(receiver, 50100000000); err != nil {
		return err
	}
	builder.AddChangeOutputByAddress(sender)
	txWithGroups, err := builder.Build(multisigConfig)
	if err != nil {
		return err
	}

	// sign transaction
	txSigner := signer.GetTransactionSignerInstance(network)
	// first signature
	ctx1, _ := transaction.NewContextWithPayload("0x4fd809631a6aa6e3bb378dd65eae5d71df895a82c91a615a1e8264741515c79c", multisigConfig)
	if _, err = txSigner.SignTransaction(txWithGroups, ctx1); err != nil {
		return err
	}
	// second signature
	ctx2, _ := transaction.NewContextWithPayload("0x7438f7b35c355e3d2fb9305167a31a72d22ddeafb80a21cc99ff6329d92e8087", multisigConfig)
	if _, err = txSigner.SignTransaction(txWithGroups, ctx2); err != nil {
		return err
	}

	// send transaction
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

	// build transaction
	builder, err := builder.NewSudtTransactionBuilderFromSudtOwnerAddress(network, iterator, builder.SudtTransactionTypeIssue, sender)
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

	// sign transaction
	txSigner := signer.GetTransactionSignerInstance(network)
	_, err = txSigner.SignTransactionByPrivateKeys(txWithGroups, "0x0c982052ffd4af5f3bbf232301dcddf468009161fc48ba1426e3ce0929fb59f8")
	if err != nil {
		return err
	}

	// send transaction
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

	// build transaction
	builder := builder.NewSudtTransactionBuilderFromSudtArgs(network, iterator, builder.SudtTransactionTypeTransfer, sudtArgs)
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

	// sign transaction
	txSigner := signer.GetTransactionSignerInstance(network)
	_, err = txSigner.SignTransactionByPrivateKeys(txWithGroups, "0x0c982052ffd4af5f3bbf232301dcddf468009161fc48ba1426e3ce0929fb59f8")
	if err != nil {
		return err
	}

	// send transaction
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

	// build transaction
	builder := builder.NewCkbTransactionBuilder(network, iterator)
	builder.FeeRate = 1000
	if err := builder.AddDaoDepositOutputByAddress(sender, 50100000000); err != nil {
		return err
	}
	builder.AddChangeOutputByAddress(sender)
	txWithGroups, err := builder.Build()
	if err != nil {
		return err
	}

	// sign transaction
	txSigner := signer.GetTransactionSignerInstance(network)
	_, err = txSigner.SignTransactionByPrivateKeys(txWithGroups, "0x6c9ed03816e3111e49384b8d180174ad08e29feb1393ea1b51cef1c505d4e36a")
	if err != nil {
		return err
	}

	// send transaction
	ckbClient, err := rpc.Dial("https://testnet.ckb.dev")
	hash, err := ckbClient.SendTransaction(context.Background(), txWithGroups.TxView)
	if err != nil {
		return err
	}
	fmt.Println("transaction hash: " + hexutil.Encode(hash.Bytes()))
	return nil
}

func WithdrawDaoExample() error {
	sender := "ckt1qzda0cr08m85hc8jlnfp3zer7xulejywt49kt2rr0vthywaa50xwsq2qf8keemy2p5uu0g0gn8cd4ju23s5269qk8rg4r"
	depositOutPoint := &types.OutPoint{
		TxHash: types.HexToHash("0xebfb7bff39985865c20bd5b6c0190e58298eb9e4b1ddb9daee16166bed658c40"),
		Index:  0,
	}

	network := types.NetworkTest
	indexerClient, err := indexer.Dial("https://testnet.ckb.dev/indexer")
	if err != nil {
		return err
	}
	ckbClient, err := rpc.Dial("https://testnet.ckb.dev")
	if err != nil {
		return err
	}
	iterator, err := collector.NewLiveCellIteratorFromAddress(indexerClient, sender)
	if err != nil {
		return err
	}

	// build transaction
	builder, err := builder.NewDaoTransactionBuilder(network, iterator, depositOutPoint, ckbClient)
	if err != nil {
		return err
	}
	builder.FeeRate = 1000
	if err := builder.AddWithdrawOutput(sender); err != nil {
		return err
	}
	builder.AddChangeOutputByAddress(sender)

	withdrawInfo, err := handler.NewWithdrawInfo(ckbClient, depositOutPoint)
	if err != nil {
		return err
	}
	txWithGroups, err := builder.Build(withdrawInfo)
	if err != nil {
		return err
	}

	// sign transaction
	txSigner := signer.GetTransactionSignerInstance(network)
	_, err = txSigner.SignTransactionByPrivateKeys(txWithGroups, "0x6c9ed03816e3111e49384b8d180174ad08e29feb1393ea1b51cef1c505d4e36a")
	if err != nil {
		return err
	}

	// send transaction
	hash, err := ckbClient.SendTransaction(context.Background(), txWithGroups.TxView)
	if err != nil {
		return err
	}
	fmt.Println("transaction hash: " + hexutil.Encode(hash.Bytes()))
	return nil
}

func ClaimDaoExample() error {
	sender := "ckt1qzda0cr08m85hc8jlnfp3zer7xulejywt49kt2rr0vthywaa50xwsq2qf8keemy2p5uu0g0gn8cd4ju23s5269qk8rg4r"
	withdrawOutPoint := &types.OutPoint{
		TxHash: types.HexToHash("0xdb5b3ad48af332df51d0c87fcf7a180d7f8e284b743c08e4338cdd36b0a4b456"),
		Index:  0,
	}

	network := types.NetworkTest
	indexerClient, err := indexer.Dial("https://testnet.ckb.dev/indexer")
	if err != nil {
		return err
	}
	ckbClient, err := rpc.Dial("https://testnet.ckb.dev")
	if err != nil {
		return err
	}

	iterator, err := collector.NewLiveCellIteratorFromAddress(indexerClient, sender)
	if err != nil {
		return err
	}

	// build transaction
	builder, err := builder.NewDaoTransactionBuilder(network, iterator, withdrawOutPoint, ckbClient)
	if err != nil {
		return err
	}
	builder.FeeRate = 1000
	builder.AddChangeOutputByAddress(sender)

	claimInfo, err := handler.NewClaimInfo(ckbClient, withdrawOutPoint)
	if err != nil {
		return err
	}
	txWithGroups, err := builder.Build(claimInfo)
	if err != nil {
		return err
	}

	// sign transaction
	txSigner := signer.GetTransactionSignerInstance(network)
	_, err = txSigner.SignTransactionByPrivateKeys(txWithGroups, "0x6c9ed03816e3111e49384b8d180174ad08e29feb1393ea1b51cef1c505d4e36a")
	if err != nil {
		return err
	}

	// send transaction
	hash, err := ckbClient.SendTransaction(context.Background(), txWithGroups.TxView)
	if err != nil {
		return err
	}
	fmt.Println("transaction hash: " + hexutil.Encode(hash.Bytes()))
	return nil
}
