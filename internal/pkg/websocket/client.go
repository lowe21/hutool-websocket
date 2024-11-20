package websocket

import (
	"context"
	"sync"
	"time"

	"github.com/gorilla/websocket"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/grpool"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"

	"websocket/internal/util"
)

func newClient(config *Config, conn *websocket.Conn, clientId string) (client *Client, err error) {
	// 初始化配置
	if config == nil {
		config = &Config{}
	}
	config.Init()

	defer func() {
		if err == nil {
			// 监听读消息
			go client.reader()
			// 监听写消息
			go client.writer()
			// 注册客户端
			manager.Register(client)
		}
	}()

	return &Client{
		config:   config,
		conn:     conn,
		clientId: clientId,
		message:  make(chan []byte),
	}, nil
}

type Client struct {
	config   *Config         // 配置
	conn     *websocket.Conn // 连接
	clientId string          // 客户端id
	message  chan []byte     // 消息
	once     sync.Once       // once
}

// GetConfig 获取配置
func (client *Client) GetConfig() *Config {
	return client.config
}

// GetClientId 获取客户端id
func (client *Client) GetClientId() string {
	return client.clientId
}

// Send 发送消息
func (client *Client) Send(message any) {
	if g.Try(context.TODO(), func(context.Context) {
		client.message <- gconv.Bytes(message)
	}) != nil {
		client.Close()
	}
}

// Close 关闭
func (client *Client) Close() {
	client.once.Do(func() {
		// 开启调试
		if client.config.Debug {
			g.Log().Info(context.TODO(), gstr.Join([]string{"output", "<-", "closed"}, " "))
		}

		// 注销客户端
		manager.Unregister(client)

		// 关闭消息通道
		close(client.message)

		// 关闭连接
		_ = client.conn.Close()
	})
}

// reader 监听读消息
func (client *Client) reader() {
	defer func() {
		client.Close()
	}()

	// 设置读消息最大大小
	client.conn.SetReadLimit(client.config.MessageMaxSize)

	// 设置读消息期限
	if client.conn.SetReadDeadline(time.Now().Add(client.config.PongWaitTime)) != nil {
		return
	}

	// 设置pong处理方法
	client.conn.SetPongHandler(func(string) error {
		// 开启调试
		if client.config.Debug {
			g.Log().Info(context.TODO(), gstr.Join([]string{"input", "->", "pong"}, " "))
		}

		// 设置读消息期限
		return client.conn.SetReadDeadline(time.Now().Add(client.config.PongWaitTime))
	})

	for {
		// 读取消息
		_, message, err := client.conn.ReadMessage()
		if err != nil {
			return
		}

		// 协程池
		if grpool.AddWithRecover(context.TODO(), func(ctx context.Context) {
			// 开启调试
			if client.config.Debug {
				g.Log().Info(ctx, gstr.Join([]string{"input", "->", gstr.TrimAll(string(message))}, " "))
			}

			// 解析数据
			input := &Input{}
			if err = gconv.Scan(message, input); err != nil {
				client.Send(Message("error", gcode.CodeInvalidParameter, "message should be json object format"))
				return
			}

			// 处理方法
			if err = client.handler(ctx, input); err != nil {
				// 开启调试
				if client.config.Debug {
					g.Log().Error(ctx, err)
				}
			}
		}, func(ctx context.Context, exception error) {
			g.Log().Error(ctx, exception)
		}) != nil {
			return
		}
	}
}

// writer 监听写消息
func (client *Client) writer() {
	// 定时器
	ticker := time.NewTicker(client.config.PingIntervalTime)
	defer func() {
		ticker.Stop()
		client.Close()
	}()

	for {
		select {
		case <-ticker.C:
			// 开启调试
			if client.config.Debug {
				g.Log().Info(context.TODO(), gstr.Join([]string{"output", "<-", "ping"}, " "))
			}

			// 设置写消息期限
			if client.conn.SetWriteDeadline(time.Now().Add(client.config.WriteWaitTime)) != nil {
				return
			}

			// 写入ping帧
			if client.conn.WriteMessage(websocket.PingMessage, []byte{}) != nil {
				return
			}
		case message, ok := <-client.message:
			// 消息通道是否关闭
			if !ok {
				_ = client.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			// 开启调试
			if client.config.Debug {
				g.Log().Info(context.TODO(), gstr.Join([]string{"output", "<-", string(message)}, " "))
			}

			// 设置写消息期限
			if client.conn.SetWriteDeadline(time.Now().Add(client.config.WriteWaitTime)) != nil {
				return
			}

			// 写入消息
			if client.conn.WriteMessage(websocket.TextMessage, message) != nil {
				return
			}
		}
	}
}

// handler 处理方法
func (client *Client) handler(ctx context.Context, input *Input) (err error) {
	defer func() {
		if err != nil {
			client.Send(Message(input.Handler, err))
		}
	}()

	if err = util.Validator(ctx, input); err != nil {
		return
	}

	handler, ok := handlers[input.Handler]
	if !ok {
		return util.Error(gcode.CodeInvalidRequest, "handler not fount")
	}

	return handler(ctx, client, input)
}
