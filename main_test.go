package ethrpc

import (
	"encoding/json"
	"os"
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

func getGolden(t *testing.T, relPath string, dst interface{}) {
	t.Helper()

	b, err := os.ReadFile("golden/" + relPath)
	require.NoError(t, err)
	require.NoError(t, json.Unmarshal(b, dst))
}
