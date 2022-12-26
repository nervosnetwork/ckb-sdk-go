package signer_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/nervosnetwork/ckb-sdk-go/v2/crypto/secp256k1"
	"github.com/nervosnetwork/ckb-sdk-go/v2/systemscript"
	"github.com/nervosnetwork/ckb-sdk-go/v2/transaction"
	"github.com/nervosnetwork/ckb-sdk-go/v2/transaction/signer"
	"github.com/nervosnetwork/ckb-sdk-go/v2/transaction/signer/omnilock"
	"github.com/nervosnetwork/ckb-sdk-go/v2/types"
	"github.com/stretchr/testify/assert"
	"os"
	"runtime/debug"
	"testing"
)

func TestIsSingleSigMatched(t *testing.T) {
	key, _ := secp256k1.HexToKey("9d8ca87d75d150692211fa62b0d30de4d1ee6c530d5678b40b8cedacf0750d0f")
	args := common.FromHex("af0b41c627807fbddcee75afa174d5a7e5135ebd")
	actual, err := signer.IsSingleSigMatched(key, args)
	assert.Equal(t, true, actual)
	assert.Nil(t, nil, err)

	key, _ = secp256k1.HexToKey("9d8ca87d75d150692211fa62b0d30de4d1ee6c530d5678b40b8cedacf0750d0f")
	args = common.FromHex("0450340178ae277261a838c89f9ccb76a190ed4b")
	actual, err = signer.IsSingleSigMatched(key, args)
	assert.Equal(t, false, actual)
	assert.Nil(t, err)

	actual, err = signer.IsSingleSigMatched(nil, args)
	assert.Equal(t, false, actual)
	assert.NotNil(t, err)

	actual, err = signer.IsSingleSigMatched(key, nil)
	assert.Equal(t, false, actual)
	assert.NotNil(t, err)
}

func TestIsPWLockMatched(t *testing.T) {
	k, err := secp256k1.HexToKey("f8f8a2f43c8376ccb0871305060d7b27b0554d2cc72bccf41b2705608452f315")
	if err != nil {
		t.Error(err)
	}
	assert.True(t, signer.IsPWLockMatched(k, common.FromHex("001d3f1ef827552ae1114027bd3ecf1f086ba0f9")))

	k, err = secp256k1.HexToKey("e0ccb2548af279947b452efda4535dd4bcadf756d919701fcd4c382833277f85")
	if err != nil {
		t.Error(err)
	}
	assert.True(t, signer.IsPWLockMatched(k, common.FromHex("adabffb9c27cb4af100ce7bca6903315220e87a2")))
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

func TestAnyoneCanPaySigner(t *testing.T) {
	testSignAndCheck(t, "acp_one_input.json")
}

func TestPWLockSigner(t *testing.T) {
	testSignAndCheck(t, "pw_one_group.json")
}

func TestOmnilockSigner(t *testing.T) {
	testSignAndCheck(t, "omnilock_secp256k1_blake160_sighash_all.json")
	testSignAndCheck(t, "omnilock_secp256k1_blake160_multisig_all_first.json")
	testSignAndCheck(t, "omnilock_secp256k1_blake160_multisig_all_second.json")
}

func testSignAndCheck(t *testing.T, fileName string) {
	checker, err := fromFile(fileName)
	if err != nil {
		t.Error(err, string(debug.Stack()))
	}
	txSigner := signer.GetTransactionSignerInstance(types.NetworkTest)
	tx := checker.Transaction
	signed, err := txSigner.SignTransaction(tx, checker.Contexts...)
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
	content, err := os.ReadFile("./fixture/" + fileName)
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
	Contexts          []*transaction.Context                   `json:"Contexts"`
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
	for _, context := range jsonObj.Contexts {
		var (
			ctx *transaction.Context
			err error
		)
		if val, ok := context["private_key"]; ok {
			if ctx, err = transaction.NewContext(val.(string)); err != nil {
				return err
			}
		} else {
			return errors.New("not find private_key")
		}
		if val, ok := context["multisig_script"]; ok {
			ctx.Payload = unmarshalMultisigConfig(val.(map[string]interface{}))
		}
		if val, ok := context["omnilock_config"]; ok {
			v := val.(map[string]interface{})
			config := new(signer.OmnilockConfiguration)
			args, err := omnilock.NewOmnilockArgsFromAgrs(hexutil.MustDecode(v["args"].(string)))
			if err != nil {
				return err
			}
			config.Args = args
			switch v["mode"].(string) {
			case "AUTH":
				config.Mode = signer.OmnolockModeAuth
			case "ADMINISTRATOR":
				config.Mode = signer.OmnolockModeAdministrator
			default:
				return fmt.Errorf("unknown mode %s", v["mode"])
			}
			if val, ok := v["multisig_script"]; ok {
				config.MultisigConfig = unmarshalMultisigConfig(val.(map[string]interface{}))
			}
			ctx.Payload = config
		}
		r.Contexts = append(r.Contexts, ctx)
	}
	return nil
}

func unmarshalMultisigConfig(input map[string]interface{}) *systemscript.MultisigConfig {
	config := systemscript.NewMultisigConfig(byte(input["first_n"].(float64)),
		byte(input["threshold"].(float64)))
	for _, h := range input["key_hashes"].([]interface{}) {
		config.AddKeyHash(common.FromHex(h.(string)))
	}
	return config
}
