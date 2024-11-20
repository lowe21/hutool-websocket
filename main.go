package main

import (
	"github.com/gogf/gf/v2/os/gctx"

	"websocket/internal/cmd"
	_ "websocket/internal/imports"
	_ "websocket/internal/logic"
)

func main() {
	cmd.Main.Run(gctx.New())
}
