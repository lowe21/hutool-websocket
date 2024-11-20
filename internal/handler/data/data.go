package data

import (
	"context"

	"websocket/internal/pkg/websocket"
)

func init() {
	websocket.SetHandler("data", dataHandler)
}

func dataHandler(ctx context.Context, client *websocket.Client, input *websocket.Input) error {
	client.Send(websocket.Message(input.Handler, input.Params))

	return nil
}
