package ethrpc

import (
	"encoding/hex"
	"encoding/json"
	"log/slog"
	"os"
	"strings"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"
	"github.com/weirdgiraffe/jsonrpc"
)

var rpcURL string

func TestMain(m *testing.M) {
	err := godotenv.Load(".env")
	if err != nil {
		slog.Warn("failed to load .env file")
	}
	rpcURL = os.Getenv("TEST_RPCURL")
	exitCode := m.Run()
	os.Exit(exitCode)
}

func testClient(t *testing.T) *Client {
	t.Helper()

	if rpcURL == "" {
		t.Skip("TEST_RPCURL env variable is not set")
	}

	impl := jsonrpc.NewHTTPClient(rpcURL)
	return NewClient(impl)
}

func getGolden(t *testing.T, relPath string) []byte {
	t.Helper()

	b, err := os.ReadFile("golden/" + relPath)
	require.NoError(t, err)
	return b
}

func getGoldenJSON(t *testing.T, relPath string, dst interface{}) {
	t.Helper()

	b := getGolden(t, relPath)
	require.NoError(t, json.Unmarshal(b, dst))
}

func getGoldenHEX(t *testing.T, relPath string) []byte {
	t.Helper()

	b := getGolden(t, relPath)
	s := strings.TrimSpace(strings.TrimPrefix(string(b), "0x"))
	b, err := hex.DecodeString(s)
	require.NoError(t, err)
	return b
}
