package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
)

const (
	CollectionType = "collection"
	ObjectType     = "object"
)

type Result struct {
	Code       int         `json:"code" default:"0"`
	Message    string      `json:"message,omitempty"`
	Data       interface{} `json:"data,omitempty"`
	Type       string      `json:"type,omitempty"`
	Pagination *Pagination `json:"pagination,omitempty"`
}

type Option = func(*gin.Context, *Result)

type Pagination struct {
	Page int  `json:"page"`
	Size int  `json:"size"`
	Desc bool `json:"desc"`
}

func OnSuccess(ctx *gin.Context, data interface{}, opts ...Option) {
	result := &Result{
		Code: 0,
		Data: data,
	}
	if data == nil {
		result.Type = ObjectType
	} else if reflect.TypeOf(data).Kind() == reflect.Array {
		result.Type = CollectionType
	} else {
		result.Type = ObjectType
	}

	for _, opt := range opts {
		if opt != nil {
			opt(ctx, result)
		}
	}
	ctx.JSON(http.StatusOK, result)
}

func OnSuccessMessage(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusOK, &Result{
		Message: message,
	})
}

func OnError(ctx *gin.Context, err error) {
	result := &Result{
		Code:    -http.StatusInternalServerError,
		Message: err.Error(),
	}
	ctx.JSON(http.StatusOK, result)
}

func OnErrorMessage(ctx *gin.Context, message string) {
	result := &Result{
		Code:    -1,
		Message: message,
	}
	ctx.JSON(http.StatusOK, result)
}

func WithPagination(pagination Pagination) Option {
	return func(_ *gin.Context, result *Result) {
		result.Pagination = &pagination
	}
}
