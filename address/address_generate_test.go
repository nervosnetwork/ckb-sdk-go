package address

import (
	"fmt"
	"testing"
)

func TestGenerateShortAddress(t *testing.T) {
	shortAddress, _ := GenerateShortAddress(Mainnet)

	fmt.Println(shortAddress)
}
