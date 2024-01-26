package golangsolanarpc

import "fmt"

func (c SolanaClient) GetEpcohSchedule() (*EpochSchedule, error) {
	request := getRPCRequest("getEpochSchedule", nil)
	response, err := c.sendRequest(request)
	if err != nil {
		return nil, err
	}
	result, err := transformRPCResponse[EpochSchedule](response)
	if err != nil {
		fmt.Println("error transforming get epoch schedule response")
		return nil, err
	}
	return result, nil
}
