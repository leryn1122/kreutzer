package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/leryn1122/kreutzer/v2/application/controller"
	"github.com/leryn1122/kreutzer/v2/application/controller/adapter"
	"github.com/leryn1122/kreutzer/v2/infra/web"
)

type AppStoreHandler struct {
	Controller controller.AppstoreController
}

func NewAppstoreHandler() AppStoreHandler {
	return AppStoreHandler{
		Controller: adapter.NewAppstoreController(),
	}
}

func (c AppStoreHandler) ListAppStores(ctx *gin.Context) {
	data, err := c.Controller.ListAppstores(ctx, &controller.ListAppstoresRequest{})
	web.OnResponse(ctx, data, err)
}

func (c AppStoreHandler) RefreshAppStore(ctx *gin.Context) {
	var request controller.RefreshAppStoreRequest
	body, err := ctx.GetRawData()
	if err != nil {
		web.OnError(ctx, err)
		return
	}
	err = json.Unmarshal(body, &request)
	if err != nil {
		web.OnError(ctx, err)
		return
	}
	data, err := c.Controller.RefreshAppStore(ctx, &request)
	web.OnResponse(ctx, data, err)
}
