package golangsolanarpc

import (
	"fmt"
)

func (c SolanaClient) GetBlockCommitment(slot uint64) (*BlockCommitmentResponse, error) {
	var params [1]interface{}
	params[0] = slot
	request := getRPCRequest("getBlockCommitment", params[:])

	response, err := c.sendRequest(request)
	if err != nil {
		return nil, err
	}
	result, err := transformRPCResponse[BlockCommitmentResponse](response)
	if err != nil {
		fmt.Println("error transforming get block commitemtn response")
		return nil, err
	}

	return result, nil
}
