package ethrpc

import (
	"context"
)

type TransactionReceipt struct {
	BlockHash   Hash   `json:"blockHash"`
	BlockNumber Number `json:"blockNumber"`

	TransactionHash  Hash    `json:"transactionHash"`
	TransactionIndex Number  `json:"transactionIndex"`
	Type             Number  `json:"type"`
	From             Address `json:"from"`
	To               Address `json:"to"`
	Status           Number  `json:"status"`
	Root             Hash    `json:"root"`

	ContractAddress Address `json:"contractAddress"`

	GasUsed           Number `json:"gasUsed"`
	CumulativeGasUsed Number `json:"cumulativeGasUsed"`
	EffectiveGasPrice Number `json:"effectiveGasPrice"`

	LogsBloom Data       `json:"logsBloom"`
	Logs      []LogEntry `json:"logs"`
}

func (c *Client) GetTransactionReceipt(ctx context.Context, txHash Hash) (*TransactionReceipt, error) {
	res, err := c.CallMethod(ctx, "eth_getTransactionReceipt", []any{txHash})
	if err != nil {
		return nil, err
	}
	return jsonUnmarshalStruct[TransactionReceipt](res)
}
