package golangsolanarpc

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/google/uuid"
)

func (c SolanaClient) GetBlockProduction() (*BlockProductionResult, error) {
	request := RPCRequest{
		Method:         "getBlockProduction",
		Params:         nil,
		Id:             uuid.NewString(),
		JsonRpcVersion: jsonRpcVersion,
	}
	response, err := c.sendRequest(request)
	if err != nil {
		return nil, err
	}
	bytes, err := json.Marshal(response.Result)
	if err != nil {
		fmt.Println("Error marshalling get block production response")
		return nil, err
	}
	var m solBlockProductionResponse
	dec := json.NewDecoder(strings.NewReader(string(bytes)))
	dec.UseNumber()
	err = dec.Decode(&m)
	if err != nil {
		fmt.Println("Could not decode block production result", err)
		return nil, err
	}
	return &m.Value, nil
}
