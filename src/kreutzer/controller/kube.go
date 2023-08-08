package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/leryn1122/kreutzer/v2/kreutzer/service"
	"github.com/leryn1122/kreutzer/v2/lib/web"
)

// KubeController Call Kubernetes APIServer to handle the Kubernetes native API.
type KubeController struct{}

func (c KubeController) ListNamespaces(ctx *gin.Context) {
	clusterID := ctx.Param("cluster")
	namespaces, err := service.KubeService{}.ListNamespaces(clusterID)
	if err != nil {
		web.OnError(ctx, err)
		return
	}
	web.OnSuccess(ctx, namespaces)
}
