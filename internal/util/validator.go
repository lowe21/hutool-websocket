package util

import (
	"context"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/gutil"
)

func Validator(ctx context.Context, pointer any, args ...any) (err error) {
	data := gmap.New()

	if len(args) > 0 {
		for _, arg := range args {
			for key, value := range gconv.Map(arg) {
				if !gutil.IsEmpty(value) {
					data.Set(key, value)
				}
			}
		}
	} else {
		if request := ghttp.RequestFromCtx(ctx); request != nil {
			for key, values := range request.Header {
				if value := garray.NewStrArrayFrom(values).FilterEmpty().Join(","); value != "" {
					data.Set(key, value)
				}
			}
			for key, value := range request.GetRequestMap() {
				if !gutil.IsEmpty(value) {
					data.Set(key, value)
				}
			}
		} else {
			for key, value := range gconv.Map(pointer) {
				data.Set(key, value)
			}
		}
	}

	if err = g.Validator().Bail().Data(pointer).Assoc(data).Run(ctx); err != nil {
		return
	}

	return gconv.Scan(data, pointer)
}
