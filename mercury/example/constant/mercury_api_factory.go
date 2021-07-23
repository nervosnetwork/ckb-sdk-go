package constant

import "github.com/nervosnetwork/ckb-sdk-go/mercury"

const MERCURY_URL = "http://8.210.169.63:8116"

type MercuryApiFactory struct {
	clent mercury.MercuryApi
}

func GetMercuryApiInstance() mercury.MercuryApi {
	api, err := mercury.NewMercuryApi(MERCURY_URL)
	if err != nil {
		panic(err)
	}

	return api
}
