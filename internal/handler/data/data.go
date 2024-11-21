package data

import (
	"context"

	"websocket/internal/pkg/websocket"
	"websocket/internal/util"
)

func init() {
	websocket.SetHandler("data", dataHandler)
}

func dataHandler(_ context.Context, client *websocket.Client, input *websocket.Input) (err error) {
	defer func() {
		if exception := recover(); exception != nil {
			err = util.Error(exception)
		}
	}()

	client.Send(websocket.Message(input.Handler, input.Params))

	return
}
