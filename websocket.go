package ethrpc

import (
	"bytes"
	"context"
	"encoding/json"
	"sync"
	"sync/atomic"

	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
	"github.com/weirdgiraffe/jsonrpc"
)

type Websocket struct {
	seq  uint64
	conn *websocket.Conn

	responses     chan *jsonrpc.Response
	notifications chan *jsonrpc.Notification

	done chan struct{}
	err  error

	wMx  sync.Mutex
	dump *Dumper
}

func DialWebsocket(ctx context.Context, url string, opt ...DumpOption) (*Websocket, error) {
	conn, _, err := websocket.DefaultDialer.DialContext(ctx, url, nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to dial websocket")
	}
	ws := &Websocket{
		conn:          conn,
		responses:     make(chan *jsonrpc.Response, 100),
		notifications: make(chan *jsonrpc.Notification, 100),
		done:          make(chan struct{}, 1),
		dump:          NewDumper(opt...),
	}
	go ws.run()

	return ws, nil
}

func (ws *Websocket) Close() error {
	select {
	case <-ws.done:
	default:
		_ = ws.conn.Close()
	}
	return nil
}

func (ws *Websocket) Done() <-chan struct{} {
	return ws.done
}

func (ws *Websocket) Err() error {
	return ws.err
}

func (ws *Websocket) Call(_ context.Context, method string, params ...interface{}) (*jsonrpc.Response, error) {
	return ws.call(method, params...)
}

func (ws *Websocket) call(method string, params ...interface{}) (*jsonrpc.Response, error) {
	select {
	case <-ws.done:
		return nil, ws.err
	default:
	}
	req := jsonrpc.NewRequest(
		atomic.AddUint64(&ws.seq, 1),
		method,
		params...)

	jreq, err := json.Marshal(req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to encode request")
	}

	ws.dump.DumpRequest(jreq)

	ws.wMx.Lock()
	err = ws.conn.WriteMessage(websocket.TextMessage, jreq)
	if err != nil {
		ws.wMx.Unlock()
		return nil, errors.Wrap(err, "failed to write request")
	}
	ws.wMx.Unlock()

	for {
		var res *jsonrpc.Response
		select {
		case <-ws.done:
			return nil, errors.New("connection closed")
		case res = <-ws.responses:
		}

		if res.ID == req.ID {
			return res, nil
		}

		ws.responses <- res
	}
}

func (ws *Websocket) Subscribe(subType string, params ...interface{}) <-chan []byte {
	out := make(chan []byte)

	res, err := ws.call("eth_subscribe", append([]interface{}{subType}, params...)...)
	if err != nil {
		close(out)
		return out
	}

	if res.Error != nil {
		close(out)
		return out
	}

	var subID string
	if err = json.Unmarshal(res.Result, &subID); err != nil {
		close(out)
		return out
	}

	go func() {
		defer close(out)

		for {
			var n *jsonrpc.Notification

			select {
			case <-ws.done:
				return
			case n = <-ws.notifications:
			}

			var subNotification struct {
				Subscription string          `json:"subscription"`
				Result       json.RawMessage `json:"result"`
			}

			err := json.Unmarshal(n.Params, &subNotification)
			if err != nil {
				return
			}

			if subNotification.Subscription == subID {
				out <- subNotification.Result
				continue
			}

			ws.notifications <- n
		}
	}()

	return out
}

func (ws *Websocket) run() {
	for {
		t, msg, err := ws.conn.ReadMessage()
		if err != nil {
			ws.err = errors.Wrap(err, "failed to read websocket message")
			close(ws.done)
			return
		}
		if t != websocket.TextMessage {
			continue
		}

		jmsg, err := jsonrpc.DecodeFrom(bytes.NewReader(msg))
		if err != nil {
			continue
		}

		switch v := jmsg.(type) {
		case *jsonrpc.Notification:
			ws.dump.DumpNotification(msg)
			ws.notifications <- v
		case *jsonrpc.Response:
			ws.dump.DumpResponse(msg)
			ws.responses <- v
		}
	}
}
