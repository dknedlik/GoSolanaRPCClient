package golangsolanarpc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBlockProduction(t *testing.T) {
	client := SolanaClient{
		RpcEndpoint: test_node,
	}
	production, err := client.GetBlockProduction()
	if err != nil {
		t.Fatal("Error getting block production")
	}
	assert.NotNil(t, production)
	assert.NotNil(t, production.ByIdentity)
	assert.NotNil(t, production.Range)
	assert.NotNil(t, production.Range.FirstSlot)
	assert.NotNil(t, production.Range.LastSlot)
}
