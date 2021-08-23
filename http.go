package ethrpc

import (
	"context"
	"sync/atomic"

	"github.com/weirdgiraffe/jsonrpc"
)

type HTTP struct {
	seq  uint64
	http *jsonrpc.ClientHTTP
	dump *Dumper
}

func NewHTTP(url string, opt ...DumpOption) *HTTP {
	return &HTTP{
		http: jsonrpc.NewClientHTTP(url),
		dump: NewDumper(opt...),
	}
}

func (c *HTTP) Call(ctx context.Context, method string, params ...interface{}) (*jsonrpc.Response, error) {
	req := jsonrpc.NewRequest(
		atomic.AddUint64(&c.seq, 1),
		method,
		params...,
	)
	c.dump.DumpRequest(req)
	res, err := c.http.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	c.dump.DumpResponse(res)
	return res, err
}
