package mercury

import (
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/action"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/resp"
	"github.com/pkg/errors"
)

type MercuryApi interface {
	GetBalance(payload *model.GetBalancePayload) (*resp.GetBalanceResponse, error)
	BuildTransferTransaction(payload *model.TransferPayload) (*resp.TransferCompletionResponse, error)
	BuildWalletCreationTransaction(payload *model.CreateWalletPayload) (*resp.TransferCompletionResponse, error)
	RegisterAddresses(normalAddresses []string) ([]string, error)
}

type DefaultMercuryApi struct {
	c *rpc.Client
}

func (cli *DefaultMercuryApi) RegisterAddresses(normalAddresses []string) ([]string, error) {
	var scriptHash []string
	err := cli.c.Call(&scriptHash, "register_addresses", normalAddresses)
	if err != nil {
		return scriptHash, err
	}

	return scriptHash, err
}

func (cli *DefaultMercuryApi) GetBalance(payload *model.GetBalancePayload) (*resp.GetBalanceResponse, error) {
	var balance resp.GetBalanceResponse
	err := cli.c.Call(&balance, "get_balance", payload)
	if err != nil {
		return &balance, err
	}

	return &balance, err
}

func (cli *DefaultMercuryApi) BuildTransferTransaction(payload *model.TransferPayload) (*resp.TransferCompletionResponse, error) {
	var resp resp.TransferCompletionResponse
	if payload.UdtHash == "" {
		for _, item := range payload.Items {
			if item.To.Action == action.Lend_by_from || item.To.Action == action.Pay_by_to {
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

func (cli *DefaultMercuryApi) BuildWalletCreationTransaction(payload *model.CreateWalletPayload) (*resp.TransferCompletionResponse, error) {
	var resp resp.TransferCompletionResponse
	err := cli.c.Call(&resp, "build_wallet_creation_transaction", payload)
	if err != nil {
		return &resp, err
	}

	return &resp, err
}

func NewMercuryApi(address string) (MercuryApi, error) {
	dial, err := rpc.Dial(address)
	if err != nil {
		return &DefaultMercuryApi{}, err
	}

	return &DefaultMercuryApi{dial}, err
}
