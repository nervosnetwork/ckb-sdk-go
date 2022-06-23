package mercury

import (
	"context"
	"github.com/nervosnetwork/ckb-sdk-go/types"

	"github.com/ethereum/go-ethereum/rpc"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/common"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/resp"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/source"
	"github.com/pkg/errors"
)

type Client interface {
	GetBalance(payload *model.GetBalancePayload) (*resp.GetBalanceResponse, error)
	BuildTransferTransaction(payload *model.TransferPayload) (*resp.TransferCompletionResponse, error)
	BuildSimpleTransferTransaction(payload *model.SimpleTransferPayload) (*resp.TransferCompletionResponse, error)
	BuildAdjustAccountTransaction(*model.BuildAdjustAccountPayload) (*resp.TransferCompletionResponse, error)
	BuildSudtIssueTransaction(payload *model.BuildSudtIssueTransactionPayload) (*resp.TransferCompletionResponse, error)
	RegisterAddresses(normalAddresses []string) ([]string, error)
	GetTransactionInfo(txHash types.Hash) (*resp.GetTransactionInfoResponse, error)
	GetSpentTransactionWithTransactionInfo(*model.GetSpentTransactionPayload) (*resp.TransactionInfoWrapper, error)
	GetSpentTransactionWithTransactionView(*model.GetSpentTransactionPayload) (*resp.TransactionViewWrapper, error)
	GetBlockInfo(payload *model.GetBlockInfoPayload) (*resp.BlockInfo, error)
	GetAccountInfo(payload *model.GetAccountInfoPayload) (*resp.AccountInfo, error)
	QueryTransactionsWithTransactionInfo(payload *model.QueryTransactionsPayload) (*resp.PaginationResponseTransactionInfo, error)
	QueryTransactionsWithTransactionView(payload *model.QueryTransactionsPayload) (*resp.PaginationResponseTransactionView, error)
	GetDbInfo() (*resp.DBInfo, error)
	GetMercuryInfo() (*resp.MercuryInfo, error)
	GetSyncState() (*resp.MercurySyncState, error)
	BuildDaoDepositTransaction(payload *model.DaoDepositPayload) (*resp.TransferCompletionResponse, error)
	BuildDaoWithdrawTransaction(payload *model.DaoWithdrawPayload) (*resp.TransferCompletionResponse, error)
	BuildDaoClaimTransaction(payload *model.DaoClaimPayload) (*resp.TransferCompletionResponse, error)
}
type client struct {
	c *rpc.Client
}

func (cli *client) BuildDaoDepositTransaction(payload *model.DaoDepositPayload) (*resp.TransferCompletionResponse, error) {
	var resp resp.TransferCompletionResponse
	err := cli.c.Call(&resp, "build_dao_deposit_transaction", payload)
	if err != nil {
		return nil, err
	}

	return &resp, err
}

func (cli *client) BuildDaoWithdrawTransaction(payload *model.DaoWithdrawPayload) (*resp.TransferCompletionResponse, error) {
	var resp resp.TransferCompletionResponse
	err := cli.c.Call(&resp, "build_dao_withdraw_transaction", payload)
	if err != nil {
		return nil, err
	}

	return &resp, err
}

func (cli *client) BuildDaoClaimTransaction(payload *model.DaoClaimPayload) (*resp.TransferCompletionResponse, error) {
	var resp resp.TransferCompletionResponse
	err := cli.c.Call(&resp, "build_dao_claim_transaction", payload)
	if err != nil {
		return nil, err
	}

	return &resp, err
}

func (cli *client) GetDbInfo() (*resp.DBInfo, error) {
	var resp resp.DBInfo
	err := cli.c.Call(&resp, "get_db_info")
	if err != nil {
		return nil, err
	}

	return &resp, err
}

func (cli *client) GetMercuryInfo() (*resp.MercuryInfo, error) {
	var resp resp.MercuryInfo
	err := cli.c.Call(&resp, "get_mercury_info")
	if err != nil {
		return nil, err
	}

	return &resp, err
}

func (cli *client) GetSyncState() (*resp.MercurySyncState, error) {
	var resp resp.MercurySyncState
	err := cli.c.Call(&resp, "get_sync_state")
	if err != nil {
		return nil, err
	}

	return &resp, err
}

func (cli *client) GetBalance(payload *model.GetBalancePayload) (*resp.GetBalanceResponse, error) {
	var balance resp.GetBalanceResponse
	err := cli.c.Call(&balance, "get_balance", payload)
	if err != nil {
		return nil, err
	}

	return &balance, err
}

func (cli *client) BuildTransferTransaction(payload *model.TransferPayload) (*resp.TransferCompletionResponse, error) {
	if payload.From.Source == source.Claimable && payload.AssetInfo.AssetType == common.CKB {
		return nil, errors.New("The transaction does not support ckb")
	}

	var resp resp.TransferCompletionResponse
	err := cli.c.Call(&resp, "build_transfer_transaction", payload)
	if err != nil {
		return &resp, err
	}

	return &resp, err
}

func (cli *client) BuildSimpleTransferTransaction(payload *model.SimpleTransferPayload) (*resp.TransferCompletionResponse, error) {
	var resp resp.TransferCompletionResponse
	err := cli.c.Call(&resp, "build_simple_transfer_transaction", payload)
	if err != nil {
		return &resp, err
	}

	return &resp, err
}

func (cli *client) BuildAdjustAccountTransaction(payload *model.BuildAdjustAccountPayload) (*resp.TransferCompletionResponse, error) {
	var resp resp.TransferCompletionResponse
	err := cli.c.Call(&resp, "build_adjust_account_transaction", payload)
	if err != nil {
		return &resp, err
	}

	return &resp, err
}

func (cli *client) BuildSudtIssueTransaction(payload *model.BuildSudtIssueTransactionPayload) (*resp.TransferCompletionResponse, error) {
	var resp resp.TransferCompletionResponse
	err := cli.c.Call(&resp, "build_sudt_issue_transaction", payload)
	if err != nil {
		return &resp, err
	}

	return &resp, err
}

func (cli *client) RegisterAddresses(normalAddresses []string) ([]string, error) {
	var scriptHash []string
	err := cli.c.Call(&scriptHash, "register_address", normalAddresses)
	if err != nil {
		return scriptHash, err
	}

	return scriptHash, err
}

func (cli *client) GetBlockInfo(payload *model.GetBlockInfoPayload) (*resp.BlockInfo, error) {
	var block resp.BlockInfo
	err := cli.c.Call(&block, "get_block_info", payload)
	if err != nil {
		return nil, err
	}

	return &block, err
}

func (cli *client) GetAccountInfo(payload *model.GetAccountInfoPayload) (*resp.AccountInfo, error) {
	var account resp.AccountInfo
	err := cli.c.Call(&account, "get_account_info", payload)
	if err != nil {
		return nil, err
	}

	return &account, err
}

func (cli *client) GetTransactionInfo(txHash types.Hash) (*resp.GetTransactionInfoResponse, error) {
	var tx *resp.GetTransactionInfoResponse
	err := cli.c.Call(&tx, "get_transaction_info", txHash)
	if err != nil {
		return nil, err
	}
	return tx, err
}

func (cli *client) GetSpentTransactionWithTransactionInfo(payload *model.GetSpentTransactionPayload) (*resp.TransactionInfoWrapper, error) {
	payload.StructureType = model.DoubleEntry
	var tx *resp.TransactionInfoWrapper
	err := cli.c.Call(&tx, "get_spent_transaction", payload)
	if err != nil {
		return nil, err
	}
	return tx, err
}

func (cli *client) GetSpentTransactionWithTransactionView(payload *model.GetSpentTransactionPayload) (*resp.TransactionViewWrapper, error) {
	payload.StructureType = model.Native
	var tx *resp.TransactionViewWrapper
	err := cli.c.Call(&tx, "get_spent_transaction", payload)
	if err != nil {
		return nil, err
	}
	return tx, err
}

func (cli *client) QueryTransactionsWithTransactionView(payload *model.QueryTransactionsPayload) (*resp.PaginationResponseTransactionView, error) {
	payload.StructureType = model.Native
	var resp resp.PaginationResponseTransactionView
	err := cli.c.Call(&resp, "query_transactions", payload)
	if err != nil {
		return &resp, err
	}

	return &resp, err
}

func (cli *client) QueryTransactionsWithTransactionInfo(payload *model.QueryTransactionsPayload) (*resp.PaginationResponseTransactionInfo, error) {
	payload.StructureType = model.DoubleEntry
	var resp resp.PaginationResponseTransactionInfo
	err := cli.c.Call(&resp, "query_transactions", payload)
	if err != nil {
		return &resp, err
	}

	return &resp, err
}

//func (cli *client) toTransferPayload(payload *model.SimpleTransferPayload) (*model.TransferPayload, error) {
//	fromBalances, err := cli.getBalance(payload.From, payload.AssetInfo)
//	if err != nil {
//		return nil, err
//	}
//
//	toAddresses := make([]string, len(payload.To))
//	for i, to := range payload.To {
//		toAddresses[i] = to.Address
//	}
//	toBalances, err := cli.getBalance(toAddresses, payload.AssetInfo)
//	if err != nil {
//		return nil, err
//	}
//
//	err = cli.feePay(fromBalances, toBalances, payload)
//	if err != nil {
//		return nil, err
//	}
//
//	source := cli.getSource(fromBalances, payload.To, payload.AssetInfo.AssetType)
//
//	builder := model.NewTransferBuilder()
//	if payload.AssetInfo.AssetType == common.Udt {
//		builder.AssetInfo = payload.AssetInfo.AssetInfo
//	}
//	builder.AddFrom(payload.From, source)
//	for _, to := range payload.To {
//		if payload.AssetInfo.AssetType == common.Udt {
//			number, err := cli.GetAccountNumber(to.Address)
//			if err != nil {
//				return nil, err
//			}
//			if number > 0 {
//				builder.AddTo(to.Address, mode.Pay_by_to, to.Amount)
//			} else {
//				builder.AddTo(to.Address, mode.Pay_by_from, to.Amount)
//			}
//		} else {
//			builder.AddTo(to.Address, mode.Pay_by_from, to.Amount)
//		}
//	}
//
//	return builder.Build(), nil
//}
//
//func (cli *client) getSource(fromBalances []*resp.GetBalanceResponse, to []*model.ToInfo, assetType common.AssetType) string {
//	fromBalance := cli.getBalanceByAssetTypeAndBalanceType(fromBalances, assetType, "claimable")
//	totalAmount := big.NewInt(0)
//	for _, smartTo := range to {
//		totalAmount = totalAmount.Add(totalAmount, smartTo.Amount)
//	}
//
//	if fromBalance.Cmp(totalAmount) >= 0 {
//		return source.Fleeting
//	} else {
//		return source.Unconstrained
//	}
//}
//
//func (cli *client) feePay(fromBalances, toBalances []*resp.GetBalanceResponse, payload *model.SimpleTransferPayload) error {
//	from := cli.getBalanceByAssetTypeAndBalanceType(fromBalances, common.Ckb, "free")
//	to := cli.getBalanceByAssetTypeAndBalanceType(toBalances, common.Ckb, "free")
//
//	feeThreshold := amount.CkbWithDecimalToShannon(0.0001)
//
//	if from.Cmp(feeThreshold) < 0 && to.Cmp(feeThreshold) < 0 {
//		return errors.New("CKB Insufficient balance to pay the fee")
//	}
//	if from.Cmp(feeThreshold) < 0 && to.Cmp(feeThreshold) >= 0 {
//		for _, getBalanceResponse := range toBalances {
//			for _, balance := range getBalanceResponse.Balances {
//				payload.From = append(payload.From, balance.Address)
//			}
//		}
//	}
//
//	return nil
//}
//
//func (cli *client) getBalanceByAssetTypeAndBalanceType(balances []*resp.GetBalanceResponse, assetType common.AssetType, balanceType string) *big.Int {
//	amount := big.NewInt(0)
//	for _, getBalanceResponse := range balances {
//		for _, balance := range getBalanceResponse.Balances {
//			if balance.AssetInfo.AssetType == assetType {
//				if balanceType == "free" {
//					amount = amount.Add(amount, balance.Free)
//				}
//
//				if balanceType == "claimable" {
//					amount = amount.Add(amount, balance.Claimable)
//				}
//
//				if balanceType == "freezed" {
//					amount = amount.Add(amount, balance.Freezed)
//				}
//			}
//		}
//	}
//	return amount
//}
//
//func (cli *client) getBalance(addresses []string, assetInfo *common.AssetInfo) ([]*resp.GetBalanceResponse, error) {
//	result := make([]*resp.GetBalanceResponse, len(addresses))
//	for i, address := range addresses {
//		builder := model.NewGetBalancePayloadBuilder()
//		builder.AddItem(address)
//		builder.AddAssetInfo(common.NewCkbAsset())
//		if assetInfo != nil {
//			builder.AddAssetInfo(assetInfo)
//		}
//
//		balance, err := cli.GetBalance(builder.Build())
//		if err != nil {
//			return nil, err
//		}
//
//		result[i] = balance
//	}
//
//	return result, nil
//}

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
