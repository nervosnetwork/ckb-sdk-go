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
		Amount:  20000000000,
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

func TestGetSpentTransactionWithTransactionView(t *testing.T) {
	payload := &model.GetSpentTransactionPayload{
		OutPoint: types.OutPoint{
			TxHash: types.HexToHash("0xb2e952a30656b68044e1d5eed69f1967347248967785449260e3942443cbeece"),
			Index:  1,
		},
	}
	resp, err := c.GetSpentTransactionWithTransactionView(payload)
	checkError(t, err)
	assert.NotNil(t, resp.Value.Transaction)
	assert.Equal(t, types.HexToHash("0x407033c3baa6104c9f46d3c7948b812274556148d74b0db251f50fc6e7507233"), resp.Value.TxStatus.BlockHash)
	assert.Equal(t, types.TransactionStatusCommitted, resp.Value.TxStatus.Status)
	assert.Equal(t, uint64(0x17bc67c4078), resp.Value.TxStatus.Timestamp)
}

func TestGetSpentTransactionWithTransactionInfo(t *testing.T) {
	payload := &model.GetSpentTransactionPayload{
		OutPoint: types.OutPoint{
			TxHash: types.HexToHash("0xb2e952a30656b68044e1d5eed69f1967347248967785449260e3942443cbeece"),
			Index:  1,
		},
	}
	resp, err := c.GetSpentTransactionWithTransactionInfo(payload)
	checkError(t, err)
	assert.Equal(t, types.HexToHash("0x2c4e242e034e70a7b8ae5f899686c256dad2a816cc36ddfe2c1460cbbbbaaaed"), resp.Value.TxHash)
	assert.Equal(t, 3, len(resp.Value.Records))
	assert.Equal(t, big.NewInt(0xd9ac33e984), resp.Value.Records[0].Amount)
	assert.Equal(t, uint64(0x2877b6), resp.Value.Records[0].BlockNumber)
	assert.Equal(t, uint64(0x70804bf000af6), resp.Value.Records[0].EpochNumber)
	assert.Equal(t, 3, len(resp.Value.Records))
	assert.Equal(t, uint64(0x1f5), resp.Value.Fee)
	assert.Equal(t, uint64(0x17bc67c4078), resp.Value.Timestamp)
}

func TestQueryTransactions(t *testing.T) {
	item, err := req.NewAddressItem("ckt1qzda0cr08m85hc8jlnfp3zer7xulejywt49kt2rr0vthywaa50xwsqg6flmrtx8y8tuu6s3jf2ahv4l6sjw9hsc3t4tqv")
	checkError(t, err)
	payload := &model.QueryTransactionsPayload{
		Item:       item,
		AssetInfos: []*common.AssetInfo{common.NewCkbAsset()},
		BlockRange: &model.BlockRange{
			From: 2778100,
			To:   3636218,
		},
		Pagination: &model.PaginationRequest{
			Order:       model.ASC,
			Limit:       2,
			ReturnCount: true,
		},
	}
	resp1, err := c.QueryTransactionsWithTransactionView(payload)
	checkError(t, err)
	assert.NotNil(t, uint64(0x02), resp1.Count)
	assert.NotNil(t, 2, len(resp1.Response))
	resp2, err := c.QueryTransactionsWithTransactionInfo(payload)
	checkError(t, err)
	assert.NotNil(t, uint64(0x02), resp2.Count)
	assert.NotNil(t, 2, len(resp2.Response))
}

func TestQueryTransactionsWithPage(t *testing.T) {
	item, err := req.NewIdentityItemByPublicKeyHash("0x1a4ff63598e43af9cd42324abb7657fa849c5bc3")
	checkError(t, err)
	payload := &model.QueryTransactionsPayload{
		Item:       item,
		AssetInfos: []*common.AssetInfo{},
		Pagination: &model.PaginationRequest{
			Order:       model.DESC,
			Limit:       1,
			ReturnCount: true,
		},
	}
	resp, err := c.QueryTransactionsWithTransactionView(payload)
	checkError(t, err)
	assert.Equal(t, 1, len(resp.Response))

	payload = &model.QueryTransactionsPayload{
		Item:       item,
		AssetInfos: []*common.AssetInfo{},
		Pagination: &model.PaginationRequest{
			Cursor:      resp.NextCursor,
			Order:       model.DESC,
			Limit:       2,
			ReturnCount: true,
		},
	}
	resp, err = c.QueryTransactionsWithTransactionView(payload)
	checkError(t, err)
	assert.Equal(t, 2, len(resp.Response))
	assert.Equal(t, types.HexToHash("0x88638e32403336912f8387ab5298ac3d3e1588082361d2fc0840808671467e54"),
		resp.Response[0].Value.Transaction.Hash)
	assert.Equal(t, types.HexToHash("0xeedfaf24add85ceea295b46a30c0b0c88bb5006edbbddb069092eb39f77a0f66"),
		resp.Response[1].Value.Transaction.Hash)
}

func TestGetTransactionInfo(t *testing.T) {
	resp, err := c.GetTransactionInfo(types.HexToHash("0x4329e4c751c95384a51072d4cbc9911a101fd08fc32c687353d016bf38b8b22c"))
	checkError(t, err)
	assert.NotNil(t, resp.Transaction)
	assert.Equal(t, types.HexToHash("0x4329e4c751c95384a51072d4cbc9911a101fd08fc32c687353d016bf38b8b22c"), resp.Transaction.TxHash)
	assert.Equal(t, types.TransactionStatusCommitted, resp.Status)
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

func TestSyncStateInfo(t *testing.T) {
	resp, err := c.GetSyncState()
	checkError(t, err)
	assert.NotEqual(t, "", resp.State)
}

func checkError(t *testing.T, err error) {
	if err != nil {
		t.Error(err, string(debug.Stack()))
	}
}
