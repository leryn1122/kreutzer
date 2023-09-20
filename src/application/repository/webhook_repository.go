package repository

import (
	"context"
	"github.com/leryn1122/kreutzer/v2/application/entity"
)

type WebhookRepository interface {
	Get(ctx context.Context, entity *entity.Webhook) (*entity.Webhook, error)
	List(ctx context.Context, entity *entity.Webhook) ([]*entity.Webhook, error)
	Save(ctx context.Context, entity *entity.Webhook) error
	Update(ctx context.Context, entity *entity.Webhook) error
	Remove(ctx context.Context, entity *entity.Webhook) error
}
