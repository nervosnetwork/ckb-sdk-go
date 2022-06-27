package mercury

import (
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/common"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/req"
	"github.com/nervosnetwork/ckb-sdk-go/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

var c, _ = Dial("https://mercury-testnet.ckbapp.dev/0.4")

func TestBuildAdjustAccountTransaction(t *testing.T) {
	item, _ := req.NewIdentityItemByPublicKeyHash("0xb0f8a32e7f9e8f3ab3a641f6eb02fcdb921d5589")
	from, _ := req.NewIdentityItemByPublicKeyHash("0x202647fecc5b9d8cbdb4ae7167e40f5ab1e4baaf")
	payload := &model.BuildAdjustAccountPayload{
		Item:          item,
		From:          []*req.Item{from},
		AssetInfo:     common.NewUdtAsset1(types.HexToHash("0xf21e7350fa9518ed3cbb008e0e8c941d7e01a12181931d5608aa366ee22228bd")),
		AccountNumber: 1,
		ExtraCKB:      20000000000,
		FeeRate:       1000,
	}
	tx, _ := c.BuildAdjustAccountTransaction(payload)

	assert.NotNil(t, tx.TxView)
	assert.NotNil(t, tx.ScriptGroups)
	assert.True(t, len(tx.ScriptGroups) >= 1)
}

func TestGetBalance(t *testing.T) {
	item, _ := req.NewIdentityItemByPublicKeyHash("0x839f1806e85b40c13d3c73866045476cc9a8c214")
	payload := &model.GetBalancePayload{
		Item: item,
		AssetInfos: []*common.AssetInfo{
			common.NewUdtAsset1(types.HexToHash("0xf21e7350fa9518ed3cbb008e0e8c941d7e01a12181931d5608aa366ee22228bd"))},
		TipBlockNumber: 0,
	}
	resp, _ := c.GetBalance(payload)
	assert.Equal(t, 2, len(resp.Balances))
	assert.Equal(t, types.HexToHash("0xf21e7350fa9518ed3cbb008e0e8c941d7e01a12181931d5608aa366ee22228bd"),
		resp.Balances[0].AssetInfo.UdtHash)
}
