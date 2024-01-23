package golangsolanarpc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBlocksWithLimti(t *testing.T) {
	client := SolanaClient{
		RpcEndpoint: test_node,
	}
	start := uint64(262972100)
	limit := uint64(5)
	response, err := client.GetBlocksWithLimit(start, limit)
	if err != nil {
		t.Fatal("Error getting blocks", err)
	}
	assert.NotNil(t, response)
	assert.Equal(t, 5, len(response.Blocks))
	assert.Equal(t, uint64(262972101), response.Blocks[1])
}
