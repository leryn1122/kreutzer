package controller

import (
	"github.com/gin-gonic/gin"
)

type GenericController struct{}

type Handler = func(*gin.Context, any) (any, error)

func (c GenericController) handle(Handler) {
}
