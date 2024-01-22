package golangsolanarpc

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/google/uuid"
)

func (c SolanaClient) GetBalance(id string) (uint64, error) {
	var params [2]interface{}
	params[0] = id
	params[1] = RpcEncoding{Encoding: "base64"}
	request := RPCRequest{
		Method:         "getBalance",
		Params:         params,
		Id:             uuid.NewString(),
		JsonRpcVersion: jsonRpcVersion,
	}
	response, err := c.sendRequest(request)
	if err != nil {
		return 0, err
	}
	bytes, err := json.Marshal(response.Result)
	if err != nil {
		fmt.Println("error marshalling get balance response")
		return 0, err
	}
	var m BalanceRPCResult
	dec := json.NewDecoder(strings.NewReader(string(bytes)))
	dec.UseNumber()
	err = dec.Decode(&m)
	if err != nil {
		fmt.Println("error decoding get balance response")
		return 0, err
	}
	return m.Value, nil
}
