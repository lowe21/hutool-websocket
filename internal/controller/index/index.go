package index

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"

	"websocket/internal/pkg/websocket"
)

func init() {
	g.Server().BindObjectRest("/", &indexController{})
}

type indexController struct{}

func (*indexController) Get(request *ghttp.Request) {
	defer func() {
		if exception := recover(); exception != nil {
			g.Log().Error(request.GetCtx(), exception)
		}
	}()

	websocket.Connect(request)
}
