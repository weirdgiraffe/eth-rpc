package ethrpc

import (
	"context"
	"encoding/json"

	"github.com/pkg/errors"
	"github.com/weirdgiraffe/jsonrpc"
)

type MethodCaller interface {
	Call(ctx context.Context, method string, params ...interface{}) (*jsonrpc.Response, error)
}

type LogsClient struct {
	impl MethodCaller
}

func NewLogsClient(impl MethodCaller) *LogsClient {
	return &LogsClient{impl: impl}
}

type LogEntry struct {
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

func (c *LogsClient) GetLogsForBlockNumber(ctx context.Context, n BlockNumber, topic ...Hash) ([]LogEntry, error) {
	return c.GetLogsForBlockRange(ctx, n, n, topic...)
}

func (c *LogsClient) GetLogsForBlockRange(ctx context.Context, from, to BlockNumber, topic ...Hash) ([]LogEntry, error) {
	l := make([]string, len(topic))
	for i := range topic {
		l[i] = topic[i].String()
	}
	res, err := c.impl.Call(ctx, "eth_getLogs",
		map[string]interface{}{
			"fromBlock": from.String(),
			"toBlock":   to.String(),
			"topics":    l,
		})
	if err != nil {
		return nil, err
	}
	if res.Error != nil {
		return nil, res.Error
	}

	var out []LogEntry
	err = json.Unmarshal(res.Result, &out)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode rpc result")
	}
	return out, nil
}

/*
func convertLogEntry(raw rawLogEntry) (out models.LogEntry, err error) {
	defer func() {
		if r := recover(); r != nil {
			if e, ok := r.(error); ok {
				err = e
				return
			}
		}
	}()
	handleErr := func(err error, reason string) {
		if err != nil {
			panic(errors.Wrap(err, reason))
		}
	}

	out.Address, err = AddressFromString(raw.Address)
	handleErr(err, "address")
	out.BlockHash, err = HashFromString(raw.BlockHash)
	handleErr(err, "blockHash")
	out.BlockNumber, err = Uint64FromString(raw.BlockNumber)
	handleErr(err, "blockNumber")
	out.Data, err = DataFromString(raw.Data)
	handleErr(err, "data")
	out.LogIndex, err = Uint64FromString(raw.LogIndex)
	handleErr(err, "logIndex")
	out.Removed = raw.Removed
	out.Topics = make([][]byte, len(raw.Topics))
	for i := range raw.Topics {
		out.Topics[i], err = TopicFromString(raw.Topics[i])
		handleErr(err, fmt.Sprintf("topic[%d]", i))
	}
	out.TxHash, err = HashFromString(raw.TransactionHash)
	handleErr(err, "TransactionHash")
	out.TxIndex, err = Uint64FromString(raw.TransactionIndex)
	handleErr(err, "TransactionIndex")
	return out, nil
}
*/
