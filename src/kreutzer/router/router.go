package router

import (
	"github.com/gin-gonic/gin"
	api "github.com/leryn1122/kreutzer/v2/kreutzer/controller"
)

func RegisterRoutes(route *gin.Engine) {
	registerRoutesWithHelm(route)
	registerRoutesWithDeploy(route)
	registerRoutesWithWebhook(route)
	registerRoutesWithKube(route)
}

func registerRoutesWithHelm(route *gin.Engine) {
	subRoute := route.Group("/helm")

	envs := subRoute.Group("/env")
	{
		// helm env
		envs.GET("", api.HelmController{}.GetEnvs)
	}

	repos := subRoute.Group("/repo")
	{
		// helm repo list
		repos.GET("/list", api.HelmController{}.ListRepos)
		// helm search repo
		repos.GET("/:repo/charts", api.HelmController{}.ListCharts)
		// helm repo update
		repos.PUT("/:repo/update", api.HelmController{}.UpdateRepo)
	}

	charts := subRoute.Group("/charts")
	{
		// helm show
		charts.GET("", api.HelmController{}.ShowChartInfo)
		//
		//charts.GET("")
		//
		//charts.GET("")
	}

	releases := subRoute.Group("/namespaces/:namespace/releases")
	{
		// helm list
		releases.GET("/list", api.HelmController{}.ListReleases)
		// helm show
		releases.GET("/:release", api.HelmController{}.ShowReleaseInfo)
		// helm install
		releases.POST("/:release", dummy)
		// helm upgrade
		releases.PUT("/:release", dummy)
		// helm uninstall
		releases.DELETE("/:release", dummy)
		// helm rollback
		releases.PUT("/:release/versions/:revision", dummy)
		// helm status
		releases.GET("/:release/status", api.HelmController{}.GetReleaseStatus)
		// helm release history
		releases.GET("/:release/histories", api.HelmController{}.ListReleaseHistories)
	}
}

func registerRoutesWithDeploy(route *gin.Engine) {
	subRoute := route.Group("/deploy")
	subRoute.GET("/env/list", api.DeployController{}.ListEnvInfo)
	subRoute.POST("/env/create", api.DeployController{}.CreateEnvInfo)
}

func registerRoutesWithWebhook(route *gin.Engine) {
	route.POST("/webhooks", api.WebhookController{}.ReceiveWebhook)
}

func registerRoutesWithKube(route *gin.Engine) {
	subRoute := route.Group("/kube")
	subRoute.GET("/namespaces/:cluster", api.KubeController{}.ListNamespaces)
}

func dummy(ctx *gin.Context) {
	_, _ = ctx.Writer.WriteString("TODO")
}
