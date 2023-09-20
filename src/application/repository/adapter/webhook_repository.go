package adapter

import (
	"context"
	"github.com/leryn1122/kreutzer/v2/application/entity"
	"github.com/leryn1122/kreutzer/v2/application/model"
	"github.com/leryn1122/kreutzer/v2/application/repository"
	"github.com/leryn1122/kreutzer/v2/infra/db"
)

var webhookRepository *WebhookRepositoryImpl

type WebhookRepositoryImpl struct {
	client *db.DBClient
}

func NewWebhookRepository() repository.WebhookRepository {
	if webhookRepository != nil {
		return webhookRepository
	}
	webhookRepository = &WebhookRepositoryImpl{
		client: &db.Client,
	}
	return webhookRepository
}

func toModelByWebhook(e *entity.Webhook) *model.Webhook {
	return &model.Webhook{
		BaseModel: model.BaseModel{
			ID: e.ID.ID,
		},
		HookID:  e.HookID,
		Channel: e.Channel,
		Enabled: e.Enabled,
		URL:     e.URL,
	}
}

func toEntityByWebhook(m *model.Webhook, e *entity.Webhook) *entity.Webhook {
	e.Secret = m.Secret
	e.URL = m.URL
	return e
}

func (r WebhookRepositoryImpl) Get(ctx context.Context, entity *entity.Webhook) (*entity.Webhook, error) {
	m := toModelByWebhook(entity)
	tx := r.client.Find(m)
	return toEntityByWebhook(m, entity), tx.Error
}

func (r WebhookRepositoryImpl) List(ctx context.Context, entity *entity.Webhook) ([]*entity.Webhook, error) {
	//TODO implement me
	panic("implement me")
}

func (r WebhookRepositoryImpl) Save(ctx context.Context, entity *entity.Webhook) error {
	//TODO implement me
	panic("implement me")
}

func (r WebhookRepositoryImpl) Update(ctx context.Context, entity *entity.Webhook) error {
	//TODO implement me
	panic("implement me")
}

func (r WebhookRepositoryImpl) Remove(ctx context.Context, entity *entity.Webhook) error {
	//TODO implement me
	panic("implement me")
}
