package golangsolanarpc

import (
	"testing"

	"github.com/zeebo/assert"
)

func TestGetBalance(t *testing.T) {
	client := SolanaClient{
		RpcEndpoint: test_node,
	}
	resp, err := client.GetBalance("vines1vzrYbzLMRdu58ou5XTby4qAqVRLmqo36NKPTg")
	if err != nil {
		t.Fatal("Error getting balance")
	}

	assert.Equal(t, uint64(88820743038050), resp)
}
