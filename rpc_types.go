package golangsolanarpc

type Commitment string

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

type RPCResponse struct {
	JSONRPC string      `json:"jsonrpc"`
	Result  interface{} `json:"result,omitempty"`
	Error   *RPCError   `json:"error,omitempty"`
	ID      string      `json:"id"`
}

type AccountInfoRPCResult struct {
	Context interface{} `json:"context"`
	Value   AccountInfo `json:"value"`
}

type BalanceRPCResult struct {
	Context interface{} `json:"context"`
	Value   uint64      `json:"value"`
}

type RpcEncoding struct {
	Encoding string `json:"encoding"`
}

type SolAccountInfoResponse struct {
	Context interface{} `json:"context"`
	Value   AccountInfo `json:"value"`
}

// Transaction represents a Solana transaction
type Transaction struct {
	Signatures          []string             `json:"signatures"`
	Message             Message              `json:"message"`
	AddressTableLookups []AddressTableLookup `json:"addressTableLookups,omitempty"`
}

// Message represents the content of the transaction
type Message struct {
	AccountKeys       []string           `json:"accountKeys"`
	Header            Header             `json:"header"`
	RecentBlockhash   string             `json:"recentBlockhash"`
	Instructions      []Instruction      `json:"instructions"`
	InnerInstructions []InnerInstruction `json:"innerInstructions,omitempty"`
	TokenBalances     []TokenBalance     `json:"tokenBalances,omitempty"`
}

// Header details the account types and signatures required by the transaction
type Header struct {
	NumRequiredSignatures       int `json:"numRequiredSignatures"`
	NumReadonlySignedAccounts   int `json:"numReadonlySignedAccounts"`
	NumReadonlyUnsignedAccounts int `json:"numReadonlyUnsignedAccounts"`
}

// Instruction is a program instruction that will be executed
type Instruction struct {
	ProgramIdIndex int    `json:"programIdIndex"`
	Accounts       []int  `json:"accounts"`
	Data           string `json:"data"`
	StackHeight    *int   `json:"stackHeight"`
}

// InnerInstruction represents the cross-program instructions invoked
type InnerInstruction struct {
	Index        int           `json:"index"`
	Instructions []Instruction `json:"instructions"`
}

// TokenBalance represents the token balance for an account
type TokenBalance struct {
	AccountIndex  int           `json:"accountIndex"`
	Mint          string        `json:"mint"`
	Owner         *string       `json:"owner,omitempty"`
	ProgramId     *string       `json:"programId,omitempty"`
	UiTokenAmount UiTokenAmount `json:"uiTokenAmount"`
}

// UiTokenAmount represents the amount of tokens in different formats
type UiTokenAmount struct {
	Amount         string   `json:"amount"`
	Decimals       int      `json:"decimals"`
	UiAmount       *float64 `json:"uiAmount,omitempty"`
	UiAmountString string   `json:"uiAmountString"`
}

// AddressTableLookup represents the address table lookups used by a transaction
type AddressTableLookup struct {
	AccountKey      string `json:"accountKey"`
	WritableIndexes []int  `json:"writableIndexes"`
	ReadonlyIndexes []int  `json:"readonlyIndexes"`
}

type BlockResponse struct {
	Blockhash         string               `json:"blockhash"`
	PreviousBlockhash string               `json:"previousBlockhash"`
	ParentSlot        uint64               `json:"parentSlot"`
	Transactions      []TransactionDetails `json:"transactions"`
	Signatures        []string             `json:"signatures,omitempty"`
	Rewards           []Reward             `json:"rewards,omitempty"`
	BlockTime         *int64               `json:"blockTime"`
	BlockHeight       *uint64              `json:"blockHeight"`
}

type TransactionDetails struct {
	Transaction Transaction `json:"transaction"`
	Meta        Meta        `json:"meta"`
}

// Meta represents the transaction status metadata
type Meta struct {
	Err                  *TransactionError   `json:"err"`
	Fee                  uint64              `json:"fee"`
	PreBalances          []uint64            `json:"preBalances"`
	PostBalances         []uint64            `json:"postBalances"`
	InnerInstructions    *[]InnerInstruction `json:"innerInstructions"`
	PreTokenBalances     []TokenBalance      `json:"preTokenBalances,omitempty"`
	PostTokenBalances    []TokenBalance      `json:"postTokenBalances,omitempty"`
	LogMessages          []string            `json:"logMessages,omitempty"`
	Rewards              []Reward            `json:"rewards,omitempty"`
	LoadedAddresses      *LoadedAddresses    `json:"loadedAddresses,omitempty"`
	ReturnData           *ReturnData         `json:"returnData,omitempty"`
	ComputeUnitsConsumed *uint64             `json:"computeUnitsConsumed,omitempty"`
	Version              interface{}         `json:"version,omitempty"` // could be string or number
}
type Reward struct {
	Pubkey      string  `json:"pubkey"`
	Lamports    int64   `json:"lamports"`
	PostBalance uint64  `json:"postBalance"`
	RewardType  *string `json:"rewardType,omitempty"`
	Commission  *uint8  `json:"commission,omitempty"`
}

type LoadedAddresses struct {
	Writable []string `json:"writable"`
	Readonly []string `json:"readonly"`
}

type ReturnData struct {
	ProgramId string   `json:"programId"`
	Data      []string `json:"data"` // base-64 encoded binary data
}

type TransactionError string

const (
	AccountInUse                       TransactionError = "Account in use"
	AccountLoadedTwice                 TransactionError = "Account loaded twice"
	AccountNotFound                    TransactionError = "Account not found"
	ProgramAccountNotFound             TransactionError = "Program account not found"
	InsufficientFundsForFee            TransactionError = "Insufficient funds for fee"
	InvalidAccountForFee               TransactionError = "Invalid account for fee"
	AlreadyProcessed                   TransactionError = "Already processed"
	BlockhashNotFound                  TransactionError = "Blockhash not found"
	InstructionError                   TransactionError = "Instruction error"
	CallChainTooDeep                   TransactionError = "Call chain too deep"
	MissingSignatureForFee             TransactionError = "Missing signature for fee"
	InvalidAccountIndex                TransactionError = "Invalid account index"
	SignatureFailure                   TransactionError = "Signature failure"
	InvalidProgramForExecution         TransactionError = "Invalid program for execution"
	SanitizeFailure                    TransactionError = "Sanitize failure"
	ClusterMaintenance                 TransactionError = "Cluster maintenance"
	AccountBorrowOutstanding           TransactionError = "Account borrow outstanding"
	WouldExceedMaxBlockCostLimit       TransactionError = "Would exceed max block cost limit"
	UnsupportedVersion                 TransactionError = "Unsupported version"
	InvalidWritableAccount             TransactionError = "Invalid writable account"
	WouldExceedMaxAccountCostLimit     TransactionError = "Would exceed max account cost limit"
	WouldExceedMaxAccountDataCostLimit TransactionError = "Would exceed max account data cost limit"
)
