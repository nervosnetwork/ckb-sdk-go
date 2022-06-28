package mercury

import (
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/common"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/req"
	"github.com/nervosnetwork/ckb-sdk-go/types"
	"github.com/stretchr/testify/assert"
	"math/big"
	"runtime/debug"
	"testing"
)

var c, _ = Dial("https://mercury-testnet.ckbapp.dev/0.4")

func TestBuildAdjustAccountTransaction(t *testing.T) {
	item, err := req.NewIdentityItemByPublicKeyHash("0xb0f8a32e7f9e8f3ab3a641f6eb02fcdb921d5589")
	checkError(t, err)
	from, err := req.NewIdentityItemByPublicKeyHash("0x202647fecc5b9d8cbdb4ae7167e40f5ab1e4baaf")
	checkError(t, err)
	payload := &model.BuildAdjustAccountPayload{
		Item:          item,
		From:          []*req.Item{from},
		AssetInfo:     common.NewUdtAsset1(types.HexToHash("0xf21e7350fa9518ed3cbb008e0e8c941d7e01a12181931d5608aa366ee22228bd")),
		AccountNumber: 1,
		ExtraCKB:      20000000000,
		FeeRate:       1000,
	}
	resp, err := c.BuildAdjustAccountTransaction(payload)
	checkError(t, err)
	assert.NotNil(t, resp.TxView)
	assert.NotNil(t, resp.ScriptGroups)
	assert.True(t, len(resp.ScriptGroups) >= 1)
}

func TestGetBalance(t *testing.T) {
	item, err := req.NewIdentityItemByPublicKeyHash("0x839f1806e85b40c13d3c73866045476cc9a8c214")
	checkError(t, err)
	payload := &model.GetBalancePayload{
		Item: item,
		AssetInfos: []*common.AssetInfo{
			common.NewUdtAsset1(types.HexToHash("0xf21e7350fa9518ed3cbb008e0e8c941d7e01a12181931d5608aa366ee22228bd"))},
		TipBlockNumber: 0,
	}
	resp, err := c.GetBalance(payload)
	checkError(t, err)
	assert.Equal(t, 2, len(resp.Balances))
	assert.Equal(t, types.HexToHash("0xf21e7350fa9518ed3cbb008e0e8c941d7e01a12181931d5608aa366ee22228bd"),
		resp.Balances[0].AssetInfo.UdtHash)
}

func TestBuildSudtIssueTransaction(t *testing.T) {
	address := "ckt1qzda0cr08m85hc8jlnfp3zer7xulejywt49kt2rr0vthywaa50xwsqg958atl2zdh8jn3ch8lc72nt0cf864ecqdxm9zf"
	item, err := req.NewIdentityItemByAddress(address)
	checkError(t, err)
	payload := &model.BuildSudtIssueTransactionPayload{
		Owner: address,
		From:  []*req.Item{item},
		To: []*model.ToInfo{
			{
				Address: "ckt1qzda0cr08m85hc8jlnfp3zer7xulejywt49kt2rr0vthywaa50xwsqg6flmrtx8y8tuu6s3jf2ahv4l6sjw9hsc3t4tqv",
				Amount:  big.NewInt(1),
			},
		},
		OutputCapacityProvider: model.OutputCapacityProviderFrom,
		FeeRate:                1000,
	}
	resp, err := c.BuildSudtIssueTransaction(payload)
	checkError(t, err)
	assert.NotNil(t, resp.TxView)
	assert.NotNil(t, resp.ScriptGroups)
}

func TestBuildSimpleTransferTransaction(t *testing.T) {
	payload := &model.SimpleTransferPayload{
		AssetInfo: common.NewCkbAsset(),
		From:      []string{"ckt1qzda0cr08m85hc8jlnfp3zer7xulejywt49kt2rr0vthywaa50xwsqfqyerlanzmnkxtmd9ww9n7gr66k8jt4tclm9jnk"},
		To: []*model.ToInfo{
			{
				Address: "ckt1qzda0cr08m85hc8jlnfp3zer7xulejywt49kt2rr0vthywaa50xwsqg958atl2zdh8jn3ch8lc72nt0cf864ecqdxm9zf",
				Amount:  big.NewInt(10000000000),
			},
		},
		FeeRate: 500,
	}
	resp, err := c.BuildSimpleTransferTransaction(payload)
	checkError(t, err)
	assert.NotNil(t, resp.TxView)
	assert.NotNil(t, resp.ScriptGroups)
}

func TestBuildTransferTransaction(t *testing.T) {
	address := "ckt1qzda0cr08m85hc8jlnfp3zer7xulejywt49kt2rr0vthywaa50xwsqfqyerlanzmnkxtmd9ww9n7gr66k8jt4tclm9jnk"
	item, err := req.NewAddressItem(address)
	checkError(t, err)
	payload := &model.TransferPayload{
		AssetInfo: common.NewCkbAsset(),
		From:      []*req.Item{item},
		To: []*model.ToInfo{
			{
				Address: "ckt1qzda0cr08m85hc8jlnfp3zer7xulejywt49kt2rr0vthywaa50xwsqg958atl2zdh8jn3ch8lc72nt0cf864ecqdxm9zf",
				Amount:  big.NewInt(100),
			},
		},
		PayFee:  model.PayFeeFrom,
		FeeRate: 1100,
	}
	resp, err := c.BuildTransferTransaction(payload)
	checkError(t, err)
	assert.NotNil(t, resp.TxView)
	assert.NotNil(t, resp.ScriptGroups)
}

func TestBuildDaoDepositTransaction(t *testing.T) {
	from, err := req.NewAddressItem("ckt1qzda0cr08m85hc8jlnfp3zer7xulejywt49kt2rr0vthywaa50xwsqfqyerlanzmnkxtmd9ww9n7gr66k8jt4tclm9jnk")
	checkError(t, err)
	payload := &model.DaoDepositPayload{
		From:    []*req.Item{from},
		To:      "ckt1qzda0cr08m85hc8jlnfp3zer7xulejywt49kt2rr0vthywaa50xwsqvrnuvqd6zmgrqn60rnsesy23mvex5vy9q0g8hfd",
		Amount: 20000000000,
		FeeRate: 1100,
	}
	resp, err := c.BuildDaoDepositTransaction(payload)
	checkError(t, err)
	assert.NotNil(t, resp.TxView)
	assert.NotNil(t, resp.ScriptGroups)
}

func TestBuildDaoWithdrawTransaction(t *testing.T) {
	from, err := req.NewAddressItem("ckt1qzda0cr08m85hc8jlnfp3zer7xulejywt49kt2rr0vthywaa50xwsqfqyerlanzmnkxtmd9ww9n7gr66k8jt4tclm9jnk")
	checkError(t, err)
	payload := &model.DaoWithdrawPayload{
		From:    []*req.Item{from},
		FeeRate: 1100,
	}
	resp, err := c.BuildDaoWithdrawTransaction(payload)
	checkError(t, err)
	assert.NotNil(t, resp.TxView)
	assert.NotNil(t, resp.ScriptGroups)
}

func TestBuildDaoClaimTransaction(t *testing.T) {
	from, err := req.NewAddressItem("ckt1qzda0cr08m85hc8jlnfp3zer7xulejywt49kt2rr0vthywaa50xwsqfqyerlanzmnkxtmd9ww9n7gr66k8jt4tclm9jnk")
	checkError(t, err)
	payload := &model.DaoClaimPayload{
		From:    []*req.Item{from},
		FeeRate: 1100,
	}
	resp, err := c.BuildDaoClaimTransaction(payload)
	checkError(t, err)
	assert.NotNil(t, resp.TxView)
	assert.NotNil(t, resp.ScriptGroups)
}

func TestGetAccountInfo(t *testing.T) {
	item, err := req.NewAddressItem("ckt1qzda0cr08m85hc8jlnfp3zer7xulejywt49kt2rr0vthywaa50xwsq06y24q4tc4tfkgze35cc23yprtpzfrzygljdjh9")
	checkError(t, err)
	payload := &model.GetAccountInfoPayload{
		Item:      item,
		AssetInfo: common.NewUdtAsset1(types.HexToHash("0xf21e7350fa9518ed3cbb008e0e8c941d7e01a12181931d5608aa366ee22228bd")),
	}
	resp, err := c.GetAccountInfo(payload)
	checkError(t, err)
	// asset not nil
	assert.NotEqual(t, "", resp.AccountAddress)
	assert.NotEqual(t, "", resp.AccountType)
}

func TestGetBlockInfoByNumber(t *testing.T) {
	payload := &model.GetBlockInfoPayload{
		BlockNumber: 2172093,
	}
	resp, err := c.GetBlockInfo(payload)
	checkError(t, err)
	assert.NotNil(t, resp)
	assert.NotEqual(t, types.Hash{}, resp.ParentHash)
	assert.NotEqual(t, types.Hash{}, resp.BlockHash)
	assert.Equal(t, 3, len(resp.Transactions))
}

func TestGetBlockInfoByHash(t *testing.T) {
	payload := &model.GetBlockInfoPayload{
		BlockHash: types.HexToHash("0xee8adba356105149cb9dc1cb0d09430a6bd01182868787ace587961c0d64e742"),
	}
	resp, err := c.GetBlockInfo(payload)
	checkError(t, err)
	assert.NotNil(t, resp)
	assert.NotEqual(t, types.Hash{}, resp.ParentHash)
	assert.NotEqual(t, types.Hash{}, resp.BlockHash)
	assert.Equal(t, 3, len(resp.Transactions))
}

func TestGetDbInfo(t *testing.T) {
	resp, err := c.GetDbInfo()
	checkError(t, err)
	assert.NotEqual(t, "", resp.Version)
	assert.NotEqual(t, 0, resp.ConnSize)
}

func TestMercuryInfo(t *testing.T) {
	resp, err := c.GetMercuryInfo()
	checkError(t, err)
	assert.NotEqual(t, "", resp.MercuryVersion)
	assert.NotEqual(t, "", resp.CkbNodeVersion)
	assert.NotEqual(t, "", resp.NetworkType)
}

func checkError(t *testing.T, err error) {
	if err != nil {
		t.Error(err, string(debug.Stack()))
	}
}
