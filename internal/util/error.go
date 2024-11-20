package util

import (
	"errors"

	"github.com/apache/dubbo-go-hessian2/java_exception"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/text/gstr"
)

type ErrorCode struct {
	code    int
	subCode string
	message string
	detail  any
}

func (errorCode ErrorCode) Code() int {
	return errorCode.code
}

func (errorCode ErrorCode) SubCode() string {
	return errorCode.subCode
}

func (errorCode ErrorCode) Message() string {
	return errorCode.message
}

func (errorCode ErrorCode) Detail() any {
	return errorCode.detail
}

func Error(args ...any) error {
	var (
		code    int
		subCode string
		message string
	)

	if argsLen := len(args); argsLen > 0 {
		switch value := args[0].(type) {
		case error:
			var throwable *java_exception.Throwable
			if errors.As(value, &throwable) {
				desc := gstr.StrEx(throwable.Error(), "desc = ")
				code = 999
				subCode = gstr.StrTillEx(desc, "@")
				message = gstr.StrEx(desc, "@")
			} else {
				switch val := gerror.Code(value).(type) {
				case ErrorCode:
					code = val.Code()
					subCode = val.SubCode()
				case gcode.Code:
					if val == gcode.CodeNil {
						val = gcode.CodeInternalError
					}
					code = val.Code()
					subCode = val.Message()
				}
				message = value.Error()
			}
		case gcode.Code:
			if value == gcode.CodeNil {
				value = gcode.CodeInternalError
			}
			code = value.Code()
			subCode = value.Message()
			message = value.Message()
		case string:
			if argsLen > 1 {
				code = gcode.CodeNil.Code()
				subCode = value
			} else {
				code = gcode.CodeInternalError.Code()
				subCode = gcode.CodeInternalError.Message()
				message = value
			}
		}
		if argsLen > 1 {
			if value, ok := args[1].(string); ok {
				message = value
			}
		}
	}

	if subCode == "" {
		code = gcode.CodeUnknown.Code()
		subCode = gcode.CodeUnknown.Message()
		message = gcode.CodeUnknown.Message()
	}

	return gerror.NewCode(ErrorCode{
		code:    code,
		subCode: gstr.CaseSnakeScreaming(subCode),
		message: message,
	}, message)
}
