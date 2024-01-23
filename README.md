This is a work in progress. Anything checked in is working and has unit tests attached. But not all endpoints are implemented yet.

Basic functionality is simply create a client with the URL of the node you want to query, and use the provided funtions.

```
	client := SolanaClient{
		RpcEndpoint: "https://api.devnet.solana.com",
	}
	height, err := client.GetBlockHeight()
	if err != nil {
		t.Fatal("Error getting block height")
	}
	fmt.Println("Block height is ", height)
```

Details of the Solana RPC API can be found here:
https://docs.solana.com/api/http

I am in no way associated with Solana, this is a side project to learn more about the blockchain and Golang.
