package ethrpc

import (
	"context"

	"github.com/weirdgiraffe/jsonrpc"
)

type MethodCaller interface {
	Call(ctx context.Context, method string, params ...interface{}) (*jsonrpc.Response, error)
}

type Client struct {
	impl MethodCaller
}

func NewClient(impl MethodCaller) *Client {
	return &Client{impl: impl}
}
