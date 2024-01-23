package golangsolanarpc

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

func (c SolanaClient) GetBlockTime(slot uint64) (uint64, error) {
	var params [1]interface{}
	params[0] = slot
	request := getRPCRequest("getBlockTime", params[:])
	response, err := c.sendRequest(request)
	if err != nil {
		return 0, err
	}
	if response.Error != nil {
		fmt.Println("Get block time failed with error: ", response.Error.Message)
		return 0, errors.New(response.Error.Message)
	}
	v, ok := response.Result.(json.Number)
	if !ok {
		fmt.Println("Get block time response was not valid number")
		return 0, errors.New("get block time response was not valid number")
	}
	number, err := strconv.ParseUint(string(v.String()), 10, 64)
	if err != nil {
		fmt.Println("get block time invalid format received for block time")
		return 0, err
	}

	return number, nil
}
