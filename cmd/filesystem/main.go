// Copyright 2021, The Go Authors. All rights reserved.
// Author: crochee
// Date: 2021/5/16

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/crochee/file/pkg/transport"
	"github.com/crochee/file/pkg/transport/httpx"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // 全局取消
	// 初始化配置
	// 初始化系统日志

	httpSrv, err := httpx.New(ctx, ":9520", nil)
	if err != nil {
		_, _ = fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
	app := transport.NewApp(
		transport.Context(ctx),
		transport.Servers(httpSrv),
	)
	if err = app.Run(); err != nil {
		_, _ = fmt.Fprint(os.Stderr, err)
	}
	os.Exit(0)
}
