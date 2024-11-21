package cmd

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/util/grand"

	"websocket/internal/pkg/websocket"
)

var Main = gcmd.Command{
	Func: func(context.Context, *gcmd.Parser) error {
		go func() {
			for {
				<-time.After(1000 * time.Millisecond)
				data := make([]map[string]int, 0, 10)
				for i := 0; i < 10; i++ {
					data = append(data, map[string]int{"x": i, "y": grand.Intn(100)})
				}
				websocket.Notice(websocket.Message("data", data))

				websocket.Notice(websocket.Message("result", map[string]int{
					"result": grand.N(-1, 1),
				}))
			}
		}()

		server := g.Server()
		server.Use(ghttp.MiddlewareCORS)
		server.Run()

		return nil
	},
}
