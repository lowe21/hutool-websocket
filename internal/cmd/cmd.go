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
	Func: func(context.Context, *gcmd.Parser) (err error) {
		go func() {
			for {
				<-time.After(time.Second)

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
		if err = server.Start(); err != nil {
			return
		}

		web := g.Server("web")
		web.SetPort(8188)
		web.SetServerRoot("/web")
		web.SetDumpRouterMap(false)
		if err = web.Start(); err != nil {
			return
		}

		g.Wait()

		return
	},
}
