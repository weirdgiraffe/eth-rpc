package ethrpc

import (
	"context"
	"encoding/json"

	"github.com/pkg/errors"
)

type TraceAction struct {
	CallType string  `json:"callType"`
	From     Address `json:"from"`
	Gas      Number  `json:"gas"`
	Input    Data    `json:"input"`
	To       Address `json:"to"`
	Value    Data    `json:"value"`
}

type TraceResult struct {
	GasUsed Number `json:"gasUsed"`
	Output  Data   `json:"output"`
}

type Trace struct {
	Action              TraceAction `json:"action"`
	BlockHash           Hash        `json:"blockHash"`
	BlockNumber         int64       `json:"blockNumber"`
	Result              TraceResult `json:"result"`
	Subtraces           int         `json:"subtraces"`
	TraceAddress        []int       `json:"traceAddress"`
	TransactionHash     Hash        `json:"transactionHash"`
	TransactionPosition int         `json:"transactionPosition"`
	Type                string      `json:"type"`
}

func (c *Client) TraceBlock(ctx context.Context, b BlockNumber) ([]Trace, error) {
	res, err := c.impl.Call(ctx, "trace_block", b)
	if err != nil {
		return nil, err
	}
	if res.Error != nil {
		return nil, res.Error
	}

	var out []Trace
	err = json.Unmarshal(res.Result, &out)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode rpc result")
	}
	return out, nil
}
