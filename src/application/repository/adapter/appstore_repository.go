package adapter

import (
	"context"
	"github.com/leryn1122/kreutzer/v2/application/entity"
	"github.com/leryn1122/kreutzer/v2/application/model"
	"github.com/leryn1122/kreutzer/v2/application/repository"
	"github.com/leryn1122/kreutzer/v2/infra/db"
)

type AppStoreRepository interface {
	Get(ctx context.Context, entity *entity.AppStore) (*entity.AppStore, error)
	List(ctx context.Context, entity *entity.AppStore) ([]*entity.AppStore, error)
	Save(ctx context.Context, entity *entity.AppStore) error
	Update(ctx context.Context, entity *entity.AppStore) error
	Remove(ctx context.Context, entity *entity.AppStore) error
}

var appStoreRepository AppStoreRepository

type AppStoreRepositoryImpl struct {
	client *db.DBClient
}

func toModelByAppStore(e *entity.AppStore) *model.AppStore {
	return &model.AppStore{
		BaseModel: model.BaseModel{},
		Name:      e.Name,
		URL:       e.URL,
		LastSync:  e.LastSync,
	}
}

func toEntityByAppStore(m *model.AppStore) *entity.AppStore {
	return &entity.AppStore{
		Name:     m.Name,
		URL:      m.URL,
		LastSync: m.LastSync,
	}
}

func (r AppStoreRepositoryImpl) Get(ctx context.Context, e *entity.AppStore) (*entity.AppStore, error) {
	var mo *model.AppStore
	tx := r.client.Model(&model.AppStore{}).
		Where(toModelByAppStore(e)).
		Find(&mo)
	return toEntityByAppStore(mo), tx.Error
}

func (r AppStoreRepositoryImpl) List(ctx context.Context, e *entity.AppStore) ([]*entity.AppStore, error) {
	var mos []model.AppStore
	tx := r.client.Model(&model.AppStore{}).
		Where(toModelByAppStore(e)).
		Order("name").
		Find(&mos)
	if tx.Error != nil {
		return nil, tx.Error
	}
	es := make([]*entity.AppStore, len(mos))
	for i, mo := range mos {
		es[i] = toEntityByAppStore(&mo)
	}
	return es, tx.Error
}

func (r AppStoreRepositoryImpl) Save(ctx context.Context, e *entity.AppStore) error {
	//TODO implement me
	panic("implement me")
}

func (r AppStoreRepositoryImpl) Update(ctx context.Context, e *entity.AppStore) error {
	tx := r.client.Model(&model.AppStore{}).
		Where("name = ?", toModelByAppStore(e).Name).
		Updates(toModelByAppStore(e))
	return tx.Error
}

func (r AppStoreRepositoryImpl) Remove(ctx context.Context, e *entity.AppStore) error {
	tx := r.client.Model(&model.AppStore{}).
		Delete("name = ?", toModelByAppStore(e).Name)
	return tx.Error
}

func NewAppStoreRepository() repository.AppStoreRepository {
	if appStoreRepository != nil {
		return appStoreRepository
	}
	appStoreRepository = &AppStoreRepositoryImpl{
		client: &db.Client,
	}
	return appStoreRepository
}
