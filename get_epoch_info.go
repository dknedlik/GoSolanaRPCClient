package golangsolanarpc

import "fmt"

func (c SolanaClient) GetEpcohInfo() (*EpochInfo, error) {
	request := getRPCRequest("getEpochInfo", nil)
	response, err := c.sendRequest(request)
	if err != nil {
		return nil, err
	}
	result, err := transformRPCResponse[EpochInfo](response)
	if err != nil {
		fmt.Println("error transforming get epoch info response")
		return nil, err
	}
	return result, nil
}
