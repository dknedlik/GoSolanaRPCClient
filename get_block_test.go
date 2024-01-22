package golangsolanarpc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBlock(t *testing.T) {
	client := SolanaClient{
		RpcEndpoint: "https://api.devnet.solana.com",
	}
	response, err := client.GetBlock(262727887)
	if err != nil {
		t.Fatal("Error getting block")
	}
	//Check the header data
	assert.Equal(t, int64(1701612610), *response.BlockTime)
	assert.Equal(t, uint64(251033213), *response.BlockHeight)
	assert.Equal(t, "AU4no55BVtW8mfWLKgwuMtmk4nVNamDez9KwgyXDaLnb", response.Blockhash)
	assert.Equal(t, uint64(262727886), response.ParentSlot)
	assert.Equal(t, "9F7VaceiyEdhnKVj3peZG85WCEcviCAz5HVfpfZifhKN", response.PreviousBlockhash)

	//check the rewards data
	assert.Equal(t, 6, len(response.Rewards))
	var r Reward
	for i := range response.Rewards {
		if response.Rewards[i].Pubkey == "dv1ZAGvdsz5hHLwWXsVnM94hWf1pjbKVau1QVkaMJ92" {
			r = response.Rewards[i]
		}
	}
	assert.NotNil(t, r)
	assert.Equal(t, "Rent", *r.RewardType)
	assert.Equal(t, uint64(2748747296401), r.PostBalance)
	assert.Nil(t, r.Commission)

	//check the transaction data
	assert.Equal(t, 30, len(response.Transactions))
	var txn TransactionDetails
	for i := range response.Transactions {
		if response.Transactions[i].Transaction.Signatures[0] == "41bGBMeNKVz9RUyCgi7Aua9HQc4pj5Mt376Hu6YAaHWpCUWC6ew2udrrtm9oopwKLFMcsDB7FPJbg5NUigXXHj8q" {
			txn = response.Transactions[i]
		}
	}
	assert.NotNil(t, txn)
	//test meta
	assert.Equal(t, uint64(79344), *txn.Meta.ComputeUnitsConsumed)
	assert.Equal(t, uint64(5001), txn.Meta.Fee)
	assert.Nil(t, txn.Meta.Err)
	assert.NotNil(t, txn.Meta.InnerInstructions)
	inner := (*txn.Meta.InnerInstructions)[0]
	assert.Equal(t, 2, inner.Index)
	assert.NotNil(t, inner.Instructions[0])
	ii := inner.Instructions[0]
	assert.Equal(t, "3px3hhrW2tYw", ii.Data)
	assert.Equal(t, 16, ii.ProgramIdIndex)
	assert.Equal(t, 2, *ii.StackHeight)
	assert.Equal(t, 3, len(ii.Accounts))
	v := []int{5, 10, 12}
	assert.Equal(t, ii.Accounts, v)

	assert.Equal(t, 19, len(txn.Meta.LogMessages))
	assert.Equal(t, "Program ComputeBudget111111111111111111111111111111 invoke [1]", txn.Meta.LogMessages[0])
	assert.Equal(t, 18, len(txn.Meta.PostBalances))
	assert.Equal(t, uint64(638240793899), txn.Meta.PostBalances[0])

	assert.Equal(t, 3, len(txn.Meta.PostTokenBalances))
	pb := txn.Meta.PostTokenBalances[0]
	assert.NotNil(t, pb)
	assert.Equal(t, 3, pb.AccountIndex)
	assert.Equal(t, "So11111111111111111111111111111111111111112", pb.Mint)
	assert.Equal(t, "CyZuD7RPDcrqCGbNvLCyqk6Py9cEZTKmNKujfPi3ynDd", *pb.Owner)
	assert.Equal(t, "TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA", *pb.ProgramId)

	assert.Equal(t, "14749747276", pb.UiTokenAmount.Amount)
	assert.Equal(t, 9, pb.UiTokenAmount.Decimals)
	assert.Equal(t, float64(14.749747276), *pb.UiTokenAmount.UiAmount)
	assert.Equal(t, "14.749747276", pb.UiTokenAmount.UiAmountString)

	assert.Equal(t, 18, len(txn.Meta.PreBalances))
	assert.Equal(t, uint64(5317440), txn.Meta.PreBalances[2])

	assert.Equal(t, 3, len(txn.Meta.PreTokenBalances))
	pb = txn.Meta.PreTokenBalances[1]

	assert.NotNil(t, pb)
	assert.Equal(t, 5, pb.AccountIndex)
	assert.Equal(t, "So11111111111111111111111111111111111111112", pb.Mint)
	assert.Equal(t, "CyZuD7RPDcrqCGbNvLCyqk6Py9cEZTKmNKujfPi3ynDd", *pb.Owner)
	assert.Equal(t, "TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA", *pb.ProgramId)

	assert.Equal(t, "3314250000", pb.UiTokenAmount.Amount)
	assert.Equal(t, 9, pb.UiTokenAmount.Decimals)
	assert.Equal(t, float64(3.31425), *pb.UiTokenAmount.UiAmount)
	assert.Equal(t, "3.31425", pb.UiTokenAmount.UiAmountString)
	//test transaction object
	assert.Equal(t, 18, len(txn.Transaction.Message.AccountKeys))
	assert.Equal(t, "67DVysERtYeq7vbUQJtrBzcFGUkPZGniz4qeJAomgCDJ", txn.Transaction.Message.AccountKeys[3])
	assert.Equal(t, 0, txn.Transaction.Message.Header.NumReadonlySignedAccounts)
	assert.Equal(t, 7, txn.Transaction.Message.Header.NumReadonlyUnsignedAccounts)
	assert.Equal(t, 1, txn.Transaction.Message.Header.NumRequiredSignatures)

	assert.Equal(t, 3, len(txn.Transaction.Message.Instructions))
	mi := txn.Transaction.Message.Instructions[2]
	testAccounts := []int{1,
		6,
		0,
		17,
		0,
		8,
		13,
		9,
		5,
		16,
		12,
		7,
		14,
		2,
		6,
		3,
		10,
		4}
	assert.Equal(t, testAccounts, mi.Accounts)
	assert.Equal(t, "3hAw8ppjnbTHpwEQYxS9UcDGdCi2Ts1ocY56oZyvyFEpfoNrfNbuNmFMK9hXWUPBRywAnQfKPiVDuwGBsUM9pBseWmAg74dR5cfT88Asv3iFAvXKGAkQ6kqfg1HmMpn2SBj6y96zjy7eXNd6RZW8p", mi.Data)
	assert.Equal(t, "CWnX2TA5LVeiSE26qRL8AffbkfGtBE3A38vtSoyihP7G", txn.Transaction.Message.RecentBlockhash)

	assert.Equal(t, 1, len(txn.Transaction.Signatures))
	assert.Equal(t, "41bGBMeNKVz9RUyCgi7Aua9HQc4pj5Mt376Hu6YAaHWpCUWC6ew2udrrtm9oopwKLFMcsDB7FPJbg5NUigXXHj8q", txn.Transaction.Signatures[0])
}
