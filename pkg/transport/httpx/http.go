package httpx

import (
	"context"
	"crypto/tls"
	"net"
	"net/http"

	"github.com/crochee/object/pkg/logger"
)

type option struct {
	tlsConfig  *tls.Config
	requestLog logger.Builder
}

// TlsConfig
func TlsConfig(cfg *tls.Config) func(*option) {
	return func(o *option) { o.tlsConfig = cfg }
}

// RequestLog
func RequestLog(log logger.Builder) func(*option) {
	return func(o *option) { o.requestLog = log }
}

type httpServer struct {
	*http.Server
	net.Listener
	ctx context.Context
	option
}

// New new http AppServer
func New(ctx context.Context, host string, handler http.Handler, opts ...func(*option)) (*httpServer, error) {
	srv := &httpServer{
		Server: &http.Server{
			Handler: handler,
			BaseContext: func(_ net.Listener) context.Context {
				return ctx
			},
		},
		ctx: ctx,
	}
	for _, opt := range opts {
		opt(&srv.option)
	}
	ln, err := net.Listen("tcp", host)
	if err != nil {
		return nil, err
	}
	if srv.tlsConfig != nil {
		ln = tls.NewListener(ln, srv.tlsConfig)
	}
	srv.Listener = ln
	if srv.requestLog != nil {
		srv.ConnContext = func(ctx context.Context, c net.Conn) context.Context {
			return logger.Context(ctx, srv.requestLog)
		}
	}
	return srv, nil
}

func (h *httpServer) Start() error {
	return h.Serve(h.Listener)
}

func (h *httpServer) Stop() error {
	return h.Shutdown(h.ctx)
}
