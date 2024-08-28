package ethrpc

import (
	"context"
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
	res, err := c.CallMethod(ctx, "eth_getTransactionByHash", []any{h})
	if err != nil {
		return nil, err
	}
	return jsonUnmarshalStruct[Transaction](res)
}
