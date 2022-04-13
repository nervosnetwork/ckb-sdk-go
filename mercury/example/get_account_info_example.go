package test

import (
	"testing"

	"github.com/nervosnetwork/ckb-sdk-go/mercury/example/constant"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/common"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/model/req"
)

func TestGetAccountInfoByAddress1(t *testing.T) {
	builder := model.NewGetAccountInfoPayloadBuilder()
	item, _ := req.NewAddressItem("ckt1qzda0cr08m85hc8jlnfp3zer7xulejywt49kt2rr0vthywaa50xwsq06y24q4tc4tfkgze35cc23yprtpzfrzygljdjh9")
	payload := builder.SetItem(item).
		AddAssetInfo(common.NewUdtAsset("0xf21e7350fa9518ed3cbb008e0e8c941d7e01a12181931d5608aa366ee22228bd")).
		Build()
	printJson(payload)
	account, err := constant.GetMercuryApiInstance().GetAccountInfo(payload)
	if err != nil {
		t.Error(err)
	}
	printJson(account)
}

func TestGetAccountInfoByAddress2(t *testing.T) {
	builder := model.NewGetAccountInfoPayloadBuilder()
	item, _ := req.NewAddressItem("ckt1qq6pngwqn6e9vlm92th84rk0l4jp2h8lurchjmnwv8kq3rt5psf4vq06y24q4tc4tfkgze35cc23yprtpzfrzygsptkzn")
	payload := builder.SetItem(item).
		AddAssetInfo(common.NewUdtAsset("0xf21e7350fa9518ed3cbb008e0e8c941d7e01a12181931d5608aa366ee22228bd")).
		Build()
	printJson(payload)
	account, err := constant.GetMercuryApiInstance().GetAccountInfo(payload)
	if err != nil {
		t.Error(err)
	}
	printJson(account)
}
