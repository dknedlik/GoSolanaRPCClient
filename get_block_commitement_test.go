package golangsolanarpc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBlockCommitement(t *testing.T) {
	client := SolanaClient{
		RpcEndpoint: "https://api.devnet.solana.com",
	}
	response, err := client.GetBlockCommitment(500)
	if err != nil {
		t.Fatal("Error getting block")
	}
	//I haven't found a commitemnt that has a non null commitment field
	assert.Equal(t, uint64(158079506343876433), *response.TotalStake)
}
