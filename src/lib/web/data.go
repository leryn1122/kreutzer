package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
	"time"
)

type DataType string

const (
	CollectionType DataType = "collection"
	ObjectType     DataType = "object"
	StringType     DataType = "string"
)

type Result struct {
	Code       int         `json:"code" default:"0"`
	Message    string      `json:"message,omitempty"`
	Data       interface{} `json:"data,omitempty"`
	Timestamp  time.Time   `json:"timestamp,omitempty"`
	Type       DataType    `json:"type,omitempty"`
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
		Code:      0,
		Data:      data,
		Timestamp: time.Now(),
		Type:      determineType(data),
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
		Message:   message,
		Timestamp: time.Now(),
	})
}

func OnError(ctx *gin.Context, err error) {
	result := &Result{
		Code:      -http.StatusInternalServerError,
		Message:   err.Error(),
		Timestamp: time.Now(),
	}
	ctx.JSON(http.StatusOK, result)
}

func OnErrorMessage(ctx *gin.Context, message string) {
	result := &Result{
		Code:      -1,
		Message:   message,
		Timestamp: time.Now(),
	}
	ctx.JSON(http.StatusOK, result)
}

func WithPagination(pagination Pagination) Option {
	return func(_ *gin.Context, result *Result) {
		result.Pagination = &pagination
	}
}

func determineType(data interface{}) DataType {
	var result DataType

	if data == nil {
		result = ObjectType
	}

	kind := reflect.TypeOf(data).Kind()
	if kind == reflect.Slice || kind == reflect.Array {
		result = CollectionType
	} else if kind == reflect.Pointer {
		elem := reflect.TypeOf(data).Elem().Kind()
		if elem == reflect.Slice || elem == reflect.Array {
			result = CollectionType
		}
	} else if kind == reflect.String {
		result = StringType
	} else {
		result = ObjectType
	}
	return result
}
