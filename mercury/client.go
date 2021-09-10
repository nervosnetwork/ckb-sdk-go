package mercury

import (
	"context"
	"errors"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/action"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/common"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/resp"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/source"
	"github.com/nervosnetwork/ckb-sdk-go/utils/amount"
	"math/big"
)

type Client interface {
	GetBalance(payload *model.GetBalancePayload) (*resp.GetBalanceResponse, error)
	BuildTransferTransaction(payload *model.TransferPayload) (*resp.TransferCompletionResponse, error)
	BuildSmartTransferTransaction(payload *model.SmartTransferPayload) (*resp.TransferCompletionResponse, error)
	BuildAdjustAccountTransaction(payload *model.AdjustAccountPayload) (*resp.TransferCompletionResponse, error)
	BuildAssetCollectionTransaction(payload *model.CollectAssetPayload) (*resp.TransferCompletionResponse, error)
	RegisterAddresses(normalAddresses []string) ([]string, error)
	GetTransactionInfo(txHash string) (*resp.TransactionInfoWithStatusResponse, error)
	GetBlockInfo(payload *model.GetBlockInfoPayload) (*resp.BlockInfoResponse, error)
	QueryGenericTransactions(payload *model.QueryGenericTransactionsPayload) (*resp.QueryGenericTransactionsResponse, error)
	GetAccountNumber(address string) (uint, error)
}

type client struct {
	c *rpc.Client
}

func (cli *client) GetBalance(payload *model.GetBalancePayload) (*resp.GetBalanceResponse, error) {
	var balance rpcGetBalanceResponse
	err := cli.c.Call(&balance, "get_balance", payload)
	if err != nil {
		return nil, err
	}

	result := &resp.GetBalanceResponse{}
	for _, balanceResp := range balance.Balances {
		var asset *common.AssetInfo
		if balanceResp.UdtHash == "" {
			asset = common.NewCkbAsset()
		} else {
			asset = common.NewUdtAsset(balanceResp.UdtHash)
		}

		free := new(big.Int)
		free.SetString(balanceResp.Unconstrained, 0)

		claimable := new(big.Int)
		claimable.SetString(balanceResp.Fleeting, 0)

		freezed := new(big.Int)
		freezed.SetString(balanceResp.Locked, 0)

		result.Balances = append(result.Balances, &resp.BalanceResp{
			Address:   balanceResp.KeyAddress,
			AssetInfo: asset,
			Free:      free,
			Claimable: claimable,
			Freezed:   freezed,
		})
	}

	return result, err
}

func (cli *client) BuildTransferTransaction(payload *model.TransferPayload) (*resp.TransferCompletionResponse, error) {
	var resp resp.TransferCompletionResponse
	if payload.UdtHash == "" {
		for _, item := range payload.Items {
			if !item.To.IsPayBayFrom() || !item.To.IsPayBayFrom() {
				return &resp, errors.New("The transaction does not support ckb")
			}
		}
	}

	err := cli.c.Call(&resp, "build_transfer_transaction", payload)
	if err != nil {
		return &resp, err
	}

	return &resp, err
}

func (cli *client) BuildSmartTransferTransaction(payload *model.SmartTransferPayload) (*resp.TransferCompletionResponse, error) {
	transferPayload, err := cli.toTransferPayload(payload)
	if err != nil {
		return nil, err
	}

	return cli.BuildTransferTransaction(transferPayload)
}

func (cli *client) BuildAdjustAccountTransaction(payload *model.AdjustAccountPayload) (*resp.TransferCompletionResponse, error) {
	var resp resp.TransferCompletionResponse
	err := cli.c.Call(&resp, "build_asset_account_creation_transaction", payload)
	if err != nil {
		return &resp, err
	}

	return &resp, err
}

func (cli *client) BuildAssetCollectionTransaction(payload *model.CollectAssetPayload) (*resp.TransferCompletionResponse, error) {
	var resp resp.TransferCompletionResponse

	err := cli.c.Call(&resp, "build_asset_collection_transaction", payload)
	if err != nil {
		return &resp, err
	}

	return &resp, err
}

func (cli *client) RegisterAddresses(normalAddresses []string) ([]string, error) {
	var scriptHash []string
	err := cli.c.Call(&scriptHash, "register_addresses", normalAddresses)
	if err != nil {
		return scriptHash, err
	}

	return scriptHash, err
}

func (cli *client) GetBlockInfo(payload *model.GetBlockInfoPayload) (*resp.BlockInfoResponse, error) {
	var block *rpcBlockInfoResponse
	err := cli.c.Call(&block, "get_generic_block", payload)
	if err != nil {
		return nil, err
	}

	result := resp.BlockInfoResponse{
		BlockNumber:     block.BlockNumber,
		BlockHash:       block.BlockHash,
		ParentBlockHash: block.ParentBlockHash,
		Timestamp:       block.Timestamp,
	}

	for _, transaction := range block.Transactions {
		tx, err := toTransactionInfoResponse(transaction.Operations, transaction.TxHash)
		if err != nil {
			return nil, err
		}
		result.Transactions = append(result.Transactions, tx)
	}

	return &result, err
}

func (cli *client) GetTransactionInfo(txHash string) (*resp.TransactionInfoWithStatusResponse, error) {
	var tx *rpcTransactionInfoWithStatusResponse
	err := cli.c.Call(&tx, "get_generic_transaction", txHash)
	if err != nil {
		return nil, err
	}

	result, err := toTransactionInfoWithStatusResponse(tx)
	if err != nil {
		return nil, err
	}

	return result, err
}

func (cli *client) QueryGenericTransactions(payload *model.QueryGenericTransactionsPayload) (*resp.QueryGenericTransactionsResponse, error) {
	var queryGenericTransactionsResponse resp.QueryGenericTransactionsResponse
	err := cli.c.Call(&queryGenericTransactionsResponse, "query_generic_transactions", payload)
	if err != nil {
		return &queryGenericTransactionsResponse, err
	}

	return &queryGenericTransactionsResponse, err
}

func (cli *client) GetAccountNumber(address string) (uint, error) {
	var number uint
	err := cli.c.Call(&number, "get_account_number", address)
	if err != nil {
		return 0, err
	}

	return number, err
}

func (cli *client) toTransferPayload(payload *model.SmartTransferPayload) (*model.TransferPayload, error) {
	fromBalances, err := cli.getBalance(payload.From, payload.AssetInfo)
	if err != nil {
		return nil, err
	}

	toAddresses := make([]string, len(payload.To))
	for i, to := range payload.To {
		toAddresses[i] = to.Address
	}
	toBalances, err := cli.getBalance(toAddresses, payload.AssetInfo)
	if err != nil {
		return nil, err
	}

	err = cli.feePay(fromBalances, toBalances, payload)
	if err != nil {
		return nil, err
	}

	source := cli.getSource(fromBalances, payload.To, payload.AssetInfo.AssetType)

	builder := model.NewTransferBuilder()
	if payload.AssetInfo.AssetType == common.Udt {
		builder.UdtHash = payload.AssetInfo.UdtHash
	}
	builder.AddFromKeyAddresses(payload.From, source)
	for _, to := range payload.To {
		if payload.AssetInfo.AssetType == common.Udt {
			number, err := cli.GetAccountNumber(to.Address)
			if err != nil {
				return nil, err
			}
			if number > 0 {
				builder.AddToKeyAddressItem(to.Address, action.Pay_by_to, to.Amount)
			} else {
				builder.AddToKeyAddressItem(to.Address, action.Pay_by_from, to.Amount)
			}
		} else {
			builder.AddToKeyAddressItem(to.Address, action.Pay_by_from, to.Amount)
		}
	}

	return builder.Build(), nil
}

func (cli *client) getSource(fromBalances []*resp.GetBalanceResponse, to []*model.SmartTo, assetType common.AssetType) string {
	fromBalance := cli.getBalanceByAssetTypeAndBalanceType(fromBalances, assetType, "claimable")
	totalAmount := big.NewInt(0)
	for _, smartTo := range to {
		totalAmount = totalAmount.Add(totalAmount, smartTo.Amount)
	}

	if fromBalance.Cmp(totalAmount) >= 0 {
		return source.Fleeting
	} else {
		return source.Unconstrained
	}
}

func (cli *client) feePay(fromBalances, toBalances []*resp.GetBalanceResponse, payload *model.SmartTransferPayload) error {
	from := cli.getBalanceByAssetTypeAndBalanceType(fromBalances, common.Ckb, "free")
	to := cli.getBalanceByAssetTypeAndBalanceType(toBalances, common.Ckb, "free")

	feeThreshold := amount.CkbWithDecimalToShannon(0.0001)

	if from.Cmp(feeThreshold) < 0 && to.Cmp(feeThreshold) < 0 {
		return errors.New("CKB Insufficient balance to pay the fee")
	}
	if from.Cmp(feeThreshold) < 0 && to.Cmp(feeThreshold) >= 0 {
		for _, getBalanceResponse := range toBalances {
			for _, balance := range getBalanceResponse.Balances {
				payload.From = append(payload.From, balance.Address)
			}
		}
	}

	return nil
}

func (cli *client) getBalanceByAssetTypeAndBalanceType(balances []*resp.GetBalanceResponse, assetType common.AssetType, balanceType string) *big.Int {
	amount := big.NewInt(0)
	for _, getBalanceResponse := range balances {
		for _, balance := range getBalanceResponse.Balances {
			if balance.AssetInfo.AssetType == assetType {
				if balanceType == "free" {
					amount = amount.Add(amount, balance.Free)
				}

				if balanceType == "claimable" {
					amount = amount.Add(amount, balance.Claimable)
				}

				if balanceType == "freezed" {
					amount = amount.Add(amount, balance.Freezed)
				}
			}
		}
	}
	return amount
}

func (cli *client) getBalance(addresses []string, assetInfo *common.AssetInfo) ([]*resp.GetBalanceResponse, error) {
	result := make([]*resp.GetBalanceResponse, len(addresses))
	for i, address := range addresses {
		builder := model.NewGetBalancePayloadBuilder()
		builder.AddAddress(address)
		builder.AddAssetInfo(common.NewCkbAsset())
		if assetInfo != nil {
			builder.AddAssetInfo(assetInfo)
		}

		balance, err := cli.GetBalance(builder.Build())
		if err != nil {
			return nil, err
		}

		result[i] = balance
	}

	return result, nil
}

func Dial(url string) (Client, error) {
	return DialContext(context.Background(), url)
}

func DialContext(ctx context.Context, url string) (Client, error) {
	c, err := rpc.DialContext(ctx, url)
	if err != nil {
		return nil, err
	}
	return NewClient(c), nil
}

func NewClient(c *rpc.Client) Client {
	return &client{
		c: c,
	}
}
