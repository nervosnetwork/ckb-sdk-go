package test

import (
	"fmt"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/test/constant"
	"testing"
)

func TestGetBalance(t *testing.T) {
	mercuryApi := constant.GetMercuryApiInstance()
	balance, err := mercuryApi.GetBalance(nil, constant.TEST_ADDRESS0)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(balance.Unconstrained)
}

func TestGetSudtBalance(t *testing.T) {
	mercuryApi := constant.GetMercuryApiInstance()
	balance, err := mercuryApi.GetBalance("0xf21e7350fa9518ed3cbb008e0e8c941d7e01a12181931d5608aa366ee22228bd", constant.TEST_ADDRESS0)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(balance.Unconstrained)
}
