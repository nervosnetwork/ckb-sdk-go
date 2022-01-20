package indexer

import (
	"context"
	"fmt"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/example/constant"
	"testing"
)

func TestGetTip(t *testing.T) {
	mercuryApi := constant.GetMercuryApiInstance()
	tip, _ := mercuryApi.GetTip(context.Background())
	fmt.Println(tip)
}
