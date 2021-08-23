package ethrpc

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseContract(t *testing.T) {
	usdt := getGoldenHEX(t, "usdt.hex")
	info, err := contractInfo(usdt)
	require.NoError(t, err)
	for _, fn := range info.Func {
		t.Logf("%s", hex.EncodeToString(fn))
	}
	for _, evt := range info.Event {
		t.Logf("%s", hex.EncodeToString(evt))
	}
}
