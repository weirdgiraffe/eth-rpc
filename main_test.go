package ethrpc

import (
	"encoding/hex"
	"encoding/json"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

var TestHttpURL string

func TestMain(m *testing.M) {
	TestHttpURL = os.Getenv("TEST_RPCURL")
	os.Exit(m.Run())

}

func testClient(t *testing.T) *Client {
	t.Helper()

	if TestHttpURL == "" {
		t.Skip("TEST_RPCURL env variable is not set")
	}
	return NewClient(NewHTTP(TestHttpURL))
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
