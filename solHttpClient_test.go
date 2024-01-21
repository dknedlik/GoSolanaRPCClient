package golangsolanarpc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAccountInfo(t *testing.T) {
	client := SolanaClient{
		RpcEndpoint: "https://api.devnet.solana.com",
	}
	resp, err := client.GetAccountInfo("vines1vzrYbzLMRdu58ou5XTby4qAqVRLmqo36NKPTg")
	if err != nil {
		t.Fatal("Error getting account information")
	}
	assert.Equal(t, resp.RentEpoch, uint64(18446744073709551615))
	assert.Equal(t, resp.Owner, "11111111111111111111111111111111")
	assert.Equal(t, resp.Space, uint64(0))
	assert.Equal(t, resp.Data[0], "")
	assert.Equal(t, resp.Data[1], "base64")
	assert.Equal(t, resp.Executable, false)
	assert.Equal(t, resp.Lamports, uint64(88820743038050))
}

func TestGetBalance(t *testing.T) {
	client := SolanaClient{
		RpcEndpoint: "https://api.devnet.solana.com",
	}
	resp, err := client.GetBalance("vines1vzrYbzLMRdu58ou5XTby4qAqVRLmqo36NKPTg")
	if err != nil {
		t.Fatal("Error getting balance")
	}

	assert.Equal(t, uint64(88820743038050), resp)
}

func TestGetBlockHeight(t *testing.T) {
	client := SolanaClient{
		RpcEndpoint: "https://api.devnet.solana.com",
	}
	height, err := client.GetBlockHeight()
	if err != nil {
		t.Fatal("Error getting block height")
	}
	assert.NotEqual(t, uint64(0), height)
}
