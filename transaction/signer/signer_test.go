package signer

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/nervosnetwork/ckb-sdk-go/crypto/secp256k1"
	"github.com/nervosnetwork/ckb-sdk-go/transaction"
	"github.com/nervosnetwork/ckb-sdk-go/types"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"runtime/debug"
	"testing"
)

func TestIsMatch(t *testing.T) {
	key, _ := secp256k1.HexToKey("9d8ca87d75d150692211fa62b0d30de4d1ee6c530d5678b40b8cedacf0750d0f")
	args := common.FromHex("af0b41c627807fbddcee75afa174d5a7e5135ebd")
	actual, err := IsMatch(key, args)
	assert.Equal(t, true, actual)
	assert.Nil(t, nil, err)

	key, _ = secp256k1.HexToKey("9d8ca87d75d150692211fa62b0d30de4d1ee6c530d5678b40b8cedacf0750d0f")
	args = common.FromHex("0450340178ae277261a838c89f9ccb76a190ed4b")
	actual, err = IsMatch(key, args)
	assert.Equal(t, false, actual)
	assert.Nil(t, err)

	actual, err = IsMatch(nil, args)
	assert.Equal(t, false, actual)
	assert.NotNil(t, err)

	actual, err = IsMatch(key, nil)
	assert.Equal(t, false, actual)
	assert.NotNil(t, err)
}

func TestMultiScriptDecode(t *testing.T) {
	bytes := common.FromHex("0x000002029b41c025515b00c24e2e2042df7b221af5c1891fe732dcd15b7618eb1d7a11e6a68e4579b5be0114")
	m, err := DecodeToMultisigScript(bytes)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, byte(0), m.FirstN)
	assert.Equal(t, byte(2), m.Threshold)
	assert.Equal(t, getKeysHashes(), m.KeysHashes)

	bytes = common.FromHex("0x000002039b41c025515b00c24e2e2042df7b221af5c1891fe732dcd15b7618eb1d7a11e6a68e4579b5be0114")
	_, err = DecodeToMultisigScript(bytes)
	assert.Error(t, err)

	bytes = common.FromHex("0x000002029b41c025515b00c24e2e2042df7b221af5c1891f")
	_, err = DecodeToMultisigScript(bytes)
	assert.Error(t, err)
}

func TestMultiScriptEncode(t *testing.T) {
	m := &MultisigScript{
		Version:    0,
		FirstN:     0,
		Threshold:  2,
		KeysHashes: getKeysHashes(),
	}
	encoded := m.encode()
	assert.Equal(t, common.FromHex("0x000002029b41c025515b00c24e2e2042df7b221af5c1891fe732dcd15b7618eb1d7a11e6a68e4579b5be0114"), encoded)
	hash, err := m.ComputeHash()
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, common.FromHex("0x35ed7b939b4ac9cb447b82340fd8f26d344f7a62"), hash)
}

func getKeysHashes() [][20]byte {
	keysHashes := make([][20]byte, 0)
	b := [20]byte{}
	copy(b[:], common.FromHex("0x9b41c025515b00c24e2e2042df7b221af5c1891f"))
	keysHashes = append(keysHashes, b)
	copy(b[:], common.FromHex("0xe732dcd15b7618eb1d7a11e6a68e4579b5be0114"))
	keysHashes = append(keysHashes, b)
	return keysHashes
}

func TestSecp256k1Blake160SighashAllSigner(t *testing.T) {
	testSignAndCheck(t, "secp256k1_blake160_sighash_all_one_input.json")
	testSignAndCheck(t, "secp256k1_blake160_sighash_all_one_group.json")
	testSignAndCheck(t, "secp256k1_blake160_sighash_all_two_groups.json")
	testSignAndCheck(t, "secp256k1_blake160_sighash_all_extra_witness.json")
}

func TestSecp256k1Blake160MultisigAllSigner(t *testing.T) {
	testSignAndCheck(t, "secp256k1_blake160_multisig_all_first.json")
	testSignAndCheck(t, "secp256k1_blake160_multisig_all_second.json")
}

func testSignAndCheck(t *testing.T, fileName string) {
	checker, err := fromFile(fileName)
	if err != nil {
		t.Error(err, string(debug.Stack()))
	}
	txSigner := GetTransactionSignerInstance(types.NetworkTest)
	tx := checker.Transaction
	signed, err := txSigner.signTransaction(tx, checker.Contexts)
	if err != nil {
		t.Error(err, string(debug.Stack()))
	}
	assert.Equal(t, len(tx.ScriptGroups), len(signed))
	signedMap := make(map[int]bool)
	for _, v := range signed {
		signedMap[v] = true
	}
	for i, _ := range tx.ScriptGroups {
		assert.True(t, signedMap[i], fmt.Sprintf("group #%d is not signed", i))
	}
	assert.Equal(t, len(checker.ExpectedWitnesses), len(tx.TxView.Witnesses))
	for i, w := range checker.ExpectedWitnesses {
		assert.Equal(t, w, hexutil.Encode(tx.TxView.Witnesses[i]))
	}
}

func fromFile(fileName string) (*signerChecker, error) {
	content, err := ioutil.ReadFile("./test-fixture/" + fileName)
	if err != nil {
		return nil, err
	}
	var checker signerChecker
	if err = json.Unmarshal(content, &checker); err != nil {
		return nil, err
	}
	return &checker, nil
}

type signerChecker struct {
	Transaction       *transaction.TransactionWithScriptGroups `json:"raw_transaction"`
	ExpectedWitnesses []string                                 `json:"expected_witnesses"`
	Contexts          transaction.Contexts                     `json:"Contexts"`
}

func (r *signerChecker) UnmarshalJSON(input []byte) error {
	var jsonObj struct {
		Transaction       *transaction.TransactionWithScriptGroups `json:"raw_transaction"`
		ExpectedWitnesses []string                                 `json:"expected_witnesses"`
		Contexts          []map[string]interface{}                 `json:"contexts"`
	}
	if err := json.Unmarshal(input, &jsonObj); err != nil {
		return err
	}
	r.Transaction = jsonObj.Transaction
	r.ExpectedWitnesses = jsonObj.ExpectedWitnesses
	r.Contexts = transaction.NewContexts()
	for _, c := range jsonObj.Contexts {
		var (
			ctx *transaction.Context
			err error
		)
		if val, ok := c["private_key"]; ok {
			if ctx, err = transaction.NewContext(val.(string)); err != nil {
				return err
			}
		} else {
			return errors.New("not find private_key")
		}
		if val, ok := c["multisig_script"]; ok {
			v := val.(map[string]interface{})
			m := NewMultisigScript(byte(v["first_n"].(float64)),
				byte(v["threshold"].(float64)))
			for _, h := range v["key_hashes"].([]interface{}) {
				if err := m.AddKeyHashBySlice(common.FromHex(h.(string))); err != nil {
					return err
				}
			}
			ctx.Payload = m
		}
		r.Contexts.Add(ctx)
	}
	return nil
}
