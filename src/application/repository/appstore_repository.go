package repository

import (
	"context"
	"github.com/leryn1122/kreutzer/v2/application/entity"
)

type AppStoreRepository interface {
	Get(ctx context.Context, entity *entity.AppStore) (*entity.AppStore, error)
	List(ctx context.Context, entity *entity.AppStore) ([]*entity.AppStore, error)
	Save(ctx context.Context, entity *entity.AppStore) error
	Update(ctx context.Context, entity *entity.AppStore) error
	Remove(ctx context.Context, entity *entity.AppStore) error
}
