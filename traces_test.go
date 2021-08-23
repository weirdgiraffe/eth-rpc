package ethrpc

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTraceBlock(t *testing.T) {
	client := testClient(t)
	ctx := context.Background()

	var exp []Trace
	getGoldenJSON(t, "trace-block.json", &exp)

	got, err := client.TraceBlock(ctx, BlockNumber(exp[0].BlockNumber))
	require.NoError(t, err)
	require.Equal(t, exp, got)
}
