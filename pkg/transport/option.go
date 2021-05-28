// Copyright 2021, The Go Authors. All rights reserved.
// Author: crochee
// Date: 2021/5/16

package transport

import (
	"context"
	"os"
)

type option struct {
	signals    []os.Signal
	serverList []AppServer
	ctx        context.Context
	shutdown   func(context.Context) error
}

// Signal with exit signals.
func Signal(sigList ...os.Signal) func(*option) {
	return func(o *option) { o.signals = sigList }
}

// Servers with transport servers.
func Servers(servers ...AppServer) func(*option) {
	return func(o *option) { o.serverList = servers }
}

// Context with service context.
func Context(ctx context.Context) func(*option) {
	return func(o *option) { o.ctx = ctx }
}

// shutdown register app shutdown function
// you must promise ctx to cancel,otherwise goroutine deadlock
func shutdown(f func(ctx context.Context) error) func(*option) {
	return func(o *option) { o.shutdown = f }
}
