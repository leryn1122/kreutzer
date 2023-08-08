package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-errors/errors"
	"github.com/leryn1122/kreutzer/v2/kreutzer/service"
	"github.com/leryn1122/kreutzer/v2/lib/web"
	"regexp"
)

type WebhookController struct{}

const (
	UserAgent = "User-Agent"
)

func (c WebhookController) ReceiveWebhook(ctx *gin.Context) {
	userAgent := ctx.GetHeader(UserAgent)
	if userAgent == "" {
		web.OnError(ctx, errors.Errorf("unknown webhook origin"))
		return
	}

	// GitHub webhook
	// To determine `User-Agent: GitHub-Hookshot/0000`
	// Verify the signature is still required.
	matched, err := regexp.MatchString("GitHub-Hookshot/[0-9]+", userAgent)
	if err == nil && matched {
		msg, err := service.GithubWebhookService{}.HandlerGithubWebhook(ctx.Request)
		if err != nil {
			web.OnError(ctx, err)
			return
		}
		web.OnSuccess(ctx, *msg)
		return
	}

	web.OnError(ctx, errors.Errorf("unknown webhook origin"))
}
