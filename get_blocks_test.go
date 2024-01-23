package golangsolanarpc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBlocks(t *testing.T) {
	client := SolanaClient{
		RpcEndpoint: test_node,
	}
	start := uint64(262972100)
	end := uint64(262972180)
	response, err := client.GetBlocks(start, &end)
	if err != nil {
		t.Fatal("Error getting blocks", err)
	}
	assert.NotNil(t, response)
	assert.Equal(t, response.Blocks[1], uint64(262972101))
}
