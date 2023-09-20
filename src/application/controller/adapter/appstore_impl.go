package adapter

import (
	"github.com/gin-gonic/gin"
	"github.com/leryn1122/kreutzer/v2/application/controller"
	"github.com/leryn1122/kreutzer/v2/application/service"
	"github.com/leryn1122/kreutzer/v2/application/service/adapter"
)

type appstoreController struct {
	AppstoreService service.AppStoreService
}

func NewAppstoreController() controller.AppstoreController {
	return appstoreController{
		AppstoreService: adapter.NewAppstoreService(),
	}
}

func (c appstoreController) ListAppstores(ctx *gin.Context, request *controller.ListAppstoresRequest) (*controller.ListAppstoresResponse, error) {
	appstores, err := c.AppstoreService.ListAppstores()
	return &appstores, err
}

func (c appstoreController) RefreshAppStore(ctx *gin.Context, request *controller.RefreshAppStoreRequest) (*controller.RefreshAppStoreResponse, error) {
	lastSync, err := c.AppstoreService.RefreshAppStore(request.Name)
	return &controller.RefreshAppStoreResponse{
		LastSync: lastSync,
	}, err
}
