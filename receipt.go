package ethrpc

import (
	"context"
	"encoding/json"

	"github.com/pkg/errors"
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
	res, err := c.impl.Call(ctx, "eth_getTransactionReceipt", txHash)
	if err != nil {
		return nil, err
	}
	if res.Error != nil {
		return nil, res.Error
	}
	var out TransactionReceipt
	err = json.Unmarshal(res.Result, &out)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode rpc result")
	}
	return &out, nil
}
