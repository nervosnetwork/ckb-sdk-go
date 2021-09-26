package constant

import (
	"github.com/nervosnetwork/ckb-sdk-go/api"
)

const MERCURY_URL = "http://127.0.0.1:8116"

const INDEXER_URL = "https://mercury-testnet.ckbapp.dev"

const CKB_URL = "https://mercury-testnet.ckbapp.dev"

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
