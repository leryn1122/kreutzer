package adapter

import (
	"context"
	"github.com/leryn1122/kreutzer/v2/application/entity"
	"github.com/leryn1122/kreutzer/v2/application/repository"
	"github.com/leryn1122/kreutzer/v2/application/repository/adapter"
	"github.com/leryn1122/kreutzer/v2/application/service"
	"github.com/leryn1122/kreutzer/v2/application/vo"
	"github.com/leryn1122/kreutzer/v2/infra/helm"
	"time"
)

var appstoreService *AppStoreServiceImpl

type AppStoreServiceImpl struct {
	AppStoreRepository repository.AppStoreRepository
	HelmClient         *helm.Client
}

func NewAppstoreService() service.AppStoreService {
	if appstoreService != nil {
		return appstoreService
	}
	appstoreService = &AppStoreServiceImpl{
		AppStoreRepository: adapter.NewAppStoreRepository(),
		HelmClient:         helm.NewHelmClient(),
	}
	return appstoreService
}

func (s *AppStoreServiceImpl) ListAppstores() ([]*vo.AppStore, error) {
	appstores, err := s.AppStoreRepository.List(context.Background(), &entity.AppStore{})
	if err != nil {
		return nil, err
	}

	res := make([]*vo.AppStore, len(appstores))
	for i, appstore := range appstores {
		res[i] = appstore.ToViewFromAppStore()
	}
	return res, nil
}

func (s *AppStoreServiceImpl) RefreshAppStore(name string) (*time.Time, error) {
	appstore, err := s.AppStoreRepository.Get(context.Background(), &entity.AppStore{
		Name: name,
	})
	if err != nil {
		return nil, err
	}

	err = s.HelmClient.UpdateRepo(appstore.Name)
	if err != nil {
		return nil, err
	}

	appstore.LastSync = time.Now()
	err = s.AppStoreRepository.Update(context.Background(), appstore)
	return &appstore.LastSync, err
}

func (s *AppStoreServiceImpl) RefreshAllAppStores() (*time.Time, error) {
	panic("todo")
}
