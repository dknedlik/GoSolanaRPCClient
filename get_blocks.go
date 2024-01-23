package golangsolanarpc

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

func (c SolanaClient) GetBlocks(startSlot uint64, endSlot *uint64) (*BlocksResponse, error) {
	var params []interface{}
	params = append(params, startSlot)
	if endSlot != nil {
		params = append(params, *endSlot)
	}
	request := getRPCRequest("getBlocks", params[:])
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
		fmt.Println("Could not read uint64 array from results")
		return nil, errors.New("could not read uint64 array from results")
	}
	blocks := []uint64{}
	for _, v := range result {
		number, err := strconv.ParseUint(string(v.(json.Number).String()), 10, 64)
		if err != nil {
			fmt.Println("invalid format received for block id")
			return nil, err
		}
		blocks = append(blocks, number)
	}
	return &BlocksResponse{Blocks: blocks}, nil
}
