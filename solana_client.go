package golangsolanarpc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/google/uuid"
)

const (
	test_node = "https://api.devnet.solana.com"
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
	GetBlocks(startSlot uint64, endSlot *uint64) (*BlocksResponse, error)
	GetBlocksWithLimit(startSlot uint64, limit uint64) (*BlocksResponse, error)
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

func transformRPCResponse[T any](response *RPCResponse) (*T, error) {
	bytes, err := json.Marshal(response.Result)
	if err != nil {
		fmt.Println("error marshalling response")
		return nil, err
	}
	var m T
	dec := json.NewDecoder(strings.NewReader(string(bytes)))
	dec.UseNumber()
	err = dec.Decode(&m)
	if err != nil {
		fmt.Println("error decoding response")
		return nil, err
	}
	return &m, nil
}

func getRPCRequest(method string, params []interface{}) RPCRequest {
	return RPCRequest{
		Method:         method,
		Params:         params,
		Id:             uuid.NewString(),
		JsonRpcVersion: jsonRpcVersion,
	}
}
