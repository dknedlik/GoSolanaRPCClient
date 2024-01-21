package golangsolanarpc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/google/uuid"
)

const (
	Finalized Commitment = "finalized"
	Confirmed Commitment = "confirmed"
	Processed Commitment = "processed"
)

const (
	jsonRpcVersion = "2.0"
)

func (c SolanaClient) GetAccountInfo(id string) (*AccountInfo, error) {
	var params [2]interface{}
	params[0] = id
	params[1] = RpcEncoding{Encoding: "base64"}
	request := RPCRequest{
		Method:         "getAccountInfo",
		Params:         params,
		Id:             uuid.NewString(),
		JsonRpcVersion: jsonRpcVersion,
	}
	response, err := c.sendRequest(request)
	if err != nil {
		return nil, err
	}
	var m AcctInfoRPCResponse
	dec := json.NewDecoder(strings.NewReader(string(*response)))
	dec.UseNumber()

	err = dec.Decode(&m)
	if err != nil {
		fmt.Println("error decoding json", err)
		return nil, err
	}
	return &m.Result.Value, nil
}

func (c SolanaClient) sendRequest(req RPCRequest) (*[]byte, error) {
	reqJSON, error := json.Marshal(req)
	if error != nil {
		fmt.Printf("Error marshalling request data %s/n", error)
		return nil, error
	}
	request, error := http.NewRequest("POST", c.RpcEndpoint, bytes.NewBuffer(reqJSON))
	if error != nil {
		fmt.Printf("Error creating request %s/n", error)
		return nil, error
	}
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}
	response, error := client.Do(request)
	if error != nil {
		fmt.Printf("Error sending request %s/n", error)
		return nil, error
	}
	defer response.Body.Close()

	//fmt.Println("response Status:", response.Status)
	//fmt.Println("response Headers:", response.Header)
	body, _ := io.ReadAll(response.Body)
	fmt.Println("response Body:", string(body))
	return &body, nil
}
