package golangsolanarpc

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/google/uuid"
)

func (c SolanaClient) GetAccountInfo(id string) (*AccountInfo, error) {
	var params [2]interface{}
	params[0] = id
	params[1] = RpcEncoding{Encoding: "base64"}
	request := RPCRequest{
		Method:         "getAccountInfo",
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
		fmt.Println("error marshalling get account info response")
		return nil, err
	}
	var m solAccountInfoResponse
	dec := json.NewDecoder(strings.NewReader(string(bytes)))
	dec.UseNumber()
	err = dec.Decode(&m)
	if err != nil {
		fmt.Println("error decoding get account info response")
		return nil, err
	}

	return &m.Value, nil
}
