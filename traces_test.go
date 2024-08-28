package ethrpc

import (
	"context"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"github.com/weirdgiraffe/jsonrpc"
)

func TestTraceBlock(t *testing.T) {
	client := testClient(t)
	ctx := context.Background()

	var exp []Trace
	getGoldenJSON(t, "trace-block.json", &exp)

	got, err := client.TraceBlock(ctx, BlockNumber(exp[0].BlockNumber))
	if err != nil {
		var jsonrpcErr *jsonrpc.Error
		if errors.As(err, &jsonrpcErr) && jsonrpcErr.Code == -32601 {
			t.Skip("trace_block method is not supported by RPC server")
		}
		t.Fatal("unexpected error:", err)
	}
	require.Equal(t, exp, got)
}
