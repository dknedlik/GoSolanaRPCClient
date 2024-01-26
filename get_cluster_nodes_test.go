package golangsolanarpc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetClusterNodes(t *testing.T) {
	client := SolanaClient{
		RpcEndpoint: test_node,
	}
	response, err := client.GetClusterNodes()
	if err != nil {
		t.Fatal("Error getting cluster nodes", err)
	}
	assert.NotNil(t, response)
	assert.Equal(t, 94, len(response.Nodes))
	n := response.Nodes[0]
	assert.Equal(t, "39cvwUEpgka9bU7Sn4my82VViMDWaCxi4YoPevfZxLf3", n.Pubkey)
	assert.Equal(t, uint32(1337574167), *n.FeatureSet)
}
