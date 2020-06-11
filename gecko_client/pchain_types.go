package gecko_client

// ============= Blockchain ====================
type UnsignedTransactionInfo struct {
	UnsignedTx string 	`json:"unsignedTx"`
}

type CreateUnsignedTransactionResponse struct {
	JsonRpcVersion string                  `json:"jsonrpc"`
	Result         UnsignedTransactionInfo `json:"result"`
	Id             int                     `json:"id"`
}

type CreateBlockchainResponse struct {
	JsonRpcVersion string                  `json:"jsonrpc"`
	Result         UnsignedTransactionInfo `json:"result"`
	Id             int                     `json:"id"`
}

type BlockchainStatus struct {
	Status string	`json:"status"`
}

type GetBlockchainStatusResponse struct {
	JsonRpcVersion string	`json:"jsonrpc"`
	Result BlockchainStatus	`json:"result"`
	Id int	`json:"id"`
}

// ============= Accounts ====================
type AccountAddressInfo struct {
	Address string 	`json:"address"`
}

type CreateAccountResponse struct {
	JsonRpcVersion string             `json:"jsonrpc"`
	Result         AccountAddressInfo `json:"result"`
	Id             int                `json:"id"`
}

type ImportKeyResponse struct {
	JsonRpcVersion string             `json:"jsonrpc"`
	Result         AccountAddressInfo `json:"result"`
	Id             int                `json:"id"`
}

type PrivateKeyInfo struct {
	PrivateKey string	`json:"privateKey"`
}

type ExportKeyResponse struct {
	JsonRpcVersion string	`json:"jsonrpc"`
	Result PrivateKeyInfo	`json:"result"`
	Id int	`json:"id"`
}

type AccountInfo struct {
	Address string 	`json:"address"`
	Nonce string 	`json:"nonce"`
	Balance string 	`json:"balance"`
}

type GetAccountResponse struct {
	JsonRpcVersion string      `json:"jsonrpc"`
	Result         AccountInfo `json:"result"`
	Id             int         `json:"id"`
}

type AccountList struct {
	Accounts	[]AccountInfo	`json:"accounts"`
}

type ListAccountsResponse struct {
	JsonRpcVersion string      `json:"jsonrpc"`
	Result         AccountList `json:"result"`
	Id             int         `json:"id"`
}

// ============= Validators ====================
type Validator struct {
	StartTime string `json:"startTime"`
	EndTime string	`json:"endTime"`
	StakeAmount string	`json:"stakeAmount"`
	Id string	`json:"id"`
}

type ValidatorList struct {
	Validators []Validator	`json:"validators"`
}

type GetValidatorsResponse struct {
	JsonRpcVersion string	`json:"jsonrpc"`
	Result ValidatorList	`json:"result"`
	Id int	`json:"id"`
}

// ============= Subnet ========================
type Subnet struct {
	Id string `json:"id"`
	ControlKeys []string `json:"controlKeys"`
	Threshold string `json:"threshold"`
}

type SubnetList struct {
	Subnets []Subnet	`json:"subnets"`
}

type GetSubnetsResponse struct {
	JsonRpcVersion string	`json:"jsonrpc"`
	Result SubnetList	`json:"result"`
	Id int	`json:"id"`
}
