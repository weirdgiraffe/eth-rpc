package ethrpc

import (
	"context"
	"encoding/json"
	"fmt"
	"sync/atomic"

	"github.com/weirdgiraffe/jsonrpc"
)

type Caller interface {
	Call(ctx context.Context, req *jsonrpc.Request) (*jsonrpc.Response, error)
	BatchCall(ctx context.Context, batch []*jsonrpc.Request) ([]*jsonrpc.Response, error)
}

type Batch struct {
	rs     *RequestSequencer
	caller Caller
	Calls  []*jsonrpc.Request
}

func (b *Batch) Add(method string, params ...any) *Batch {
	req := b.rs.NewRequest(method, params)
	b.Calls = append(b.Calls, req)
	return b
}

func (b *Batch) Submit(ctx context.Context) ([]*jsonrpc.Response, error) {
	res, err := b.caller.BatchCall(ctx, b.Calls)
	if err != nil {
		return nil, err
	}
	return res, err
}

type RequestSequencer struct {
	seq atomic.Uint64
}

func (rs *RequestSequencer) NewRequest(method string, params any) *jsonrpc.Request {
	return jsonrpc.NewRequest(rs.seq.Add(1), method, params)
}

type Client struct {
	rs   *RequestSequencer
	impl Caller
}

func NewClient(impl Caller) *Client {
	return &Client{
		rs:   &RequestSequencer{},
		impl: impl,
	}
}

func (c *Client) CallMethod(ctx context.Context, method string, params any) ([]byte, error) {
	req := c.rs.NewRequest(method, params)

	res, err := c.impl.Call(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to call jsonrpc: %w", err)
	}
	if res.Error != nil {
		return nil, fmt.Errorf("jsonrpc error: %w", res.Error)
	}
	return res.Result, nil
}

func jsonUnmarshalStruct[T any](b []byte) (*T, error) {
	var dst T
	err := json.Unmarshal(b, &dst)
	if err != nil {
		fmt.Println("error is here", err)
		return nil, fmt.Errorf("failed to decode rpc result: %w", err)
	}
	return &dst, nil
}

func jsonUnmarshalSlice[T any](b []byte) ([]T, error) {
	var dst []T
	err := json.Unmarshal(b, &dst)
	if err != nil {
		return nil, fmt.Errorf("failed to decode rpc result: %w", err)
	}
	return dst, nil
}
