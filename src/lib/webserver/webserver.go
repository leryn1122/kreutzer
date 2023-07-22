package webserver

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/leryn1122/kreutzer/v2/lib/web"
	"net/http"
)

func CreateRoute() *gin.Engine {
	return gin.Default()
}

func StartRestfulWebServer(router *gin.Engine) {
	gin.SetMode(gin.DebugMode)

	_ = router.SetTrustedProxies(nil)

	router.NoRoute(func(ctx *gin.Context) {
		path := ctx.Request.URL.Path
		method := ctx.Request.Method
		ctx.JSON(http.StatusOK, &web.Result{
			Code:    -http.StatusNotFound,
			Message: fmt.Sprintf("%s %s not found.", method, path),
		})
	})

	router.Use(Cors())
	router.GET("/healthz", Healthz)

	_ = router.Run(fmt.Sprintf("%s:%s", "0.0.0.0", "8080"))
}

func Healthz(ctx *gin.Context) {
	web.OnSuccessMessage(ctx, "Health check passed!")
}
