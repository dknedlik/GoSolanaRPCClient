package golangsolanarpc

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

func (c SolanaClient) GetBlockHeight() (uint64, error) {
	request := getRPCRequest("getBlockHeight", nil)
	response, err := c.sendRequest(request)
	if err != nil {
		return 0, err
	}
	height, ok := response.Result.(json.Number)
	if !ok {
		fmt.Println("Unable to convert response to uint64")
		return 0, errors.New("unable to convert response to uint64")
	}
	number, err := strconv.ParseUint(string(height.String()), 10, 64)
	if err != nil {
		fmt.Println("Unable to parse block height to uint64", err)
		return 0, err
	}

	return number, nil
}
