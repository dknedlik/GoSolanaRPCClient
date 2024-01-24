package golangsolanarpc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBlockCommitement(t *testing.T) {
	client := SolanaClient{
		RpcEndpoint: test_node,
	}
	response, err := client.GetBlockCommitment(500)
	if err != nil {
		t.Fatal("Error getting block commitment")
	}
	//the actual values are not stable enough so just check that it is not nil
	//Have verified the payload previously.
	assert.NotNil(t, response)
	// assert.Equal(t, uint64(158079506343876433), *response.TotalStake)
}
