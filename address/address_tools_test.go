package address

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateShortAddress(t *testing.T) {
	shortAddress, _ := GenerateShortAddress(Mainnet)
	fmt.Println(shortAddress)
}

func TestGenerateAcpAddress(t *testing.T) {
	address := "ckt1qyqqtg06h75ymw098r3w0l3u4xklsj04tnsqctqrmc"
	acpAddress, err := GenerateAcpAddress(address)
	assert.Nil(t, err)
	assert.Equal(t, "ckt1qypqtg06h75ymw098r3w0l3u4xklsj04tnsqkm65q6", acpAddress)
}

func TestGenerateChequeAddress(t *testing.T) {
	senderAddress := "ckt1qyq27z6pccncqlaamnh8ttapwn260egnt67ss2cwvz"
	receiverAddress := "ckt1qyqqtg06h75ymw098r3w0l3u4xklsj04tnsqctqrmc"
	acpAddress, err := GenerateChequeAddress(senderAddress, receiverAddress)
	assert.Nil(t, err)
	assert.Equal(t, "ckt1q3sdtuu7lnjqn3v8ew02xkwwlh4dv5x2z28shkwt8p2nfruccux4k5kw5xmckqjq7gwpe990sn88xssv96try4l46hu6nnudr2huau238a4prwus9pqts3uptms", acpAddress)
}

func TestParseAddress(t *testing.T) {

}
