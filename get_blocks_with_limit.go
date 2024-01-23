package golangsolanarpc

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

func (c SolanaClient) GetBlocksWithLimit(startSlot uint64, limit uint64) (*BlocksResponse, error) {
	var params [2]interface{}
	params[0] = startSlot
	params[1] = limit
	request := getRPCRequest("getBlocksWithLimit", params[:])
	response, err := c.sendRequest(request)
	if err != nil {
		return nil, err
	}
	if response.Result == nil {
		fmt.Println("No blocks found")
		return &BlocksResponse{}, nil
	}
	result, ok := response.Result.([]interface{})
	if !ok {
		fmt.Println("GetBlocksWithLimit-> Could not read uint64 array from results")
		return nil, errors.New("getBlocksWithLimit -> could not read array from results")
	}
	blocks := []uint64{}
	for _, v := range result {
		number, err := strconv.ParseUint(string(v.(json.Number).String()), 10, 64)
		if err != nil {
			fmt.Println("GetBlocksWithLimit-> invalid format received for block id")
			return nil, err
		}
		blocks = append(blocks, number)
	}
	return &BlocksResponse{Blocks: blocks}, nil
}
