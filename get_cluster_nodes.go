package golangsolanarpc

import "fmt"

func (c SolanaClient) GetClusterNodes() (*ClusterNodesResponse, error) {
	request := getRPCRequest("getClusterNodes", nil)
	response, err := c.sendRequest(request)
	if err != nil {
		return nil, err
	}
	result, err := transformRPCResponse[[]NodeInfo](response)
	if err != nil {
		fmt.Println("error transforming get cluster nodes response")
		return nil, err
	}
	return &ClusterNodesResponse{Nodes: *result}, nil
}
