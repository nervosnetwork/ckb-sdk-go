package dao

import (
	"fmt"
	"github.com/nervosnetwork/ckb-sdk-go/mercury/example/constant"
	"github.com/nervosnetwork/ckb-sdk-go/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetDaoDepositCellInfo(t *testing.T) {
	client := constant.GetMercuryApiInstance()
	daoHelper := DaoHelper{Client: client}

	outpoint := types.OutPoint{
		TxHash: types.HexToHash("0x41bbccdf7015ea8458d7ef3499dc80cb2d3dc10cf48eb2b7f8f74468b24027fc"),
		Index:  0,
	}
	withdrawBlockHash := types.HexToHash("0xbaef9b22ee3d04d8fc3ad8c04f8403ad3b3b39c5ace51406c5305920976105f7")

	daoCellInfo, err := daoHelper.GetDaoDepositCellInfo(&outpoint, &withdrawBlockHash)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, outpoint, daoCellInfo.Outpoint)
	assert.Equal(t, withdrawBlockHash, daoCellInfo.withdrawBlockHash)
	assert.Equal(t, uint64(2383851), daoCellInfo.Compensation)
	assert.Equal(t, uint64(11055500000), daoCellInfo.DepositCapacity)
	assert.Equal(t, uint32(247), daoCellInfo.NextClaimableEpochNumber)
	assert.Equal(t, uint64(171182), daoCellInfo.NextClaimableBlock)
}

func TestGetDaoDepositCellInfoWithWithdrawOutpoint(t *testing.T) {
	client := constant.GetMercuryApiInstance()
	daoHelper := DaoHelper{Client: client}

	outpoint := types.OutPoint{
		TxHash: types.HexToHash("0x41bbccdf7015ea8458d7ef3499dc80cb2d3dc10cf48eb2b7f8f74468b24027fc"),
		Index:  0,
	}

	withdrawOutpoint := types.OutPoint{
		TxHash: types.HexToHash("0x88b071409d8bda119c1b4d613ccd78abbc01442566defcb7f745f6084c81adcb"),
		Index:  0,
	}

	daoCellInfo, err := daoHelper.GetDaoDepositCellInfoWithWithdrawOutpoint(&outpoint, &withdrawOutpoint)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, outpoint, daoCellInfo.Outpoint)
	assert.Equal(t, uint64(2383851), daoCellInfo.Compensation)
	assert.Equal(t, uint64(11055500000), daoCellInfo.DepositCapacity)
	assert.Equal(t, uint32(247), daoCellInfo.NextClaimableEpochNumber)
	assert.Equal(t, uint64(171182), daoCellInfo.NextClaimableBlock)
}

func TestResolveEpoch(t *testing.T) {
	epochLength, blockIndexInEpoch, epochNumber := ResolveEpoch(1979138915175034)

	assert.Equal(t, uint16(1800), epochLength)
	assert.Equal(t, uint16(1072), blockIndexInEpoch)
	assert.Equal(t, uint32(2682), epochNumber)
}

func TestExtractArFromDaoData(t *testing.T) {
	daoData := types.HexToHash("8268d571c743a32ee1e547ea57872300989ceafa3e710000005d6a650b53ff06")
	x := extractArFromDaoData(&daoData)
	fmt.Println(x)
}
