package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/leryn1122/kreutzer/v2/lib/web"
	"net/http"
	"reflect"
)

type GenericController struct{}

type AnyFunc func(...any) (any, error)

var (
	GinContextType = reflect.TypeOf(&gin.Context{})
	ErrorType      = reflect.TypeOf((*error)(nil)).Elem()
)

func GinWrap(handler AnyFunc) gin.HandlerFunc {
	if reflect.TypeOf(handler).Kind() != reflect.Func {
		panic("not a function")
	}

	funcType := reflect.ValueOf(handler).Type()
	if funcType.NumIn() != 2 {
		panic("not a convertable API endpoint")
	}

	requestType := funcType.In(0)
	if funcType.In(1) != GinContextType {
		panic("second parameters is not `*gin.Context`")
	}
	if funcType.Out(1).AssignableTo(ErrorType) {
		panic("second return value is not `error`")
	}

	return func(ctx *gin.Context) {
		method := ctx.Request.Method
		if method == http.MethodGet {
		} else {
			request := reflect.New(requestType)
			data, err := ctx.GetRawData()
			if err != nil {
				web.OnError(ctx, err)
				return
			}
			err = json.Unmarshal(data, &request)
			if err != nil {
				web.OnError(ctx, err)
				return
			}
			resp, err := handler(&request, ctx)
			if err != nil {
				web.OnError(ctx, err)
				return
			}
			web.OnSuccess(ctx, resp)
		}
	}
}
