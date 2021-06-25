package ethrpc

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"sync/atomic"
	"time"

	"github.com/pkg/errors"
	"github.com/weirdgiraffe/jsonrpc"
)

type HTTP struct {
	seq     uint64
	baseURL string
	http    *http.Client
	dump    *Dumper
}

func NewHTTP(url string, opt ...DumpOption) *HTTP {
	return &HTTP{
		baseURL: url,
		http:    &http.Client{Timeout: 30 * time.Second},
		dump:    NewDumper(opt...),
	}
}

func (c *HTTP) Call(ctx context.Context, method string, params ...interface{}) (*jsonrpc.Response, error) {
	jreq, err := json.Marshal(jsonrpc.NewRequest(
		atomic.AddUint64(&c.seq, 1),
		method,
		params...,
	))
	if err != nil {
		return nil, errors.Wrap(err, "failed to encode request")
	}

	req, err := http.NewRequest(http.MethodPost, c.baseURL, bytes.NewReader(jreq))
	if err != nil {
		return nil, err
	}
	req.Header.Set("content-type", "application/json")

	c.dump.DumpRequest(jreq)
	res, err := c.http.Do(req.WithContext(ctx))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	c.dump.DumpResponse(body)

	jres, err := jsonrpc.DecodeFrom(bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	return jres.(*jsonrpc.Response), nil
}
