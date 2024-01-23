package golangsolanarpc

import (
	"testing"

	"github.com/zeebo/assert"
)

func TestGetAccountInfo(t *testing.T) {
	client := SolanaClient{
		RpcEndpoint: test_node,
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
