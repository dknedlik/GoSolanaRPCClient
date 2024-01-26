package golangsolanarpc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEpochSchedule(t *testing.T) {
	client := SolanaClient{
		RpcEndpoint: test_node,
	}
	response, err := client.GetEpcohInfo()
	if err != nil {
		t.Fatal("Error getting epoch schedule", err)
	}
	assert.NotNil(t, response)
}
