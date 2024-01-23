package golangsolanarpc

import (
	"fmt"
)

func (c SolanaClient) GetBalance(id string) (uint64, error) {
	var params [2]interface{}
	params[0] = id
	params[1] = RpcEncoding{Encoding: "base64"}
	request := getRPCRequest("getBalance", params[:])
	response, err := c.sendRequest(request)
	if err != nil {
		return 0, err
	}
	result, err := transformRPCResponse[BalanceRPCResult](response)
	if err != nil {
		fmt.Println("error transforming get balance response")
		return 0, err
	}
	return result.Value, nil
}
