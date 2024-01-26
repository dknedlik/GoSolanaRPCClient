package golangsolanarpc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBlockTome(t *testing.T) {
	client := SolanaClient{
		RpcEndpoint: test_node,
	}
	slot := uint64(262972100)
	response, err := client.GetBlockTime(slot)
	if err != nil {
		t.Fatal("Error getting block time", err)
	}
	assert.NotNil(t, response)
	assert.Equal(t, uint64(1701703179), response)
	slot = uint64(999999999)
	response, err = client.GetBlockTime(slot)
	assert.Equal(t, uint64(0), response)
	assert.Equal(t, "Block not available for slot 999999999", err.Error())
}
