package web

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type invalidArgument struct {
	Field string `json:"field"`
	Value string `json:"value"`
	Tag   string `json:"tag"`
	Param string `json:"param"`
}

func BindData(ctx *gin.Context, request interface{}) bool {
	// Manually handle the data.
	if ctx.ContentType() != "application/json" {
		return true
	}

	if err := ctx.ShouldBind(request); err != nil {
		logrus.Printf("fail to binding data: %+v\n", err)
		if errs, ok := err.(validator.ValidationErrors); ok {
			var invalidArgs []invalidArgument

			for _, err := range errs {
				invalidArgs = append(invalidArgs, invalidArgument{
					err.Field(),
					err.Value().(string),
					err.Tag(),
					err.Param(),
				})
			}

			OnError(ctx, errors.Errorf("%+v", invalidArgs))
			return false
		}
		OnError(ctx, errors.Errorf("fail to binding data: %+v\n", err))
		return false
	}
	return true
}
