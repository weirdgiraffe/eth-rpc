package ethrpc

import (
	"fmt"
	"io"
	"sync"
)

type Dumper struct {
	dumpRequests      bool
	dumpResponses     bool
	dumpNotifications bool

	mx     sync.Mutex
	output io.Writer
}

type DumpOption func(*Dumper)

func DumpRequests() DumpOption {
	return func(d *Dumper) {
		d.dumpRequests = true
	}
}

func DumpResponses() DumpOption {
	return func(d *Dumper) {
		d.dumpResponses = true
	}
}

func DumpNotifications() DumpOption {
	return func(d *Dumper) {
		d.dumpNotifications = true
	}
}

func DumpMessages() DumpOption {
	return func(d *Dumper) {
		d.dumpRequests = true
		d.dumpResponses = true
	}
}

func DumpOutput(w io.Writer) DumpOption {
	return func(d *Dumper) {
		d.output = w
	}
}

func NewDumper(opt ...DumpOption) *Dumper {
	d := &Dumper{output: io.Discard}
	for i := range opt {
		opt[i](d)
	}
	return d
}

func (d *Dumper) DumpRequest(req []byte) {
	if d.dumpRequests {
		d.mx.Lock()
		defer d.mx.Unlock()
		fmt.Fprintln(d.output, string(req))
	}
}

func (d *Dumper) DumpResponse(res []byte) {
	if d.dumpResponses {
		d.mx.Lock()
		defer d.mx.Unlock()
		fmt.Fprintln(d.output, string(res))
	}
}

func (d *Dumper) DumpNotification(notify []byte) {
	if d.dumpNotifications {
		d.mx.Lock()
		defer d.mx.Unlock()
		fmt.Fprintln(d.output, string(notify))
	}
}
