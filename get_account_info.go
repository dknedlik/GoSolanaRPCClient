package golangsolanarpc

import (
	"fmt"
)

func (c SolanaClient) GetAccountInfo(id string) (*AccountInfo, error) {
	var params [2]interface{}
	params[0] = id
	params[1] = RpcEncoding{Encoding: "base64"}
	request := getRPCRequest("getAccountInfo", params[:])
	response, err := c.sendRequest(request)
	if err != nil {
		return nil, err
	}
	result, err := transformRPCResponse[solAccountInfoResponse](response)
	if err != nil {
		fmt.Println("error transforming get account info response")
		return nil, err
	}

	return &result.Value, nil
}
