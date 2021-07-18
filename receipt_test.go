package ethrpc

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetTransactionReceipt(t *testing.T) {
	client := testClient(t)
	ctx := context.Background()

	var exp TransactionReceipt
	getGolden(t, "receipt.json", &exp)

	got, err := client.GetTransactionReceipt(ctx, exp.TransactionHash)
	require.NoError(t, err)
	require.Equal(t, &exp, got)
}
