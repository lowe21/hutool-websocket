package websocket

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"

	"websocket/internal/util"
)

type Input struct {
	Handler string `json:"handler" valid:"required"` // 处理方法
	Params  any    `json:"params"  valid:"json"`     // 参数
}

type Output struct {
	Handler string `json:"handler"` // 处理方法
	Code    string `json:"code"`    // 返回码
	Message string `json:"message"` // 返回信息
	Data    any    `json:"data"`    // 返回数据
}

// Message 消息
func Message(handler string, args ...any) []byte {
	var (
		code    = gcode.CodeOK.Message()
		message string
		data    any
	)

	if argsLen := len(args); argsLen > 0 {
		switch value := args[0].(type) {
		case error:
			errCode := gerror.Code(util.Error(value)).(util.ErrorCode)
			code = errCode.SubCode()
			message = errCode.Message()
		case gcode.Code:
			if value == gcode.CodeNil {
				value = gcode.CodeInternalError
			}
			code = value.Message()
			if argsLen > 1 {
				message = gconv.String(args[1])
			} else {
				message = value.Message()
			}
		case string:
			if argsLen > 1 {
				code = value
				message = gconv.String(args[1])
			} else {
				message = value
			}
		default:
			switch output := value.(type) {
			case Output:
				if handler == "" {
					handler = output.Handler
				}
				code = output.Code
				message = output.Message
				data = output.Data
			case *Output:
				if handler == "" {
					handler = output.Handler
				}
				code = output.Code
				message = output.Message
				data = output.Data
			default:
				data = value
			}
		}
	}

	return gjson.MustEncode(&Output{
		Handler: handler,
		Code:    gstr.CaseSnakeScreaming(code),
		Message: message,
		Data:    data,
	})
}
