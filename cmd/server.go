// Package cmd
package cmd

import (
	"context"

	"github.com/crochee/object/pkg/transport"
	"github.com/crochee/object/pkg/transport/httpx"
)

func Server() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // 全局取消
	// 初始化配置
	// 初始化系统日志

	httpSrv, err := httpx.New(ctx, ":9520", nil)
	if err != nil {
		return err
	}
	app := transport.NewApp(
		transport.Context(ctx),
		transport.Servers(httpSrv),
	)
	return app.Run()
}
