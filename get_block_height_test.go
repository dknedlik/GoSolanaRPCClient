package golangsolanarpc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBlockHeight(t *testing.T) {
	client := SolanaClient{
		RpcEndpoint: test_node,
	}
	height, err := client.GetBlockHeight()
	if err != nil {
		t.Fatal("Error getting block height")
	}
	assert.NotEqual(t, uint64(0), height)
}
