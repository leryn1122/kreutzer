package controller

import (
	"github.com/gin-gonic/gin"
)

type TestController interface {
	HealthCheck(request HealthCheckRequest, ctx *gin.Context) (HealthCheckResponse, error)
}

type testControllerImpl struct{}

type HealthCheckRequest struct {
}

type HealthCheckResponse struct {
}

func (c *testControllerImpl) HealthCheck(request *HealthCheckRequest, ctx *gin.Context) (*HealthCheckResponse, error) {
	return &HealthCheckResponse{}, nil
}
