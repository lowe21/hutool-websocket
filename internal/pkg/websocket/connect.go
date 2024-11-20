package websocket

import (
	"net/http"

	"github.com/gorilla/websocket"

	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
)

var ws = websocket.Upgrader{
	CheckOrigin: func(*http.Request) bool {
		return true
	},
}

// Connect 连接
func Connect(request *ghttp.Request) {
	conn, err := ws.Upgrade(request.Response.Writer, request.Request, nil)
	if err != nil {
		return
	}
	defer func() {
		if err != nil {
			_ = conn.WriteMessage(websocket.TextMessage, Message("connect", err))
			_ = conn.Close()
		}
	}()

	client, err := newClient(&Config{}, conn, gctx.CtxId(request.GetCtx()))
	if err != nil {
		return
	}
	client.Send(Message("connect", "connect succeed"))
}
