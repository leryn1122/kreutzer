package service

import (
	"github.com/leryn1122/kreutzer/v2/application/entity"
	"net/http"
)

type WebhookService interface {
	Receive(request *http.Request) error
	VerifyWebhook(webhook *entity.Webhook) error
}
