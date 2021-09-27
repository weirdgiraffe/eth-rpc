package ethrpc

import (
	"context"
	"encoding/json"

	"github.com/pkg/errors"
)

type Transaction struct {
	BlockHash        Hash        `json:"blockHash"`
	BlockNumber      BlockNumber `json:"blockNumber"`
	TransactionIndex Number      `json:"transactionIndex"`

	Hash  Hash    `json:"hash"`
	Type  Number  `json:"type"`
	Nonce Number  `json:"nonce"`
	From  Address `json:"from"`
	To    Address `json:"to"`
	Value Data    `json:"value"`
	Input Data    `json:"input"`

	Gas      Number `json:"gas"`
	GasPrice Number `json:"gasPrice"`

	MaxPriorityFeePerGas Number `json:"maxPriorityFeePerGas"`
	MaxFeePerGas         Number `json:"maxFeePerGas"`
	TransactionType      Number `json:"transactionType"`

	R Data `json:"r"`
	S Data `json:"s"`
	V Data `json:"v"`
}

func (c *Client) GetTransactionByHash(ctx context.Context, h Hash) (*Transaction, error) {
	res, err := c.impl.Call(ctx, "eth_getTransactionByHash", h)
	if err != nil {
		return nil, err
	}
	if res.Error != nil {
		return nil, res.Error
	}
	var out Transaction
	err = json.Unmarshal(res.Result, &out)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode rpc result")
	}
	return &out, nil
}
