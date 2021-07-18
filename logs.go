package ethrpc

type LogEntry struct {
	Address          Address     `json:"address"`
	BlockHash        Hash        `json:"blockHash"`
	BlockNumber      BlockNumber `json:"blockNumber"`
	Data             Data        `json:"data"`
	LogIndex         Number      `json:"logIndex"`
	Removed          bool        `json:"removed"`
	Topics           []Hash      `json:"topics"`
	TransactionHash  Hash        `json:"transactionHash"`
	TransactionIndex Number      `json:"transactionIndex"`
}
