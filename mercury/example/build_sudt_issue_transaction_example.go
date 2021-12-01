package test

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/example/constant"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/example/utils"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model"
	mcommon "github.com/nervosnetwork/ckb-sdk-go/mercury/model/common"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/mode"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/req"
	"github.com/nervosnetwork/ckb-sdk-go/types"
	"github.com/nervosnetwork/ckb-sdk-go/utils/amount"
	"testing"
)

func TestSudtIssue(t *testing.T) {
	secpCodeHash := "0x9bd7e06f3ecf4be0f2fcd2188b23f1b9fcc88e5d4b65a8637b17723bbda3cce8"
	sudtTestnetCodeHash := "0xc5e5dcf215925f7ef4dfaf5f4b4f105bc321c02776d6e7d52a1db3fcd9d011a4"
	adminAddress := constant.TEST_ADDRESS3
	adminPublicKey := constant.TEST_PUBKEY3
	receiverAddress := constant.TEST_ADDRESS0
	issueUdtAmount := uint64(1000000000)
	item2, _ := req.NewIdentityItemByAddress(receiverAddress)
	fmt.Println(item2)

	// 1 - get admin lock hash
	adminScript := &types.Script{
		CodeHash: types.HexToHash(secpCodeHash),
		HashType: types.HashTypeType,
		Args:     common.FromHex(adminPublicKey),
	}
	adminLockHash, _ := adminScript.Hash()
	t.Log("admin lock hash: " + adminLockHash.String())

	// 2 - get UDT hash
	udtScript := &types.Script{
		CodeHash: types.HexToHash(sudtTestnetCodeHash),
		HashType: types.HashTypeType,
		Args:     common.FromHex(adminLockHash.String()),
	}
	udtHash, _ := udtScript.Hash()
	t.Log("udt hash: " + adminLockHash.String())

	// 3 - send udt script to chain if this SUDT is newly generated
	sudtIssuePayloadBuilder := model.NewBuildSudtIssueTransactionPayloadBuilder()
	sudtIssuePayloadBuilder.AddOwner(adminAddress)
	sudtIssuePayloadBuilder.AddTo(
		mode.HoldByFrom, model.NewToInfo(adminAddress, amount.CkbToShannon(issueUdtAmount)),
	)
	sudtIssuePayload := sudtIssuePayloadBuilder.Build()
	t.Log("build_sudt_issue_transactions payload: " + toJson(sudtIssuePayload))
	// 3.1 send transaction
	tx, err := constant.GetMercuryApiInstance().BuildSudtIssueTransaction(sudtIssuePayload)
	if err != nil {
		t.Error(err)
	}
	transaction := utils.Sign(tx)

	hash, err := constant.GetMercuryApiInstance().SendTransaction(context.Background(), transaction)
	if err != nil {
		t.Error(err)
	}
	t.Log("hash of transaction creating UDT script" + hash.String())

	item, _ := req.NewIdentityItemByCkb(constant.TEST_PUBKEY4)
	builder := model.NewBuildAdjustAccountPayloadBuilder()
	builder.AddItem(item)
	builder.AddAssetInfo(mcommon.NewUdtAsset(constant.UDT_HASH))
	builder.AddAccountNumber(1)

	// 4 - create acp cell for sudt receiver
	adjustAccountPayloadBuilder := model.NewBuildAdjustAccountPayloadBuilder()
	item, _ = req.NewIdentityItemByAddress(receiverAddress)
	adjustAccountPayloadBuilder.AddItem(item)
	adjustAccountPayloadBuilder.AddAssetInfo(mcommon.NewUdtAsset(udtHash.String()))
	adjustAccountPayloadBuilder.AddAccountNumber(1)
	payload := adjustAccountPayloadBuilder.Build()
	t.Log("adjust_account_transactions payload: " + toJson(payload))
	// 4.1 - send transaction
	tx, err = constant.GetMercuryApiInstance().BuildAdjustAccountTransaction(payload)
	if err != nil {
		t.Error(err)
	}
	transaction = utils.Sign(tx)
	hash, err = constant.GetMercuryApiInstance().SendTransaction(context.Background(), transaction)
	if err != nil {
		t.Error(err)
	}
	t.Log("hash of transaction creating ACP cell for receiver: " + hash.String())

	// 5.1 - issue SUDT to receiver
	sudtIssuePayloadBuilder = model.NewBuildSudtIssueTransactionPayloadBuilder()
	sudtIssuePayloadBuilder.AddOwner(adminAddress)
	sudtIssuePayloadBuilder.AddTo(
		mode.HoldByTo, model.NewToInfo(receiverAddress, amount.CkbToShannon(issueUdtAmount)),
	)
	sudtIssuePayload = sudtIssuePayloadBuilder.Build()
	t.Log("build_sudt_issue_transactions payload: " + toJson(sudtIssuePayload))
	// 3.1 send transaction
	tx, err = constant.GetMercuryApiInstance().BuildSudtIssueTransaction(sudtIssuePayload)
	if err != nil {
		t.Error(err)
	}
	transaction = utils.Sign(tx)
	hash, err = constant.GetMercuryApiInstance().SendTransaction(context.Background(), transaction)
	if err != nil {
		t.Error(err)
	}
	t.Log("hash of transaction issue UDT to receiver: " + toJson(hash))
}

func toJson(s interface{}) string {
	b, err := json.Marshal(s)
	if err != nil {
		return err.Error()
	}
	return string(b)
}
