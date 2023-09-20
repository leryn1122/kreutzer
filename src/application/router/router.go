package router

import (
	"github.com/gin-gonic/gin"
	api "github.com/leryn1122/kreutzer/v2/application/handler"
)

func RegisterRoutes(route *gin.Engine) {
	registerRoutesWithAppstores(route)
	registerRoutesWithWebhook(route)
	registerRoutesWithPipeline(route)
}

func registerRoutesWithAppstores(route *gin.Engine) {
	handler := api.NewAppstoreHandler()
	subroute := route.Group("/appstore")

	subroute.GET("/list", handler.ListAppStores)
	subroute.POST("/refresh", handler.RefreshAppStore)
}

func registerRoutesWithWebhook(route *gin.Engine) {
	handler := api.NewWebhookHandler()
	route.POST("/webhooks", handler.ReceiveWebhook)
}

func registerRoutesWithPipeline(route *gin.Engine) {
	handler := api.NewPipelineHandler()
	subroute := route.Group("/pipeline")

	subroute.GET("/list", handler.ListPipelines)
}
