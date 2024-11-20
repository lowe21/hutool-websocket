package websocket

import (
	"context"
	"fmt"
	"sync"
)

type Handler func(context.Context, *Client, *Input) error

var (
	handlers = map[string]Handler{}
	mutex    sync.Mutex
)

// SetHandler 设置处理程序
func SetHandler(name string, handler Handler) {
	mutex.Lock()
	defer mutex.Unlock()

	if name == "" {
		panic("websocket handler name should not be empty")
	}

	if _, ok := handlers[name]; ok {
		panic(fmt.Sprintf("duplicate websocket handler name %s", name))
	}

	handlers[name] = handler
}
