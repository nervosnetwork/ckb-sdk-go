package example

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/nervosnetwork/ckb-sdk-go/v2/address"
	"github.com/nervosnetwork/ckb-sdk-go/v2/collector"
	"github.com/nervosnetwork/ckb-sdk-go/v2/collector/builder"
	"github.com/nervosnetwork/ckb-sdk-go/v2/collector/handler"
	"github.com/nervosnetwork/ckb-sdk-go/v2/lightclient"
	"github.com/nervosnetwork/ckb-sdk-go/v2/rpc"
	"github.com/nervosnetwork/ckb-sdk-go/v2/systemscript"
	"github.com/nervosnetwork/ckb-sdk-go/v2/transaction"
	"github.com/nervosnetwork/ckb-sdk-go/v2/transaction/signer"
	"github.com/nervosnetwork/ckb-sdk-go/v2/transaction/signer/omnilock"
	"github.com/nervosnetwork/ckb-sdk-go/v2/types"
	"math/big"
)

func SendCkbExample() error {
	sender := "ckt1qzda0cr08m85hc8jlnfp3zer7xulejywt49kt2rr0vthywaa50xwsq2qf8keemy2p5uu0g0gn8cd4ju23s5269qk8rg4r"
	receiver := "ckt1qzda0cr08m85hc8jlnfp3zer7xulejywt49kt2rr0vthywaa50xwsq2qf8keemy2p5uu0g0gn8cd4ju23s5269qk8rg4r"
	network := types.NetworkTest
	client, err := rpc.Dial("https://testnet.ckb.dev")
	if err != nil {
		return err
	}
	iterator, err := collector.NewLiveCellIteratorFromAddress(client, sender)
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
	hash, err := client.SendTransaction(context.Background(), txWithGroups.TxView)
	if err != nil {
		return err
	}
	fmt.Println("transaction hash: " + hexutil.Encode(hash.Bytes()))
	return nil
}

func SendCkbByLightClientExample() error {
	sender := "ckt1qzda0cr08m85hc8jlnfp3zer7xulejywt49kt2rr0vthywaa50xwsq2qf8keemy2p5uu0g0gn8cd4ju23s5269qk8rg4r"
	receiver := "ckt1qzda0cr08m85hc8jlnfp3zer7xulejywt49kt2rr0vthywaa50xwsq2qf8keemy2p5uu0g0gn8cd4ju23s5269qk8rg4r"
	network := types.NetworkTest
	lightClient, err := lightclient.Dial("http://localhost:9000/")
	if err != nil {
		return err
	}
	senderAddress, err := address.Decode(sender)
	if err != nil {
		return err
	}
	senderScriptDetail := &lightclient.ScriptDetail{
		Script:      senderAddress.Script,
		ScriptType:  types.ScriptTypeLock,
		BlockNumber: 0,
	}
	// Set script to let light client sync information about this script on chain.
	lightClient.SetScripts(context.Background(), []*lightclient.ScriptDetail{senderScriptDetail})
	iterator, err := collector.NewLiveCellIteratorByLightClientFromAddress(lightClient, sender)
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
	hash, err := lightClient.SendTransaction(context.Background(), txWithGroups.TxView)
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
	sender, _ := address.Address{
		Script: &types.Script{
			CodeHash: systemscript.GetCodeHash(network, systemscript.Secp256k1Blake160MultisigAll),
			HashType: types.HashTypeType,
			Args:     args,
		},
		Network: network,
	}.Encode()

	receiver := "ckt1qzda0cr08m85hc8jlnfp3zer7xulejywt49kt2rr0vthywaa50xwsq2qf8keemy2p5uu0g0gn8cd4ju23s5269qk8rg4r"
	client, err := rpc.Dial("https://testnet.ckb.dev")
	if err != nil {
		return err
	}
	iterator, err := collector.NewLiveCellIteratorFromAddress(client, sender)
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
	hash, err := client.SendTransaction(context.Background(), txWithGroups.TxView)
	if err != nil {
		return err
	}
	fmt.Println("transaction hash: " + hexutil.Encode(hash.Bytes()))
	return nil
}

func IssueSudtExample() error {
	sender := "ckt1qzda0cr08m85hc8jlnfp3zer7xulejywt49kt2rr0vthywaa50xwsqdamwzrffgc54ef48493nfd2sd0h4cjnxg4850up"
	network := types.NetworkTest
	client, err := rpc.Dial("https://testnet.ckb.dev")
	if err != nil {
		return err
	}
	iterator, err := collector.NewLiveCellIteratorFromAddress(client, sender)
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
	hash, err := client.SendTransaction(context.Background(), txWithGroups.TxView)
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
	client, err := rpc.Dial("https://testnet.ckb.dev")
	if err != nil {
		return err
	}
	iterator, err := collector.NewLiveCellIteratorFromAddress(client, sender)
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
	hash, err := client.SendTransaction(context.Background(), txWithGroups.TxView)
	if err != nil {
		return err
	}
	fmt.Println("transaction hash: " + hexutil.Encode(hash.Bytes()))
	return nil
}

func DepositDaoExample() error {
	sender := "ckt1qzda0cr08m85hc8jlnfp3zer7xulejywt49kt2rr0vthywaa50xwsq2qf8keemy2p5uu0g0gn8cd4ju23s5269qk8rg4r"
	network := types.NetworkTest
	client, err := rpc.Dial("https://testnet.ckb.dev")
	if err != nil {
		return err
	}
	iterator, err := collector.NewLiveCellIteratorFromAddress(client, sender)
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
	hash, err := client.SendTransaction(context.Background(), txWithGroups.TxView)
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
	client, err := rpc.Dial("https://testnet.ckb.dev")
	if err != nil {
		return err
	}
	iterator, err := collector.NewLiveCellIteratorFromAddress(client, sender)
	if err != nil {
		return err
	}

	// build transaction
	builder, err := builder.NewDaoTransactionBuilder(network, iterator, depositOutPoint, client)
	if err != nil {
		return err
	}
	builder.FeeRate = 1000
	if err := builder.AddWithdrawOutput(sender); err != nil {
		return err
	}
	builder.AddChangeOutputByAddress(sender)

	withdrawInfo, err := handler.NewWithdrawInfo(client, depositOutPoint)
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
	hash, err := client.SendTransaction(context.Background(), txWithGroups.TxView)
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
	client, err := rpc.Dial("https://testnet.ckb.dev")
	if err != nil {
		return err
	}

	iterator, err := collector.NewLiveCellIteratorFromAddress(client, sender)
	if err != nil {
		return err
	}

	// build transaction
	builder, err := builder.NewDaoTransactionBuilder(network, iterator, withdrawOutPoint, client)
	if err != nil {
		return err
	}
	builder.FeeRate = 1000
	builder.AddChangeOutputByAddress(sender)

	claimInfo, err := handler.NewClaimInfo(client, withdrawOutPoint)
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
	hash, err := client.SendTransaction(context.Background(), txWithGroups.TxView)
	if err != nil {
		return err
	}
	fmt.Println("transaction hash: " + hexutil.Encode(hash.Bytes()))
	return nil
}

func SendCkbOmnilockExample() error {
	sender := "ckt1qrejnmlar3r452tcg57gvq8patctcgy8acync0hxfnyka35ywafvkqgqgpy7m88v3gxnn3apazvlpkkt32xz3tg5qq3kzjf3"
	receiver := "ckt1qzda0cr08m85hc8jlnfp3zer7xulejywt49kt2rr0vthywaa50xwsq2qf8keemy2p5uu0g0gn8cd4ju23s5269qk8rg4r"

	network := types.NetworkTest
	client, err := rpc.Dial("https://testnet.ckb.dev")
	if err != nil {
		return err
	}
	iterator, err := collector.NewLiveCellIteratorFromAddress(client, sender)
	if err != nil {
		return err
	}

	// prepare context for building and signing transaction
	config := new(signer.OmnilockConfiguration)
	if args, err := omnilock.NewOmnilockArgsFromAddress(sender); err != nil {
		return err
	} else {
		config.Args = args
	}
	config.Mode = signer.OmnolockModeAuth

	// build transaction
	builder := builder.NewCkbTransactionBuilder(network, iterator)
	builder.FeeRate = 1000
	if err := builder.AddOutputByAddress(receiver, 50100000000); err != nil {
		return err
	}
	builder.AddChangeOutputByAddress(sender)
	txWithGroups, err := builder.Build(config)
	if err != nil {
		return err
	}

	// sign transaction
	txSigner := signer.GetTransactionSignerInstance(network)
	ctx, err := transaction.NewContextWithPayload("0x6c9ed03816e3111e49384b8d180174ad08e29feb1393ea1b51cef1c505d4e36a", config)
	if err != nil {
		return err
	}
	if _, err = txSigner.SignTransaction(txWithGroups, ctx); err != nil {
		return err
	}

	// send transaction
	hash, err := client.SendTransaction(context.Background(), txWithGroups.TxView)
	if err != nil {
		return err
	}
	fmt.Println("transaction hash: " + hexutil.Encode(hash.Bytes()))
	return nil
}

func SendCkbMultisigOmnilockExample() error {
	sender := "ckt1qrejnmlar3r452tcg57gvq8patctcgy8acync0hxfnyka35ywafvkqgxhjvp3k9pf88upngryvuxc346q7fq5qmlqqlrhr0p"
	receiver := "ckt1qzda0cr08m85hc8jlnfp3zer7xulejywt49kt2rr0vthywaa50xwsq2qf8keemy2p5uu0g0gn8cd4ju23s5269qk8rg4r"

	network := types.NetworkTest
	client, err := rpc.Dial("https://testnet.ckb.dev")
	if err != nil {
		return err
	}
	iterator, err := collector.NewLiveCellIteratorFromAddress(client, sender)
	if err != nil {
		return err
	}

	// prepare context for building and signing transaction
	config := new(signer.OmnilockConfiguration)
	if args, err := omnilock.NewOmnilockArgsFromAddress(sender); err != nil {
		return err
	} else {
		config.Args = args
	}
	config.Mode = signer.OmnolockModeAuth
	multisigConfig := systemscript.NewMultisigConfig(0, 2)
	multisigConfig.AddKeyHash(hexutil.MustDecode("0x7336b0ba900684cb3cb00f0d46d4f64c0994a562"))
	multisigConfig.AddKeyHash(hexutil.MustDecode("0x5724c1e3925a5206944d753a6f3edaedf977d77f"))
	config.MultisigConfig = multisigConfig

	// build transaction
	builder := builder.NewCkbTransactionBuilder(network, iterator)
	builder.FeeRate = 1000
	if err := builder.AddOutputByAddress(receiver, 50100000000); err != nil {
		return err
	}
	builder.AddChangeOutputByAddress(sender)
	txWithGroups, err := builder.Build(config)
	if err != nil {
		return err
	}

	// sign transaction
	txSigner := signer.GetTransactionSignerInstance(network)
	// first signature
	ctx, err := transaction.NewContextWithPayload("0x7438f7b35c355e3d2fb9305167a31a72d22ddeafb80a21cc99ff6329d92e8087", config)
	if err != nil {
		return err
	}
	if _, err = txSigner.SignTransaction(txWithGroups, ctx); err != nil {
		return err
	}
	// second signature
	ctx, err = transaction.NewContextWithPayload("0x4fd809631a6aa6e3bb378dd65eae5d71df895a82c91a615a1e8264741515c79c", config)
	if err != nil {
		return err
	}
	if _, err = txSigner.SignTransaction(txWithGroups, ctx); err != nil {
		return err
	}

	// send transaction
	hash, err := client.SendTransaction(context.Background(), txWithGroups.TxView)
	if err != nil {
		return err
	}
	fmt.Println("transaction hash: " + hexutil.Encode(hash.Bytes()))
	return nil
}

func SendChainedTransactionExample() error {
	address1 := "ckt1qzda0cr08m85hc8jlnfp3zer7xulejywt49kt2rr0vthywaa50xwsq2qf8keemy2p5uu0g0gn8cd4ju23s5269qk8rg4r"
	address2 := "ckt1qzda0cr08m85hc8jlnfp3zer7xulejywt49kt2rr0vthywaa50xwsqdamwzrffgc54ef48493nfd2sd0h4cjnxg4850up"
	address3 := "ckt1qrejnmlar3r452tcg57gvq8patctcgy8acync0hxfnyka35ywafvkqgxhjvp3k9pf88upngryvuxc346q7fq5qmlqqlrhr0p"

	network := types.NetworkTest
	client, err := rpc.Dial("https://testnet.ckb.dev")
	if err != nil {
		return err
	}

	offChainInputCollector := collector.NewOffChainInputCollector(client)

	var iterator, err_ = collector.NewOffChainInputIteratorFromAddress(client, address1, offChainInputCollector, true)
	if err_ != nil {
		return err_
	}

	txWithGroupsBuilder := builder.NewCkbTransactionBuilder(network, iterator)
	err = txWithGroupsBuilder.AddOutputByAddress(address2, 50100000000)
	if err != nil {
		return err
	}

	err = txWithGroupsBuilder.AddChangeOutputByAddress(address1)
	if err != nil {
		return err
	}

	config := new(signer.TransactionSigner)
	txWithGroups, err := txWithGroupsBuilder.Build(config)

	if err != nil {
		return err
	}

	// sign transaction
	if _, err = signer.GetTransactionSignerInstance(network).SignTransactionByPrivateKeys(txWithGroups, "0x6c9ed03816e3111e49384b8d180174ad08e29feb1393ea1b51cef1c505d4e36a"); err != nil {
		return err
	}

	// send transaction
	hash, err := client.SendTransaction(context.Background(), txWithGroups.TxView)

	if err != nil {
		return err
	}

	blockNumber, err := client.GetTipBlockNumber(context.Background())

	if err != nil {
		return err
	}

	offChainInputCollector.ApplyOffChainTransaction(blockNumber, *txWithGroups.TxView)
	fmt.Println("transaction hash: " + hexutil.Encode(hash.Bytes()))

	it2, _ := collector.NewLiveCellIteratorFromAddress(client, address2)
	cell_iterator2 := it2.(*collector.LiveCellIterator)
	var iterator2 = collector.OffChainInputIterator{
		Iterator:                    cell_iterator2,
		Collector:                   offChainInputCollector,
		ConsumeOffChainCellsFirstly: true,
	}

	txWithGroupsBuilder = builder.NewCkbTransactionBuilder(network, &iterator2)
	err = txWithGroupsBuilder.AddOutputByAddress(address3, 100000000000)
	if err != nil {
		return err
	}

	txWithGroupsBuilder.AddChangeOutputByAddress(address2)

	txWithGroups, err = txWithGroupsBuilder.Build(config)

	// sign transaction
	txSigner := signer.GetTransactionSignerInstance(network)
	_, err = txSigner.SignTransactionByPrivateKeys(txWithGroups, "0x0c982052ffd4af5f3bbf232301dcddf468009161fc48ba1426e3ce0929fb59f8")
	if err != nil {
		return err
	}

	// send transaction
	hash, err = client.SendTransaction(context.Background(), txWithGroups.TxView)

	if err != nil {
		return err
	}

	blockNumber, err = client.GetTipBlockNumber(context.Background())

	if err != nil {
		return err
	}

	offChainInputCollector.ApplyOffChainTransaction(blockNumber, *txWithGroups.TxView)
	fmt.Println("transaction hash: " + hexutil.Encode(hash.Bytes()))

	return nil
}
