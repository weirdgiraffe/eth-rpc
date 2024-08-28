package ethrpc

import (
	"context"
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
	res, err := c.CallMethod(ctx, "trace_block", []any{b})
	if err != nil {
		return nil, err
	}
	return jsonUnmarshalSlice[Trace](res)
}
