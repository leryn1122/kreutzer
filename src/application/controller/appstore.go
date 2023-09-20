package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/leryn1122/kreutzer/v2/application/vo"
	"github.com/leryn1122/kreutzer/v2/infra/web"
	"time"
)

type AppstoreController interface {
	ListAppstores(ctx *gin.Context, request *ListAppstoresRequest) (*ListAppstoresResponse, error)
	RefreshAppStore(ctx *gin.Context, request *RefreshAppStoreRequest) (*RefreshAppStoreResponse, error)
}

type ListAppstoresRequest struct {
	Pagination web.Pagination `json:"pagination"`
}

type ListAppstoresResponse = []*vo.AppStore

type RefreshAppStoreRequest struct {
	Name string `json:"name"`
}

type RefreshAppStoreResponse struct {
	LastSync *time.Time `json:"lastSync"`
}
