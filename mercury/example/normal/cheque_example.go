package normal

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/nervosnetwork/ckb-sdk-go/address"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/example/constant"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/action"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/resp"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/source"
	"github.com/nervosnetwork/ckb-sdk-go/types"
	"testing"
	"time"
)

const (
	senderAddress             = constant.TEST_ADDRESS1
	chequeCellReceiverAddress = constant.TEST_ADDRESS2
	receiverAddress           = constant.TEST_ADDRESS3
	udtHash                   = "0xf21e7350fa9518ed3cbb008e0e8c941d7e01a12181931d5608aa366ee22228bd"
)

func TestFleeting(t *testing.T) {
	printBalance()
	issuingChequeCell()
	printBalance()
	claimChequeCell()
	printBalance()
}

func issuingChequeCell() {
	mercuryApi := constant.GetMercuryApiInstance()
	ckbNode := constant.GetCkbNodeInstance()

	builder := model.NewTransferBuilder()
	builder.AddUdtHash(udtHash)
	builder.AddFromKeyAddresses([]string{senderAddress}, source.Unconstrained)
	builder.AddToKeyAddressItem(chequeCellReceiverAddress, action.Lend_by_from, 100)
	transferPayload := builder.Build()
	transferCompletion, err := mercuryApi.BuildTransferTransaction(transferPayload)
	if err != nil {
		fmt.Println(err)
	}

	tx := constant.Sign(transferCompletion)

	hash, err := ckbNode.SendTransaction(context.Background(), tx)
	if err != nil {
		fmt.Println(err)
	}

	var txStatus types.TransactionStatus = "pending"
	for {
		transaction, _ := ckbNode.GetTransaction(context.Background(), *hash)
		if transaction.TxStatus.Status != txStatus {
			break
		}
		fmt.Println("Awaiting transaction results")
		time.Sleep(1 * 1e9)

	}

	time.Sleep(60 * 1e9)
	fmt.Printf("send hash of cheque cell transactions: %s\n", hash.String())
}

func claimChequeCell() {
	mercuryApi := constant.GetMercuryApiInstance()
	ckbNode := constant.GetCkbNodeInstance()

	builder := model.NewTransferBuilder()
	builder.AddUdtHash(udtHash)
	builder.AddFromNormalAddresses([]string{getChequeAddress()})
	builder.AddToKeyAddressItem(receiverAddress, action.Pay_by_from, 100)
	transferPayload := builder.Build()
	transferCompletion, err := mercuryApi.BuildTransferTransaction(transferPayload)
	if err != nil {
		fmt.Println(err)
	}

	tx := constant.Sign(transferCompletion)

	hash, err := ckbNode.SendTransaction(context.Background(), tx)
	if err != nil {
		fmt.Println(err)
	}

	var txStatus types.TransactionStatus = "pending"
	for {
		transaction, _ := ckbNode.GetTransaction(context.Background(), *hash)
		if transaction.TxStatus.Status != txStatus {
			break
		}
		fmt.Println("Awaiting transaction results")
		time.Sleep(1 * 1e9)
	}

	time.Sleep(60 * 1e9)
	fmt.Printf("claim hash of cheque cell transactions: %s\n", hash.String())
}

func printBalance() {
	ckbBalanceA := getCkbBalance(senderAddress)
	udtBalanceA := getUdtBalance(senderAddress, udtHash)

	fmt.Printf("sender ckb balance: %s\n", getJsonStr(ckbBalanceA))
	fmt.Printf("sender udt balance: %s\n", getJsonStr(udtBalanceA))

	ckbBalanceB := getCkbBalance(chequeCellReceiverAddress)
	udtBalanceB := getUdtBalance(chequeCellReceiverAddress, udtHash)

	fmt.Printf("chequeCellReceiverAddress ckb balance: %s\n", getJsonStr(ckbBalanceB))
	fmt.Printf("chequeCellReceiverAddress udt balance: %s\n", getJsonStr(udtBalanceB))
}

func getCkbBalance(addr string) *resp.GetBalanceResponse {
	builder := model.NewGetBalancePayloadBuilder()
	builder.AddAddress(addr)
	payload, err := builder.Build()
	if err != nil {
		panic(err)
	}
	balance, _ := constant.GetMercuryApiInstance().GetBalance(payload)

	return balance
}

func getUdtBalance(addr, udtHash string) *resp.GetBalanceResponse {
	builder := model.NewGetBalancePayloadBuilder()
	builder.AddAddress(addr)
	builder.AddUdtHash(udtHash)
	payload, err := builder.Build()
	if err != nil {
		panic(err)
	}
	balance, _ := constant.GetMercuryApiInstance().GetBalance(payload)

	return balance
}

func getJsonStr(balance *resp.GetBalanceResponse) string {
	jsonStr, _ := json.Marshal(balance)
	return string(jsonStr)
}

func getChequeAddress() string {
	senderScript, _ := address.Parse(senderAddress)
	receiverScript, _ := address.Parse(chequeCellReceiverAddress)

	senderScriptHash, _ := senderScript.Script.Hash()
	receiverScriptHash, _ := receiverScript.Script.Hash()

	fmt.Printf("senderScriptHash: %s\n", senderScriptHash)
	fmt.Printf("receiverScript: %s\n", receiverScriptHash)

	s1 := senderScriptHash.String()[0:42]
	s2 := receiverScriptHash.String()[0:42]

	args := bytesCombine(common.FromHex(s2), common.FromHex(s1))
	pubKey := common.Bytes2Hex(args)
	fmt.Printf("pubKey: %s\n", pubKey)

	chequeLock := &types.Script{
		CodeHash: types.HexToHash("0x60d5f39efce409c587cb9ea359cefdead650ca128f0bd9cb3855348f98c70d5b"),
		HashType: types.HashTypeType,
		Args:     common.FromHex(pubKey),
	}

	address, _ := address.Generate(address.Testnet, chequeLock)

	fmt.Printf("address: %s\n", address)
	return address
}

func bytesCombine(pBytes ...[]byte) []byte {
	return bytes.Join(pBytes, []byte(""))
}
