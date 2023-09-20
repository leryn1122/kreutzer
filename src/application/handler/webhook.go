package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/leryn1122/kreutzer/v2/application/service"
	"github.com/leryn1122/kreutzer/v2/application/service/adapter"
	"github.com/leryn1122/kreutzer/v2/infra/web"
)

// WebhookController `WebhookController` is something different from other controllers. We don't
// explicitly declare the request body with struct.
type WebhookController struct {
	WebhookService service.WebhookService
}

func NewWebhookHandler() WebhookController {
	return WebhookController{
		WebhookService: adapter.NewWebhookService(),
	}
}

type WebhookHandler = WebhookController

func (c WebhookController) ReceiveWebhook(ctx *gin.Context) {
	err := c.WebhookService.Receive(ctx.Request)
	if err != nil {
		web.OnError(ctx, err)
		return
	}
	web.OnSuccess(ctx, nil)
}
