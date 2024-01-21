package golangsolanarpc

type Commitment string
type SolanaClient struct {
	RpcEndpoint string
}

type AccountInfo struct {
	Lamports   uint64   `json:"lamports"`
	Owner      string   `json:"owner"`
	Data       []string `json:"data"`
	Executable bool     `json:"executable"`
	RentEpoch  uint64   `json:"rentEpoch"`
	Space      uint64   `json:"space"`
}

type RPCError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type RPCRequest struct {
	Method         string      `json:"method"`
	Params         interface{} `json:"params,omitempty"`
	Id             string      `json:"id"`
	JsonRpcVersion string      `json:"jsonrpc"`
}

type AcctInfoRPCResponse struct {
	JsonRpcVersion string                 `json:"jsonrpc"`
	Result         SolAccountInfoResponse `json:"result,omitempty"`
	Error          *RPCError              `json:"error,omitempty"`
	Id             string                 `json:"id"`
}

type AccountInfoRPCResult struct {
	Context interface{} `json:"context"`
	Value   AccountInfo `json:"value"`
}

type RpcEncoding struct {
	Encoding string `json:"encoding"`
}

type SolAccountInfoResponse struct {
	Context interface{} `json:"context"`
	Value   AccountInfo `json:"value"`
}
