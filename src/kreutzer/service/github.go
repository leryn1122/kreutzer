package service

import (
	"github.com/leryn1122/kreutzer/v2/kreutzer/dao"
	"github.com/leryn1122/kreutzer/v2/lib/db"
	"github.com/leryn1122/kreutzer/v2/lib/webhook"
	"github.com/sirupsen/logrus"
	"net/http"
)

type GithubWebhookService struct{}

// HandlerGithubWebhook Handles the GitHub webhook from `http.Request`
func (s GithubWebhookService) HandlerGithubWebhook(request *http.Request) (*string, error) {
	hook, err := webhook.NewGithubWebhook(request)
	if err != nil {
		return nil, err
	}

	secret, err := s.fetchWebhookSecret(hook.HookId)
	if err != nil {
		return nil, err
	}

	hook, err = hook.Verify([]byte(secret))
	if err != nil {
		return nil, err
	}
	return &hook.HookId, err
}

// fetchWebhookSecret Fetch the webhook secret
// TODO Read the secret from vault in later version.
func (s GithubWebhookService) fetchWebhookSecret(hookId string) (string, error) {
	var webhookDao dao.Webhook
	tx := db.DBClient.
		Select("hook_id, secret").
		Where("hook_id = ?", hookId).
		Where("channel = ?", "github").
		Limit(5).
		Find(&webhookDao)
	if err := tx.Error; err != nil {
		logrus.Errorf("failed to fetch webhook details: %+v", err)
		return "", err
	}
	return webhookDao.WebhookId, nil
}
