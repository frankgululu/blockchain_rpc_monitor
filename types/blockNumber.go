package types

type BlockNumResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	Id      int    `json:"id"`
	Result  struct {
		BaseFeePerGas         string        `json:"baseFeePerGas"`
		BlobGasUsed           string        `json:"blobGasUsed"`
		Difficulty            string        `json:"difficulty"`
		ExcessBlobGas         string        `json:"excessBlobGas"`
		ExtraData             string        `json:"extraData"`
		GasLimit              string        `json:"gasLimit"`
		GasUsed               string        `json:"gasUsed"`
		Hash                  string        `json:"hash"`
		LogsBloom             string        `json:"logsBloom"`
		Miner                 string        `json:"miner"`
		MixHash               string        `json:"mixHash"`
		Nonce                 string        `json:"nonce"`
		Number                string        `json:"number"`
		ParentBeaconBlockRoot string        `json:"parentBeaconBlockRoot"`
		ParentHash            string        `json:"parentHash"`
		ReceiptsRoot          string        `json:"receiptsRoot"`
		Sha3Uncles            string        `json:"sha3Uncles"`
		Size                  string        `json:"size"`
		StateRoot             string        `json:"stateRoot"`
		Timestamp             string        `json:"timestamp"`
		TotalDifficulty       string        `json:"totalDifficulty"`
		Transactions          []string      `json:"transactions"`
		TransactionsRoot      string        `json:"transactionsRoot"`
		Uncles                []interface{} `json:"uncles"`
		Withdrawals           []interface{} `json:"withdrawals"`
		WithdrawalsRoot       string        `json:"withdrawalsRoot"`
	} `json:"result"`
}
