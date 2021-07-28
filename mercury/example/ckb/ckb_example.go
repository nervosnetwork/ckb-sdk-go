package ckb

import (
	"context"
	"fmt"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/example/constant"
	"testing"
)

func TestCkb(t *testing.T) {
	number, _ := constant.GetMercuryApiInstance().GetTipBlockNumber(context.Background())
	fmt.Println(number)
}
