package ethrpc

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetBlockByHash(t *testing.T) {
	client := testClient(t)
	ctx := context.Background()

	var exp Block
	getGolden(t, "block.json", &exp)

	got, err := client.GetBlockByHash(ctx, exp.Hash)
	require.NoError(t, err)
	require.Equal(t, &exp, got)
}

func TestGetDetailedBlockByHash(t *testing.T) {
	client := testClient(t)
	ctx := context.Background()

	var exp DetailedBlock
	getGolden(t, "detailed-block.json", &exp)

	got, err := client.GetDetailedBlockByHash(ctx, exp.Hash)
	require.NoError(t, err)
	require.Equal(t, &exp, got)
}
