package constant

import "github.com/nervosnetwork/ckb-sdk-go/rpc"

const NODE_URL = "http://127.0.0.1:8114"

type CkbNodeFactory struct {
	clent rpc.Client
}

func GetCkbNodeInstance() rpc.Client {
	client, err := rpc.Dial(NODE_URL)
	if err != nil {
		panic(err)
	}

	return client
}
