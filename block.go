package ethrpc

import (
	"context"
	"encoding/json"

	"github.com/pkg/errors"
)

type BlockInfo struct {
	Difficulty       Number  `json:"difficulty"`
	ExtraData        Data    `json:"extraData"`
	GasLimit         Number  `json:"gasLimit"`
	GasUsed          Number  `json:"gasUsed"`
	Hash             Hash    `json:"hash"`
	LogsBloom        Data    `json:"logsBloom"`
	Miner            Address `json:"miner"`
	MixHash          Hash    `json:"mixHash"`
	Nonce            Number  `json:"nonce"`
	Number           Number  `json:"number"`
	ParentHash       Hash    `json:"parentHash"`
	ReceiptsRoot     Hash    `json:"receiptsRoot"`
	Sha3Uncles       Hash    `json:"sha3Uncles"`
	Size             Number  `json:"size"`
	StateRoot        Hash    `json:"stateRoot"`
	Timestamp        Number  `json:"timestamp"`
	TotalDifficulty  Data    `json:"totalDifficulty"`
	TransactionsRoot Hash    `json:"transactionsRoot"`
	Uncles           []Hash  `json:"uncles"`
}

type Block struct {
	BlockInfo
	Transactions []Hash `json:"transactions"`
}

type DetailedBlock struct {
	BlockInfo
	Transactions []Transaction `json:"transactions"`
}

func (c *Client) GetBlockByHash(ctx context.Context, h Hash) (*Block, error) {
	res, err := c.impl.Call(ctx, "eth_getBlockByHash", h, false)
	if err != nil {
		return nil, err
	}
	if res.Error != nil {
		return nil, res.Error
	}
	var out Block
	err = json.Unmarshal(res.Result, &out)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode rpc result")
	}
	return &out, nil
}

func (c *Client) GetBlockByNumber(ctx context.Context, n BlockNumber) (*Block, error) {
	res, err := c.impl.Call(ctx, "eth_getBlockByNumber", n, false)
	if err != nil {
		return nil, err
	}
	if res.Error != nil {
		return nil, res.Error
	}
	var out Block
	err = json.Unmarshal(res.Result, &out)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode rpc result")
	}
	return &out, nil
}

func (c *Client) GetDetailedBlockByHash(ctx context.Context, h Hash) (*DetailedBlock, error) {
	res, err := c.impl.Call(ctx, "eth_getBlockByHash", h, true)
	if err != nil {
		return nil, err
	}
	if res.Error != nil {
		return nil, res.Error
	}
	var out DetailedBlock
	err = json.Unmarshal(res.Result, &out)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode rpc result")
	}
	return &out, nil
}
