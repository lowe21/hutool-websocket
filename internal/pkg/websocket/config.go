package websocket

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

const (
	patternDebug            = "websocket.debug"            // 配置是否开启调试
	patternMessageMaxSize   = "websocket.messageMaxSize"   // 配置消息最大大小
	patternPingIntervalTime = "websocket.pingIntervalTime" // 配置ping间隔时间
	patternPongWaitTime     = "websocket.pongWaitTime"     // 配置pong等待时间
	patternWriteWaitTime    = "websocket.writeWaitTime"    // 配置写消息等待时间
	defaultDebug            = false                        // 默认是否开启调试
	defaultMessageMaxSize   = 512                          // 默认消息最大大小
	defaultPingIntervalTime = "60s"                        // 默认ping间隔时间
	defaultPongWaitTime     = "90s"                        // 默认pong等待时间
	defaultWriteWaitTime    = "10s"                        // 默认写消息等待时间
)

type Config struct {
	Debug            bool          // 是否开启调试
	MessageMaxSize   int64         // 消息最大大小
	PingIntervalTime time.Duration // ping间隔时间
	PongWaitTime     time.Duration // pong等待时间
	WriteWaitTime    time.Duration // 写消息等待时间
}

// Init 初始化配置
func (config *Config) Init() {
	ctx := context.TODO()

	if config.Debug = g.Config().MustGet(ctx, patternDebug).Bool(); !config.Debug {
		config.Debug = defaultDebug
	}
	if config.MessageMaxSize = g.Config().MustGet(ctx, patternMessageMaxSize).Int64(); config.MessageMaxSize <= 0 {
		config.MessageMaxSize = defaultMessageMaxSize
	}
	if config.PingIntervalTime = g.Config().MustGet(ctx, patternPingIntervalTime).Duration(); config.PingIntervalTime <= 0 {
		config.PingIntervalTime = gconv.Duration(defaultPingIntervalTime)
	}
	if config.PongWaitTime = g.Config().MustGet(ctx, patternPongWaitTime).Duration(); config.PongWaitTime <= 0 {
		config.PongWaitTime = gconv.Duration(defaultPongWaitTime)
	}
	if config.WriteWaitTime = g.Config().MustGet(ctx, patternWriteWaitTime).Duration(); config.WriteWaitTime <= 0 {
		config.WriteWaitTime = gconv.Duration(defaultWriteWaitTime)
	}
}
