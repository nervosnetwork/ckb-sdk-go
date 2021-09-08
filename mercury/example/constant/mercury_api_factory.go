package constant

import (
	"github.com/nervosnetwork/ckb-sdk-go/api"
)

const MERCURY_URL = "http://127.0.0.1:8116"

type MercuryApiFactory struct {
	clent api.CkbApi
}

func GetMercuryApiInstance() api.CkbApi {
	api, err := api.NewCkbApi(MERCURY_URL)
	if err != nil {
		panic(err)
	}

	return api
}
