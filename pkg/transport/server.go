package transport

import (
	"context"
	"errors"
	"os"
	"os/signal"
	"syscall"

	"github.com/crochee/object/pkg/routine"
)

type AppServer interface {
	Start() error
	Stop() error
}

type app struct {
	option
	ctx    context.Context
	cancel context.CancelFunc
}

func NewApp(opts ...func(*option)) *app {
	a := &app{
		option: option{
			signals: []os.Signal{syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT},
			ctx:     context.Background(),
		},
	}
	for _, opt := range opts {
		opt(&a.option)
	}
	a.ctx, a.cancel = context.WithCancel(a.option.ctx)
	return a
}

func (a *app) Run() error {
	g := routine.NewGroup(a.ctx)
	for _, srv := range a.serverList {
		realSrv := srv
		g.Go(func(ctx context.Context) error {
			<-ctx.Done()
			return realSrv.Stop()
		})
		g.Go(func(ctx context.Context) error {
			return realSrv.Start()
		})
	}
	if a.shutdown != nil {
		g.Go(a.shutdown)
	} else {
		g.Go(func(ctx context.Context) error {
			quit := make(chan os.Signal, 1)
			signal.Notify(quit, a.option.signals...)
			for {
				select {
				case <-ctx.Done():
					close(quit)
					return ctx.Err()
				case <-quit:
					a.cancel()
				}
			}
		})
	}
	if err := g.Wait(); err != nil && !errors.Is(err, context.Canceled) {
		return err
	}
	return nil
}
