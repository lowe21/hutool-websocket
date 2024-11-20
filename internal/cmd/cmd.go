package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var Main = gcmd.Command{
	Func: func(context.Context, *gcmd.Parser) error {
		server := g.Server()
		server.Use(ghttp.MiddlewareCORS)
		server.Run()

		return nil
	},
}
