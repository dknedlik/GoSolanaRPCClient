package golangsolanarpc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type SolanaClient struct {
	RpcEndpoint string
}

type SolanaRPCClient interface {
	GetAccountInfo(id string) (*AccountInfo, error)
	GetBalance(id string) (uint64, error)
	GetBlockHeight() (uint64, error)
	GetBlock(slot uint64) (*BlockResponse, error)
	GetBlockProduction() (*BlockProductionResult, error)
	GetBlockCommitment(slot uint64) (*BlockCommitmentResponse, error)
}

const (
	Finalized Commitment = "finalized"
	Confirmed Commitment = "confirmed"
	Processed Commitment = "processed"
)

const (
	jsonRpcVersion = "2.0"
)

func (c SolanaClient) sendRequest(req RPCRequest) (*RPCResponse, error) {
	reqJSON, error := json.Marshal(req)
	if error != nil {
		fmt.Printf("Error marshalling request data %s/n", error)
		return nil, error
	}
	request, error := http.NewRequest("POST", c.RpcEndpoint, bytes.NewBuffer(reqJSON))
	if error != nil {
		fmt.Printf("Error creating request %s/n", error)
		return nil, error
	}
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}
	response, error := client.Do(request)
	if error != nil {
		fmt.Printf("Error sending request %s/n", error)
		return nil, error
	}
	defer response.Body.Close()

	//fmt.Println("response Status:", response.Status)
	//fmt.Println("response Headers:", response.Header)
	body, _ := io.ReadAll(response.Body)
	var m RPCResponse
	dec := json.NewDecoder(strings.NewReader(string(body)))
	dec.UseNumber()

	error = dec.Decode(&m)
	if error != nil {
		fmt.Println("Error decoding response", error)
		return nil, error
	}
	//fmt.Println("response Body:", string(body))
	return &m, nil
}
