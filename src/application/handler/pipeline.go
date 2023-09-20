package handler

import (
  "github.com/gin-gonic/gin"
  "github.com/leryn1122/kreutzer/v2/application/controller"
)

type PipelineHandler struct {
  controller controller.PipelineController
}

func NewPipelineHandler() PipelineHandler {
  return PipelineHandler{
    controller: controller.NewPipelineHandler(),
  }
}

func (c PipelineHandler) ListPipelines(ctx *gin.Context) {
}
