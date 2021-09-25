package mercury

import (
	"context"
	"errors"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/types/req"
	resp2 "github.com/nervosnetwork/ckb-sdk-go/mercury/model/types/resp"
	"math/big"

	"github.com/ethereum/go-ethereum/rpc"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/action"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/source"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/types"
	"github.com/nervosnetwork/ckb-sdk-go/utils/amount"
)

type Client interface {
	GetBalance(payload *req.GetBalancePayload) (*resp2.GetBalanceResponse, error)
	BuildTransferTransaction(payload *req.TransferPayload) (*resp2.TransferCompletionResponse, error)
	BuildSmartTransferTransaction(payload *req.SmartTransferPayload) (*resp2.TransferCompletionResponse, error)
	BuildAdjustAccountTransaction(payload *req.AdjustAccountPayload) (*resp2.TransferCompletionResponse, error)
	BuildAssetCollectionTransaction(payload *req.CollectAssetPayload) (*resp2.TransferCompletionResponse, error)
	RegisterAddresses(normalAddresses []string) ([]string, error)
	GetTransactionInfo(txHash string) (*resp2.GetTransactionInfoResponse, error)
	GetBlockInfo(payload *req.GetBlockInfoPayload) (*resp2.GetBlockInfoResponse, error)
	QueryGenericTransactions(payload *req.QueryGenericTransactionsPayload) (*resp2.QueryGenericTransactionsResponse, error)
	GetAccountNumber(address string) (uint, error)
	GetDbInfo() (*resp2.DBDriver, error)
	GetMercuryInfo() (*resp2.MercuryInfo, error)
}

type client struct {
	c *rpc.Client
}

func (cli *client) GetDbInfo() (*resp2.DBDriver, error) {
	var resp resp2.DBDriver
	err := cli.c.Call(&resp, "get_balance")
	if err != nil {
		return nil, err
	}

	return &resp, err
}

func (cli *client) GetMercuryInfo() (*resp2.MercuryInfo, error) {
	var resp resp2.MercuryInfo
	err := cli.c.Call(&resp, "get_balance")
	if err != nil {
		return nil, err
	}

	return &resp, err
}

func (cli *client) GetBalance(payload *req.GetBalancePayload) (*resp2.GetBalanceResponse, error) {
	var resp resp2.GetBalanceResponse
	err := cli.c.Call(&resp, "get_balance", payload)
	if err != nil {
		return nil, err
	}

	return &resp, err
}

func (cli *client) BuildTransferTransaction(payload *req.TransferPayload) (*resp2.TransferCompletionResponse, error) {
	var resp resp2.TransferCompletionResponse
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

func (cli *client) BuildSmartTransferTransaction(payload *req.SmartTransferPayload) (*resp2.TransferCompletionResponse, error) {
	transferPayload, err := cli.toTransferPayload(payload)
	if err != nil {
		return nil, err
	}

	return cli.BuildTransferTransaction(transferPayload)
}

func (cli *client) BuildAdjustAccountTransaction(payload *req.AdjustAccountPayload) (*resp2.TransferCompletionResponse, error) {
	var resp resp2.TransferCompletionResponse
	err := cli.c.Call(&resp, "build_asset_account_creation_transaction", payload)
	if err != nil {
		return &resp, err
	}

	return &resp, err
}

func (cli *client) BuildAssetCollectionTransaction(payload *req.CollectAssetPayload) (*resp2.TransferCompletionResponse, error) {
	var resp resp2.TransferCompletionResponse

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

func (cli *client) GetBlockInfo(payload *req.GetBlockInfoPayload) (*resp2.GetBlockInfoResponse, error) {
	var resp resp2.GetBlockInfoResponse
	err := cli.c.Call(&resp, "get_block_info", payload)
	if err != nil {
		return nil, err
	}

	return &resp, err
}

func (cli *client) GetTransactionInfo(txHash string) (*resp2.GetTransactionInfoResponse, error) {
	var resp *resp2.GetTransactionInfoResponse
	err := cli.c.Call(&resp, "get_transaction_info", txHash)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (cli *client) QueryGenericTransactions(payload *req.QueryGenericTransactionsPayload) (*resp2.QueryGenericTransactionsResponse, error) {
	var queryGenericTransactionsResponse resp2.QueryGenericTransactionsResponse
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

func (cli *client) toTransferPayload(payload *req.SmartTransferPayload) (*req.TransferPayload, error) {
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

	builder := req.NewTransferBuilder()
	if payload.AssetInfo.AssetType == types.UDT {
		builder.UdtHash = payload.AssetInfo.UdtHash
	}
	builder.AddFromKeyAddresses(payload.From, source)
	for _, to := range payload.To {
		if payload.AssetInfo.AssetType == types.UDT {
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

func (cli *client) getSource(fromBalances []*resp2.GetBalanceResponse, to []*req.SmartTo, assetType types.AssetType) string {
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

func (cli *client) feePay(fromBalances, toBalances []*resp2.GetBalanceResponse, payload *req.SmartTransferPayload) error {
	from := cli.getBalanceByAssetTypeAndBalanceType(fromBalances, types.CKB, "free")
	to := cli.getBalanceByAssetTypeAndBalanceType(toBalances, types.CKB, "free")

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

func (cli *client) getBalanceByAssetTypeAndBalanceType(balances []*resp2.GetBalanceResponse, assetType types.AssetType, balanceType string) *big.Int {
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

func (cli *client) getBalance(addresses []string, assetInfo *types.AssetInfo) ([]*resp2.GetBalanceResponse, error) {
	result := make([]*resp2.GetBalanceResponse, len(addresses))
	for i, address := range addresses {
		builder := req.NewGetBalancePayloadBuilder()
		builder.SetItemAsAddress(address)
		builder.AddAssetInfo(types.NewCkbAsset())
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
