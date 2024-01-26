package golangsolanarpc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEpochInfo(t *testing.T) {
	client := SolanaClient{
		RpcEndpoint: test_node,
	}
	response, err := client.GetEpcohInfo()
	if err != nil {
		t.Fatal("Error getting epoch info", err)
	}
	assert.NotNil(t, response)
}
