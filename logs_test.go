package ethrpc

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/weirdgiraffe/jsonrpc"
)

type MockMethodCaller struct {
	Result []byte
}

func (mock MockMethodCaller) Call(_ context.Context, _ string, _ ...interface{}) (*jsonrpc.Response, error) {
	return &jsonrpc.Response{
		Version: "2.0",
		ID:      1,
		Result:  mock.Result,
		Error:   nil,
	}, nil
}

func TestConvertLogEntry(t *testing.T) {
	caller := MockMethodCaller{Result: []byte(`[{
	"address":"0xdac17f958d2ee523a2206206994597c13d831ec7",
	"blockHash":"0xf3739a5b6e0bf365f459c957ce991ba771f5e669889b1dfc8aedb33cabd057f9",
	"blockNumber":"0x98c6b9",
	"data":"0x000000000000000000000000000000000000000000000000000000012a05f200",
	"logIndex":"0x3",
	"removed":false,
	"topics":[
		"0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef",
		"0x000000000000000000000000fdb16996831753d5331ff813c29a93c76834a0ad",
		"0x000000000000000000000000af398e7933e6a5cce3e5730878cd90680012e305"
	],
	"transactionHash":"0xda329cb39bf9703f88db4ca8da41754ca09fedddad039fa118f202f12f619fb4",
	"transactionIndex":"0x11"
}]`)}
	c := NewLogsClient(caller)
	ctx := context.Background()
	_, err := c.GetLogsForBlockNumber(ctx, EarliestBlock)
	require.NoError(t, err)
}
