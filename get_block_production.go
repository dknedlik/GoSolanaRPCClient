package golangsolanarpc

import (
	"fmt"
)

func (c SolanaClient) GetBlockProduction() (*BlockProductionResult, error) {
	request := getRPCRequest("getBlockProduction", nil)
	response, err := c.sendRequest(request)
	if err != nil {
		return nil, err
	}
	result, err := transformRPCResponse[solBlockProductionResponse](response)
	if err != nil {
		fmt.Println("Error decoding block production result", err)
		return nil, err
	}
	return &result.Value, nil
}
