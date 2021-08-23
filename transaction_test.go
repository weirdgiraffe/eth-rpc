package ethrpc

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetTransactionByHash(t *testing.T) {
	client := testClient(t)
	ctx := context.Background()

	var exp Transaction
	getGoldenJSON(t, "transaction.json", &exp)

	got, err := client.GetTransactionByHash(ctx, exp.Hash)
	require.NoError(t, err)
	require.Equal(t, &exp, got)
}
