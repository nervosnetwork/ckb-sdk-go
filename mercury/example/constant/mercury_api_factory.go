package constant

import (
	"github.com/nervosnetwork/ckb-sdk-go/api"
)

const MERCURY_URL = "https://mercury-testnet.ckbapp.dev/"

const INDEXER_URL = "https://testnet.ckb.dev/indexer"

const CKB_URL = "https://testnet.ckb.dev/"

type MercuryApiFactory struct {
	clent api.CkbApi
}

func GetMercuryApiInstance() api.CkbApi {
	api, err := api.NewCkbApi(CKB_URL, MERCURY_URL, INDEXER_URL)
	if err != nil {
		panic(err)
	}

	return api
}
