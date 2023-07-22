package router

import (
	"github.com/gin-gonic/gin"
	api "github.com/leryn1122/kreutzer/v2/kreutzer/controller"
)

func RegisterRoutes(route *gin.Engine) {

	envs := route.Group("/envs")
	{
		// helm env
		envs.GET("", api.HelmController{}.GetEnvs)
	}

	repos := route.Group("/repo")
	{
		// helm repo list
		repos.GET("/list", api.HelmController{}.ListRepo)
		// helm search repo
		repos.GET("/:repo/charts", api.HelmController{}.ListCharts)
		// helm repo update
		repos.PUT("/:repo/update", api.HelmController{}.UpdateRepo)
	}

	charts := route.Group("/charts")
	{
		// helm show
		charts.GET("", api.HelmController{}.ShowChartInfo)
		////
		//charts.GET("")
		////
		//charts.GET("")
	}

	releases := route.Group("/namespaces/:namespace/releases")
	{
		// helm list
		releases.GET("/list", api.HelmController{}.ListReleases)
		// helm show
		releases.GET("/:release", api.HelmController{}.ShowReleaseInfo)
		// helm install
		releases.POST("/:release")
		// helm upgrade
		releases.PUT("/:release")
		// helm uninstall
		releases.DELETE("/:release")
		// helm rollback
		releases.PUT("/:release/versions/:revision")
		// helm status
		releases.GET("/:release/status", api.HelmController{}.GetReleaseStatus)
		// helm release history
		releases.GET("/:release/histories", api.HelmController{}.ListReleaseHistories)
	}

}

func dummy(ctx *gin.Context) {
	_, _ = ctx.Writer.WriteString("TODO")
}
