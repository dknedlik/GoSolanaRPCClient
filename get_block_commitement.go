package golangsolanarpc

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/google/uuid"
)

func (c SolanaClient) GetBlockCommitment(slot uint64) (*BlockCommitmentResponse, error) {
	var params [1]interface{}
	params[0] = slot
	request := RPCRequest{
		Method:         "getBlockCommitment",
		Params:         params,
		Id:             uuid.NewString(),
		JsonRpcVersion: jsonRpcVersion,
	}
	response, err := c.sendRequest(request)
	if err != nil {
		return nil, err
	}
	bytes, err := json.Marshal(response.Result)
	if err != nil {
		fmt.Println("Error marshalling get block response")
		return nil, err
	}
	var m BlockCommitmentResponse
	dec := json.NewDecoder(strings.NewReader(string(bytes)))
	dec.UseNumber()
	err = dec.Decode(&m)
	if err != nil {
		fmt.Println("error decoding get block response")
		return nil, err
	}

	return &m, nil

}
