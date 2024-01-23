package golangsolanarpc

import (
	"fmt"
)

func (c SolanaClient) GetBlock(slot uint64) (*BlockResponse, error) {
	var params [1]interface{}
	params[0] = slot
	request := getRPCRequest("getBlock", params[:])
	response, err := c.sendRequest(request)
	if err != nil {
		return nil, err
	}
	if response.Result == nil {
		fmt.Println("No block found")
		return nil, nil
	}
	result, err := transformRPCResponse[BlockResponse](response)
	if err != nil {
		fmt.Println("error transforming get block response")
		return nil, err
	}

	return result, nil
}
