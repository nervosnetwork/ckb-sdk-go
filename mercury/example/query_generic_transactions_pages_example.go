package test

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/nervosnetwork/ckb-sdk-go/address"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/example/constant"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model"
	"github.com/nervosnetwork/ckb-sdk-go/types"
	"testing"
)

func TestQueryGenericTransactionsWithCkb(t *testing.T) {
	builder := model.NewQueryGenericTransactionsPayloadBuilder()
	builder.AddAddress(constant.QUERY_TRANSACTION_ADDRESS)

	payload, err := builder.Build()
	if err != nil {
		t.Error(err)
	}

	marshal, _ := json.Marshal(payload)
	fmt.Println(string(marshal))

	transactions, err := constant.GetMercuryApiInstance().QueryGenericTransactions(payload)
	if err != nil {
		t.Error(err)
	}

	json, err := json.Marshal(transactions)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(len(transactions.Txs))
	fmt.Println(string(json))

}

func TestQueryGenericTransactionsWithUdt(t *testing.T) {
	builder := model.NewQueryGenericTransactionsPayloadBuilder()
	builder.AddAddress(constant.QUERY_TRANSACTION_ADDRESS)
	builder.AddUdtHash(constant.UdtHash)

	payload, err := builder.Build()
	if err != nil {
		t.Error(err)
	}

	marshal, _ := json.Marshal(payload)
	fmt.Println(string(marshal))

	transactions, err := constant.GetMercuryApiInstance().QueryGenericTransactions(payload)
	if err != nil {
		t.Error(err)
	}

	json, err := json.Marshal(transactions)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(len(transactions.Txs))
	fmt.Println(string(json))

}

func TestQueryGenericTransactionsWithAll(t *testing.T) {
	builder := model.NewQueryGenericTransactionsPayloadBuilder()
	builder.AddAddress(constant.QUERY_TRANSACTION_ADDRESS)
	builder.AllTransactionType()

	payload, err := builder.Build()
	if err != nil {
		t.Error(err)
	}

	marshal, _ := json.Marshal(payload)
	fmt.Println(string(marshal))

	transactions, err := constant.GetMercuryApiInstance().QueryGenericTransactions(payload)
	if err != nil {
		t.Error(err)
	}

	json, err := json.Marshal(transactions)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(len(transactions.Txs))
	fmt.Println(string(json))

}

func TestQueryGenericTransactionsWithChequeAddress(t *testing.T) {
	builder := model.NewQueryGenericTransactionsPayloadBuilder()
	builder.AddAddress(chequeAddress())
	builder.AllTransactionType()

	payload, err := builder.Build()
	if err != nil {
		t.Error(err)
	}

	marshal, _ := json.Marshal(payload)
	fmt.Println(string(marshal))

	transactions, err := constant.GetMercuryApiInstance().QueryGenericTransactions(payload)
	if err != nil {
		t.Error(err)
	}

	json, err := json.Marshal(transactions)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(len(transactions.Txs))
	fmt.Println(string(json))

}

func TestQueryGenericTransactionsWithAcpAddress(t *testing.T) {
	builder := model.NewQueryGenericTransactionsPayloadBuilder()
	builder.AddAddress(getQueryTransactionAcpAddress())
	builder.AllTransactionType()

	payload, err := builder.Build()
	if err != nil {
		t.Error(err)
	}

	marshal, _ := json.Marshal(payload)
	fmt.Println(string(marshal))

	transactions, err := constant.GetMercuryApiInstance().QueryGenericTransactions(payload)
	if err != nil {
		t.Error(err)
	}

	json, err := json.Marshal(transactions)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(len(transactions.Txs))
	fmt.Println(string(json))

}

func TestQueryGenericTransactionsWithFromBlock(t *testing.T) {
	builder := model.NewQueryGenericTransactionsPayloadBuilder()
	builder.AddAddress(getQueryTransactionAcpAddress())
	builder.AllTransactionType()
	builder.AddFromBlock(2224987)

	payload, err := builder.Build()
	if err != nil {
		t.Error(err)
	}

	marshal, _ := json.Marshal(payload)
	fmt.Println(string(marshal))

	transactions, err := constant.GetMercuryApiInstance().QueryGenericTransactions(payload)
	if err != nil {
		t.Error(err)
	}

	json, err := json.Marshal(transactions)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(len(transactions.Txs))
	fmt.Println(string(json))

}

func TestQueryGenericTransactionsWithToBlock(t *testing.T) {
	builder := model.NewQueryGenericTransactionsPayloadBuilder()
	builder.AddAddress(getQueryTransactionAcpAddress())
	builder.AllTransactionType()
	builder.AddToBlock(2224987)

	payload, err := builder.Build()
	if err != nil {
		t.Error(err)
	}

	marshal, _ := json.Marshal(payload)
	fmt.Println(string(marshal))

	transactions, err := constant.GetMercuryApiInstance().QueryGenericTransactions(payload)
	if err != nil {
		t.Error(err)
	}

	json, err := json.Marshal(transactions)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(len(transactions.Txs))
	fmt.Println(string(json))

}

func TestQueryGenericTransactionsWithFromBlockAndToBlock(t *testing.T) {
	builder := model.NewQueryGenericTransactionsPayloadBuilder()
	builder.AddAddress(getQueryTransactionAcpAddress())
	builder.AllTransactionType()
	builder.AddFromBlock(2224993)
	builder.AddToBlock(2225023)

	payload, err := builder.Build()
	if err != nil {
		t.Error(err)
	}

	marshal, _ := json.Marshal(payload)
	fmt.Println(string(marshal))

	transactions, err := constant.GetMercuryApiInstance().QueryGenericTransactions(payload)
	if err != nil {
		t.Error(err)
	}

	json, err := json.Marshal(transactions)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(len(transactions.Txs))
	fmt.Println(string(json))

}

func TestQueryGenericTransactionsWithLimit(t *testing.T) {
	builder := model.NewQueryGenericTransactionsPayloadBuilder()
	builder.AddAddress(getQueryTransactionAcpAddress())
	builder.AllTransactionType()
	// default limit 50
	builder.AddLimit(2)

	payload, err := builder.Build()
	if err != nil {
		t.Error(err)
	}

	marshal, _ := json.Marshal(payload)
	fmt.Println(string(marshal))

	transactions, err := constant.GetMercuryApiInstance().QueryGenericTransactions(payload)
	if err != nil {
		t.Error(err)
	}

	json, err := json.Marshal(transactions)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(len(transactions.Txs))
	fmt.Println(string(json))

}

func TestQueryGenericTransactionsWithOrder(t *testing.T) {
	builder := model.NewQueryGenericTransactionsPayloadBuilder()
	builder.AddAddress(constant.QUERY_TRANSACTION_ADDRESS)
	builder.AllTransactionType()
	// default order desc
	builder.AddOrder("asc")

	payload, err := builder.Build()
	if err != nil {
		t.Error(err)
	}

	marshal, _ := json.Marshal(payload)
	fmt.Println(string(marshal))

	transactions, err := constant.GetMercuryApiInstance().QueryGenericTransactions(payload)
	if err != nil {
		t.Error(err)
	}

	json, err := json.Marshal(transactions)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(len(transactions.Txs))
	fmt.Println(string(json))

}

func TestQueryGenericTransactionsWithOffset(t *testing.T) {
	builder := model.NewQueryGenericTransactionsPayloadBuilder()
	builder.AddAddress(constant.QUERY_TRANSACTION_ADDRESS)
	builder.AllTransactionType()
	builder.AddLimit(1)
	// Offset start from 0
	builder.AddOffset(1)

	payload, err := builder.Build()
	if err != nil {
		t.Error(err)
	}

	marshal, _ := json.Marshal(payload)
	fmt.Println(string(marshal))

	transactions, err := constant.GetMercuryApiInstance().QueryGenericTransactions(payload)
	if err != nil {
		t.Error(err)
	}

	json, err := json.Marshal(transactions)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(len(transactions.Txs))
	fmt.Println(string(json))

}

func getQueryTransactionAcpAddress() string {
	pubKey := "0x8d135c59240be2229b19eec0be5a006b34b3b0cb"

	acpLock := &types.Script{
		CodeHash: types.HexToHash("0x3419a1c09eb2567f6552ee7a8ecffd64155cffe0f1796e6e61ec088d740c1356"),
		HashType: types.HashTypeType,
		Args:     common.FromHex(pubKey),
	}

	address, _ := address.Generate(address.Testnet, acpLock)

	return address
}

func chequeAddress() string {
	senderScript, _ := address.Parse(constant.TEST_ADDRESS0)
	receiverScript, _ := address.Parse(constant.QUERY_TRANSACTION_ADDRESS)

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
