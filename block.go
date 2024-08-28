package ethrpc

import (
	"context"
)

type Block struct {
	Transactions []Hash `json:"transactions"`
	BlockInfo
}

type DetailedBlock struct {
	Transactions []Transaction `json:"transactions"`
	BlockInfo
}

type BlockInfo struct {
	ParentHash       Hash    `json:"parentHash"`
	TransactionsRoot Hash    `json:"transactionsRoot"`
	Uncles           []Hash  `json:"uncles"`
	TotalDifficulty  Data    `json:"totalDifficulty"`
	Hash             Hash    `json:"hash"`
	LogsBloom        Data    `json:"logsBloom"`
	Miner            Address `json:"miner"`
	StateRoot        Hash    `json:"stateRoot"`
	Sha3Uncles       Hash    `json:"sha3Uncles"`
	ExtraData        Data    `json:"extraData"`
	MixHash          Hash    `json:"mixHash"`
	ReceiptsRoot     Hash    `json:"receiptsRoot"`
	BasFeePerGas     Number  `json:"baseFeePerGas"`
	Size             Number  `json:"size"`
	Difficulty       Number  `json:"difficulty"`
	Timestamp        Number  `json:"timestamp"`
	GasUsed          Number  `json:"gasUsed"`
	Number           Number  `json:"number"`
	GasLimit         Number  `json:"gasLimit"`
	Nonce            Number  `json:"nonce"`
}

func (c *Client) BlockNumber(ctx context.Context) (*BlockNumber, error) {
	res, err := c.CallMethod(ctx, "eth_blockNumber", nil)
	if err != nil {
		return nil, err
	}
	return jsonUnmarshalStruct[BlockNumber](res)
}

func (c *Client) GetBlockByHash(ctx context.Context, h Hash) (*Block, error) {
	res, err := c.CallMethod(ctx, "eth_getBlockByHash", []any{h, false})
	if err != nil {
		return nil, err
	}
	return jsonUnmarshalStruct[Block](res)
}

func (c *Client) GetDetailedBlockByHash(ctx context.Context, h Hash) (*DetailedBlock, error) {
	res, err := c.CallMethod(ctx, "eth_getBlockByHash", []any{h, true})
	if err != nil {
		return nil, err
	}
	return jsonUnmarshalStruct[DetailedBlock](res)
}

func (c *Client) GetBlockByNumber(ctx context.Context, n BlockNumber) (*Block, error) {
	res, err := c.CallMethod(ctx, "eth_getBlockByNumber", []any{n, false})
	if err != nil {
		return nil, err
	}
	return jsonUnmarshalStruct[Block](res)
}

func (c *Client) GetDetailedBlockByNumber(ctx context.Context, n BlockNumber) (*DetailedBlock, error) {
	res, err := c.CallMethod(ctx, "eth_getBlockByNumber", []any{n, true})
	if err != nil {
		return nil, err
	}
	return jsonUnmarshalStruct[DetailedBlock](res)
}
