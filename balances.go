package ethrpc

import (
	"context"
	"encoding/json"
	"math/big"
	"strings"

	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
)

type Balance struct {
	Address          Address `json:"address"`
	BlockHash        Hash    `json:"blockHash"`
	BlockNumber      Number  `json:"blockNumber"`
	Data             Data    `json:"data"`
	LogIndex         Number  `json:"logIndex"`
	Removed          bool    `json:"removed"`
	Topics           []Hash  `json:"topics"`
	TransactionHash  Hash    `json:"transactionHash"`
	TransactionIndex Number  `json:"transactionIndex"`
}

type GetBalanceRequest struct {
	Token         Address
	Addr          Address
	Block         BlockNumber
	TokenDecimals int32
}

func (c *Client) GetBalanceETH(ctx context.Context, addr Address, block BlockNumber) (out decimal.Decimal, err error) {
	res, err := c.CallMethod(ctx, "eth_getBalance", []any{addr.String(), block.String()})
	if err != nil {
		return out, err
	}

	var evmValue string
	err = json.Unmarshal(res, &evmValue)
	if err != nil {
		return out, errors.Wrap(err, "failed to decode rpc result")
	}
	i, err := decodeBigint(evmValue)
	if err != nil {
		return out, errors.Wrap(err, "failed to decode balance")
	}
	return decimal.NewFromBigInt(i, -18), nil
}

func (c *Client) GetBalanceERC20(ctx context.Context, req GetBalanceRequest) (decimal.Decimal, error) {
	res, err := c.CallMethod(ctx, "eth_call", []any{
		map[string]any{
			"to":   req.Token.String(),
			"data": "0x70a08231" + padTo32(req.Addr.String()),
		},
		req.Block.String(),
	})
	if err != nil {
		return decimal.Zero, err
	}

	var evmValue string
	err = json.Unmarshal(res, &evmValue)
	if err != nil {
		return decimal.Zero, errors.Wrap(err, "failed to decode rpc result")
	}
	i, err := decodeBigint(evmValue)
	if err != nil {
		return decimal.Zero, errors.Wrap(err, "failed to decode balance")
	}
	return decimal.NewFromBigInt(i, -req.TokenDecimals), nil
}

func decodeBigint(hex string) (*big.Int, error) {
	if hex == "0x" {
		return new(big.Int).SetInt64(0), nil
	}
	i, ok := new(big.Int).SetString(hex, 0)
	if !ok {
		return nil, errors.New("failed to convert hex string to big.Int")
	}
	return i, nil
}

func padTo32(addr string) string {
	addr = strings.TrimPrefix(addr, "0x")
	padding := 64 - len(addr)
	return strings.Repeat("0", padding) + addr
}
