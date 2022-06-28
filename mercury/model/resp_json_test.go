package model

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJsonDaoState(t *testing.T) {
	jsonText := []byte(`
{
	"type": "Deposit",
	"value": "0x100"
}`)
	var v DaoState
	json.Unmarshal(jsonText, &v)
	assert.Equal(t, DaoStateTypeDeposit, v.Type)
	assert.Equal(t, 1, len(v.Value))
	assert.Equal(t, uint64(0x100), v.Value[0])

	jsonText = []byte(`
{
	"type": "Deposit",
	"value": ["0x100", "0x400"]
}`)
	json.Unmarshal(jsonText, &v)
	assert.Equal(t, DaoStateTypeDeposit, v.Type)
	assert.Equal(t, 2, len(v.Value))
	assert.Equal(t, uint64(0x100), v.Value[0])
	assert.Equal(t, uint64(0x400), v.Value[1])
}
