package golangsolanarpc

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/google/uuid"
)

func (c SolanaClient) GetBlock(slot uint64) (*BlockResponse, error) {
	var params [1]interface{}
	params[0] = slot
	request := RPCRequest{
		Method:         "getBlock",
		Params:         params,
		Id:             uuid.NewString(),
		JsonRpcVersion: jsonRpcVersion,
	}
	response, err := c.sendRequest(request)
	if err != nil {
		return nil, err
	}
	if response.Result == nil {
		fmt.Println("No block found")
		return nil, nil
	}
	bytes, err := json.Marshal(response.Result)
	if err != nil {
		fmt.Println("Error marshalling get block response")
		return nil, err
	}
	var m BlockResponse
	dec := json.NewDecoder(strings.NewReader(string(bytes)))
	dec.UseNumber()
	err = dec.Decode(&m)
	if err != nil {
		fmt.Println("error decoding get block response")
		return nil, err
	}

	fmt.Println("Response bytes = ", m.BlockHeight)
	return &m, nil
}
