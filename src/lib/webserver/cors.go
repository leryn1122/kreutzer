package webserver

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

// CorsOptions CORS options.
type CorsOptions struct {
	AllowedOrigins   []string
	AllowedMethods   []string
	AllowedHeaders   []string
	ExposedHeaders   []string
	AllowCredentials bool
	MaxAge           int
}

func Cors() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Add("Access-Control-Allow-Credentials", strconv.FormatBool(true))
		ctx.Writer.Header().Add("Access-Control-Allow-Origin", strings.Join(
			[]string{"*"}, ","))
		ctx.Writer.Header().Add("Access-Control-Allow-Methods", strings.Join(
			[]string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions}, ","))
		ctx.Writer.Header().Add("Access-Control-Allow-Headers", strings.Join(
			[]string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"}, ","))
		ctx.Writer.Header().Add("Access-Control-Expose-Headers", strings.Join(
			[]string{"X-Access-Token"}, ","))
		ctx.Writer.Header().Add("Access-Control-Max-Age", strconv.Itoa(300))

		method := ctx.Request.Method
		if http.MethodOptions == method {
			ctx.AbortWithStatus(http.StatusNoContent)
		}

		ctx.Next()
	}
}
